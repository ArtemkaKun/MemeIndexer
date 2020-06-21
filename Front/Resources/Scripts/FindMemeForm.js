$(document).ready(function () {
    $("#findMemeForm").submit(function(event){
        event.preventDefault();

        $.ajax({
            url : $(this).attr("action"),
            type: $(this).attr("method"),
            data : $(this).serialize()
        }).done(function(response){ //
            $("#findResults").html("<img id='foundedMeme' src=''>");
            response
            $("#foundedMeme").attr('src', `data:image/webp;base64,${response}`);
        });

        $(this).trigger("reset");
    });
});