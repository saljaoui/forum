import fetchData  from './forum.js';
function likes(likes, dislikes, post_liked) {
    let is_liked = false
    if(post_liked==="1 liked "){
        is_liked=true
         likes.style.color = "var(--color-action-hover)"
    }
    
    likes.addEventListener("click", async () => {
         if (is_liked == false && post_liked==="0 liked ") {
            likes.style.color = "var(--color-action-hover)"
            is_liked = true
            likes.classList.remove("active")
            let card_id = likes.dataset.id_card
            addLikes(card_id, 1)
        }
        else if (is_liked == true ) {
            console.log(post_liked==="1 liked ");
            is_liked = false
            likes.style.color = "var(--color-muted)"
            likes.classList.add("active")
            let card_id = likes.dataset.id_card
            deletLikes(card_id)
        }
    })


}
async function addLikes(card_id, liked) {
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
        fetchData()
        let data = await response.json()
        console.log(data);

    } else {
        let data = await response.json()
        console.log(data);
    }
}
async function deletLikes(card_id) {
    console.log(card_id);
    
    let response = await fetch("/api/deleted", {
        method: "DELETE",
        headers: {
            "Content-Type": "application/json",
            'Accept': 'application/json',
        },
        body:JSON.stringify({
            card_id: +card_id
        })
    })
    if (response.ok) {
        fetchData()
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