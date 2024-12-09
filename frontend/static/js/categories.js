import { navigate } from "./home.js"
import { cards } from "./card.js";
import { likes } from "./likes.js";

const profileNav = document.querySelectorAll(".profile-nav a");
navigate()
let content = ""
profileNav.forEach((navItem) => {
  navItem.addEventListener("click", () => {
    navItem.className = "active";
    fetchData(navItem.textContent)
    content = navItem.textContent
    profileNav.forEach((item) => {
      if (item != navItem) {
        item.className = "";
      }
    });
  });
});

export default async function fetchData(categoryName) {
  const responce = await fetch("http://localhost:3333/api/category", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ category: categoryName }),
  });
  if (responce.ok) {
    let data = await responce.json();
    let user_info = document.querySelector(".main");
    console.log(data);
    
    //user_info.innerHTML = "";
    content = cards(data,user_info)
    let like = document.querySelectorAll("#likes");
    likes(like)
    
  } 
  // else if(responce.status===401) {
  //  // location.href="/login"
  // }
}


//fetchData(categoryName);