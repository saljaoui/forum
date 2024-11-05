

document.getElementById('fetchDataBtn').addEventListener('click', function() {

    fetch('http://localhost:3333/')
        .then(response => response.json())
        .then(data => {
            const responseDataDiv = document.getElementById('responseData');
            responseDataDiv.innerHTML = `
                <p><strong>Message:</strong> ${data.message}</p>
                <p><strong>Status:</strong> ${data.status}</p>
            `;
        })
        .catch(error => {
            console.error('Error fetching data:', error);
        });
});