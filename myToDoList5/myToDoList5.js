function drawfield(result) {
    var content = "";
    for (i = 0; i < result.task.length; i++) {
        content += '<div class="list-group-item"><input val="' + i + '"  type="checkbox">' + result.task[i] + "</input></div>";
    }

    $("#target").html('<div class="list-group">' + content + "</div>");

}


$(document).ready(function () {
    $("#addTaskButton").click(function () {
        $.post("http://localhost:8080/tasks/add",
            {
                //task: "Donald Duck"
                task: $("#task").val()
            },
            function (result) {
                drawfield(result);
            });



    });

    $("#resetTaskButton").click(function () {
        $.post("http://localhost:8080/tasks/reset",
          {
                //task: "Donald Duck"
                task: $("#task").val()
            },
                      
            function (result) {
                drawfield(result);
        
        });


    });


});


