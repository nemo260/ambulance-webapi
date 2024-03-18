package ambulance_wl

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Nasledujúci kód je kópiou vygenerovaného a zakomentovaného kódu zo súboru api_ambulance_conditions.go
func (this *implAmbulanceConditionsAPI) GetConditions(ctx *gin.Context) {
	updateAmbulanceFunc(ctx, func(
		ctx *gin.Context,
		ambulance *Ambulance,
	) (updatedAmbulance *Ambulance, responseContent interface{}, status int) {
		result := ambulance.PredefinedConditions
		if result == nil {
			result = []Condition{}
		}
		return nil, result, http.StatusOK
	})
}

func (this *implAmbulanceConditionsAPI) AddConditions(ctx *gin.Context) {
	updateAmbulanceFunc(ctx, func(
		ctx *gin.Context,
		ambulance *Ambulance,
	) (updatedAmbulance *Ambulance, responseContent interface{}, status int) {
		var conditions []Condition
		err := ctx.BindJSON(&conditions)
		if err != nil {
			return nil, err, http.StatusBadRequest
		}
		ambulance.PredefinedConditions = conditions
		return ambulance, ambulance, http.StatusOK
	})
}

func (this *implAmbulanceConditionsAPI) DeleteConditions(ctx *gin.Context) {
	updateAmbulanceFunc(ctx, func(
		ctx *gin.Context,
		ambulance *Ambulance,
	) (updatedAmbulance *Ambulance, responseContent interface{}, status int) {
		ambulance.PredefinedConditions = nil
		return ambulance, nil, http.StatusOK
	})
}

func (this *implAmbulanceConditionsAPI) UpdateConditions(ctx *gin.Context) {
	updateAmbulanceFunc(ctx, func(
		ctx *gin.Context,
		ambulance *Ambulance,
	) (updatedAmbulance *Ambulance, responseContent interface{}, status int) {
		var conditions []Condition
		err := ctx.BindJSON(&conditions)
		if err != nil {
			return nil, err, http.StatusBadRequest
		}
		ambulance.PredefinedConditions = conditions
		return ambulance, ambulance, http.StatusOK
	})
}