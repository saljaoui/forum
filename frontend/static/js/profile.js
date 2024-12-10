import { navigate } from "./home.js";
import { cards } from "./card.js";
import { likes } from "./likescomment.js";

const profileNav = document.querySelectorAll(".profile-nav a");
navigate();
let content = []
fetchData("posts");
profileNav.forEach((navItem) => {
  navItem.addEventListener("click", (e) => {
    const navId = navItem.getAttribute("id");
    if (navId === undefined) {
      return;
    }
    fetchData(navId);
    navItem.className = "active";
    profileNav.forEach((item) => {
      if (item != navItem) {
        item.className = "";
      }
    });
  });
});
//--------------------------------------
 async function fetchData(id) {
  const responce = await fetch("/api/profile/"+id, {
    method: "GET",
  });
  if (responce.ok) {
    let user_info = document.querySelector(".profile");
    let data = await responce.json();
    cards(data, user_info)
    let like = document.querySelectorAll("#likes");
    likes(like)
  } 
  // else if (responce.status === 401) {
  //  // location.href = "/login";
  // } else {
  //   let data = responce.json();
  //   console.log(data);
  // }
}
