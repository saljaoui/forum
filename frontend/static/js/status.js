export async function status(response) {
    let statuscode = null
    let message = null
    if (response === 404) {
        statuscode = 404
        message = "Page Not Found"
     } if (typeof response === "function") {
        statuscode = response.status
        d = await response.json();
        message=d.message
    }

    let data = await fetch("/api/err", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            "Accept": "application/json",
        },
        body: JSON.stringify({
            code: statuscode,
            msg: message
        })
    });

    if (!data.ok) {
        let re = await data.json()
        window.history.pushState(
            { data: re, code: statuscode}, // State object
            "",                                  // Title (optional, not used here)
            `/err`                               // URL for error page
        );

        location.href = "/err"
    }

}
