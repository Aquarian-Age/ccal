package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"
	"time"

	"liangzi.local/ts/qimen"

	"liangzi.local/nongli/ccal"
	"liangzi.local/nongli/dimu"
	"liangzi.local/nongli/ganzhi"
	"liangzi.local/nongli/lunar"
	"liangzi.local/nongli/solar"
	"liangzi.local/nongli/today"
	xlr "liangzi.local/nongli/zeji"
	"liangzi.local/sjqm"
	"liangzi.local/ts/jfj"
	"liangzi.local/ts/xjbfs"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("home.html")
		t.Execute(w, nil)
	} else {
		handhome(w, r)
	}
}

func handhome(w http.ResponseWriter, r *http.Request) {
	//解析表单
	r.ParseForm()
	//农历年
	ly, err := strconv.Atoi(r.Form["ly"][0])
	if err != nil {
		log.Fatalln("ly:", err)
	}
	//农历月
	lm, err := strconv.Atoi(r.Form["lm"][0])
	if err != nil {
		log.Fatalln(err)
	}

	//fmt.Printf("农历: %d月\n", lm)
	ld, err := strconv.Atoi(r.Form["ld"][0])
	if err != nil {
		log.Fatalln(err)
	}
	//时辰 子时1 丑时2 寅时3...
	lh, err := strconv.Atoi(r.Form["lh"][0])
	if err != nil {
		log.Fatalln("时辰异常:", err)
	}
	//生肖
	sxs := r.Form["la"][0]
	sxn, err := strconv.Atoi(sxs)
	if err != nil {
		log.Fatal("生肖异常:", err)
	}
	var sx string
	var zhi = []string{"err", "鼠", "牛", "虎", "兔", "龙", "蛇", "马", "羊", "猴", "鸡", "狗", "猪"}
	for index := 0; index <= 13; index++ {
		if sxn == index {
			sx = zhi[index]
			break
		}
	}
	//闰月
	sb := r.Form["le"][0]
	var leapb bool
	if strings.EqualFold(sb, "t") {
		leapb = true
	} else if strings.EqualFold(sb, "f") {
		leapb = false
	}
	//ccal农历基本纪年信息
	_, s, l, g, jq := ccal.Input(ly, lm, ld, lh, sx, leapb)

	///////////////////////////
	ygz := fmt.Sprintf("%s%s", g.YearGanM, g.YearZhiM)
	mgz := g.MonthGanZhiM
	dgz := fmt.Sprintf("%s%s", g.DayGanM, g.DayZhiM)
	hgz := g.HourGanZhiM
	yeargan := g.YearGanM
	yearzhi := g.YearZhiM

	aliaslmonth := convYmc(l.LMonth)
	aliaslday := convRmc(l.LDay)
	aliaslhour := convHourZhi(g.HourGanZhiM)
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
	zryl = NewZRYL(ygz, mgz, dgz, hgz, yeargan, yearzhi,
		aliaslmonth, aliaslday, aliaslhour, aliaslydxs,
		lyear, lmonth, lday, lhour, lydx,
		stime, syear, smonth, sday, sweek,
		leapmb, leapmonth, lunarmjd)

	//纪年信息 纳音
	sc, lc, leap, gc := zryl.Lunar()
	nayin := zryl.NaYin()
	lunarCalendar := sc + "<br>" + lc + "<br>" + gc + "<br>" + nayin + "<br>" + leap + "<br>"
	jinianinfo := lunarCalendar

	//地母经
	dmg := g.YearGan
	dmz := g.YearZhi
	infodmj := dimu.DimuInfo(dmg, dmz)
	dmjinfo := fmt.Sprintf("%s", infodmj)

	//24节气
	jq24info := solar.ShowJieqi24(jq.Jqt, jq.Jq11t)

	//农历月历表
	listday := lunar.ListLunarDay(jq.LunarmJd, l.LYdx)

	//小六壬择吉
	wd := int(s.SolarDayT.Weekday()) //周几
	stars := xlr.NewStars(wd, dgz)
	starName := stars.Name //值宿名称//
	zhisus := stars.Info   //当日值宿信息//
	//本日星数字
	m := l.LMonth
	d := l.LDay
	h := l.LHour
	total := xlr.NewTotal(ygz, m, d, h)
	b1 := total.YiPan()
	b2 := total.ErPan()
	b3 := total.SanPan()
	zeji := stars.GoodNumberDay(b1, b2, b3) //
	//本月吉干信息
	jg1, jg2, jg3 := xlr.JiGanNumber(m)
	jgName := fmt.Sprintf("本月吉干:%s-%s-%s", ganzhi.Gan[jg1], ganzhi.Gan[jg2], ganzhi.Gan[jg3]) //

	//没有去除七煞的吉干map
	ydx := l.LYdx
	jgmap, qsarr, qsdaymap := xlr.FindMonthJG(lunarmjd, m, ydx)
	//去除金神七煞的吉干map
	newjgmap := xlr.DelQs(jgmap, qsarr)
	//去除生肖相冲的吉干map
	good := xlr.DelSX(newjgmap, sx)
	//本月吉干数组
	goodarr := xlr.GoodJG(good)
	goodjg := goodarr
	//本月七煞数组
	qsdays := xlr.GoodJG(qsdaymap) //

	//择日 协纪辩方书
	ji, taiSuiWuGui, sha, xiong := zryl.XJBF年表(jz60) //协纪辩方 年表
	djc, jcb := zryl.JC12M()                         //协纪辩方 建除十二神煞(日)
	yjh, rj, rx := zryl.XJBF月表(jcb)                  //协纪辩方 月表
	yjh = append(yjh, djc)                           //加日建除
	//加日黄黑
	zhimap := xjbfs.MakeHHMap()
	hhd := xjbfs.XJBF黄黑起法(mgz, dgz, zhimap)
	yjh = append(yjh, hhd)

	hcj, hcx := zryl.XJBF日表(jz60) //协纪辩方 日表
	//时辰黄黑
	zhihmap := xjbfs.MakeHHMap()
	hgxs := xjbfs.XJBF黄黑起法(dgz, hgz, zhihmap)
	//时孤虚
	_, gx := qimen.SJQM孤虚法(dgz, hgz)
	gu := "孤:" + gx["孤"]
	xu := "虚:" + gx["虚"]
	gus := gu + " " + xu
	hgx := []string{hgxs, gus}

	bw1, bw2, bw3 := zryl.BianWei() //协纪辩方 辩伪+其他
	bianwei := bw1 + "<br>" + bw2 + "<br>" + bw3 + "<br>"
	xieJiBFS := XJBF{
		Ji:          ji,
		TaiSuiWuGui: taiSuiWuGui,
		Sha:         sha,
		Xiong:       xiong,
		Yjh:         yjh,
		Rj:          rj,
		Rx:          rx,
		Hgx:         hgx,
		Hcj:         hcj,
		Hcx:         hcx,
		BW:          bianwei,
	}

	/* 	///通书内容
	   	rszl := Rszl(aliaslmonth, dgz, aliaslday, aliaslhour) //通书日时总览
	   	trj, trx := Zsyl(zryl.DGZ, zryl.HGZ)                  //择时要览
	   	tsinfo := TSInfo{
	   		TSRs: rszl,
	   		TSsj: trj,
	   		TSsx: trx,
	   	} */

	////月将
	jqt := xjbfs.JQT(ly)
	solarT := time.Date(s.SYear, time.Month(s.SMonth), s.SDay, 0, 0, 0, 0, time.UTC)
	yjs := xjbfs.NewYueJiang(solarT, jqt)
	yjname := yjs.Name //月将名
	star := yjs.Star   //十二宫星
	yj := YJ{
		YjName:   yjname,
		Zhi:      yjs.Zhi,
		StarName: star,
		JieQi:    yjs.JieQi,
	}
	//贵登天门
	s1, s0 := yjs.GuiDengTianMen(dgz)
	gdtm := s1 + "|" + s0

	///时家奇门
	y := l.LYear
	dgzm := fmt.Sprintf("%s%s", g.DayGanM, g.DayZhiM)
	hgzm := g.HourGanZhiM
	//这里的s.SHour不是准确的阳历时间 是由输入的时辰转换而来
	st := time.Date(s.SYear, time.Month(s.SMonth), s.SDay, s.SHour, 0, 0, 0, time.Local)
	G, _ := sjqm.Result(y, dgzm, hgzm, st)
	Sjqm := fmt.Sprintf("<br><b>时家奇门</b><br><br>"+
		"节气:%s %s %s %d局 旬首:%s 值符:%s 值使:%s <br>"+
		"一宫 ==> 九星:%s 八门:%s 暗干支:%s 天盘奇仪:%s 八神:%s 地盘奇仪:%s <br>"+
		"八宫 ==> 九星:%s 八门:%s 暗干支:%s 天盘奇仪:%s 八神:%s 地盘奇仪:%s <br>"+
		"三宫 ==> 九星:%s 八门:%s 暗干支:%s 天盘奇仪:%s 八神:%s 地盘奇仪:%s <br>"+
		"四宫 ==> 九星:%s 八门:%s 暗干支:%s 天盘奇仪:%s 八神:%s 地盘奇仪:%s <br>"+
		"五宫 ==> 九星:%s 八门:%s 暗干支:%s 地盘奇仪:%s <br>"+
		"九宫 ==> 九星:%s 八门:%s 暗干支:%s 天盘奇仪:%s 八神:%s 地盘奇仪:%s <br>"+
		"二宫 ==> 九星:%s 八门:%s 暗干支:%s 天盘奇仪:%s 八神:%s 地盘奇仪:%s <br>"+
		"七宫 ==> 九星:%s 八门:%s 暗干支:%s 天盘奇仪:%s 八神:%s 地盘奇仪:%s <br>"+
		"六宫 ==> 九星:%s 八门:%s 暗干支:%s 天盘奇仪:%s 八神:%s 地盘奇仪:%s <br>",
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

	//Other
	//金符九星
	ymc := convYmc(l.LMonth)
	jfs := jfj.JinFuJing(ymc, dgz)

	//本月七煞日 这里是根据年干+日支计算
	days := jfj.FindDays(lunarmjd, ydx)
	dzs, _ := jfj.QiShaDay(ygz, dgz)
	jsarr := jfj.ListQS(days, dzs)
	other := Other{
		JinFu:     jfs,
		JinShenQS: jsarr,
	}

	//今日信息
	lunary, lunarm, lunard, lunarh, leapmB := today.TodayT(T)
	var lts string
	if leapmB == true {
		lts = fmt.Sprintf("--今日农历:%d年%d月%d日%d时-(%s时) 当前月份是闰月", lunary, lunarm, lunard, lunarh, convHToName(lunarh))
	} else if leapmB == false {
		lts = fmt.Sprintf("今日农历:%d年%d月%d日%d时-(%s时) 当前月份不是闰月", lunary, lunarm, lunard, lunarh, convHToName(lunarh))

	}
	//关于
	about := about()
	/////
	resp := Resp{
		TodayS:   lts,
		JiNian:   jinianinfo,
		Dmj:      dmjinfo,
		Jq:       jq24info,
		ListDay:  listday,
		StarName: starName,
		StarInfo: zhisus,
		Zeji:     zeji,
		JiGan:    jgName,
		JiRi:     goodjg,
		QiSha:    qsdays,
		XJBF:     xieJiBFS,
		//TSInfo:   tsinfo,
		YJ:    yj,
		SJQM:  Sjqm,
		Gdtm:  gdtm,
		Other: other,
		About: about,
	}
	json.NewEncoder(w).Encode(resp)

}
func NewZRYL(ygz, mgz, dgz, hgz, yeargan, yearzhi string,
	aliaslmonth, aliaslday, aliaslhour, aliaslydxs string,
	lyear, lmonth, lday, lhour, lydx int,
	stime time.Time, syear, smonth, sday int, sweek string,
	leapmb bool, leapmonth int, lunarmjd float64) *xjbfs.ZRYL {
	//fmt.Printf(stime.Format("2006-01-02"))
	return xjbfs.NewZRYL(ygz, mgz, dgz, hgz, yeargan, yearzhi,
		aliaslmonth, aliaslday, aliaslhour, aliaslydxs,
		lyear, lmonth, lday, lhour, lydx,
		stime, syear, smonth, sday, sweek,
		leapmb, leapmonth, lunarmjd)
}
