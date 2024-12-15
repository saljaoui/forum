export async function status(response) {
    let data = await response.json()
    window.history.pushState(
        { data: data, code: response.status }, // State object
        "",                                  // Title (optional, not used here)
        `/err?code=${response.status}`       // URL for error page
    );

    location.href = `/err?code=${response.status}`
    return;
}