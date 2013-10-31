
function save(name, dat) {
    $.post("api/" + name, dat, function (data) { } );
}

function load(name, f) {
    $.ajax({
        url: "api/" + name, 
        data: "",
        success: f,
        cache: false
    });
}

function today() {
    var now = new Date();
    var date = now.getDate();
    var month = now.getMonth() + 1;
    var year = now.getFullYear();
    
    function s(i) {
        var ret = "" + i;
        while (ret.length < 2) {
            ret = "0" + ret;
        }
        return ret;
    }

    var ret = year + "-" + s(month) + "-" + s(date);
    return ret;
}

function saveMain() {
    save(today(), $("#main").val());
}

function showInMain(dat) {
    $("#main").val(dat);
}

function main() {
    var t = today();
    load(t, showInMain);
    $("#main").keyup(saveMain);
    $("#main").change(saveMain);
    $("#main").focus();
    $("#date").html(t);
}

$(document).ready(main);
