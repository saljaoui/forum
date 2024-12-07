import { GetComments } from "./comment.js";
//import { NewPost } from "./forum.js";
import { likes,addLikes ,deletLikes} from "./likescomment.js";
const urlParams = new URLSearchParams(window.location.search);
const cardData = urlParams.get("card_id");
let user_id = localStorage.getItem("user_id")
const creategategory = document.querySelector(".postReply")
const creatPostPopup = document.getElementById('creatPost-popup')//categories-popup
const categories_popup = document.getElementById('categories-popup')
const create_btn = document.querySelector('.create-post')
const openCategories = document.querySelector('.openCategories')
const post_close = document.querySelector('.post-close')
const comment = document.querySelector(".create-comment")
categories_popup.style.display = "none"
creategategory.addEventListener("click", () => {
    create_btn.textContent = "Comment"
    creatPostPopup.style.display = "flex"
    openCategories.style.display = "none"
})

post_close.addEventListener("click", () => {
    creatPostPopup.style.display = "none"
})
comment.addEventListener("click", () => {
    createComment(content.value)
})


function NewPost(postInfo, parent) {
    console.log(postInfo);
    
}

export async function LoadPage(cardid) {
    let main = document.querySelector(".allcomment")
    /// console.log(main, cardData);
    await GetComments(cardid)
}

//await LoadPage(+cardData)


export function InitialComment(ele, comments) {
    ele.map((data) => {
        let div = document.createElement("div")
        div.className = "commens-card"
        div.innerHTML = `
             <div  class="commentFromPost">
                            <div class="post-header">
                                <img src="../static/imgs/profilePic.png"
                                    class="avatar" alt="Profile picture" />
                                <div class="user-info">
                                    <div class="display-name">${data.firstName + " " + data.lastName}</div>
                                    <span class="username">@${data.firstName}</span>
                                    <span class="timestamp">2h</span>
                                </div>
                            </div>
                            <div class="post-content">
                                ${data.content}
                            </div>
                            <div class="post-actions">
                                <div class="action active is_liked" data-context="comment" id="likes" data-liked="false" data-like="like" data-id_card="${data.id}" >
                                    <svg width="17" height="17" viewBox="0 0 20 20"
                                        fill="currentColor">
                                        <path
                                            d="M10 19c-.072 0-.145 0-.218-.006A4.1 4.1 0 0 1 6 14.816V11H2.862a1.751 1.751 0 0 1-1.234-2.993L9.41.28a.836.836 0 0 1 1.18 0l7.782 7.727A1.751 1.751 0 0 1 17.139 11H14v3.882a4.134 4.134 0 0 1-.854 2.592A3.99 3.99 0 0 1 10 19Zm0-17.193L2.685 9.071a.251.251 0 0 0 .177.429H7.5v5.316A2.63 2.63 0 0 0 9.864 17.5a2.441 2.441 0 0 0 1.856-.682A2.478 2.478 0 0 0 12.5 15V9.5h4.639a.25.25 0 0 0 .176-.429L10 1.807Z"></path>
                                    </svg>
                                    <span id="is_liked">${data.likes}</span>
                                </div>
                                <div class="action disliked" id="likes" data-liked="false"  data-like="Dislikes" data-id_card="${data.id}">
                                    <svg width="17" height="17" viewBox="0 0 20 20"
                                        fill="currentColor">
                                        <path
                                            d="M10 1c.072 0 .145 0 .218.006A4.1 4.1 0 0 1 14 5.184V9h3.138a1.751 1.751 0 0 1 1.234 2.993L10.59 19.72a.836.836 0 0 1-1.18 0l-7.782-7.727A1.751 1.751 0 0 1 2.861 9H6V5.118a4.134 4.134 0 0 1 .854-2.592A3.99 3.99 0 0 1 10 1Zm0 17.193 7.315-7.264a.251.251 0 0 0-.177-.429H12.5V5.184A2.631 2.631 0 0 0 10.136 2.5a2.441 2.441 0 0 0-1.856.682A2.478 2.478 0 0 0 7.5 5v5.5H2.861a.251.251 0 0 0-.176.429L10 18.193Z"></path>
                                    </svg>
                                    <span  id="is_Dislikes" data-disliked="disliked">${data.dislikes}</span>
                                </div>
                                <div class="action">
                                    <svg width="17" height="17" viewBox="0 0 20 20"
                                        fill="currentColor">
                                        <path
                                            d="M10 19H1.871a.886.886 0 0 1-.798-.52.886.886 0 0 1 .158-.941L3.1 15.771A9 9 0 1 1 10 19Zm-6.549-1.5H10a7.5 7.5 0 1 0-5.323-2.219l.54.545L3.451 17.5Z"></path>
                                    </svg>
                                    <span>
                                    <a href="/comment?card_id=${data.id}" >${data.comments}</a>
                                    </span>
                                </div>
                            </div>
                        </div> 
            `
        comments.appendChild(div)
    })
    let like = document.querySelectorAll("#likes");
    likes(like)
}
export async function fetchCard(card) {
    try {
        NewPost(card,"main")
        let cardId = card.getAttribute("data-id_card");
        const response = await fetch(`/api/card?id=${cardId}`, {
            method: "GET",
        });

        if (!response.ok) {
            const errorData = await response.json();
            console.error("Error:", errorData);
            alert(`Error: ${errorData.message || "Request failed"}`);
            return;
        }
        const cardData = await response.json();
        console.log("Fetched card data:", cardData);
        let cardElement = card.closest(".commens-card");
        if (cardElement) {
            updateCard(cardElement, cardData, card); // Update only necessary parts of the card
        }
    } catch (error) {
        console.error("Fetch Error:", error);
        alert("An error occurred while fetching the card data.");
    }
}

function updateCard(cardElement, cardData, cardClicked) {
    const postActions = cardElement.querySelector(".post-actions");
    if (postActions) {
        postActions.innerHTML = `
                <div class="action active is_liked ${cardData.likes===1?"clicked":""}" id="likes" data-liked="false" data-like="like" data-id_card="${cardData.id}">
                    <svg width="17" height="17" viewBox="0 0 20 20" fill="currentColor">
                        <path d="M10 19c-.072 0-.145 0-.218-.006A4.1 4.1 0 0 1 6 14.816V11H2.862a1.751 1.751 0 0 1-1.234-2.993L9.41.28a.836.836 0 0 1 1.18 0l7.782 7.727A1.751 1.751 0 0 1 17.139 11H14v3.882a4.134 4.134 0 0 1-.854 2.592A3.99 3.99 0 0 1 10 19Zm0-17.193L2.685 9.071a.251.251 0 0 0 .177.429H7.5v5.316A2.63 2.63 0 0 0 9.864 17.5a2.441 2.441 0 0 0 1.856-.682A2.478 2.478 0 0 0 12.5 15V9.5h4.639a.25.25 0 0 0 .176-.429L10 1.807Z"></path>
                    </svg>
                    <span id="is_liked">${cardData.likes}</span>
                </div>
                <div class="action disliked ${cardData.dislikes===1?"clicked_disliked":""} " id="likes" data-liked="false" data-like="Dislikes" data-id_card="${cardData.id}">
                    <svg width="17" height="17" viewBox="0 0 20 20" fill="currentColor">
                        <path d="M10 1c.072 0 .145 0 .218.006A4.1 4.1 0 0 1 14 5.184V9h3.138a1.751 1.751 0 0 1 1.234 2.993L10.59 19.72a.836.836 0 0 1-1.18 0l-7.782-7.727A1.751 1.751 0 0 1 2.861 9H6V5.118a4.134 4.134 0 0 1 .854-2.592A3.99 3.99 0 0 1 10 1Zm0 17.193 7.315-7.264a.251.251 0 0 0-.177-.429H12.5V5.184A2.631 2.631 0 0 0 10.136 2.5a2.441 2.441 0 0 0-1.856.682A2.478 2.478 0 0 0 7.5 5v5.5H2.861a.251.251 0 0 0-.176.429L10 18.193Z"></path>
                    </svg>
                    <span id="is_Dislikes" data-disliked="disliked">${cardData.dislikes}</span>
                </div>
                <div class="action">
                    <svg width="17" height="17" viewBox="0 0 20 20" fill="currentColor">
                        <path d="M10 19H1.871a.886.886 0 0 1-.798-.52.886.886 0 0 1 .158-.941L3.1 15.771A9 9 0 1 1 10 19Zm-6.549-1.5H10a7.5 7.5 0 1 0-5.323-2.219l.54.545L3.451 17.5Z"></path>
                    </svg>
                    <span>
                        <a href="/comment?card_id=${cardData.id}">${cardData.comments}</a>
                    </span>
                </div>
    ` ;
    }
    console.log("Updated card actions:", postActions);
}


async function createComment(content) {
    const response = await fetch("/api/addcomment", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            'Accept': 'application/json',
        },
        body: JSON.stringify({
            user_id: +user_id,
            content: content,
            target_id: +cardData
        })
    })
    if (response.ok) {
        creatPostPopup.style.display = "none"
        const data = await response.json();
        console.log("Success:", data);

    } else {
        const errorData = response.json();
        console.error("Error:", errorData);
        alert(`Error: ${errorData.message || "Request failed"}`);
    }
}

// Event handler to add likes/dislikes and update card details
document.body.addEventListener("click", async (e) => {
    const click = e.target.closest(".action"); // Ensure the clicked element is an action button
    if (!click || !click.matches(".is_liked, .disliked")) return; // Ignore unrelated clicks

    e.preventDefault();

    const user_data = localStorage.getItem("user_id");
    const card_id = click.getAttribute("data-id_card");
    const like = click.getAttribute("data-like");
    const data_liked = click.getAttribute("data-liked");

    try {
        if (like === "like") {
            if (data_liked === "true") {
                await deletLikes(user_data, card_id, click);
                console.log("Removed like");
            } else {
                await addLikes(card_id, 1, true, false, click);
                console.log("Added like");
            }
        } else if (like === "Dislikes") {
            if (data_liked === "true") {
                await deletLikes(user_data, card_id, click);
                console.log("Removed dislike");
            } else {
                await addLikes(card_id, -1, false, true, click);
                console.log("Added dislike");
            }
        }

        await fetchCard(click); // Update the card data after liking/disliking
    } catch (error) {
        console.error("Error handling like/dislike:", error);
    }
});

// export function addLike_card(){
//     const user_data = localStorage.getItem("user_id");
//     document.body.addEventListener("click", async (e) => {
//         const click = e.target.closest(".action"); 
//         if (!click || !click.matches(".is_liked, .disliked")) return;
//         e.preventDefault();
//          let card_id = click.getAttribute("data-id_card");
//          let like = click.getAttribute("data-like");
//         let data_liked = click.getAttribute("data-liked");
//         if (like === "like") {
//             if (data_liked === "true") {
//                 await deletLikes(user_data, card_id, click);
//                 console.log("Removed like");
//             } else if (data_liked === "false") {
//                 await addLikes(card_id, 1, true, false, click);
//                 console.log("Added like");
//             }
//         } else if (like === "Dislikes") {
//             if (data_liked === "true") {
//                 await deletLikes(user_data, card_id, click);
//                 console.log("Removed dislike");
//             } else if (data_liked === "false") {
//                 await addLikes(card_id, -1, false, true, click);
//                 console.log("Added dislike");
//             }
//         }
//         fetchCard(click)
//     });
// }
// addLike_card()