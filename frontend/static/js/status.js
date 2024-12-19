export async function status(response) {
  let d=await response.json();
    console.log(d);
    
    let data = await fetch("/api/err", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            "Accept": "application/json",
        },
        body: JSON.stringify({
            code: response.status,
            msg: d.message
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
    }  

}
