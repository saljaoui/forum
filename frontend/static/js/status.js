export async function status(response) {
    let data = await fetch("/api/err", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            "Accept": "application/json",
        },
        body: JSON.stringify({
            code: response.status
        })
    });

    if (!data.ok) {
        let re = await data.json()
        window.history.pushState(
            { data: re, code: response.status }, // State object
            "",                                  // Title (optional, not used here)
            `/err`                               // URL for error page
        );
   
        location.href="/err"
    } else {
        console.error("Failed to send status code:", data.status, data.statusText);
    }

}
