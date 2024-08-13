package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gitlab.local/dhamith93/devops-playground/app/internal/logger"
	"gitlab.local/dhamith93/devops-playground/app/internal/maths"
)

var MATHS_API_ENDPOINT string

type ClientServer struct {
	Port             string
	MathsApiEndpoint string
}

func (c *ClientServer) Run() {
	c.MathsApiEndpoint = os.Getenv("MATH_API_ENDPOINT")
	c.handleRequests()
}

func (c *ClientServer) handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/health", c.handleHealth)
	router.HandleFunc("/solve", c.handleCalculation)
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./frontend/")))

	server := http.Server{}
	server.Addr = c.Port
	server.Handler = handlers.CompressHandler(router)
	server.SetKeepAlivesEnabled(false)

	logger.Info("server started on port " + c.Port)
	log.Fatal(server.ListenAndServe())
}

func (c *ClientServer) handleCalculation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	expression := r.PostFormValue("expression")
	logger.Info(expression)

	req, err := http.NewRequest("POST", c.MathsApiEndpoint+"/solve", bytes.NewBuffer(
		[]byte(`{ "expression_str": "`+expression+`" }`),
	))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		logger.Error(err.Error())
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.Error("error sending request: " + err.Error())
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	var exp maths.Expression
	if err := decoder.Decode(&exp); err != nil {
		logger.Error("error decoding response: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
	}
	defer r.Body.Close()

	if exp.Error != nil {
		logger.Error("error with expression: " + exp.Error.Error())
		w.WriteHeader(http.StatusBadRequest)
	}
	json.NewEncoder(w).Encode(&exp)
}

func (c *ClientServer) handleHealth(w http.ResponseWriter, r *http.Request) {
	logger.Info(r.RemoteAddr + " checked health")
	w.WriteHeader(http.StatusOK)
}
