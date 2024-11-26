const urlParams = new URLSearchParams(window.location.search);
const id = urlParams.get("target_id");
let comment = document.querySelector(".comments")
async function fetchdata(){
  const response= await fetch(`/api/comment?target_id=${id}`, {
        method: "GET",
    });
    
    if(response.ok){
        const data =await response.json();       
        data.forEach(element => {
            console.log(element);
            
        });
        let div = document.createElement("div")
        let title = document.createElement("h2")
        let content = document.createElement("p")
     
    }
    else{
        console.log("err");
        
    }
}
fetchdata()