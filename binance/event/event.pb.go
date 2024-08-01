package event

import (
	"github.com/gogo/protobuf/proto"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Trade struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
	Marker bool    `json:"marker"`
	Qty    float64 `json:"qty"`
}

type MarketPrice struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
}

// Reset resets the MarketPrice struct.
func (m *MarketPrice) Reset() {
	*m = MarketPrice{}
}

// String returns the string representation of the MarketPrice struct.
func (m *MarketPrice) String() string {
	return proto.CompactTextString(m)
}

// ProtoMessage is a required method for proto.Message interface.
func (*MarketPrice) ProtoMessage() {}
