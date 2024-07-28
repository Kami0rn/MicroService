package payment

type (
	ItemServiceReq struct {
		Item []*ItemserviceReqDatum `json:"items" validate:"required"`
	}

	ItemserviceReqDatum struct {
		ItemId string `json:"item_id" validate:"required,max=64"`
		Price float64 `json:"price"`
	}
)