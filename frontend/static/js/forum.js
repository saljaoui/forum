import { checkandAdd } from "./addlikes.js";
import { likes } from "./likescomment.js";
import { cards } from "./card.js";
import { checklogin } from "./checklogin.js";
import { search } from "./search.js";
let content = []
checklogin()
const searchInput = document.querySelector("[data-search]")
if (searchInput) {
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
}

export async function fetchData(page=1) {
  console.log(page);
  
  const responce = await fetch(`/api/home?page=${page}`, {
    method: "GET",
  });
  if (responce.ok) {
    let path = window.location.pathname
    if (path !== "/profile") {


      let data = await responce.json();
      console.log(data);
      
      let user_info = document.querySelector(".main");
      content = cards(data.posts, user_info)

      let like = document.querySelectorAll("#likes");
      likes(like)
      search(content)
      renderPagination(data, user_info);
    }
  }
  // else if (responce.status === 401) {
  //   let body = document.querySelector("body")
  //   body.style.display = "none"
  //    //location.href = "/login"
  // }

}
fetchData()
document.addEventListener("DOMContentLoaded", () => {
  checkandAdd();
});


function renderPagination(data, container) {
  // Create pagination container if it doesn't exist
  let paginationDiv = document.querySelector('.pagination-controls');
  if (!paginationDiv) {
    paginationDiv = document.createElement('div');
    paginationDiv.className = 'pagination-controls';
    container.appendChild(paginationDiv);
  }
  
  let paginationHTML = '';
  
  // Previous button
  if (data.currentPage > 1) {
    paginationHTML += `
      <button onclick="window.fetchData(${data.currentPage - 1})" class="pagination-btn">
        Previous
      </button>
    `;
  }
  
  // Page numbers
  for (let i = 1; i <= data.totalPages; i++) {
    paginationHTML += `
      <button 
        onclick="window.fetchData(${i})" 
        class="pagination-btn ${i === data.currentPage ? 'active' : ''}"
      >
        ${i}
      </button>
    `;
  }
  
  // Next button
  if (data.currentPage < data.totalPages) {
    paginationHTML += `
      <button onclick="window.fetchData(${data.currentPage + 1})" class="pagination-btn">
        Next
      </button>
    `;
  }
  
  paginationDiv.innerHTML = paginationHTML;
}

// Make fetchData available globally
window.fetchData = fetchData;
// if (document.cookie) {
//   let join = document.querySelector(".join");
//   join.style.display = "none";
//   let aside_nav = document.querySelector(".aside-nav");
//   aside_nav.style.display = "block";
//   while (join.firstChild) {
//     join.removeChild(join.firstChild);
//   }
//   // location.href="/home"
//   let tokens = document.cookie.split("; ");
//   let token = null;
//   let userId = null;
//   tokens.forEach((ele) => {
//     let [key, value] = ele.split("=");
//     if (key === "token") {
//       token = value;
//     } else if (key === "user_id") {
//       userId = value;
//     }
//   });
//   // console.log(token, userId);
// } else {

//   let join = document.querySelector(".join");
//   join.style.display = "block";
//   let aside_nav = document.querySelector(".aside-nav");
//   aside_nav.style.display = "none";
//   while (aside_nav.firstChild) {
//     aside_nav.removeChild(aside_nav.firstChild);
//   }
// }