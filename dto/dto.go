package dto

type DeliveryRequest struct {
	Origin      []string `json:"origin" validate:"required"`
	Destination []string `json:"destination" validate:"required"`
}

type DeliveryResponse struct {
	Id       uint64 `json:"id"`
	Distance int    `json:"distance"`
	Status   string `json:"status"`
}

type TakeOrder struct {
	Status string `json:"status"`
}
