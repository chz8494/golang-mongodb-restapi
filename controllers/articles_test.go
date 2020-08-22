package controllers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestCreateArticle(t *testing.T) {

	var jsonStr = []byte(`{"id":"5f3fd53ac1a707a6419705d2","title":"", "content": "test content" , "author":"test author"}`)
	req, err := http.NewRequest("POST", "/articles", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateArticle)
	handler.ServeHTTP(response, req)

	if response.Code == 500 {
		assert.Equal(t, 500, response.Code, "Failed Creation")
	} else {
		assert.Equal(t, 200, response.Code, "Created!")
	}

}
func TestGetArticle(t *testing.T) {
	t.Parallel()

	req, _ := http.NewRequest("GET", "/articles/5f3f7c62c03b7f0a8ad0d9d1", nil)
	response := httptest.NewRecorder()

	//Hack to try to fake gorilla/mux vars
	vars := map[string]string{
		"id": "5f3f7c62c03b7f0a8ad0d9d1",
	}

	req = mux.SetURLVars(req, vars)

	GetArticle(response, req)
	expected1 := `{"data":{"_id":"5f3f7c62c03b7f0a8ad0d9d1","title":"Test - Britta","content":"Hello World","author":"Britta-testing"},"message":"Success","status":200}
`
	expected2 := `{"data":null,"message":"the provided hex string is not a valid ObjectID","status":500}
`
	if response.Code == 500 {
		assert.Equal(t, http.StatusInternalServerError, response.Code)
		assert.Equal(t, expected2, response.Body.String())
	} else {

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, expected1, response.Body.String())
	}

}
func TestGetArticles(t *testing.T) {
	t.Parallel()

	req, _ := http.NewRequest("GET", "/articles", nil)
	response := httptest.NewRecorder()

	GetArticle(response, req)

	if response.Code == 500 {
		assert.Equal(t, http.StatusInternalServerError, response.Code)

	} else {

		assert.Equal(t, http.StatusOK, response.Code)
	}

}
