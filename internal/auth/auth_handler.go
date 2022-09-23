package auth

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/maheswaradevo/hacktiv8-assignment2/internal/constant"
	"github.com/maheswaradevo/hacktiv8-assignment2/internal/dto"
	"github.com/maheswaradevo/hacktiv8-assignment2/pkg/utils"
)

type authHandler struct {
	r  *mux.Router
	as AuthService
}

func ProvideAuthHandler(r *mux.Router, as AuthService) *authHandler {
	return &authHandler{r: r, as: as}
}

func (a *authHandler) InitHandler() {
	route := a.r.PathPrefix(constant.AUTH_USER_API_PATH).Subrouter()

	route.HandleFunc("/register", a.newUser()).Methods(http.MethodPost)
	route.HandleFunc("/login", a.loginUser()).Methods(http.MethodPost)
}

func (a *authHandler) newUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := dto.UserRegistrationRequest{}
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			log.Printf("[newUser] failed to parse JSON data, err => %v", err)
		}
		res, err := a.as.NewUser(r.Context(), &data)
		if err != nil {
			log.Printf("[newUser] failed to store new user, err => %v", err)
		}
		utils.NewBaseResponse(http.StatusCreated, "SUCCESS", nil, res).SendResponse(&w)
	}
}

func (a *authHandler) loginUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &dto.UserLoginRequest{}
		err := json.NewDecoder(r.Body).Decode(data)

		if err != nil {
			log.Printf("[loginUser] failed to parse the JSON data, err => %v", err)
			return
		}

		res, err := a.as.LoginUser(r.Context(), data)
		if err != nil {
			log.Printf("[loginUser] failed to login, err => %v", err)
			return
		}
		utils.NewBaseResponse(http.StatusOK, "SUCCESS", nil, res).SendResponse(&w)
	}
}
