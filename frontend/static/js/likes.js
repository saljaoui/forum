import {fetchData} from './forum.js';
function likes(likes, disliked) {
    const user_data = localStorage.getItem("user_id");

    likes.forEach(async click => {
        let card_id = click.getAttribute("data-id_card");

        let like = click.getAttribute("data-like")
        const responce = await fetch("api/likes", {
            method: "POST",
            body: JSON.stringify({
                "card_id": +card_id
            })
        });
        if (responce.ok) {
            let data = await responce.json()
            data.forEach(el => {
                if (el.User_id === +user_data) {
                    if (el.UserLiked && like === "like") {
                        click.classList.add("clicked")
                        click.setAttribute("data-liked", "true")
                    } else if (el.UserDisliked && like === "Dislikes") {
                        click.classList.add("clicked_disliked")
                        click.setAttribute("data-liked", "true")
                    }
                }
            })
        }
        else if (responce.status === 401) {
            location.href = "/login"
        }
        click.addEventListener("click", async (e) => {
            e.preventDefault()
            let card_id = click.getAttribute("data-id_card");
            let like = click.getAttribute("data-like")
            let data_liked = click.getAttribute("data-liked");
            console.log(data_liked);
            
            if (like === "like") {
                if (data_liked === "true") {
                    await deletLikes(user_data, card_id)
                } else {
                    await addLikes(card_id, 1, true, false)
                }
            } else if (like === "Dislikes") {
                if (data_liked === "true") {
                    await deletLikes(user_data, card_id)
                } else {

                    await addLikes(card_id, -1, false, true)
                }
            }
        })
    })
}
async function addLikes(card_id, liked, lik, dislk) {
    try {
        let response = await fetch("/api/like", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                'Accept': 'application/json',
            },
            body: JSON.stringify({
                is_liked: +liked,
                card_id: +card_id,
                UserLiked: lik,
                Userdisliked: dislk
            })
        })
        if (response.ok) {
            //fetchData()
            let data = await response.json()
            console.log(data);
        }
        else if (response.status === 401) {
            location.href = "/login"
        }

    } catch (error) {
        console.log(error);

    }

}
async function deletLikes(user_id, card_id) {
    let response = await fetch("/api/deleted", {
        method: "DELETE",
        headers: {
            "Content-Type": "application/json",
            'Accept': 'application/json',
        },
        body: JSON.stringify({
            user_id: +user_id,
            card_id: +card_id
        })
    })
    if (response.ok) {
        //fetchData()
        let data = await response.json()
        console.log(data);

    } else if (response.status === 401) {
        location.href = "/login"
    }
}

export {
    likes
}