const creategategory = document.querySelector(".postReply")
const creatPostPopup = document.getElementById('creatPost-popup')
const create_btn = document.querySelector('.create-post')
const openCategories = document.querySelector('.openCategories')
const post_close = document.querySelector('.post-close')
const  comment=document.querySelector(".create-comment")
let user_id=localStorage.getItem("user_id")



const urlParams = new URLSearchParams(window.location.search);
const id = urlParams.get("card_id");

let content = document.querySelector("#content")
creategategory.addEventListener("click", () => {
    create_btn.textContent="Comment"
    creatPostPopup.style.display = "flex"
    openCategories.style.display = "none"
})

post_close.addEventListener("click",()=>{
    creatPostPopup.style.display = "none"
 })
 comment.addEventListener("click",()=>{
    console.log("content is : ",content.value);
    console.log("userid is : ",+user_id);
    createComment(content.value)
 })

 async function createComment(content) {
    console.log("postid is : ",+id);

    const response = await fetch("/api/addcomment", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            'Accept': 'application/json',
        },
        body: JSON.stringify({
            user_id:+user_id,
            content:content,
            target_id:+id
        })
    })

    if (response.ok) {
        const data = await response.json();
        console.log("Success:", data);
      
    } else {
        const errorData = response.json();
        console.error("Error:", errorData);
        alert(`Error: ${errorData.message || "Request failed"}`);
    }
}