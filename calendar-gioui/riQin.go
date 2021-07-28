package main

import "strings"

//name:当日值宿名称  qsName：七煞名称 b=true七煞 false非七煞
func QinXingDay(weekN int, dgz string) (name string, b bool, qsName string) {
	//weekN:当前周几　0周日 dayz:当日地支
	//周几的值宿=当前局的基数+7*当前周几+当前周几 如果大于28取余数即为28宿索引号

	zhi := []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}
	var dayz string
	for i := 0; i < len(zhi); i++ {
		if strings.ContainsAny(dgz, zhi[i]) {
			dayz = zhi[i]
			break
		}
	}
	switch dayz {
	case "亥", "卯", "未": //木局-->周日:昴日鸡值宿
		_w0 := 17                   //周日值宿对应的基数
		wn := _w0 + 7*weekN + weekN //当前周数字对应的28宿索引
		if wn > 28 {
			wn = wn % 28
		}
		name = XingSu28[wn] //当日值宿名称
		b, qsName = QiShaNameDay(wn)
	case "巳", "酉", "丑": //金局-->周日:房日兔值宿
		_w0 := 3
		wn := _w0 + 7*weekN + weekN
		if wn > 28 {
			wn = wn % 28
		}
		name = XingSu28[wn]
		b, qsName = QiShaNameDay(wn)
	case "寅", "午", "戌": //火局-->周日:星日马值宿
		_w0 := 24
		wn := _w0 + 7*weekN + weekN
		if wn > 28 {
			wn = wn % 28
		}
		name = XingSu28[wn]
		b, qsName = QiShaNameDay(wn)
	case "申", "子", "辰": //水局-->周日:虚日鼠值宿
		_w0 := 10
		wn := _w0 + 7*weekN + weekN
		if wn > 28 {
			wn = wn % 28
		}
		name = XingSu28[wn]
		b, qsName = QiShaNameDay(wn)
	}
	return
}

//七煞索引数字 角:0 亢:1 牛:8 奎:14  鬼:22 柳:23 星:24
//索引数字与二十八宿中的索引值对比 找出七煞 b=true当前为七煞
//wn:0-->周日
func QiShaNameDay(wn int) (b bool, qsname string) {
	if wn == 0 || wn == 1 || wn == 8 || wn == 14 || wn == 22 || wn == 23 || wn == 24 {
		qsname = XingSu28[wn]
		b = true
	} else {
		b = false
	}
	return
}
