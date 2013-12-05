package main

var style = `
body {
    background: #222;
    margin: 0;
    padding: 0;
}

div.test {
    /* font-family: "Microsoft YaHei"; */
    font-size: 80%;
}

div#date {
    position: absolute;
    top: 50%;
    left: 50%;
    margin-left: -2.9in;
    margin-top: -3.92in;
    width: 5.6in;
    text-align: right;
    font-size: 14px;
    font-family: "Consolas", "monospace";
    color: #ccc;
}

textarea#main {
    position: absolute;
    top: 50%;
    left: 50%;
    margin-left: -2.9in;
    margin-top: -3.7in;
    padding: .2in;
    height: 7in;
    width: 5.4in;

    font-family: "Consolas", "Microsoft YaHei", "sans-serif";
    font-size: 14px;
    font-weight: normal;
    line-height: 22px;
    border: 0;
    border-radius: .15in;
    outline: none;
    background: #eee;
    color: #333;
}

form {
    margin: auto;
    text-align: center;
    vertical-align: middle;
}
`

var script = `
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
`
