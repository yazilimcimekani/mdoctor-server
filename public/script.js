window.onload = () => {
    console.log('Serving your markdown files.\nHave a problem? https://github.com/yazilimcimekani/mdoctor-server/issues')
}

function switchTheme() {
    const body = document.body;
    const themeIcon = document.getElementById('theme-icon');

    body.classList.toggle('light-theme');
    themeIcon.classList.toggle('fa-sun');
    themeIcon.classList.toggle('fa-moon');
}
