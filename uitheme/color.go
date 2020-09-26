package uitheme

import "image/color"

type darkTheme struct {
	Primary   color.Color
	Secondary color.Color
	Success   color.Color
	Info      color.Color
	Warning   color.Color
	Danger    color.Color
}
type lightTheme struct {
	Primary   color.Color
	Secondary color.Color
	Success   color.Color
	Info      color.Color
	Warning   color.Color
	Danger    color.Color
}
var (
	Dark = darkTheme{
		Primary:   &color.RGBA{R: 0x0d, G: 0X47, B: 0xa1, A: 255}, // 主要	#0d47a1
		Secondary: &color.RGBA{R: 0x99, G: 0X33, B: 0xCC, A: 255}, // 次要	#9933CC
		Success:   &color.RGBA{R: 0x00, G: 0X7E, B: 0x33, A: 255}, // 成功	#007E33
		Info:      &color.RGBA{R: 0x00, G: 0X99, B: 0xCC, A: 255}, // 信息	#0099CC
		Warning:   &color.RGBA{R: 0xFF, G: 0X88, B: 0x00, A: 255}, // 警告	#FF8800
		Danger:    &color.RGBA{R: 0xCC, G: 0X00, B: 0x00, A: 255}, // 危险	#CC0000
	}
	Light = lightTheme{
		Primary:   &color.RGBA{R: 0x0d, G: 0x47, B: 0xa1, A: 255}, // 主要	#0d47a1
		Secondary: &color.RGBA{R: 0x99, G: 0x33, B: 0xCC, A: 255}, // 次要	#9933CC
		Success:   &color.RGBA{R: 0x00, G: 0x7E, B: 0x33, A: 255}, // 成功	#007E33
		Info:      &color.RGBA{R: 0x00, G: 0x99, B: 0xCC, A: 255}, // 信息	#0099CC
		Warning:   &color.RGBA{R: 0xFF, G: 0x88, B: 0x00, A: 255}, // 警告	#FF8800
		Danger:    &color.RGBA{R: 0xCC, G: 0x00, B: 0x00, A: 255}, // 危险	#CC0000
	}
)

var (
	ColorBlue   = &color.RGBA{R: 0x00, G: 0x7b, B: 0xff, A: 255} // 蓝	#007bff
	ColorCyan   = &color.RGBA{R: 0x17, G: 0xa2, B: 0xb8, A: 255} // 青	#17a2b8
	ColorDark   = &color.RGBA{R: 0x34, G: 0x3a, B: 0x40, A: 255} // 黑	#343a40
	ColorGray   = &color.RGBA{R: 0x6c, G: 0x75, B: 0x7d, A: 255} // 灰	#6c757d
	ColorGreen  = &color.RGBA{R: 0x28, G: 0xa7, B: 0x45, A: 255} // 绿	#28a745
	ColorIndigo = &color.RGBA{R: 0x66, G: 0x10, B: 0xf2, A: 255} // 靛蓝	#6610f2
	ColorOrange = &color.RGBA{R: 0xfd, G: 0x7e, B: 0x14, A: 255} // 橙	#fd7e14
	ColorPink   = &color.RGBA{R: 0xe8, G: 0x3e, B: 0x8c, A: 255} // 粉	#e83e8c
	ColorPurple = &color.RGBA{R: 0x6f, G: 0x42, B: 0xc1, A: 255} // 紫	#6f42c1
	ColorRed    = &color.RGBA{R: 0xdc, G: 0x35, B: 0x45, A: 255} // 红	#dc3545
	ColorTeal   = &color.RGBA{R: 0x20, G: 0xc9, B: 0x97, A: 255} // 蓝绿	#20c997
	ColorWhite  = &color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 255}  // 白	#fff
	ColorYellow = &color.RGBA{R: 0xff, G: 0xc1, B: 0x07, A: 255} // 黄	#ffc107
	ColorBlack = &color.RGBA{R: 0, G: 0, B: 0, A: 255} // 黄	#000

	ColorPrimary   = &color.RGBA{R: 0x00, G: 0x7b, B: 0xff, A: 255} // 主要	#007bff
	ColorSecondary = &color.RGBA{R: 0x6c, G: 0x75, B: 0x7d, A: 255} // 次要	#6c757d
	ColorSuccess   = &color.RGBA{R: 0x28, G: 0xa7, B: 0x45, A: 255} // 成功	#28a745
	ColorInfo      = &color.RGBA{R: 0x17, G: 0xa2, B: 0xb8, A: 255} // 信息	#17a2b8
	ColorWarning   = &color.RGBA{R: 0xff, G: 0xc1, B: 0x07, A: 255} // 警告	#ffc107
	ColorDanger    = &color.RGBA{R: 0xdc, G: 0x35, B: 0x45, A: 255} // 危险	#dc3545
	ColorLight     = &color.RGBA{R: 0xf8, G: 0xf9, B: 0xfa, A: 255} // 浅色	#f8f9fa

	//ColorPrimaryLight   = &color.RGBA{R: 0xf8, G: 0Xf9, B: 0xfa, A: 255} // 主要	#4285F4
	//ColorSecondaryLight = &color.RGBA{R: 0xf8, G: 0Xf9, B: 0xfa, A: 255} // 次要	#aa66cc
	//ColorSuccessLight   = &color.RGBA{R: 0xf8, G: 0Xf9, B: 0xfa, A: 255} // 成功	#00C851
	//ColorInfoLight      = &color.RGBA{R: 0xf8, G: 0Xf9, B: 0xfa, A: 255} // 信息	#33b5e5
	//ColorWarningLight   = &color.RGBA{R: 0xf8, G: 0Xf9, B: 0xfa, A: 255} // 警告	#ffbb33
	//ColorDangerLight    = &color.RGBA{R: 0xf8, G: 0Xf9, B: 0xfa, A: 255} // 危险	#ff4444
	//
	//ColorPrimaryDark   = &color.RGBA{R: 0xf8, G: 0Xf9, B: 0xfa, A: 255} // 主要	#0d47a1
	//ColorSecondaryDark = &color.RGBA{R: 0xf8, G: 0Xf9, B: 0xfa, A: 255} // 次要	#9933CC
	//ColorSuccessDark   = &color.RGBA{R: 0xf8, G: 0Xf9, B: 0xfa, A: 255} // 成功	#007E33
	//ColorInfoDark      = &color.RGBA{R: 0xf8, G: 0Xf9, B: 0xfa, A: 255} // 信息	#0099CC
	//ColorWarningDark   = &color.RGBA{R: 0xf8, G: 0Xf9, B: 0xfa, A: 255} // 警告	#FF8800
	//ColorDangerDark    = &color.RGBA{R: 0xf8, G: 0Xf9, B: 0xfa, A: 255} // 危险	#CC0000
	//
	//ColorDefaultLight = &color.RGBA{R: 0xf8, G: 0Xf9, B: 0xfa, A: 255} // 错误	#2BBBAD
	//ColorDefaultDark  = &color.RGBA{R: 0xf8, G: 0Xf9, B: 0xfa, A: 255} // 错误	#00695c
)
