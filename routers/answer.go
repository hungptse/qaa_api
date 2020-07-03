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

func AnswerRoute(r chi.Router)  {
	r.With(helper.AuthMiddleware).Post("/", createAnswer)
	r.Get("/", viewAnswerByQuestion)
}

func createAnswer(writer http.ResponseWriter, request *http.Request)  {
	accountId := helper.GetAccountIDLogged(request)
	type RequestBody struct {
		QuestionId string `json:"question_id"`
		Content string `json:"content"`
	}
	type Response struct {
		Message string `json:"message"`
		AnswerId string `json:"answer_id"`
	}
	var response Response
	var body RequestBody
	if err := json.NewDecoder(request.Body).Decode(&body); err != nil  {
		log.Println(err)
	}
	isSuccess, answerId := api.CreateAnswer(helper.ParseObjectID(body.QuestionId), accountId,body.Content)
	if isSuccess {
		response = Response{Message: "Created answer", AnswerId:answerId}
	} else {
		response = Response{
			Message: "Can't create answer at this time. Request can be incompliance parameter",
			AnswerId: "",
		}
	}
	json.NewEncoder(writer).Encode(response)
}

func viewAnswerByQuestion(writer http.ResponseWriter, request *http.Request)  {
	type RequestBody struct {
		QuestionId string `json:"question_id"`
	}
	type Response struct {
		Answers []models.AnswerWithVote `json:"answers"`
	}
	var response Response
	var body RequestBody
	if err := json.NewDecoder(request.Body).Decode(&body); err != nil  {
		log.Println(err)
	}
	response.Answers = api.ViewAnswerByQuestion(helper.ParseObjectID(body.QuestionId))
	json.NewEncoder(writer).Encode(response)
}