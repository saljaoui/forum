
export default async function fetchData() {

  const responce = await fetch("http://localhost:3333/api/home", { method: "GET" });
  if (responce.ok) {
   // SeccesCreatPost()
    const user_data = history.state
    let data = await responce.json();
    let user_info = document.querySelector(".main");
    user_info.innerHTML=""
     
    data.map(ele => {
      let date = new Date(ele.CreatedAt)
    //  console.log(date.getHours());
    // Dislikes     int
    // Likes        int
      let contents = document.createElement("div")
      contents.innerHTML = `
        <div class="post">
         <div class="post-header">
          <img src="../static/imgs/profilePic.png" class="avatar" alt="Profile picture" />
          <div class="user-info">
            <div class="display-name">${ele.FirstName + " " + ele.LastName}</div>
            <span class="username">@aoc.bsky.social</span>
            <span class="timestamp">· ${date.getHours()}h</span>
          </div>
        </div>
        <div class="post-content">
            ${ele.Content}
        </div>
        <div class="post-actions">
          <div class="action">
            <svg width="18" height="18" viewBox="0 0 24 24" fill="currentColor">
              <path
                d="M12 21.638h-.014C9.403 21.59 1.95 14.856 1.95 8.478c0-3.064 2.525-5.754 5.403-5.754 2.29 0 3.83 1.58 4.646 2.73.814-1.148 2.354-2.73 4.645-2.73 2.88 0 5.404 2.69 5.404 5.755 0 6.376-7.454 13.11-10.037 13.157H12z" />
            </svg>
            <span>${ele.Likes} liked </span>
          </div>
           <div class="action">
            <svg width="18" height="18" viewBox="0 0 24 24" fill="currentColor">
              <path
                d="M12 21.638h-.014C9.403 21.59 1.95 14.856 1.95 8.478c0-3.064 2.525-5.754 5.403-5.754 2.29 0 3.83 1.58 4.646 2.73.814-1.148 2.354-2.73 4.645-2.73 2.88 0 5.404 2.69 5.404 5.755 0 6.376-7.454 13.11-10.037 13.157H12z" />
            </svg>
            <span>${ele.Dislikes} disliked</span>
          </div>
          <div class="action">
            <svg width="18" height="18" viewBox="0 0 24 24" fill="currentColor">
              <path
                d="M14.046 2.242l-4.148-.01h-.002c-4.374 0-7.8 3.427-7.8 7.802 0 4.098 3.186 7.206 7.465 7.37v3.828c0 .108.044.286.12.403.142.225.384.347.632.347.138 0 .277-.038.402-.118.264-.168 6.473-4.14 8.088-5.506 1.902-1.61 3.04-3.97 3.043-6.312v-.017c-.006-4.367-3.43-7.787-7.8-7.788zm3.787 12.972c-1.134.96-4.862 3.405-6.772 4.643V16.67c0-.414-.335-.75-.75-.75h-.396c-3.66 0-6.318-2.476-6.318-5.886 0-3.534 2.768-6.302 6.3-6.302l4.147.01h.002c3.532 0 6.3 2.766 6.302 6.296-.003 1.91-.942 3.844-2.514 5.176z" />
            </svg>
            <span>27.1K</span>
          </div>
        </div>
         </div>
        `
      user_info.appendChild(contents)
    })
   // console.log(data);
  } else {
    let data = responce.json()
     console.log(data);

  }
}
fetchData()
if (document.cookie) {
  let join = document.querySelector(".join")
  join.style.display = "none"
  let aside_nav = document.querySelector(".aside-nav")
  aside_nav.style.display = "block"
  while (join.firstChild) {
    join.removeChild(join.firstChild)
  }
 // location.href="/home"
  let tokens = document.cookie.split("; ")
  let token = null;
  let userId = null;
  tokens.forEach(ele => {
    let [key, value] = ele.split("=")
    if (key === 'token') {
      token = value
    }
    else if (key === "user_id") {
      userId = value
    }
  })
 // console.log(token, userId);

} else {
  let join = document.querySelector(".join")
  join.style.display = "block"
  let aside_nav = document.querySelector(".aside-nav")
  aside_nav.style.display = "none"
  while (aside_nav.firstChild) {
    aside_nav.removeChild(aside_nav.firstChild)
  }

}
