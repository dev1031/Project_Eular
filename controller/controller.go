package controller

import (
	"encoding/json"
	"net/http"
)

func GetLandingPage(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}
