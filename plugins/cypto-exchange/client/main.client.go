package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/LeVanHieu0509/backend-go/plugins/cypto-exchange/server"
)

const Endpoint = "http://localhost:3000"

type PlaceOrderParams struct {
	UserID int64
	Bid    bool

	//Price only needed for placing LIMIT order.
	Price float64
	Size  float64
}

type Client struct {
	*http.Client
}

func NewClient() *Client {
	return &Client{
		Client: http.DefaultClient,
	}
}

// func (c *Client) GetTrades(market string) ([]*orderbook.Trade, error) {
// 	e := fmt.Sprintf("%/trades/%s", Endpoint, market)
// 	req, err := http.NewRequest(http.MethodGet, e, nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	resp, err := c.Do(req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	trades := []*orderbook.Trade{}
// 	if err := json.NewDecoder(resp.Body).Decode(&trades); err != nil {
// 		return nil, err
// 	}

// 	return trades, nil

// }

func (c *Client) GetOrders(userID int64) (*server.GetOrdersResponse, error) {
	e := fmt.Sprintf("%/order/%s", Endpoint, userID)
	req, err := http.NewRequest(http.MethodGet, e, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	orders := server.GetOrdersResponse{}

	if err := json.NewDecoder(resp.Body).Decode(&orders); err != nil {
		return nil, err
	}

	return &orders, nil
}

func (c *Client) GetBestAsk() (float64, error) {
	e := fmt.Sprintf("%/book/ETH/ask", Endpoint)
	req, err := http.NewRequest(http.MethodGet, e, nil)
	if err != nil {
		return 0, err
	}

	resp, err := c.Do(req)
	if err != nil {
		return 0, err
	}

	priceResp := server.PriceResponse{}

	if err := json.NewDecoder(resp.Body).Decode(&priceResp); err != nil {
		return 0, err
	}

	return priceResp.Price, err
}

func (c *Client) GetBestBid() (float64, error) {
	e := fmt.Sprintf("%/book/ETH/bid", Endpoint)
	req, err := http.NewRequest(http.MethodGet, e, nil)
	if err != nil {
		return 0, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return 0, err
	}

	priceResp := server.PriceResponse{}

	if err := json.NewDecoder(resp.Body).Decode(&priceResp); err != nil {
		return 0, err
	}

	return priceResp.Price, err

}
