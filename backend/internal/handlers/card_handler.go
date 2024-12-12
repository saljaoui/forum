package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"forum-project/backend/internal/repository/cards"
)

func GetCard_handler(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		HandleError(res, req, "Status Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	id, err := strconv.Atoi(req.FormValue("id"))
	if err != nil {
		HandleError(res, req, "Status Bad Request", http.StatusBadRequest)
		return
	}
	card := cards.GetOneCard(id)
	if card.Id == -1 {
		HandleError(res, req, "Status Bad Request", http.StatusBadRequest)
		return
	}
	json.NewEncoder(res).Encode(card)
}
