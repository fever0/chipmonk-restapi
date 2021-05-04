package main

import (
	"encoding/json"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

//login user
func login(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	session, err := store.Get(request, "session")
	if err != nil {
		response.WriteHeader(http.StatusUnauthorized)
		response.Write([]byte(`{"message":"Login Required !!"}`))
		return
	}

	var user, dbUser User

	err = json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{"message":"` + err.Error() + `"}`))
		return
	}

	err = db.Table("users").Select("id, name, email,password").Where("username = ?", user.Username).Find(&dbUser).Error
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message":"` + err.Error() + `"}`))
		return
	}

	userPass := []byte(user.Password)
	dbPass := []byte(dbUser.Password)

	passErr := bcrypt.CompareHashAndPassword(dbPass, userPass)
	if passErr != nil {
		log.Println(passErr)
		response.WriteHeader(http.StatusForbidden)
		response.Write([]byte(`{"response":"Wrong Password!"}`))
		return
	}

	err = db.Model(&user).Where("username = ?", user.Username).Update("session", 1).Error
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{"message":"` + err.Error() + `"}`))
		return
	}

	// Set user as authenticated
	session.Values["authenticated"] = true
	session.Save(request, response)

	response.WriteHeader(http.StatusOK)

	res := make(map[string]interface{})
	res["response"] = "Login Successful!"
	res["user_id"] = dbUser.ID
	res["user_name"] = dbUser.Name
	res["user_email"] = dbUser.Email

	des, err := json.Marshal(res)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message":"` + err.Error() + `"}`))
		return
	}
	response.Write([]byte(des))

}
