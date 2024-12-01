
const profileNav = document.querySelectorAll(".profile-nav a");
navigate()

profileNav.forEach((navItem) => {
  navItem.addEventListener("click", () => {
    navItem.className = "active";
    profileNav.forEach((item) => {
      if (item != navItem) {
        item.className = "";
      }
    });
  });
});

