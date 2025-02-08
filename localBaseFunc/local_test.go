package localbasefunc

import (
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
