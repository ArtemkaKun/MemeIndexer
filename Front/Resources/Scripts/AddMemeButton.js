$(document).ready(function () {
    $("#addMemeButton").click(function () {
        CheckLogin();
    });
});

function CheckLogin() {
    let userLogin = prompt("Admin login");

    if (userLogin === "") {
        return;
    }

    let userPass = prompt("Admin pass");

    if (userPass === "") {
        return;
    }

    $.get("/userAuth", {login: userLogin, pass: userPass})
        .done(function () {
            $("#addMemeForm").show();
            $("#findMemeForm").hide()
        })
        .fail(function (response) {
            alert(response.responseText);
        });
}