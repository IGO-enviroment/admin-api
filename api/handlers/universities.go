package handlers

import (
	"admin-api/config"
	"admin-api/gen"
	"admin-api/usecases/universities"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func AddStudents(logger *log.Logger, setting *config.Settings, universityService universities.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims, err := GetTokenClaims(r, setting)
		if err != nil {
			logger.Println(fmt.Sprintf("%v", err.Error()))
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if !claims.IsUniversity {
			logger.Println("request from not university")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		var requestBody gen.AddStudent
		if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
			logger.Println(fmt.Sprintf("%v", err.Error()))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		response, err := universityService.AddStudents(requestBody)
		if err != nil {
			logger.Println(fmt.Sprintf("%v", err.Error()))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		b, err := json.Marshal(response)
		if err != nil {
			logger.Println(fmt.Sprintf("%v", err.Error()))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		_, err = w.Write(b)
		if err != nil {
			logger.Println(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	})
}
