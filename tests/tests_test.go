package tests

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"origin-auth/auth"
	"testing"

	"github.com/udonetsm/help/models"
)

func TestAuthorize(t *testing.T) {
	body, err := json.Marshal(models.Auth{Password: "testpassword", Email: "test@email"})
	if err != nil {
		log.Fatal(err)
	}
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8383/authorize/?", bytes.NewReader(body))
	w := httptest.NewRecorder()
	auth.Authorize(w, req)
	res := w.Result()
	defer res.Body.Close()
	ans := models.ResponseAuth{}
	json.NewDecoder(w.Body).Decode(&ans)
	t.Log(ans)
}
