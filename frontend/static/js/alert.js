

export function alertPopup(data) {
    console.log(data,"hhere");
    let alert=document.querySelector(".alert")
    console.log(alert.textContent);
    
    alert.innerHTML+=`
        <div class="popup-container" id="popup">
            <div class="popup-message">
            <div class="success-icon">
                <i class="fas fa-check"></i>
            </div>
            <div class="message-content">
                <h3 class="meesage">Error</h3>
                <p class="error">${data.message}</p>
            </div>
            </div>
        </div>
    `
    function showPopup() {
        const popup = document.getElementById('popup');
        popup.classList.add('show');
    }

    function hidePopup() {
        const popup = document.getElementById('popup');
        popup.classList.remove('show');
         location.href="/home"
    }
    showPopup()
    setTimeout(hidePopup, 900);

} 