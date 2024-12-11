export function checklogin() {
    console.log("error here 2");
    const value = getcookies()
    const token = value[0]
    const userId = value[1]
    //console.log(userId);
    if (token != null && userId != null) {
        let aside_nav = document.querySelector(".aside-nav");
        aside_nav.style.display = "block";
        
        let join = document.querySelector(".join");
        join.style.display = "none";
        while (join.firstChild) {
            join.removeChild(join.firstChild);
        }
    } else {
        //console.log("error here 1");
        let join = document.querySelector(".join");
        join.style.display = "block";
        let aside_nav = document.querySelector(".aside-nav");
        aside_nav.style.display = "none";
        while (aside_nav.firstChild) {
            aside_nav.removeChild(aside_nav.firstChild);
        }
        let post_comment = document.querySelector(".postReply");
        if (post_comment) {
            post_comment.remove()
        }
    }
}

function getcookies() {
    
    
    let tokens = document.cookie.split("; ");
    let token = null;
    let userId = null;
    tokens.forEach((ele) => {
        let [key, value] = ele.split("=");
        if (key === "token") {
            token = value;
        } else if (key === "user_id") {
            userId = value;
        }
    });
    console.log(token,userId);
    
    return [token, userId]
}