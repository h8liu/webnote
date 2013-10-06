
function save(name, dat) {
    $.post("api/" + name, dat, function (data) { } );
}

function load(name, f) {
    $.get("api/" + name, "", f );
}

function today() {
    return "2013-10-05";
}

var testText = "这个似乎还是保序的……";
function main() {
    save(today(), testText);
    load(today(), function(dat) {
        $("#main").text(dat);
    });
}

$(document).ready(main);
