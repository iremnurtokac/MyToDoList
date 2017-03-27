function drawfield(result) {
    $("#target").html('<table class="table"><tr>' +
        '<td><taskbutton val="1"  class="btn btn-default">' + result.field[0] + "</taskbutton></td><tr>" +
        '<td><taskbutton val="2"  class="btn btn-default">' + result.field[1] + "</taskbutton></td><tr>" +
        '<td><taskbutton val="3"  class="btn btn-default">' + result.field[2] + "</taskbutton></td><tr>" +
        '<td><taskbutton val="4"  class="btn btn-default">' + result.field[3] + "</taskbutton></td><tr>" +
        '<td><taskbutton val="5"  class="btn btn-default">' + result.field[4] + "</taskbutton></td><tr>" +
        '<td><taskbutton val="6"  class="btn btn-default">' + result.field[5] + "</taskbutton></td><tr>" +
        '<td><taskbutton val="7"  class="btn btn-default">' + result.field[6] + "</taskbutton></td><tr>" +
        '<td><taskbutton val="8"  class="btn btn-default">' + result.field[7] + "</taskbutton></td><tr>" +
        '<td><taskbutton val="9"  class="btn btn-default">' + result.field[8] + "</taskbutton></td><tr>"
        + "</tr> </table><div>" + result.currentPlayer + "</div>" + "<div>" + result.message + "</div>");
   
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