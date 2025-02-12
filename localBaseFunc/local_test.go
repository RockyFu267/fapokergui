package localbasefunc

import (
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
	pubCard := []Card{{Suit: "？", Rank: 0}, {Suit: "?", Rank: 0}}
	playerList := []Players{
		{Hand: HandCard{HandCard: []Card{{Suit: "?", Rank: 0}, {Suit: "?", Rank: 0}}}},
		{Hand: HandCard{HandCard: []Card{{Suit: "?", Rank: 0}, {Suit: "?", Rank: 0}}}},
	}

	// 执行
	playerList, pubCard = ShortLocalDealCards(pubCard, playerList)

	fmt.Println(playerList)
	for k, v := range pubCard {
		fmt.Println(k, v.CardTranslate())
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
		{ID: "AAA", Hand: HandCard{HandCard: []Card{{Rank: 2, Suit: "红桃"}, {Rank: 3, Suit: "红桃"}}}},
		{ID: "BBB", Hand: HandCard{HandCard: []Card{{Rank: 4, Suit: "红桃"}, {Rank: 5, Suit: "红桃"}}}},
	}
	publicCards := []Card{
		{Rank: 6, Suit: "红桃"},
		{Rank: 7, Suit: "红桃"},
		{Rank: 8, Suit: "红桃"},
		{Rank: 9, Suit: "红桃"},
		{Rank: 10, Suit: "红桃"},
	}
	winner := shuffleJudgeGUI01(players, publicCards, true)
	for i, p := range winner {
		fmt.Println("winner:", i)
		fmt.Println(p.ID, p.Hand.HandCard[0].CardTranslate(), p.Hand.HandCard[1].CardTranslate())
		fmt.Println(p.Grade, p.Card5)
	}

	// 测试用例2：相同的最高牌型，不同的五张牌
	players = []Players{
		{ID: "AAA", Hand: HandCard{HandCard: []Card{{Rank: 10, Suit: "黑桃"}, {Rank: 9, Suit: "黑桃"}}}},
		{ID: "BBB", Hand: HandCard{HandCard: []Card{{Rank: 11, Suit: "红桃"}, {Rank: 8, Suit: "黑桃"}}}},
	}
	winner = shuffleJudgeGUI01(players, publicCards, false)
	for i, p := range winner {
		fmt.Println("winner:", i)
		fmt.Println(p.ID, p.Hand.HandCard[0].CardTranslate(), p.Hand.HandCard[1].CardTranslate())
		fmt.Println(p.Grade, p.Card5)
	}
	// 测试用例3：相同的最高牌型和五张牌
	players = []Players{
		{ID: "AAA", Hand: HandCard{HandCard: []Card{{Rank: 10, Suit: "黑桃"}, {Rank: 9, Suit: "黑桃"}}}},
		{ID: "BBB", Hand: HandCard{HandCard: []Card{{Rank: 12, Suit: "黑桃"}, {Rank: 13, Suit: "黑桃"}}}},
	}
	winner = shuffleJudgeGUI01(players, publicCards, false)
	for i, p := range winner {
		fmt.Println("winner:", i)
		fmt.Println(p.ID, p.Hand.HandCard[0].CardTranslate(), p.Hand.HandCard[1].CardTranslate())
		fmt.Println(p.Grade, p.Card5)
	}

}
func TestHandWinRateSimulationWeb01(t *testing.T) {
	// 设置测试数据
	player1 := Players{ID: "1", Hand: HandCard{HandCard: []Card{{Rank: 13, Suit: "黑桃"}, {Rank: 13, Suit: "红桃"}}}}
	player2 := Players{ID: "2", Hand: HandCard{HandCard: []Card{{Rank: 12, Suit: "黑桃"}, {Rank: 12, Suit: "红桃"}}}}
	player3 := Players{ID: "3", Hand: HandCard{HandCard: []Card{{Rank: 14, Suit: "黑桃"}, {Rank: 14, Suit: "红桃"}}}}
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

	fmt.Println(result.PlayersRes)
	fmt.Println(result.DrawCount)
	fmt.Println(result.WinGradeList)
}
