$(document).ready(function () {
    $("#addMemeButton").click(function () {
        CheckLogin();
    });
});

function CheckLogin() {
    let userLogin = prompt("Логин");

    if (userLogin === "") {
        return;
    }

    let userPass = prompt("Пароль");

    if (userPass === "") {
        return;
    }

    $.get("/userAuth", {login: userLogin, pass: userPass})
        .done(function () {
            $("#foundedMeme").attr('src', ``);
            $("#addMemeForm").show();
            $("#findMemeForm").hide()
        })
        .fail(function (response) {
            alert(response.responseText);
        });
}