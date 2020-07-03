package routers

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"qaa_api/api"
	"qaa_api/helper"
	"qaa_api/models"
)
func QuestionRoute(r chi.Router)  {
	r.With(helper.AuthMiddleware).Post("/",createQuestion)
	r.Get("/", viewQuestion)
}

func createQuestion(writer http.ResponseWriter, request *http.Request)  {
	accountId := helper.GetAccountIDLogged(request)
	type RequestBody struct {
		Name string `json:"name"`
		Content string `json:"content"`
	}
	type Response struct {
		Message string `json:"message"`
		IsOk bool `json:"is_ok"`
		QuestionId string `json:"question_id"`
	}
	var response Response
	var body RequestBody
	if err := json.NewDecoder(request.Body).Decode(&body); err != nil  {
		log.Println(err)
	}
	isSuccess, questionId := api.CreateQuestion(body.Name,body.Content,true,accountId)
	if isSuccess {
		response = Response{IsOk: true, Message:"Created question", QuestionId:questionId}
	} else {
		response = Response{
			Message: "Internal server error",
			IsOk: false,
		}
	}
	json.NewEncoder(writer).Encode(response)
}

func viewQuestion(writer http.ResponseWriter, request *http.Request) {
	type Response struct {
		Questions []models.Question `json:"questions"`
	}
	var response Response
	size := request.URL.Query().Get("size")
	page := request.URL.Query().Get("page")
	isOpen := request.URL.Query().Get("is_open")
	questions := api.ViewQuestion(helper.ParseInt(page),helper.ParseInt(size),helper.ParseBool(isOpen))
	response = Response{Questions:questions}
	json.NewEncoder(writer).Encode(response)
}