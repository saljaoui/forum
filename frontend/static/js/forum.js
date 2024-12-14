import { checkandAdd } from "./addlikes.js";
import { likes } from "./likescomment.js";
import { cards } from "./card.js";
import { checklogin } from "./checklogin.js";
import { search } from "./search.js";
import { status } from "./status.js";
let content = []
checklogin()
const searchInput = document.querySelector("[data-search]")
search(searchInput)

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
  } else if (!responce.ok) {
    status(responce)
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