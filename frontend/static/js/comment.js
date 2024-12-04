import { InitialComment } from "./createcomment.js"

const urlParams = new URLSearchParams(window.location.search);
const cardData = urlParams.get("card_id");
console.log("card data : ",cardData)

function NewComment() {
    let comment = document.createElement("div")
    comment.className = "comment"
    let main = document.querySelector("main")
    main.appendChild(comment)
}
NewComment()
export async function GetComments() {
    const response = await fetch(`/api/comment?target_id=${+cardData}`, {
        method: "GET",
    });
    if (response.ok) {
        let datacomment = await response.json()
        console.log("datacomment : ",datacomment);
        console.log("rani ndkhel ")
        let comments = document.querySelector(".comment")
        comments.innerHTML = ""
        datacomment.map(ele => {
            InitialComment(ele)
       })
    }
    else {
        console.log("err");
    }
}

await GetComments()