package main

import (
	"fmt"

	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"

	"forum-project/backend/internal/database"
	"forum-project/backend/internal/handlers"
	like "forum-project/backend/internal/repository/likes"
	//"forum-project/backend/internal/repository/cards"
	//comment "forum-project/backend/internal/repository/comments"
)


func main() {
	myLike := like.NewLike(1,2)
	myerr := myLike.SetIsLike(0)
	if myerr != nil {
		fmt.Println(myerr.Error())
		return
	}
	fmt.Println(myLike.GetIsLike())
	Err := database.InitDB()
	if Err != nil {
		fmt.Println(Err)
	}  //18
	fmt.Printf("1 & 7 : %v \n", 18 & 16) //     1 2 4 8 16 32 64 128
										//    1 1 1 0 0  0  0   0
										//    1 1 0 0 0  0  0   0
	/*MyComment := comment.NewComment(2,"omar",1)
	MyComment.Add()
	myCard := cards.GetCard(MyComment.Card_Id)
	myCard.PrintInfo()*/
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.Middleware)
	mux.HandleFunc("/api/register", handlers.HandleRegister)
	mux.HandleFunc("/api/login", handlers.LoginHandle)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../../frontend/static"))))
	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../../frontend/templates/login.html")
	})
	fmt.Println("Server running at :3333")
	fmt.Println("http://localhost:3333")
	err := http.ListenAndServe(":3333", mux)
	if err != nil {
		log.Fatal("ListenAndServe Error: ", err)
	}
}
