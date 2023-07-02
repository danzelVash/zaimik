function update_subscription() {
    const id = 1;
    let new_expired_date = new Date(document.querySelector('#new_expired_date').value).toISOString();
    let moderated = document.querySelector('#moderated').checked;
    // alert(new_expired_date)
    if (moderated) {
        fetch("http://localhost/admin/panel/subscriptions/"+id+"/", {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                "id": id,
                "expired_date": new_expired_date
            })
        })
            .then(res => {
                if (res.redirected) {
                    document.location = res.url
                } else if (res.status === 200) {
                    alert("Данные о подписке обновлены")
                } else if (res.status !== 200) {
                    alert(res.json())
                }
            })
            .catch(err => {
                alert(err)
            })
    } else {
        alert("отметьте кнопку 'Проверено'")
    }


}