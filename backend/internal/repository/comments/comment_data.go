package comment

import "forum-project/backend/internal/database"


func insertComment(card_id,target_id int) int {
    query := "INSERT INTO comment(card_id,target_id) VALUES(?,?);"
    resl,_ := database.Exec(query,card_id,target_id)
    id,err := resl.LastInsertId()
    if err != nil {
        return -1
    }
    return int(id)
}