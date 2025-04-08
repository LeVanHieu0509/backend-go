package model

// Get ticketItems returns
type TicketItemsOutput struct {
	TicketId       int    `json:"id"`
	TicketName     string `json:"name"`
	StockAvailable int    `json:"stock_available"`
	StockInitial   int    `json:"stock_initial"`
}
