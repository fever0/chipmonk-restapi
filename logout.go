package main

import (
	"encoding/json"
	"net/http"
)

//logout user
func logout(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	session, err := store.Get(request, "session")
	if err != nil {
		response.WriteHeader(http.StatusUnauthorized)
		response.Write([]byte(`{"message":"Login Required !!"}`))
		return
	}
	var user User

	err = json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{"message":"` + err.Error() + `"}`))
		return
	}

	err = db.Model(&user).Where("username = ?", user.Username).Update("session", 0).Error
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{"message":"` + err.Error() + `"}`))
		return
	}
	// Revoke users authentication
	session.Values["authenticated"] = false
	session.Save(request, response)
	response.Write([]byte(`{"message":"Successfully logged out"}`))

}
