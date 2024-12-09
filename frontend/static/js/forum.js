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

// if (document.cookie) {
//   let join = document.querySelector(".join");
//   join.style.display = "none";
//   let aside_nav = document.querySelector(".aside-nav");
//   aside_nav.style.display = "block";
//   while (join.firstChild) {
//     join.removeChild(join.firstChild);
//   }
//   // location.href="/home"
//   let tokens = document.cookie.split("; ");
//   let token = null;
//   let userId = null;
//   tokens.forEach((ele) => {
//     let [key, value] = ele.split("=");
//     if (key === "token") {
//       token = value;
//     } else if (key === "user_id") {
//       userId = value;
//     }
//   });
//   // console.log(token, userId);
// } else {

//   let join = document.querySelector(".join");
//   join.style.display = "block";
//   let aside_nav = document.querySelector(".aside-nav");
//   aside_nav.style.display = "none";
//   while (aside_nav.firstChild) {
//     aside_nav.removeChild(aside_nav.firstChild);
//   }
// }