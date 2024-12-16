import { navigate } from "./home.js"
import { cards } from "./card.js";
import { likes } from "./likescomment.js";
import { search } from "./search.js";
import { status } from "./status.js";
const profileNav = document.querySelectorAll(".profile-nav a");
navigate()
let content = []
profileNav.forEach((navItem) => {
  navItem.addEventListener("click", () => {
    navItem.className = "active";
    fetchData(navItem.textContent)
    profileNav.forEach((item) => {
      if (item != navItem) {
        item.className = "";
      }
    });
  });
});

const categories = {
  "General": 0,"Sports": 1,"Entertainment": 2,"Politics": 3,"Technology": 4,"Business": 5,"Science": 6,"Health": 7,"Food": 8,"Travel": 9,"Fashion": 10,"Art": 11,"Music": 12
};
export default async function fetchData(categoryName) {
  if (categories[categoryName] === undefined) {
    alert("don't touch in inspect");
    return;
  }
  const responce = await fetch("/api/category", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ category: categoryName }),
  });
  if (responce.ok) {
    let data = await responce.json();
    let user_info = document.querySelector(".main");
    content = cards(data, user_info)
    search(content)
    let like = document.querySelectorAll("#likes");
    likes(like)
  } else if (!responce.ok) {
    status(responce)
  }
}
