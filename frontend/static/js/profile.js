import { navigate } from "./home.js";
import { likes } from "./likescomment.js";
import {cards}from "./card.js";
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
    // SeccesCreatPost()
    //const user_data = history.state;
    // console.log(user_data);

    let data = await responce.json();           
    let user_info = document.querySelector(".main");
    content = cards(data, user_info)

    let like = document.querySelectorAll("#likes"); 
    // console.log(data);
    likes(like );
  } else {
    let data = responce.json();
    console.log(data);
  }
}
