package main

import (
	"net/http"
)

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

func (app *application) Login(w http.ResponseWriter, r *http.Request) {
	type credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var creds credentials
	var payload jsonResponse

	err := app.readJSON(w, r, &creds)
	if err != nil {
		app.errorLog.Println(err)
		payload.Error = true
		payload.Message = "JSON неправильний, або його немає"
		_ = app.writeJSON(w, http.StatusBadRequest, payload)
	}

	// TODO авторизувати
	app.infoLog.Println(creds.Email, creds.Password)

	// надіслати успішну відповідь
	payload.Error = false
	payload.Message = "Ви ввійшли"

	err = app.writeJSON(w, http.StatusOK, payload)
	if err != nil {
		app.errorLog.Println(err)
	}
}
