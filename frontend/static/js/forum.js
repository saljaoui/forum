import { checkandAdd } from "./addlikes.js";
import { likes } from "./likescomment.js";
import { cards } from "./card.js";
import { checklogin } from "./checklogin.js";
import { search } from "./search.js";
let content = []
checklogin()
const searchInput = document.querySelector("[data-search]")
if (searchInput) {
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
}

export async function fetchData() {
  const responce = await fetch("/api/home", {
    method: "GET",
  });
  if (responce.ok) {
    let path = window.location.pathname
    if (path !== "/profile") {
      let data = await responce.json();
      let user_info = document.querySelector(".main");
      content = cards(data, user_info)

      let like = document.querySelectorAll("#likes");
      likes(like)
      search(content)
    }
  }
  // else if (responce.status === 401) {
  //   let body = document.querySelector("body")
  //   body.style.display = "none"
  //    //location.href = "/login"
  // }

}
fetchData()
document.addEventListener("DOMContentLoaded", () => {
  checkandAdd();
});