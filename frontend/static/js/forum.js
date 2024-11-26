import addpost from './post.js';
let containr = document.querySelector(".containr")
async function fetchData() {
    const responce = await fetch("http://localhost:3333/api/home", { method: "GET" });
    if (responce.ok) {
        let data = await responce.json();
        let post = document.querySelector("#post");
        post.innerHTML = "";
        data.map(ele => {
            let div = document.createElement("div")
            let title = document.createElement("h2")
            let content = document.createElement("p")
            let a = document.createElement("a")
            a.href = `comment?target_id=${ele.ID}`
            div.classList = "card"
            title.style.color = "#ffff"
            content.style.color = "#ffff"
            let date = new Date(ele.CreatedAt)
            title.textContent = ele.Title + "====>" + date.getDate() + "  " + date.getHours() + ":" + date.getMinutes() + ":" + date.getSeconds()
            content.textContent = ele.Content
            div.append(title, content)
            a.append(div)
            post.appendChild(a)
        })
        containr.appendChild(post)

    } else {
        let data = responce.json()
        console.log(data);

    }
}
document.addEventListener("DOMContentLoaded", fetchData);

if (document.cookie) {
    let tokens = document.cookie.split("; ")
    let token = null;
    let userId = null;
    tokens.forEach(ele => {
        let [key, value] = ele.split("=")
        if (key === 'token') {
            token = value
        }
        else if (key === "user_id") {
            userId = value
        }
    })

    addpost()
    console.log(token, userId);

} else {
    const container = document.querySelector(".form");
    while (container.firstChild) {
        container.removeChild(container.firstChild)
    }
}



export {
    fetchData
}