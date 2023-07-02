function upload_review() {
    let reviewer_name = document.querySelector('#reviewer_name').value;
    let reviewer_phone = document.querySelector('#reviewer_phone').value;
    let review = document.querySelector('#review').value;
    let moderated = document.querySelector('#moderated').checked;

    if (moderated) {
        fetch("http://localhost/admin/panel/reviews/upload/", {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                "reviewer_name": reviewer_name,
                "reviewer_phone": reviewer_phone,
                "review": review,
                "moderated": moderated,
            })
        })
            .then(res => {
                if (res.redirected) {
                    document.location = res.url
                } else if (res.status === 200) {
                    alert("Данные об отзыве обновлены")
                } else if (res.status !== 200) {
                    alert(res.body)
                }
            })
            .catch(err => {
                alert(err)
            })
    } else {
        alert("отметьте кнопку 'Проверено'")
    }
}