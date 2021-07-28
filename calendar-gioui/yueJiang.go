package main

import (
	"fmt"
	"strings"
	"time"
)

//二十四节气名称
var JMC = []string{
	"冬至", "小寒", "大寒", "立春", "雨水", "惊蛰",
	"春分", "清明", "谷雨", "立夏", "小满", "芒种",
	"夏至", "小暑", "大暑", "立秋", "处暑", "白露",
	"秋分", "寒露", "霜降", "立冬", "小雪", "大雪",
}

//从上年冬至到下年立春的节气名称　节气时间戳
// 冬至 小寒 大寒 立春 雨水 惊蛰 春分 清明 谷雨 立夏 小满 芒种
//夏至 小暑 大暑 立秋 处暑 白露 秋分 寒露 霜降 立冬 小雪 大雪
//冬至 小寒 大寒 立春
type YueJiangJQ struct {
	Name []string    //节气名称数组
	Time []time.Time //节气时间数组 精确到分钟秒钟
}

////////////
//月将名称 月将对应的地支 十二宫星
type YueJiang struct {
	Name  string `json:"name"`   //月将
	Zhi   string `json:"zhi"`    //月将对应的地支
	Star  string `json:"star"`   //黄道十二宫星名称
	JieQi string `json:"jie_qi"` //节气
}

//节气时间数组　精确到分钟
//上一年冬至(含)开始 0:冬至　1:小寒　2:大寒...
func YueJiangJQArrT(y int) []time.Time { //y:农历年数字
	var j24t []time.Time
	data := Data(y)
	for i := 1; i <= 25; i++ { //<=25冬至到冬至
		hs := data[0] + data[i]  //UTC合朔
		hs8 := JdToLocalTime(hs) //CST+8
		j24t = append(j24t, hs8)
	}

	var j24t1 []time.Time
	data1 := Data(y + 1)
	for j := 1; j <= 25; j++ {
		hs := data1[0] + data1[j] //UTC合朔
		hs8 := JdToLocalTime(hs)  //CST+8
		j24t1 = append(j24t1, hs8)
	}

	//去重
	var kx int
I:
	for k := 0; k < len(j24t); k++ {
		xk := j24t[k]
		for k1 := 0; k1 < len(j24t1); k1++ {
			xk1 := j24t1[k1]
			if b := xk.Equal(xk1); b == true {
				kx = k1
				break I
			}
		}
	}
	//上一年冬至(含)开始 0:冬至　1:小寒　2:大寒...
	alljqt := append(j24t, j24t1[kx+1:]...)
	return alljqt
}

//上一年冬至到下一年立春的节气
func NewYueJiangJQ(jqt []time.Time) *YueJiangJQ {
	jq := new(YueJiangJQ)

	var jqmc []string     //节气名称
	var jqmct []time.Time //节气时间戳
	//i=28到下一年立春
	for i := 0; i < 28; i++ {
		x := i
		if x > 23 {
			x = i - 24
		}
		jqmc = append(jqmc, JMC[x])
		jqmct = append(jqmct, jqt[i])
	}
	jq = &YueJiangJQ{
		Name: jqmc,
		Time: jqmct,
	}
	return jq
}

//月将
func (jq *YueJiangJQ) YueJiang(cust time.Time) (yj *YueJiang) {
	yj = new(YueJiang)
	//jqt:24节气的时间数组,由JQT()函数生成 cust:当前阳历时间戳
	jqt := jq.Time
	dht := jqt[2]           //上一年大寒
	lichunt := jqt[3]       //立春日
	yushuit := jqt[4]       //雨水日
	jingzhet := jqt[5]      //惊蛰
	chunfent := jqt[6]      //春分
	qingmingt := jqt[7]     //清明
	guyut := jqt[8]         //谷雨
	lixiat := jqt[9]        //立夏
	xiaomant := jqt[10]     //小满
	mangzhongt := jqt[11]   //芒种
	xiazhit := jqt[12]      //夏至
	xiaoshut := jqt[13]     //小暑
	dashut := jqt[14]       //大暑
	liqiut := jqt[15]       //立秋
	chushut := jqt[16]      //处暑
	bailut := jqt[17]       //白露
	qiufent := jqt[18]      //秋分
	hanlut := jqt[19]       //寒露
	shuangjiangt := jqt[20] //霜降
	lidongt := jqt[21]      //立冬
	xiaoxuet := jqt[22]     //小雪
	daxuet := jqt[23]       //大雪
	dongzhit := jqt[24]     //冬至
	xiaohant := jqt[25]     //小寒
	dahant := jqt[26]       //大寒
	lichun2t := jqt[27]     //次年立春

	// 时间精确到日
	lichunt = time.Date(lichunt.Year(), lichunt.Month(), lichunt.Day(), 0, 0, 0, 0, time.UTC)
	yushuit = time.Date(yushuit.Year(), yushuit.Month(), yushuit.Day(), 0, 0, 0, 0, time.UTC)
	jingzhet = time.Date(jingzhet.Year(), jingzhet.Month(), jingzhet.Day(), 0, 0, 0, 0, time.UTC)
	chunfent = time.Date(chunfent.Year(), chunfent.Month(), chunfent.Day(), 0, 0, 0, 0, time.UTC)
	qingmingt = time.Date(qingmingt.Year(), qingmingt.Month(), qingmingt.Day(), 0, 0, 0, 0, time.UTC)
	guyut = time.Date(guyut.Year(), guyut.Month(), guyut.Day(), 0, 0, 0, 0, time.UTC)
	lixiat = time.Date(lixiat.Year(), lixiat.Month(), lixiat.Day(), 0, 0, 0, 0, time.UTC)
	xiaomant = time.Date(xiaomant.Year(), xiaomant.Month(), xiaomant.Day(), 0, 0, 0, 0, time.UTC)
	mangzhongt = time.Date(mangzhongt.Year(), mangzhongt.Month(), mangzhongt.Day(), 0, 0, 0, 0, time.UTC)
	xiazhit = time.Date(xiazhit.Year(), xiazhit.Month(), xiazhit.Day(), 0, 0, 0, 0, time.UTC)
	xiaoshut = time.Date(xiaoshut.Year(), xiaoshut.Month(), xiaoshut.Day(), 0, 0, 0, 0, time.UTC)
	dashut = time.Date(dashut.Year(), dashut.Month(), dashut.Day(), 0, 0, 0, 0, time.UTC)
	liqiut = time.Date(liqiut.Year(), liqiut.Month(), liqiut.Day(), 0, 0, 0, 0, time.UTC)
	chushut = time.Date(chushut.Year(), chushut.Month(), chushut.Day(), 0, 0, 0, 0, time.UTC)
	bailut = time.Date(bailut.Year(), bailut.Month(), bailut.Day(), 0, 0, 0, 0, time.UTC)
	qiufent = time.Date(qiufent.Year(), qiufent.Month(), qiufent.Day(), 0, 0, 0, 0, time.UTC)
	hanlut = time.Date(hanlut.Year(), hanlut.Month(), hanlut.Day(), 0, 0, 0, 0, time.UTC)
	shuangjiangt = time.Date(shuangjiangt.Year(), shuangjiangt.Month(), shuangjiangt.Day(), 0, 0, 0, 0, time.UTC)
	lidongt = time.Date(lidongt.Year(), lidongt.Month(), lidongt.Day(), 0, 0, 0, 0, time.UTC)
	xiaoxuet = time.Date(xiaoxuet.Year(), xiaoxuet.Month(), xiaoxuet.Day(), 0, 0, 0, 0, time.UTC)
	daxuet = time.Date(daxuet.Year(), daxuet.Month(), daxuet.Day(), 0, 0, 0, 0, time.UTC)
	dongzhit = time.Date(dongzhit.Year(), dongzhit.Month(), dongzhit.Day(), 0, 0, 0, 0, time.UTC)
	xiaohant = time.Date(xiaohant.Year(), xiaohant.Month(), xiaohant.Day(), 0, 0, 0, 0, time.UTC)
	dahant = time.Date(dahant.Year(), dahant.Month(), dahant.Day(), 0, 0, 0, 0, time.UTC)
	lichun2t = time.Date(lichun2t.Year(), lichun2t.Month(), lichun2t.Day(), 0, 0, 0, 0, time.UTC)

	cust = time.Date(cust.Year(), cust.Month(), cust.Day(), 0, 0, 0, 0, time.UTC)

	//十二神将
	sj12 := []string{"登明", "河魁", "从魁", "传送", "小吉", "胜光", "太乙", "天罡", "太冲", "功曹", "大吉", "神后"}
	//黄道十二宫星
	star := []string{"白羊宫", "金牛宫", "双子宫", "巨蟹宫", "狮子宫", "室女宫", "天秤宫", "天蝎宫", "人马宫", "魔羯宫", "宝瓶宫", "双鱼宫"}
	//天月将地支
	tyj := []string{"亥", "戌", "酉", "申", "未", "午", "巳", "辰", "卯", "寅", "丑", "子"}
	//上一年大寒后 今年立春前
	//大寒后至雨水前 宝瓶座
	//大寒后太阳在子 神后为天月将
	if (cust.After(dht) || cust.Equal(dht)) && cust.Before(lichunt) {
		yj = &YueJiang{
			Name:  sj12[11],
			Zhi:   tyj[11],
			Star:  star[10],
			JieQi: fmt.Sprintf("大寒(%v)", dht.Format("2006-01-02")),
		}
	}
	//立春后（含立春日）雨水前
	//大寒后至雨水前 宝瓶座
	//立春后太阳在子 神后为天月将
	if (cust.After(lichunt) || cust.Equal(lichunt)) && cust.Before(yushuit) {
		yj = &YueJiang{
			Name:  sj12[11],
			Zhi:   tyj[11],
			Star:  star[10],
			JieQi: fmt.Sprintf("立春/节气(%v)", lichunt.Format("2006-01-02")),
		}
	}
	//雨水后(含雨水日)惊蛰前
	//雨水后至春分点前 双鱼宫
	//雨水后太阳过亥 登明为天月将
	if (cust.After(yushuit) || cust.Equal(yushuit)) && cust.Before(jingzhet) {
		yj = &YueJiang{
			Name:  sj12[0],
			Zhi:   tyj[0],
			Star:  star[11],
			JieQi: fmt.Sprintf("雨水/中气(%v)", yushuit.Format("2006-01-02")),
		}
	}
	//惊蛰后(含)春分前
	//雨水后至春分点前 双鱼宫
	//惊蛰后太阳在亥 登明为天月将
	if (cust.After(jingzhet) || cust.Equal(jingzhet)) && cust.Before(chunfent) {
		yj = &YueJiang{
			Name:  sj12[0],
			Zhi:   tyj[0],
			Star:  star[11],
			JieQi: fmt.Sprintf("惊蛰/节气(%v)", jingzhet.Format("2006-01-02")),
		}
	}
	//春分后清明前
	//春分后至谷雨前 白羊宫
	//春分后太阳过戌 河魁为天月将
	if (cust.After(chunfent) || cust.Equal(chunfent)) && cust.Before(qingmingt) {
		yj = &YueJiang{
			Name:  sj12[1],
			Zhi:   tyj[1],
			Star:  star[0],
			JieQi: fmt.Sprintf("春分/中气(%v)", chunfent.Format("2006-01-02")),
		}
	}
	//清明后谷雨前
	//春分后至谷雨前 白羊宫
	//清明后后太阳在戌 河魁为天月将
	if (cust.After(qingmingt) || cust.Equal(qingmingt)) && cust.Before(guyut) {
		yj = &YueJiang{
			Name:  sj12[1],
			Zhi:   tyj[1],
			Star:  star[0],
			JieQi: fmt.Sprintf("清明/节气(%v)", qingmingt.Format("2006-01-02")),
		}
	}
	//谷雨后立夏前
	//谷雨后至小满前 金牛宫
	//谷雨后太阳过酉 从魁为天月将
	if (cust.After(guyut) || cust.Equal(guyut)) && cust.Before(lixiat) {
		yj = &YueJiang{
			Name:  sj12[2],
			Zhi:   tyj[2],
			Star:  star[1],
			JieQi: fmt.Sprintf("谷雨/中气(%v)", guyut.Format("2006-01-02")),
		}
	}
	//立夏后小满前
	//谷雨后至小满前 金牛宫
	//立夏后太阳在酉 从魁为天月将
	if (cust.After(lixiat) || cust.Equal(lixiat)) && cust.Before(xiaomant) {
		yj = &YueJiang{
			Name:  sj12[2],
			Zhi:   tyj[2],
			Star:  star[1],
			JieQi: fmt.Sprintf("立夏/节气(%v)", lixiat.Format("2006-01-02")),
		}
	}
	//小满后忙钟前
	//小满后至夏至前 双子宫
	//小满后太阳过申 传送为天月将
	if (cust.After(xiaomant) || cust.Equal(xiaomant)) && cust.Before(mangzhongt) {
		yj = &YueJiang{
			Name:  sj12[3],
			Zhi:   tyj[3],
			Star:  star[2],
			JieQi: fmt.Sprintf("小满/中气(%v)", xiaomant.Format("2006-01-02")),
		}
	}
	//芒种后夏至前
	//小满后至夏至前 双子宫
	//芒种后太阳在申 传送为天月将
	if (cust.After(mangzhongt) || cust.Equal(mangzhongt)) && cust.Before(xiazhit) {
		yj = &YueJiang{
			Name:  sj12[3],
			Zhi:   tyj[3],
			Star:  star[2],
			JieQi: fmt.Sprintf("忙种/节气(%v)", mangzhongt.Format("2006-01-02")),
		}
	}
	//夏至后小暑前
	//夏至后至大暑前 巨蟹宫
	//夏至后太阳过未 小吉为天月将
	if (cust.After(xiazhit) || cust.Equal(xiazhit)) && cust.Before(xiaoshut) {
		yj = &YueJiang{
			Name:  sj12[4],
			Zhi:   tyj[4],
			Star:  star[3],
			JieQi: fmt.Sprintf("夏至/中气(%v)", xiazhit.Format("2006-01-02")),
		}
	}
	//小暑后大暑前
	//夏至后至大暑前 巨蟹宫
	//小暑后太阳在未 小吉为天月将
	if (cust.After(xiaoshut) || cust.Equal(xiaoshut)) && cust.Before(dashut) {
		yj = &YueJiang{
			Name:  sj12[4],
			Zhi:   tyj[4],
			Star:  star[3],
			JieQi: fmt.Sprintf("小暑/节气(%v)", xiaoshut.Format("2006-01-02")),
		}
	}
	//大暑后立秋前
	//大暑后至处暑前 狮子宫
	//大暑后太阳过午 胜光为天月将
	if (cust.After(dashut) || cust.Equal(dashut)) && cust.Before(liqiut) {
		yj = &YueJiang{
			Name:  sj12[5],
			Zhi:   tyj[5],
			Star:  star[4],
			JieQi: fmt.Sprintf("大暑/中气(%v)", dashut.Format("2006-01-02")),
		}
	}
	//立秋后处暑前
	//大暑后至处暑前 狮子宫
	//立秋后太阳在午 胜光为天月将
	if (cust.After(liqiut) || cust.Equal(liqiut)) && cust.Before(chushut) {
		yj = &YueJiang{
			Name:  sj12[5],
			Zhi:   tyj[5],
			Star:  star[4],
			JieQi: fmt.Sprintf("立秋/节气(%v)", liqiut.Format("2006-01-02")),
		}
	}
	//处暑后白露前
	//处暑后至秋分前 室女宫
	//处暑后太阳过巳 太乙为天月将
	if (cust.After(chushut) || cust.Equal(chushut)) && cust.Before(bailut) {
		yj = &YueJiang{
			Name:  sj12[6],
			Zhi:   tyj[6],
			Star:  star[5],
			JieQi: fmt.Sprintf("处暑/中气(%v)", chushut.Format("2006-01-02")),
		}
	}
	//白露后秋分前
	//处暑后至秋分前 室女宫
	//白露后太阳在巳 太乙为天月将
	if (cust.After(bailut) || cust.Equal(bailut)) && cust.Before(qiufent) {
		yj = &YueJiang{
			Name:  sj12[6],
			Zhi:   tyj[6],
			Star:  star[5],
			JieQi: fmt.Sprintf("白露/节气(%v)", bailut.Format("2006-01-02")),
		}
	}
	//秋分后寒露前
	//秋分直至霜降 天秤宫
	//秋分后太阳过辰 天罡为天月将
	if (cust.After(qiufent) || cust.Equal(qiufent)) && cust.Before(hanlut) {
		yj = &YueJiang{
			Name:  sj12[7],
			Zhi:   tyj[7],
			Star:  star[6],
			JieQi: fmt.Sprintf("秋分/中气(%v)", qiufent.Format("2006-01-02")),
		}
	}
	//寒露后霜降前
	//秋分直至霜降 天秤宫
	//寒露后太阳在辰 天罡为天月将
	if (cust.After(hanlut) || cust.Equal(hanlut)) && cust.Before(shuangjiangt) {
		yj = &YueJiang{
			Name:  sj12[7],
			Zhi:   tyj[7],
			Star:  star[6],
			JieQi: fmt.Sprintf("寒露/节气(%v)", hanlut.Format("2006-01-02")),
		}
	}
	//霜降后立冬前
	//霜降后至小雪前 天蝎宫
	//霜降后太阳过卯 太冲为天月将
	if (cust.After(shuangjiangt) || cust.Equal(shuangjiangt)) && cust.Before(lidongt) {
		yj = &YueJiang{
			Name:  sj12[8],
			Zhi:   tyj[8],
			Star:  star[7],
			JieQi: fmt.Sprintf("霜降/中气(%v)", shuangjiangt.Format("2006-01-02")),
		}
	}
	//立冬后小雪前
	//霜降后至小雪前 天蝎宫
	//立东后太阳在卯 太冲为天月将
	if (cust.After(lidongt) || cust.Equal(lidongt)) && cust.Before(xiaoxuet) {
		yj = &YueJiang{
			Name:  sj12[8],
			Zhi:   tyj[8],
			Star:  star[7],
			JieQi: fmt.Sprintf("立冬/节气(%v)", lidongt.Format("2006-01-02")),
		}
	}
	//小雪后大雪前
	//小雪后至冬至前 人马宫
	//小雪后太阳过寅 功曹为天月将
	if (cust.After(xiaoxuet) || cust.Equal(xiaoxuet)) && cust.Before(daxuet) {
		yj = &YueJiang{
			Name:  sj12[9],
			Zhi:   tyj[9],
			Star:  star[8],
			JieQi: fmt.Sprintf("小雪/中气(%v)", xiaoxuet.Format("2006-01-02")),
		}
	}
	//大雪后冬至前
	//小雪后至冬至前 人马宫
	//大雪后太阳在寅 功曹为天月将
	if (cust.After(daxuet) || cust.Equal(daxuet)) && cust.Before(dongzhit) {
		yj = &YueJiang{
			Name:  sj12[9],
			Zhi:   tyj[9],
			Star:  star[8],
			JieQi: fmt.Sprintf("大雪/节气(%v)", daxuet.Format("2006-01-02")),
		}
	}
	//冬至后小寒前
	//冬至后大寒前 魔羯宫
	//冬至后太阳过丑 大吉为天月将
	if (cust.After(dongzhit) || cust.Equal(dongzhit)) && cust.Before(xiaohant) {
		yj = &YueJiang{
			Name:  sj12[10],
			Zhi:   tyj[10],
			Star:  star[9],
			JieQi: fmt.Sprintf("冬至/中气(%v)", dongzhit.Format("2006-01-02")),
		}
	}
	//小寒后大寒前
	//冬至后大寒前 魔羯宫
	//小寒后太阳在丑 大吉为天月将
	if (cust.After(xiaohant) || cust.Equal(xiaohant)) && cust.Before(dahant) {
		yj = &YueJiang{
			Name:  sj12[10],
			Zhi:   tyj[10],
			Star:  star[9],
			JieQi: fmt.Sprintf("小寒/节气(%v)", xiaohant.Format("2006-01-02")),
		}
	}
	//大寒后立冬2前
	//大寒后至雨水前 宝瓶座
	//大寒后太阳过子 神后为天月将
	if (cust.After(dahant) || cust.Equal(dahant)) && cust.Before(lichun2t) {
		yj = &YueJiang{
			Name:  sj12[11],
			Zhi:   tyj[11],
			Star:  star[10],
			JieQi: fmt.Sprintf("大寒/中气(%v)", dahant.Format("2006-01-02")),
		}
	}
	return
}

//当前节气
func (jq *YueJiangJQ) TodayJQ(y int, cust time.Time) (jt string) {
	jqt := YueJiangJQArrT(y)
	dht := jqt[2]           //上一年大寒
	lichunt := jqt[3]       //立春日
	yushuit := jqt[4]       //雨水日
	jingzhet := jqt[5]      //惊蛰
	chunfent := jqt[6]      //春分
	qingmingt := jqt[7]     //清明
	guyut := jqt[8]         //谷雨
	lixiat := jqt[9]        //立夏
	xiaomant := jqt[10]     //小满
	mangzhongt := jqt[11]   //芒种
	xiazhit := jqt[12]      //夏至
	xiaoshut := jqt[13]     //小暑
	dashut := jqt[14]       //大暑
	liqiut := jqt[15]       //立秋
	chushut := jqt[16]      //处暑
	bailut := jqt[17]       //白露
	qiufent := jqt[18]      //秋分
	hanlut := jqt[19]       //寒露
	shuangjiangt := jqt[20] //霜降
	lidongt := jqt[21]      //立冬
	xiaoxuet := jqt[22]     //小雪
	daxuet := jqt[23]       //大雪
	dongzhit := jqt[24]     //冬至
	xiaohant := jqt[25]     //小寒
	dahant := jqt[26]       //大寒
	lichun2t := jqt[27]     //次年立春

	// 时间精确到日
	lichunt = time.Date(lichunt.Year(), lichunt.Month(), lichunt.Day(), 0, 0, 0, 0, time.Local)
	yushuit = time.Date(yushuit.Year(), yushuit.Month(), yushuit.Day(), 0, 0, 0, 0, time.Local)
	jingzhet = time.Date(jingzhet.Year(), jingzhet.Month(), jingzhet.Day(), 0, 0, 0, 0, time.Local)
	chunfent = time.Date(chunfent.Year(), chunfent.Month(), chunfent.Day(), 0, 0, 0, 0, time.Local)
	qingmingt = time.Date(qingmingt.Year(), qingmingt.Month(), qingmingt.Day(), 0, 0, 0, 0, time.Local)
	guyut = time.Date(guyut.Year(), guyut.Month(), guyut.Day(), 0, 0, 0, 0, time.Local)
	lixiat = time.Date(lixiat.Year(), lixiat.Month(), lixiat.Day(), 0, 0, 0, 0, time.Local)
	xiaomant = time.Date(xiaomant.Year(), xiaomant.Month(), xiaomant.Day(), 0, 0, 0, 0, time.Local)
	mangzhongt = time.Date(mangzhongt.Year(), mangzhongt.Month(), mangzhongt.Day(), 0, 0, 0, 0, time.Local)
	xiazhit = time.Date(xiazhit.Year(), xiazhit.Month(), xiazhit.Day(), 0, 0, 0, 0, time.Local)
	xiaoshut = time.Date(xiaoshut.Year(), xiaoshut.Month(), xiaoshut.Day(), 0, 0, 0, 0, time.Local)
	dashut = time.Date(dashut.Year(), dashut.Month(), dashut.Day(), 0, 0, 0, 0, time.Local)
	liqiut = time.Date(liqiut.Year(), liqiut.Month(), liqiut.Day(), 0, 0, 0, 0, time.Local)
	chushut = time.Date(chushut.Year(), chushut.Month(), chushut.Day(), 0, 0, 0, 0, time.Local)
	bailut = time.Date(bailut.Year(), bailut.Month(), bailut.Day(), 0, 0, 0, 0, time.Local)
	qiufent = time.Date(qiufent.Year(), qiufent.Month(), qiufent.Day(), 0, 0, 0, 0, time.Local)
	hanlut = time.Date(hanlut.Year(), hanlut.Month(), hanlut.Day(), 0, 0, 0, 0, time.Local)
	shuangjiangt = time.Date(shuangjiangt.Year(), shuangjiangt.Month(), shuangjiangt.Day(), 0, 0, 0, 0, time.Local)
	lidongt = time.Date(lidongt.Year(), lidongt.Month(), lidongt.Day(), 0, 0, 0, 0, time.Local)
	xiaoxuet = time.Date(xiaoxuet.Year(), xiaoxuet.Month(), xiaoxuet.Day(), 0, 0, 0, 0, time.Local)
	daxuet = time.Date(daxuet.Year(), daxuet.Month(), daxuet.Day(), 0, 0, 0, 0, time.Local)
	dongzhit = time.Date(dongzhit.Year(), dongzhit.Month(), dongzhit.Day(), 0, 0, 0, 0, time.Local)
	xiaohant = time.Date(xiaohant.Year(), xiaohant.Month(), xiaohant.Day(), 0, 0, 0, 0, time.Local)
	dahant = time.Date(dahant.Year(), dahant.Month(), dahant.Day(), 0, 0, 0, 0, time.Local)
	lichun2t = time.Date(lichun2t.Year(), lichun2t.Month(), lichun2t.Day(), 0, 0, 0, 0, time.Local)

	cust = time.Date(cust.Year(), cust.Month(), cust.Day(), 0, 0, 0, 0, time.Local)

	//上一年大寒后 今年立春前
	if (cust.After(dht) || cust.Equal(dht)) && cust.Before(lichunt) {
		jt = fmt.Sprintf("大寒(%v)", dht.Format("2006-01-02"))

	}
	//立春后（含立春日）雨水前
	if (cust.After(lichunt) || cust.Equal(lichunt)) && cust.Before(yushuit) {
		jt = fmt.Sprintf("立春/节气(%v)", lichunt.Format("2006-01-02"))

	}
	//雨水后(含雨水日)惊蛰前
	if (cust.After(yushuit) || cust.Equal(yushuit)) && cust.Before(jingzhet) {
		jt = fmt.Sprintf("雨水/中气(%v)", yushuit.Format("2006-01-02"))

	}
	//惊蛰后(含)春分前
	if (cust.After(jingzhet) || cust.Equal(jingzhet)) && cust.Before(chunfent) {
		jt = fmt.Sprintf("惊蛰/节气(%v)", jingzhet.Format("2006-01-02"))

	}
	//春分后清明前
	if (cust.After(chunfent) || cust.Equal(chunfent)) && cust.Before(qingmingt) {
		jt = fmt.Sprintf("春分/中气(%v)", chunfent.Format("2006-01-02"))

	}
	//清明后谷雨前
	if (cust.After(qingmingt) || cust.Equal(qingmingt)) && cust.Before(guyut) {
		jt = fmt.Sprintf("清明/节气(%v)", qingmingt.Format("2006-01-02"))

	}
	//谷雨后立夏前
	if (cust.After(guyut) || cust.Equal(guyut)) && cust.Before(lixiat) {
		jt = fmt.Sprintf("谷雨/中气(%v)", guyut.Format("2006-01-02"))

	}
	//立夏后小满前
	if (cust.After(lixiat) || cust.Equal(lixiat)) && cust.Before(xiaomant) {
		jt = fmt.Sprintf("立夏/节气(%v)", lixiat.Format("2006-01-02"))

	}
	//小满后忙钟前
	if (cust.After(xiaomant) || cust.Equal(xiaomant)) && cust.Before(mangzhongt) {
		jt = fmt.Sprintf("小满/中气(%v)", xiaomant.Format("2006-01-02"))

	}
	//芒种后夏至前
	if (cust.After(mangzhongt) || cust.Equal(mangzhongt)) && cust.Before(xiazhit) {
		jt = fmt.Sprintf("忙种/节气(%v)", mangzhongt.Format("2006-01-02"))

	}
	//夏至后小暑前
	if (cust.After(xiazhit) || cust.Equal(xiazhit)) && cust.Before(xiaoshut) {
		jt = fmt.Sprintf("夏至/中气(%v)", xiazhit.Format("2006-01-02"))

	}
	//小暑后大暑前
	if (cust.After(xiaoshut) || cust.Equal(xiaoshut)) && cust.Before(dashut) {
		jt = fmt.Sprintf("小暑/节气(%v)", xiaoshut.Format("2006-01-02"))

	}
	//大暑后立秋前
	if (cust.After(dashut) || cust.Equal(dashut)) && cust.Before(liqiut) {
		jt = fmt.Sprintf("大暑/中气(%v)", dashut.Format("2006-01-02"))

	}
	//立秋后处暑前
	if (cust.After(liqiut) || cust.Equal(liqiut)) && cust.Before(chushut) {
		jt = fmt.Sprintf("立秋/节气(%v)", liqiut.Format("2006-01-02"))

	}
	//处暑后白露前
	if (cust.After(chushut) || cust.Equal(chushut)) && cust.Before(bailut) {
		jt = fmt.Sprintf("处暑/中气(%v)", chushut.Format("2006-01-02"))

	}
	//白露后秋分前
	if (cust.After(bailut) || cust.Equal(bailut)) && cust.Before(qiufent) {
		jt = fmt.Sprintf("白露/节气(%v)", bailut.Format("2006-01-02"))

	}
	//秋分后寒露前
	if (cust.After(qiufent) || cust.Equal(qiufent)) && cust.Before(hanlut) {
		jt = fmt.Sprintf("秋分/中气(%v)", qiufent.Format("2006-01-02"))

	}
	//寒露后霜降前
	if (cust.After(hanlut) || cust.Equal(hanlut)) && cust.Before(shuangjiangt) {
		jt = fmt.Sprintf("寒露/节气(%v)", hanlut.Format("2006-01-02"))

	}
	//霜降后立冬前
	if (cust.After(shuangjiangt) || cust.Equal(shuangjiangt)) && cust.Before(lidongt) {
		jt = fmt.Sprintf("霜降/中气(%v)", shuangjiangt.Format("2006-01-02"))

	}
	//立冬后小雪前
	if (cust.After(lidongt) || cust.Equal(lidongt)) && cust.Before(xiaoxuet) {
		jt = fmt.Sprintf("立冬/节气(%v)", lidongt.Format("2006-01-02"))

	}
	//小雪后大雪前
	if (cust.After(xiaoxuet) || cust.Equal(xiaoxuet)) && cust.Before(daxuet) {
		jt = fmt.Sprintf("小雪/中气(%v)", xiaoxuet.Format("2006-01-02"))

	}
	//大雪后冬至前
	if (cust.After(daxuet) || cust.Equal(daxuet)) && cust.Before(dongzhit) {
		jt = fmt.Sprintf("大雪/节气(%v)", daxuet.Format("2006-01-02"))

	}
	//冬至后小寒前
	if (cust.After(dongzhit) || cust.Equal(dongzhit)) && cust.Before(xiaohant) {
		jt = fmt.Sprintf("冬至/中气(%v)", dongzhit.Format("2006-01-02"))

	}
	//小寒后大寒前
	if (cust.After(xiaohant) || cust.Equal(xiaohant)) && cust.Before(dahant) {
		jt = fmt.Sprintf("小寒/节气(%v)", xiaohant.Format("2006-01-02"))

	}
	//大寒后立冬2前
	if (cust.After(dahant) || cust.Equal(dahant)) && cust.Before(lichun2t) {
		jt = fmt.Sprintf("大寒/中气(%v)", dahant.Format("2006-01-02"))

	}
	return
}

//全年24节气 文本自动换行
func (jq *YueJiangJQ) JQ24() string {
	format := "2006-01-02 15:04:05"
	var jqArr []string
	for i := 3; i <= 26; i++ {
		jqx := jq.Name[i] + ":" + jq.Time[i].Format(format)
		jqArr = append(jqArr, jqx)
	}
	jqs := strings.Join(jqArr, "\n")
	return jqs
}
