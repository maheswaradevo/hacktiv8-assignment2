package ping

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/maheswaradevo/hacktiv8-assignment2/internal/constant"
	"github.com/maheswaradevo/hacktiv8-assignment2/pkg/utils"
)

type PingHandler struct {
	r  *mux.Router
	ps PingService
}

func (pc *PingHandler) InitHandler() {
	routes := pc.r.PathPrefix(constant.PING_API_PATH).Subrouter()
	routes.HandleFunc("", pc.PingHandler).Methods(http.MethodGet)
}

func (pc PingHandler) PingHandler(w http.ResponseWriter, r *http.Request) {
	res := pc.ps.Ping(r.Context())
	utils.NewBaseResponse(http.StatusOK, res, nil, nil).SendResponse(&w)
}

func NewPingHandler(r *mux.Router, ps PingService) *PingHandler {
	return &PingHandler{
		r:  r,
		ps: ps,
	}
}
