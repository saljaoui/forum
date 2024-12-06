import { InitialComment,LoadPage } from "./createcomment.js"
const urlParams = new URLSearchParams(window.location.search);
const cardData = urlParams.get("card_id");
console.log("card data : ",cardData)


//await LoadPage(+cardData)

export function NewComment() {
    let comment = document.createElement("div")
    comment.className = "comment"
    let main = document.querySelector("main")
    main.appendChild(comment)
}
//NewComment()
export async function GetComments(cardid) {
    
    const response = await fetch(`/api/comment?target_id=${cardid}`, {
        method: "GET",
    });
    if (response.ok) {
        let datacomment = await response.json()
        console.log("datacomment : ",datacomment);
        console.log("rani ndkhel ")
        let comments = document.querySelector(".comment")
        comments.innerHTML = ""
        InitialComment(datacomment)
    //     datacomment.map(ele => {
    //    })
    }
    else {
        console.log("err");
    }
}
