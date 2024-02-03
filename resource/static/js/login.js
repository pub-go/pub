(function () {
    (function () {
        const form = document.getElementById('login-form')
        if (!form) { return }
        const fieldset = document.getElementById('fieldset')
        fieldset.disabled = false
        const plain = document.getElementById('plain')
        const passwd = document.getElementById('passwd')
        form.addEventListener('submit', e => {
            passwd.value = sha256(form.username.value + sha256(plain.value) + form.salt.value)
            return true
        })
    }());
}())
