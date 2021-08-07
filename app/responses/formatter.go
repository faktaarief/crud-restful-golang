package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func success(res http.ResponseWriter, statusCode int, data interface{}) {
	res.WriteHeader(statusCode)
	err := json.NewEncoder(res).Encode(data)
	if err != nil {
		fmt.Println("Response Formatter Error: ", err)
	}
}

func failed(res http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		success(res, statusCode, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}

	success(res, http.StatusBadRequest, nil)
}
