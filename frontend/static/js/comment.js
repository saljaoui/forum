const urlParams = new URLSearchParams(window.location.search);
const id = urlParams.get("target_id");
let comment = document.querySelectorAll(".comments")


comment.forEach((c) => {
    c.addEventListener('click', console.log("hello"))
})
async function fetchdata(){
  const response= await fetch(`/api/comment?target_id=${id}`, {
        method: "GET",
    });
    
    if(response.ok){
        const data =await response.json();      
        if (Array.isArray(data)) {

            data.forEach(element => {
                console.log("Comment:", element);                
            });
        } 
        let div = document.createElement("div")
        let title = document.createElement("h2")
        let content = document.createElement("p")
     
    }
    else{
        console.log("err");
        
    }
}
fetchdata()