package itemHandler

import "github.com/Kami0rn/MicroService/module/item/itemRepository"

type (
	ItemHandlerService interface{}

	itemHandler struct {
		itemRepository itemRepository.ItemRepositoryService
	}
)