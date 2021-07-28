package main

import (
	"fmt"
	"strings"
)

//值太岁
func ZhiTaiSui(ygz string) (zs string) {
	zhi := []string{"寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥", "子", "丑"}
	for i := 0; i < len(zhi); i++ {
		if strings.ContainsAny(ygz, zhi[i]) {
			zis := convZhiToSX(zhi[i])
			zs = fmt.Sprintf("生肖:%s值太岁", zis)
			break
		}
	}
	return
}

//冲太岁
func ChongTaiSui(ygz string) (cs string) {
	chmap := map[string]string{
		"子": "午", "丑": "未", "寅": "申", "卯": "酉", "辰": "戌", "巳": "亥",
		"午": "子", "未": "丑", "申": "寅", "酉": "卯", "戌": "辰", "亥": "巳",
	}
	for yz, ch := range chmap {
		if strings.ContainsAny(ygz, yz) {
			ch = convZhiToSX(ch)
			cs = fmt.Sprintf("生肖:%s冲太岁", ch)
			break
		}
	}
	return
}

//害太岁
func HaiTaiSui(ygz string) (hs string) {
	hmap := map[string]string{
		"子": "未", "丑": "午", "寅": "巳", "卯": "辰", "辰": "卯", "巳": "寅",
		"午": "丑", "未": "子", "申": "亥", "酉": "戌", "戌": "酉", "亥": "申",
	}
	for mz, h := range hmap {
		if strings.ContainsAny(ygz, mz) {
			h = convZhiToSX(h)
			hs = fmt.Sprintf("生肖:%s害太岁", h)
			break
		}
	}
	return
}

//刑太岁
func XingTaiSui(ygz string) (xs string) {
	xmap := map[string]string{
		"子": "卯", "丑": "戌", "寅": "巳", "卯": "子", "辰": "辰", "巳": "申",
		"午": "午", "未": "丑", "申": "寅", "酉": "酉", "戌": "未", "亥": "亥",
	}
	for mz, x := range xmap {
		if strings.ContainsAny(ygz, mz) {
			x = convZhiToSX(x)
			xs = fmt.Sprintf("生肖:%s刑太岁", x)
			break
		}
	}
	return
}

//地支转生肖
func convZhiToSX(s string) (sx string) {
	sxmap := map[string]string{
		"子": "鼠", "丑": "牛", "寅": "虎", "卯": "兔", "辰": "龙", "巳": "蛇",
		"午": "马", "未": "羊", "申": "猴", "酉": "鸡", "戌": "狗", "亥": "猪",
	}
	for k, v := range sxmap {
		if strings.ContainsAny(s, k) {
			sx = v
		}
	}
	return
}
