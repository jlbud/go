package handler

import (
	"net/http"
	"github.com/jinzhu/gorm"
	"xiaoshi/app/model"
	"encoding/json"
	"xiaoshi/app/model/response"
)

func CreateFeedback(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("X-AccessToken")
	respFeedback := response.RespFeedback{}
	if checkToken(db, token) {
		feedback := model.Feedbacks{}
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&feedback); err != nil {
			respondError(w, http.StatusBadRequest, err.Error())
		}
		defer r.Body.Close()
		if err := db.Save(&feedback).Error; err != nil {
			respondError(w, http.StatusInternalServerError, err.Error())
			return
		}
		respFeedback.Data.Feedback = feedback
		respFeedback.Message = "pass"
		respFeedback.Success = "0"
		respondJSON(w, http.StatusCreated, respFeedback)
	} else {
		respFeedback.Message = "reject"
		respFeedback.Success = "1"
		respondJSON(w, http.StatusInternalServerError, respFeedback)
	}
}

func GetAllFeedback(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("X-AccessToken")
	respFeedbacks := response.RespFeedbacks{}
	if checkToken(db, token) {
		feedbacs := []model.Feedbacks{}
		db.Find(&feedbacs)

		respFeedbacks.Data.Feedbacks = feedbacs
		respFeedbacks.Message = "pass"
		respFeedbacks.Success = "0"
		respondJSON(w, http.StatusCreated, respFeedbacks)
	} else {
		respFeedbacks.Message = "reject"
		respFeedbacks.Success = "1"
		respondJSON(w, http.StatusInternalServerError, respFeedbacks)
	}
}

var statusCode int = 0

func checkToken(db *gorm.DB, token string) bool {
	//todo
	statusCode++
	if statusCode == 5 || statusCode == 10 || statusCode > 15 {
		return false
	}
	return true
}
