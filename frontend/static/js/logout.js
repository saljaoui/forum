import { status } from "./status";

const LogoutItem = document.querySelector(".signOut");

export default async function logout() {
    let userId = localStorage.getItem("user_id");
    const response = await fetch("http://localhost:3333/api/logout", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({ id: +userId }),
    });
    
    if (response.ok) {
        console.log("Logout successful");
        localStorage.clear();
        window.location.href = "/login";
    }else if (!response.ok) {
        status(response)
      }
}

if (LogoutItem) {
    LogoutItem.addEventListener("click", () => {
        logout();
    });
} else {
    console.error("Logout button not found");
}