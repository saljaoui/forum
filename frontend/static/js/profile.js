import { navigate } from "./home.js";
import { likes } from "./likescomment.js";
import { cards } from "./card.js";
import { status } from "./status.js";
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

function profileInfo() {
  const profileInfo = document.querySelector('.profile-info');
  let dataUser = JSON.parse(localStorage.getItem("data"))
  console.log(dataUser.firstname);
  const profileName = document.createElement('h1');
  profileName.classList.add('profile-name');
  profileName.textContent = dataUser.firstname + " " + dataUser.lastname;
  
  const profileHandle = document.createElement('p');
  profileHandle.classList.add('profile-handle');
  profileHandle.textContent = `🌟 ` + dataUser.email;

  profileInfo.appendChild(profileName);
  profileInfo.appendChild(profileHandle);
}

profileInfo()



//--------------------------------------
async function fetchData(id) {
  const responce = await fetch("/api/profile/" + id, {
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
    likes(like);
  } else if (!responce.ok) {
    status(responce)
  } else {
    let data = responce.json();
    console.log(data);
  }
}
