const navItems = document.querySelectorAll('.nav-item');

const creatPostPopup = document.getElementById('creatPost-popup')
const openCategorie = document.getElementById('categories-popup')

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

function openCategories() {
    openCategorie.style.display = "flex"
}

function closeCategories() {
    openCategorie.style.display = "none"
}

const categoryItems = document.querySelectorAll('.category-item');
categoryItems.forEach(item => {
  item.addEventListener('click', () => {
    item.classList.toggle('selected');
  });
});