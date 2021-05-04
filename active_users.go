package main

import (
	"encoding/json"
	"net/http"
)

//show active users
func activeUsers(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	session, err := store.Get(request, "session")
	if err != nil {
		response.WriteHeader(http.StatusUnauthorized)
		response.Write([]byte(`{"message":"Login Required !!"}`))
		return
	}
	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(response, "Forbidden", http.StatusForbidden)
		return
	}

	var activeUser []ActiveUser
	err = db.Table("users").Select("id, name, email, username").Where("session = 1").Find(&activeUser).Error
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message":"` + err.Error() + `"}`))
		return
	}

	res := make(map[string]interface{})
	res["active_users"] = activeUser
	des, err := json.Marshal(res)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message":"` + err.Error() + `"}`))
		return
	}
	response.Write([]byte(des))
}
