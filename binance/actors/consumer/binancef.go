package consumer

import (
	"errors"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"

	"github.com/LeVanHieu0509/backend-go/binance/actors/symbol"
	"github.com/LeVanHieu0509/backend-go/binance/event"
	"github.com/asynkron/protoactor-go/actor"
	"github.com/gorilla/websocket"
	"github.com/valyala/fastjson"
)

/*
Summary: Mã này tạo một hệ thống dựa trên các actor để theo dõi các sự kiện từ Binance qua WebSocket.
Nó khởi tạo kết nối WebSocket, tạo các actor con cho từng symbol,
và gửi các sự kiện MarketPrice đến các actor tương ứng khi nhận được thông điệp từ WebSocket.
Việc sử dụng thư viện actor giúp mã dễ dàng mở rộng và quản lý tốt hơn các tác vụ song song.
*/

//22p: https://www.youtube.com/watch?v=16kAjS2lNAs

// 1. Add coin
// 2. % theo 1 tiếng tính từ giá cuối cùng tới giá 1 tiếng.
// 3. check theo 3-4 nến giảm 5% thì ping về telegram
// 4. Ping lên giá nào có lệnh bán lớn ở mức giá nào
//

// when volum > x and marketPrice > x and OpenInterest > x -> alert me
// bot and algo makers (tự động hoá quá trình giao dịch) and webhook or subscribe WS (đăng kí dựa trên 1 điều kiện)

// Đây là phần cố định của URL WebSocket để kết nối với Binance.
const wsEndpoint = "wss://fstream.binance.com/stream?streams="

// Mảng các cặp giao dịch mà bạn muốn theo dõi. Trong trường hợp này là btcusdt và ethusdt.
var symbols = []string{"wusdt"}

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
	ws      *websocket.Conn       //Kết nối WebSocket.
	ctx     actor.Context         //Ngữ cảnh của actor.
	symbols map[string]*actor.PID //bản đồ symbols để lưu các PID của các symbol.
}

// Tạo một producer cho actor Binancef.
func NewBinancef() actor.Producer {
	// actor Binancef có vai trò chính trong việc kết nối và nhận thông điệp từ WebSocket
	return func() actor.Actor {
		// Binancef actor quản lý kết nối WebSocket và khởi tạo các actor con cho từng symbol.
		return &Binancef{
			symbols: make(map[string]*actor.PID),
		}
	}
}

// Hàm nhận thông điệp và thực hiện hành động tương ứng. Khi actor bắt đầu, nó thiết lập ngữ cảnh và bắt đầu kết nối WebSocket.
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

// Vòng lặp để đọc thông điệp từ WebSocket.
// Nó phân tích cú pháp thông điệp, tách thông tin stream thành symbol và kind, và gửi sự kiện MarketPrice đến PID tương ứng.
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
		data := v.Get("data")
		// fmt.Println("data", data)
		// Kiểm tra nếu loại thông điệp nhận được là aggTrade.
		if kind == "aggTrade" {
			// convert data difficult => data ez
			if pid, ok := a.symbols[symbol]; ok {
				price, _ := strconv.ParseFloat(string(data.GetStringBytes("p")), 64)
				maker, _ := strconv.ParseBool(string(data.GetStringBytes("m")))
				qty, _ := strconv.ParseFloat(string(data.GetStringBytes("q")), 64)

				// Các actor con tương ứng với từng symbol sẽ nhận và xử lý các sự kiện MarketPrice cụ thể.
				a.ctx.Send(pid, &event.Trade{
					Symbol: symbol,
					Price:  price,
					Maker:  maker,
					Qty:    qty,
				})
			}
		}
		if kind == "markPrice" {
			_ = symbol
			fmt.Println("symbol actor started:", symbol)
			// fmt.Printf("%s => %s\n", symbol, kind)

			if pid, ok := a.symbols[symbol]; ok {
				price, _ := strconv.ParseFloat(string(data.GetStringBytes("p")), 64)

				// Các actor con tương ứng với từng symbol sẽ nhận và xử lý các sự kiện MarketPrice cụ thể.
				a.ctx.Send(pid, &event.MarketPrice{
					Symbol: symbol,
					Price:  price,
				})
			}
		}
		// data := v.Get("data")
		// fmt.Println(data)
	}
}

// Kết nối tới WebSocket endpoint của Binance và khởi tạo các actor con cho từng symbol.
func (a *Binancef) start() {
	c, _, err := websocket.DefaultDialer.Dial(createWSEndPoint(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}

	a.ws = c

	for _, s := range symbols {
		pid, _ := a.ctx.SpawnNamed(actor.PropsFromProducer(symbol.New(s)), s)
		a.symbols[s] = pid
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

/*
	Actor model cung cấp một cách tiếp cận mạnh mẽ và linh hoạt để quản lý các tác vụ đồng thời và không đồng bộ trong hệ thống của bạn.
	Nó giúp bạn quản lý trạng thái một cách an toàn, xử lý các tác vụ một cách hiệu quả, và dễ dàng mở rộng hệ thống.
	Việc sử dụng actor trong hệ thống của bạn giúp mã dễ đọc, dễ bảo trì và có hiệu năng cao.
*/
