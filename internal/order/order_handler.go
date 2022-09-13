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

// CreateNewOrder godoc
// @Summary Create a New Order
// @Description Create a new order from request body as a JSON
// @Tags orders
// @Accept json
// @Produce json
// @Param data body dto.CreateOrderRequest true "Create order"
// @Success 200 {object} dto.OrderResponse
// @Router /orders [post]
func (o *OrderHandler) CreateNewOrder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := dto.CreateOrderRequest{}
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			log.Printf("[CreateNewOrder] failed to parse JSON data, err => %+v", err)
			panic(utils.NewErrorResponse(http.StatusBadRequest, "BAD_REQUEST", utils.NewErrorResponseValue("request body", "invalid json format")))
		}
		res, err := o.os.CreateNewOrder(r.Context(), &data)
		if err != nil {
			panic(
				utils.NewErrorResponse(http.StatusInternalServerError,
					"INTERNAL_SERVER_ERROR",
					utils.NewErrorResponseValue("internal", "server error")),
			)
		}
		utils.NewBaseResponse(http.StatusOK, "SUCCESS", nil, res).SendResponse(&w)
	}
}

// ViewAllOrders godoc
// @Summary View all orders
// @Description View all orders and return a JSON
// @Tags orders
// @Accept json
// @Produce json
// @Success 200 {object} dto.OrderDetails
// @Router /orders [get]
func (o *OrderHandler) ViewAllOrders() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := o.os.ViewAllOrders(r.Context())
		if err != nil {
			panic(
				utils.NewErrorResponse(http.StatusInternalServerError,
					"Internal Server Error",
					utils.NewErrorResponseValue("Internal Server Error", err.Error())))
		}
		utils.NewBaseResponse(http.StatusOK, "SUCCESS", nil, res).SendResponse(&w)
	}
}
