package routers

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"qaa_api/api"
	"qaa_api/constant"
	"qaa_api/helper"
)

func VoteRoute(r chi.Router)  {
	r.With(helper.AuthMiddleware).Post("/",createVote)
}

func createVote(writer http.ResponseWriter, request *http.Request)  {
	accountId := helper.GetAccountIDLogged(request)
	type RequestBody struct {
		Id string `json:"id"`
		VoteType constant.VoteType `json:"vote_type"`
	}
	type Response struct {
		Message string `json:"message"`
		IsOk bool `json:"is_ok"`
	}
	var response Response
	var body RequestBody
	if err := json.NewDecoder(request.Body).Decode(&body); err != nil  {
		log.Println(err)
	}
	isSuccess, _ := api.CreateVote(accountId,helper.ParseObjectID(body.Id),body.VoteType)
	if isSuccess {
		response = Response{IsOk: true, Message:"Created vote"}
	} else {
		response = Response{
			Message: "The question or answer not existed or you already vote for them",
			IsOk: false,
		}
	}
	json.NewEncoder(writer).Encode(response)
}