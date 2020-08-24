package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ashishb26/rzpbool/auth"
	"github.com/ashishb26/rzpbool/controller"
	"github.com/ashishb26/rzpbool/models"
	"github.com/stretchr/testify/assert"
)

type UserToken struct {
	Token string `json:"token"`
}

func TestUserLogin(t *testing.T) {
	server := setServer()
	defer server.DB.Close()

	w := httptest.NewRecorder()
	userCred := []byte(`{"username": "root", "password": "password"}`)

	req, err := http.NewRequest("POST", "/user/login", bytes.NewBuffer(userCred))

	if err != nil {
		t.Fail()
	}
	req.Header.Set("Content-Type", "application/json")

	server.Router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	tok := &UserToken{}
	if err := json.Unmarshal([]byte(w.Body.String()), tok); err != nil {
		t.Fail()
	}

	err = auth.ValidateToken(tok.Token)
	if err != nil {
		t.Error("Token invalid")
	}

}

func TestPostBool(t *testing.T) {
	server := setServer()
	defer server.DB.Close()
	user := models.User{
		Username: "root",
		Password: "password",
	}
	tokenString, err := auth.GetToken(user)
	if err != nil {
		t.Error("Error getting authenticated")
	}
	var bearer = "Bearer: " + tokenString

	newBool := []byte(`{
		"value": true,
		"key": "Sample key"
	}`)

	req, err := http.NewRequest("POST", "/api/", bytes.NewBuffer(newBool))
	if err != nil {
		t.Error(err.Error())
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", bearer)
	w := httptest.NewRecorder()
	server.Router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var record models.BoolRecord
	if err := json.Unmarshal([]byte(w.Body.String()), &record); err != nil {
		t.Error(err.Error())
	}

	// data := url.Values{}
	// data.Add("id", record.ID)
	// req, err = http.NewRequest("GET", "/api/", strings.NewReader(data.Encode()))
	// req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	// req.Header.Set("Authorization", bearer)
	// if err != nil {
	// 	t.Error(err.Error())
	// }
	// w = httptest.NewRecorder()
	// server.Router.ServeHTTP(w, req)
	// assert.Equal(t, 200, w.Code)

	// var respRecord models.BoolRecord
	// if err := json.Unmarshal([]byte(w.Body.String()), &respRecord); err != nil {
	// 	t.Error(err.Error())
	// }

	// if record != respRecord {
	// 	t.Error(fmt.Sprintln("Expected record:", record, "Response record:", respRecord))
	// }
}

func setServer() *controller.Server {
	server := controller.NewServer()
	server.InitRoutes()
	server.Router.Run("8080")
	return server
}
