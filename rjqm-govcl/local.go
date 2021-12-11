/*
 * The code is automatically generated by the Goland.
 * Copyright © Aquarian-Age. All Rights Reserved.
 * Licensed under MIT.
 */

package main

import "github.com/ying32/govcl/vcl"

const (
	topl1  = 140
	topl2  = 2 * topl1
	topl3  = 3 * topl1
	leftl1 = 30
	leftl2 = 160
	leftl3 = 290
)

func (f *TMainForm) local() {
	f.local1 = vcl.NewLabel(f)
	f.local1.SetParent(f)
	f.local1.SetTop(topl3)
	f.local1.SetLeft(leftl2)
	f.local1.SetCaption(`子 坎`)

	f.local8 = vcl.NewLabel(f)
	f.local8.SetParent(f)
	f.local8.SetTop(topl3 - 20)
	f.local8.SetLeft(leftl1)
	l8 := `寅 艮
丑`
	f.local8.SetCaption(l8)

	f.local3 = vcl.NewLabel(f)
	f.local3.SetParent(f)
	f.local3.SetTop(topl2)
	f.local3.SetLeft(leftl1)
	l3 := `卯 震`
	f.local3.SetCaption(l3)

	f.local4 = vcl.NewLabel(f)
	f.local4.SetParent(f)
	f.local4.SetTop(topl1 - 30)
	f.local4.SetLeft(leftl1)
	l4 := `巳 巽
辰`
	f.local4.SetCaption(l4)

	f.local9 = vcl.NewLabel(f)
	f.local9.SetParent(f)
	f.local9.SetTop(topl1)
	f.local9.SetLeft(leftl2)
	f.local9.SetCaption(`午 离`)

	f.local2 = vcl.NewLabel(f)
	f.local2.SetParent(f)
	f.local2.SetTop(topl1 - 30)
	f.local2.SetLeft(leftl3)
	l2 := `未 坤
申`
	f.local2.SetCaption(l2)

	f.local7 = vcl.NewLabel(f)
	f.local7.SetParent(f)
	f.local7.SetTop(topl2)
	f.local7.SetLeft(leftl3)
	f.local7.SetCaption(`酉 兑`)

	f.local6 = vcl.NewLabel(f)
	f.local6.SetParent(f)
	f.local6.SetTop(topl3 - 20)
	f.local6.SetLeft(leftl3)
	l6 := `戌 乾
亥`
	f.local6.SetCaption(l6)
}