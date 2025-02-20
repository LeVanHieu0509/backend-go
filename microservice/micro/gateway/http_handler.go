package main

import (
	"errors"
	"net/http"

	"github.com/LeVanHieu0509/backend-go/microservice/micro/common"
	pb "github.com/LeVanHieu0509/backend-go/microservice/micro/common/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Handler là một struct có một trường client, là client gRPC của dịch vụ
type Handler struct {
	// Nó giúp kết nối và gọi các phương thức gRPC từ backend
	client pb.OrderServiceClient
}

// tạo một handler mới với client gRPC và trả về con trỏ đến Handler.
func NewHandler(client pb.OrderServiceClient) *Handler {
	return &Handler{
		client,
	}
}

// Đăng ký một route HTTP POST với URL mẫu
func (h *Handler) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/customers/{customerID}/orders", h.HandleCreateOrder)
}

func (h *Handler) HandleCreateOrder(w http.ResponseWriter, r *http.Request) {
	// Đọc customerID từ URL.
	customerID := r.PathValue("customerID")
	var items []*pb.ItemsWithQuantity

	// Đọc dữ liệu JSON từ body của yêu cầu HTTP và giải mã vào biến items (danh sách các mặt hàng và số lượng).
	if err := common.ReadJSON(r, &items); err != nil {
		common.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := validateItems(items); err != nil {
		common.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	// gọi phương thức gRPC CreateOrder để tạo đơn hàng với các dữ liệu: customerID và items.
	o, err := h.client.CreateOrder(r.Context(), &pb.CreateOrderRequest{
		CustomerID: customerID,
		Items:      items,
	})

	// Đoạn mã sử dụng status.Convert(err) để chuyển đổi lỗi từ gRPC thành status.Status
	rStatus := status.Convert(err)

	if rStatus != nil {
		if rStatus.Code() != codes.InvalidArgument {

			common.WriteError(w, http.StatusInternalServerError, rStatus.Message())
			return
		}
	}

	// trả về phản hồi thành công với mã trạng thái HTTP 200 và dữ liệu đơn hàng vừa tạo.
	common.WriteJSON(w, http.StatusOK, o)
}

func validateItems(items []*pb.ItemsWithQuantity) error {
	if len(items) == 0 {
		return common.ErrNoItems
	}

	for _, i := range items {
		if i.ID == "" {
			return errors.New("item ID is required")
		}
		if i.Quantity <= 0 {
			return errors.New("items must have a valid quantity")
		}
	}

	return nil
}
