package main

import (
	"github.com/ying32/govcl/vcl"
)

type TMainForm struct {
	*vcl.TForm

	pbox                                                                                          *vcl.TPaintBox
	input                                                                                         *vcl.TEdit //获取输入日期　年四位数　月　日　时　各两位数
	BtnOK                                                                                         *vcl.TButton
	btnHai, btnZi, btnChou, btnYin, btnMao, btnChen, btnSi, btnWu, btnWei, btnShen, btnYou, btnXu *vcl.TButton //寅首　十二长生

	haiSiHua, ziSiHua, chouSiHua, yinSiHua, maoSiHua, chenSiHua *vcl.TLabel //四化
	siSiHua, wuSiHua, weiSiHua, shenSiHua, youSiHua, xuSiHua    *vcl.TLabel //四化

	haiMg, ziMg, chouMg, yinMg, maoMg, chenMg, siMg, wuMg, weiMg, shenMg, youMg, xuMg                         *vcl.TLabel //十二命宫
	haiStar, ziStar, chouStar, yinStar, maoStar, chenStar, siStar, wuStar, weiStar, shenStar, youStar, xuStar *vcl.TLabel //十四主星

	haiYearGanStar, ziYearGanStar, chouYearGanStar, yinYearGanStar, maoYearGanStar, chenYearGanStar *vcl.TLabel //年干系诸星
	siYearGanStar, wuYearGanStar, weiYearGanStar, shenYearGanStar, youYearGanStar, xuYearGanStar    *vcl.TLabel //年干系诸星

	haiYearZhiStar, ziYearZhiStar, chouYearZhiStar, yinYearZhiStar, maoYearZhiStar, chenYearZhiStar *vcl.TLabel //年支系诸星
	siYearZhiStar, wuYearZhiStar, weiYearZhiStar, shenYearZhiStar, youYearZhiStar, xuYearZhiStar    *vcl.TLabel //年支系诸星

	haiMonthStar, ziMonthStar, chouMonthStar, yinMonthStar, maoMonthStar, chenMonthStar *vcl.TLabel //月系诸星
	siMonthStar, wuMonthStar, weiMonthStar, shenMonthStar, youMonthStar, xuMonthStar    *vcl.TLabel //月系诸星

	haiHourStar, ziHourStar, chouHourStar, yinHourStar, maoHourStar, chenHourStar *vcl.TLabel //时系诸星
	siHourStar, wuHourStar, weiHourStar, shenHourStar, youHourStar, xuHourStar    *vcl.TLabel //时系诸星

	haiOtherStar, ziOtherStar, chouOtherStar, yinOtherStar, maoOtherStar, chenOtherStar *vcl.TLabel //其他系诸星
	siOtherStar, wuOtherStar, weiOtherStar, shenOtherStar, youOtherStar, xuOtherStar    *vcl.TLabel //其他系诸星

	haiDaXian, ziDaXian, chouDaXian, yinDaXian, maoDaXian, chenDaXian *vcl.TLabel //大限
	siDaXian, wuDaXian, weiDaXian, shenDaXian, youDaXian, xuDaXian    *vcl.TLabel //大限

	haiXiaoXian, ziXiaoXian, chouXiaoXian, yinXiaoXian, maoXiaoXian, chenXiaoXian *vcl.TLabel //小限
	siXiaoXian, wuXiaoXian, weiXiaoXian, shenXiaoXian, youXiaoXian, xuXiaoXian    *vcl.TLabel //小限

	//中心区域
	label1 *vcl.TLabel //纪年信息　化禄化权
	label2 *vcl.TLabel
	label3 *vcl.TLabel
	label4 *vcl.TLabel
	//颜色说明
	c14 *vcl.TLabel //14主星 colors.ClRed
	cyg *vcl.TLabel //年干系诸星 blue RGB(72, 118, 255)
	cyz *vcl.TLabel //年支系诸星 skyBlue RGB(135, 206, 255)
	cm  *vcl.TLabel //月系诸星 limeGreen RGB(50, 205, 50)
	ch  *vcl.TLabel //时系诸星 orchid RGB(218, 112, 214)
	co  *vcl.TLabel //其他星	orangeRed RGB(255, 69, 0)
	//关于
	about *vcl.TLabel
}

var (
	mainForm *TMainForm
)

const (
	top       = 10
	width     = 60
	height    = 120
	leftInput = 130
	leftok    = 235
)

//画
const (
	//巳宫
	l1 = 10
	r1 = 220
	//未宫
	l2 = 2 * r1
	r2 = 3 * r1
	b1 = 215
	//卯宫
	t3 = top + 2*b1
	b3 = 3 * b1
	//丑宫
	t4 = top + 3*b1
	l4 = r1
	r4 = l4 + b1
	b4 = 4 * b1
	//亥宫
	l5 = r2
	r5 = l5 + b1
	//酉宫
	t6 = b1
	b6 = t3
)
