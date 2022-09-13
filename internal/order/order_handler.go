package order

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/maheswaradevo/hacktiv8-assignment2/internal/constant"
	"github.com/maheswaradevo/hacktiv8-assignment2/internal/dto"
	"github.com/maheswaradevo/hacktiv8-assignment2/pkg/utils"
)

type OrderHandler struct {
	r  *mux.Router
	os OrderService
}

func (o *OrderHandler) InitHandler() {
	routes := o.r.PathPrefix(constant.ORDER_API_PATH).Subrouter()

	//Order
	routes.HandleFunc("", o.CreateNewOrder()).Methods(http.MethodPost)
	routes.HandleFunc("", o.ViewAllOrders()).Methods(http.MethodGet)
}

func ProvideOrderHandler(r *mux.Router, os OrderService) *OrderHandler {
	return &OrderHandler{
		r:  r,
		os: os,
	}
}

func (o *OrderHandler) CreateNewOrder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := dto.CreateOrderRequest{}
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			log.Printf("[StoreNewProduct] failed to parse JSON data, err => %+v", err)
			panic(err)
		}
		res, err := o.os.CreateNewOrder(r.Context(), &data)
		if err != nil {
			utils.NewErrorResponse(http.StatusInternalServerError, "Internal Server Error", utils.NewErrorResponseValue("Internal Server Error", err.Error())).SendResponse(&w)
		}
		utils.NewBaseResponse(http.StatusOK, "Success Insert New Order", nil, res).SendResponse(&w)
	}
}

func (o *OrderHandler) ViewAllOrders() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := o.os.ViewAllOrders(r.Context())
		if err != nil {
			panic(utils.NewErrorResponse(http.StatusInternalServerError, "Internal Server Error", utils.NewErrorResponseValue("Internal Server Error", err.Error())))
		}
		utils.NewBaseResponse(http.StatusOK, "Success", nil, res).SendResponse(&w)
	}
}
