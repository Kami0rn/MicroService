package paymentHandler

import (
	"github.com/Kami0rn/MicroService/config"
	"github.com/Kami0rn/MicroService/module/payment/paymentUsecase"
)

type (
	PaymentQueueHandlerService interface{}

	paymentQueueHandler struct {
		cfg *config.Config
		paymentUsecase paymentUsecase.PaymentUsecaseService
	}
)

func NewPaymentQueueHandler(cfg *config.Config , paymentUsecase paymentUsecase.PaymentUsecaseService) PaymentQueueHandlerService {
	return &paymentQueueHandler{
		cfg: cfg,
		paymentUsecase: paymentUsecase,
	}
}