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
        .fail(function() {
            alert("Login or password incorrect!");
        });

    // $.get("/login", {login: userLogin, pass: userPass})
    //     .done(function (data) {
    //         if (data === "login") {
    //             CheckPass();
    //         }
    //     })
    //     .fail(function() {
    //         alert("Admin login incorrect!");
    //     });
}

function CheckPass() {
    let adminPass = prompt("Admin pass");

    if (adminPass === "") {
        return;
    }

    $.get("/adminPass", {pass: adminPass})
        .done(function (data) {
            if (data === "pass") {
                $("#addMemeForm").show();
                $("#findMemeForm").hide()
            }
        })
        .fail(function() {
            alert("Admin password incorrect!");
        });
}