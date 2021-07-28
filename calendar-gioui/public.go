package main

import (
	"4d63.com/tz"
	"fmt"
	jl "github.com/soniakeys/meeus/v3/julian"
	"math"
	"os"
	"time"
)

//从上年冬至到下年立春的节气名称
func NewJQArr(y int) *JQArr {
	jqArrT := JqT(y)
	var jmc []string
	var jqt []time.Time
	for i := 0; i < len(jqArrT); i++ {
		jmc = append(jmc, Jmc[i])
		jqt = append(jqt, jqArrT[i])
	}
	jqarr := &JQArr{
		Name: jmc,
		Time: jqt,
	}
	return jqarr
}

//节气相关计算
//阳历年份节气数组 冬至--冬至
func JqT(y int) []time.Time {
	var j24t []time.Time
	data := Data(y)
	for i := 1; i <= 25; i++ {
		hs := data[0] + data[i]
		hs8 := JdToLocalTime(hs)
		j24t = append(j24t, hs8)
	}
	return j24t
}

//立春和指定时间的比较值 true在立春之后或者等于立春 false在立春之前 同时返回立春时间戳(精确到日)
func (jqarr *JQArr) LinChuT(cust time.Time) (bool, time.Time) {
	m1t := jqarr.Time[3]
	m1t = time.Date(m1t.Year(), time.Month(int(m1t.Month())), m1t.Day(), 0, 0, 0, 0, time.Local)
	cust = time.Date(cust.Year(), cust.Month(), cust.Day(), 0, 0, 0, 0, time.Local)

	var b bool
	if cust.Equal(m1t) || cust.After(m1t) {
		b = true
	} else {
		b = false
	}
	return b, m1t
}

//节气数组 0大雪 1小寒 2立春 3惊蛰 4清明 5立夏 6芒种 7小暑 8立秋 9白露 10寒露 11立东
func (jqarr *JQArr) Jie12ArrT() []time.Time {
	var jqArrT []time.Time
	for i := 1; i < len(jqarr.Time); i += 2 {
		jqt := jqarr.Time[i]
		jqArrT = append(jqArrT, jqt)
	}

	daxuet := jqArrT[11]
	times := jqArrT[:11]

	var j12T []time.Time
	j12T = append(j12T, daxuet)
	j12T = append(j12T, times...)
	return j12T
}

//中气 从上年冬至开始 13个中气 0:冬至 大寒 雨水 春分 谷雨 小满 夏至 大暑 处暑 秋分 霜降 小雪 冬至
func (jqarr *JQArr) ZhongQiArrT() []time.Time {
	var zqArrT []time.Time
	for i := 0; i < len(jqarr.Time); i += 2 {
		zqt := jqarr.Time[i]
		zqArrT = append(zqArrT, zqt)
	}
	return zqArrT
}

//y,m,d时间对比12节时间 true在节之后 false在节之前
func DiffJieT(y, m, d int, moonShuotj12arrt []*MoonShuoTJ12T) (bool, int, int) {
	t := time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.Local)
	var b bool
	var indexMGZ int
	var indexswN int
	for i := 0; i < len(moonShuotj12arrt); i++ {
		jiet := moonShuotj12arrt[i].MoonJieT
		jiet = time.Date(jiet.Year(), jiet.Month(), jiet.Day(), 0, 0, 0, 0, time.Local)
		difft := t.Sub(jiet).Hours() / 24
		m := i + 11
		if m > 12 {
			m -= 12
		}
		if (int(difft) >= 0 && int(difft) < 30) && (t.Equal(jiet) || t.Before(jiet)) {
			indexswN = i
			indexMGZ = m
			b = true
		} else if (int(difft) >= 0 && int(difft) < 30) && (t.After(jiet)) {
			indexswN = i
			indexMGZ = m
			b = true
		} else if int(difft) < 0 && int(math.Abs(difft)) < 30 {
			indexswN = i - 1
			indexMGZ = m - 1
			if indexMGZ == 0 {
				indexMGZ = 12
			}
			b = true
		}
	}
	return b, indexMGZ, indexswN
}

//农历朔和本月节  0:上年农历十一月(这里没计算农历闰月 所以农历月份要待计算闰月之后才能最终确定)
//这里的T都是阳历时间戳 精确到秒
func MoonShuoTJ12ArrT(j12arrt []time.Time, shuowangt []*ShuoWangT) []*MoonShuoTJ12T {
	var msjArrT []*MoonShuoTJ12T
	for i := 0; i < len(j12arrt); i++ { //0:农历上年11月的节
		for j := i; j < len(shuowangt); j++ {
			moonshuot := shuowangt[j].ShuoT
			moonj12t := j12arrt[j]
			moonsj12T := &MoonShuoTJ12T{
				MoonShuoT: moonshuot,
				MoonJieT:  moonj12t,
			}
			msjArrT = append(msjArrT, moonsj12T)
			break
		}
	}
	return msjArrT
}

//十四个月的朔 上弦 望 下弦 时间精确到分钟
//0:上年十一月 ... (含闰月的年份 14:本年十二月) (不含润月的年份 14:下一年正月)
func MoonShuoWangT(shuoWObj []*ShuoWangF) []*ShuoWangT {
	var swTArrT []*ShuoWangT
	for i := 0; i < len(shuoWObj); i++ {
		shuojd := shuoWObj[i].ShuoF
		shangjd := shuoWObj[i].ShangXianF
		wangjd := shuoWObj[i].WangF
		xiajd := shuoWObj[i].XiaXianF
		shuot := JdToLocalTime(shuojd)
		shangt := JdToLocalTime(shangjd)
		wangt := JdToLocalTime(wangjd)
		xiat := JdToLocalTime(xiajd)
		swT := &ShuoWangT{
			ShuoT:      shuot,
			ShangXianT: shangt,
			WangT:      wangt,
			XiaXianT:   xiat,
		}
		swTArrT = append(swTArrT, swT)
	}
	return swTArrT
}

//十五个月的朔望数据 0:农历上年十一月:{朔 :上弦 :望 :下弦} 下标1:农历上年十二月 下标3农历本年正月
func MoonShuoWangF(shuoWangObjArr []*ShuoWangF) []*ShuoWangF {
	var moonArr []*ShuoWangF
	for i := 3; i < len(shuoWangObjArr); i += 4 { //从索引3开始+=4取最后这组结构体值
		moonArr = append(moonArr, shuoWangObjArr[i])
	}
	return moonArr
}

//JD数据 从农历上一年十一月开始到本年十月结束
func NewShuoWangF(index int, data []float64) []*ShuoWangF {
	var indexShuo, indexShangXian, indexWang, indexXiaXian float64
	var swArr []*ShuoWangF
	for i := index; i < len(data); i++ {
		switch (i - index) % 4 {
		case 0:
			indexShuo = data[0] + data[i]
		case 1:
			indexShangXian = data[0] + data[i]
		case 2:
			indexWang = data[0] + data[i]
		case 3:
			indexXiaXian = data[0] + data[i]
		}
		sw := &ShuoWangF{
			ShuoF:      indexShuo,
			ShangXianF: indexShangXian,
			WangF:      indexWang,
			XiaXianF:   indexXiaXian,
		}
		swArr = append(swArr, sw)
	}
	return swArr
}

func LeapmB(y int) bool {
	data := Data(y)
	_, m1t, _ := m1(data)
	m11t, _ := m11(data)
	d := m11t.Sub(m1t).Hours() / 24
	l := d / 29.53
	x := math.Dim(l, 12.0) > 0.9 //true两冬至间有不含中气的闰月
	return x
}

//农历全年合朔(2个冬至间的朔数据)
func AllShuo(y int) []float64 {
	data1 := Data(y)
	data11 := Data(y + 1)
	index1, _, m1jd0 := m1(data1)
	_, m1jd1 := m11(data1)
	index11, _, m11jd0 := m1(data11)
	_, m11jd1 := m11(data11)
	shuo1 := Shuox(index1, m1jd0, m1jd1, data1)
	shuo11 := Shuox(index11, m11jd0, m11jd1, data11)

	var allshuo []float64
	shuo1b := EqualT(shuo1[len(shuo1)-1], shuo11[0])
	shuo11b := EqualT(shuo1[len(shuo1)-2], shuo11[0])
	if shuo1b == true || shuo11b == true { //过滤掉重复的数据
		allshuo = append(shuo1, shuo11[1:]...)
	}
	return allshuo
}

//朔的时间戳比较 时间精确到小时
func EqualT(jd1, jd11 float64) bool {
	t1 := JdToLocalTime(jd1)
	t11 := JdToLocalTime(jd11)
	t1 = time.Date(t1.Year(), t1.Month(), t1.Day(), t1.Hour(), 0, 0, 0, time.Local)
	t11 = time.Date(t11.Year(), t11.Month(), t11.Day(), t11.Hour(), 0, 0, 0, time.Local)

	return t1.Equal(t11)
}

//两冬至间朔 冬至前距离最近的朔 11月到11月
func Shuox(index int, m1jd, m11jd float64, datax []float64) (lunarShuox []float64) {
	var hsjdarr []float64
	for i := 26; i < len(datax); i += 4 {
		_hsjd := datax[0] + datax[i]
		hsjdarr = append(hsjdarr, _hsjd)
		jd11 := datax[0] + datax[i]
		if math.Dim(jd11, m11jd) < 0.9 && jd11 == m11jd {
			break
		}
	}
	lunarShuox = hsjdarr
	return
}

//第一个冬至前的朔 时间精确到日
func m1(data []float64) (index int, m1t time.Time, m1jd float64) {
	jd0 := data[0]
	dz0 := jd0 + data[1]
	dz0t := JdToLocalTime(dz0)
	for i := 26; i < len(data); i += 4 {
		jdx := jd0 + data[i]
		jdxt := JdToLocalTime(jdx)
		dz0t = time.Date(dz0t.Year(), dz0t.Month(), dz0t.Day(), 0, 0, 0, 0, time.UTC)
		jdxt = time.Date(jdxt.Year(), jdxt.Month(), jdxt.Day(), 0, 0, 0, 0, time.UTC)

		if jdxt.Equal(dz0t) == true {
			index = i
			m1jd = jdx
			m1t = JdToLocalTime(jdx)
		} else if jdxt.Equal(dz0t) == false && jdxt.Before(dz0t) {
			index = i
			m1jd = jdx
			m1t = JdToLocalTime(jdx)
		}
	}
	return
}

//第二个冬至前的朔 时间精确到日
func m11(data []float64) (m11t time.Time, m11jd float64) {
	jd0 := data[0]
	dz1 := jd0 + data[25]
	dz1t := JdToLocalTime(dz1)
	for i := 74; i < len(data); i += 4 {
		jdx := jd0 + data[i]
		jdxt := JdToLocalTime(jdx)
		dz1t = time.Date(dz1t.Year(), dz1t.Month(), dz1t.Day(), 0, 0, 0, 0, time.UTC)
		jdxt = time.Date(jdxt.Year(), jdxt.Month(), jdxt.Day(), 0, 0, 0, 0, time.UTC)
		if jdxt.Equal(dz1t) == true {
			m11jd = jdx
			m11t = JdToLocalTime(jdx)
		} else if jdxt.Equal(dz1t) == false && jdxt.Before(dz1t) {
			m11jd = jdx
			m11t = JdToLocalTime(jdx)
		}
	}
	return
}

//获取年份数据
func Data(y int) (data []float64) {
	for k, v := range mapOfYears {
		if k == y {
			data = v
			break
		}
	}
	return
}

//儒略日转time.Time 这里精确到秒
func JdToLocalTime(jd float64) time.Time {
	FormatT := "2006-01-02 15:04:05"
	cst := localTime(jd)
	t, err := time.Parse(FormatT, cst)
	if err != nil {
		fmt.Println("JDToLocalTime解析异常:", err)
		os.Exit(0)
	}
	return t
}

//转为为本地时间(Asia/Shanghai)
func localTime(jd float64) string {
	utc := jl.JDToTime(jd)
	local := utc
	location, err := tz.LoadLocation("Asia/Shanghai")
	if err == nil {
		local = local.In(location)
	}
	s := local.Format("2006-01-02 15:04:05")
	return s
}

//阳历年月日转换为ind类型的JD数据
func SolarYmdToJD(y, m, d int) (solarJd int) {
	ra := (14 - m) / 12
	ry := y + 4800 - ra
	rm := m + 12*ra - 3
	solarJd = (d + (153*rm+2)/5 + 365*ry + (ry / 4) - (ry / 100) + (ry / 400) - 32045)
	return
}
func covnH24Toh12(h int) int {
	var h12 int
	switch h {
	case 23:
		h12 = 1
	case 00:
		h12 = 1
	case 1:
		h12 = 2
	case 2:
		h12 = 2
	case 3:
		h12 = 3
	case 4:
		h12 = 3
	case 5:
		h12 = 4
	case 6:
		h12 = 4
	case 7:
		h12 = 5
	case 8:
		h12 = 5
	case 9:
		h12 = 6
	case 10:
		h12 = 6
	case 11:
		h12 = 7
	case 12:
		h12 = 7
	case 13:
		h12 = 8
	case 14:
		h12 = 8
	case 15:
		h12 = 9
	case 16:
		h12 = 9
	case 17:
		h12 = 10
	case 18:
		h12 = 10
	case 19:
		h12 = 11
	case 20:
		h12 = 11
	case 21:
		h12 = 12
	case 22:
		h12 = 12
	}
	return h12
}
