function sign_in_admin() {
    let login = document.querySelector('#admin_login');
    let password = document.querySelector('#admin_password');

    fetch("http://localhost/admin/get-code/", {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({"admin_login": login.value, "admin_password": password.value})
    })
        .then(res => {
            if (res.redirected) {
                document.location = res.url
            } else if (res.status === 200) {
                let hid = document.querySelectorAll('.hidden')
                hid.forEach(h => h.style.display = 'block')
            } else if (res.status !== 200) {
                alert("Неправильные логин или пароль")
            }
        })
        .catch(err => {
            alert(err)
        })
}

function check_admin_auth_code() {
    let code = document.querySelector('#auth_code');

    fetch("http://localhost/admin/", {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({"code": code.value})
    })
        .then(res => {
            if (res.redirected) {
                document.location = res.url
            }
        })
        .catch(err => {
            alert(err)
        })
}