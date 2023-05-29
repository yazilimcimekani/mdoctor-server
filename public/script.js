window.onload = () => {
    console.log('Serving your markdown files.\nHave a problem? https://github.com/yazilimcimekani/mdoctor-server/issues')
}

function switchTheme() {
    let themeSwitchButton = document.getElementById('theme-switch-button')
    let body = document.getElementById('body')
    let isLightTheme = document.getElementById('body').classList.contains('light-theme') ? true : false

    if (isLightTheme) {
        body.classList.remove('light-theme')
        body.classList.add('dark-theme')
        themeSwitchButton.textContent = 'Light Theme'
    } else {
        body.classList.remove('dark-theme')
        body.classList.add('light-theme')
        themeSwitchButton.textContent = 'Dark Theme'
    }
}
