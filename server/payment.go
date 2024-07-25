package server

import (
	"github.com/Kami0rn/MicroService/module/payment/paymentRepository"
	"github.com/Kami0rn/MicroService/module/payment/paymentUsecase"
	"github.com/Kami0rn/MicroService/module/payment/paymentHandler"
)

func (s *server) paymentService() {
	paymentRepository := paymentRepository.NewPaymentRepository(s.db)
	paymentUsecase := paymentUsecase.NewPaymentUsecase(paymentRepository)
	paymentHttpHandler := paymentHandler.NewPaymentHttpHandler(s.cfg, paymentUsecase)
	paymentQueue := paymentHandler.NewPaymentQueueHandler(s.cfg, paymentUsecase)

	_ = paymentHttpHandler
	_ = paymentQueue

	payment := s.app.Group("/payment_v1")

	//Health Check
	payment.GET("" , s.healthCheckService)

}