package handlers

import (
	"admin-api/gen"
	"admin-api/usecases/students"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Login(logger *log.Logger, studentService students.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var loginModel gen.Login
		if err := json.NewDecoder(r.Body).Decode(&loginModel); err != nil {
			logger.Println(fmt.Sprintf("%v", err.Error()))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		tokenResponse, err := studentService.Authenticate(loginModel)
		if err != nil {
			logger.Println(fmt.Sprintf("%v", err.Error()))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		b, err := json.Marshal(tokenResponse)
		if err != nil {
			logger.Println(fmt.Sprintf("%v", err.Error()))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.Header().Add(ContentTypeHeader, JsonContentType)
		w.Write(b)

		return
	})
}
