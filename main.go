package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// Card 定义扑克牌结构体
type Card struct {
	Suit string
	Rank string
}

// generateDeck 生成一副 52 张扑克牌
func generateDeck() []Card {
	suits := []string{"♥", "♦", "♣", "♠"}
	ranks := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	deck := make([]Card, 0, 52)
	for _, suit := range suits {
		for _, rank := range ranks {
			deck = append(deck, Card{Suit: suit, Rank: rank})
		}
	}
	return deck
}

// 自定义主题，用于修改按钮背景颜色
type customTheme struct{}

// Color 实现 fyne.Theme 接口的 Color 方法
func (c customTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	if name == theme.ColorNameButton && variant == theme.VariantLight {
		// 设置按钮背景颜色为金色
		return color.RGBA{R: 255, G: 215, B: 0, A: 255}
	} else if name == theme.ColorNameBackground {
		// 设置窗口背景颜色为灰色
		return color.RGBA{R: 200, G: 200, B: 200, A: 255}
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
	// 生成一副扑克牌
	deck := generateDeck()
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
	// 创建一个可点击的按钮来模拟长方形
	clickableRectflop02 := widget.NewButton("", func() {
		// 这里可以添加按钮点击后的处理逻辑
		// 例如打印日志，当前仅作示例
		println("长方形被点击了！")
	})
	// 创建一个可点击的按钮来模拟长方形
	clickableRectflop03 := widget.NewButton("", func() {
		// 这里可以添加按钮点击后的处理逻辑
		// 例如打印日志，当前仅作示例
		println("长方形被点击了！")
	})
	// 创建一个可点击的按钮来模拟长方形
	clickableRectflop04 := widget.NewButton("", func() {
		// 这里可以添加按钮点击后的处理逻辑
		// 例如打印日志，当前仅作示例
		println("长方形被点击了！")
	})
	// 创建一个可点击的按钮来模拟长方形
	clickableRectflop05 := widget.NewButton("", func() {
		// 这里可以添加按钮点击后的处理逻辑
		// 例如打印日志，当前仅作示例
		println("长方形被点击了！")
	})
	// 设置按钮的大小为 50x90
	clickableRectflop01.Resize(fyne.NewSize(50, 90))
	// 设置按钮的大小为 50x90
	clickableRectflop02.Resize(fyne.NewSize(50, 90))
	// 设置按钮的大小为 50x90
	clickableRectflop03.Resize(fyne.NewSize(50, 90))
	// 设置按钮的大小为 50x90
	clickableRectflop04.Resize(fyne.NewSize(50, 90))
	// 设置按钮的大小为 50x90
	clickableRectflop05.Resize(fyne.NewSize(50, 90))

	// 创建一个容器用于绝对定位
	positionContainer := container.NewWithoutLayout()
	positionContainer.Add(clickableRectflop01)
	positionContainer.Add(clickableRectflop02)
	positionContainer.Add(clickableRectflop03)
	positionContainer.Add(clickableRectflop04)
	positionContainer.Add(clickableRectflop05)
	// 将按钮移动到左上角 (0, 0) 的位置
	clickableRectflop01.Move(fyne.NewPos(0, 30))
	// 将按钮移动到左上角 (55, 0) 的位置
	clickableRectflop02.Move(fyne.NewPos(55, 30))
	// 将按钮移动到左上角 (110, 0) 的位置
	clickableRectflop03.Move(fyne.NewPos(110, 30))
	// 将按钮移动到左上角 (165, 0) 的位置
	clickableRectflop04.Move(fyne.NewPos(165, 30))
	// 将按钮移动到左上角 (220, 0) 的位置
	clickableRectflop05.Move(fyne.NewPos(220, 30))

	// 记录新增按钮的行数
	newRowCount := 0
	// 创建 "+add" 按钮
	addButton := widget.NewButton("+add", func() {
		println("+add 按钮被点击了！")

		for i := 0; i < 2; i++ {
			newButton := widget.NewButton("", nil) // 先创建按钮
			newX := float32(i * 55)
			newY := float32(30 + 90 + 10 + newRowCount*(90+10))

			newButton.OnTapped = func() {
				println("新长方形被点击了！")

				// 生成扑克牌选项
				cardOptions := make([]string, len(deck))
				for j, card := range deck {
					cardOptions[j] = card.Suit + card.Rank
				}

				// 先声明 selectCard 变量
				var selectCard *widget.Select
				selectCard = widget.NewSelect(cardOptions, func(selected string) {
					newButton.SetText(selected)          // 更新按钮文本
					positionContainer.Remove(selectCard) // 选择完后移除下拉菜单
					positionContainer.Refresh()
				})

				selectCard.Resize(fyne.NewSize(50, 30))
				selectCard.Move(fyne.NewPos(newX, newY+100)) // 让下拉菜单显示在按钮下方

				positionContainer.Add(selectCard)
				positionContainer.Refresh()
			}

			newButton.Resize(fyne.NewSize(50, 90))
			newButton.Move(fyne.NewPos(newX, newY))
			newButton.Importance = widget.HighImportance // 让字体填充按钮
			positionContainer.Add(newButton)
		}

		newRowCount++
		positionContainer.Refresh()
	})

	// 设置按钮大小
	addButton.Resize(fyne.NewSize(50, 30))
	// 设置按钮位置，放置在现有按钮下方
	addButton.Move(fyne.NewPos(280, 30))

	// 将 "+add" 按钮添加到绝对定位容器中
	positionContainer.Add(addButton)

	// 使用 NewPadded 容器将绝对定位容器居中显示在窗口中
	content := container.NewPadded(positionContainer)

	// 设置窗口的内容为包含按钮的容器
	w.SetContent(content)

	// 显示窗口并运行应用程序
	w.ShowAndRun()
}
