function update_companies_priority() {
    fetch("http://localhost/admin/panel/companies/update/", {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify([
            {"id": 1, "priority": 2},
            {"id": 2, "priority": 1}
        ])
    })
        .then(res => {
            if (res.redirected) {
                document.location = res.url
            } else if (res.status === 200) {
                alert("данные обновлены")
            }
        })
        .catch(err => {
            alert(err)
        })
}