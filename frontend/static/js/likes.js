import fetchData from './forum.js';
function likes(likes, UserID, user_login) {
    // let is_liked = false
    // UserID.forEach(el=>{
    //     let card_id = click.getAttribute("data-id_card");
    //     if(el.UserID===user_login){
    //         likes.forEach(click => {
    //             click.classList.add("clicked");
    //         })
    //     }
        
    // })
    // likes.forEach(click => {
    //     click.addEventListener("click", async (e) => {
    //         e.preventDefault()
           
    //         let islike = click.querySelector("#is_liked");
    //         let card_id = click.getAttribute("data-id_card");
    //         if (click.classList.contains("clicked")) {
               
    //             click.classList.remove("clicked");
    //             deletLikes(card_id)
    //             islike.textContent = parseInt(islike.textContent) - 1;
    //             islike.setAttribute('data-liked', 'false');
    //         } else {
    //             console.log("add");
    //             click.classList.add("clicked");
    //             islike.textContent = parseInt(islike.textContent) + 1;
    //             islike.setAttribute('data-liked', 'true');
    //             addLikes(card_id, 1)
    //         }
    //     })
    // })
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
                card_id: +card_id
            })
        })
        if (response.ok) {
            //fetchData()
            let data = await response.json()
            console.log(data);

        } else if (response.status === 400) {
            //await deletLikes(card_id)
        }
        else {
            let data = await response.json()
            console.log(data);
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
        //fetchData()
        let data = await response.json()
        console.log(data);

    } else {
        let data = await response.json()
        console.log(data);
    }
}

export {
    likes
}