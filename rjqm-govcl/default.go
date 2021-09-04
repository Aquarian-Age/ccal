/*
 * The code is automatically generated by the Goland.
 * Copyright © Aquarian-Age. All Rights Reserved.
 * Licensed under MIT.
 */

package main

import (
	"fmt"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types/colors"
	"liangzi.local/qx/pkg/rjqm"
	"strings"
)

func (f *TMainForm) OnFormDefault(y int, m int, d int, h int) {

	rqm := rjqm.NewRjqm(y, m, d, h)

	ymds := fmt.Sprintf("%s--%s--%s--%s", rqm.Ygz, rqm.Mgz, rqm.Dgz, rqm.Hgz)
	ymds += " " + rqm.YinYang + "遁"
	f.ymdl = vcl.NewLabel(f)
	f.ymdl.SetParent(f)
	f.ymdl.SetTop(740)
	f.ymdl.SetLeft(left1 - 20)
	f.ymdl.SetCaption(ymds)

	//--- 九星 ---
	star := rqm.Star

	f.kanStar = vcl.NewLabel(f)
	f.kanStar.SetParent(f)
	f.kanStar.SetTop(top3)
	f.kanStar.SetLeft(left2)
	//吉色 红
	if strings.EqualFold(star[1], "天乙") || strings.EqualFold(star[1], "青龙") ||
		strings.EqualFold(star[1], "太乙") || strings.EqualFold(star[1], "太阴") {
		f.kanStar.Font().SetColor(colors.ClRed)
	}
	//平色 黄
	if strings.EqualFold(star[1], "轩辕") || strings.EqualFold(star[1], "招摇") {
		f.kanStar.Font().SetColor(colors.ClGoldenrod)
	}
	f.kanStar.SetCaption(star[1])

	f.genStar = vcl.NewLabel(f)
	f.genStar.SetParent(f)
	f.genStar.SetTop(top3)
	f.genStar.SetLeft(left1)
	//吉色 红
	if strings.EqualFold(star[8], "天乙") || strings.EqualFold(star[8], "青龙") ||
		strings.EqualFold(star[8], "太乙") || strings.EqualFold(star[8], "太阴") {
		f.genStar.Font().SetColor(colors.ClRed)
	}
	//平色 黄
	if strings.EqualFold(star[8], "轩辕") || strings.EqualFold(star[8], "招摇") {
		f.genStar.Font().SetColor(colors.ClGoldenrod)
	}
	f.genStar.SetCaption(star[8])

	f.zhenStar = vcl.NewLabel(f)
	f.zhenStar.SetParent(f)
	f.zhenStar.SetTop(top2)
	f.zhenStar.SetLeft(left1)
	//吉色 红
	if strings.EqualFold(star[3], "天乙") || strings.EqualFold(star[3], "青龙") ||
		strings.EqualFold(star[3], "太乙") || strings.EqualFold(star[3], "太阴") {
		f.zhenStar.Font().SetColor(colors.ClRed)
	}
	//平色 黄
	if strings.EqualFold(star[3], "轩辕") || strings.EqualFold(star[3], "招摇") {
		f.zhenStar.Font().SetColor(colors.ClGoldenrod)
	}
	f.zhenStar.SetCaption(star[3])

	f.xunStar = vcl.NewLabel(f)
	f.xunStar.SetParent(f)
	f.xunStar.SetTop(top1)
	f.xunStar.SetLeft(left1)
	//吉色 红
	if strings.EqualFold(star[4], "天乙") || strings.EqualFold(star[4], "青龙") ||
		strings.EqualFold(star[4], "太乙") || strings.EqualFold(star[4], "太阴") {
		f.xunStar.Font().SetColor(colors.ClRed)
	}
	//平色 黄
	if strings.EqualFold(star[4], "轩辕") || strings.EqualFold(star[4], "招摇") {
		f.xunStar.Font().SetColor(colors.ClGoldenrod)
	}
	f.xunStar.SetCaption(star[4])

	f.liStar = vcl.NewLabel(f)
	f.liStar.SetParent(f)
	f.liStar.SetTop(top1)
	f.liStar.SetLeft(left2)
	//吉色 红
	if strings.EqualFold(star[9], "天乙") || strings.EqualFold(star[9], "青龙") ||
		strings.EqualFold(star[9], "太乙") || strings.EqualFold(star[9], "太阴") {
		f.liStar.Font().SetColor(colors.ClRed)
	}
	//平色 黄
	if strings.EqualFold(star[9], "轩辕") || strings.EqualFold(star[9], "招摇") {
		f.liStar.Font().SetColor(colors.ClGoldenrod)
	}
	f.liStar.SetCaption(star[9])

	f.kunStar = vcl.NewLabel(f)
	f.kunStar.SetParent(f)
	f.kunStar.SetTop(top1)
	f.kunStar.SetLeft(left3)
	//吉色 红
	if strings.EqualFold(star[2], "天乙") || strings.EqualFold(star[2], "青龙") ||
		strings.EqualFold(star[2], "太乙") || strings.EqualFold(star[2], "太阴") {
		f.kunStar.Font().SetColor(colors.ClRed)
	}
	//平色 黄
	if strings.EqualFold(star[2], "轩辕") || strings.EqualFold(star[2], "招摇") {
		f.kunStar.Font().SetColor(colors.ClGoldenrod)
	}
	f.kunStar.SetCaption(star[2])

	f.duiStar = vcl.NewLabel(f)
	f.duiStar.SetParent(f)
	f.duiStar.SetTop(top2)
	f.duiStar.SetLeft(left3)
	//吉色 红
	if strings.EqualFold(star[7], "天乙") || strings.EqualFold(star[7], "青龙") ||
		strings.EqualFold(star[7], "太乙") || strings.EqualFold(star[7], "太阴") {
		f.duiStar.Font().SetColor(colors.ClRed)
	}
	//平色 黄
	if strings.EqualFold(star[7], "轩辕") || strings.EqualFold(star[7], "招摇") {
		f.duiStar.Font().SetColor(colors.ClGoldenrod)
	}
	f.duiStar.SetCaption(star[7])

	f.qianStar = vcl.NewLabel(f)
	f.qianStar.SetParent(f)
	f.qianStar.SetTop(top3)
	f.qianStar.SetLeft(left3)
	//吉色 红
	if strings.EqualFold(star[6], "天乙") || strings.EqualFold(star[6], "青龙") ||
		strings.EqualFold(star[6], "太乙") || strings.EqualFold(star[6], "太阴") {
		f.qianStar.Font().SetColor(colors.ClRed)
	}
	//平色 黄
	if strings.EqualFold(star[6], "轩辕") || strings.EqualFold(star[6], "招摇") {
		f.qianStar.Font().SetColor(colors.ClGoldenrod)
	}
	f.qianStar.SetCaption(star[6])

	f.zhongStar = vcl.NewLabel(f)
	f.zhongStar.SetParent(f)
	f.zhongStar.SetTop(top2)
	f.zhongStar.SetLeft(left2)
	//吉色 红
	if strings.EqualFold(star[5], "天乙") || strings.EqualFold(star[5], "青龙") ||
		strings.EqualFold(star[5], "太乙") || strings.EqualFold(star[5], "太阴") {
		f.zhongStar.Font().SetColor(colors.ClRed)
	}
	//平色 黄
	if strings.EqualFold(star[5], "轩辕") || strings.EqualFold(star[5], "招摇") {
		f.zhongStar.Font().SetColor(colors.ClGoldenrod)
	}
	f.zhongStar.SetCaption(star[5])

	//--- 八门 ---
	door := rqm.Door

	f.kanDoor = vcl.NewLabel(f)
	f.kanDoor.SetParent(f)
	f.kanDoor.SetTop(top3 + hTop)
	f.kanDoor.SetLeft(left2)
	//吉色 红
	if strings.EqualFold(door[1], "开") || strings.EqualFold(door[1], "休") || strings.EqualFold(door[1], "生") {
		f.kanDoor.Font().SetColor(colors.ClRed)
	}
	f.kanDoor.SetCaption(door[1])

	f.genDoor = vcl.NewLabel(f)
	f.genDoor.SetParent(f)
	f.genDoor.SetTop(top3 + hTop)
	f.genDoor.SetLeft(left1)
	//吉色 红
	if strings.EqualFold(door[8], "开") || strings.EqualFold(door[8], "休") || strings.EqualFold(door[8], "生") {
		f.genDoor.Font().SetColor(colors.ClRed)
	}
	f.genDoor.SetCaption(door[8])

	f.zhenDoor = vcl.NewLabel(f)
	f.zhenDoor.SetParent(f)
	f.zhenDoor.SetTop(top2 + hTop)
	f.zhenDoor.SetLeft(left1)
	//吉色 红
	if strings.EqualFold(door[3], "开") || strings.EqualFold(door[3], "休") || strings.EqualFold(door[3], "生") {
		f.zhenDoor.Font().SetColor(colors.ClRed)
	}
	f.zhenDoor.SetCaption(door[3])

	f.xunDoor = vcl.NewLabel(f)
	f.xunDoor.SetParent(f)
	f.xunDoor.SetTop(top1 + hTop)
	f.xunDoor.SetLeft(left1)
	//吉色 红
	if strings.EqualFold(door[4], "开") || strings.EqualFold(door[4], "休") || strings.EqualFold(door[4], "生") {
		f.xunDoor.Font().SetColor(colors.ClRed)
	}
	f.xunDoor.SetCaption(door[4])

	f.liDoor = vcl.NewLabel(f)
	f.liDoor.SetParent(f)
	f.liDoor.SetTop(top1 + hTop)
	f.liDoor.SetLeft(left2)
	//吉色 红
	if strings.EqualFold(door[9], "开") || strings.EqualFold(door[9], "休") || strings.EqualFold(door[9], "生") {
		f.liDoor.Font().SetColor(colors.ClRed)
	}
	f.liDoor.SetCaption(door[9])

	f.kunDoor = vcl.NewLabel(f)
	f.kunDoor.SetParent(f)
	f.kunDoor.SetTop(top1 + hTop)
	f.kunDoor.SetLeft(left3)
	//吉色 红
	if strings.EqualFold(door[2], "开") || strings.EqualFold(door[2], "休") || strings.EqualFold(door[2], "生") {
		f.kunDoor.Font().SetColor(colors.ClRed)
	}
	f.kunDoor.SetCaption(door[2])

	f.duiDoor = vcl.NewLabel(f)
	f.duiDoor.SetParent(f)
	f.duiDoor.SetTop(top2 + hTop)
	f.duiDoor.SetLeft(left3)
	//吉色 红
	if strings.EqualFold(door[7], "开") || strings.EqualFold(door[7], "休") || strings.EqualFold(door[7], "生") {
		f.duiDoor.Font().SetColor(colors.ClRed)
	}
	f.duiDoor.SetCaption(door[7])

	f.qianDoor = vcl.NewLabel(f)
	f.qianDoor.SetParent(f)
	f.qianDoor.SetTop(top3 + hTop)
	f.qianDoor.SetLeft(left3)
	//吉色 红
	if strings.EqualFold(door[6], "开") || strings.EqualFold(door[6], "休") || strings.EqualFold(door[6], "生") {
		f.qianDoor.Font().SetColor(colors.ClRed)
	}
	f.qianDoor.SetCaption(door[6])

	//--- 喜神 ---
	xiShen := rqm.XiShen()

	f.genXiShen = vcl.NewLabel(f)
	f.genXiShen.SetParent(f)
	f.genXiShen.SetTop(top3 + 2*hTop)
	f.genXiShen.SetLeft(left1)
	f.genXiShen.Font().SetColor(colors.ClRed)
	if _, ok := xiShen[8]; ok {
		f.genXiShen.SetCaption(xiShen[8] + "(寅时)")
	}
	//f.genXiShen.SetCaption(xiShen[8] + "(寅时)")

	f.xunXiShen = vcl.NewLabel(f)
	f.xunXiShen.SetParent(f)
	f.xunXiShen.SetTop(top1 + 2*hTop)
	f.xunXiShen.SetLeft(left1)
	f.xunXiShen.Font().SetColor(colors.ClRed)
	if _, ok := xiShen[4]; ok {
		f.xunXiShen.SetCaption(xiShen[4] + "(辰时)")
	}

	f.liXiShen = vcl.NewLabel(f)
	f.liXiShen.SetParent(f)
	f.liXiShen.SetTop(top1 + 2*hTop)
	f.liXiShen.SetLeft(left2)
	f.liXiShen.Font().SetColor(colors.ClRed)
	if _, ok := xiShen[9]; ok {
		f.liXiShen.SetCaption(xiShen[9] + "(午时)")
	}

	f.kunXiShen = vcl.NewLabel(f)
	f.kunXiShen.SetParent(f)
	f.kunXiShen.SetTop(top1 + 2*hTop)
	f.kunXiShen.SetLeft(left3)
	f.kunXiShen.Font().SetColor(colors.ClRed)
	if _, ok := xiShen[2]; ok {
		f.kunXiShen.SetCaption(xiShen[2] + "(申时)")
	}

	f.qianXiShen = vcl.NewLabel(f)
	f.qianXiShen.SetParent(f)
	f.qianXiShen.SetTop(top3 + 2*hTop)
	f.qianXiShen.SetLeft(left3)
	f.qianXiShen.Font().SetColor(colors.ClRed)
	if _, ok := xiShen[6]; ok {
		f.qianXiShen.SetCaption(xiShen[6] + "(戌时)")
	}
	//--- 黄黑 ---
	huangHei := rqm.HuangHeiH()
	harr := rqm.HgzArr

	f.huangHeiZi = vcl.NewLabel(f)
	f.huangHeiZi.SetParent(f)
	f.huangHeiZi.SetTop(top1)
	f.huangHeiZi.SetLeft(left4)
	if strings.EqualFold(huangHei["子"], "司命") || strings.EqualFold(huangHei["子"], "青龙") ||
		strings.EqualFold(huangHei["子"], "金匮") || strings.EqualFold(huangHei["子"], "明堂") ||
		strings.EqualFold(huangHei["子"], "天德") || strings.EqualFold(huangHei["子"], "玉堂") {
		f.huangHeiZi.Font().SetColor(colors.ClRed)
	}
	f.huangHeiZi.SetCaption(harr[0] + huangHei["子"])

	f.huangHeiChou = vcl.NewLabel(f)
	f.huangHeiChou.SetParent(f)
	f.huangHeiChou.SetTop(top1 * 2)
	f.huangHeiChou.SetLeft(left4)
	if strings.EqualFold(huangHei["丑"], "司命") || strings.EqualFold(huangHei["丑"], "青龙") ||
		strings.EqualFold(huangHei["丑"], "金匮") || strings.EqualFold(huangHei["丑"], "明堂") ||
		strings.EqualFold(huangHei["丑"], "天德") || strings.EqualFold(huangHei["丑"], "玉堂") {
		f.huangHeiChou.Font().SetColor(colors.ClRed)
	}
	f.huangHeiChou.SetCaption(harr[1] + huangHei["丑"])

	f.huangHeiYin = vcl.NewLabel(f)
	f.huangHeiYin.SetParent(f)
	f.huangHeiYin.SetTop(top1 * 3)
	f.huangHeiYin.SetLeft(left4)
	if strings.EqualFold(huangHei["寅"], "司命") || strings.EqualFold(huangHei["寅"], "青龙") ||
		strings.EqualFold(huangHei["寅"], "金匮") || strings.EqualFold(huangHei["寅"], "明堂") ||
		strings.EqualFold(huangHei["寅"], "天德") || strings.EqualFold(huangHei["寅"], "玉堂") {
		f.huangHeiYin.Font().SetColor(colors.ClRed)
	}
	f.huangHeiYin.SetCaption(harr[2] + huangHei["寅"])

	f.huangHeiMao = vcl.NewLabel(f)
	f.huangHeiMao.SetParent(f)
	f.huangHeiMao.SetTop(top1 * 4)
	f.huangHeiMao.SetLeft(left4)
	if strings.EqualFold(huangHei["卯"], "司命") || strings.EqualFold(huangHei["卯"], "青龙") ||
		strings.EqualFold(huangHei["卯"], "金匮") || strings.EqualFold(huangHei["卯"], "明堂") ||
		strings.EqualFold(huangHei["卯"], "天德") || strings.EqualFold(huangHei["卯"], "玉堂") {
		f.huangHeiMao.Font().SetColor(colors.ClRed)
	}
	f.huangHeiMao.SetCaption(harr[3] + huangHei["卯"])

	f.huangHeiChen = vcl.NewLabel(f)
	f.huangHeiChen.SetParent(f)
	f.huangHeiChen.SetTop(top1 * 5)
	f.huangHeiChen.SetLeft(left4)
	if strings.EqualFold(huangHei["辰"], "司命") || strings.EqualFold(huangHei["辰"], "青龙") ||
		strings.EqualFold(huangHei["辰"], "金匮") || strings.EqualFold(huangHei["辰"], "明堂") ||
		strings.EqualFold(huangHei["辰"], "天德") || strings.EqualFold(huangHei["辰"], "玉堂") {
		f.huangHeiChen.Font().SetColor(colors.ClRed)
	}
	f.huangHeiChen.SetCaption(harr[4] + huangHei["辰"])

	f.huangHeiSi = vcl.NewLabel(f)
	f.huangHeiSi.SetParent(f)
	f.huangHeiSi.SetTop(top1 * 6)
	f.huangHeiSi.SetLeft(left4)
	if strings.EqualFold(huangHei["巳"], "司命") || strings.EqualFold(huangHei["巳"], "青龙") ||
		strings.EqualFold(huangHei["巳"], "金匮") || strings.EqualFold(huangHei["巳"], "明堂") ||
		strings.EqualFold(huangHei["巳"], "天德") || strings.EqualFold(huangHei["巳"], "玉堂") {
		f.huangHeiSi.Font().SetColor(colors.ClRed)
	}
	f.huangHeiSi.SetCaption(harr[5] + huangHei["巳"])

	f.huangHeiWu = vcl.NewLabel(f)
	f.huangHeiWu.SetParent(f)
	f.huangHeiWu.SetTop(top1 * 7)
	f.huangHeiWu.SetLeft(left4)
	if strings.EqualFold(huangHei["午"], "司命") || strings.EqualFold(huangHei["午"], "青龙") ||
		strings.EqualFold(huangHei["午"], "金匮") || strings.EqualFold(huangHei["午"], "明堂") ||
		strings.EqualFold(huangHei["午"], "天德") || strings.EqualFold(huangHei["午"], "玉堂") {
		f.huangHeiWu.Font().SetColor(colors.ClRed)
	}
	f.huangHeiWu.SetCaption(harr[6] + huangHei["午"])

	f.huangHeiWei = vcl.NewLabel(f)
	f.huangHeiWei.SetParent(f)
	f.huangHeiWei.SetTop(top1 * 8)
	f.huangHeiWei.SetLeft(left4)
	if strings.EqualFold(huangHei["未"], "司命") || strings.EqualFold(huangHei["未"], "青龙") ||
		strings.EqualFold(huangHei["未"], "金匮") || strings.EqualFold(huangHei["未"], "明堂") ||
		strings.EqualFold(huangHei["未"], "天德") || strings.EqualFold(huangHei["未"], "玉堂") {
		f.huangHeiWei.Font().SetColor(colors.ClRed)
	}
	f.huangHeiWei.SetCaption(harr[7] + huangHei["未"])

	f.huangHeiShen = vcl.NewLabel(f)
	f.huangHeiShen.SetParent(f)
	f.huangHeiShen.SetTop(top1 * 9)
	f.huangHeiShen.SetLeft(left4)
	if strings.EqualFold(huangHei["申"], "司命") || strings.EqualFold(huangHei["申"], "青龙") ||
		strings.EqualFold(huangHei["申"], "金匮") || strings.EqualFold(huangHei["申"], "明堂") ||
		strings.EqualFold(huangHei["申"], "天德") || strings.EqualFold(huangHei["申"], "玉堂") {
		f.huangHeiShen.Font().SetColor(colors.ClRed)
	}
	f.huangHeiShen.SetCaption(harr[8] + huangHei["申"])

	f.huangHeiYou = vcl.NewLabel(f)
	f.huangHeiYou.SetParent(f)
	f.huangHeiYou.SetTop(top1 * 10)
	f.huangHeiYou.SetLeft(left4)
	if strings.EqualFold(huangHei["酉"], "司命") || strings.EqualFold(huangHei["酉"], "青龙") ||
		strings.EqualFold(huangHei["酉"], "金匮") || strings.EqualFold(huangHei["酉"], "明堂") ||
		strings.EqualFold(huangHei["酉"], "天德") || strings.EqualFold(huangHei["酉"], "玉堂") {
		f.huangHeiYou.Font().SetColor(colors.ClRed)
	}
	f.huangHeiYou.SetCaption(harr[9] + huangHei["酉"])

	f.huangHeiXu = vcl.NewLabel(f)
	f.huangHeiXu.SetParent(f)
	f.huangHeiXu.SetTop(top1 * 11)
	f.huangHeiXu.SetLeft(left4)
	if strings.EqualFold(huangHei["戌"], "司命") || strings.EqualFold(huangHei["戌"], "青龙") ||
		strings.EqualFold(huangHei["戌"], "金匮") || strings.EqualFold(huangHei["戌"], "明堂") ||
		strings.EqualFold(huangHei["戌"], "天德") || strings.EqualFold(huangHei["戌"], "玉堂") {
		f.huangHeiXu.Font().SetColor(colors.ClRed)
	}
	f.huangHeiXu.SetCaption(harr[10] + huangHei["戌"])

	f.huangHeiHai = vcl.NewLabel(f)
	f.huangHeiHai.SetParent(f)
	f.huangHeiHai.SetTop(top1 * 12)
	f.huangHeiHai.SetLeft(left4)
	if strings.EqualFold(huangHei["亥"], "司命") || strings.EqualFold(huangHei["亥"], "青龙") ||
		strings.EqualFold(huangHei["亥"], "金匮") || strings.EqualFold(huangHei["亥"], "明堂") ||
		strings.EqualFold(huangHei["亥"], "天德") || strings.EqualFold(huangHei["亥"], "玉堂") {
		f.huangHeiHai.Font().SetColor(colors.ClRed)
	}
	f.huangHeiHai.SetCaption(harr[11] + huangHei["亥"])

	//--- 截路空 ---
	f.jieLuZi = vcl.NewLabel(f)
	f.jieLuZi.SetParent(f)
	f.jieLuZi.SetTop(top1)
	f.jieLuZi.SetLeft(left5)
	if strings.ContainsAny(harr[0], "壬") || strings.ContainsAny(harr[0], "癸") {
		f.jieLuZi.SetCaption("截路空")
	}

	f.jieLuChou = vcl.NewLabel(f)
	f.jieLuChou.SetParent(f)
	f.jieLuChou.SetTop(top1 * 2)
	f.jieLuChou.SetLeft(left5)
	if strings.ContainsAny(harr[1], "壬") || strings.ContainsAny(harr[1], "癸") {
		f.jieLuChou.SetCaption("截路空")
	}

	f.jieLuYin = vcl.NewLabel(f)
	f.jieLuYin.SetParent(f)
	f.jieLuYin.SetTop(top1 * 3)
	f.jieLuYin.SetLeft(left5)
	if strings.ContainsAny(harr[2], "壬") || strings.ContainsAny(harr[2], "癸") {
		f.jieLuYin.SetCaption("截路空")
	}

	f.jieLuMao = vcl.NewLabel(f)
	f.jieLuMao.SetParent(f)
	f.jieLuMao.SetTop(top1 * 4)
	f.jieLuMao.SetLeft(left5)
	if strings.ContainsAny(harr[3], "壬") || strings.ContainsAny(harr[3], "癸") {
		f.jieLuMao.SetCaption("截路空")
	}

	f.jieLuChen = vcl.NewLabel(f)
	f.jieLuChen.SetParent(f)
	f.jieLuChen.SetTop(top1 * 5)
	f.jieLuChen.SetLeft(left5)
	if strings.ContainsAny(harr[4], "壬") || strings.ContainsAny(harr[4], "癸") {
		f.jieLuChen.SetCaption("截路空")
	}

	f.jieLuSi = vcl.NewLabel(f)
	f.jieLuSi.SetParent(f)
	f.jieLuSi.SetTop(top1 * 6)
	f.jieLuSi.SetLeft(left5)
	if strings.ContainsAny(harr[5], "壬") || strings.ContainsAny(harr[5], "癸") {
		f.jieLuSi.SetCaption("截路空")
	}

	f.jieLuWu = vcl.NewLabel(f)
	f.jieLuWu.SetParent(f)
	f.jieLuWu.SetTop(top1 * 7)
	f.jieLuWu.SetLeft(left5)
	if strings.ContainsAny(harr[6], "壬") || strings.ContainsAny(harr[6], "癸") {
		f.jieLuWu.SetCaption("截路空")
	}

	f.jieLuWei = vcl.NewLabel(f)
	f.jieLuWei.SetParent(f)
	f.jieLuWei.SetTop(top1 * 8)
	f.jieLuWei.SetLeft(left5)
	if strings.ContainsAny(harr[7], "壬") || strings.ContainsAny(harr[7], "癸") {
		f.jieLuWei.SetCaption("截路空")
	}

	f.jieLuShen = vcl.NewLabel(f)
	f.jieLuShen.SetParent(f)
	f.jieLuShen.SetTop(top1 * 9)
	f.jieLuShen.SetLeft(left5)
	if strings.ContainsAny(harr[8], "壬") || strings.ContainsAny(harr[8], "癸") {
		f.jieLuShen.SetCaption("截路空")
	}

	f.jieLuYou = vcl.NewLabel(f)
	f.jieLuYou.SetParent(f)
	f.jieLuYou.SetTop(top1 * 10)
	f.jieLuYou.SetLeft(left5)
	if strings.ContainsAny(harr[9], "壬") || strings.ContainsAny(harr[9], "癸") {
		f.jieLuYou.SetCaption("截路空")
	}

	f.jieLuXu = vcl.NewLabel(f)
	f.jieLuXu.SetParent(f)
	f.jieLuXu.SetTop(top1 * 11)
	f.jieLuXu.SetLeft(left5)
	if strings.ContainsAny(harr[10], "壬") || strings.ContainsAny(harr[10], "癸") {
		f.jieLuXu.SetCaption("截路空")
	}

	f.jieLuHai = vcl.NewLabel(f)
	f.jieLuHai.SetParent(f)
	f.jieLuHai.SetTop(top1 * 12)
	f.jieLuHai.SetLeft(left5)
	if strings.ContainsAny(harr[11], "壬") || strings.ContainsAny(harr[11], "癸") {
		f.jieLuHai.SetCaption("截路空")
	}
	//--- 天乙贵人 ---
	tianYi := rqm.TianYi()
	f.tianYiZi = vcl.NewLabel(f)
	f.tianYiZi.SetParent(f)
	f.tianYiZi.SetTop(top1)
	f.tianYiZi.SetLeft(left6)
	f.tianYiZi.Font().SetColor(colors.ClRed)
	if _, ok := tianYi[harr[0]]; ok {
		f.tianYiZi.SetCaption(tianYi[harr[0]])
	}

	f.tianYiChou = vcl.NewLabel(f)
	f.tianYiChou.SetParent(f)
	f.tianYiChou.SetTop(top1 * 2)
	f.tianYiChou.SetLeft(left6)
	f.tianYiChou.Font().SetColor(colors.ClRed)
	if _, ok := tianYi[harr[1]]; ok {
		f.tianYiChou.SetCaption(tianYi[harr[1]])
	}

	f.tianYiYin = vcl.NewLabel(f)
	f.tianYiYin.SetParent(f)
	f.tianYiYin.SetTop(top1 * 3)
	f.tianYiYin.SetLeft(left6)
	f.tianYiYin.Font().SetColor(colors.ClRed)
	if _, ok := tianYi[harr[2]]; ok {
		f.tianYiYin.SetCaption(tianYi[harr[2]])
	}

	f.tianYiMao = vcl.NewLabel(f)
	f.tianYiMao.SetParent(f)
	f.tianYiMao.SetTop(top1 * 4)
	f.tianYiMao.SetLeft(left6)
	f.tianYiMao.Font().SetColor(colors.ClRed)
	if _, ok := tianYi[harr[3]]; ok {
		f.tianYiMao.SetCaption(tianYi[harr[3]])
	}

	f.tianYiChen = vcl.NewLabel(f)
	f.tianYiChen.SetParent(f)
	f.tianYiChen.SetTop(top1 * 5)
	f.tianYiChen.SetLeft(left6)
	f.tianYiChen.Font().SetColor(colors.ClRed)
	if _, ok := tianYi[harr[4]]; ok {
		f.tianYiChen.SetCaption(tianYi[harr[4]])
	}

	f.tianYiSi = vcl.NewLabel(f)
	f.tianYiSi.SetParent(f)
	f.tianYiSi.SetTop(top1 * 6)
	f.tianYiSi.SetLeft(left6)
	f.tianYiSi.Font().SetColor(colors.ClRed)
	if _, ok := tianYi[harr[5]]; ok {
		f.tianYiSi.SetCaption(tianYi[harr[5]])
	}

	f.tianYiWu = vcl.NewLabel(f)
	f.tianYiWu.SetParent(f)
	f.tianYiWu.SetTop(top1 * 7)
	f.tianYiWu.SetLeft(left6)
	f.tianYiWu.Font().SetColor(colors.ClRed)
	if _, ok := tianYi[harr[6]]; ok {
		f.tianYiWu.SetCaption(tianYi[harr[6]])
	}

	f.tianYiWei = vcl.NewLabel(f)
	f.tianYiWei.SetParent(f)
	f.tianYiWei.SetTop(top1 * 8)
	f.tianYiWei.SetLeft(left6)
	f.tianYiWei.Font().SetColor(colors.ClRed)
	if _, ok := tianYi[harr[7]]; ok {
		f.tianYiWei.SetCaption(tianYi[harr[7]])
	}

	f.tianYiShen = vcl.NewLabel(f)
	f.tianYiShen.SetParent(f)
	f.tianYiShen.SetTop(top1 * 9)
	f.tianYiShen.SetLeft(left6)
	f.tianYiShen.Font().SetColor(colors.ClRed)
	if _, ok := tianYi[harr[8]]; ok {
		f.tianYiShen.SetCaption(tianYi[harr[8]])
	}

	f.tianYiYou = vcl.NewLabel(f)
	f.tianYiYou.SetParent(f)
	f.tianYiYou.SetTop(top1 * 10)
	f.tianYiYou.SetLeft(left6)
	f.tianYiYou.Font().SetColor(colors.ClRed)
	if _, ok := tianYi[harr[9]]; ok {
		f.tianYiYou.SetCaption(tianYi[harr[9]])
	}

	f.tianYiXu = vcl.NewLabel(f)
	f.tianYiXu.SetParent(f)
	f.tianYiXu.SetTop(top1 * 11)
	f.tianYiXu.SetLeft(left6)
	f.tianYiXu.Font().SetColor(colors.ClRed)
	if _, ok := tianYi[harr[10]]; ok {
		f.tianYiXu.SetCaption(tianYi[harr[10]])
	}

	f.tianYiHai = vcl.NewLabel(f)
	f.tianYiHai.SetParent(f)
	f.tianYiHai.SetTop(top1 * 12)
	f.tianYiHai.SetLeft(left6)
	f.tianYiHai.Font().SetColor(colors.ClRed)
	if _, ok := tianYi[harr[11]]; ok {
		f.tianYiHai.SetCaption(tianYi[harr[11]])
	}
	// --- 五不遇 ---
	wuBuYu := rqm.WuBuYu()

	f.wuBuYuZi = vcl.NewLabel(f)
	f.wuBuYuZi.SetParent(f)
	f.wuBuYuZi.SetTop(top1)
	f.wuBuYuZi.SetLeft(left7)
	if _, ok := wuBuYu[harr[0]]; ok {
		f.wuBuYuZi.SetCaption(wuBuYu[harr[0]])
	}

	f.wuBuYuChou = vcl.NewLabel(f)
	f.wuBuYuChou.SetParent(f)
	f.wuBuYuChou.SetTop(top1 * 2)
	f.wuBuYuChou.SetLeft(left7)
	if _, ok := wuBuYu[harr[1]]; ok {
		f.wuBuYuChou.SetCaption(wuBuYu[harr[1]])
	}

	f.wuBuYuYin = vcl.NewLabel(f)
	f.wuBuYuYin.SetParent(f)
	f.wuBuYuYin.SetTop(top1 * 3)
	f.wuBuYuYin.SetLeft(left7)
	if _, ok := wuBuYu[harr[2]]; ok {
		f.wuBuYuYin.SetCaption(wuBuYu[harr[2]])
	}

	f.wuBuYuMao = vcl.NewLabel(f)
	f.wuBuYuMao.SetParent(f)
	f.wuBuYuMao.SetTop(top1 * 4)
	f.wuBuYuMao.SetLeft(left7)
	if _, ok := wuBuYu[harr[3]]; ok {
		f.wuBuYuMao.SetCaption(wuBuYu[harr[3]])
	}

	f.wuBuYuChen = vcl.NewLabel(f)
	f.wuBuYuChen.SetParent(f)
	f.wuBuYuChen.SetTop(top1 * 5)
	f.wuBuYuChen.SetLeft(left7)
	if _, ok := wuBuYu[harr[4]]; ok {
		f.wuBuYuChen.SetCaption(wuBuYu[harr[4]])
	}

	f.wuBuYuSi = vcl.NewLabel(f)
	f.wuBuYuSi.SetParent(f)
	f.wuBuYuSi.SetTop(top1 * 6)
	f.wuBuYuSi.SetLeft(left7)
	if _, ok := wuBuYu[harr[5]]; ok {
		f.wuBuYuSi.SetCaption(wuBuYu[harr[5]])
	}

	f.wuBuYuWu = vcl.NewLabel(f)
	f.wuBuYuWu.SetParent(f)
	f.wuBuYuWu.SetTop(top1 * 7)
	f.wuBuYuWu.SetLeft(left7)
	if _, ok := wuBuYu[harr[6]]; ok {
		f.wuBuYuWu.SetCaption(wuBuYu[harr[6]])
	}

	f.wuBuYuWei = vcl.NewLabel(f)
	f.wuBuYuWei.SetParent(f)
	f.wuBuYuWei.SetTop(top1 * 8)
	f.wuBuYuWei.SetLeft(left7)
	if _, ok := wuBuYu[harr[7]]; ok {
		f.wuBuYuWei.SetCaption(wuBuYu[harr[7]])
	}

	f.wuBuYuShen = vcl.NewLabel(f)
	f.wuBuYuShen.SetParent(f)
	f.wuBuYuShen.SetTop(top1 * 9)
	f.wuBuYuShen.SetLeft(left7)
	if _, ok := wuBuYu[harr[8]]; ok {
		f.wuBuYuShen.SetCaption(wuBuYu[harr[8]])
	}

	f.wuBuYuYou = vcl.NewLabel(f)
	f.wuBuYuYou.SetParent(f)
	f.wuBuYuYou.SetTop(top1 * 10)
	f.wuBuYuYou.SetLeft(left7)
	if _, ok := wuBuYu[harr[9]]; ok {
		f.wuBuYuYou.SetCaption(wuBuYu[harr[9]])
	}

	f.wuBuYuXu = vcl.NewLabel(f)
	f.wuBuYuXu.SetParent(f)
	f.wuBuYuXu.SetTop(top1 * 11)
	f.wuBuYuXu.SetLeft(left7)
	if _, ok := wuBuYu[harr[10]]; ok {
		f.wuBuYuXu.SetCaption(wuBuYu[harr[10]])
	}

	f.wuBuYuHai = vcl.NewLabel(f)
	f.wuBuYuHai.SetParent(f)
	f.wuBuYuHai.SetTop(top1 * 12)
	f.wuBuYuHai.SetLeft(left7)
	if _, ok := wuBuYu[harr[11]]; ok {
		f.wuBuYuHai.SetCaption(wuBuYu[harr[11]])
	}

	//---九星 八门 吉凶 ---
	infoStarFei := rqm.InfoJiuXingFeiGong()
	infoStar := rqm.InfoJiuXing()
	infoDoor := rqm.InfoBaMen()
	infoDS1 := rqm.InfoBaMenJiuXing(1)
	infoDS8 := rqm.InfoBaMenJiuXing(8)
	infoDS3 := rqm.InfoBaMenJiuXing(3)
	infoDS4 := rqm.InfoBaMenJiuXing(4)
	infoDS9 := rqm.InfoBaMenJiuXing(9)
	infoDS2 := rqm.InfoBaMenJiuXing(2)
	infoDS7 := rqm.InfoBaMenJiuXing(7)
	infoDS6 := rqm.InfoBaMenJiuXing(6)

	f.kanStar.SetHint(infoStarFei[1] + infoStar[1])
	f.genStar.SetHint(infoStarFei[8] + infoStar[8])
	f.zhenStar.SetHint(infoStarFei[3] + infoStar[3])
	f.xunStar.SetHint(infoStarFei[4] + infoStar[4])
	f.liStar.SetHint(infoStarFei[9] + infoStar[9])
	f.kunStar.SetHint(infoStarFei[2] + infoStar[2])
	f.duiStar.SetHint(infoStarFei[7] + infoStar[7])
	f.qianStar.SetHint(infoStarFei[6] + infoStar[6])
	f.zhongStar.SetHint(infoStarFei[5] + infoStar[5])

	if infoDS1 != "" {
		f.kanDoor.SetHint(infoDoor[1] + infoDS1)
	} else {
		f.kanDoor.SetHint(infoDoor[1])
	}
	if infoDS8 != "" {
		f.genDoor.SetHint(infoDoor[8] + infoDS8)
	} else {
		f.genDoor.SetHint(infoDoor[8])
	}
	if infoDS3 != "" {
		f.zhenDoor.SetHint(infoDoor[3] + infoDS3)
	} else {
		f.zhenDoor.SetHint(infoDoor[3])
	}
	if infoDS4 != "" {
		f.xunDoor.SetHint(infoDoor[4] + infoDS4)
	} else {
		f.xunDoor.SetHint(infoDoor[4])
	}
	if infoDS9 != "" {
		f.liDoor.SetHint(infoDoor[9] + infoDS9)
	} else {
		f.liDoor.SetHint(infoDoor[9])
	}
	if infoDS2 != "" {
		f.kunDoor.SetHint(infoDoor[2] + infoDS2)
	} else {
		f.kunDoor.SetHint(infoDoor[2])
	}
	if infoDS7 != "" {
		f.duiDoor.SetHint(infoDoor[7] + infoDS7)
	} else {
		f.duiDoor.SetHint(infoDoor[7])
	}
	if infoDS6 != "" {
		f.qianDoor.SetHint(infoDoor[6] + infoDS6)
	} else {
		f.qianDoor.SetHint(infoDoor[6])
	}

	f.huangHeiZi.SetHint(rqm.InfoHuangHei(huangHei["子"]))
	f.huangHeiChou.SetHint(rqm.InfoHuangHei(huangHei["丑"]))
	f.huangHeiYin.SetHint(rqm.InfoHuangHei(huangHei["寅"]))
	f.huangHeiMao.SetHint(rqm.InfoHuangHei(huangHei["卯"]))
	f.huangHeiChen.SetHint(rqm.InfoHuangHei(huangHei["辰"]))
	f.huangHeiSi.SetHint(rqm.InfoHuangHei(huangHei["巳"]))
	f.huangHeiWu.SetHint(rqm.InfoHuangHei(huangHei["午"]))
	f.huangHeiWei.SetHint(rqm.InfoHuangHei(huangHei["未"]))
	f.huangHeiShen.SetHint(rqm.InfoHuangHei(huangHei["申"]))
	f.huangHeiYou.SetHint(rqm.InfoHuangHei(huangHei["酉"]))
	f.huangHeiXu.SetHint(rqm.InfoHuangHei(huangHei["戌"]))
	f.huangHeiHai.SetHint(rqm.InfoHuangHei(huangHei["亥"]))
}
