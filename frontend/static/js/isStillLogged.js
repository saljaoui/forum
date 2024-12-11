export async function isLogged() {
    const responce = await fetch("/api/isLogged", {
        method: "GET",
    });
    if (responce.ok) {
        if (!responce.message) {
        const response = await fetch("/api/logout", {
            method: "POST",
        });
        }
    } else {
        const response = await fetch("/api/logout", {
            method: "POST",
    })
}
}