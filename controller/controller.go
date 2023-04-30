package controller

import (
	"./Models"
	"./service"
	"bytes"
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
	"time"
)

func ConvertImage(w http.ResponseWriter, req *http.Request) {
	body, err := getRequestBody(req)
	if err != nil {
		sendErrorMessage(w, http.StatusBadRequest, err)
		return
	}
	resp, err := service.ConvertImage(body)
	if err != nil {
		sendErrorMessage(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Add("Content-Type", "image/"+resp.ContentType)
	w.Header().Set("Content-Disposition", "filename="+resp.FileName)
	w.WriteHeader(http.StatusOK)
	http.ServeContent(w, req, resp.FileName, time.Now(), bytes.NewReader(resp.File))
	return
}

func getRequestBody(req *http.Request) (*Models.RequestBody, error) {
	v := validator.New()
	requestBody := &Models.RequestBody{}
	file, handler, err := req.FormFile("file")
	if err != nil {
		return nil, err
	}
	requestBody.File = file
	requestBody.FileName = handler.Filename
	requestBody.Format = req.FormValue("format")
	requestBody.Quality, _ = strconv.Atoi(req.FormValue("quality"))
	requestBody.Height, _ = strconv.Atoi(req.FormValue("height"))
	requestBody.Width, _ = strconv.Atoi(req.FormValue("width"))

	err = v.Struct(requestBody)
	fmt.Println(err)
	if err != nil {
		return nil, err
	}

	return requestBody, nil
}

func sendErrorMessage(w http.ResponseWriter, code int, error error) {
	w.WriteHeader(code)
	fmt.Fprintf(w, error.Error())
}
