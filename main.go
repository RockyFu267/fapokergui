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
	// 记录选中的牌，key 是 "A♠"
	selectedCards := make(map[string]bool)

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
		println("clickableRectflop01 被点击了！")
		clickableRectflop01 := widget.NewButton("", nil)
		var popup *widget.PopUp
		println("新长方形被点击了！")

		// 生成所有扑克牌选项，并根据状态变灰
		cardButtons := []fyne.CanvasObject{}
		pokerRanks := []string{"A", "K", "Q", "J", "10", "9", "8", "7", "6", "5", "4", "3", "2"}
		suits := []string{"♠", "♥", "♣", "♦"}

		for _, rank := range pokerRanks {
			for _, suit := range suits {
				cardText := rank + suit
				cardButton := widget.NewButton(cardText, func() {
					// 如果按钮上已有牌，则先释放它
					if oldCard := clickableRectflop01.Text; oldCard != "" {
						delete(selectedCards, oldCard)
					}

					clickableRectflop01.SetText(cardText) // 选择后更新按钮文本
					selectedCards[cardText] = true        // 标记该牌已被选中
					popup.Hide()                          // 关闭弹窗
				})

				if selectedCards[cardText] {
					cardButton.Disable() // 变灰并禁用
				}

				cardButton.Importance = widget.HighImportance
				cardButtons = append(cardButtons, cardButton)
			}
		}

		// 创建 4 列网格布局
		cardGrid := container.NewGridWithColumns(4, cardButtons...)
		popup = widget.NewModalPopUp(cardGrid, w.Canvas())
		popup.Show()
		clickableRectflop01.Resize(fyne.NewSize(50, 90))
		clickableRectflop01.Move(fyne.NewPos(0, 30))
		clickableRectflop01.Importance = widget.HighImportance
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
	// 将按钮移动到左上角 (0, 30) 的位置
	clickableRectflop01.Move(fyne.NewPos(0, 30))
	// 将按钮移动到左上角 (55, 30) 的位置
	clickableRectflop02.Move(fyne.NewPos(55, 30))
	// 将按钮移动到左上角 (110, 30) 的位置
	clickableRectflop03.Move(fyne.NewPos(110, 30))
	// 将按钮移动到左上角 (165, 30) 的位置
	clickableRectflop04.Move(fyne.NewPos(165, 30))
	// 将按钮移动到左上角 (220, 30) 的位置
	clickableRectflop05.Move(fyne.NewPos(220, 30))

	// 记录新增按钮的行数
	newRowCount := 0

	// 维护所有行的数据结构
	var allRows [][]fyne.CanvasObject

	// 创建 "+add" 按钮
	addButton := widget.NewButton("+add", func() {
		println("+add 按钮被点击了！")

		rowButtons := []fyne.CanvasObject{}

		for i := 0; i < 2; i++ {
			newButton := widget.NewButton("", nil)
			newX := float32(i * 55)
			newY := float32(30 + 90 + 10 + newRowCount*(90+10))

			var popup *widget.PopUp

			newButton.OnTapped = func() {
				println("新长方形被点击了！")

				// 生成所有扑克牌选项，并根据状态变灰
				cardButtons := []fyne.CanvasObject{}
				pokerRanks := []string{"A", "K", "Q", "J", "10", "9", "8", "7", "6", "5", "4", "3", "2"}
				suits := []string{"♠", "♥", "♣", "♦"}

				for _, rank := range pokerRanks {
					for _, suit := range suits {
						cardText := rank + suit
						cardButton := widget.NewButton(cardText, func() {
							// 如果按钮上已有牌，则先释放它
							if oldCard := newButton.Text; oldCard != "" {
								delete(selectedCards, oldCard)
							}

							newButton.SetText(cardText)    // 选择后更新按钮文本
							selectedCards[cardText] = true // 标记该牌已被选中
							popup.Hide()                   // 关闭弹窗
						})

						if selectedCards[cardText] {
							cardButton.Disable() // 变灰并禁用
						}

						cardButton.Importance = widget.HighImportance
						cardButtons = append(cardButtons, cardButton)
					}
				}

				// 创建 4 列网格布局
				cardGrid := container.NewGridWithColumns(4, cardButtons...)
				popup = widget.NewModalPopUp(cardGrid, w.Canvas())
				popup.Show()
			}

			newButton.Resize(fyne.NewSize(50, 90))
			newButton.Move(fyne.NewPos(newX, newY))
			newButton.Importance = widget.HighImportance

			positionContainer.Add(newButton)

			rowButtons = append(rowButtons, newButton)
		}

		deleteButton := widget.NewButton("del", func() {
			rowIndex := -1
			for i, row := range allRows {
				if len(row) > 0 && row[0] == rowButtons[0] {
					rowIndex = i
					break
				}
			}

			if rowIndex != -1 {
				for _, btn := range allRows[rowIndex] {
					if textBtn, ok := btn.(*widget.Button); ok {
						if textBtn.Text != "" {
							delete(selectedCards, textBtn.Text) // 释放牌
						}
					}
					positionContainer.Remove(btn)
				}
				allRows = append(allRows[:rowIndex], allRows[rowIndex+1:]...)

				for i := rowIndex; i < len(allRows); i++ {
					for j, btn := range allRows[i] {
						newX := float32(j * 55)
						newY := float32(30 + 90 + 10 + i*(90+10))
						btn.Move(fyne.NewPos(newX, newY))
					}
				}

				newRowCount--
				positionContainer.Refresh()
			}
		})

		deleteButton.Resize(fyne.NewSize(25, 15))
		deleteButton.Move(fyne.NewPos(110, 30+90+10+float32(newRowCount*(90+10))))
		positionContainer.Add(deleteButton)

		rowButtons = append(rowButtons, deleteButton)
		allRows = append(allRows, rowButtons)
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
