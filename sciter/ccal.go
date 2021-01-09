package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/window"
	"liangzi.local/nongli/ccal"
	"liangzi.local/nongli/ganzhi"
	"liangzi.local/nongli/lunar"
	"liangzi.local/nongli/solar"
	"liangzi.local/nongli/today"
	"liangzi.local/qx/qx"
	"liangzi.local/sjqm"
	"liangzi.local/ts/jfj"
	"liangzi.local/ts/jx"
	"liangzi.local/ts/xjbfs"
)

func main() {
	w, err := window.New(sciter.SW_TITLEBAR|sciter.SW_RESIZEABLE|sciter.SW_CONTROLS|sciter.SW_MAIN,
		&sciter.Rect{Left: 0, Top: 0, Right: 1280, Bottom: 800}, //设置初始窗口大小
	)
	if err != nil {
		log.Fatal(err)
	}
	//加载文件
	w.LoadFile("ccal.html")
	//设置标题
	w.SetTitle("农历择日")

	//窗口根元素
	root, _ := w.GetRootElement()
	setElementHandlers(root)

	SetWinHandler(w)
	//吉凶信息
	SetWinHandlerInfo(w)
	//显示窗口
	w.Show()
	//运行窗口
	w.Run()
}

//######################
/*月吉 顺序同html脚本相同
"天道","天赦","天德", "天德合", "天恩", "天愿", "月德", "月德合", "月恩",
"月空", "母仓", "时德", "阴德", "阳德", "时阳生气", "益后", "续世",
"四相", "天仓", "要安", "敬安", "三合", "五合", "六合", "天医天喜",
"五富", "玉宇", "福德天巫", "六仪", "金堂", "天马", "时阴", "驿马",
"普护", "福生", "解神", "除神", "聖心", "吠鸣日", "吠鸣对日", "临日",
"王官守相民日", "建禄", "宝", "养日", "专日", "制日"*/

/*月凶
"天罡河魁死神", "河魁天罡", "月建", "月刑", "月害", "月破", "月厌地火",
"月煞月虚", "天火灾煞", "小时", "血支", "天贼", "往亡", "咸池大时大败",
"厌对招摇", "九空", "九坎", "游祸", "劫煞", "重日", "管符死气", "大耗",
"小耗", "复日", "天吏致死", "大煞", "八龙", "四穷", "四耗", "四废",
"四忌", "五虚", "五离", "八风", "土符", "归忌", "血忌", "五墓",
"八专", "触水龙", "兵禁", "月忌日", "无禄日", "上朔", "伐日",
*/
func SetWinHandlerInfo(w *window.Window) {
	w.DefineFunction("reg", func(args ...*sciter.Value) *sciter.Value {
		var js string
		for _, v := range args {
			js = v.String()
		}
		//解析前端传来的json数据
		var regObj YearJXObj
		json.Unmarshal([]byte(js), &regObj)
		//岁吉凶煞
		var ji, xiong, sha string
		if len(regObj.SuiJi) < 1 {
			ji = "岁德"
		} else {
			ji = regObj.SuiJi
		}
		if regObj.SuiXiong == "" {
			xiong = "力士"
		} else {
			xiong = regObj.SuiXiong
		}
		if regObj.SuiSha == "" {
			sha = "岁煞"
		} else {
			sha = regObj.SuiSha
		}
		//岁吉凶
		a, b, c := jx.SuiJiXiong(ji, xiong, sha)
		//月吉凶
		var mj, mx string
		if regObj.Mjs == "" {
			mj = "天德"
		} else {
			mj = regObj.Mjs
		}
		if regObj.Mjs == "" {
			mx = "天贼"
		} else {
			mx = regObj.Mxs
		}
		mjs, mxs := jx.YueJiXiong(mj, mx)
		//fmt.Println("web-->", regObj.Mjs, regObj.Mxs, mjs, mxs)
		//返回处理之后的值
		s0 := a + "\n" + b + "\n" + c + "\n" //岁吉凶煞
		s1 := mjs + mxs + "\n"               //月吉凶
		s := s0 + s1
		return sciter.NewValue(s)
	})
	//干支八卦查询
	w.DefineFunction("ganZhiGua", ganZhiGua)
}
func ganZhiGua(args ...*sciter.Value) *sciter.Value {
	var jsgzg string
	for _, v := range args {
		jsgzg = v.String()
	}
	//解析前端传来的json数据
	var gzgObj GanZhiGua
	json.Unmarshal([]byte(jsgzg), &gzgObj)

	var gyy, zyy, guayy string //阴阳
	gan := gzgObj.Gan
	zhi := gzgObj.Zhi
	gua := gzgObj.Gua
	ganInfo := ganzhi.NewGAN(gan)
	zhiInfo := ganzhi.NewZHI(zhi)
	guaInfo := ganzhi.NewGUA(gua)
	if ganInfo.YinYang == true {
		gyy = "阳"
	} else if ganInfo.YinYang == false {
		gyy = "阴"
	}
	if zhiInfo.YinYang == true {
		zyy = "阳"
	} else if zhiInfo.YinYang == false {
		zyy = "阴"
	}
	if guaInfo.YinYang == true {
		guayy = "阳"
	} else if guaInfo.YinYang == false {
		guayy = "阴"
	}
	gans := fmt.Sprintf("天干:%s 方位:%s 五行:%s 八卦:%s %s\n", ganInfo.Name, ganInfo.FangXiang, ganInfo.WuXing, ganInfo.Gua, gyy)
	zhis := fmt.Sprintf("地支:%s 方位:%s 五行:%s 八卦:%s %s\n", zhiInfo.Name, zhiInfo.FangXiang, zhiInfo.WuXing, zhiInfo.Gua, zyy)
	guas := fmt.Sprintf("八卦:%s 天干:%s 支:%s 方位:%s 五行:%s %s 节气:%s\n",
		guaInfo.Name, guaInfo.Gan, guaInfo.Zhi, guaInfo.FangXiang, guaInfo.WuXing, guayy, guaInfo.JieQi)

	all := gans + zhis + guas
	return sciter.NewValue(all)
}

type GanZhiGua struct {
	Gan string `json:"gan"`
	Zhi string `json:"zhi"`
	Gua string `json:"gua"`
}

type YearJXObj struct {
	SuiJi    string `json:"suiji"`    //岁吉
	SuiXiong string `json:"suixiong"` //岁凶
	SuiSha   string `json:"suisha"`   //岁煞
	Mjs      string `json:"mjs"`      //月吉
	Mxs      string `json:"mxs"`      //月凶
}

//########################################
//设置元素的处理程序
func setElementHandlers(root *sciter.Element) {
	//获取下拉值 基本纪年
	btn1, err := root.SelectById("btn1")
	if err != nil {
		log.Fatal("btn1:", err)
	}
	btn1.DefineMethod("ymdh", ymdh)
	//获取下拉值 小六壬择吉
	btn2, err := root.SelectById("btn2")
	if err != nil {
		log.Fatal("btn2:", err)
	}
	btn2.DefineMethod("xlrzj", xlrzj)
	//获取下拉值 协纪辩方书
	btn3, err := root.SelectById("btn3")
	if err != nil {
		log.Fatal("btn3:", err)
	}
	btn3.DefineMethod("xjbfshu", xjbfshu)
	//今日农历
	btn4, err := root.SelectById("btn4")
	if err != nil {
		log.Fatal("btn4:", err)
	}
	btn4.DefineMethod("todayt", todayt)
	//24节气
	btn5, err := root.SelectById("btn5")
	if err != nil {
		log.Fatal("btn5:", err)
	}
	btn5.DefineMethod("jieqi", jieqi)
	//关于
	btn6, err := root.SelectById("btn6")
	if err != nil {
		log.Fatal("btn6:", err)
	}
	btn6.DefineMethod("about", about)
	//奇门
	btn7, err := root.SelectById("btn7")
	if err != nil {
		log.Fatal("btn7:", err)
	}
	btn7.DefineMethod("qimen", qimen)
	//禽星
	btn8, err := root.SelectById("btn8")
	if err != nil {
		log.Fatal("btn8:", err)
	}
	btn8.DefineMethod("qinxing", qinxing)
}

//处理下拉列表获取的数据 元素的处理程序
func ymdh(args ...*sciter.Value) *sciter.Value {
	var ly, lm, ld, lh, sx, lb string
	//获取下拉值
	for k, arg := range args {
		//String()把参数转换成字符串
		//fmt.Printf("-->ymdh k=%d v=%s\n", k, arg.String())
		switch k {
		case 0:
			ly = arg.String()
		case 1:
			lm = arg.String()
		case 2:
			ld = arg.String()
		case 3:
			lh = arg.String()
		case 4:
			sx = arg.String()
		case 5:
			lb = arg.String()
		}
	}
	//fmt.Printf("--> %s-%d %s-%d %s-%d %s-%d %s-%d %s-%d\n", ly, len(ly), lm, len(lm), ld, len(ld), lh, len(lh), sx, len(sx), lb, len(lb))

	if len(lm) < 2 {
		lm = "0" + lm
	}
	if len(ld) < 2 {
		ld = "0" + ld
	}
	if len(lh) < 2 {
		lh = "0" + lh
	}
	if len(sx) < 2 {
		sx = "0" + sx
	}
	//fmt.Printf("-->++ %s-%d %s-%d %s-%d %s-%d %s-%d %s-%d\n", ly, len(ly), lm, len(lm), ld, len(ld), lh, len(lh), sx, len(sx), lb, len(lb))
	sinfo := fmt.Sprintf("%s%s%s%s%s%s", ly, lm, ld, lh, sx, lb)
	return sciter.NewValue(sinfo)

}
func xlrzj(args ...*sciter.Value) *sciter.Value {
	var ly, lm, ld, lh, sx, lb string
	//获取下拉值
	for k, arg := range args {
		switch k {
		case 0:
			ly = arg.String()
		case 1:
			lm = arg.String()
		case 2:
			ld = arg.String()
		case 3:
			lh = arg.String()
		case 4:
			sx = arg.String()
		case 5:
			lb = arg.String()
		}
	}
	if len(lm) < 2 {
		lm = "0" + lm
	}
	if len(ld) < 2 {
		ld = "0" + ld
	}
	if len(lh) < 2 {
		lh = "0" + lh
	}
	if len(sx) < 2 {
		sx = "0" + sx
	}
	sinfo := fmt.Sprintf("%s%s%s%s%s%s", ly, lm, ld, lh, sx, lb)
	return sciter.NewValue(sinfo)
}
func xjbfshu(args ...*sciter.Value) *sciter.Value {
	var ly, lm, ld, lh, sx, lb string
	//获取下拉值
	for k, arg := range args {
		switch k {
		case 0:
			ly = arg.String()
		case 1:
			lm = arg.String()
		case 2:
			ld = arg.String()
		case 3:
			lh = arg.String()
		case 4:
			sx = arg.String()
		case 5:
			lb = arg.String()
		}
	}
	if len(lm) < 2 {
		lm = "0" + lm
	}
	if len(ld) < 2 {
		ld = "0" + ld
	}
	if len(lh) < 2 {
		lh = "0" + lh
	}
	if len(sx) < 2 {
		sx = "0" + sx
	}
	sinfo := fmt.Sprintf("%s%s%s%s%s%s", ly, lm, ld, lh, sx, lb)
	return sciter.NewValue(sinfo)
}
func todayt(args ...*sciter.Value) *sciter.Value {
	return sciter.NullValue()
}
func jieqi(args ...*sciter.Value) *sciter.Value {
	return sciter.NewValue(args[0].String())
}
func about(args ...*sciter.Value) *sciter.Value {
	return sciter.NullValue()
}
func qimen(args ...*sciter.Value) *sciter.Value {
	var ly, lm, ld, lh, sx, lb string
	//获取下拉值
	for k, arg := range args {
		switch k {
		case 0:
			ly = arg.String()
		case 1:
			lm = arg.String()
		case 2:
			ld = arg.String()
		case 3:
			lh = arg.String()
		case 4:
			sx = arg.String()
		case 5:
			lb = arg.String()
		}
	}
	if len(lm) < 2 {
		lm = "0" + lm
	}
	if len(ld) < 2 {
		ld = "0" + ld
	}
	if len(lh) < 2 {
		lh = "0" + lh
	}
	if len(sx) < 2 {
		sx = "0" + sx
	}
	sinfo := fmt.Sprintf("%s%s%s%s%s%s", ly, lm, ld, lh, sx, lb)
	return sciter.NewValue(sinfo)
}
func qinxing(args ...*sciter.Value) *sciter.Value {
	var ly, lm, ld, lh, sx, lb string
	//获取下拉值
	for k, arg := range args {
		switch k {
		case 0:
			ly = arg.String()
		case 1:
			lm = arg.String()
		case 2:
			ld = arg.String()
		case 3:
			lh = arg.String()
		case 4:
			sx = arg.String()
		case 5:
			lb = arg.String()
		}
	}
	if len(lm) < 2 {
		lm = "0" + lm
	}
	if len(ld) < 2 {
		ld = "0" + ld
	}
	if len(lh) < 2 {
		lh = "0" + lh
	}
	if len(sx) < 2 {
		sx = "0" + sx
	}
	sinfo := fmt.Sprintf("%s%s%s%s%s%s", ly, lm, ld, lh, sx, lb)
	return sciter.NewValue(sinfo)
}

//#################################
func SetWinHandler(w *window.Window) {
	//纪年信息
	w.DefineFunction("ymdinfo", ymdinfo)
	//小六壬择吉信息
	w.DefineFunction("xlrzjinfo", xlrzjinfo)
	//协纪辩方书
	w.DefineFunction("xjbfsinfo", xjbfsinfo)
	//今日农历
	w.DefineFunction("todaytinfo", todaytinfo)
	//24节气
	w.DefineFunction("jieqiinfo", jieqiinfo)
	//关于
	w.DefineFunction("aboutinfo", aboutinfo)
	//奇门
	w.DefineFunction("qimeninfo", qimeninfo)
	//禽星
	w.DefineFunction("qinxinginfo", qinxinginfo)
}

//纪年信息
func ymdinfo(args ...*sciter.Value) *sciter.Value {
	var s string
	for _, v := range args {
		//fmt.Printf("-->ymdinfo v:%s len=%d\n", v.String(), len(v.String()))
		s = v.String()
	}
	y, m, d, h, sx, b := convs(s)
	err, solar, lu, g, _ := ccal.Input(y, m, d, h, sx, b)
	if err != nil {
		log.Fatal("ccal-input:", err)
	}
	ygz := fmt.Sprintf("%s%s", g.YearGanM, g.YearZhiM) //年干支
	mgz := g.MonthGanZhiM                              //月干支
	dgz := fmt.Sprintf("%s%s", g.DayGanM, g.DayZhiM)   //日干支
	hgz := g.HourGanZhiM                               //时干支
	var aliasM string
	if lu.Leapmb == true {
		aliasM = "是"
	} else {
		aliasM = "否"
	}
	jng := fmt.Sprintf("干支纪年:%s年-%s月-%s日-%s时", ygz, mgz, dgz, hgz)
	jns := fmt.Sprintf("阳历纪年:%d年-%d月-%d日-周%s", solar.SYear, solar.SMonth, solar.SDay, solar.SWeek)
	jnl := fmt.Sprintf("农历纪年: %d年%s月(%s)%s %s时(%d时)",
		lu.LYear, lunar.Ymc[lu.LMonth-1], lu.LYdxs, lunar.Rmc[lu.LDay-1], lu.LaliasHour, lu.LHour)
	jnmb := fmt.Sprintf("本年是否有闰月:%s 闰%d月", aliasM, lu.LeapMonth)
	//纳音
	ygzny := ganzhi.GZ纳音(ygz)
	mgzny := ganzhi.GZ纳音(mgz)
	dgzny := ganzhi.GZ纳音(dgz)
	hgzny := ganzhi.GZ纳音(hgz)
	ny := fmt.Sprintf("干支纳音: %s %s %s %s", ygzny[ygz], mgzny[mgz], dgzny[dgz], hgzny[hgz])

	jn := JN{
		Sjn: jns,
		Ljn: jnl,
		Gjn: jng,
		Lmb: jnmb,
		Ny:  ny,
	}
	jnjson, err := json.Marshal(jn)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("jnjson:%s\n", string(jnjson))

	return sciter.NewValue(string(jnjson))
}

//小六壬择吉信息
func xlrzjinfo(args ...*sciter.Value) *sciter.Value {
	var as string
	for _, v := range args {
		as = v.String()
	}
	y, m, d, h, sx, b := convs(as)
	err, s, l, g, jq := ccal.Input(y, m, d, h, sx, b)
	if err != nil {
		log.Fatal("ccal-input:", err)
	}
	ygz := fmt.Sprintf("%s%s", g.YearGanM, g.YearZhiM)
	mgz := g.MonthGanZhiM
	dgz := fmt.Sprintf("%s%s", g.DayGanM, g.DayZhiM)
	hgz := g.HourGanZhiM
	yeargan := g.YearGanM
	yearzhi := g.YearZhiM

	aliaslmonth := lunar.ConvYmc(l.LMonth)
	aliaslday := lunar.ConvRmc(l.LDay)
	aliaslhour := lunar.ConvHourZhi(g.HourGanZhiM)
	aliaslydxs := l.LYdxs

	lyear := l.LYear
	lmonth := l.LMonth
	lday := l.LDay
	lhour := l.LHour
	lydx := l.LYdx

	stime := time.Date(s.SYear, time.Month(s.SMonth), s.SDay, 0, 0, 0, 0, time.Local)
	syear := s.SYear
	smonth := s.SMonth
	sday := s.SDay
	sweek := s.SWeek

	leapmb := l.Leapmb
	leapmonth := l.LeapMonth
	lunarmjd := jq.LunarmJd
	zryl := NewZRYL(ygz, mgz, dgz, hgz, yeargan, yearzhi,
		aliaslmonth, aliaslday, aliaslhour, aliaslydxs,
		lyear, lmonth, lday, lhour, lydx,
		stime, syear, smonth, sday, sweek,
		leapmb, leapmonth, lunarmjd)

	///############################
	///小六壬七煞
	qsOBj := zryl.XiaoLRStarObj()
	////择吉结果 显示本日吉星数字
	b1, b2, b3 := zryl.XiaoLRBool()
	zeji := qsOBj.GoodNumberDay(b1, b2, b3)
	//本月吉干 七煞数组
	jgArr, qsArr := zryl.XiaoLRJiGanQiShaArr(sx)

	jgArrs := "本月吉干: " + strings.Join(jgArr, " ")
	qsArrs := "本月七煞: " + strings.Join(qsArr, " ")

	xlrzj := XLRZJ{
		XstarName: "本日值宿: " + qsOBj.Name,
		Xinfo:     qsOBj.Info,
		Xzeji:     zeji,
		XjiGanArr: jgArrs,
		XQiShaArr: qsArrs,
	}
	xlrzjJson, err := json.Marshal(xlrzj)
	if err != nil {
		log.Fatal("xlrzjJson:", err)
	}
	return sciter.NewValue(string(xlrzjJson))
}

//协纪辩方书
func xjbfsinfo(args ...*sciter.Value) *sciter.Value {
	var as string
	for _, v := range args {
		as = v.String()
	}
	y, m, d, h, sx, b := convs(as)
	err, s, l, g, jq := ccal.Input(y, m, d, h, sx, b)
	if err != nil {
		log.Fatal("ccal-input:", err)
	}
	ygz := fmt.Sprintf("%s%s", g.YearGanM, g.YearZhiM)
	mgz := g.MonthGanZhiM
	dgz := fmt.Sprintf("%s%s", g.DayGanM, g.DayZhiM)
	hgz := g.HourGanZhiM
	yeargan := g.YearGanM
	yearzhi := g.YearZhiM

	aliaslmonth := lunar.ConvYmc(l.LMonth)
	aliaslday := lunar.ConvRmc(l.LDay)
	aliaslhour := lunar.ConvHourZhi(g.HourGanZhiM)
	aliaslydxs := l.LYdxs

	lyear := l.LYear
	lmonth := l.LMonth
	lday := l.LDay
	lhour := l.LHour
	lydx := l.LYdx

	stime := time.Date(s.SYear, time.Month(s.SMonth), s.SDay, 0, 0, 0, 0, time.Local)
	syear := s.SYear
	smonth := s.SMonth
	sday := s.SDay
	sweek := s.SWeek

	leapmb := l.Leapmb
	leapmonth := l.LeapMonth
	lunarmjd := jq.LunarmJd
	zryl := NewZRYL(ygz, mgz, dgz, hgz, yeargan, yearzhi,
		aliaslmonth, aliaslday, aliaslhour, aliaslydxs,
		lyear, lmonth, lday, lhour, lydx,
		stime, syear, smonth, sday, sweek,
		leapmb, leapmonth, lunarmjd)
	jz60 := ganzhi.MakeJZ60() //六十甲子

	//择日 协纪辩方书
	ji, taiSuiWuGui, sha, xiong := zryl.XJBF年表(jz60) //协纪辩方 年表
	djc, jcb := zryl.JC12M()                         //协纪辩方 日建除
	yjh, rj, rx := zryl.XJBF月表(jcb)                  //协纪辩方 月表
	hcj, hcx := zryl.XJBF日表(jz60)                    //协纪辩方 日表
	//bw1, bw2, bw3 := zryl.BianWei()                  //协纪辩方 辩伪

	//转换
	//年表
	jis := xjbfs.ConvArrToS(ji)
	taiSuiWuGuis := xjbfs.ConvArrToS(taiSuiWuGui)
	shas := xjbfs.ConvArrToS(sha)
	xiongs := xjbfs.ConvArrToS(xiong)
	yminfo := "岁吉: " + jis + "\n" +
		"岁煞: " + shas + "\n" +
		"岁凶: " + xiongs + "\n" +
		"金神五鬼: " + taiSuiWuGuis + "\n"

	//月将
	jqt := xjbfs.JQT(y)
	solarT := stime
	yjs := xjbfs.NewYueJiang(solarT, jqt)
	yjnames := yjs.Name //月将名
	stars := yjs.Star   //十二宫星
	jqs := yjs.JieQi
	YueJiangs := "月将: " + yjnames + "\n" + "十二宫: " + stars + "\n" + "节气: " + jqs + "\n"

	//月表
	yjhs := xjbfs.ConvArrToS(yjh) //月总论
	yjhs = xjbfs.DelString(yjhs, "<br>", "\n")
	rjs := xjbfs.ConvArrToS(rj)                       //月日论 日吉
	rxs := xjbfs.ConvArrToS(rx)                       //月日论 日凶
	dhh := zryl.HuangHeiDay()                         //日黄黑
	djcs := "日建除:" + djc + "\n" + "日黄黑:" + dhh + "\n" //日建除+日黄黑
	mdinfo := YueJiangs + "\n" + "月总论:\n" + yjhs + "\n" + djcs + "日吉: " + rjs + "\n" + "日凶:" + rxs

	//日表
	hgxs := zryl.HuangHeiHour()   //时辰黄黑
	hcjs := xjbfs.ConvArrToS(hcj) //时辰吉
	hcxs := xjbfs.ConvArrToS(hcx) //时辰凶
	dhinfo := "时辰黄黑: " + hgxs + "\n" + "时辰吉: " + hcjs + "\n" + "时辰凶: " + hcxs

	//咸池桃花
	xcth := xjbfs.XCTH咸池桃花(zryl.YGZ, zryl.MGZ, zryl.DGZ, zryl.HGZ)
	xcth = "咸池桃花: " + xcth
	//时孤虚
	guxu := zryl.GuXu()
	//金符九星
	ymc := zryl.AliasLMonth
	jfs := jfj.JinFuJing(ymc, zryl.DGZ)
	jfs = "金符九星: " + jfs
	//本月七煞日 这里是根据年干+日支计算
	days := jfj.FindDays(lunarmjd, zryl.Lydx)
	dzs, _ := jfj.QiShaDay(ygz, dgz)
	jsarr := jfj.ListQS(days, dzs)
	jinShens := xjbfs.ConvArrToS(jsarr)
	jinShens = "金神七煞日: " + jinShens

	//贵登天门
	sm1, sm0 := yjs.GuiDengTianMen(dgz)
	var dtms string
	if len(sm1) > 0 && len(sm0) > 0 {
		dtms = sm1 + "|" + sm0
	}

	bws := guxu + "\n" + xcth + "\n" + jfs + "\n" + jinShens + "\n" + dtms

	xj := XJBFS{
		Nb: yminfo,
		Yb: mdinfo,
		Rb: dhinfo,
		Bw: bws,
	}
	xjJson, err := json.Marshal(xj)
	if err != nil {
		log.Fatal("jxJson:", err)
	}
	return sciter.NewValue(string(xjJson))
}

//今日农历
func todaytinfo(args ...*sciter.Value) *sciter.Value {
	////今日信息
	var T = time.Now().Local()
	lunary, lunarm, lunard, lunarh, leapmB := today.TodayT(T)
	var lts string
	if leapmB == true {
		lts = fmt.Sprintf("--今日农历:%d年%d月%d日%d时-(%s时) 当前月份是闰月",
			lunary, lunarm, lunard, lunarh, lunar.ConvHToName(lunarh))
	} else if leapmB == false {
		lts = fmt.Sprintf("今日农历:%d年%d月%d日%d时-(%s时) 当前月份不是闰月",
			lunary, lunarm, lunard, lunarh, lunar.ConvHToName(lunarh))

	}
	return sciter.NewValue(lts)
}

//24节气
func jieqiinfo(args ...*sciter.Value) *sciter.Value {
	y, err := strconv.Atoi(args[0].String())
	if err != nil {
		log.Fatal("节气年份:", err)
	}
	jqtarr := solar.JQT(y)
	jqt := solar.NewJQ(jqtarr)
	jqArr := jqt.JieQi()
	jqs := strings.Join(jqArr, "\n")
	return sciter.NewValue(jqs)
}

//关于
func aboutinfo(args ...*sciter.Value) *sciter.Value {
	ab := xjbfs.AboutMe()
	abouts := ab.Info + "\n" + "下载链接: " + ab.Url1 + "\n" + "农历编算参考: " + ab.Url2 + "\n" + "Mail: " + ab.Me + "\n"
	return sciter.NewValue(abouts)
}

//奇门
func qimeninfo(args ...*sciter.Value) *sciter.Value {
	var as string
	for _, v := range args {
		as = v.String()
	}
	y, m, d, h, sx, b := convs(as)
	err, s, _, g, _ := ccal.Input(y, m, d, h, sx, b)
	if err != nil {
		log.Fatal("ccal-input:", err)
	}
	yg := g.YearGanM
	yz := g.YearZhiM
	ygz := fmt.Sprintf("%s%s", yg, yz)
	mgz := g.MonthGanZhiM
	///时家奇门
	dgzm := fmt.Sprintf("%s%s", g.DayGanM, g.DayZhiM)
	hgzm := g.HourGanZhiM
	//这里的s.SHour是由输入的时辰转换而来
	st := time.Date(s.SYear, time.Month(s.SMonth), s.SDay, s.SHour, 0, 0, 0, time.Local)
	G, _ := sjqm.Result(y, dgzm, hgzm, st)
	qms := fmt.Sprintf("时家奇门\n"+
		"节气:%s %s %s %d局 旬首:%s 值符:%s 值使:%s\n"+
		"一宫 ==> 九星:%s 八门:%s 暗干支:%s 天盘奇仪:%s 八神:%s 地盘奇仪:%s\n"+
		"八宫 ==> 九星:%s 八门:%s 暗干支:%s 天盘奇仪:%s 八神:%s 地盘奇仪:%s\n"+
		"三宫 ==> 九星:%s 八门:%s 暗干支:%s 天盘奇仪:%s 八神:%s 地盘奇仪:%s\n"+
		"四宫 ==> 九星:%s 八门:%s 暗干支:%s 天盘奇仪:%s 八神:%s 地盘奇仪:%s\n"+
		"五宫 ==> 九星:%s 八门:%s 暗干支:%s 地盘奇仪:%s\n"+
		"九宫 ==> 九星:%s 八门:%s 暗干支:%s 天盘奇仪:%s 八神:%s 地盘奇仪:%s\n"+
		"二宫 ==> 九星:%s 八门:%s 暗干支:%s 天盘奇仪:%s 八神:%s 地盘奇仪:%s\n"+
		"七宫 ==> 九星:%s 八门:%s 暗干支:%s 天盘奇仪:%s 八神:%s 地盘奇仪:%s\n"+
		"六宫 ==> 九星:%s 八门:%s 暗干支:%s 天盘奇仪:%s 八神:%s 地盘奇仪:%s\n",
		G.JieQi, G.YinYang, G.YUAN, G.N, G.XS, G.ZHIFU, G.ZHISHI,
		G.G1[0], G.G1[1], G.G1[2], G.G1[3], G.G1[4], G.G1[5],
		G.G8[0], G.G8[1], G.G8[2], G.G8[3], G.G8[4], G.G8[5],
		G.G3[0], G.G3[1], G.G3[2], G.G3[3], G.G3[4], G.G3[5],
		G.G4[0], G.G4[1], G.G4[2], G.G4[3], G.G4[4], G.G4[5],
		G.G5[0], G.G5[1], G.G5[2], G.G5[3],
		G.G9[0], G.G9[1], G.G9[2], G.G9[3], G.G9[4], G.G9[5],
		G.G2[0], G.G2[1], G.G2[2], G.G2[3], G.G2[4], G.G2[5],
		G.G7[0], G.G7[1], G.G7[2], G.G7[3], G.G7[4], G.G7[5],
		G.G6[0], G.G6[1], G.G6[2], G.G6[3], G.G6[4], G.G6[5],
	)
	gzs := fmt.Sprintf("%s-%s-%s-%s\n", ygz, mgz, dgzm, hgzm)
	return sciter.NewValue(gzs + qms)
}

//禽星
func qinxinginfo(args ...*sciter.Value) *sciter.Value {
	var as string
	for _, v := range args {
		as = v.String()
	}
	y, m, d, h, sx, b := convs(as)
	err, _, l, g, _ := ccal.Input(y, m, d, h, sx, b)
	if err != nil {
		log.Fatal("ccal-input:", err)
	}
	yg := g.YearGanM
	yz := g.YearZhiM
	ygz := fmt.Sprintf("%s%s", yg, yz)
	mgz := g.MonthGanZhiM
	dg := g.DayGanM
	dz := g.DayZhiM
	dgz := fmt.Sprintf("%s%s", dg, dz)
	sy := l.LYear
	sm := l.LMonth
	qy := qx.FindNianQin(sy)
	qm := qx.FindYueQin(sy, sm)
	meta := qx.YearMetaNumber(sy)
	qd, _ := qx.FindRiQin(meta, dgz)
	qxs := fmt.Sprintf("%s年%s日%s时\n%d元\n年禽:%s 月禽:%s 日禽:%s\n", ygz, mgz, dgz, meta, qy, qm, qd)
	return sciter.NewValue(qxs)
}
func NewZRYL(ygz, mgz, dgz, hgz, yeargan, yearzhi string,
	aliaslmonth, aliaslday, aliaslhour, aliaslydxs string,
	lyear, lmonth, lday, lhour, lydx int,
	stime time.Time, syear, smonth, sday int, sweek string,
	leapmb bool, leapmonth int, lunarmjd float64) *xjbfs.ZRYL {
	return xjbfs.NewZRYL(ygz, mgz, dgz, hgz, yeargan, yearzhi,
		aliaslmonth, aliaslday, aliaslhour, aliaslydxs,
		lyear, lmonth, lday, lhour, lydx,
		stime, syear, smonth, sday, sweek,
		leapmb, leapmonth, lunarmjd)
}

//协纪辩方书
type XJBFS struct {
	Nb string `json:"nb"`
	Yb string `json:"yb"`
	Rb string `json:"rb"`
	Bw string `json:"bw"`
}

//小六壬择吉
type XLRZJ struct {
	XstarName string `json:"xstar_name"`
	Xinfo     string `json:"xinfo"`
	Xzeji     string `json:"xzeji"`
	XjiGanArr string `json:"x_ji_gan_arr"`
	XQiShaArr string `json:"x_qi_sha_arr"`
}

//纪年信息
type JN struct {
	Sjn string `json:"sjn"`
	Ljn string `json:"ljn"`
	Gjn string `json:"gjn"`
	Lmb string `json:"lmb"`
	Ny  string `json:"ny"`
}

func convs(s string) (int, int, int, int, string, bool) {
	rs := []rune(s)
	ys := string(rs[:4])  //年數字
	ms := string(rs[4:6]) //月數字
	ds := string(rs[6:8])
	hs := string(rs[8:10]) //時辰數字1子時　2丑時...１２亥時
	sxs := string(rs[10:12])
	bs := string(rs[12:13])
	//fmt.Printf("--conv ly=%d lm=%d ld=%d lh=%d lsx=%d lb=%d\n", len(ys), len(ms), len(ds), len(hs), len(sxs), len(bs))
	y, err := strconv.Atoi(ys)
	if err != nil {
		log.Fatal("年份時間解析:", err)
	}

	m, err := strconv.Atoi(ms)
	if err != nil {
		log.Fatal("月份時間解析:", err)
	}
	d, err := strconv.Atoi(ds)
	if err != nil {
		log.Fatal("日期時間解析:", err)
	}
	h, err := strconv.Atoi(hs)
	if err != nil {
		log.Fatal("時辰解析:", err)
	}

	si, err := strconv.Atoi(sxs)
	if err != nil {
		log.Fatal("生肖解析:", err)
	}
	sx := lunar.ConvSX(si)

	bi, err := strconv.Atoi(bs)
	if err != nil {
		log.Fatal("闰月数字解析:", err)
	}
	if bi == 0 {
		bs = "f"
	} else if bi == 1 {
		bs = "t"
	}
	b, err := strconv.ParseBool(bs)
	if err != nil {
		log.Fatal("閏月bool解析:", err)
	}

	//fmt.Printf("--> convs: y:%v m:%v d:%v h:%v sx:%v b:%t\n", y, m, d, h, sx, b)
	return y, m, d, h, sx, b
}

func conv(ly, lm, ld, lh, sx, lb string) (y, m, d, h int, s string, b bool) {
	y, err := strconv.Atoi(ly)
	if err != nil {
		log.Fatal(err)
	}
	m, err = strconv.Atoi(lm)
	if err != nil {
		log.Fatal(err)
	}
	d, err = strconv.Atoi(ld)
	if err != nil {
		log.Fatal(err)
	}
	h, err = strconv.Atoi(lh)
	if err != nil {
		log.Fatal(err)
	}
	si, err := strconv.Atoi(sx)
	if err != nil {
		log.Fatal(err)
	}
	bi, err := strconv.Atoi(lb) //0:f 1:t
	if err != nil {
		log.Fatal(err)
	}
	s = lunar.ConvSX(si)
	if bi == 0 {
		b = false
	} else if bi == 1 {
		b = true
	}
	return

}
