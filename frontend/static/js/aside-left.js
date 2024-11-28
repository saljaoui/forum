//import { addpost } from './post.js';
const navItems = document.querySelectorAll('.nav-item');
const categoryItems = document.querySelectorAll('.category-item');
const creatPostPopup = document.getElementById('creatPost-popup')
const openCategorie = document.getElementById('categories-popup')
let content = document.querySelector("#content")
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

function openCreatPost() {
    creatPostPopup.style.display = "flex"
}
function closeCreatPost() {
    creatPostPopup.style.display = "none"
}

function doneCategories() {
    defaultCategories()
    categoriesList.forEach(category => {
        if (category.classList.contains('selected')) {
            categoriesSelected.push(category.textContent)
        }
    });
}

function SeccesCreatPost() {
     if (categoriesSelected.length > 0 && content.value.length > 0) {
        creatPost(categoriesSelected)

        creatPostPopup.style.display = "none"
        closeCategories()

        content.value = ""
    } else {
        openCategories()
    }
    console.log(doneCategories());

}

function openCategories() {
    openCategorie.style.display = "flex"
}

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

async function creatPost(categories) {
    console.log(content.value, categoriesSelected);
    const response = await fetch("/api/post", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            'Accept': 'application/json',
        },
        body: JSON.stringify({
            title: "t",
            content: "testtt",
            name: [
                "hrlo",'frefrr'
            ]
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