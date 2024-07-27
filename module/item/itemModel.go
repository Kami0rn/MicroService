package item

type (
	CreateItemReq struct {
		Title    string  `json:"title" vladate:"required,max=64"`
		Price    float64 `json:"price" vlaidate:"required"`
		ImageUrl string  `json:"image_url" validate:"required,max=255"`
		Damage   int     `json:"damage" validate:"required"`
	}

	ItemShowCase struct {
		ItemId string `json:"item_id"`
	}
)
