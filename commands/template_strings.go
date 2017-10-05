package commands

import (
	"encoding/json"

	pb "github.com/sonm-io/core/proto"
)

var slotTemplate = &pb.Slot{
	StartTime:      &pb.Timestamp{Seconds: 10000},
	EndTime:        &pb.Timestamp{Seconds: 20000},
	BuyerRating:    10,
	SupplierRating: 20,
	Geo:            &pb.Geo{Country: "RU", City: "Moscow", Lat: 55.7494055, Lon: 37.6242552},
	Resources: &pb.Resources{
		CpuCores:      1,
		RamBytes:      100 * 1024 * 1024,
		GpuCount:      0,
		Storage:       100,
		NetTrafficIn:  100,
		NetTrafficOut: 100,
		NetworkType:   pb.NetworkType_OUTBOUND,
	},
}

var orderTemplate = &pb.Order{
	ByuerID:    "1",
	SupplierID: "2",
	Price:      100,
	OrderType:  pb.OrderType_ASK,
	Slot:       slotTemplate,
}

func getSlotTemplate() string {
	b, _ := json.MarshalIndent(slotTemplate, "", "\t")
	return string(b)
}

func getOrderTemplate() string {
	b, _ := json.MarshalIndent(orderTemplate, "", "\t")
	return string(b)
}
