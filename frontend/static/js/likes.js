import fetchData from './forum.js';
function likes(likes, disliked) {
    const user_data = localStorage.getItem("user_id");
    likes.forEach(async click => {
        let card_id = click.getAttribute("data-id_card");
       /// console.log(card_id);
        
        // const responce = await fetch("http://localhost:3333/api/likes", {
        //     method: "POST",
        //     body: JSON.stringify({
        //         "user_id": +user_data,
        //         "card_id": +card_id
        //     })
        // });
        // if (responce.ok) {
        //     console.log(likes);
        //     let data =await responce.json()
        //     if(data.message.UserLiked){
        //         console.log("liked");
                
        //     }
        //     if(data.message.UserDisliked){
        //         console.log("Disliked");
                
        //     }
           // console.log(data.message);
            
             
        //}
        //clicked_disliked
        //  let check_likes = click.getAttribute("data-like");

        // click.addEventListener("click", async (e) => {
        //     e.preventDefault()
        //     let card_id = click.getAttribute("data-id_card");
        //     if (check_likes === "like") {
        //         let data_liked = click.getAttribute("data-liked");
        //         if (data_liked === "true") {
        //             await deletLikes(card_id)
        //         } else {
        //             await addLikes(card_id, 1)
        //         }
        //     } else if (check_likes === "Dislikes") {
        //         console.log("dislike");

        //     }
        // })
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