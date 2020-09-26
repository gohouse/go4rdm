package fmcolor

import "image/color"

type symbolicColor struct {
	Primary   *color.RGBA
	Secondary *color.RGBA
	Success   *color.RGBA
	Info      *color.RGBA
	Warning   *color.RGBA
	Danger    *color.RGBA
	Light     *color.RGBA
}

type namedColor struct {
	Blue   *color.RGBA
	Cyan   *color.RGBA
	Dark   *color.RGBA
	Gray   *color.RGBA
	Green  *color.RGBA
	Indigo *color.RGBA
	Orange *color.RGBA
	Pink   *color.RGBA
	Purple *color.RGBA
	Red    *color.RGBA
	Teal   *color.RGBA
	White  *color.RGBA
	Yellow *color.RGBA
	Black  *color.RGBA
}

type fmColor struct {
	symbolicColor
	namedColor
}

var (
	Dark = symbolicColor{
		Primary:   &color.RGBA{R: 0x0d, G: 0X47, B: 0xa1, A: 255}, // 主要	#0d47a1
		Secondary: &color.RGBA{R: 0x99, G: 0X33, B: 0xCC, A: 255}, // 次要	#9933CC
		Success:   &color.RGBA{R: 0x00, G: 0X7E, B: 0x33, A: 255}, // 成功	#007E33
		Info:      &color.RGBA{R: 0x00, G: 0X99, B: 0xCC, A: 255}, // 信息	#0099CC
		Warning:   &color.RGBA{R: 0xFF, G: 0X88, B: 0x00, A: 255}, // 警告	#FF8800
		Danger:    &color.RGBA{R: 0xCC, G: 0X00, B: 0x00, A: 255}, // 危险	#CC0000
	}
	Light = symbolicColor{
		Primary:   &color.RGBA{R: 0x42, G: 0x85, B: 0xF4, A: 255}, // 主要	#4285F4
		Secondary: &color.RGBA{R: 0xaa, G: 0x66, B: 0xcc, A: 255}, // 次要	#aa66cc
		Success:   &color.RGBA{R: 0x00, G: 0xC8, B: 0x51, A: 255}, // 成功	#00C851
		Info:      &color.RGBA{R: 0x33, G: 0xb5, B: 0xe5, A: 255}, // 信息	#33b5e5
		Warning:   &color.RGBA{R: 0xff, G: 0xbb, B: 0x33, A: 255}, // 警告	#ffbb33
		Danger:    &color.RGBA{R: 0xff, G: 0x44, B: 0x44, A: 255}, // 危险	#ff4444
	}
	Normal = fmColor{
		symbolicColor: symbolicColor{
			Primary:   &color.RGBA{R: 0x00, G: 0x7b, B: 0xff, A: 255}, // 主要	#007bff
			Secondary: &color.RGBA{R: 0x6c, G: 0x75, B: 0x7d, A: 255}, // 次要	#6c757d
			Success:   &color.RGBA{R: 0x28, G: 0xa7, B: 0x45, A: 255}, // 成功	#28a745
			Info:      &color.RGBA{R: 0x17, G: 0xa2, B: 0xb8, A: 255}, // 信息	#17a2b8
			Warning:   &color.RGBA{R: 0xff, G: 0xc1, B: 0x07, A: 255}, // 警告	#ffc107
			Danger:    &color.RGBA{R: 0xdc, G: 0x35, B: 0x45, A: 255}, // 危险	#dc3545
			Light:     &color.RGBA{R: 0xf8, G: 0xf9, B: 0xfa, A: 255}, // 浅色	#f8f9fa
		},
		namedColor: namedColor{
			Blue:   &color.RGBA{R: 0x00, G: 0x7b, B: 0xff, A: 255}, // 蓝	#007bff
			Cyan:   &color.RGBA{R: 0x17, G: 0xa2, B: 0xb8, A: 255}, // 青	#17a2b8
			Dark:   &color.RGBA{R: 0x34, G: 0x3a, B: 0x40, A: 255}, // 黑	#343a40
			Gray:   &color.RGBA{R: 0x6c, G: 0x75, B: 0x7d, A: 255}, // 灰	#6c757d
			Green:  &color.RGBA{R: 0x28, G: 0xa7, B: 0x45, A: 255}, // 绿	#28a745
			Indigo: &color.RGBA{R: 0x66, G: 0x10, B: 0xf2, A: 255}, // 靛蓝	#6610f2
			Orange: &color.RGBA{R: 0xfd, G: 0x7e, B: 0x14, A: 255}, // 橙	#fd7e14
			Pink:   &color.RGBA{R: 0xe8, G: 0x3e, B: 0x8c, A: 255}, // 粉	#e83e8c
			Purple: &color.RGBA{R: 0x6f, G: 0x42, B: 0xc1, A: 255}, // 紫	#6f42c1
			Red:    &color.RGBA{R: 0xdc, G: 0x35, B: 0x45, A: 255}, // 红	#dc3545
			Teal:   &color.RGBA{R: 0x20, G: 0xc9, B: 0x97, A: 255}, // 蓝绿	#20c997
			White:  &color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 255}, // 白	#fff
			Yellow: &color.RGBA{R: 0xff, G: 0xc1, B: 0x07, A: 255}, // 黄	#ffc107
			Black:  &color.RGBA{R: 0, G: 0, B: 0, A: 255},          // 黄	#000
		},
	}
)
