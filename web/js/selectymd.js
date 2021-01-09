window.onload = function () {
    //年份选择
   function sleectY() {
       var d = new Date();
       var year = d.getFullYear(); //当前年数字
       var yearSelect = document.getElementById("yearid");

       for (var i = 1601; i <= 3498; i++) {
           var opt = document.createElement("option"); //生成option
           opt.value = i;
           opt.text = i + "年";
           yearSelect.appendChild(opt); //添加option到select
       }
       //添加select到from1表单
       document.body.appendChild(document.getElementById("form1"));

       //默认显示本年数字
       var def = year;
       document.getElementsByName("ly")[0].value = def;
   } 
selectY();
}