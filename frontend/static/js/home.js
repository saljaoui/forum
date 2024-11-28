// import div from './asideRight.js';
 const navItems = document.querySelectorAll('.nav-item');

const creatPostPopup = document.getElementById('creatPost-popup')

function activeByDefault() {
    navItems.forEach(navItem => {
        const outlineIcon = navItem.querySelector('ion-icon[name$="-outline"], ion-icon[name$="-sharp"]');
        const filledIcon = navItem.querySelector('ion-icon:not([name$="-outline"]):not([name$="-sharp"])');
        filledIcon.classList.remove('active');
        outlineIcon.classList.add('active');
        navItem.classList.remove('active')
    });
}

navItems.forEach(navItem => {
    navItem.addEventListener('click', function() {
        activeByDefault();
        navItem.classList.add('active')
        const outlineIcon = this.querySelector('ion-icon[name$="-outline"], ion-icon[name$="-sharp"]');
        const filledIcon = this.querySelector('ion-icon:not([name$="-outline"]):not([name$="-sharp"])');
        if (outlineIcon.classList.contains('active')) {
            outlineIcon.classList.remove('active');
            filledIcon.classList.add('active');
        }
    });
});

function openCreatPost() {
    creatPostPopup.style.display = "flex"
}
function closeCreatPost() {
    creatPostPopup.style.display = "none"
}


// let side =document.querySelector("#side-right")

// console.log(div());
