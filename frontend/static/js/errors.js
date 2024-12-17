function errors() {
  const user_data = history.state;
  let status = document.querySelector(".status-code")
  let message = document.querySelector(".message")
  if (user_data) {
    message.textContent = user_data.data.message
    status.textContent = user_data.code
  }
  console.log(user_data);
}
errors()