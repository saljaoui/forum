import { status } from "./status.js";

async function errors() {
  const user_data = history.state;
  let statuscode = document.querySelector(".status-code")
  let message = document.querySelector(".message")
  if (user_data) {
    message.textContent = user_data.data.message
    statuscode.textContent = user_data.code
  }else{
    let code=404
    await status(code)
    // message.textContent="Page Not Found "
    // status.textContent=404
  }
   
}
errors()