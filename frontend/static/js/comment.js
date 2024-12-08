import { InitialComment } from "./createcomment.js"
const urlParams = new URLSearchParams(window.location.search);
const cardData = urlParams.get("card_id");
async function fetchdata() {
    let fullname = document.querySelector(".full-name")
    let content =document.querySelector(".content")
    let time=document.querySelector(".time")
    let  username=document.querySelector(".username")
    let cards=document.querySelectorAll("#likes")
    let is_liked=document.querySelector("#is_liked")
    let is_Dislikes=document.querySelector("#is_Dislikes")
    console.log(is_Dislikes,is_liked);
    
    cards.forEach(async (card)=>{
        let cardid = card.getAttribute("data-id_card");
        console.log(cardid);
        let card_liked=card.getAttribute("data-liked")
        console.log(card_liked);
        
        const response = await fetch(`/api/card?id=${cardData}`, {
            method: "GET",
        })
        if (response.ok) {
            const data = await response.json();
            fullname.textContent=data.lastName+" "+data.firstName
            content.textContent=data.content
            username.textContent=data.lastName
            is_liked.textContent=data.likes
            is_Dislikes.textContent=data.dislikes
console.log(data.dislikes);

            card.setAttribute("data-id_card", data.id)
           // cardid.dataset.id_card = data.id;
            console.log("data is : ", data);
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
        // console.log("datacomment : ",datacomment);
        // console.log("rani ndkhel ")
        let comments = document.querySelector(".allcomment")
        comments.innerHTML = ""
        InitialComment(datacomment, comments)
        //     datacomment.map(ele => {
        //    })
    }
    else {
        console.log("err");
    }
}
GetComments()