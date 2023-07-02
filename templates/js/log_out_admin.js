function log_out() {
    fetch("http://localhost/admin/panel/sign-out/", {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        }
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