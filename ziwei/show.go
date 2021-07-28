package main

import (
	"fmt"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types/colors"
)

//中心区域　基本信息
func (f *TMainForm) OnFormInfoShow(label1s, label2s, label3s string) {
	f.label1.Free()
	f.label1 = vcl.NewLabel(f)
	f.label1.SetParent(f)
	f.label1.SetTop(b1 + 60)
	f.label1.SetLeft(r1)
	f.label1.SetWidth(r1)
	f.label1.SetTextBuf(label1s)

	f.label2.Free()
	f.label2 = vcl.NewLabel(f)
	f.label2.SetParent(f)
	f.label2.SetTop(b1)
	f.label2.SetLeft(2 * r1)
	f.label2.SetWidth(r1)
	f.label2.SetTextBuf(label2s)

	f.label3.Free()
	f.label3 = vcl.NewLabel(f)
	f.label3.SetParent(f)
	f.label3.SetTop(2 * b1)
	f.label3.SetLeft(r1)
	f.label3.SetWidth(r1)
	f.label3.SetTextBuf(label3s)
	f.label3.Font().SetColor(colors.ClRed)
}

//寅首(月干支)　十二长生　btn
func (f *TMainForm) OnFormMgzShow(mgz, ch []string) {
	f.btnHai.Free()
	f.btnHai = vcl.NewButton(f)
	f.btnHai.SetParent(f)
	f.btnHai.SetTop(t4)
	f.btnHai.SetLeft(l5)
	f.btnHai.SetTextBuf(mgz[0] + " " + ch[0])
	f.btnHai.SetOnClick(f.btnHaiOnClick)

	f.btnZi.Free()
	f.btnZi = vcl.NewButton(f)
	f.btnZi.SetParent(f)
	f.btnZi.SetTop(t4)
	f.btnZi.SetLeft(r4)
	f.btnZi.SetTextBuf(mgz[1] + " " + ch[1])
	f.btnZi.SetOnClick(f.btnZiOnClick)

	f.btnChou.Free()
	f.btnChou = vcl.NewButton(f)
	f.btnChou.SetParent(f)
	f.btnChou.SetTop(t4)
	f.btnChou.SetLeft(l4)
	f.btnChou.SetTextBuf(mgz[2] + " " + ch[2])
	f.btnChou.SetOnClick(f.btnChouOnClick)

	f.btnYin.Free()
	f.btnYin = vcl.NewButton(f)
	f.btnYin.SetParent(f)
	f.btnYin.SetTop(t4)
	f.btnYin.SetLeft(l1)
	f.btnYin.SetTextBuf(mgz[3] + " " + ch[3])
	f.btnYin.SetOnClick(f.btnYinOnClick)

	f.btnMao.Free()
	f.btnMao = vcl.NewButton(f)
	f.btnMao.SetParent(f)
	f.btnMao.SetTop(t3)
	f.btnMao.SetLeft(l1)
	f.btnMao.SetTextBuf(mgz[4] + " " + ch[4])
	f.btnMao.SetOnClick(f.btnMaoOnClick)

	f.btnChen.Free()
	f.btnChen = vcl.NewButton(f)
	f.btnChen.SetParent(f)
	f.btnChen.SetTop(t6)
	f.btnChen.SetLeft(l1)
	f.btnChen.SetTextBuf(mgz[5] + " " + ch[5])
	f.btnChen.SetOnClick(f.btnChenOnClick)

	f.btnSi.Free()
	f.btnSi = vcl.NewButton(f)
	f.btnSi.SetParent(f)
	f.btnSi.SetTop(top)
	f.btnSi.SetLeft(l1)
	f.btnSi.SetTextBuf(mgz[6] + " " + ch[6])
	f.btnSi.SetOnClick(f.btnSiOnClick)

	f.btnWu.Free()
	f.btnWu = vcl.NewButton(f)
	f.btnWu.SetParent(f)
	f.btnWu.SetTop(top)
	f.btnWu.SetLeft(r1)
	f.btnWu.SetTextBuf(mgz[7] + " " + ch[7])
	f.btnWu.SetOnClick(f.btnWuOnClick)

	f.btnWei.Free()
	f.btnWei = vcl.NewButton(f)
	f.btnWei.SetParent(f)
	f.btnWei.SetTop(top)
	f.btnWei.SetLeft(l2)
	f.btnWei.SetTextBuf(mgz[8] + " " + ch[8])
	f.btnWei.SetOnClick(f.btnWeiOnClick)

	f.btnShen.Free()
	f.btnShen = vcl.NewButton(f)
	f.btnShen.SetParent(f)
	f.btnShen.SetTop(top)
	f.btnShen.SetLeft(l5)
	f.btnShen.SetTextBuf(mgz[9] + " " + ch[9])
	f.btnShen.SetOnClick(f.btnShenOnClick)

	f.btnYou.Free()
	f.btnYou = vcl.NewButton(f)
	f.btnYou.SetParent(f)
	f.btnYou.SetTop(t6)
	f.btnYou.SetLeft(l5)
	f.btnYou.SetTextBuf(mgz[10] + " " + ch[10])
	f.btnYou.SetOnClick(f.btnYouOnClick)

	f.btnXu.Free()
	f.btnXu = vcl.NewButton(f)
	f.btnXu.SetParent(f)
	f.btnXu.SetTop(t3)
	f.btnXu.SetLeft(l5)
	f.btnXu.SetTextBuf(mgz[11] + " " + ch[11])
	f.btnXu.SetOnClick(f.btnXuOnClick)
}

//四化
func (f *TMainForm) OnFormSiHuaShow(shmap map[string]string) {
	const height1 = 5 * 30
	const step = 60
	f.haiSiHua.Free()
	f.haiSiHua = vcl.NewLabel(f)
	f.haiSiHua.SetParent(f)
	f.haiSiHua.SetTop(t4 + height1)
	f.haiSiHua.SetLeft(l5 + step)
	f.haiSiHua.Font().SetColor(colors.ClDarkorange)

	f.ziSiHua.Free()
	f.ziSiHua = vcl.NewLabel(f)
	f.ziSiHua.SetParent(f)
	f.ziSiHua.SetTop(t4 + height1)
	f.ziSiHua.SetLeft(r4 + step)
	f.ziSiHua.Font().SetColor(colors.ClDarkorange)

	f.chouSiHua.Free()
	f.chouSiHua = vcl.NewLabel(f)
	f.chouSiHua.SetParent(f)
	f.chouSiHua.SetTop(t4 + height1)
	f.chouSiHua.SetLeft(l4 + step)
	f.chouSiHua.Font().SetColor(colors.ClDarkorange)

	f.yinSiHua.Free()
	f.yinSiHua = vcl.NewLabel(f)
	f.yinSiHua.SetParent(f)
	f.yinSiHua.SetTop(t4 + height1)
	f.yinSiHua.SetLeft(l1 + step)
	f.yinSiHua.Font().SetColor(colors.ClDarkorange)

	f.maoSiHua.Free()
	f.maoSiHua = vcl.NewLabel(f)
	f.maoSiHua.SetParent(f)
	f.maoSiHua.SetTop(t3 + height1)
	f.maoSiHua.SetLeft(l1 + step)
	f.maoSiHua.Font().SetColor(colors.ClDarkorange)

	f.chenSiHua.Free()
	f.chenSiHua = vcl.NewLabel(f)
	f.chenSiHua.SetParent(f)
	f.chenSiHua.SetTop(t6 + height1)
	f.chenSiHua.SetLeft(l1 + step)
	f.chenSiHua.Font().SetColor(colors.ClDarkorange)

	f.siSiHua.Free()
	f.siSiHua = vcl.NewLabel(f)
	f.siSiHua.SetParent(f)
	f.siSiHua.SetTop(top + height1)
	f.siSiHua.SetLeft(l1 + step)
	f.siSiHua.Font().SetColor(colors.ClDarkorange)

	f.wuSiHua.Free()
	f.wuSiHua = vcl.NewLabel(f)
	f.wuSiHua.SetParent(f)
	f.wuSiHua.SetTop(top + height1)
	f.wuSiHua.SetLeft(r1 + step)
	f.wuSiHua.Font().SetColor(colors.ClDarkorange)

	f.weiSiHua.Free()
	f.weiSiHua = vcl.NewLabel(f)
	f.weiSiHua.SetParent(f)
	f.weiSiHua.SetTop(top + height1)
	f.weiSiHua.SetLeft(l2 + step)
	f.weiSiHua.Font().SetColor(colors.ClDarkorange)

	f.shenSiHua.Free()
	f.shenSiHua = vcl.NewLabel(f)
	f.shenSiHua.SetParent(f)
	f.shenSiHua.SetTop(top + height1)
	f.shenSiHua.SetLeft(l5 + step)
	f.shenSiHua.Font().SetColor(colors.ClDarkorange)

	f.youSiHua.Free()
	f.youSiHua = vcl.NewLabel(f)
	f.youSiHua.SetParent(f)
	f.youSiHua.SetTop(t6 + height1)
	f.youSiHua.SetLeft(l5 + step)
	f.youSiHua.Font().SetColor(colors.ClDarkorange)

	f.xuSiHua.Free()
	f.xuSiHua = vcl.NewLabel(f)
	f.xuSiHua.SetParent(f)
	f.xuSiHua.SetTop(t3 + height1)
	f.xuSiHua.SetLeft(l5 + step)
	f.xuSiHua.Font().SetColor(colors.ClDarkorange)

	var arr, arr1, arr2, arr3, arr4, arr5, arr6, arr7, arr8, arr9, arr10, arr11 []string
	for siHuaS, zhi := range shmap {
		switch zhi {
		case "亥":
			arr = append(arr, siHuaS)
			siHuaS1 := arrToS(arr, " ")
			f.haiSiHua.SetTextBuf(siHuaS1)
		case "子":
			arr1 = append(arr1, siHuaS)
			siHuaS2 := arrToS(arr1, " ")
			f.ziSiHua.SetTextBuf(siHuaS2)
		case "丑":
			arr2 = append(arr2, siHuaS)
			siHuaS2 := arrToS(arr2, " ")
			f.chouSiHua.SetTextBuf(siHuaS2)
		case "寅":
			arr3 = append(arr3, siHuaS)
			siHuaS3 := arrToS(arr3, " ")
			f.yinSiHua.SetTextBuf(siHuaS3)
		case "卯":
			arr4 = append(arr4, siHuaS)
			siHuaS4 := arrToS(arr4, " ")
			f.maoSiHua.SetTextBuf(siHuaS4)
		case "辰":
			arr5 = append(arr5, siHuaS)
			siHuaS5 := arrToS(arr5, " ")
			f.chenSiHua.SetTextBuf(siHuaS5)
		case "巳":
			arr6 = append(arr6, siHuaS)
			siHuaS6 := arrToS(arr6, " ")
			f.siSiHua.SetTextBuf(siHuaS6)
		case "午":
			arr7 = append(arr7, siHuaS)
			siHuaS7 := arrToS(arr, " ")
			f.wuSiHua.SetTextBuf(siHuaS7)
		case "未":
			arr8 = append(arr8, siHuaS)
			siHuaS8 := arrToS(arr8, " ")
			f.weiSiHua.SetTextBuf(siHuaS8)
		case "申":
			arr9 = append(arr9, siHuaS)
			siHuaS9 := arrToS(arr9, " ")
			f.shenSiHua.SetTextBuf(siHuaS9)
		case "酉":
			arr10 = append(arr10, siHuaS)
			siHuaS10 := arrToS(arr10, " ")
			f.youSiHua.SetTextBuf(siHuaS10)
		case "戌":
			arr11 = append(arr11, siHuaS)
			siHuaS11 := arrToS(arr11, " ")
			f.xuSiHua.SetTextBuf(siHuaS11)
		}
	}
}

//十二命宫....
func (f *TMainForm) OnFormMingGongShow(mpmap map[string]string) {
	const step = 150
	f.haiMg.Free()
	f.haiMg = vcl.NewLabel(f)
	f.haiMg.SetParent(f)
	f.haiMg.SetTop(t4)
	f.haiMg.SetLeft(l5 + step)

	f.ziMg.Free()
	f.ziMg = vcl.NewLabel(f)
	f.ziMg.SetParent(f)
	f.ziMg.SetTop(t4)
	f.ziMg.SetLeft(r4 + step)

	f.chouMg.Free()
	f.chouMg = vcl.NewLabel(f)
	f.chouMg.SetParent(f)
	f.chouMg.SetTop(t4)
	f.chouMg.SetLeft(l4 + step)

	f.yinMg.Free()
	f.yinMg = vcl.NewLabel(f)
	f.yinMg.SetParent(f)
	f.yinMg.SetTop(t4)
	f.yinMg.SetLeft(l1 + step)

	f.maoMg.Free()
	f.maoMg = vcl.NewLabel(f)
	f.maoMg.SetParent(f)
	f.maoMg.SetTop(t3)
	f.maoMg.SetLeft(l1 + step)

	f.chenMg.Free()
	f.chenMg = vcl.NewLabel(f)
	f.chenMg.SetParent(f)
	f.chenMg.SetTop(t6)
	f.chenMg.SetLeft(l1 + step)

	f.siMg.Free()
	f.siMg = vcl.NewLabel(f)
	f.siMg.SetParent(f)
	f.siMg.SetTop(top)
	f.siMg.SetLeft(l1 + step)

	f.wuMg.Free()
	f.wuMg = vcl.NewLabel(f)
	f.wuMg.SetParent(f)
	f.wuMg.SetTop(top)
	f.wuMg.SetLeft(r1 + step)

	f.weiMg.Free()
	f.weiMg = vcl.NewLabel(f)
	f.weiMg.SetParent(f)
	f.weiMg.SetTop(top)
	f.weiMg.SetLeft(l2 + step)

	f.shenMg.Free()
	f.shenMg = vcl.NewLabel(f)
	f.shenMg.SetParent(f)
	f.shenMg.SetTop(top)
	f.shenMg.SetLeft(l5 + step)

	f.youMg.Free()
	f.youMg = vcl.NewLabel(f)
	f.youMg.SetParent(f)
	f.youMg.SetTop(t6)
	f.youMg.SetLeft(l5 + step)

	f.xuMg.Free()
	f.xuMg = vcl.NewLabel(f)
	f.xuMg.SetParent(f)
	f.xuMg.SetTop(t3)
	f.xuMg.SetLeft(l5 + step)

	for zhi, star := range mpmap {
		switch zhi {
		case "亥":
			f.haiMg.SetTextBuf(star)
		case "子":
			f.ziMg.SetTextBuf(star)
		case "丑":
			f.chouMg.SetTextBuf(star)
		case "寅":
			f.yinMg.SetTextBuf(star)
		case "卯":
			f.maoMg.SetTextBuf(star)
		case "辰":
			f.chenMg.SetTextBuf(star)
		case "巳":
			f.siMg.SetTextBuf(star)
		case "午":
			f.wuMg.SetTextBuf(star)
		case "未":
			f.weiMg.SetTextBuf(star)
		case "申":
			f.shenMg.SetTextBuf(star)
		case "酉":
			f.youMg.SetTextBuf(star)
		case "戌":
			f.xuMg.SetTextBuf(star)
		}
	}
}

//十四主星+庙 红色
func (f *TMainForm) OnFormStar14Show(smap, mmap map[string]string) {
	const height1 = 30

	f.haiStar.Free()
	f.haiStar = vcl.NewLabel(f)
	f.haiStar.SetParent(f)
	f.haiStar.SetTop(t4 + height1)
	f.haiStar.SetLeft(l5)
	f.haiStar.Font().SetColor(colors.ClRed)

	f.ziStar.Free()
	f.ziStar = vcl.NewLabel(f)
	f.ziStar.SetParent(f)
	f.ziStar.SetTop(t4 + height1)
	f.ziStar.SetLeft(r4)
	f.ziStar.Font().SetColor(colors.ClRed)

	f.chouStar.Free()
	f.chouStar = vcl.NewLabel(f)
	f.chouStar.SetParent(f)
	f.chouStar.SetTop(t4 + height1)
	f.chouStar.SetLeft(l4)
	f.chouStar.Font().SetColor(colors.ClRed)

	f.yinStar.Free()
	f.yinStar = vcl.NewLabel(f)
	f.yinStar.SetParent(f)
	f.yinStar.SetTop(t4 + height1)
	f.yinStar.SetLeft(l1)
	f.yinStar.Font().SetColor(colors.ClRed)

	f.maoStar.Free()
	f.maoStar = vcl.NewLabel(f)
	f.maoStar.SetParent(f)
	f.maoStar.SetTop(t3 + height1)
	f.maoStar.SetLeft(l1)
	f.maoStar.Font().SetColor(colors.ClRed)

	f.chenStar.Free()
	f.chenStar = vcl.NewLabel(f)
	f.chenStar.SetParent(f)
	f.chenStar.SetTop(t6 + height1)
	f.chenStar.SetLeft(l1)
	f.chenStar.Font().SetColor(colors.ClRed)

	f.siStar.Free()
	f.siStar = vcl.NewLabel(f)
	f.siStar.SetParent(f)
	f.siStar.SetTop(top + height1)
	f.siStar.SetLeft(l1)
	f.siStar.Font().SetColor(colors.ClRed)

	f.wuStar.Free()
	f.wuStar = vcl.NewLabel(f)
	f.wuStar.SetParent(f)
	f.wuStar.SetTop(top + height1)
	f.wuStar.SetLeft(r1)
	f.wuStar.Font().SetColor(colors.ClRed)

	f.weiStar.Free()
	f.weiStar = vcl.NewLabel(f)
	f.weiStar.SetParent(f)
	f.weiStar.SetTop(top + height1)
	f.weiStar.SetLeft(l2)
	f.weiStar.Font().SetColor(colors.ClRed)

	f.shenStar.Free()
	f.shenStar = vcl.NewLabel(f)
	f.shenStar.SetParent(f)
	f.shenStar.SetTop(top + height1)
	f.shenStar.SetLeft(l5)
	f.shenStar.Font().SetColor(colors.ClRed)

	f.youStar.Free()
	f.youStar = vcl.NewLabel(f)
	f.youStar.SetParent(f)
	f.youStar.SetTop(t6 + height1)
	f.youStar.SetLeft(l5)
	f.youStar.Font().SetColor(colors.ClRed)

	f.xuStar.Free()
	f.xuStar = vcl.NewLabel(f)
	f.xuStar.SetParent(f)
	f.xuStar.SetTop(t3 + height1)
	f.xuStar.SetLeft(l5)
	f.xuStar.Font().SetColor(colors.ClRed)

	var shai, szi, schou, syin, smao, schen, ssi, swu, swei, sshen, syou, sxu []string
	for star, zhi := range smap {
		switch zhi {
		case "亥":
			shai = append(shai, star)
			sm := addMiaoYG(shai, mmap)
			f.haiStar.SetTextBuf(sm)
		case "子":
			szi = append(szi, star)
			sm := addMiaoYG(szi, mmap)
			f.ziStar.SetTextBuf(sm)
		case "丑":
			schou = append(schou, star)
			sm := addMiaoYG(schou, mmap)
			f.chouStar.SetTextBuf(sm)
		case "寅":
			syin = append(syin, star)
			sm := addMiaoYG(syin, mmap)
			f.yinStar.SetTextBuf(sm)
		case "卯":
			smao = append(smao, star)
			sm := addMiaoYG(smao, mmap)
			f.maoStar.SetTextBuf(sm)
		case "辰":
			schen = append(schen, star)
			sm := addMiaoYG(schen, mmap)
			f.chenStar.SetTextBuf(sm)
		case "巳":
			ssi = append(ssi, star)
			sm := addMiaoYG(ssi, mmap)
			f.siStar.SetTextBuf(sm)
		case "午":
			swu = append(swu, star)
			sm := addMiaoYG(swu, mmap)
			f.wuStar.SetTextBuf(sm)
		case "未":
			swei = append(swei, star)
			sm := addMiaoYG(swei, mmap)
			f.weiStar.SetTextBuf(sm)
		case "申":
			sshen = append(sshen, star)
			sm := addMiaoYG(sshen, mmap)
			f.shenStar.SetTextBuf(sm)
		case "酉":
			syou = append(syou, star)
			sm := addMiaoYG(syou, mmap)
			f.youStar.SetTextBuf(sm)
		case "戌":
			sxu = append(sxu, star)
			sm := addMiaoYG(sxu, mmap)
			f.xuStar.SetTextBuf(sm)
		}
	}
}

//年干系诸星+庙 蓝色
func (f *TMainForm) OnFormYearGanShow(ygmap, mmap map[string]string) {
	const height1 = 2 * 30
	f.haiYearGanStar.Free()
	f.haiYearGanStar = vcl.NewLabel(f)
	f.haiYearGanStar.SetParent(f)
	f.haiYearGanStar.SetTop(t4 + height1)
	f.haiYearGanStar.SetLeft(l5)
	blue := f.haiYearGanStar.Color().RGB(72, 118, 255)
	f.haiYearGanStar.Font().SetColor(blue)

	f.ziYearGanStar.Free()
	f.ziYearGanStar = vcl.NewLabel(f)
	f.ziYearGanStar.SetParent(f)
	f.ziYearGanStar.SetTop(t4 + height1)
	f.ziYearGanStar.SetLeft(r4)
	f.ziYearGanStar.Font().SetColor(blue)

	f.chouYearGanStar.Free()
	f.chouYearGanStar = vcl.NewLabel(f)
	f.chouYearGanStar.SetParent(f)
	f.chouYearGanStar.SetTop(t4 + height1)
	f.chouYearGanStar.SetLeft(l4)
	f.chouYearGanStar.Font().SetColor(blue)

	f.yinYearGanStar.Free()
	f.yinYearGanStar = vcl.NewLabel(f)
	f.yinYearGanStar.SetParent(f)
	f.yinYearGanStar.SetTop(t4 + height1)
	f.yinYearGanStar.SetLeft(l1)
	f.yinYearGanStar.SetColor(blue)
	f.yinYearGanStar.Font().SetColor(blue)

	f.maoYearGanStar.Free()
	f.maoYearGanStar = vcl.NewLabel(f)
	f.maoYearGanStar.SetParent(f)
	f.maoYearGanStar.SetTop(t3 + height1)
	f.maoYearGanStar.SetLeft(l1)
	f.maoYearGanStar.Font().SetColor(blue)

	f.chenYearGanStar.Free()
	f.chenYearGanStar = vcl.NewLabel(f)
	f.chenYearGanStar.SetParent(f)
	f.chenYearGanStar.SetTop(t6 + height1)
	f.chenYearGanStar.SetLeft(l1)
	f.chenYearGanStar.Font().SetColor(blue)

	f.siYearGanStar.Free()
	f.siYearGanStar = vcl.NewLabel(f)
	f.siYearGanStar.SetParent(f)
	f.siYearGanStar.SetTop(top + height1)
	f.siYearGanStar.SetLeft(l1)
	f.siYearGanStar.Font().SetColor(blue)

	f.wuYearGanStar.Free()
	f.wuYearGanStar = vcl.NewLabel(f)
	f.wuYearGanStar.SetParent(f)
	f.wuYearGanStar.SetTop(top + height1)
	f.wuYearGanStar.SetLeft(r1)
	f.wuYearGanStar.Font().SetColor(blue)

	f.weiYearGanStar.Free()
	f.weiYearGanStar = vcl.NewLabel(f)
	f.weiYearGanStar.SetParent(f)
	f.weiYearGanStar.SetTop(top + height1)
	f.weiYearGanStar.SetLeft(l2)
	f.weiYearGanStar.Font().SetColor(blue)

	f.shenYearGanStar.Free()
	f.shenYearGanStar = vcl.NewLabel(f)
	f.shenYearGanStar.SetParent(f)
	f.shenYearGanStar.SetTop(top + height1)
	f.shenYearGanStar.SetLeft(l5)
	f.shenYearGanStar.Font().SetColor(blue)

	f.youYearGanStar.Free()
	f.youYearGanStar = vcl.NewLabel(f)
	f.youYearGanStar.SetParent(f)
	f.youYearGanStar.SetTop(t6 + height1)
	f.youYearGanStar.SetLeft(l5)
	f.youYearGanStar.Font().SetColor(blue)

	f.xuYearGanStar.Free()
	f.xuYearGanStar = vcl.NewLabel(f)
	f.xuYearGanStar.SetParent(f)
	f.xuYearGanStar.SetTop(t3 + height1)
	f.xuYearGanStar.SetLeft(l5)
	f.xuYearGanStar.Font().SetColor(blue)

	var shai, szi, schou, syin, smao, schen, ssi, swu, swei, sshen, syou, sxu []string
	for star, zhi := range ygmap {
		switch zhi {
		case "亥":
			shai = append(shai, star)
			sm := addMiaoYG(shai, mmap)
			f.haiYearGanStar.SetTextBuf(sm)
		case "子":
			szi = append(szi, star)
			sm := addMiaoYG(szi, mmap)
			f.ziYearGanStar.SetTextBuf(sm)
		case "丑":
			schou = append(schou, star)
			sm := addMiaoYG(schou, mmap)
			f.chouYearGanStar.SetTextBuf(sm)
		case "寅":
			syin = append(syin, star)
			sm := addMiaoYG(syin, mmap)
			f.yinYearGanStar.SetTextBuf(sm)
		case "卯":
			smao = append(smao, star)
			sm := addMiaoYG(smao, mmap)
			f.maoYearGanStar.SetTextBuf(sm)
		case "辰":
			schen = append(schen, star)
			sm := addMiaoYG(schen, mmap)
			f.chenYearGanStar.SetTextBuf(sm)
		case "巳":
			ssi = append(ssi, star)
			sm := addMiaoYG(ssi, mmap)
			f.siYearGanStar.SetTextBuf(sm)
		case "午":
			swu = append(swu, star)
			sm := addMiaoYG(swu, mmap)
			f.wuYearGanStar.SetTextBuf(sm)
		case "未":
			swei = append(swei, star)
			sm := addMiaoYG(swei, mmap)
			f.weiYearGanStar.SetTextBuf(sm)
		case "申":
			sshen = append(sshen, star)
			sm := addMiaoYG(sshen, mmap)
			f.shenYearGanStar.SetTextBuf(sm)
		case "酉":
			syou = append(syou, star)
			sm := addMiaoYG(syou, mmap)
			f.youYearGanStar.SetTextBuf(sm)
		case "戌":
			sxu = append(sxu, star)
			sm := addMiaoYG(sxu, mmap)
			f.xuYearGanStar.SetTextBuf(sm)
		}
	}
}

//年支系诸星　浅蓝色
func (f *TMainForm) OnFormYearZhiShow(yzmap map[string]string) {
	const height1 = 3 * 30

	f.haiYearZhiStar.Free()
	f.haiYearZhiStar = vcl.NewLabel(f)
	f.haiYearZhiStar.SetParent(f)
	f.haiYearZhiStar.SetTop(t4 + height1)
	f.haiYearZhiStar.SetLeft(l5)
	skyBlue := f.haiYearZhiStar.Color().RGB(135, 206, 255)
	f.haiYearZhiStar.Font().SetColor(skyBlue)

	f.ziYearZhiStar.Free()
	f.ziYearZhiStar = vcl.NewLabel(f)
	f.ziYearZhiStar.SetParent(f)
	f.ziYearZhiStar.SetTop(t4 + height1)
	f.ziYearZhiStar.SetLeft(r4)
	f.ziYearZhiStar.Font().SetColor(skyBlue)

	f.chouYearZhiStar.Free()
	f.chouYearZhiStar = vcl.NewLabel(f)
	f.chouYearZhiStar.SetParent(f)
	f.chouYearZhiStar.SetTop(t4 + height1)
	f.chouYearZhiStar.SetLeft(l4)
	f.chouYearZhiStar.Font().SetColor(skyBlue)

	f.yinYearZhiStar.Free()
	f.yinYearZhiStar = vcl.NewLabel(f)
	f.yinYearZhiStar.SetParent(f)
	f.yinYearZhiStar.SetTop(t4 + height1)
	f.yinYearZhiStar.SetLeft(l1)
	f.yinYearZhiStar.Font().SetColor(skyBlue)

	f.maoYearZhiStar.Free()
	f.maoYearZhiStar = vcl.NewLabel(f)
	f.maoYearZhiStar.SetParent(f)
	f.maoYearZhiStar.SetTop(t3 + height1)
	f.maoYearZhiStar.SetLeft(l1)
	f.maoYearZhiStar.Font().SetColor(skyBlue)

	f.chenYearZhiStar.Free()
	f.chenYearZhiStar = vcl.NewLabel(f)
	f.chenYearZhiStar.SetParent(f)
	f.chenYearZhiStar.SetTop(t6 + height1)
	f.chenYearZhiStar.SetLeft(l1)
	f.chenYearZhiStar.Font().SetColor(skyBlue)

	f.siYearZhiStar.Free()
	f.siYearZhiStar = vcl.NewLabel(f)
	f.siYearZhiStar.SetParent(f)
	f.siYearZhiStar.SetTop(top + height1)
	f.siYearZhiStar.SetLeft(l1)
	f.siYearZhiStar.Font().SetColor(skyBlue)

	f.wuYearZhiStar.Free()
	f.wuYearZhiStar = vcl.NewLabel(f)
	f.wuYearZhiStar.SetParent(f)
	f.wuYearZhiStar.SetTop(top + height1)
	f.wuYearZhiStar.SetLeft(r1)
	f.wuYearZhiStar.Font().SetColor(skyBlue)

	f.weiYearZhiStar.Free()
	f.weiYearZhiStar = vcl.NewLabel(f)
	f.weiYearZhiStar.SetParent(f)
	f.weiYearZhiStar.SetTop(top + height1)
	f.weiYearZhiStar.SetLeft(l2)
	f.weiYearZhiStar.Font().SetColor(skyBlue)

	f.shenYearZhiStar.Free()
	f.shenYearZhiStar = vcl.NewLabel(f)
	f.shenYearZhiStar.SetParent(f)
	f.shenYearZhiStar.SetTop(top + height1)
	f.shenYearZhiStar.SetLeft(l5)
	f.shenYearZhiStar.Font().SetColor(skyBlue)

	f.youYearZhiStar.Free()
	f.youYearZhiStar = vcl.NewLabel(f)
	f.youYearZhiStar.SetParent(f)
	f.youYearZhiStar.SetTop(t6 + height1)
	f.youYearZhiStar.SetLeft(l5)
	f.youYearZhiStar.Font().SetColor(skyBlue)

	f.xuYearZhiStar.Free()
	f.xuYearZhiStar = vcl.NewLabel(f)
	f.xuYearZhiStar.SetParent(f)
	f.xuYearZhiStar.SetTop(t3 + height1)
	f.xuYearZhiStar.SetLeft(l5)
	f.xuYearZhiStar.Font().SetColor(skyBlue)

	var arr0, arr1, arr2, arr3, arr4, arr5, arr6, arr7, arr8, arr9, arr10, arr11 []string
	for star, zhi := range yzmap {
		switch zhi {
		case "亥":
			arr0 = append(arr0, star)
			star0 := arrToS(arr0, " ")
			f.haiYearZhiStar.SetTextBuf(star0)
		case "子":
			arr1 = append(arr1, star)
			star1 := arrToS(arr1, " ")
			f.ziYearZhiStar.SetTextBuf(star1)
		case "丑":
			arr2 = append(arr2, star)
			star2 := arrToS(arr2, " ")
			f.chouYearZhiStar.SetTextBuf(star2)
		case "寅":
			arr3 = append(arr3, star)
			star3 := arrToS(arr3, " ")
			f.yinYearZhiStar.SetTextBuf(star3)
		case "卯":
			arr4 = append(arr4, star)
			star4 := arrToS(arr4, " ")
			f.maoYearZhiStar.SetTextBuf(star4)
		case "辰":
			arr5 = append(arr5, star)
			star5 := arrToS(arr5, " ")
			f.chenYearZhiStar.SetTextBuf(star5)
		case "巳":
			arr6 = append(arr6, star)
			star6 := arrToS(arr6, " ")
			f.siYearZhiStar.SetTextBuf(star6)
		case "午":
			arr7 = append(arr7, star)
			star7 := arrToS(arr7, " ")
			f.wuYearZhiStar.SetTextBuf(star7)
		case "未":
			arr8 = append(arr8, star)
			star8 := arrToS(arr8, " ")
			f.weiYearZhiStar.SetTextBuf(star8)
		case "申":
			arr9 = append(arr9, star)
			star9 := arrToS(arr9, " ")
			f.shenYearZhiStar.SetTextBuf(star9)
		case "酉":
			arr10 = append(arr10, star)
			star10 := arrToS(arr10, " ")
			f.youYearZhiStar.SetTextBuf(star10)
		case "戌":
			arr11 = append(arr11, star)
			star11 := arrToS(arr11, " ")
			f.xuYearZhiStar.SetTextBuf(star11)
		}
	}
}

//月系诸星　绿色
func (f *TMainForm) OnFormMonthShow(msmap map[string]string) {
	const height1 = 3 * 30
	const step = 140

	f.haiMonthStar.Free()
	f.haiMonthStar = vcl.NewLabel(f)
	f.haiMonthStar.SetParent(f)
	f.haiMonthStar.SetTop(t4 + height1)
	f.haiMonthStar.SetLeft(l5 + step)
	limeGreen := f.haiMonthStar.Color().RGB(50, 205, 50)
	f.haiMonthStar.Font().SetColor(limeGreen)

	f.ziMonthStar.Free()
	f.ziMonthStar = vcl.NewLabel(f)
	f.ziMonthStar.SetParent(f)
	f.ziMonthStar.SetTop(t4 + height1)
	f.ziMonthStar.SetLeft(r4 + step)
	f.ziMonthStar.Font().SetColor(limeGreen)

	f.chouMonthStar.Free()
	f.chouMonthStar = vcl.NewLabel(f)
	f.chouMonthStar.SetParent(f)
	f.chouMonthStar.SetTop(t4 + height1)
	f.chouMonthStar.SetLeft(l4 + step)
	f.chouMonthStar.Font().SetColor(limeGreen)

	f.yinMonthStar.Free()
	f.yinMonthStar = vcl.NewLabel(f)
	f.yinMonthStar.SetParent(f)
	f.yinMonthStar.SetTop(t4 + height1)
	f.yinMonthStar.SetLeft(l1 + step)
	f.yinMonthStar.Font().SetColor(limeGreen)

	f.maoMonthStar.Free()
	f.maoMonthStar = vcl.NewLabel(f)
	f.maoMonthStar.SetParent(f)
	f.maoMonthStar.SetTop(t3 + height1)
	f.maoMonthStar.SetLeft(l1 + step)
	f.maoMonthStar.Font().SetColor(limeGreen)

	f.chenMonthStar.Free()
	f.chenMonthStar = vcl.NewLabel(f)
	f.chenMonthStar.SetParent(f)
	f.chenMonthStar.SetTop(t6 + height1)
	f.chenMonthStar.SetLeft(l1 + step)
	f.chenMonthStar.Font().SetColor(limeGreen)

	f.siMonthStar.Free()
	f.siMonthStar = vcl.NewLabel(f)
	f.siMonthStar.SetParent(f)
	f.siMonthStar.SetTop(top + height1)
	f.siMonthStar.SetLeft(l1 + step)
	f.siMonthStar.Font().SetColor(limeGreen)

	f.wuMonthStar.Free()
	f.wuMonthStar = vcl.NewLabel(f)
	f.wuMonthStar.SetParent(f)
	f.wuMonthStar.SetTop(top + height1)
	f.wuMonthStar.SetLeft(r1 + step)
	f.wuMonthStar.Font().SetColor(limeGreen)

	f.weiMonthStar.Free()
	f.weiMonthStar = vcl.NewLabel(f)
	f.weiMonthStar.SetParent(f)
	f.weiMonthStar.SetTop(top + height1)
	f.weiMonthStar.SetLeft(l2 + step)
	f.weiMonthStar.Font().SetColor(limeGreen)

	f.shenMonthStar.Free()
	f.shenMonthStar = vcl.NewLabel(f)
	f.shenMonthStar.SetParent(f)
	f.shenMonthStar.SetTop(top + height1)
	f.shenMonthStar.SetLeft(l5 + step)
	f.shenMonthStar.Font().SetColor(limeGreen)

	f.youMonthStar.Free()
	f.youMonthStar = vcl.NewLabel(f)
	f.youMonthStar.SetParent(f)
	f.youMonthStar.SetTop(t6 + height1)
	f.youMonthStar.SetLeft(l5 + step)
	f.youMonthStar.Font().SetColor(limeGreen)

	f.xuMonthStar.Free()
	f.xuMonthStar = vcl.NewLabel(f)
	f.xuMonthStar.SetParent(f)
	f.xuMonthStar.SetTop(t3 + height1)
	f.xuMonthStar.SetLeft(l5 + step)
	f.xuMonthStar.Font().SetColor(limeGreen)

	var arr, arr1, arr2, arr3, arr4, arr5, arr6, arr7, arr8, arr9, arr10, arr11 []string
	for star, zhi := range msmap {
		switch zhi {
		case "亥":
			arr = append(arr, star)
			star1 := arrToS(arr, " ")
			f.haiMonthStar.SetTextBuf(star1)
		case "子":
			arr1 = append(arr1, star)
			star2 := arrToS(arr1, " ")
			f.ziMonthStar.SetTextBuf(star2)
		case "丑":
			arr2 = append(arr2, star)
			star2 := arrToS(arr2, " ")
			f.chouMonthStar.SetTextBuf(star2)
		case "寅":
			arr3 = append(arr3, star)
			star3 := arrToS(arr3, " ")
			f.yinMonthStar.SetTextBuf(star3)
		case "卯":
			arr4 = append(arr4, star)
			star4 := arrToS(arr4, " ")
			f.maoMonthStar.SetTextBuf(star4)
		case "辰":
			arr5 = append(arr5, star)
			star5 := arrToS(arr5, " ")
			f.chenMonthStar.SetTextBuf(star5)
		case "巳":
			arr6 = append(arr6, star)
			star6 := arrToS(arr6, " ")
			f.siMonthStar.SetTextBuf(star6)
		case "午":
			arr7 = append(arr7, star)
			star7 := arrToS(arr, " ")
			f.wuMonthStar.SetTextBuf(star7)
		case "未":
			arr8 = append(arr8, star)
			star8 := arrToS(arr8, " ")
			f.weiMonthStar.SetTextBuf(star8)
		case "申":
			arr9 = append(arr9, star)
			star9 := arrToS(arr9, " ")
			f.shenMonthStar.SetTextBuf(star9)
		case "酉":
			arr10 = append(arr10, star)
			star10 := arrToS(arr10, " ")
			f.youMonthStar.SetTextBuf(star10)
		case "戌":
			arr11 = append(arr11, star)
			star11 := arrToS(arr11, " ")
			f.xuMonthStar.SetTextBuf(star11)
		}
	}
}

//时系诸星 颜色:orchid
func (f *TMainForm) OnFormHourShow(hmap map[string]string) {
	const height1 = 4 * 30

	f.haiHourStar.Free()
	f.haiHourStar = vcl.NewLabel(f)
	f.haiHourStar.SetParent(f)
	f.haiHourStar.SetTop(t4 + height1)
	f.haiHourStar.SetLeft(l5)
	orchid := f.haiHourStar.Color().RGB(218, 112, 214)
	f.haiHourStar.Font().SetColor(orchid)

	f.ziHourStar.Free()
	f.ziHourStar = vcl.NewLabel(f)
	f.ziHourStar.SetParent(f)
	f.ziHourStar.SetTop(t4 + height1)
	f.ziHourStar.SetLeft(r4)
	f.ziHourStar.Font().SetColor(orchid)

	f.chouHourStar.Free()
	f.chouHourStar = vcl.NewLabel(f)
	f.chouHourStar.SetParent(f)
	f.chouHourStar.SetTop(t4 + height1)
	f.chouHourStar.SetLeft(l4)
	f.chouHourStar.Font().SetColor(orchid)

	f.yinHourStar.Free()
	f.yinHourStar = vcl.NewLabel(f)
	f.yinHourStar.SetParent(f)
	f.yinHourStar.SetTop(t4 + height1)
	f.yinHourStar.SetLeft(l1)
	f.yinHourStar.Font().SetColor(orchid)

	f.maoHourStar.Free()
	f.maoHourStar = vcl.NewLabel(f)
	f.maoHourStar.SetParent(f)
	f.maoHourStar.SetTop(t3 + height1)
	f.maoHourStar.SetLeft(l1)
	f.maoHourStar.Font().SetColor(orchid)

	f.chenHourStar.Free()
	f.chenHourStar = vcl.NewLabel(f)
	f.chenHourStar.SetParent(f)
	f.chenHourStar.SetTop(t6 + height1)
	f.chenHourStar.SetLeft(l1)
	f.chenHourStar.Font().SetColor(orchid)

	f.siHourStar.Free()
	f.siHourStar = vcl.NewLabel(f)
	f.siHourStar.SetParent(f)
	f.siHourStar.SetTop(top + height1)
	f.siHourStar.SetLeft(l1)
	f.siHourStar.Font().SetColor(orchid)

	f.wuHourStar.Free()
	f.wuHourStar = vcl.NewLabel(f)
	f.wuHourStar.SetParent(f)
	f.wuHourStar.SetTop(top + height1)
	f.wuHourStar.SetLeft(r1)
	f.wuHourStar.Font().SetColor(orchid)

	f.weiHourStar.Free()
	f.weiHourStar = vcl.NewLabel(f)
	f.weiHourStar.SetParent(f)
	f.weiHourStar.SetTop(top + height1)
	f.weiHourStar.SetLeft(l2)
	f.weiHourStar.Font().SetColor(orchid)

	f.shenHourStar.Free()
	f.shenHourStar = vcl.NewLabel(f)
	f.shenHourStar.SetParent(f)
	f.shenHourStar.SetTop(top + height1)
	f.shenHourStar.SetLeft(l5)
	f.shenHourStar.Font().SetColor(orchid)

	f.youHourStar.Free()
	f.youHourStar = vcl.NewLabel(f)
	f.youHourStar.SetParent(f)
	f.youHourStar.SetTop(t6 + height1)
	f.youHourStar.SetLeft(l5)
	f.youHourStar.Font().SetColor(orchid)

	f.xuHourStar.Free()
	f.xuHourStar = vcl.NewLabel(f)
	f.xuHourStar.SetParent(f)
	f.xuHourStar.SetTop(t3 + height1)
	f.xuHourStar.SetLeft(l5)
	f.xuHourStar.Font().SetColor(orchid)
	var arr, arr1, arr2, arr3, arr4, arr5, arr6, arr7, arr8, arr9, arr10, arr11 []string
	for star, zhi := range hmap {
		switch zhi {
		case "亥":
			arr = append(arr, star)
			star1 := arrToS(arr, " ")
			f.haiHourStar.SetTextBuf(star1)
		case "子":
			arr1 = append(arr1, star)
			star2 := arrToS(arr1, " ")
			f.ziHourStar.SetTextBuf(star2)
		case "丑":
			arr2 = append(arr2, star)
			star2 := arrToS(arr2, " ")
			f.chouHourStar.SetTextBuf(star2)
		case "寅":
			arr3 = append(arr3, star)
			star3 := arrToS(arr3, " ")
			f.yinHourStar.SetTextBuf(star3)
		case "卯":
			arr4 = append(arr4, star)
			star4 := arrToS(arr4, " ")
			f.maoHourStar.SetTextBuf(star4)
		case "辰":
			arr5 = append(arr5, star)
			star5 := arrToS(arr5, " ")
			f.chenHourStar.SetTextBuf(star5)
		case "巳":
			arr6 = append(arr6, star)
			star6 := arrToS(arr6, " ")
			f.siHourStar.SetTextBuf(star6)
		case "午":
			arr7 = append(arr7, star)
			star7 := arrToS(arr, " ")
			f.wuHourStar.SetTextBuf(star7)
		case "未":
			arr8 = append(arr8, star)
			star8 := arrToS(arr8, " ")
			f.weiHourStar.SetTextBuf(star8)
		case "申":
			arr9 = append(arr9, star)
			star9 := arrToS(arr9, " ")
			f.shenHourStar.SetTextBuf(star9)
		case "酉":
			arr10 = append(arr10, star)
			star10 := arrToS(arr10, " ")
			f.youHourStar.SetTextBuf(star10)
		case "戌":
			arr11 = append(arr11, star)
			star11 := arrToS(arr11, " ")
			f.xuHourStar.SetTextBuf(star11)
		}
	}

}

//其他星系＋庙　颜色:OrangeRed1	255 69 0
func (f *TMainForm) OnFormOtherShow(omap, mmap map[string]string) {
	const height1 = 4 * 30
	const step = 165

	f.haiOtherStar.Free()
	f.haiOtherStar = vcl.NewLabel(f)
	f.haiOtherStar.SetParent(f)
	f.haiOtherStar.SetTop(t4 + height1)
	f.haiOtherStar.SetLeft(l5 + step)
	orangeRed := f.haiOtherStar.Color().RGB(255, 69, 0)
	f.haiOtherStar.Font().SetColor(orangeRed)

	f.ziOtherStar.Free()
	f.ziOtherStar = vcl.NewLabel(f)
	f.ziOtherStar.SetParent(f)
	f.ziOtherStar.SetTop(t4 + height1)
	f.ziOtherStar.SetLeft(r4 + step)
	f.ziOtherStar.Font().SetColor(orangeRed)

	f.chouOtherStar.Free()
	f.chouOtherStar = vcl.NewLabel(f)
	f.chouOtherStar.SetParent(f)
	f.chouOtherStar.SetTop(t4 + height1)
	f.chouOtherStar.SetLeft(l4 + step)
	f.chouOtherStar.Font().SetColor(orangeRed)

	f.yinOtherStar.Free()
	f.yinOtherStar = vcl.NewLabel(f)
	f.yinOtherStar.SetParent(f)
	f.yinOtherStar.SetTop(t4 + height1)
	f.yinOtherStar.SetLeft(l1 + step)
	f.yinOtherStar.Font().SetColor(orangeRed)

	f.maoOtherStar.Free()
	f.maoOtherStar = vcl.NewLabel(f)
	f.maoOtherStar.SetParent(f)
	f.maoOtherStar.SetTop(t3 + height1)
	f.maoOtherStar.SetLeft(l1 + step)
	f.maoOtherStar.Font().SetColor(orangeRed)

	f.chenOtherStar.Free()
	f.chenOtherStar = vcl.NewLabel(f)
	f.chenOtherStar.SetParent(f)
	f.chenOtherStar.SetTop(t6 + height1)
	f.chenOtherStar.SetLeft(l1 + step)
	f.chenOtherStar.Font().SetColor(orangeRed)

	f.siOtherStar.Free()
	f.siOtherStar = vcl.NewLabel(f)
	f.siOtherStar.SetParent(f)
	f.siOtherStar.SetTop(top + height1)
	f.siOtherStar.SetLeft(l1 + step)
	f.siOtherStar.Font().SetColor(orangeRed)

	f.wuOtherStar.Free()
	f.wuOtherStar = vcl.NewLabel(f)
	f.wuOtherStar.SetParent(f)
	f.wuOtherStar.SetTop(top + height1)
	f.wuOtherStar.SetLeft(r1 + step)
	f.wuOtherStar.Font().SetColor(orangeRed)

	f.weiOtherStar.Free()
	f.weiOtherStar = vcl.NewLabel(f)
	f.weiOtherStar.SetParent(f)
	f.weiOtherStar.SetTop(top + height1)
	f.weiOtherStar.SetLeft(l2 + step)
	f.weiOtherStar.Font().SetColor(orangeRed)

	f.shenOtherStar.Free()
	f.shenOtherStar = vcl.NewLabel(f)
	f.shenOtherStar.SetParent(f)
	f.shenOtherStar.SetTop(top + height1)
	f.shenOtherStar.SetLeft(l5 + step)
	f.shenOtherStar.Font().SetColor(orangeRed)

	f.youOtherStar.Free()
	f.youOtherStar = vcl.NewLabel(f)
	f.youOtherStar.SetParent(f)
	f.youOtherStar.SetTop(t6 + height1)
	f.youOtherStar.SetLeft(l5 + step)
	f.youOtherStar.Font().SetColor(orangeRed)

	f.xuOtherStar.Free()
	f.xuOtherStar = vcl.NewLabel(f)
	f.xuOtherStar.SetParent(f)
	f.xuOtherStar.SetTop(t3 + height1)
	f.xuOtherStar.SetLeft(l5 + step)
	f.xuOtherStar.Font().SetColor(orangeRed)

	var shai, szi, schou, syin, smao, schen, ssi, swu, swei, sshen, syou, sxu []string
	for star, zhi := range omap {
		switch zhi {
		case "亥":
			shai = append(shai, star)
			sm := addMiaoYG(shai, mmap)
			f.haiOtherStar.SetTextBuf(sm)
		case "子":
			szi = append(szi, star)
			sm := addMiaoYG(szi, mmap)
			f.ziOtherStar.SetTextBuf(sm)
		case "丑":
			schou = append(schou, star)
			sm := addMiaoYG(schou, mmap)
			f.chouOtherStar.SetTextBuf(sm)
		case "寅":
			syin = append(syin, star)
			sm := addMiaoYG(syin, mmap)
			f.yinOtherStar.SetTextBuf(sm)
		case "卯":
			smao = append(smao, star)
			sm := addMiaoYG(smao, mmap)
			f.maoOtherStar.SetTextBuf(sm)
		case "辰":
			schen = append(schen, star)
			sm := addMiaoYG(schen, mmap)
			f.chenOtherStar.SetTextBuf(sm)
		case "巳":
			ssi = append(ssi, star)
			sm := addMiaoYG(ssi, mmap)
			f.siOtherStar.SetTextBuf(sm)
		case "午":
			swu = append(swu, star)
			sm := addMiaoYG(swu, mmap)
			f.wuOtherStar.SetTextBuf(sm)
		case "未":
			swei = append(swei, star)
			sm := addMiaoYG(swei, mmap)
			f.weiOtherStar.SetTextBuf(sm)
		case "申":
			sshen = append(sshen, star)
			sm := addMiaoYG(sshen, mmap)
			f.shenOtherStar.SetTextBuf(sm)
		case "酉":
			syou = append(syou, star)
			sm := addMiaoYG(syou, mmap)
			f.youOtherStar.SetTextBuf(sm)
		case "戌":
			sxu = append(sxu, star)
			sm := addMiaoYG(sxu, mmap)
			f.xuOtherStar.SetTextBuf(sm)
		}
	}
}

//大限
func (f *TMainForm) OnFormDaXianShow(dxmap map[string]string) {
	const height1 = 5 * 30

	f.haiDaXian.Free()
	f.haiDaXian = vcl.NewLabel(f)
	f.haiDaXian.SetParent(f)
	f.haiDaXian.SetTop(t4 + height1)
	f.haiDaXian.SetLeft(l5)

	f.ziDaXian.Free()
	f.ziDaXian = vcl.NewLabel(f)
	f.ziDaXian.SetParent(f)
	f.ziDaXian.SetTop(t4 + height1)
	f.ziDaXian.SetLeft(r4)

	f.chouDaXian.Free()
	f.chouDaXian = vcl.NewLabel(f)
	f.chouDaXian.SetParent(f)
	f.chouDaXian.SetTop(t4 + height1)
	f.chouDaXian.SetLeft(l4)

	f.yinDaXian.Free()
	f.yinDaXian = vcl.NewLabel(f)
	f.yinDaXian.SetParent(f)
	f.yinDaXian.SetTop(t4 + height1)
	f.yinDaXian.SetLeft(l1)

	f.maoDaXian.Free()
	f.maoDaXian = vcl.NewLabel(f)
	f.maoDaXian.SetParent(f)
	f.maoDaXian.SetTop(t3 + height1)
	f.maoDaXian.SetLeft(l1)

	f.chenDaXian.Free()
	f.chenDaXian = vcl.NewLabel(f)
	f.chenDaXian.SetParent(f)
	f.chenDaXian.SetTop(t6 + height1)
	f.chenDaXian.SetLeft(l1)

	f.siDaXian.Free()
	f.siDaXian = vcl.NewLabel(f)
	f.siDaXian.SetParent(f)
	f.siDaXian.SetTop(top + height1)
	f.siDaXian.SetLeft(l1)

	f.wuDaXian.Free()
	f.wuDaXian = vcl.NewLabel(f)
	f.wuDaXian.SetParent(f)
	f.wuDaXian.SetTop(top + height1)
	f.wuDaXian.SetLeft(r1)

	f.weiDaXian.Free()
	f.weiDaXian = vcl.NewLabel(f)
	f.weiDaXian.SetParent(f)
	f.weiDaXian.SetTop(top + height1)
	f.weiDaXian.SetLeft(l2)

	f.shenDaXian.Free()
	f.shenDaXian = vcl.NewLabel(f)
	f.shenDaXian.SetParent(f)
	f.shenDaXian.SetTop(top + height1)
	f.shenDaXian.SetLeft(l5)

	f.youDaXian.Free()
	f.youDaXian = vcl.NewLabel(f)
	f.youDaXian.SetParent(f)
	f.youDaXian.SetTop(t6 + height1)
	f.youDaXian.SetLeft(l5)

	f.xuDaXian.Free()
	f.xuDaXian = vcl.NewLabel(f)
	f.xuDaXian.SetParent(f)
	f.xuDaXian.SetTop(t3 + height1)
	f.xuDaXian.SetLeft(l5)

	for zhi, xian := range dxmap {
		switch zhi {
		case "亥":
			f.haiDaXian.SetTextBuf(xian)
		case "子":
			f.ziDaXian.SetTextBuf(xian)
		case "丑":
			f.chouDaXian.SetTextBuf(xian)
		case "寅":
			f.yinDaXian.SetTextBuf(xian)
		case "卯":
			f.maoDaXian.SetTextBuf(xian)
		case "辰":
			f.chenDaXian.SetTextBuf(xian)
		case "巳":
			f.siDaXian.SetTextBuf(xian)
		case "午":
			f.wuDaXian.SetTextBuf(xian)
		case "未":
			f.weiDaXian.SetTextBuf(xian)
		case "申":
			f.shenDaXian.SetTextBuf(xian)
		case "酉":
			f.youDaXian.SetTextBuf(xian)
		case "戌":
			f.xuDaXian.SetTextBuf(xian)
		}
	}
}

//颜色说明
func (f *TMainForm) OnFormColorShow() {
	const (
		heightx = 855
		w14     = iota * 120
		wyg
		wyz
		wcm
		wch
		wco
	)
	f.c14.Free()
	f.c14 = vcl.NewLabel(f)
	f.c14.SetParent(f)
	f.c14.SetLeft(l1)
	f.c14.SetWidth(w14)
	f.c14.SetTop(top + heightx)
	f.c14.SetTextBuf("颜色说明: 14主星颜色")
	f.c14.Font().SetColor(colors.ClRed)

	f.cyg.Free()
	f.cyg = vcl.NewLabel(f)
	f.cyg.SetParent(f)
	f.cyg.SetLeft(wyg)
	f.cyg.SetTop(top + heightx)
	f.cyg.SetTextBuf("年干系诸星颜色")
	blue := f.cyg.Font().Color().RGB(72, 118, 255)
	f.cyg.Font().SetColor(blue)

	f.cyz.Free()
	f.cyz = vcl.NewLabel(f)
	f.cyz.SetParent(f)
	f.cyz.SetLeft(wyz)
	f.cyz.SetTop(top + heightx)
	f.cyz.SetTextBuf("年支系诸星颜色")
	skyBlue := f.cyz.Font().Color().RGB(135, 206, 255)
	f.cyz.Font().SetColor(skyBlue)

	f.cm.Free()
	f.cm = vcl.NewLabel(f)
	f.cm.SetParent(f)
	f.cm.SetLeft(wcm)
	f.cm.SetTop(top + heightx)
	f.cm.SetTextBuf("月系诸星颜色")
	limeGreen := f.cm.Font().Color().RGB(50, 205, 50)
	f.cm.Font().SetColor(limeGreen)

	f.ch.Free()
	f.ch = vcl.NewLabel(f)
	f.ch.SetParent(f)
	f.ch.SetLeft(wch)
	f.ch.SetTop(top + heightx)
	f.ch.SetTextBuf(" 时系诸星颜色")
	orchid := f.ch.Font().Color().RGB(218, 112, 214)
	f.ch.Font().SetColor(orchid)

	f.co.Free()
	f.co = vcl.NewLabel(f)
	f.co.SetParent(f)
	f.co.SetLeft(wco)
	f.co.SetTop(top + heightx)
	f.co.SetTextBuf("其他诸星颜色")
	orangeRed := f.co.Font().Color().RGB(255, 69, 0)
	f.co.Font().SetColor(orangeRed)
}

//---------------------------------------------------------------------------
func (f *TMainForm) btnZiOnClick(sender vcl.IObject) {

}

func (f *TMainForm) btnHaiOnClick(sender vcl.IObject) {
	fmt.Println("btnHai clicked")
}

func (f *TMainForm) btnChouOnClick(sender vcl.IObject) {

}

func (f *TMainForm) btnYinOnClick(sender vcl.IObject) {

}

func (f *TMainForm) btnMaoOnClick(sender vcl.IObject) {

}

func (f *TMainForm) btnChenOnClick(sender vcl.IObject) {

}

func (f *TMainForm) btnSiOnClick(sender vcl.IObject) {

}

func (f *TMainForm) btnWuOnClick(sender vcl.IObject) {

}

func (f *TMainForm) btnWeiOnClick(sender vcl.IObject) {

}

func (f *TMainForm) btnShenOnClick(sender vcl.IObject) {

}

func (f *TMainForm) btnYouOnClick(sender vcl.IObject) {

}

func (f *TMainForm) btnXuOnClick(sender vcl.IObject) {

}
