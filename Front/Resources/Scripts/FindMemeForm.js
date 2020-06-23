$(document).ready(function () {
    $("#findMemeForm").submit(function(event){
        event.preventDefault();

        $.ajax({
            url : $(this).attr("action"),
            type: $(this).attr("method"),
            data : $(this).serialize()
        }).done(function(response){
            $("#findResults").html("<img id='foundedMeme' src=''>");
            $("#foundedMeme").attr('src', `data:image/webp;base64,${response}`);
        }).fail(function(response) {
            alert(response.getResponseHeader('err'));
        });

        $(this).trigger("reset");
    });
});