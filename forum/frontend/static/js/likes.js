import { fetchData } from "./forum.js";
 export function likes(likeElements) {
    const user_data = localStorage.getItem("user_id");

    likeElements.forEach(async (click) => {
        let card_id = click.getAttribute("data-id_card");
        let like = click.getAttribute("data-like");
         const response = await fetch("api/likes", {
            method: "POST",
            body: JSON.stringify({ card_id: +card_id }),
        });

        if (response.ok) {
            let data = await response.json();
            data.forEach((el) => {
                if (el.User_id === +user_data) {
                    if (el.UserLiked && like === "like") {
                        click.classList.add("clicked");
                        click.setAttribute("data-liked", "true");
                    } else if (el.UserDisliked && like === "Dislikes") {
                        click.classList.add("clicked_disliked");
                        click.setAttribute("data-liked", "true");
                    }
                }
            });
        } else if (response.status === 401) {
            location.href = "/login";
        }
        // Add click event listener
        click.addEventListener("click", async (e) => {
            e.preventDefault();
            let card_id = click.getAttribute("data-id_card");
            let like = click.getAttribute("data-like");
            let data_liked = click.getAttribute("data-liked");
            let context = click.getAttribute("data-context"); // Capture context again
            // Handle like and dislike actions
             if (like === "like") {
                if (data_liked === "true") {
                    await deletLikes(user_data, card_id, context);
                    console.log("Removed like");
                } else if (data_liked === "false") {
                    await addLikes(card_id, 1, true, false, context);
                    console.log("Added like");
                }
            } else if (like === "Dislikes") {
                if (data_liked === "true") {
                    await deletLikes(user_data, card_id, context);
                    console.log("Removed dislike");
                } else if (data_liked === "false") {
                    await addLikes(card_id, -1, false, true, context);
                    console.log("Added dislike");
                }
            }
        });
    });
}


async function addLikes(card_id, liked, lik, dislk, context) {
    try {
        let response = await fetch("/api/like", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Accept: "application/json",
            },
            body: JSON.stringify({
                is_liked: +liked,
                card_id: +card_id,
                UserLiked: lik,
                Userdisliked: dislk,
            }),
        });

        if (response.ok) {
            let data = await response.json();
            fetchData()
        } else if (response.status === 401) {
            location.href = "/login";
        }
    } catch (error) {
        console.log(error);
    }
}

async function deletLikes(user_id, card_id, context) {
    let response = await fetch("/api/deleted", {
        method: "DELETE",
        headers: {
            "Content-Type": "application/json",
            Accept: "application/json",
        },
        body: JSON.stringify({ user_id: +user_id, card_id: +card_id }),
    });

    if (response.ok) {
        let data = await response.json();
        fetchData()
    } else if (response.status === 401) {
        location.href = "/login";
    }
}
