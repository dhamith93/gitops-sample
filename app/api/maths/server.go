package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gitlab.local/dhamith93/devops-playground/app/internal/logger"
	"gitlab.local/dhamith93/devops-playground/app/internal/maths"
	"gitlab.local/dhamith93/devops-playground/app/internal/req"
)

func Run(port string) {
	handleRequests(port)
}

func handleRequests(port string) {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/health", handleHealth)
	router.HandleFunc("/solve", solve)

	server := http.Server{}
	server.Addr = port
	server.Handler = handlers.CompressHandler(router)
	server.SetKeepAlivesEnabled(false)

	logger.Info("maths API server started on port " + port)
	log.Fatal(server.ListenAndServe())
}

func solve(w http.ResponseWriter, r *http.Request) {
	var expressionReq req.ExpressionToSolve
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&expressionReq); err != nil {
		logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	logger.Info(r.RemoteAddr + " " + expressionReq.ExpressionStr)

	expression := maths.Expression{}
	expression.Parse(expressionReq.ExpressionStr)
	expression.Solve()

	if expression.Error != nil {
		logger.Error(expression.Error.Error())
	}

	logger.Info(expression.Result)
	json.NewEncoder(w).Encode(&expression)
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	logger.Info(r.RemoteAddr + " checked health")
	w.WriteHeader(http.StatusOK)
}
