const toggle = document.getElementById('nav-toggle');
const menu = document.getElementById('nav-menu');
const burger = toggle.querySelector('.nav__burger');
const close = toggle.querySelector('.nav__close');

toggle.addEventListener('click', () => {
    menu.classList.toggle('show-menu');

    if(menu.classList.contains('show-menu')) {
        burger.style.opacity = '0';
        close.style.opacity = '1';
    } else {
        burger.style.opacity = '1';
        close.style.opacity = '0';
    }
});
