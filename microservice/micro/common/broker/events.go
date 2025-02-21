package broker

const (
	OrderCreatedEvent = "order.created"
	OrderPaidEvent    = "order.paid"
)

const (
	OrderCreatedRoutingKey = "order-created-routing-key"
	OrderPaidRoutingKey    = "order-paid-routing-key"
)
