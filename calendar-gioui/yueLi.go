package main

import (
	"fmt"
	"time"
)

//月历表
func NewYueLi(y, m, d int) *YueLi {
	cust := time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.Local)
	allshuo := AllShuo(y)
	leapmb := LeapmB(y)

	ydxarr := YueDaXiao(allshuo)
	jqarrOBj := NewJQArr(y)
	zqArrT := jqarrOBj.ZhongQiArrT() //len:13
	index, _, _ := m1(Data(y))
	shuoWangFArr := NewShuoWangF(index, Data(y))
	moonSWFarr := MoonShuoWangF(shuoWangFArr)
	moonSwArrT := MoonShuoWangT(moonSWFarr)
	leapmN, leapmT := findLeapmN(leapmb, zqArrT, moonSwArrT)
	mtarr := findLunarMN(allshuo)
	_, lm := LunarMN(cust, mtarr, leapmb, leapmT)
	ydx := findYdx(leapmb, leapmN, ydxarr, lm)

	ld, sarr, gzarr, rqarr, qsbarr := yueli(cust, ydx, allshuo) //这里不计算闰月
	var ylb = new(YueLi)
	ylb = &YueLi{
		LunarD: ld,
		SolarD: sarr,
		GzD:    gzarr,
		RiQin:  rqarr,
		QiShaB: qsbarr,
	}

	return ylb
}

//闰月月历
func leapYueli(leapmT, cust time.Time, leapmN int, mtarr []time.Time, ydxarr []int) ([]string, []string, []string) {
	var leapydx int
	var idx int
	for j := 0; j < 15; j++ {
		index := j + 11
		if index > 12 {
			index -= 12
		}
		if leapmN == index-1 && leapmT.Equal(mtarr[j]) {
			leapydx = ydxarr[j]
			idx = j
			break
		}
	}
	var ldayArr = []string{} //农历月数组 长度等于农历月大小数字
	var dayArr = []string{}  //阳历月数组 长度等于农历月大小数字
	var dgzArr = []string{}
	for i := 0; i < leapydx-1; i++ {
		t := mtarr[idx]
		t = time.Date(t.Year(), t.Month(), t.Day()+i, 0, 0, 0, 0, time.Local)
		dayjd := SolarYmdToJD(t.Year(), int(t.Month()), t.Day())
		dayt := JdToLocalTime(float64(dayjd)) //阳历时间
		dayt = time.Date(dayt.Year(), dayt.Month(), dayt.Day(), 0, 0, 0, 0, time.Local)
		day := fmt.Sprintf("%d/%d", int(dayt.Month()), dayt.Day())
		//阳历
		dayArr = append(dayArr, day)
		//农历
		ldayArr = append(ldayArr, "闰"+Rmc[i])
		dgz, _ := DayGZ(t.Year(), int(t.Month()), t.Day())
		dgzArr = append(dgzArr, dgz)
	}
	return ldayArr, dayArr, dgzArr
}

//农历日名称 对应的阳历日数字 农历日数组 阳历日数组 干支数组 日禽 七煞判断
func yueli(cust time.Time, ydx int, allshuo []float64) ([]string, []string, []string, []string, []bool) {
	cust = time.Date(cust.Year(), cust.Month(), cust.Day(), 0, 0, 0, 0, time.Local)
	var idx int //月大小数组索引值
	var ldayArr = []string{} //农历月数组 长度等于农历月大小数字
	var dayArr = []string{}  //阳历月数组 长度等于农历月大小数字
	var dgzArr = []string{}  //日干支
	var rqArr = []string{}   //日禽
	var qsArr = []bool{}     //七煞判断 true 七煞
	for i := 0; i < 15; i++ {
		jd0 := allshuo[i]
		jd1 := allshuo[i+1]
		t := JdToLocalTime(jd0)
		t1 := JdToLocalTime(jd1)
		t = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
		t1 = time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, time.Local)
		if (cust.Equal(t) || cust.After(t)) && cust.Before(t1) {
			//使用月大小计算
			index := i + 11
			if index > 12 {
				index -= 12
			}
			idx = i
			break
		}
	}
	for j := 0; j <= ydx-1; j++ { //j农历日数字
		jd1 := allshuo[idx] //月初jd
		t := JdToLocalTime(jd1)
		t = time.Date(t.Year(), t.Month(), t.Day()+j, 0, 0, 0, 0, time.Local)
		dayjd := SolarYmdToJD(t.Year(), int(t.Month()), t.Day())
		dayt := JdToLocalTime(float64(dayjd)) //阳历时间
		dayt = time.Date(dayt.Year(), dayt.Month(), dayt.Day(), 0, 0, 0, 0, time.Local)
		//农历
		ldayArr = append(ldayArr, Rmc[j])
		//阳历
		day := fmt.Sprintf("%d/%d", int(dayt.Month()), dayt.Day())
		dayArr = append(dayArr, day)
		//干支
		dgz, _ := DayGZ(t.Year(), int(t.Month()), t.Day())
		dgzArr = append(dgzArr, dgz)
		//日禽
		weekN := int(t.Weekday())
		rq, qsb, _ := QinXingDay(weekN, dgz)
		rqArr = append(rqArr, rq)
		qsArr = append(qsArr, qsb)
	}
	return ldayArr, dayArr, dgzArr, rqArr, qsArr
}

