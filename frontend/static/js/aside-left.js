import fetchData  from './forum.js';
const navItems = document.querySelectorAll('.nav-item');
const categoryItems = document.querySelectorAll('.category-item');
const creatPostPopup = document.getElementById('creatPost-popup')
const openCategorie = document.getElementById('categories-popup')
const newPost = document.querySelector('.newPost-popup')
const cancel_btn = document.querySelector('.cancel-btn')
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
// function openCreatPost() {
// }
cancel_btn.addEventListener("click",()=>{
    creatPostPopup.style.display = "none"

})
// function closeCreatPost() {
// }
done_btn.addEventListener("click",()=>{
    defaultCategories()
    categoriesList.forEach(category => {
        if (category.classList.contains('selected')) {
            categoriesSelected.push(category.textContent)
        }
    });
})
// function doneCategories() {
  
// }

create_btn.addEventListener("click",()=>{
    if (categoriesSelected.length > 0 && content.value.length > 0) {
       creatPost()
       creatPostPopup.style.display = "none"
       closeCategories()
       content.value = ""
      // fetchData()
    } else {
       openCategories()
    }
})
// function SeccesCreatPost() {
// }
openCategories.addEventListener('click',()=>{
    openCategorie.style.display = "flex"

})
// function openCategories() {
// }

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

async function creatPost() {
    let content = document.querySelector("#content")
    console.log(content.value, categoriesSelected);
    const response = await fetch("/api/post", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            'Accept': 'application/json',
        },
        body: JSON.stringify({
            title: "title 2",
            content: content.value,
            name: categoriesSelected
        })
    })

    if (response.ok) {
      await  fetchData()
        const data = await response.json();
        console.log("Success:", data);
      
    } else {
        const errorData = response.json();
        console.error("Error:", errorData);
        alert(`Error: ${errorData.message || "Request failed"}`);
    }
}

// export {
//     creatPost
// }