function likes(likes,dislikes){
    likes.addEventListener("click",()=>{
        console.log("hello");

    })
    
    let response =fetch("/api/like",{
        method:"POST",
        body:{
            is_liked:1,
            card_id:214
          }
    })
}
  export{
    likes
 }