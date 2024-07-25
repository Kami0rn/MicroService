package paymentHandler

import (
	"github.com/Kami0rn/MicroService/config"
	"github.com/Kami0rn/MicroService/module/payment/paymentUsecase"
)

type (
	PaymentHttpHandlerService interface {
	}

	paymentHttpHandler struct {
		cfg *config.Config
		paymentUsecase paymentUsecase.PaymentUsecaseService
	}
)

func NewPaymentHttpHandler (cfg *config.Config , paymentUsecase paymentUsecase.PaymentUsecaseService) PaymentHttpHandlerService {
	return &paymentHttpHandler {
		cfg : cfg ,
		paymentUsecase : paymentUsecase ,
	}
}