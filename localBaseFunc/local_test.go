package localbasefunc

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func TestSortCards(t *testing.T) {
	tests := []struct {
		cards    []Card
		expected []Card
	}{
		{
			cards: []Card{
				{Rank: 2, Suit: "梅花"},
				{Rank: 3, Suit: "红桃"},
				{Rank: 2, Suit: "黑桃"},
				{Rank: 4, Suit: "方片"},
			},
			expected: []Card{
				{Rank: 4, Suit: "方片"},
				{Rank: 3, Suit: "红桃"},
				{Rank: 2, Suit: "黑桃"},
				{Rank: 2, Suit: "梅花"},
			},
		},
		{
			cards: []Card{
				{Rank: 5, Suit: "梅花"},
				{Rank: 5, Suit: "黑桃"},
				{Rank: 5, Suit: "红桃"},
				{Rank: 5, Suit: "方片"},
			},
			expected: []Card{
				{Rank: 5, Suit: "黑桃"},
				{Rank: 5, Suit: "红桃"},
				{Rank: 5, Suit: "方片"},
				{Rank: 5, Suit: "梅花"},
			},
		},
		{
			cards: []Card{
				{Rank: 10, Suit: "梅花"},
			},
			expected: []Card{
				{Rank: 10, Suit: "梅花"},
			},
		},
		{
			cards:    []Card{},
			expected: []Card{},
		},
	}

	for _, test := range tests {
		result := sortCards(test.cards)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("sortCards(%v) = %v, want %v", test.cards, result, test.expected)
		}
	}
}

func TestConvertInputToCard_ValidInput_ReturnsCard(t *testing.T) {
	tests := []struct {
		input    string
		expected Card
	}{
		{"♥A", Card{Rank: 14, Suit: "红桃"}},
		{"♦10", Card{Rank: 10, Suit: "方片"}},
		{"♣2", Card{Rank: 2, Suit: "梅花"}},
		{"♠K", Card{Rank: 13, Suit: "黑桃"}},
		{"?", Card{Rank: 0, Suit: "?"}},
		{"", Card{}},
	}

	for _, test := range tests {
		card, err := ConvertInputToCard(test.input)
		if err != nil {

			fmt.Println(test.input, err)
		}
		if card != test.expected {
			t.Errorf("convertInputToCard(%q) = %v, want %v", test.input, card, test.expected)
		}
	}
}

func Test_DealCards(t *testing.T) {

	res := ShuffleCard()
	// fmt.Println(res)
	for _, i := range res {
		cardView := i.CardTranslate()
		fmt.Print(cardView, " ")
	}

	res01, res02 := DealCards(res, 10)
	fmt.Println(len(res01))
	for k, v := range res01 {
		fmt.Println(k, v.HandCard[0].CardTranslate(), v.HandCard[1].CardTranslate())
	}
	fmt.Println("--------")
	for _, v := range res02 {
		fmt.Print(v.CardTranslate(), " ")
	}

	fmt.Println("xxxxxxxxxxxxxxxxxxx")

	inputTest := []Card{
		{Rank: 4, Suit: "方片"},
		{Rank: 4, Suit: "梅花"},
		{Rank: 4, Suit: "黑桃"},
		{Rank: 8, Suit: "方片"},
		{Rank: 7, Suit: "梅花"},
		{Rank: 4, Suit: "红桃"},
		{Rank: 9, Suit: "方片"},
	}
	resA := ShortOfShuffleCard(inputTest)
	for _, i := range resA {
		cardView := i.CardTranslate()
		fmt.Print(cardView, " ")
	}

	resA01, resA02 := DealCards(resA, 5)
	fmt.Println(len(resA01))
	for k, v := range resA01 {
		fmt.Println(k, v.HandCard[0].CardTranslate(), v.HandCard[1].CardTranslate())
	}
	fmt.Println(len(resA), len(resA01), len(resA02))
	fmt.Println("xxxxxxxxxxxxxxxxxxx")
	for _, v := range resA02 {
		fmt.Print(v.CardTranslate(), " ")
	}
}

func TestShortLocalDealCards_EmptyHands_NewHandsAssigned(t *testing.T) {
	// 准备
	// pubCard := []Card{{Suit: "红桃", Rank: 2}, {Suit: "黑桃", Rank: 3}}
	// playerList := []Players{
	// 	{Hand: HandCard{HandCard: []Card{{Suit: "红桃", Rank: 14}, {Suit: "黑桃", Rank: 14}}}},
	// 	{Hand: HandCard{HandCard: []Card{{Rank: 0}, {Rank: 0}}}},
	// }
	pubCard := []Card{{Suit: "？", Rank: 0}, {Suit: "?", Rank: 0}, {Suit: "？", Rank: 0}, {Suit: "?", Rank: 0}, {Suit: "？", Rank: 0}}
	playerList := []Players{
		{Hand: HandCard{HandCard: []Card{{Suit: "?", Rank: 0}, {Suit: "?", Rank: 0}}}},
		{Hand: HandCard{HandCard: []Card{{Suit: "?", Rank: 0}, {Suit: "?", Rank: 0}}}},
	}
	playerList[0].ID = "111"
	playerList[1].ID = "222"

	for i := 0; i < 20; i++ {

		fmt.Println(pubCard, playerList)
		// 执行
		PlayersRes, pubCardRes := ShortLocalDealCards(pubCard, playerList)

		for _, player := range PlayersRes {
			fmt.Println(player.ID + "xxxxxxxxxx")
			fmt.Println(player.Hand.HandCard[0].CardTranslate(), player.Hand.HandCard[1].CardTranslate())

			fmt.Println(pubCardRes[0].CardTranslate(), pubCardRes[1].CardTranslate(), pubCardRes[2].CardTranslate(), pubCardRes[3].CardTranslate(), pubCardRes[4].CardTranslate())
		}
		fmt.Println("--------")
	}

}

func TestSortCards7(t *testing.T) {
	tests := []struct {
		cards    [7]Card
		expected [7]Card
	}{
		{
			cards: [7]Card{
				{Rank: 2, Suit: "梅花"},
				{Rank: 3, Suit: "红桃"},
				{Rank: 4, Suit: "方片"},
				{Rank: 5, Suit: "黑桃"},
				{Rank: 6, Suit: "梅花"},
				{Rank: 7, Suit: "红桃"},
				{Rank: 8, Suit: "方片"},
			},
			expected: [7]Card{
				{Rank: 8, Suit: "方片"},
				{Rank: 7, Suit: "红桃"},
				{Rank: 6, Suit: "梅花"},
				{Rank: 5, Suit: "黑桃"},
				{Rank: 4, Suit: "方片"},
				{Rank: 3, Suit: "红桃"},
				{Rank: 2, Suit: "梅花"},
			},
		},
		{
			cards: [7]Card{
				{Rank: 5, Suit: "梅花"},
				{Rank: 5, Suit: "红桃"},
				{Rank: 5, Suit: "方片"},
				{Rank: 5, Suit: "黑桃"},
				{Rank: 5, Suit: "梅花"},
				{Rank: 5, Suit: "红桃"},
				{Rank: 5, Suit: "方片"},
			},
			expected: [7]Card{
				{Rank: 5, Suit: "黑桃"},
				{Rank: 5, Suit: "红桃"},
				{Rank: 5, Suit: "红桃"},
				{Rank: 5, Suit: "方片"},
				{Rank: 5, Suit: "方片"},
				{Rank: 5, Suit: "梅花"},
				{Rank: 5, Suit: "梅花"},
			},
		},
		{
			cards: [7]Card{
				{Rank: 10, Suit: "梅花"},
				{Rank: 10, Suit: "红桃"},
				{Rank: 10, Suit: "方片"},
				{Rank: 10, Suit: "黑桃"},
				{Rank: 10, Suit: "梅花"},
				{Rank: 10, Suit: "红桃"},
				{Rank: 10, Suit: "方片"},
			},
			expected: [7]Card{
				{Rank: 10, Suit: "黑桃"},
				{Rank: 10, Suit: "红桃"},
				{Rank: 10, Suit: "红桃"},
				{Rank: 10, Suit: "方片"},
				{Rank: 10, Suit: "方片"},
				{Rank: 10, Suit: "梅花"},
				{Rank: 10, Suit: "梅花"},
			},
		},
	}

	for _, test := range tests {
		result := sortCards7(test.cards)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("sortCards7(%v) = %v, want %v", test.cards, result, test.expected)
		}
	}
}
func TestShuffleJudgeGUI01(t *testing.T) {

	// 测试用例1：不同的牌型
	players := []Players{
		// {ID: "AAA", Hand: HandCard{HandCard: []Card{{Rank: 2, Suit: "红桃"}, {Rank: 3, Suit: "红桃"}}}},
		// {ID: "BBB", Hand: HandCard{HandCard: []Card{{Rank: 4, Suit: "红桃"}, {Rank: 5, Suit: "红桃"}}}},
		{ID: "AAA", Hand: HandCard{HandCard: []Card{{Rank: 0, Suit: "?"}, {Rank: 0, Suit: "?"}}}},
		{ID: "BBB", Hand: HandCard{HandCard: []Card{{Rank: 0, Suit: "?"}, {Rank: 0, Suit: "?"}}}},
	}
	publicCards := []Card{
		// {Rank: 6, Suit: "红桃"},
		// {Rank: 7, Suit: "红桃"},
		// {Rank: 8, Suit: "红桃"},
		// {Rank: 9, Suit: "红桃"},
		// {Rank: 10, Suit: "红桃"},
		{Rank: 0, Suit: "?"},
		{Rank: 0, Suit: "?"},
		{Rank: 0, Suit: "?"},
		{Rank: 0, Suit: "?"},
		{Rank: 0, Suit: "?"},
	}
	var a, b, c int
	for i := 0; i < 100; i++ {
		winner, playres := shuffleJudgeGUI01(players, publicCards, false)
		if len(winner) > 1 {
			c++
			continue
		}
		if winner[0].ID == "AAA" {
			a++
		} else {
			b++
		}
		fmt.Println(playres)
	}
	fmt.Println(a, b, c)

	// winner := shuffleJudgeGUI01(players, publicCards, true)
	// for i, p := range winner {
	// 	fmt.Println("winner:", i)
	// 	fmt.Println(p.ID, p.Hand.HandCard[0].CardTranslate(), p.Hand.HandCard[1].CardTranslate())
	// 	fmt.Println(p.Grade, p.Card5)
	// }

	// // 测试用例2：相同的最高牌型，不同的五张牌
	// players = []Players{
	// 	{ID: "AAA", Hand: HandCard{HandCard: []Card{{Rank: 10, Suit: "黑桃"}, {Rank: 9, Suit: "黑桃"}}}},
	// 	{ID: "BBB", Hand: HandCard{HandCard: []Card{{Rank: 11, Suit: "红桃"}, {Rank: 8, Suit: "黑桃"}}}},
	// }
	// winner = shuffleJudgeGUI01(players, publicCards, false)
	// for i, p := range winner {
	// 	fmt.Println("winner:", i)
	// 	fmt.Println(p.ID, p.Hand.HandCard[0].CardTranslate(), p.Hand.HandCard[1].CardTranslate())
	// 	fmt.Println(p.Grade, p.Card5)
	// }
	// // 测试用例3：相同的最高牌型和五张牌
	// players = []Players{
	// 	{ID: "AAA", Hand: HandCard{HandCard: []Card{{Rank: 10, Suit: "黑桃"}, {Rank: 9, Suit: "黑桃"}}}},
	// 	{ID: "BBB", Hand: HandCard{HandCard: []Card{{Rank: 12, Suit: "黑桃"}, {Rank: 13, Suit: "黑桃"}}}},
	// }
	// winner = shuffleJudgeGUI01(players, publicCards, false)
	// for i, p := range winner {
	// 	fmt.Println("winner:", i)
	// 	fmt.Println(p.ID, p.Hand.HandCard[0].CardTranslate(), p.Hand.HandCard[1].CardTranslate())
	// 	fmt.Println(p.Grade, p.Card5)
	// }

}
func TestHandWinRateSimulationWeb01(t *testing.T) {
	// 设置测试数据
	player1 := Players{ID: "1", Hand: HandCard{HandCard: []Card{{Rank: 13, Suit: "黑桃"}, {Rank: 13, Suit: "红桃"}}}}
	player2 := Players{ID: "2", Hand: HandCard{HandCard: []Card{{Rank: 12, Suit: "黑桃"}, {Rank: 12, Suit: "红桃"}}}}
	player3 := Players{ID: "3", Hand: HandCard{HandCard: []Card{{Rank: 14, Suit: "黑桃"}, {Rank: 14, Suit: "红桃"}}}}
	// player1 := Players{ID: "1", Hand: HandCard{HandCard: []Card{{Rank: 0, Suit: "?"}, {Rank: 0, Suit: "?"}}}}
	// player2 := Players{ID: "2", Hand: HandCard{HandCard: []Card{{Rank: 0, Suit: "?"}, {Rank: 0, Suit: "?"}}}}
	// player3 := Players{ID: "3", Hand: HandCard{HandCard: []Card{{Rank: 0, Suit: "?"}, {Rank: 0, Suit: "?"}}}}
	player4 := Players{ID: "4", Hand: HandCard{HandCard: []Card{{Rank: 0, Suit: "?"}, {Rank: 0, Suit: "?"}}}}

	input := HandConfig{
		PlayerList: []Players{player1, player2, player3, player4},
		// PublicCard:  []Card{{Rank: 14, Suit: "梅花"}, {Rank: 14, Suit: "方片"}, {Rank: 4, Suit: "黑桃"}, {Rank: 5, Suit: "红桃"}, {Rank: 6, Suit: "梅花"}},
		PublicCard:  []Card{{Rank: 0, Suit: "?"}, {Rank: 0, Suit: "?"}, {Rank: 0, Suit: "?"}, {Rank: 0, Suit: "?"}, {Rank: 0, Suit: "?"}},
		RoundNumber: 10000,
		DebugSwitch: false,
	}

	// 调用被测方法
	result, err := HandWinRateSimulationWeb01(input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// fmt.Println(result.PlayersRes)
	// 将结构体转换为JSON字节切片
	jsonData, err := json.MarshalIndent(result.PlayersRes, "", "  ")
	if err != nil {
		fmt.Println("转换为JSON失败:", err)
		return
	}
	fmt.Println(string(jsonData))
	fmt.Println(result.DrawCount)
	fmt.Println(result.WinGradeList)
}

func TestShortOfShuffleCard(t *testing.T) {
	// 测试输入为空的情况
	input1 := []Card{{Rank: 0, Suit: "?"}, {Rank: 0, Suit: "?"}}
	for i := 0; i < 1000; i++ {
		result1 := ShortOfShuffleCard(input1)
		if len(result1) != 52 {
			t.Errorf("Expected length of result1 to be 2, but got %d", len(result1))
		}
		fmt.Println(result1[0].CardTranslate(), result1[1].CardTranslate(), result1[2].CardTranslate(), result1[3].CardTranslate(), result1[4].CardTranslate(), result1[51].CardTranslate())
	}

}
