package main

import "strings"

func about() About {
	ccal := "农历 择吉 可计算时间范围:1601～3498"
	data := "下载: https://github.com/Aquarian-Age/ccal-gui/releases"
	xlr := "小六壬择吉 依据道家小六壬择法卷"
	xjbf := "择日 依据协纪辩方书"
	ck := "农历编算参考: https://ytliu0.github.io/ChineseCalendar/index_simp.html"
	me := "UVHnvqQgNzgzMjk0NjU1Cg=="
	return About{
		Ccal: ccal,
		Data: data,
		Xlr:  xlr,
		Xjbf: xjbf,
		Ck:   ck,
		Me:   me,
	}
}

func convHToName(h int) (hs string) {
	zhi := []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}
	for i := 1; i <= 12; i++ {
		if i == h {
			hs = zhi[i-1]
		}
	}
	return
}

//时辰地支
func convHourZhi(hourgz string) (alias string) {
	zhi := []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}
	for i := 0; i < len(zhi); i++ {
		if strings.ContainsAny(hourgz, zhi[i]) {
			alias = zhi[i]
			break
		}
	}
	return
}

//日名称 这里用的是廿
func convRmc(n int) (alias string) {
	rmc := []string{
		"初一", "初二", "初三", "初四", "初五", "初六", "初七", "初八", "初九", "初十",
		"十一", "十二", "十三", "十四", "十五", "十六", "十七", "十八", "十九", "二十",
		"廿一", "廿二", "廿三", "廿四", "廿五", "廿六", "廿七", "廿八", "廿九", "三十"}
	for i := 0; i < len(rmc); i++ {
		if i+1 == n {
			alias = rmc[i]
			break
		}
	}
	return
}

//农历数字专汉字
func convYmc(n int) (alias string) {
	ymc := []string{"正", "二", "三", "四", "五", "六", "七", "八", "九", "十", "十一", "十二"}
	for i := 0; i < len(ymc); i++ {
		if i+1 == n {
			alias = ymc[i]
			break
		}
	}
	return
}

//干支对应的月名称 正 二 三...十一 十二
func gztoYmc(gz string) (ym string) {
	ymc := []string{"正", "二", "三", "四", "五", "六", "七", "八", "九", "十", "十一", "十二"}
	zhi := []string{"寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥", "子", "丑"}
	for i := 0; i < len(zhi); i++ {
		if strings.ContainsAny(gz, zhi[i]) {
			switch i { //0正月
			case 0:
				ym = ymc[0]
			case 1:
				ym = ymc[1]
			case 2:
				ym = ymc[2]
			case 3:
				ym = ymc[3]
			case 4:
				ym = ymc[4]
			case 5:
				ym = ymc[5]
			case 6:
				ym = ymc[6]
			case 7:
				ym = ymc[7]
			case 8:
				ym = ymc[8]
			case 9:
				ym = ymc[9]
			case 10:
				ym = ymc[10]
			case 11:
				ym = ymc[11]
			}
			break
		}
	}
	return
}
