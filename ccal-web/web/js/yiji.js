window.onload = function () {

    function jinshen() {
        var request = new XMLHttpRequest();
        request.onreadystatechange = function () {
            if (request.readyState === 4 && request.status === 200) {

                var info = request.responseText;
                var js = JSON.parse(info);

                var j1 = js.JinShen1;
                var j2 = js.JinShen2;
                var j3 = js.JinShen3;
                var jinsheninfo = j1 + "<br>" + j2 + "<br>" + j3;

                var p = document.getElementById("y1");
                p.innerHTML = jinsheninfo;

            };
        };
        // 发送请求:
        request.open('POST', '/yiji', true);
        request.send();
    }
    function taisuiCY() {
        var xhr = new XMLHttpRequest();
        xhr.onreadystatechange = function () {
            if (xhr.readyState == 4 && xhr.status == 200) {
                var info = xhr.responseText;
                var js = JSON.parse(info);
                var taisuichuyou = js.TaiSui;
                var p = document.getElementById("y1");
                p.innerHTML = taisuichuyou;
            };
        };
        xhr.open("POST", "/yiji", true);
        xhr.send();
    }
    function home() {
        window.open("/")
    }
    jinshen();
    taisuiCY();
    home();
}


/* document.getElementById("btny1").onclick = function jinshen() {
    var request = new XMLHttpRequest();
    request.open("POST", '/ccal', true);
    //var data=JSON.stringify(data);
    request.setRequestHeader("Content-type", "application/json");
    request.send();
    request.onreadystatechange = function () {
        if (request.readyState === 4) {
            if (request.status === 200) {
                var data = JSON.parse(request.responseText);
                console.log(data.JinShen1);
                var j1 = data.JinShen1;
                var j2 = data.JinShen2;
                var j3 = data.JinShen3;
                var jinsheninfo = j1 + "<br>" + j2 + "<br>" + j3;

                var p = document.getElementById("y1");
                p.innerHTML = jinsheninfo;
            }
        }
    }
} */
/* $.ajax({
    type: "POST",
    // dataType: "json",
    url: "/ccal",
    data: $('#btny1').serialize(),
    contentType: 'application/json;charset=utf-8',
    success: function (data) {
        var p = document.getElementById("p1");
        p.innerHTML = data.todayInfo;
    },
    error: function () {
        alert("今日信息异常");
    }
}); */