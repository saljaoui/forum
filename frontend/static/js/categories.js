import { navigate } from "./home.js"
import { cards } from "./card.js";
import { likes } from "./likescomment.js";
import { search } from "./search.js";
import { status } from "./status.js";
import { alertPopup } from "./alert.js";
const profileNav = document.querySelectorAll(".profile-nav a");
navigate()
let content = []
profileNav.forEach((navItem) => {
  navItem.addEventListener("click", async () => {
    navItem.className = "active";
    await fetchData(navItem.textContent)
    profileNav.forEach((item) => {
      if (item != navItem) {
        item.className = "";
      }
    });
  });
});

async function fetchData(categoryName) {
  const response = await fetch("/api/category", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ category: categoryName }),
  });
  if (response.ok) {
    let data = await response.json();
    let user_info = document.querySelector(".main");
    content = cards(data, user_info)
    search(content)
    let like = document.querySelectorAll("#likes");
      likes(like)
  } else if (!response.ok && !response.status === 409 && !response.status === 400) {
    await status(response)
 }else if( response.status === 409 || response.status === 400) {
     const data = await response.json();
      alertPopup(data,)
  }


}
