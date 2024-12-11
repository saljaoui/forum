
export async function likes(likeElements) {
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
                    localStorage.setItem("user_login", el.User_id);
                    if (el.UserLiked && like === "like") {
                        click.classList.add("clicked");
                        click.setAttribute("data-liked", "true");
                    } else if (el.UserDisliked && like === "Dislikes") {
                        click.classList.add("clicked_disliked");
                        click.setAttribute("data-liked", "true");
                    }
                }
            });
        }
    });
}


export async function addLikes(card_id, liked, lik, dislk, click) {
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
        } 
        // else if (response.status === 401) {
        //     //location.href = "/login";
        // }
    } catch (error) {
        console.log(error);
    }
}

export async function deletLikes(user_id, card_id, click) {
    console.log(user_id);
    
    let response = await fetch("/api/deleted", {
        method: "DELETE",
        headers: {
            "Content-Type": "application/json",
            Accept: "application/json",
        },
        body: JSON.stringify({ uuid: +user_id, card_id: +card_id }),
    });

    if (response.ok) {
        let data = await response.json();
        // fetchCard(card_id,click)
        // //console.log(data);
        // // if (context === "post") {
        // //     fetchData(); // Refresh post section
        // //     return
        // //  } else if (context === "comment") {
        // //       // Reload only the comment section
        // //      fetchCard(card_id)
        // //  }
    }
    //  else if (response.status === 401) {
    //     //location.href = "/login";
    // }
}
