export async function isLogged() {
    const responce = await fetch("/api/isLogged", {
        method: "GET",
      });
      if (responce.ok) {
        let path = window.location.pathname
        if (path !== "/profile") {
    
    
          let data = await responce.json();
          let user_info = document.querySelector(".main");
          content = cards(data, user_info)
    
          let like = document.querySelectorAll("#likes");
          likes(like)
          search(content)
        }
      }
}