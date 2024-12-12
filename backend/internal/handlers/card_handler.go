package handlers


import(
	"encoding/json"
	"forum-project/backend/internal/repository/cards"
	"net/http"
	"strconv"
)

func GetCard_handler(res http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.FormValue("id"))
	if err != nil {
		HandleError(res, "Status Bad Request", http.StatusBadRequest)
		return
	}
	card := cards.GetOneCard(id)
	if card.Id == -1  {
		HandleError(res, "Status Not Found", http.StatusNotFound)
		return
	}
	json.NewEncoder(res).Encode(card)
}