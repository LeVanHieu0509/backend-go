package server

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/labstack/echo"
)

const (
	MarketETH   Market    = "ETH"
	MarketOrder OrderType = "MARKET"
	LimitOrder  OrderType = "LIMIT"

	exchangePrivateKey = ""
)

type (
	OrderType string
	Market    string

	PlaceOrderRequest struct {
		UserID int64
		Type   OrderType //limit or market
		Bid    bool
		Size   float64
		Price  float64
		Market Market
	}

	Order struct {
		UserID    int64
		ID        int64
		Price     float64
		Size      float64
		Bid       bool
		Timestamp int64
	}

	OrderbookData struct {
		TotalBidVolume float64
		TotalAskVolume float64
		Asks           []*Order
		Bids           []*Order
	}
	MatchedOrder struct {
		UserID int64
		Price  float64
		Size   float64
		ID     int64
	}
	APIError struct {
		Error string
	}
)

// 30p https://www.youtube.com/watch?v=n0SqZmEBWpk&list=PL0xRBLFXXsP6sG4IG1-gOfOCz5zAhsx5T&index=1
func StartServer() {
	e := echo.New()
	e.HTTPErrorHandler = httpErrorHandler

	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}

	ex, err := NewExchange(exchangePrivateKey, client)
	if err != nil {
		log.Fatal(err)
	}

	buyerAddressString := ""
	buyerBalance, err := client.BalanceAt(context.Background(), common.HexToAddress(buyerAddressString), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("buyer:", buyerBalance)

	sellerAddressStr := ""
	sellerBalance, err := client.BalanceAt(context.Background(), common.HexToAddress(sellerAddressStr), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("seller:", sellerBalance)

	pkStr8 := ""
	user8 := NewUser(pkStr8, 8)
	ex.Users(user8.ID) = user8

	joinAddress := ""
	joinBalance, err := client.BalanceAt(context.Background(), common.HexToAddress(joinAddress))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("seller:", joinBalance)

	e.POST("/order", ex.handlePlaceOrder)
	e.GET("/trades/:market", ex.handleGetTraders)
	e.GET("/order/:userID", ex.handleGetOrders)
	e.GET("/book/:market", ex.handleGetBook)
	e.GET("/book/:market/bid", ex.handleGetBestBid)
	e.GET("/book/:market/ask", ex.handleGetBestAsk)

	e.DELETE("/order/:id", ex.handleGetBestAsk)
	e.Start(":8000")

}

type User struct {
	ID         int64
	PrivateKey *ecdsa.PrivateKey
}

func NewUser(privKey string, id int64) *User {
	pk, err := crypto.HexToECDSA(privKey)
	if err != nil {
		panic(err)
	}

	return &User{
		ID:         id,
		PrivateKey: pk,
	}
}

func httpErrorHandler(err error, c echo.Context) {
	fmt.Println(err)
}

type Exchange struct {
	Client *ethclient.Client
	mu     sync.RWMutex
	Users  map[int64]*User
	//Orders maps a user his order
	Orders     map[int64][]*orderbook.Order
	PrivateKey *ecdsa.PrivateKey
	orderbooks map[Market]*orderbook.OrderbookData
}

func NewExchange(privateKey string, client *ethclient.Client) (*Exchange, error) {
	orderbooks := make(map[Market]*orderbook.OrderBook)
	orderbooks[MarketETH] = orderbook.NewOrderbook()

	pk, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, err
	}

	return &Exchange{
		Client:     client,
		Users:      make(map[int64]*User),
		Orders:     make(map[int64][]*orderbook.Order),
		PrivateKey: pk,
		orderbooks: orderbooks,
	}
}

type GetOrdersResponse struct {
	Asks []Order
	Bids []Order
}

func (ex Exchange) registerUser(pk string, userId int64) {
	user := NewUser(pk, userId)
	ex.Users[userId] = user
}

func (ex Exchange) handleGetTraders(c echo.Context) error {
	market := Market(c.Param("market"))
	ob, ok := ex.orderbooks[market]

	if !ok {
		return c.JSON(http.StatusBadRequest, APIError{Error: "Orderbook not found"})
	}

	return c.JSON(http.StatusOK, ob.trades)
}

func (ex Exchange) handleGetOrders(c echo.Context) error {
	userIDStr := c.Param("userID")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return err
	}

	ex.mu.RLock()
	orderbookOrders := ex.Orders[int64(userID)]
	ordersResp := &GetOrdersResponse{
		Asks: []Order{},
		Bids: []Order{},
	}

	for i := 0; i < len(orderbookOrders); i++ {
		if orderbookOrders[i].Limit == nil {
			continue
		}

		order := Order{
			ID:        orderbookOrders[i].ID,
			UserID:    orderbookOrders[i].UserID,
			Price:     orderbookOrders[i].Price,
			Size:      orderbookOrders[i].Size,
			Timestamp: orderbookOrders[i].Timestamp,
			Bid:       orderbookOrders[i].Bid,
		}

		if order.Bid {
			ordersResp.Bids = append(ordersResp.Bids, order)

		} else {
			ordersResp.Asks = append(ordersResp.Asks, order)
		}
	}
	ex.mu.RUnlock()

	return c.JSON(http.StatusOK, ordersResp)
}

func (ex Exchange) handleGetBook(c echo.Context) error {
	market := Market(c.Param("market"))
	ob, ok := ex.orderbooks[market]

	if !ok {
		return c.JSON(http.StatusBadRequest, APIError{Error: "Market not found"})
	}

	orderbookData := OrderbookData{
		TotalBidVolume: ob.BidTotalVolume(),
		TotalAskVolume: ob.AskTotalVolume(),
		Asks:           []*Order{},
		Bids:           []*Order{},
	}

	for _, limit := range ob.Asks() {
		for _, order := range limit.Orders {
			o := Order{
				ID:        order.ID,
				UserID:    order.UserID,
				Price:     order.Price,
				Size:      order.Size,
				Timestamp: order.Timestamp,
				Bid:       order.Bid,
			}
			orderbookData.Asks = append(orderbookData.Asks, &o)
		}
	}

	for _, limit := range ob.Bids() {
		for _, order := range limit.Orders {
			o := Order{
				ID:        order.ID,
				UserID:    order.UserID,
				Price:     order.Price,
				Size:      order.Size,
				Timestamp: order.Timestamp,
				Bid:       order.Bid,
			}
			orderbookData.Bids = append(orderbookData.Bids, &o)
		}
	}
	return c.JSON(http.StatusOK, orderbookData)
}

type PriceResponse struct {
	Price float64
}

func (ex Exchange) handleGetBestBid(c echo.Context) error {
	market := Market(c.Param("market"))
	ob := ex.orderbooks[market]
	pr := PriceResponse{
		Price: 0.0,
	}

	if len(ob.Bids()) == 0 {
		return fmt.Errorf("The bids are empty")
	}

	pr.Price = ob.Bids()[0].Price
	return c.JSON(http.StatusOK, pr)
}

func (ex Exchange) handleGetBestAsk(c echo.Context) error {
	market := Market(c.Param("market"))
	ob := ex.orderbooks[market]
	pr := PriceResponse{
		Price: 0.0,
	}

	if len(ob.Asks()) == 0 {
		return fmt.Errorf("The Asks are empty")
	}

	pr.Price = ob.Bids()[0].Price

	return c.JSON(http.StatusOK, pr)
}
