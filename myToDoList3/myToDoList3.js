function drawfield(result) {
    $("#target").html('<table class="table"><tr>' +
        '<td><taskbutton val="1"  class="checkbox">' + result.task[0] + "</taskbutton></td><tr>" +
        '<td><taskbutton val="2"  class="checkbox">' + result.task[1] + "</taskbutton></td><tr>" +
        '<td><taskbutton val="3"  class="checkbox">' + result.task[2] + "</taskbutton></td><tr>" +
        '<td><taskbutton val="4"  class="checkbox">' + result.task[3] + "</taskbutton></td><tr>" +
        '<td><taskbutton val="5"  class="checkbox">' + result.task[4] + "</taskbutton></td><tr>" +
        '<td><taskbutton val="6"  class="checkbox">' + result.task[5] + "</taskbutton></td><tr>" +
        '<td><taskbutton val="7"  class="checkbox">' + result.task[6] + "</taskbutton></td><tr>" +
        '<td><taskbutton val="8"  class="checkbox">' + result.task[7] + "</taskbutton></td><tr>" +
        '<td><taskbutton val="9"  class="checkbox">' + result.task[8] + "</taskbutton></td><tr>");
   
}


function myfunction() {
    $("taskbutton").click(function (event) {
        $.ajax({
            type: 'POST',
            crossOrigin: true,
            url: 'http://localhost:8080/users/' + $("#task").val() + '/move/' + $(event.target).attr("val"),
            success: function (result) {

                drawfield(result);
                myfunction();
            },
            error: function () {

            }
        });
    });
};



$(document).ready(function () {
    $("#task").keydown(function () {$("#target").html("");});
    $("#addTaskButton").click(function (event) {
        $.ajax({
            type: 'GET',
            crossOrigin: true,
            url: 'http://localhost:8080/users/' + $("#task").val(),
            success: function (result) {

                drawfield(result);
                myfunction();

            },
            error: function () {

            }

        });


    });
    $("#resetTaskButton").click(function (event) {
        $.ajax({
            type: 'POST',
            crossOrigin: true,
            url: 'http://localhost:8080/users/' + $("#task").val() + '/reset',
            success: function (result) {

                drawfield(result);
                myfunction();

            },
            error: function () {

            }

        });


    });
    
    
});