import { fetchData } from './forum.js';
export default function addpost() {
    const form = document.querySelector("#form-submit");
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
            fetchData()
            const data = await response.json();
            console.log("Success:", data);
             

        } else {
            const errorData = await response.json();
            console.error("Error:", errorData);
            alert(`Error: ${errorData.message || "Request failed"}`);
        }

    })
} 