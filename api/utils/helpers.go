package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func ToJson(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	CheckErr(err)
}

func BodyParser(r *http.Request) ([]byte, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func RespondWithError(response http.ResponseWriter, message string, statusCode int) {
	ToJson(response, map[string]string{"error": message}, statusCode)
}

func RespondWithJSON(response http.ResponseWriter, data interface{}, statusCode int) {
	switch statusCode {
	case http.StatusNoContent:
		response.WriteHeader(statusCode)
	default:
		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(statusCode)
		if data != nil {
			err := json.NewEncoder(response).Encode(data)
			if err != nil {
				return
			}
		}
	}
}
