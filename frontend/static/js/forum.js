// import addpost from './post.js';
let containr = document.querySelector(".containr")
async function fetchData() {
    const responce = await fetch("http://localhost:3333/api/home", { method: "GET" });
    if (responce.ok) {
        const user_data = history.state

        let data = await responce.json();
        let user_info = document.querySelector(".content_post");
        // let post_content = document.querySelector(".post-content");
        data.map(ele => {
            let contents = document.createElement("div")
            //     let username = document.createElement("span")
            //     let created = document.createElement("span")
            //     let content = document.createElement("div")
            //     let a = document.createElement("a") 
            //     full_name.className="display-name"
            //     username.className="username"
            //     created.className="timestamp"
            //     full_name.textContent=ele.FirstName+" "+ele.LastName
            //     username.textContent="omrharbi"
            //     created.textContent=".11"
            //     content.textContent=ele.Content
            //     user_info.append(full_name,user_data,created)
            //     post_content.appendChild(content)
            contents.innerHTML = `
              <div class="post">
         <div class="post-header">
          <img src="../static/imgs/profilePic.png" class="avatar" alt="Profile picture" />
          <div class="user-info">
            <div class="display-name">${ele.FirstName + " " + ele.LastName}</div>
            <span class="username">@aoc.bsky.social</span>
            <span class="timestamp">· 11h</span>
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
            <span>862</span>
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


        console.log(data);


    } else {
        let data = responce.json()
        console.log(data);

    }
}
document.addEventListener("DOMContentLoaded", fetchData);

if (document.cookie) {
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
    console.log(token, userId);

} else {
    // const container = document.querySelector(".form");
    // while (container.firstChild) {
    //     container.removeChild(container.firstChild)
    // }
}



export {
    fetchData
}