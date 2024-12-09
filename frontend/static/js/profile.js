import { navigate } from "./home.js";
import { likes } from "./likes.js";
import { cards } from "./card.js";
const profileNav = document.querySelectorAll(".profile-nav a");
navigate();

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

export default async function fetchData(navIdName) {
  console.log(navIdName);
  if (navIdName === undefined) {
    return;
  }
  const responce = await fetch(`/api/profile/${navIdName}`, {
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
// fetchData();
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
