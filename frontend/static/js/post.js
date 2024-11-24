let form = document.querySelector('#form-submit')
 if( document.cookie.includes("token")){
   let tokens=document.cookie.split(";")
   let token=""
   let user_id=""
   tokens.map(elem=>{
    let t=elem.split("token=")
    let u=elem.split("user_id=")
    
})
console.log(t[1],u[1]);
    
 }else{
    console.log("no have any token");
    
 }


form.addEventListener("submit", async (e) => {
    e.preventDefault()
    let title = document.querySelector("#title")
    let content = document.querySelector("#content")
    let container = Array.from(document.querySelectorAll(".item-label"))
    let category = []
    container.map(ele => {
        category.push(ele.textContent)
    })
    const response = await fetch("/api/post", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            'Accept': 'application/json',
        },
        body: JSON.stringify({
            title: title.value,
            content: content.value,
            name: category
        })
    })

    if (response.ok) {
        const data = await response.json(); 
         console.log();
        console.log("Success:", data);
       
        } else {
        const errorData = await response.json();
        console.error("Error:", errorData);
        alert(`Error: ${errorData.message || "Request failed"}`);
    }


})