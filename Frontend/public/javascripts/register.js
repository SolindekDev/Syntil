const form = document.querySelector("#form");

form.addEventListener("submit", async () => {
    const formData = new FormData(form);
    fetch("htts://localhost:80/api/1v/account/register", {
        method: "POST",
        headers: {
            "Content-Type": "application/x-www-form-urlencoded",
        },
        body: new URLSearchParams(formData),
    }).then(res => res.json())
    .then(body => { 
        console.log(body);
        if (body.status == "400") {
            // Error
        } else {
            // Successfully response 
        }
    })
});