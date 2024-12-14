import { navigate } from "./home.js"
import { cards } from "./card.js";
import { likes } from "./likescomment.js";
import { search } from "./search.js";
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

export default async function fetchData(categoryName) {
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
     content = cards(data,user_info)
      //console.log(content);
      search(content)
     // console.log(content);
      
    let like = document.querySelectorAll("#likes");
    likes(like)
  }  


}
