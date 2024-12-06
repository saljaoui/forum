import { likes } from "./likes.js";
import { cards } from "./card.js";
const user_data = localStorage.getItem("user_id");
let content = []
const searchInput = document.querySelector("[data-search]")
searchInput.addEventListener("input", (e) => {
  const value = e.target.value.toLowerCase()
  content.forEach(data => {
    const isVisible = data.data.toLowerCase().includes(value)
    if (!isVisible) {
      data.element.style.display = "none"
    } else {
      data.element.style.display = "block"
    }
  })
})


export function NewPost(postInfo,parent) {
  let parentElem = document.querySelector(parent)
  let post = document.createElement("div");
  post.className = "post"
  post.innerHTML = `
      <div class="post-header">
              <img src="../static/imgs/profilePic.png"
                class="avatar" alt="Profile picture" />
                    <div class="user-info">
                        <div class="display-name">${postInfo.firstName + " " + postInfo.lastName}</div>
                              <span class="username">@${postInfo.firstName}</span>
                              <span class="timestamp">1h</span>
                        </div>
                    </div>
                    <div class="post-content">
                        ${postInfo.content}
                    </div>
                    <div class="post-actions">
                         <div class="action active is_liked"  id="likes" data-liked="false" data-like="like" data-id_card="${postInfo.id}" >
                              <svg width="20" height="20" viewBox="0 0 20 20"
                                  fill="currentColor">
                                   <path
                                        d="M10 19c-.072 0-.145 0-.218-.006A4.1 4.1 0 0 1 6 14.816V11H2.862a1.751 1.751 0 0 1-1.234-2.993L9.41.28a.836.836 0 0 1 1.18 0l7.782 7.727A1.751 1.751 0 0 1 17.139 11H14v3.882a4.134 4.134 0 0 1-.854 2.592A3.99 3.99 0 0 1 10 19Zm0-17.193L2.685 9.071a.251.251 0 0 0 .177.429H7.5v5.316A2.63 2.63 0 0 0 9.864 17.5a2.441 2.441 0 0 0 1.856-.682A2.478 2.478 0 0 0 12.5 15V9.5h4.639a.25.25 0 0 0 .176-.429L10 1.807Z"></path>
                                </svg>
                                <span id="is_liked">${postInfo.likes}</span>
                            </div>
                            <div class="action " id="dilike">
                                <svg width="20" height="20" viewBox="0 0 20 20"
                                    fill="currentColor">
                                    <path
                                        d="M10 1c.072 0 .145 0 .218.006A4.1 4.1 0 0 1 14 5.184V9h3.138a1.751 1.751 0 0 1 1.234 2.993L10.59 19.72a.836.836 0 0 1-1.18 0l-7.782-7.727A1.751 1.751 0 0 1 2.861 9H6V5.118a4.134 4.134 0 0 1 .854-2.592A3.99 3.99 0 0 1 10 1Zm0 17.193 7.315-7.264a.251.251 0 0 0-.177-.429H12.5V5.184A2.631 2.631 0 0 0 10.136 2.5a2.441 2.441 0 0 0-1.856.682A2.478 2.478 0 0 0 7.5 5v5.5H2.861a.251.251 0 0 0-.176.429L10 18.193Z"></path>
                                </svg>
                                <span id="is_Dislikes" data-disliked="disliked">${postInfo.dislikes}</span>
                            </div>
                            <div class="action">
                      <svg width="20" height="20" viewBox="0 0 20 20"
                         fill="currentColor">
                     <path
                  d="M10 19H1.871a.886.886 0 0 1-.798-.52.886.886 0 0 1 .158-.941L3.1 15.771A9 9 0 1 1 10 19Zm-6.549-1.5H10a7.5 7.5 0 1 0-5.323-2.219l.54.545L3.451 17.5Z"></path>
              </svg>
           <span>${postInfo.comments}</span>
         </div>
       </div>`

    let anotherPost = document.createElement("div")
    anotherPost.className = "post"
    let postReply = document.createElement("div")
    postReply.className = "postReply"
    postReply.innerHTML = `<img src="../static/imgs/profilePic.png"
                                class="avatar" alt="Profile picture" />
                              <div class="writeReply" >Write your reply</div>`
    anotherPost.appendChild(postReply)
    console.log("parent is : ",parentElem);
    
    parentElem.appendChild(post)
    parentElem.appendChild(anotherPost)
}


export   async function fetchData() {
  const responce = await fetch("/api/home", {
    method: "GET",
  });
  if (responce.ok) {
 
    let data = await responce.json();
    let user_info = document.querySelector(".main");
    
    content = cards(data,user_info)
    let like = document.querySelectorAll("#likes");
    likes(like)
  } else if (responce.status === 401) {
    let body = document.querySelector("body")
    body.style.display = "none"
     location.href = "/login"
  }
}
fetchData()