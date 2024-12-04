import fetchData from './forum.js';
function likes(likes, disliked) {

    likes.forEach(click => {
        click.addEventListener("click", async (e) => {
            e.preventDefault()

            let card_id = click.getAttribute("data-id_card");
            let check_likes = click.getAttribute("data-like");
            if (check_likes === "like") {
                let data_liked = click.getAttribute("data-liked");
                if (data_liked === "true") {
                    await deletLikes(card_id)
                } else {
                    await addLikes(card_id, 1)
                }
            } else if (check_likes === "Dislikes") {
                console.log("dislike");

            }
            // let data_liked = click.getAttribute("data-liked");
            // if (data_liked === "true") {
            //     await deletLikes(card_id)
            // } else {
            //     await addLikes(card_id, 1)
            // }
        })
    })
}
async function addLikes(card_id, liked) {
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
                UserLiked: true
            })
        })
        if (response.ok) {
            fetchData()
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
async function deletLikes(card_id) {
    let response = await fetch("/api/deleted", {
        method: "DELETE",
        headers: {
            "Content-Type": "application/json",
            'Accept': 'application/json',
        },
        body: JSON.stringify({
            card_id: +card_id
        })
    })
    if (response.ok) {
        fetchData()
        let data = await response.json()
        console.log(data);

    } else if (response.status === 401) {
        location.href = "/login"
    }
}

export {
    likes
}