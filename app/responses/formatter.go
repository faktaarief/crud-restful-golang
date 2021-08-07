package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Success(res http.ResponseWriter, statusCode int, data interface{}) {
	res.WriteHeader(statusCode)
	err := json.NewEncoder(res).Encode(data)
	if err != nil {
		fmt.Println("Response Formatter Error: ", err)
	}
}

func Failed(res http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		Success(res, statusCode, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}

	Success(res, http.StatusBadRequest, nil)
}
