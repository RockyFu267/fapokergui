package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// 自定义主题，用于修改按钮背景颜色
type customTheme struct{}

// Color 实现 fyne.Theme 接口的 Color 方法
func (c customTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	if name == theme.ColorNameButton && variant == theme.VariantLight {
		// 设置按钮背景颜色为金色
		return color.RGBA{R: 255, G: 215, B: 0, A: 255}
	}
	return theme.DefaultTheme().Color(name, variant)
}

// Font 实现 fyne.Theme 接口的 Font 方法
func (c customTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

// Icon 实现 fyne.Theme 接口的 Icon 方法
func (c customTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

// Size 实现 fyne.Theme 接口的 Size 方法
func (c customTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}

func main() {
	// 创建一个新的 Fyne 应用
	a := app.New()
	// 设置自定义主题
	a.Settings().SetTheme(customTheme{})

	// 创建一个新的窗口，窗口标题为 "FA - poker"
	w := a.NewWindow("FA-poker")

	// 设置窗口的大小为自定义分辨率
	w.Resize(fyne.NewSize(350, 780))

	// 创建一个可点击的按钮来模拟长方形
	clickableRectflop01 := widget.NewButton("", func() {
		// 这里可以添加按钮点击后的处理逻辑
		// 例如打印日志，当前仅作示例
		println("长方形被点击了！")
	})

	// 设置按钮的大小为 50x90
	clickableRectflop01.Resize(fyne.NewSize(50, 90))

	// 创建一个容器用于绝对定位
	positionContainer := container.NewWithoutLayout()
	positionContainer.Add(clickableRectflop01)
	// 将按钮移动到左上角 (0, 0) 的位置
	clickableRectflop01.Move(fyne.NewPos(0, 0))

	// 使用 NewPadded 容器将绝对定位容器居中显示在窗口中
	content := container.NewPadded(positionContainer)

	// 设置窗口的内容为包含按钮的容器
	w.SetContent(content)

	// 显示窗口并运行应用程序
	w.ShowAndRun()
}
