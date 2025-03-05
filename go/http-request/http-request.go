package httprequest

import (
	"cipher/decode"
	"encoding/json"
	"net/http"
)

type MessageRequest struct {
	Content string `json:"text"`
	Key     string `json:"key"`
	Cod     string `json:"cod"`
	Cipher  string `json: "cipher"`
}

type MessageRespone struct {
	Result string `json:"result"`
}

func HandleMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Only Post method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var msg MessageRequest

	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	repsone := MessageRespone{
		Result: "message:" + decode.Decode(
			msg.Content, msg.Key, msg.Cipher, msg.Cod),
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(repsone)
}
