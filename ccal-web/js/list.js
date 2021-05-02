window.onload = function () {
   //提交表单数据到服务器
/*   function formonsub() {
    var request = new XMLHttpRequest();
    request.onreadystatechange = function () {
      if (request.readyState === 4 && request.status === 200) {
      };
    };
    // 发送请求:
    request.open('POST', '/', true);
    request.send();
    return false;
  } 
 */
  //年
  function selectY() {
    var d = new Date();
    var year = d.getFullYear(); //当前年数字
    var yearSelect = document.getElementById("yearid");

    for (var i = 1601; i <= 3498; i++) {
      var opt = document.createElement("option"); //生成option
      opt.value = i;
      opt.text = i + "年";
      //console.log(opt.value);
      yearSelect.appendChild(opt); //添加option到select
    }
    //添加select到from1表单
    document.body.appendChild(document.getElementById("form1"));

    //默认显示本年数字
    var def = year;
    document.getElementsByName("ly")[0].value = def;
  }

  //月 value 正月1 二月2
  function selectM() {
    var monthSelect = document.getElementById("monthid");
    var ymc = new Array('正', '二', '三', '四', '五', '六', '七', '八', '九', '十', '十一', '十二');
    var zhi = new Array("寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥", "子", "丑");

    for (var i = 0; i < 12; i++) {
      var opt = document.createElement("option"); //生成option
      opt.value = i + 1;
      opt.text = ymc[i] + "月" + (zhi[i]);
      //console.log("value:",opt.value,"name:",opt.text);
      monthSelect.appendChild(opt); //添加option到select
    }
    //添加select到from1表单
    document.body.appendChild(document.getElementById("form1"));
  }

  //日 初一value=1 初二value=2
  function selectD() {
    var dSelect = document.getElementById("dayid");
    var rmc = new Array(
      '初一', '初二', '初三', '初四', '初五', '初六', '初七', '初八', '初九', '初十',
      '十一', '十二', '十三', '十四', '十五', '十六', '十七', '十八', '十九', '二十',
      '廿一', '廿二', '廿三', '廿四', '廿五', '廿六', '廿七', '廿八', '廿九', '三十');

    for (var i = 0; i < rmc.length; i++) {
      var opt = document.createElement("option"); //生成option
      opt.value = i + 1;
      opt.text = rmc[i];
      // console.log("value:",opt.value,"name:",opt.text);
      dSelect.appendChild(opt); //添加option到select
    }
    //添加select到from1表单
    document.body.appendChild(document.getElementById("form1"));
  }

  //时辰 子时value=1 丑时value=2
  function selectH() {
    var hSelect = document.getElementById("hourid");
    var zhi = new Array("子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥");
    var aliasZhi = new Array("23~1", "1~3", "3~5", "5~7", "7~9", "9~11", "11~13", "13~15", "15~17", "17~19", "19~21", "21~23");
    for (var i = 0; i < zhi.length; i++) {
      var opt = document.createElement("option");
      opt.value = i + 1;
      opt.text = zhi[i] + "时" + aliasZhi[i];
      // console.log(opt.value, opt.text)
      hSelect.appendChild(opt);
    }
    document.body.appendChild(document.getElementById("form1"));
  }

  //生肖 子水鼠value=1 丑土牛value=2 
  function shengX() {
    var sxarr = new Array(
      "子水鼠", "丑土牛", "寅木虎", "卯木兔", "辰土龙", "巳火蛇",
      "午火马", "未土羊", "申金猴", "酉金鸡", "戌土狗", "亥水猪");
    var sxSelect = document.getElementById("sxid");
    for (var i = 0; i < sxarr.length; i++) {
      var opt = document.createElement("option");
      opt.value = i + 1;
      opt.text = sxarr[i]
      console.log(opt.value, opt.text)
      sxSelect.appendChild(opt);
    }
    document.body.appendChild(document.getElementById("form1"));
  }
  //formonsub();
  selectY();
  selectM();
  selectD();
  selectH();
  shengX();
}
