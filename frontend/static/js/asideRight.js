export default function div() {
    const asidElement = document.querySelector("#asid");

    if (asidElement) {
        console.log(asidElement.textContent);  // This should print the content of the element with id "asid"
    } else {
        console.log("Element with id 'asid' not found.");
    }
}

document.addEventListener("DOMContentLoaded", function () {
    div();  // Ensure it's called when the DOM is fully loaded
});
