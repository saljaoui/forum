let register = document.querySelector("#form-submit")
  register.addEventListener('submit', async (e) => {
    e.preventDefault()
    let firstname = document.getElementById('firstname').value
    let lastname = document.getElementById('lastname').value
    let emailRegister = document.getElementById('emailRegister').value
    let passwordRegister = document.getElementById('passwordRegister').value

    const response = await fetch("/api/register", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            'Accept': 'application/json',
        },

        body: JSON.stringify({
            firstname: firstname,
            lastname: lastname,
            email: emailRegister,
            password: passwordRegister
        })
    })
    // if (response.status === 401) {
    //     const data = await response.json();
    //     localStorage.setItem("token",data.token)
    //     console.log("Success:", data);
    // }
    if (response.ok) {
        const data = await response.json();
        console.log("Success:", data);
        //localStorage.setItem("user_id",data)
    } else {
        const errorData = await response.json();
        console.error("Error:", errorData);
        alert(`Error: ${errorData.message || "Request failed"}`);
    }

})


let login = document.querySelector("#login")

login.addEventListener('submit', async (e) => {
    e.preventDefault()
    let email = document.querySelector('#email').value
    let password = document.querySelector('#password').value
    const response = await fetch("/api/login", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            'Accept': 'application/json',
        },
        body: JSON.stringify({
            email: email,
            password: password
        })
    })
    if (response.ok) {
        const data = await response.json();
        localStorage.setItem("token",data.message.uuid)
         window.history.pushState({user:data},"","/home")
        location.href="/home"
         console.log("Success:", data);
     } else {
        const errorData = await response.json();
        console.error("Error:", errorData);
        alert(`Error: ${errorData.message || "Request failed"}`);
    }

}) 
 