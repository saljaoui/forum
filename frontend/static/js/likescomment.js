
export async function likes(likeElements) {
    
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
                 let tokens = document.cookie.split("token=")
                console.log(tokens[1]===el.Uuid);
                if (el.Uuid ===tokens[1]) {
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

export async function deletLikes(card_id) {
    console.log(card_id);
    
    let response = await fetch("/api/deleted", {
        method: "DELETE",
        headers: {
            "Content-Type": "application/json",
            Accept: "application/json",
        },
        body: JSON.stringify({ card_id: +card_id }),
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
     else  {
        let data = await response.json();
        console.log(data);
         
    }
}
