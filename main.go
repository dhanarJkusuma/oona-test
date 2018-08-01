package main

import (
	"oona-test/helper"
	"oona-test/models"
	"oona-test/usecase"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.POST("/api/v1/arithmatic", handleCalculator)
	router.POST("/api/v1/high-sum", handleCalculateHighestArray)
	router.POST("/api/v1/summaries", handleSummaryTemperature)
	http.ListenAndServe(":8080", router)
}

func handleCalculator(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	var input models.Payload
	err := helper.UnMarshall(req.Body, &input)
	if err != nil {
		errMessage := struct {
			message string
		}{
			"Invalid Request",
		}
		helper.RespondJSON(w, http.StatusBadRequest, errMessage)
		return
	}
	result := usecase.Calculate(input.Body)
	message := fmt.Sprintf(input.Body+" adalah %v", result)
	response := models.Payload{
		Body: message,
	}
	helper.RespondJSON(w, http.StatusOK, response)
	return
}

func handleCalculateHighestArray(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	var input models.PayloadArray
	err := helper.UnMarshall(req.Body, &input)
	if err != nil {
		errMessage := struct {
			message string
		}{
			"Invalid Request",
		}
		helper.RespondJSON(w, http.StatusBadRequest, errMessage)
		return
	}

	val := usecase.GetHighest(input.Body)
	message := fmt.Sprintf("Result: %v", val)
	response := models.Payload{
		Body: message,
	}
	helper.RespondJSON(w, http.StatusOK, response)
	return
}

func handleSummaryTemperature(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	var input []models.PayloadTemperature
	err := helper.UnMarshall(req.Body, &input)
	if err != nil {
		errMessage := struct {
			message string
		}{
			"Invalid Request",
		}
		helper.RespondJSON(w, http.StatusBadRequest, errMessage)
		return
	}
	result := usecase.GetSummary(input)
	helper.RespondJSON(w, http.StatusOK, result)
	return
}
