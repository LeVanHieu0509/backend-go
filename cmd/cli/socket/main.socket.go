package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"golang.org/x/net/websocket"
)

// How To Build A Chat And Data Feed With WebSockets In Golang?
type Server struct {
	conns map[*websocket.Conn]bool
}

func NewServer() *Server {
	return &Server{
		conns: make(map[*websocket.Conn]bool),
	}
}

func (s *Server) handleWSOrderbook(ws *websocket.Conn) {
	fmt.Println("New incoming connection from client to orderbook feed:", ws.RemoteAddr())

	for {
		payload := fmt.Sprintf("Orderbook data => %d\n", time.Now().UnixNano()+1)
		ws.Write([]byte(payload))

		time.Sleep(time.Second * 2)
	}
}
func (s *Server) handleWS(ws *websocket.Conn) {
	fmt.Println("New incoming connection from client:", ws.RemoteAddr())

	s.conns[ws] = true

	s.readLoop(ws)
}

func (s *Server) readLoop(ws *websocket.Conn) {
	buf := make([]byte, 1024)

	for {
		n, err := ws.Read(buf)
		if err != nil {

			if err == io.EOF {
				delete(s.conns, ws)
				break
			}

			fmt.Println("Read Error:", err)
			continue
		}

		msg := buf[:n]

		s.broadcast(msg)

		fmt.Println("msg:", string(msg))
		// ws.Write([]byte("thank you for the msg!!!"))
	}
}

func (s *Server) broadcast(b []byte) {
	for ws := range s.conns {
		go func(ws *websocket.Conn) {
			if _, err := ws.Write(b); err != nil {
				fmt.Println("Write error:", err)
			}
		}(ws)
	}
}
func main() {
	server := NewServer()
	http.Handle("/ws", websocket.Handler(server.handleWS))
	http.Handle("/orderbookfeed", websocket.Handler(server.handleWSOrderbook))
	http.ListenAndServe(":3000", nil)
}

/*
	1. test 1 connect
	socket client connect: let socket = new WebSocket("ws://localhost:3000/ws")
	socket.onmessage = (event) => {console.log("received from the server ", event.data)}
	socket.send("Hello Iam hieu")
*/

/*
	2. test multi connect

*/
