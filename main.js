
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
    return "2013-10-05";
}

function saveMain() {
    save(today(), $("#main").val());
}

function showInMain(dat) {
    $("#main").val(dat);
}

function main() {
    load(today(), showInMain);
    $("#main").keyup(saveMain);
    $("#main").change(saveMain);
    $("#main").focus();
}

$(document).ready(main);
