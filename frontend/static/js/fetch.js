
let button =document.querySelector("#form-submit")

button.addEventListener('submit',(e)=>{
    let firstname=document.getElementById('firstname').value
    let lastname=document.querySelector('#lastname').value
    let email=document.querySelector('#email').value
    let password=document.querySelector('#password').value
    const response =    fetch("/api/register",{
        method:"POST",
        headers:{
             "Content-Type": "application/json",
             'Accept': 'application/json',
        },
        
        body:JSON.stringify({
            firstname:firstname,
            lastname:lastname,
            email:email,
            password:password
        })
    }) ;
    
    response.then(response=>{
        if (response.ok){
            let data=response.json
            console.log(data);
            
        }
    }).catch(err=>{
        console.log(err);
        
    }) 
    

})
console.log(firstname,"ddsddd");

