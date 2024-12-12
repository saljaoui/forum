import { InitialComment } from "./createcomment.js"
import { checklogin } from "./checklogin.js";
checklogin()
const urlParams = new URLSearchParams(window.location.search);
const cardData = urlParams.get("card_id");

async function fetchdata() {
    let fullname = document.querySelector(".full-name")
    let content = document.querySelector(".content")
    let time = document.querySelector(".time")
    let username = document.querySelector(".username")
    let cards = document.querySelectorAll("#likes")
    let is_liked = document.querySelector("#is_liked")
    let is_Dislikes = document.querySelector("#is_Dislikes")
    let comments = document.querySelector(".comments")
    let data = ""
    let path = window.location.pathname
    if (path !== "/comment") {
        return ""
    } else {
        const response = await fetch(`/api/card?id=${cardData}`, {
            method: "GET",
        })
        if (response.ok) {
            data = await response.json();
            fullname.textContent = data.lastName + " " + data.firstName
            content.textContent = data.content
            username.textContent = data.lastName
            is_liked.textContent = data.likes
            is_Dislikes.textContent = data.dislikes
            comments.textContent = data.comments
        } else if (!response.ok) {
            // Redirect to the backend error page with appropriate status and message
            window.location.href = `/error?code=${response.status}&msg=${encodeURIComponent("Failed to fetch card data.")}`;
            return;
        }
        cards.forEach(async (card) => {
            card.setAttribute("data-id_card", data.id)
        })
    }
}
fetchdata()
async function GetComments() {

    let path = window.location.pathname
    if (path !== "/comment") {
        return ""
    } else {
        const response = await fetch(`/api/comment?target_id=${cardData}`, {
            method: "GET",
        });
        if (response === null) {
            return ""
        }
        if (response.ok) {
            let datacomment = await response.json()
            let comments = document.querySelector(".allcomment")
            comments.innerHTML = ""
            InitialComment(datacomment, comments)
        } else if (!response.ok) {
            console.log('test here');

        }
        else {
            console.log("err");
        }
    }
}
GetComments()
export {
    fetchdata,
    GetComments
}