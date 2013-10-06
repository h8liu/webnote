
function save(name, dat) {
    $.post("api/" + name, dat, function (data) { } );
}

function load(name, f) {
    $.get("api/" + name, "", f );
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
}

$(document).ready(main);
