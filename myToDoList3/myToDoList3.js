function drawfield(result) {
    var content = "";
    for (i = 0; i < result.task.length; i++) {
        content += '<div class="list-group-item"><input val="' + i + '"  type="checkbox">' + result.task[i] + "</input></div>";
    }

    $("#target").html('<div class="list-group">' + content + "</div>");

}


$(document).ready(function () {
    $("#addTaskButton").click(function (event) {
        $.ajax({
            type: 'GET',
            crossOrigin: true,
            url: 'http://localhost:8080/tasks/',
            success: function (result) {

                drawfield(result);
             

            },
            error: function () {
                $("#target").html('ERROR');
            }

        });


    });
    $("#addTaskbutton").click(function(){
    $.post("http://localhost:8080/tasks/add",
    {
        //task: "Donald Duck"
        task: $("#mytextfield").val() 
    },
    function(result){
        drawfield(result);
    });
});



});