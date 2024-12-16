function errors() {
  let path = window.location.pathname


  const user_data = history.state;
  let status = document.querySelector(".status-code")
  let message = document.querySelector(".message")



  if (user_data) {
    message.textContent = user_data.data.message
    status.textContent = user_data.code
  }
  // else  if (path === "/err" && !user_data) {
  //   message.textContent = "404 Not Found: The requested resource could not be located"
  //   status.textContent = 404
  // }
  // else 
  // if (path != "/comment" && path != "/register" && path != "/login" && path != "/logout" &&
  //   path != "/about" && path != "/contact" && path != "/home" && path != "/categories"
  //   && path != "/contact" && path != "/comment" && path != "/profile"
  //   && path != "/settings" && path != "/err") {
  //   message.textContent = "Path not found"
  //   status.textContent = 404
  // }
  console.log(user_data);
}
errors()