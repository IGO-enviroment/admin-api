package handlers

import (
	"admin-api/config"
	"admin-api/usecases/students"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func GetStudentProfile(logger *log.Logger, settings *config.Settings, studentService students.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims, err := GetTokenClaims(r, settings)
		if err != nil {
			logger.Println(fmt.Sprintf("%v", err.Error()))
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		profile, err := studentService.GetStudentProfile(claims.Email)
		if err != nil {
			logger.Println(fmt.Sprintf("%v", err.Error()))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		b, err := json.Marshal(profile)
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
