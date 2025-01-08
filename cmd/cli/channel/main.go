package main

import (
	"fmt"
	"time"
)

type Course struct {
	Title string
	Price int
}

type Message struct {
	OrderId string
	Title   string
	Price   int
}

/*
- Càng ngày càng nhiều người học nhiều ngôn ngữ lập trình khác nhau (NodeJS/Java/Go là vũ khí)
- Channel: Giống như redis pub/sub
*/

func main() {
	// 1. add channel to
	ch := make(chan Course)

	// 2. create goroutine``
	go func() {
		//create data
		course := Course{Title: "Tips Go", Price: 30}

		//send data to channel
		ch <- course
	}()

	c := <-ch //receive data from channel

	fmt.Printf("Received course: title: %s, price %d\n", c.Title, c.Price) //Received course: title: Tips Go, price 30

	// ---------//
	//1. create channel
	orderChannel := make(chan Message)
	fmt.Println("Start pub sub...")

	orders := []Message{
		{OrderId: "1", Title: "Tips Go", Price: 30},
		{OrderId: "2", Title: "Tips NodeJS", Price: 31},
		{OrderId: "3", Title: "Tips Java", Price: 32},
		{OrderId: "4", Title: "Tips Go 4", Price: 33},
	}

	// 2. create goroutine``
	go publishCourse(orderChannel, orders)
	go subscribeCourse(orderChannel, "Anonystick User")

	time.Sleep(5 * time.Second)
	fmt.Println("Finished pub sub...")
}

func publishCourse(ch chan<- Message, orders []Message) {

	for _, order := range orders {
		fmt.Printf("Publishing: title: %s, price %d\n", order.Title, order.Price)
		ch <- order
		time.Sleep(1 * time.Second)
	}

	close(ch)
}

func subscribeCourse(ch <-chan Message, userName string) {
	for mes := range ch {
		fmt.Printf("User %s received: orderId: %s, title: %s, price %d\n", userName, mes.OrderId, mes.Title, mes.Price)

		time.Sleep(1 * time.Second)
	}
}
