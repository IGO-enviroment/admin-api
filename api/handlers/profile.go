package handlers

import (
	"admin-api/config"
	"admin-api/usecases/organization"
	"admin-api/usecases/students"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func GetProfile(logger *log.Logger, settings *config.Settings, studentService students.Service, orfService organization.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims, err := GetTokenClaims(r, settings)
		if err != nil {
			logger.Println(fmt.Sprintf("%v", err.Error()))
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		var res interface{}

		if claims.IsStudent {
			res, err = studentService.GetStudentProfile(claims.Email)
			if err != nil {
				logger.Println(fmt.Sprintf("%v", err.Error()))
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		}

		if claims.IsOrganization {
			res, err = orfService.GetOrgProfile(claims.Email)
			if err != nil {
				logger.Println(fmt.Sprintf("%v", err.Error()))
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		}

		w.Header().Set("Access-Control-Allow-Origin", "*")
		b, err := json.Marshal(res)
		if err != nil {
			logger.Println(fmt.Sprintf("%v", err.Error()))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.Header().Add(ContentTypeHeader, JsonContentType)
		_, err = w.Write(b)
		if err != nil {
			logger.Println(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	})

}
