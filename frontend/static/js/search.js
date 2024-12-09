export function search(content) {
console.log(content);

    const searchInput = document.querySelector("[data-search]")
    searchInput.addEventListener("input", (e) => {
        const value = e.target.value.toLowerCase()
        content.forEach(data => {
            const isVisible = data.data.toLowerCase().includes(value)
            if (!isVisible) {
                data.element.style.display = "none"
            } else {
                data.element.style.display = "block"
            }
        })
    })
}