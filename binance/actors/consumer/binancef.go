package consumer

import (
	"errors"
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/LeVanHieu0509/backend-go/binance/actors/symbol"
	"github.com/asynkron/protoactor-go/actor"
	"github.com/gorilla/websocket"
	"github.com/valyala/fastjson"
)

//22p: https://www.youtube.com/watch?v=16kAjS2lNAs

// when volum > x and marketPrice > x and OpenInterest > x -> alert me
// bot and algo makers (tự động hoá quá trình giao dịch) and webhook or subscribe WS (đăng kí dựa trên 1 điều kiện)

// Đây là phần cố định của URL WebSocket để kết nối với Binance.
const wsEndpoint = "wss://fstream.binance.com/stream?streams="

// Mảng các cặp giao dịch mà bạn muốn theo dõi. Trong trường hợp này là btcusdt và ethusdt.
var symbols = []string{"btcusdt", "ethusdt"}

/*
xây dựng một URL kết nối tới WebSocket của Binance để nhận các loại dữ liệu khác nhau
(giao dịch tổng hợp, giá thị trường và độ sâu thị trường) cho các cặp giao dịch btcusdt và ethusdt.

return: wss://fstream.binance.com/stream?streams=btcusdt@aggTrade/btcusdt@markPrice/btcusdt@depth/ethusdt@aggTrade/ethusdt@markPrice/ethusdt@depth
*/
func createWSEndPoint() string {
	//Mảng này sẽ lưu trữ các endpoint phụ cho từng loại dữ liệu của từng cặp giao dịch
	results := []string{}

	// Với mỗi cặp giao dịch, nó tạo ba endpoint phụ
	for _, symbol := range symbols {
		//fmt.Sprintf định dạng chuỗi theo mẫu "symbol@type".
		results = append(results, fmt.Sprintf("%s@aggTrade", symbol))  //Giao dịch tổng hợp (aggTrade)
		results = append(results, fmt.Sprintf("%s@markPrice", symbol)) //Giá thị trường (markPrice)
		results = append(results, fmt.Sprintf("%s@depth", symbol))     //Độ sâu thị trường (depth)
	}
	//nối các endpoint phụ trong results thành một chuỗi duy nhất, cách nhau bằng dấu gạch chéo (/).
	return fmt.Sprintf("%s%s", wsEndpoint, strings.Join(results, "/"))
}

// Tạo ra một hệ thống sử dụng WebSocket để kết nối tới Binance và nhận dữ liệu từ các stream, đồng thời sử dụng thư viện actor để quản lý các tác nhân (actors).
type Binancef struct {
	ws  *websocket.Conn //Kết nối WebSocket.
	ctx actor.Context   //Ngữ cảnh của actor.
}

// Tạo một producer cho actor Binancef.
func NewBinancef() actor.Producer {
	return func() actor.Actor {
		return &Binancef{}
	}
}

// Nhận thông điệp và thực hiện hành động tương ứng
func (a *Binancef) Receive(c actor.Context) {
	switch msg := c.Message().(type) {
	case *actor.Started:
		_ = msg
		a.ctx = c
		a.start()
	case *actor.Stopped:
		a.ws.Close()
	}

}

// Vòng lặp để đọc thông điệp từ WebSocket
func (a *Binancef) wsLoop() {
	/*
		Đọc thông điệp, nếu có lỗi, kiểm tra lỗi và xử lý phù hợp.
		Sử dụng fastjson để phân tích cú pháp thông điệp.
		Lấy thông tin stream và tách thành symbol và kind.
		Nếu kind là "markPrice", in ra symbol và kind.
	*/
	for {
		_, msg, err := a.ws.ReadMessage()
		if err != nil {
			if errors.Is(err, net.ErrClosed) {
				break
			}
			fmt.Println(err)
			continue
		}
		parser := fastjson.Parser{}
		v, err := parser.ParseBytes(msg)
		if err != nil {
			fmt.Println(err)
			continue
		}
		stream := v.GetStringBytes("stream")
		symbol, kind := splitStream(string(stream))

		if kind == "markPrice" {
			_ = symbol
			fmt.Println("symbol actor started:", symbol)
			// fmt.Printf("%s => %s\n", symbol, kind)
		}
		// data := v.Get("data")
		// fmt.Println(data)
	}
}

// Kết nối tới WebSocket endpoint của Binance
func (a *Binancef) start() {
	c, _, err := websocket.DefaultDialer.Dial(createWSEndPoint(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	a.ws = c

	for _, s := range symbols {
		pid, _ := a.ctx.SpawnNamed(actor.PropsFromProducer(symbol.New(s)), s)

		fmt.Println("spawning symbol child actor", pid)
	}
	// bắt đầu vòng lặp wsLoop trong một goroutine
	go a.wsLoop()
}

type Conditioner struct {
}

// Một actor đơn giản khác tên là Conditioner, có một hàm Receive để xử lý thông điệp Started
func NewConditioner() actor.Actor {
	return &Conditioner{}
}

func (a *Conditioner) Receive(c actor.Context) {
	switch msg := c.Message().(type) {
	case *actor.Started:
		_ = msg
	}
}

// Tách chuỗi stream thành hai phần: symbol và kind
func splitStream(stream string) (string, string) {
	parts := strings.Split(stream, "@")
	return parts[0], parts[1]
}
