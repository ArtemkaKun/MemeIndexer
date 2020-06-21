$(document).ready(function () {
    $("#addMemeForm").submit(function(event){
        event.preventDefault();

        let fileData = $("#memeFile").prop("files")[0];

        if (fileData.size > 3145728) {
            alert("Файл слишком большой! Доспустимы файлы < 3МБ");
            return;
        }

        let formData = new FormData();

        formData.append("file", fileData)
        formData.append("mainTags", $("#mainTags").prop("value"));
        formData.append("associationTags", $("#associantionTags").prop("value"))

        $.ajax({
            url : $(this).attr("action"),
            type: $(this).attr("method"),
            contentType: false,
            processData: false,
            data : formData
        }).done(function () {
            alert("Мем успешно добавлен");
        }).fail(function() {
            alert("При добавлении мема произошла ошибка");
        });

        $(this).trigger("reset");
    });
});