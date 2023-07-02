function upload_company() {
    let name = document.querySelector('#name').value;
    let link_on_company_site = document.querySelector('#link_on_company_site').value;
    let max_loan_amount = document.querySelector('#max_loan_amount').value;
    let max_loan_duration = document.querySelector('#max_loan_duration').value;
    const min_loan_percent = document.querySelector('#min_loan_percent').value;

    const logo = document.querySelector('#logo').files;
    const logo_filename = document.querySelector('#logo_filename').value.split(",");

    const formData = new FormData();

    formData.append('name', name);
    formData.append('link_on_company_site', link_on_company_site);
    formData.append('max_loan_amount', max_loan_amount);
    formData.append('max_loan_duration', max_loan_duration);
    formData.append('min_loan_percent', min_loan_percent);

    formData.append('logo', logo[0]);
    formData.append('logo_filename', logo_filename[0])

    fetch('http://localhost/admin/panel/companies/upload/', {
        method: 'POST',
        body: formData,
        redirect: "follow"

    }).then(response => {
        if (response.redirected) {
            alert("Компания добавлена")
            window.location.href = response.url
        } else if (response.ok) {
            alert("Компания добавлена")
        } else {
            alert(response.body.getReader().read())
        }
    }).catch(err => {
        alert(err)
    })

}