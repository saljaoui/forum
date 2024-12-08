import { InitialComment } from "./createcomment.js"
import { likes } from "./likescomment.js";
import { checklogin } from "./checklogin.js";
checklogin()
const urlParams = new URLSearchParams(window.location.search);
const cardData = urlParams.get("card_id");

export async function fetchdata() {
    let fullname = document.querySelector(".full-name")
    let content = document.querySelector(".content")
    let time = document.querySelector(".time")
    let username = document.querySelector(".username")
    let cards = document.querySelectorAll("#likes")
    let is_liked = document.querySelector("#is_liked")
    let is_Dislikes = document.querySelector("#is_Dislikes")
    let comments=document.querySelector(".comments")
     cards.forEach(async (card) => {
        let card_liked = card.getAttribute("data-like")
        console.log(card_liked);
        const response = await fetch(`/api/card?id=${cardData}`, {
            method: "GET",
        })
        if (response.ok) {
            likes(cards)
            const data = await response.json();
            fullname.textContent = data.lastName + " " + data.firstName
            content.textContent = data.content
            username.textContent = data.lastName
            is_liked.textContent = data.likes
            is_Dislikes.textContent = data.dislikes
            comments.textContent=data.comments
            card.setAttribute("data-id_card", data.id) 
        } else {
            const errorData = response.json();
            console.error("Error:", errorData);
            alert(`Error: ${errorData.message || "Request failed"}`);
        }
    })
}
fetchdata()
export async function GetComments() {
    const response = await fetch(`/api/comment?target_id=${cardData}`, {
        method: "GET",
    });
    if (response.ok) {
        let datacomment = await response.json()
        let comments = document.querySelector(".allcomment")
        comments.innerHTML = ""
        InitialComment(datacomment, comments)
    }
    else {
        console.log("err");
    }
}
GetComments()