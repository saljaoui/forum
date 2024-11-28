const profileNav = document.querySelectorAll(".profile-nav a");

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