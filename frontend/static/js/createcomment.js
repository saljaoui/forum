const creategategory = document.querySelector(".postReply")
const creatPostPopup = document.getElementById('creatPost-popup')
const create_btn = document.querySelector('.create-post')
const openCategories = document.querySelector('.openCategories')
const post_close = document.querySelector('.post-close')
const  createComment=document.querySelector(".create-comment")
let content = document.querySelector("#content")
creategategory.addEventListener("click", () => {
    create_btn.textContent="Comment"
    creatPostPopup.style.display = "flex"
    openCategories.style.display = "none"
})

post_close.addEventListener("click",()=>{
    creatPostPopup.style.display = "none"
 })
 createComment.addEventListener("click",()=>{
    console.log(content.value);
    
 })