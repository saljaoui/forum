import {GetComments,NewComment}  from "./comment.js";
import { NewPost } from "./forum.js";
import { likes } from "./likes.js";
import { cards } from "./card.js";
const urlParams = new URLSearchParams(window.location.search);
const cardData = urlParams.get("card_id");
console.log("cardId is : ",cardData);

export async function LoadPage(cardid) {
    let main = document.querySelector("main")
    main.innerHTML = ""
    await fetchCard(cardid)
    NewComment()
    await GetComments(cardid)
}

export async function fetchCard(cardid) {
    const response = await fetch(`/api/card?id=${cardid}`, {
        method: "GET",
    })
    if (response.ok) {
        const data = await response.json();
        console.log("data is : ",data);
        NewPost(data,"main")
    } else {
        const errorData = response.json();
        console.error("Error:", errorData);
        alert(`Error: ${errorData.message || "Request failed"}`);
    }
}
await LoadPage(+cardData)
//await fetchCard(+cardData)
const creategategory = document.querySelector(".postReply")
const creatPostPopup = document.getElementById('creatPost-popup')
const create_btn = document.querySelector('.create-post')
const openCategories = document.querySelector('.openCategories')
const post_close = document.querySelector('.post-close')
const  comment = document.querySelector(".create-comment")
let user_id = localStorage.getItem("user_id")

let content = document.querySelector("#content")
creategategory.addEventListener("click", () => {
    create_btn.textContent="Comment"
    creatPostPopup.style.display = "flex"
    openCategories.style.display = "none"
})

post_close.addEventListener("click",()=>{
    creatPostPopup.style.display = "none"
 })
 comment.addEventListener("click",()=>{
    console.log("content is : ",content.value);
    console.log("userid is : ",+user_id);
    createComment(content.value)
 })

export function InitialComment(data) {

    let comments = document.querySelector(".comment")

     cards(data,comments)
    let like = document.querySelectorAll("#likes");
    likes(like)
    // let div = document.createElement("div")
    //         div.innerHTML = `
    //          <div  class="commentFromPost">
    //                         <div class="post-header">
    //                             <img src="../static/imgs/profilePic.png"
    //                                 class="avatar" alt="Profile picture" />
    //                             <div class="user-info">
    //                                 <div class="display-name">${data.firstName+" "+ data.lastName}</div>
    //                                 <span class="username">@${data.firstName}</span>
    //                                 <span class="timestamp">2h</span>
    //                             </div>
    //                         </div>
    //                         <div class="post-content">
    //                             ${data.content}
    //                         </div>
    //                         <div class="post-actions">
    //                             <div class="action active" id="likes"
    //                                 data-id_card="" >
    //                                 <svg width="17" height="17" viewBox="0 0 20 20"
    //                                     fill="currentColor">
    //                                     <path
    //                                         d="M10 19c-.072 0-.145 0-.218-.006A4.1 4.1 0 0 1 6 14.816V11H2.862a1.751 1.751 0 0 1-1.234-2.993L9.41.28a.836.836 0 0 1 1.18 0l7.782 7.727A1.751 1.751 0 0 1 17.139 11H14v3.882a4.134 4.134 0 0 1-.854 2.592A3.99 3.99 0 0 1 10 19Zm0-17.193L2.685 9.071a.251.251 0 0 0 .177.429H7.5v5.316A2.63 2.63 0 0 0 9.864 17.5a2.441 2.441 0 0 0 1.856-.682A2.478 2.478 0 0 0 12.5 15V9.5h4.639a.25.25 0 0 0 .176-.429L10 1.807Z"></path>
    //                                 </svg>
    //                                 <span id="is_liked">${data.likes}</span>
    //                             </div>
    //                             <div class="action" id="dilike">
    //                                 <svg width="17" height="17" viewBox="0 0 20 20"
    //                                     fill="currentColor">
    //                                     <path
    //                                         d="M10 1c.072 0 .145 0 .218.006A4.1 4.1 0 0 1 14 5.184V9h3.138a1.751 1.751 0 0 1 1.234 2.993L10.59 19.72a.836.836 0 0 1-1.18 0l-7.782-7.727A1.751 1.751 0 0 1 2.861 9H6V5.118a4.134 4.134 0 0 1 .854-2.592A3.99 3.99 0 0 1 10 1Zm0 17.193 7.315-7.264a.251.251 0 0 0-.177-.429H12.5V5.184A2.631 2.631 0 0 0 10.136 2.5a2.441 2.441 0 0 0-1.856.682A2.478 2.478 0 0 0 7.5 5v5.5H2.861a.251.251 0 0 0-.176.429L10 18.193Z"></path>
    //                                 </svg>
    //                                 <span id="is_liked">${data.dislikes}</span>
    //                             </div>
    //                             <div class="action">
    //                                 <svg width="17" height="17" viewBox="0 0 20 20"
    //                                     fill="currentColor">
    //                                     <path
    //                                         d="M10 19H1.871a.886.886 0 0 1-.798-.52.886.886 0 0 1 .158-.941L3.1 15.771A9 9 0 1 1 10 19Zm-6.549-1.5H10a7.5 7.5 0 1 0-5.323-2.219l.54.545L3.451 17.5Z"></path>
    //                                 </svg>

    //                                 <span>
    //                                 <a href="/comment?card_id=${data.id}" >${data.comments}</a>
    //                                 </span>
    //                             </div>
    //                         </div>
    //                     </div> 
    //         `
    //comments.appendChild(div)
}

async function createComment(content) {
    console.log("postid is : ",+cardData);
    const response = await fetch("/api/addcomment", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            'Accept': 'application/json',
        },
        body: JSON.stringify({
            user_id:+user_id,
            content:content,
            target_id:+cardData
        })
    })
    if (response.ok) {
        //InitialComment(content)
        await LoadPage(+cardData)
        creatPostPopup.style.display = "none"
        const data = await response.json();
        console.log("Success:", data);
      
    } else {
        const errorData = response.json();
        console.error("Error:", errorData);
        alert(`Error: ${errorData.message || "Request failed"}`);
    }
}