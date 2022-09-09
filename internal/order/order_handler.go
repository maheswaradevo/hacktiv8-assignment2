package order

import (
	"github.com/gorilla/mux"
)

type OrderHandler struct {
	r  *mux.Router
	os OrderService
}

func (o *OrderHandler) InitHandler() {
	//routes := o.r.PathPrefix(constant.ORDER_API_PATH).Subrouter()

	//Order

}

func ProvideOrderHandler(r *mux.Router, os OrderService) *OrderHandler {
	return &OrderHandler{
		r:  r,
		os: os,
	}
}
