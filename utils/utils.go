package utils

import (
	"log"
	"net/http"
	"strconv"
)

func ConvertStringToFloat64(data string) float64 {
	convertedData, err := strconv.ParseFloat(data, 64)
	if err != nil {
		log.Println("Erro na convesão para float64:", err)
	}

	return convertedData
}

func ConvertStringToInt(data string) int {
	convertedData, err := strconv.Atoi(data)
	if err != nil {
		log.Println("Erro na convesão para int:", err)
	}

	return convertedData
}

func IsPost(r *http.Request) bool {
	return r.Method == "POST"
}
