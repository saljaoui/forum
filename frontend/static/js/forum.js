import { likes } from "./likes.js";
import { cards } from "./card.js";
const user_data = localStorage.getItem("user_id");
let content = []
const searchInput = document.querySelector("[data-search]")
searchInput.addEventListener("input", (e) => {
  const value = e.target.value.toLowerCase()
  content.forEach(data => {
    const isVisible = data.data.toLowerCase().includes(value)
    if (!isVisible) {
      data.element.style.display = "none"
    } else {
      data.element.style.display = "block"
    }
  })
})

export   async function fetchData() {
  const responce = await fetch("/api/home", {
    method: "GET",
  });
  if (responce.ok) {
 
    let data = await responce.json();
    let user_info = document.querySelector(".main");
    content = cards(data,user_info)
     let like = document.querySelectorAll("#likes");
     likes(like)
  } else if (responce.status === 401) {
    let body = document.querySelector("body")
    body.style.display = "none"
     location.href = "/login"
  }
}
fetchData()