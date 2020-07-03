package main

import (
    "fmt"
    "github.com/go-chi/chi"
    "net/http"
    "os"
    "qaa_api/constant"
    "qaa_api/routers"
)

func main() {
    fmt.Println("Q&A API Service")
    fmt.Println(os.Getenv(constant.DB_URL))
    r := chi.NewRouter()
    r.Get("/", func(writer http.ResponseWriter, request *http.Request) {
        writer.Write([]byte("Q&A API Service"))
    })
    r.Route("/auth",routers.AuthRoute)
    r.Route("/question",routers.QuestionRoute)
    r.Route("/answer", routers.AnswerRoute)
    r.Route("/vote",routers.VoteRoute)
    http.ListenAndServe(":3000",r)
}
