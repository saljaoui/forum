import {creatPost}  from './post.js';
const navItems = document.querySelectorAll('.nav-item');
const categoryItems = document.querySelectorAll('.category-item');
const creatPostPopup = document.getElementById('creatPost-popup')
const openCategorie = document.getElementById('categories-popup')
const newPost = document.querySelector('.newPost-popup')
const cancel_btn = document.querySelector('.category')
const post_close = document.querySelector('.post-close')
const openCategories = document.querySelector('.openCategories')
const create_btn = document.querySelector('.create-post')
const done_btn = document.querySelector('.done-post')
const categoriesList = Array.from(document.getElementsByClassName('category-item'))
let categoriesSelected = []


function activeByDefault() {
    navItems.forEach(navItem => {
        const outlineIcon = navItem.querySelector('ion-icon[name$="-outline"], ion-icon[name$="-sharp"]');
        const filledIcon = navItem.querySelector('ion-icon:not([name$="-outline"]):not([name$="-sharp"])');
        filledIcon.classList.remove('active');
        outlineIcon.classList.add('active');
        navItem.classList.remove('active')
    });
}

navItems.forEach(navItem => {
    navItem.addEventListener('click', function () {
        activeByDefault();
        navItem.classList.add('active')
        const outlineIcon = this.querySelector('ion-icon[name$="-outline"], ion-icon[name$="-sharp"]');
        const filledIcon = this.querySelector('ion-icon:not([name$="-outline"]):not([name$="-sharp"])');
        if (outlineIcon.classList.contains('active')) {
            outlineIcon.classList.remove('active');
            filledIcon.classList.add('active');
        }
    });
});
newPost.addEventListener("click",()=>{
    creatPostPopup.style.display = "flex"
})

// creategategory.addEventListener("click",()=>{
//  creatPostPopup.style.display = "flex"
//  openCategories.style.display="none"
// })
 
cancel_btn.addEventListener("click",()=>{
    defaultCategories()
   // creatPostPopup.style.display = "none"
})
post_close.addEventListener("click",()=>{
   creatPostPopup.style.display = "none"
})
 
done_btn.addEventListener("click",()=>{
    defaultCategories()
    categoriesList.forEach(category => {
        if (category.classList.contains('selected')) {
            categoriesSelected.push(category.textContent)
        }
    });
})
 
create_btn.addEventListener("click",()=>{
    if (categoriesSelected.length > 0 && content.value.length > 0) {
       creatPost(categoriesSelected)
       creatPostPopup.style.display = "none"
       closeCategories()
       content.value = ""
      // fetchData()
    }  
})
 
openCategories.addEventListener('click',()=>{
    openCategorie.style.display = "flex"

})
 

function closeCategories() {
    defaultCategories()
    categoriesList.forEach(category => {
        if (category.classList.contains('selected')) {
            category.classList = "category-item"
        }
    });
}

function defaultCategories() {
    categoriesSelected = []
    openCategorie.style.display = "none"
}

categoryItems.forEach(item => {
    item.addEventListener('click', () => {
        item.classList.toggle('selected');
    });
});

