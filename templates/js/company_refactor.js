const id = document.querySelector('#id').value

function update_company() {
    let name = document.querySelector('#name').value;
    let link_on_company_site = document.querySelector('#link_on_company_site').value;
    let max_loan_amount = document.querySelector('#max_loan_amount').value;
    let max_loan_duration = document.querySelector('#max_loan_duration').value;
    let min_loan_percent = document.querySelector('#min_loan_percent').value;
    let moderated = document.querySelector('#moderated').checked;

    if (moderated) {
        fetch("http://localhost/admin/panel/companies/"+id+"/", {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                "name": name,
                "link_on_company_site": link_on_company_site,
                "max_loan_amount": parseInt(max_loan_amount),
                "max_loan_duration": parseInt(max_loan_duration),
                "min_loan_percent": parseInt(min_loan_percent),
            })
        })
            .then(res => {
                if (res.redirected) {
                    document.location = res.url
                } else if (res.status === 200) {
                    alert("Данные об отзыве обновлены")
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

function delete_company() {
    fetch("http://localhost/admin/panel/companies/"+id+"/", {
        method: 'DELETE'
    })
        .then(res => {
            if (res.redirected) {
                alert("Данные о компании удалены")
                document.location = res.url
            } else if (res.status === 200) {
                alert("Данные о компании удалены")
            } else if (res.status !== 200) {
                alert(res.body)
            }
        })
        .catch(err => {
            alert(err)
        })
}