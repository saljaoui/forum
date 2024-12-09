import { createComment } from "./createcomment.js"
function classes() {
    const creategategory = document.querySelector(".postReply")
    const creatPostPopup = document.getElementById('creatPost-popup')//categories-popup
    const categories_popup = document.getElementById('categories-popup')
    const create_btn = document.querySelector('.create-post')
    const openCategories = document.querySelector('.openCategories')
    
    while(categories_popup.firstChild){
        categories_popup.removeChild(categories_popup.firstChild)
    }
    const post_close = document.querySelector('.post-close')
    const comment = document.querySelector(".create-comment")
    comment.addEventListener("click",()=>{
        creatPostPopup.style.display="none"
        
    })
   // categories_popup.style.display = "none"
    creategategory.addEventListener("click", () => {
        create_btn.textContent = "Comment"
        creatPostPopup.style.display = "flex"
     })

    post_close.addEventListener("click", () => {
        creatPostPopup.style.display = "none"
    })
    comment.addEventListener("click", () => {
        createComment(content.value)
        content.value=""
    })
}
classes() 