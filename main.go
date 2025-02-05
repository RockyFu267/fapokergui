package main

import (
	"fmt"
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

	// 创建一个容器用于绝对定位
	positionContainer := container.NewWithoutLayout()
	// 预定义按钮的位置
	flopPositions := []fyne.Position{
		fyne.NewPos(0, 30),
		fyne.NewPos(50, 30),
		fyne.NewPos(100, 30),
		fyne.NewPos(150, 30),
		fyne.NewPos(200, 30),
	}

	var flopButtons []*widget.Button

	// 创建清除公共牌按钮
	clearFlopButton := widget.NewButton("清除公共牌", func() {
		for _, btn := range flopButtons {
			if oldCard := btn.Text; oldCard != "?" {
				delete(selectedCards, oldCard) // 释放已选牌
			}
			btn.SetText("?") // 恢复默认状态
		}
	})

	// 调整清除按钮位置和大小
	clearFlopButton.Resize(fyne.NewSize(80, 20))
	clearFlopButton.Move(fyne.NewPos(250, 60)) // 你可以调整这个位置

	// 将清除按钮添加到绝对定位容器中
	positionContainer.Add(clearFlopButton)

	for _, pos := range flopPositions {
		btn := widget.NewButton("?", nil) // 先创建按钮，不在这里直接绑定事件

		btn.OnTapped = func(b *widget.Button) func() {
			return func() {
				var popup *widget.PopUp

				println("公共牌按钮被点击！")

				cardButtons := []fyne.CanvasObject{}
				pokerRanks := []string{"A", "K", "Q", "J", "10", "9", "8", "7", "6", "5", "4", "3", "2"}
				suits := []string{"♠", "♥", "♣", "♦"}

				for _, rank := range pokerRanks {
					for _, suit := range suits {
						cardText := suit + rank
						cardButton := widget.NewButton(cardText, func() {
							// 如果按钮上已有牌，先释放它
							if oldCard := b.Text; oldCard != "?" {
								delete(selectedCards, oldCard)
							}

							b.SetText(cardText)            // 更新按钮文本
							selectedCards[cardText] = true // 标记该牌已被选中
							popup.Hide()                   // 关闭弹窗
						})

						if selectedCards[cardText] {
							cardButton.Disable() // 已选的牌变灰
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
		}(btn) // 这里传递 `btn`，确保作用域正确

		btn.Resize(fyne.NewSize(45, 80))
		btn.Move(pos)
		flopButtons = append(flopButtons, btn)
		positionContainer.Add(btn)
	}

	// 记录新增按钮的行数
	newRowCount := 0

	// 维护所有行的数据结构
	var allRows [][]fyne.CanvasObject

	var addButton *widget.Button // 先声明变量
	// 创建 "+add" 按钮
	addButton = widget.NewButton("+add", func() {
		if newRowCount >= 10 {
			println("行数已达上限，直接返回")
			addButton.Disable() // 禁用 "+add" 按钮
			return              // 如果行数已达上限，直接返回
		}

		println("+add 按钮被点击了！")

		rowButtons := []fyne.CanvasObject{}

		for i := 0; i < 2; i++ {
			newButton := widget.NewButton("?", nil)
			newX := float32(i * 55)
			newY := float32(30 + 90 + 10 + newRowCount*(65))

			var popup *widget.PopUp

			newButton.OnTapped = func() {
				println("新长方形被点击了！")

				// 生成所有扑克牌选项，并根据状态变灰
				cardButtons := []fyne.CanvasObject{}
				pokerRanks := []string{"A", "K", "Q", "J", "10", "9", "8", "7", "6", "5", "4", "3", "2"}
				suits := []string{"♠", "♥", "♣", "♦"}

				for _, rank := range pokerRanks {
					for _, suit := range suits {
						cardText := suit + rank
						cardButton := widget.NewButton(cardText, func() {
							// 如果按钮上已有牌，则先释放它
							if oldCard := newButton.Text; oldCard != "?" {
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

			newButton.Resize(fyne.NewSize(30, 50))
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
						newY := float32(30 + 90 + 10 + i*(65))
						btn.Move(fyne.NewPos(newX, newY))
					}
				}

				newRowCount--
				if newRowCount < 10 {
					addButton.Enable() // 重新启用 "+add" 按钮
				}
				positionContainer.Refresh()
			}
		})

		deleteButton.Resize(fyne.NewSize(25, 15))
		deleteButton.Move(fyne.NewPos(110, 30+90+10+float32(newRowCount*(65))))
		positionContainer.Add(deleteButton)

		rowButtons = append(rowButtons, deleteButton)
		allRows = append(allRows, rowButtons)
		newRowCount++
		positionContainer.Refresh()
	})

	// 将 "+add" 按钮添加到绝对定位容器中
	positionContainer.Add(addButton)

	// 设置按钮大小
	addButton.Resize(fyne.NewSize(40, 20))
	// 设置按钮位置，放置在现有按钮下方
	addButton.Move(fyne.NewPos(280, 30))

	// 创建执行运算按钮
	executeButton := widget.NewButton("GO", func() {
		// 在这里执行你的运算逻辑
		println("执行运算按钮被点击！")
		// 在这里收集所有手牌和公共牌数据
		var playerHands []string
		for _, row := range allRows {
			for _, btn := range row {
				if textBtn, ok := btn.(*widget.Button); ok && textBtn.Text != "?" {
					playerHands = append(playerHands, textBtn.Text)
				}
			}
		}

		var flopCards []string
		for _, btn := range flopButtons {
			if btn.Text != "?" {
				flopCards = append(flopCards, btn.Text)
			}
		}

		// 打印手牌和公共牌
		fmt.Println("玩家手牌: ", playerHands)
		fmt.Println("公共牌: ", flopCards)

		// 这里可以将收集到的手牌和公共牌数据传入胜率计算逻辑
		// 调用你自己实现的胜率计算函数
		// calculateWinRate(playerHands, flopCards)
	})
	// 设置按钮大小
	executeButton.Resize(fyne.NewSize(80, 20))
	// 设置按钮位置，放置在现有按钮下方
	executeButton.Move(fyne.NewPos(250, 90))
	// 将执行运算按钮添加到绝对定位容器中
	positionContainer.Add(executeButton)

	// 使用 NewPadded 容器将绝对定位容器居中显示在窗口中
	content := container.NewPadded(positionContainer)

	// 设置窗口的内容为包含按钮的容器
	w.SetContent(content)

	// 显示窗口并运行应用程序
	w.ShowAndRun()
}
