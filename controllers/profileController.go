package controllers

import (
	"encoding/json"
	"go-rest-api-basic/models"
	u "go-rest-api-basic/utils"
	"net/http"
)

var CreateProfile = func(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("user").(uint)
	profile := &models.Profile{}

	err := json.NewDecoder(r.Body).Decode(profile)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	profile.UserId = userId
	resp := profile.Create()
	u.Respond(w, resp)
}

var GetProfile = func(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("user").(uint)
	profile := &models.Profile{}

	data := profile.Get(uint(userId))
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var UpdateProfile = func(w http.ResponseWriter, r *http.Request) {
	profile := &models.Profile{}

	err := json.NewDecoder(r.Body).Decode(profile)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	data := profile.Update()
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
