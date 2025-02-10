package localbasefunc

import (
	"fmt"
	"time"

	"math/rand"
)

// // 指定人数 洗牌，发牌，并比较谁的牌最大,并且可以选择指定手牌
// func shuffleJudgeGUI01(playlist []Players, pulibcCard []Card, DebugSwitch bool) (winner []Players) {
// 	playerNum := len(playlist)
// 	var Card52 []Card
// 	var resHandList []HandCard
// 	var pubHandList []Card

// 	//添加返回
// 	ALLHandMap := make(map[string]int)
// 	if appointHandCardList != nil { //填写指定发牌的逻辑
// 		var tempAppointHand []Card
// 		var handTempList []HandCard
// 		var handTemp HandCard
// 		for i := 0; i < len(appointHandCardList); i++ {
// 			tempAppointHand = append(tempAppointHand, appointHandCardList[i].HandCard[0])
// 			tempAppointHand = append(tempAppointHand, appointHandCardList[i].HandCard[1])
// 			handTemp = HandCard{
// 				HandCard: [2]Card{appointHandCardList[i].HandCard[0], appointHandCardList[i].HandCard[1]},
// 			}
// 			handTempList = append(handTempList, handTemp)
// 		}
// 		// fmt.Println("指定牌的长度: ", len(tempAppointHand)) //debug
// 		Card52 := shortOfShuffleCard(tempAppointHand)
// 		resHandListTemp, pubHandListTemp := DealCards(Card52, playerNum-len(handTempList))
// 		resHandList = append(handTempList, resHandListTemp...)
// 		pubHandList = pubHandListTemp
// 		if DebugSwitch {
// 			for k, v := range resHandList {
// 				fmt.Println("玩家ID：", playlist[k].ID, "手牌：", v.HandCard[0].CardTranslate(), v.HandCard[1].CardTranslate(), "---debug---handcard")
// 			}
// 			fmt.Println("公共手牌：", pubHandList[0].CardTranslate(), pubHandList[1].CardTranslate(), pubHandList[2].CardTranslate(), pubHandList[3].CardTranslate(), pubHandList[4].CardTranslate(), "---debug---pubhandcard")
// 		}
// 	}
// 	if len(pulibcCard) == 0 && 所有玩家手牌都是空 { //如果没有指定手牌
// 		Card52 = ShuffleCard() //洗牌
// 		resHandList, pubHandList = DealCards(Card52, playerNum)
// 		if DebugSwitch {
// 			for k, v := range resHandList {
// 				fmt.Println("玩家ID：", playlist[k].ID, "手牌：", v.HandCard[0].CardTranslate(), v.HandCard[1].CardTranslate(), "---debug---handcard")
// 			}
// 			fmt.Println("公共手牌：", pubHandList[0].CardTranslate(), pubHandList[1].CardTranslate(), pubHandList[2].CardTranslate(), pubHandList[3].CardTranslate(), pubHandList[4].CardTranslate(), "---debug---pubhandcard")
// 		}

// 	}
// 	for i := 0; i < len(resHandList); i++ {
// 		if resHandList[i].HandCard[0].Suit == resHandList[i].HandCard[1].Suit {
// 			ALLHandMap[resHandList[i].HandCard[0].CardRankTranslate()+resHandList[i].HandCard[1].CardRankTranslate()+"s"]++
// 			continue
// 		}
// 		if resHandList[i].HandCard[0].Suit != resHandList[i].HandCard[1].Suit && resHandList[i].HandCard[0].Rank != resHandList[i].HandCard[1].Rank {
// 			ALLHandMap[resHandList[i].HandCard[0].CardRankTranslate()+resHandList[i].HandCard[1].CardRankTranslate()+"o"]++
// 			continue
// 		}
// 		if resHandList[i].HandCard[0].Rank == resHandList[i].HandCard[1].Rank {
// 			ALLHandMap[resHandList[i].HandCard[0].CardRankTranslate()+resHandList[i].HandCard[1].CardRankTranslate()]++
// 			continue
// 		}
// 	}
// 	// fmt.Println("手牌组：", len(resHandList), resHandList)        //debug
// 	// fmt.Println("公共手牌长度以及手牌 ", len(pubHandList), pubHandList) //debug

// 	maxGrade := 0
// 	maxCard5 := [5]int{0, 0, 0, 0, 0}
// 	// 假装这里已经处理了座次
// 	for i := 0; i < len(resHandList); i++ {
// 		var tempCard7 [7]Card
// 		playlist[i].Hand = resHandList[i]
// 		tempCard7[0] = playlist[i].Hand.HandCard[0]
// 		tempCard7[1] = playlist[i].Hand.HandCard[1]
// 		tempCard7[2] = pubHandList[0]
// 		tempCard7[3] = pubHandList[1]
// 		tempCard7[4] = pubHandList[2]
// 		tempCard7[5] = pubHandList[3]
// 		tempCard7[6] = pubHandList[4]
// 		tempCard7 = sortCards(tempCard7)
// 		playlist[i].Card7 = tempCard7
// 		playlist[i].Grade, playlist[i].Card5 = Judge5From7(playlist[i].Card7)
// 		// fmt.Println(playlist[i].ID, playlist[i].Hand)              //debug
// 		// fmt.Println(playlist[i].Grade, "-max-", playlist[i].Card5) //debug
// 		if maxGrade == playlist[i].Grade {
// 			for j := 0; j < 5; j++ {
// 				if playlist[i].Card5[j].Rank > maxCard5[j] {
// 					maxCard5[0] = playlist[i].Card5[0].Rank
// 					maxCard5[1] = playlist[i].Card5[1].Rank
// 					maxCard5[2] = playlist[i].Card5[2].Rank
// 					maxCard5[3] = playlist[i].Card5[3].Rank
// 					maxCard5[4] = playlist[i].Card5[4].Rank
// 					break
// 				}
// 				if playlist[i].Card5[j].Rank == maxCard5[j] {
// 					continue
// 				}
// 				if playlist[i].Card5[j].Rank < maxCard5[j] {
// 					break
// 				}
// 			}
// 			continue
// 		}
// 		if maxGrade < playlist[i].Grade {
// 			maxGrade = playlist[i].Grade
// 			for j := 0; j < 5; j++ {
// 				maxCard5[j] = playlist[i].Card5[j].Rank
// 			}
// 			continue
// 		}
// 	}
// 	// fmt.Println("len ", len(winner))   //debug
// 	// fmt.Println("maxGrade ", maxGrade) //debug
// 	// fmt.Println("maxCard5 ", maxCard5) //debug

// 	for i := 0; i < len(playlist); i++ {
// 		if playlist[i].Grade == maxGrade {
// 			// fmt.Println("最大的ID ", playlist[i].ID) //debug
// 			sign := true
// 			for j := 0; j < 5; j++ {
// 				if playlist[i].Card5[j].Rank == maxCard5[j] {
// 					continue
// 				} else {
// 					sign = false
// 				}
// 			}
// 			if sign {
// 				winner = append(winner, playlist[i])
// 			}
// 		}
// 	}
// 	// fmt.Println("len2 ", len(winner)) //debug

// 	return winner, ALLHandMap
// }

// sortCards对[7]Card数组进行排序
func sortCards(cards []Card) []Card {
	for i := 0; i < len(cards)-1; i++ {
		for j := 0; j < len(cards)-i-1; j++ {
			// 先比较牌面数字大小
			if cards[j].Rank < cards[j+1].Rank {
				cards[j], cards[j+1] = cards[j+1], cards[j]
			} else if cards[j].Rank == cards[j+1].Rank {
				// 牌面数字相同，比较花色权重
				if suitWeight[cards[j].Suit] < suitWeight[cards[j+1].Suit] {
					cards[j], cards[j+1] = cards[j+1], cards[j]
				}
			}
		}
	}
	return cards
}

// ConvertInputToCard 根据输入的字符串转换为 Card 结构体
func ConvertInputToCard(input string) (Card, error) {
	// 转换为 rune 切片，按字符处理
	runes := []rune(input)
	// 确保输入长度符合要求
	if len(runes) < 2 {
		if input == "?" {
			return Card{Rank: 0, Suit: "?"}, nil
		}
		return Card{}, fmt.Errorf("输入有问题")
	}
	// 提取花色和点数
	suitSymbol := string(runes[0])
	rankStr := string(runes[1:])
	suitMap := map[string]string{
		"♥": "红桃",
		"♦": "方片",
		"♣": "梅花",
		"♠": "黑桃",
	}
	// 转换花色
	suit, ok := suitMap[suitSymbol]
	if !ok {
		return Card{}, fmt.Errorf("花色输入有问题")
	}
	rankMap := map[string]int{
		"A":  14,
		"2":  2,
		"3":  3,
		"4":  4,
		"5":  5,
		"6":  6,
		"7":  7,
		"8":  8,
		"9":  9,
		"10": 10,
		"J":  11,
		"Q":  12,
		"K":  13,
	}
	// 转换点数
	rank, ok := rankMap[rankStr]
	if !ok {
		return Card{}, fmt.Errorf("rank输入有问题")
	}
	return Card{Rank: rank, Suit: suit}, nil
}

// CardTranslate 转换卡牌显示值
func (p Card) CardTranslate() string {
	suits := map[string]string{
		"黑桃": "♠",
		"红桃": "♥",
		"梅花": "♣",
		"方片": "♦",
	}

	ranks := map[int]string{
		14: "A", 13: "K", 12: "Q", 11: "J",
		10: "10", 9: "9", 8: "8", 7: "7",
		6: "6", 5: "5", 4: "4", 3: "3", 2: "2",
	}

	suitSymbol, suitExists := suits[p.Suit]
	rankSymbol, rankExists := ranks[p.Rank]

	if suitExists && rankExists {
		return suitSymbol + rankSymbol
	}
	return "fuck card"
}

// ShuffleCard 洗牌
func ShuffleCard() (New52CardList []Card) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	//初始化52张牌
	var Card52 = []Card{
		{Suit: "黑桃", Rank: 14},
		{Suit: "黑桃", Rank: 2},
		{Suit: "黑桃", Rank: 3},
		{Suit: "黑桃", Rank: 4},
		{Suit: "黑桃", Rank: 5},
		{Suit: "黑桃", Rank: 6},
		{Suit: "黑桃", Rank: 7},
		{Suit: "黑桃", Rank: 8},
		{Suit: "黑桃", Rank: 9},
		{Suit: "黑桃", Rank: 10},
		{Suit: "黑桃", Rank: 11},
		{Suit: "黑桃", Rank: 12},
		{Suit: "黑桃", Rank: 13},
		{Suit: "红桃", Rank: 14},
		{Suit: "红桃", Rank: 2},
		{Suit: "红桃", Rank: 3},
		{Suit: "红桃", Rank: 4},
		{Suit: "红桃", Rank: 5},
		{Suit: "红桃", Rank: 6},
		{Suit: "红桃", Rank: 7},
		{Suit: "红桃", Rank: 8},
		{Suit: "红桃", Rank: 9},
		{Suit: "红桃", Rank: 10},
		{Suit: "红桃", Rank: 11},
		{Suit: "红桃", Rank: 12},
		{Suit: "红桃", Rank: 13},
		{Suit: "梅花", Rank: 14},
		{Suit: "梅花", Rank: 2},
		{Suit: "梅花", Rank: 3},
		{Suit: "梅花", Rank: 4},
		{Suit: "梅花", Rank: 5},
		{Suit: "梅花", Rank: 6},
		{Suit: "梅花", Rank: 7},
		{Suit: "梅花", Rank: 8},
		{Suit: "梅花", Rank: 9},
		{Suit: "梅花", Rank: 10},
		{Suit: "梅花", Rank: 11},
		{Suit: "梅花", Rank: 12},
		{Suit: "梅花", Rank: 13},
		{Suit: "方片", Rank: 14},
		{Suit: "方片", Rank: 2},
		{Suit: "方片", Rank: 3},
		{Suit: "方片", Rank: 4},
		{Suit: "方片", Rank: 5},
		{Suit: "方片", Rank: 6},
		{Suit: "方片", Rank: 7},
		{Suit: "方片", Rank: 8},
		{Suit: "方片", Rank: 9},
		{Suit: "方片", Rank: 10},
		{Suit: "方片", Rank: 11},
		{Suit: "方片", Rank: 12},
		{Suit: "方片", Rank: 13},
	}
	// //洗牌
	// var new52 [52]Card
	// 洗牌
	r.Shuffle(len(Card52), func(i, j int) {
		Card52[i], Card52[j] = Card52[j], Card52[i]
	})
	// b := 0
	// r := rand.New(rand.NewSource(time.Now().Unix()))
	// for _, i := range r.Perm(len(Card52)) {
	// 	val := Card52[i]
	// 	// fmt.Println(val)
	// 	// fmt.Println(i)
	// 	new52[b] = val
	// 	b = b + 1
	// }
	//fmt.Println(new52)
	return Card52
}

// ShuffleCard 少牌洗牌
func ShortOfShuffleCard(input []Card) (New52CardList []Card) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	//初始化52张牌
	var Card52 = []Card{
		{Suit: "黑桃", Rank: 14},
		{Suit: "黑桃", Rank: 2},
		{Suit: "黑桃", Rank: 3},
		{Suit: "黑桃", Rank: 4},
		{Suit: "黑桃", Rank: 5},
		{Suit: "黑桃", Rank: 6},
		{Suit: "黑桃", Rank: 7},
		{Suit: "黑桃", Rank: 8},
		{Suit: "黑桃", Rank: 9},
		{Suit: "黑桃", Rank: 10},
		{Suit: "黑桃", Rank: 11},
		{Suit: "黑桃", Rank: 12},
		{Suit: "黑桃", Rank: 13},
		{Suit: "红桃", Rank: 14},
		{Suit: "红桃", Rank: 2},
		{Suit: "红桃", Rank: 3},
		{Suit: "红桃", Rank: 4},
		{Suit: "红桃", Rank: 5},
		{Suit: "红桃", Rank: 6},
		{Suit: "红桃", Rank: 7},
		{Suit: "红桃", Rank: 8},
		{Suit: "红桃", Rank: 9},
		{Suit: "红桃", Rank: 10},
		{Suit: "红桃", Rank: 11},
		{Suit: "红桃", Rank: 12},
		{Suit: "红桃", Rank: 13},
		{Suit: "梅花", Rank: 14},
		{Suit: "梅花", Rank: 2},
		{Suit: "梅花", Rank: 3},
		{Suit: "梅花", Rank: 4},
		{Suit: "梅花", Rank: 5},
		{Suit: "梅花", Rank: 6},
		{Suit: "梅花", Rank: 7},
		{Suit: "梅花", Rank: 8},
		{Suit: "梅花", Rank: 9},
		{Suit: "梅花", Rank: 10},
		{Suit: "梅花", Rank: 11},
		{Suit: "梅花", Rank: 12},
		{Suit: "梅花", Rank: 13},
		{Suit: "方片", Rank: 14},
		{Suit: "方片", Rank: 2},
		{Suit: "方片", Rank: 3},
		{Suit: "方片", Rank: 4},
		{Suit: "方片", Rank: 5},
		{Suit: "方片", Rank: 6},
		{Suit: "方片", Rank: 7},
		{Suit: "方片", Rank: 8},
		{Suit: "方片", Rank: 9},
		{Suit: "方片", Rank: 10},
		{Suit: "方片", Rank: 11},
		{Suit: "方片", Rank: 12},
		{Suit: "方片", Rank: 13},
	}
	var result []Card
	existMap := make(map[Card]bool)
	for _, v := range input {
		key := v
		existMap[key] = true
	}
	for _, v := range Card52 {
		key := v
		if !existMap[key] {
			result = append(result, v)
		}
	}
	Card52 = result

	// //洗牌
	// var new52 [52]Card
	// 洗牌
	r.Shuffle(len(Card52), func(i, j int) {
		Card52[i], Card52[j] = Card52[j], Card52[i]
	})

	return Card52
}

// DealCards 发牌，返回玩家手牌和公共牌  常规发牌
func DealCards(New52CardList []Card, playersNumber int) (resHandCard []HandCard, resPublicCard []Card) {
	// 初始化玩家手牌
	resHandCard = make([]HandCard, playersNumber)
	// 初始化每个玩家的手牌切片
	for i := range resHandCard {
		resHandCard[i].HandCard = make([]Card, 2)
	}
	resPublicCard = make([]Card, len(New52CardList)-(playersNumber*2))

	// 每个玩家分两张牌
	for j := 1; j <= playersNumber; j++ {
		resHandCard[j-1].HandCard[0] = New52CardList[j-1]
	}
	for j := 1; j <= playersNumber; j++ {
		resHandCard[j-1].HandCard[1] = New52CardList[playersNumber-1+j]
	}
	j := 0
	for i := 1; i <= len(New52CardList)-(playersNumber*2); i++ {
		resPublicCard[j] = New52CardList[2*playersNumber-1+i]
		j++
	}

	for i, k := range resHandCard {
		tmp := sortCards(k.HandCard)
		resHandCard[i].HandCard = tmp
	}
	return resHandCard, resPublicCard
}

// DealCards 发牌，返回玩家手牌和剩余牌  非常规发牌
func ShortLocalDealCards(pubCard []Card, playeListIN []Players) (playeListOut []Players, resPublicCard []Card) {
	playersNumber := len(playeListIN)
	// 初始化玩家手牌
	resHandCard := make([]HandCard, playersNumber)
	//初始化要剔除的牌
	var ShortCard []Card
	// 初始化每个玩家的手牌切片
	for i := range resHandCard {
		resHandCard[i].HandCard = make([]Card, 2)
	}
	// resPublicCard = make([]Card, 52)
	for i := 0; i < len(playeListIN); i++ {
		if playeListIN[i].Hand.HandCard[0].Rank == 0 && playeListIN[i].Hand.HandCard[1].Rank == 0 {
			continue
		}
		ShortCard = append(ShortCard, playeListIN[i].Hand.HandCard[0], playeListIN[i].Hand.HandCard[1])
	}
	ShortCard = append(ShortCard, pubCard...)
	short52Card := ShortOfShuffleCard(ShortCard)
	// 给每个玩家发牌
	for i := 0; i < len(playeListIN); i++ {
		if playeListIN[i].Hand.HandCard[0].Rank == 0 && playeListIN[i].Hand.HandCard[1].Rank == 0 { //将来也许可以改成一个
			playeListIN[i].Hand.HandCard[0] = short52Card[0]
			playeListIN[i].Hand.HandCard[1] = short52Card[1]
			//把已经赋值过得牌从short52Card中剔除   足够长
			short52Card = short52Card[2:]
		}
	}
	resPublicCard = append(resPublicCard, pubCard...)
	resPublicCard = append(resPublicCard, short52Card...)
	for i, k := range playeListIN {
		tmp := sortCards(k.Hand.HandCard)
		playeListIN[i].Hand.HandCard = tmp
	}
	playeListOut = playeListIN

	return playeListOut, resPublicCard
}
