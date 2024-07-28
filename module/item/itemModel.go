package item

import "github.com/Kami0rn/MicroService/module/models"

type (
	CreateItemReq struct {
		Title    string  `json:"title" vladate:"required,max=64"`
		Price    float64 `json:"price" vlaidate:"required"`
		ImageUrl string  `json:"image_url" validate:"required,max=255"`
		Damage   int     `json:"damage" validate:"required"`
	}

	ItemShowCase struct {
		ItemId   string  `json:"item_id"`
		Title    string  `json:"title"`
		Price    float64 `json:"price"`
		ImageUrl string  `json:"image_url"`
		Damage   int     `json:"damage"`
	}

	ItemSearchReq struct {
		Title string `json:"title"`
		models.PaginateReq
	}

	ItemUpdateReq struct {
		Title    string  `json:"title" validate:"required,max=64"`
		Price    float64 `json:"price" validate:"required"`
		ImageUrl string  `json:"image_url" validate:"required,max=255"`
		Damage   int     `json:"damage" validate:"required"`
	}

	EnableOrDisableItemReq struct {
		UsageStatus bool `json:"usage_status"`

	}
)
