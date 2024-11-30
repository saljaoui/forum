let nav_items = document.querySelector("#home")//search
let search = document.querySelector("#search")//search
let categories = document.querySelector("#categories")//search
let profile = document.querySelector("#profile")//search

 function navigate(){

    nav_items.addEventListener("click", (e) => {
        const currentUrl = window.location.pathname;
        if (currentUrl !== "/home") {
            location.href = "/home";
        } else {
            e.preventDefault()
    
            console.log("Already on the Home page, no need to reload.");
        }
    })
    
    profile.addEventListener("click", (e) => {
        e.preventDefault()
        const currentUrl = window.location.pathname;
        if (currentUrl !== "/profile") {
            location.href = "/profile";
        } else {
            console.log("Already on the profile page, no need to reload.");
        }
    })
}
navigate()
export {
    navigate
}
// search.addEventListener("click", (e) => {
//     const currentUrl = window.location.pathname;

//     if (currentUrl !== "/search") {
//         location.href = "/search";
//     } else {
//         e.preventDefault()

//         console.log("Already on the search page, no need to reload.");
//     }
// })

// categories.addEventListener("click", (e) => {
//     e.preventDefault()
//     const currentUrl = window.location.pathname;
//     if (currentUrl !== "/categories") {
//         location.href = "/categories";
//     } else {
//         e.preventDefault()
//         console.log("Already on the categories page, no need to reload.");
//     }
// })
