package handler

import (
	"encoding/json"
	"mymodule/entity"
	"mymodule/utils"
	"net/http"
)

type zipperHandler struct {
}

func NewZipperHandler() *zipperHandler {
	return &zipperHandler{}
}

func (h *zipperHandler) ZipperByPathHandler(w http.ResponseWriter, r *http.Request) {
	// Put content header in the first line of handler
	w.Header().Set("Content-Type", "application/json")

	// Get header from request to determine the client is production or staging
	// Throw error if apiKey is not exists or valid.
	apiKey := r.Header.Get("Authorization")

	workDir, err := utils.WorkdirRouting(apiKey)
	if err != nil {
		response := utils.GenerateResponse(nil, 500, err.Error())

		utils.NewResponse(w, response)
		return
	}

	// Parse request
	var requestData entity.DocumentRequest

	json.NewDecoder(r.Body).Decode(&requestData)

	result, err := utils.CreateZip(&requestData.OutputName, requestData.Source, workDir, requestData.OutputDir)
	if err != nil {
		response := utils.GenerateResponse(result, 500, err.Error())

		utils.NewResponse(w, response)
		return
	}

	response := utils.GenerateResponse(result, 200, "Success")

	utils.NewResponse(w, response)
}
