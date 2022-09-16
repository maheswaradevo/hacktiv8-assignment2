package order

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

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
	routes.HandleFunc("", o.createNewOrder()).Methods(http.MethodPost)
	routes.HandleFunc("", o.viewAllOrders()).Methods(http.MethodGet)
	routes.HandleFunc("/{order_id}", o.deleteOrderByID()).Methods(http.MethodDelete)
	routes.HandleFunc("/{order_id}", o.updateOrderByID()).Methods(http.MethodPut)
	routes.HandleFunc("/person/{order_id}", o.getPersonOrders()).Methods(http.MethodGet)
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
// @Success 201 {object} dto.OrderResponse
// @Router /orders [post]
func (o *OrderHandler) createNewOrder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := dto.CreateOrderRequest{}
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			log.Printf("[createNewOrder] failed to parse JSON data, err => %+v", err)
			panic(err)
		}
		res, err := o.os.CreateNewOrder(r.Context(), &data)
		if err != nil {
			log.Printf("[createNewOrder] failed to create a new order, err => %v", err)
			panic(err)
		}
		utils.NewBaseResponse(http.StatusCreated, "SUCCESS", nil, res).SendResponse(&w)
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
func (o *OrderHandler) viewAllOrders() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := o.os.ViewAllOrders(r.Context())
		if err != nil {
			log.Printf("[viewAllOrders] failed to get all the orders, err => %v", err)
			panic(err)
		}
		utils.NewBaseResponse(http.StatusOK, "SUCCESS", nil, res).SendResponse(&w)
	}
}
func (o *OrderHandler) deleteOrderByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		routeVar := mux.Vars(r)
		idVar := routeVar["order_id"]
		idConv, _ := strconv.ParseUint(idVar, 10, 64)

		res, err := o.os.DeleteOrderByID(r.Context(), idConv)
		if err != nil {
			log.Printf("[deleteOrderByID] failed to delete the order by id, err => %v, id => %v", err, idConv)
			panic(err)
		}
		utils.NewBaseResponse(http.StatusAccepted, "SUCCESS", nil, res).SendResponse(&w)
	}
}

func (o *OrderHandler) updateOrderByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		routeVar := mux.Vars(r)
		idVar := routeVar["order_id"]
		idConv, _ := strconv.ParseUint(idVar, 10, 64)

		data := new(dto.UpdateOrderRequest)
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			log.Printf("[updateOrderByID] failed to parse JSON data, err => %+v", err)
			panic(err)
		}

		res, err := o.os.UpdateOrderByID(r.Context(), idConv, data)
		if err != nil {
			log.Printf("[updateOrderByID] failed to update the order by id, err => %v, id => %v", err, idConv)
			panic(err)
		}
		utils.NewBaseResponse(http.StatusCreated, "SUCCESS", nil, res).SendResponse(&w)
	}
}

func (o *OrderHandler) getPersonOrders() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		routeVar := mux.Vars(r)
		idVar := routeVar["order_id"]
		idConv, _ := strconv.ParseUint(idVar, 10, 64)

		res, err := o.os.PersonOrders(r.Context(), idConv)
		if err != nil {
			panic(err)
		}
		utils.NewBaseResponse(http.StatusOK, "SUCCESS", nil, res).SendResponse(&w)
	}
}
