package routers

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"qaa_api/api"
)

func AuthRoute(r chi.Router)  {
	r.Post("/login", login)
	r.Post("/register",register)
}
func login(writer http.ResponseWriter, request *http.Request) {
	type RequestBody struct {
		Email string `json:"email"`
		Password string `json:"password"`
	}
	type Response struct {
		Token string `json:"token"`
		Message string `json:"message"`
		IsOk bool `json:"is_ok"`
	}
	response := Response{}
	body := RequestBody{}
	if err := json.NewDecoder(request.Body).Decode(&body); err != nil  {
		log.Println(err)
	}
	isValid, token := api.Login(body.Email,body.Password)
	if isValid {
		response = Response{Token:token, IsOk: true, Message:"Login successfully"}
	} else {
		response = Response{
			Message: "Invalid username & password",
			IsOk: false,
		}
	}
	json.NewEncoder(writer).Encode(response)
}

func register(writer http.ResponseWriter, request *http.Request)  {
	type RequestBody struct {
		Email string `json:"email"`
		Password string `json:"password"`
		Name string `json:"name"`
	}
	type Response struct {
		Message string `json:"message"`
		IsSuccess bool `json:"is_ok"`
	}
	response := Response{}
	body := RequestBody{}
	if err := json.NewDecoder(request.Body).Decode(&body); err != nil  {
		log.Println(err)
	}
	if isSuccess := api.Register(body.Name, body.Email, body.Password); isSuccess {
		response = Response{
			Message:   "Register successfully",
			IsSuccess: true,
		}
	} else {
		response = Response{
			Message:   "Register fail. Email already existed",
			IsSuccess: false,
		}
	}
	json.NewEncoder(writer).Encode(response)
}
