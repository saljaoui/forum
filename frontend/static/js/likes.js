import fetchData from './forum.js';
function likes(likes, dislikes, post_liked) {
    let is_liked = false
    likes.forEach(click => {
        click.addEventListener("click", () => {
            likes.forEach(btn => btn.classList.remove("clicked"));
            click.classList.add("clicked")
            let clickd = click.classList.contains("clicked")
            if (clickd) {
                if (post_liked === "1 liked " && is_liked==false) {
                    is_liked = true
                    click.style.color = "var(--color-action-hover)"
                    let card_id = click.dataset.id_card
                    console.log(card_id);

                    addLikes(card_id, 1)
                } else if (post_liked === "0 liked " && is_liked===true) {
                    is_liked=false
                    console.log("hello");
                }
                // click.style.color = "var(--color-action-hover)"
                // click.classList.remove("active")
              
            }

        })
    })
    // 
    // if(post_liked==="1 liked "){
    //     is_liked=true
    //      likes.style.color = "var(--color-action-hover)"
    // }

    // if(post_liked==="0 liked "){
    //    console.log("hello");

    // }

    // likes.addEventListener("click", async () => {
    //  if (is_liked == false && post_liked==="0 liked ") {
    //     likes.style.color = "var(--color-action-hover)"
    //     is_liked = true
    //     likes.classList.remove("active")
    //     let card_id = likes.dataset.id_card
    //     console.log(card_id);

    //     addLikes(card_id, 1)
    // }
    //     else if (is_liked == true ) {
    //         console.log(post_liked==="1 liked ");
    //         is_liked = false
    //         likes.style.color = "var(--color-muted)"
    //         likes.classList.add("active")
    //         let card_id = likes.dataset.id_card
    //         deletLikes(card_id)
    //     }
    // })


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

    } else {
        let data = await response.json()
        console.log(data);
    }
}

export {
    likes
}