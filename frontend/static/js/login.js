// const wrapper = document.querySelector('.wrapper');
// const loginLink = document.querySelector('.login-link');
// const registerLink = document.querySelector('.register-link');

import { status } from "./status.js"

// registerLink.addEventListener('click', () => {
//   wrapper.classList.add('active');
// });

// loginLink.addEventListener('click', () => {
//   wrapper.classList.remove('active');
// });

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
        console.log(data);
        location.href = "/home"
    } else if (!response.ok) {
        status(response)
    } else {
        const errorData = await response.json();
        console.error("Error:", errorData);
        alert(`Error: ${errorData.message || "Request failed"}`);
    }

})