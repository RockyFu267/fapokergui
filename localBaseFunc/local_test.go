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
