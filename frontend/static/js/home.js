let nav_items = document.querySelectorAll(".nav-item")

nav_items.forEach(nav => {
    nav.addEventListener("click", (e) => {
        e.preventDefault()
        let url = nav.getAttribute("href")
        history.pushState({}, '', url);
        location.replace(url)
    })
}) 