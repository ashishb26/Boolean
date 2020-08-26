package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/ashishb26/rzpbool/auth"
	"github.com/ashishb26/rzpbool/controller"
	"github.com/ashishb26/rzpbool/models"
	"github.com/rs/xid"
	"github.com/stretchr/testify/assert"
)

type UserToken struct {
	Token string `json:"token"`
}

func setServer() *controller.Server {
	server := controller.NewServer()
	server.InitRoutes()
	server.Router.Run("8080")
	return server
}

func getBearerToken() (string, error) {
	user := models.User{
		Username: "root",
		Password: "password",
	}
	tokenString, err := auth.GetToken(user)
	if err != nil {
		return "", err
	}
	var bearer = "Bearer: " + tokenString
	return bearer, nil
}

func addSampleUser(s *controller.Server) *models.BoolRecord {
	record := models.BoolRecord{
		ID:        xid.New().String(),
		Value:     false,
		Key:       "Test Key",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	s.DB.Create(&record)
	return &record
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

func TestPostAndGetBool(t *testing.T) {
	server := setServer()
	defer server.DB.Close()
	bearer, err := getBearerToken()
	if err != nil {
		t.Error(err.Error())
	}
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

	req1, err := http.NewRequest("GET", "/api/"+record.ID, nil)
	req1.Header.Set("Authorization", bearer)
	if err != nil {
		t.Error(err.Error())
	}
	w1 := httptest.NewRecorder()
	server.Router.ServeHTTP(w1, req1)
	assert.Equal(t, 200, w1.Code)

	var respRecord models.BoolRecord
	if err := json.Unmarshal([]byte(w1.Body.String()), &respRecord); err != nil {
		t.Error(err.Error())
	}

	if record.ID != respRecord.ID || record.Value != respRecord.Value || record.Key != respRecord.Key {
		t.Error(fmt.Sprintln("Expected record:", record, "Response record:", respRecord))
	}
}

func TestUpdateValueOnly(t *testing.T) {
	server := setServer()
	bearer, err := getBearerToken()
	record := addSampleUser(server)

	updtBool := []byte(`{
		"value": true
	}`)
	req, err := http.NewRequest("PATCH", "/api/"+record.ID, bytes.NewBuffer(updtBool))
	if err != nil {
		t.Error(err.Error())
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", bearer)
	w := httptest.NewRecorder()
	server.Router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	var updtRecord models.BoolRecord

	if err := json.Unmarshal([]byte(w.Body.String()), &updtRecord); err != nil {
		t.Error(err.Error())
	}
	if updtRecord.ID != record.ID || updtRecord.Key != record.Key || updtRecord.Value != true {
		t.Error("Expected to change the value to `true`. But value is `false`")
	}
}

func TestUpdateKeyOnly(t *testing.T) {
	server := setServer()
	bearer, err := getBearerToken()
	record := addSampleUser(server)

	updtBool := []byte(`{
		"key": "Modified key"
	}`)
	req, err := http.NewRequest("PATCH", "/api/"+record.ID, bytes.NewBuffer(updtBool))
	if err != nil {
		t.Error(err.Error())
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", bearer)
	w := httptest.NewRecorder()
	server.Router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	var updtRecord models.BoolRecord

	if err := json.Unmarshal([]byte(w.Body.String()), &updtRecord); err != nil {
		t.Error(err.Error())
	}
	if updtRecord.ID != record.ID || updtRecord.Key != "Modified key" || updtRecord.Value != record.Value {
		t.Error("Expected to change the key to `Modified key`. But value is `Test key`")
	}
}

func TestUpdateKeyAndValue(t *testing.T) {
	server := setServer()
	bearer, err := getBearerToken()
	record := addSampleUser(server)

	updtBool := []byte(`{
		"value": true,
		"key": "Modified key"
	}`)
	req, err := http.NewRequest("PATCH", "/api/"+record.ID, bytes.NewBuffer(updtBool))
	if err != nil {
		t.Error(err.Error())
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", bearer)
	w := httptest.NewRecorder()
	server.Router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	var updtRecord models.BoolRecord

	if err := json.Unmarshal([]byte(w.Body.String()), &updtRecord); err != nil {
		t.Error(err.Error())
	}
	if updtRecord.ID != record.ID || updtRecord.Key == record.Key || updtRecord.Value == record.Value {
		t.Error("Expected to change the (value, key) to (true,`Modified key`). But (value, key) is (false,`Test key`).")
	}
}

func TestDeleteRecord(t *testing.T) {
	server := setServer()
	bearer, err := getBearerToken()
	record := addSampleUser(server)

	req, err := http.NewRequest("DELETE", "/api/"+record.ID, nil)
	if err != nil {
		t.Error(err.Error())
	}
	req.Header.Set("Authorization", bearer)
	w := httptest.NewRecorder()
	server.Router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	req1, err := http.NewRequest("GET", "/api/"+record.ID, nil)
	if err != nil {
		t.Error(err.Error())
	}
	req1.Header.Set("Authorization", bearer)
	w1 := httptest.NewRecorder()
	server.Router.ServeHTTP(w1, req1)
	if w1.Code == 200 {
		t.Error("Expected to delete the record. But failed to delete")
	}
}

func TestAccessWithoutAuth(t *testing.T) {
	server := setServer()
	newBool := []byte(`{
		"value": true,
		"key": "Sample key"
	}`)

	req, err := http.NewRequest("POST", "/api/", bytes.NewBuffer(newBool))
	if err != nil {
		t.Error(err.Error())
	}
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	server.Router.ServeHTTP(w, req)
	assert.Equal(t, 401, w.Code)
}

func TestGetRecordThatDoesntExist(t *testing.T) {
	server := setServer()
	bearer, err := getBearerToken()
	req, err := http.NewRequest("GET", "/api/"+"1824fgjewyerq814", nil)
	req.Header.Set("Authorization", bearer)
	if err != nil {
		t.Error(err.Error())
	}
	w := httptest.NewRecorder()
	server.Router.ServeHTTP(w, req)
	if w.Code != http.StatusBadRequest {
		t.Error("Expected to return an error. Actual: no error")
	}
}
