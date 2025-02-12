package localbasefunc

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"time"
)

var entertainmentIDs = []string{
	"扑克小鲁班",
	"牌桌范德彪",
	"赌圣他二舅",
	"气质男孩",
	"顺子制造机",
	"葫芦小王子",
	"皇家同花顺饲养员",
	"铁头娃",
	//以下都是台球名人的名字，姓都改成了傅，也可以是其他体育明行的名字
	"傅沙利文",
	"傅尔摩斯",
	"傅俊辉",
	"傅尔比",
	"傅金斯",
	"傅得利",
	"傅廉姆斯",
	"傅马赫",
	"傅俊辉",
}

// HandWinRateSimulationWeb01 单机模式功能1：模拟牌局统计胜率、牌型
func HandWinRateSimulationWeb01(input HandConfig) (WebRes PracticeRes, err error) {
	if len(input.PlayerList) < 1 || len(input.PlayerList) > 9 {
		return WebRes, fmt.Errorf("playNumber must 大于等于 0，小于 10")
	}
	if input.RoundNumber < 1 || input.RoundNumber > 100000 {
		return WebRes, fmt.Errorf("roundNumber must 大于等于 1,小于等于 100000")
	}
	//重复手牌校验前段已做

	// 初始化随机数生成器
	rng := rand.New(rand.NewSource(time.Now().UnixNano())) // 使用独立的随机数生成器

	// 随机分配娱乐ID
	usedIDs := make(map[string]bool) // 记录已分配的ID，避免重复

	for i := 0; i < len(input.PlayerList); i++ {
		for {
			randomIndex := rng.Intn(len(entertainmentIDs)) // 使用本地生成器生成随机索引
			randomID := entertainmentIDs[randomIndex]      // 选择对应的ID
			if !usedIDs[randomID] {                        // 检查是否已经被使用
				input.PlayerList[i].DisPlayName = input.PlayerList[i].ID + "-" + randomID // 分配给玩家
				usedIDs[randomID] = true                                                  // 标记为已使用
				break                                                                     // 跳出循环，分配下一个玩家
			}
		}
	}
	// 统计获胜的牌力类型
	winGradeList := make(map[string]int)
	// 统计获得胜利最多的玩家
	mostWinPlayer := make(map[string]int)
	type winnerKV struct {
		Key   string
		Value int
	}
	var mostWinPlayerSlice []winnerKV
	//统计获得胜利做的手牌
	mostWinHand := make(map[string]int)
	type hadnKV struct {
		Key   string
		Value int
	}
	//统计真实胜率
	allWinRealRate := make(map[string]int)
	type realWinRateKV struct {
		Key   string
		Value float64
	}

	var mostWinrHandSlice []hadnKV
	var allHandRateSlice []realWinRateKV
	var winGradeListSlice []winnerKV
	//统计平局的次数
	var tieCount int

	//所有被发出来的手牌统计
	allHandListOrigin := make(map[string]int)

	// 模拟牌局
	for i := 0; i < input.RoundNumber; i++ {
		// fmt.Println("第" + strconv.Itoa(i+1) + "局")
		winners := shuffleJudgeGUI01(input.PlayerList, input.PublicCard, input.DebugSwitch)
		if len(winners) > 1 {
			// fmt.Println("出现了多个玩家同时获得胜利的情况") //debug
			tieCount++
			switch winners[0].Grade { //统计获得胜利的牌力类型
			case 0:
				winGradeList["高牌"]++
			case 1:
				winGradeList["一对"]++
			case 2:
				winGradeList["两对"]++
			case 3:
				winGradeList["三条"]++
			case 4:
				winGradeList["顺子"]++
				// fmt.Println(winners[0].Card7[0].CardTranslate(), winners[0].Card7[1].CardTranslate(), winners[0].Card7[2].CardTranslate(), winners[0].Card7[3].CardTranslate(), winners[0].Card7[4].CardTranslate(), winners[0].Card7[5].CardTranslate(), winners[0].Card7[6].CardTranslate(), "---debug") //debug
			case 5:
				winGradeList["同花"]++
			case 6:
				winGradeList["葫芦"]++
			case 7:
				winGradeList["四条"]++
			case 8:
				winGradeList["同花顺"]++
			}
			continue
		}

		for _, v := range winners {
			// fmt.Println("--AAA---" + strconv.Itoa(i+1) + "---AAA---")
			// fmt.Println(k, v) //debug
			//统计获得胜利最多的玩家
			mostWinPlayer[v.ID]++
			//统计获得胜利做的手牌
			v.Hand.HandCard = sortCards(v.Hand.HandCard)
			mostWinHand[v.Hand.HandCard[0].CardTranslate()+v.Hand.HandCard[1].CardTranslate()]++
			switch v.Grade { //统计获得胜利的牌力类型
			case 0:
				winGradeList["高牌"]++
			case 1:
				winGradeList["一对"]++
			case 2:
				winGradeList["两对"]++
			case 3:
				winGradeList["三条"]++
			case 4:
				winGradeList["顺子"]++
				// fmt.Println(v.Card7[0].CardTranslate(), v.Card7[1].CardTranslate(), v.Card7[2].CardTranslate(), v.Card7[3].CardTranslate(), v.Card7[4].CardTranslate(), v.Card7[5].CardTranslate(), v.Card7[6].CardTranslate(), "------------------------debug") //debug
			case 5:
				winGradeList["同花"]++
			case 6:
				winGradeList["葫芦"]++
			case 7:
				winGradeList["四条"]++
			case 8:
				winGradeList["同花顺"]++
			}
			if v.Hand.HandCard[0].Suit == v.Hand.HandCard[1].Suit {
				allWinRealRate[v.Hand.HandCard[0].CardRankTranslate()+v.Hand.HandCard[1].CardRankTranslate()+"s"]++
				continue
			}
			if v.Hand.HandCard[0].Suit != v.Hand.HandCard[1].Suit && v.Hand.HandCard[0].Rank != v.Hand.HandCard[1].Rank {
				allWinRealRate[v.Hand.HandCard[0].CardRankTranslate()+v.Hand.HandCard[1].CardRankTranslate()+"o"]++
				continue
			}
			if v.Hand.HandCard[0].Rank == v.Hand.HandCard[1].Rank {
				allWinRealRate[v.Hand.HandCard[0].CardRankTranslate()+v.Hand.HandCard[1].CardRankTranslate()]++
				continue
			}
		}
	}

	for k, v := range mostWinPlayer {
		mostWinPlayerSlice = append(mostWinPlayerSlice, winnerKV{k, v})
		for i := 0; i < len(input.PlayerList); i++ {
			if input.PlayerList[i].ID == k {
				input.PlayerList[i].WinCount = v
				input.PlayerList[i].WinRate = float64(v) / float64(input.RoundNumber)
			}
		}
	}
	for k, v := range mostWinHand {
		mostWinrHandSlice = append(mostWinrHandSlice, hadnKV{k, v})
	}
	for k, v := range allWinRealRate {
		tempRate := float64(v) / float64(allHandListOrigin[k])
		allHandRateSlice = append(allHandRateSlice, realWinRateKV{k, tempRate})
	}
	for k, v := range winGradeList {
		winGradeListSlice = append(winGradeListSlice, winnerKV{k, v})
	}

	//输出排序结果
	sort.Slice(mostWinPlayerSlice, func(i, j int) bool {
		return mostWinPlayerSlice[i].Value > mostWinPlayerSlice[j].Value
	})
	sort.Slice(mostWinrHandSlice, func(i, j int) bool { //所有获得胜利的具体手牌
		return mostWinrHandSlice[i].Value > mostWinrHandSlice[j].Value
	})
	sort.Slice(allHandRateSlice, func(i, j int) bool { //所有手牌的胜率
		return allHandRateSlice[i].Value > allHandRateSlice[j].Value
	})
	sort.Slice(winGradeListSlice, func(i, j int) bool { //所有牌力的胜率
		return winGradeListSlice[i].Value > winGradeListSlice[j].Value
	})
	// 输出结果
	fmt.Println("玩家ID和对应ID:") //debug
	for k, v := range input.PlayerList {
		fmt.Println(k, v.ID)
	}
	// fmt.Println(mostWinPlayer) //debug
	// fmt.Println(mostWinHand) //debug
	fmt.Println("玩家胜利次数排序如下：") //debug
	for i := 0; i < len(mostWinPlayerSlice); i++ {
		fmt.Println(mostWinPlayerSlice[i].Key, mostWinPlayerSlice[i].Value)
		for j := 0; j < len(input.PlayerList); j++ {
			if mostWinPlayerSlice[i].Key == input.PlayerList[j].ID {
				input.PlayerList[j].WinCount = mostWinPlayerSlice[i].Value
				input.PlayerList[j].WinRate = float64(mostWinPlayerSlice[i].Value) / float64(input.RoundNumber)
			}
		}
	}
	if len(input.PlayerList) > 0 {
		n := len(mostWinrHandSlice)
		fmt.Println("获得过胜利的手牌组合数:", n)
		if n > 50 {
			n = 50
		}
		fmt.Println("胜利次数位于前列的手牌组合以及对应胜率:")
		for i := 0; i < n; i++ { //输出具体的卡牌
			fmt.Println(mostWinrHandSlice[i].Key, mostWinrHandSlice[i].Value, strconv.FormatFloat(float64(mostWinrHandSlice[i].Value)/float64(input.RoundNumber)*100, 'f', 4, 64)+"%")
		}
		//指定手牌的胜率
		fmt.Println("指定手牌的胜率如下:")
		for i := 0; i < len(input.PlayerList); i++ {
			temp := mostWinHand[input.PlayerList[i].Hand.HandCard[0].CardTranslate()+input.PlayerList[i].Hand.HandCard[1].CardTranslate()] //指定手牌获得的胜利次数
			realRate := float64(temp) / float64(input.RoundNumber)
			fmt.Println(input.PlayerList[i].Hand.HandCard[0].CardTranslate()+input.PlayerList[i].Hand.HandCard[1].CardTranslate(), temp, strconv.FormatFloat(realRate*100, 'f', 4, 64)+"%")
		}
	}
	fmt.Println("平局次数：", tieCount)
	WebRes.DrawCount = tieCount
	fmt.Println("成牌牌力分布统计：", winGradeList)
	for i := 0; i < len(winGradeListSlice); i++ {
		var tempWinGradeList WinGradeList
		tempWinGradeList.GradeName = winGradeListSlice[i].Key
		tempWinGradeList.WinCount = winGradeListSlice[i].Value
		switch winGradeListSlice[i].Key {
		case "高牌":
			tempWinGradeList.Grade = 0
		case "一对":
			tempWinGradeList.Grade = 1
		case "两对":
			tempWinGradeList.Grade = 2
		case "三条":
			tempWinGradeList.Grade = 3
		case "顺子":
			tempWinGradeList.Grade = 4
		case "同花":
			tempWinGradeList.Grade = 5
		case "葫芦":
			tempWinGradeList.Grade = 6
		case "四条":
			tempWinGradeList.Grade = 7
		case "同花顺":
			tempWinGradeList.Grade = 8

		}
		WebRes.WinGradeList = append(WebRes.WinGradeList, tempWinGradeList)
	}
	WebRes.PlayersRes = input.PlayerList
	return WebRes, nil
}

// 指定人数 洗牌，发牌，并比较谁的牌最大,并且可以选择指定手牌
func shuffleJudgeGUI01(playlist []Players, pulibcCard []Card, DebugSwitch bool) (winner []Players) {
	maxGrade := 0
	maxCard5 := [5]int{0, 0, 0, 0, 0}
	//开始发牌
	playerTemp, resCard := ShortLocalDealCards(pulibcCard, playlist)
	for i := 0; i < len(playerTemp); i++ {
		playerTemp[i].Card7[0] = playerTemp[i].Hand.HandCard[0]
		playerTemp[i].Card7[1] = playerTemp[i].Hand.HandCard[1]
		for j := 0; j < 5; j++ { //给card7赋值
			playerTemp[i].Card7[j+2] = resCard[j]
		}
		if DebugSwitch { //debug
			fmt.Println("player ", i, ": ", playerTemp[i].Card7)
		}
		playerTemp[i].Card7 = sortCards7(playerTemp[i].Card7)
		if DebugSwitch { //debug
			fmt.Println("player ", i, ": ", playerTemp[i].Card7)
		}
		playerTemp[i].Grade, playerTemp[i].Card5 = Judge5From7(playerTemp[i].Card7)
		if DebugSwitch { //debug
			fmt.Println("player ", i, ": ", playerTemp[i].Grade, "-max-", playerTemp[i].Card5)
		}
		if maxGrade == playerTemp[i].Grade {
			for j := 0; j < 5; j++ {
				if playerTemp[i].Card5[j].Rank > maxCard5[j] {
					maxCard5[0] = playerTemp[i].Card5[0].Rank
					maxCard5[1] = playerTemp[i].Card5[1].Rank
					maxCard5[2] = playerTemp[i].Card5[2].Rank
					maxCard5[3] = playerTemp[i].Card5[3].Rank
					maxCard5[4] = playerTemp[i].Card5[4].Rank
					break
				}
				if playerTemp[i].Card5[j].Rank == maxCard5[j] {
					continue
				}
				if playerTemp[i].Card5[j].Rank < maxCard5[j] {
					break
				}
			}
			continue
		}
		if maxGrade < playerTemp[i].Grade {
			maxGrade = playerTemp[i].Grade
			for j := 0; j < 5; j++ {
				maxCard5[j] = playerTemp[i].Card5[j].Rank
			}
			continue
		}

	}

	for i := 0; i < len(playerTemp); i++ {
		if playerTemp[i].Grade == maxGrade {
			// fmt.Println("最大的ID ", playerTemp[i].ID) //debug
			sign := true
			for j := 0; j < 5; j++ {
				if playerTemp[i].Card5[j].Rank == maxCard5[j] {
					continue
				} else {
					sign = false
				}
			}
			if sign {
				winner = append(winner, playerTemp[i])
			}
		}
	}
	if DebugSwitch {
		for k, v := range winner {
			fmt.Println("winner ", k, ": ", v.Hand.HandCard, v.Grade, v.Card5)
		}
	}
	// fmt.Println("len2 ", len(winner)) //debug

	return winner
}

// sortCards对[]Card数组进行排序
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

// sortCards7对[7]Card数组进行排序
func sortCards7(cards [7]Card) [7]Card {
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
	var pubTempCard []Card
	//校验pubCard是否合法
	if len(pubCard) != 0 {
		for i := 0; i < len(pubCard); i++ {
			if pubCard[i].Rank == 0 || pubCard[i].Suit == "?" {
				continue
			}
			pubTempCard = append(pubTempCard, pubCard[i])
		}
	}
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
	resPublicCard = append(resPublicCard, pubTempCard...)
	resPublicCard = append(resPublicCard, short52Card...)
	for i, k := range playeListIN {
		tmp := sortCards(k.Hand.HandCard)
		playeListIN[i].Hand.HandCard = tmp
	}
	playeListOut = playeListIN

	return playeListOut, resPublicCard
}

// Judge5From7 7选五的21种牌型的牌力，高牌的牌力为0，对子的牌力为1，两对的牌力为2，三条的牌力为3，顺子的牌力为4，同花的牌力为5，葫芦的牌力为6，四条的牌力为7，同花顺的牌力为8
func Judge5From7(playersAllCard [7]Card) (Grade int, MaxCard5 [5]Card) { // Judge5From7 7选五的21种牌型的牌力，高牌的牌力为0，对子的牌力为1，两对的牌力为2，三条的牌力为3，顺子的牌力为4，同花的牌力为5，葫芦的牌力为6，四条的牌力为7，同花顺的牌力为8
	//输入的7张牌，大小已经是按从大到小排列
	suitMap := make(map[string]int)       //定义四个花色的map，用来统计花色出现的次数
	suitListMap := make(map[string][]int) //定义四个花色的map，用来统计相同花色的rank
	sameMap := make(map[int]int)          //记录最多大小相同的牌的长度
	for i := 0; i < 7; i++ {
		suitMap[playersAllCard[i].Suit] = suitMap[playersAllCard[i].Suit] + 1                                     //记录花色
		sameMap[playersAllCard[i].Rank] = sameMap[playersAllCard[i].Rank] + 1                                     //记录大小相同的牌
		suitListMap[playersAllCard[i].Suit] = append(suitListMap[playersAllCard[i].Suit], playersAllCard[i].Rank) //记录相同花色的牌
	}

	//根据map长度来判断大小
	switch len(sameMap) { //这种写法不适合后期做娱乐技能判定，标准先这样
	case 2: //只有可能是金刚
		Grade = 7
		Value4 := 0 //不需要1
		for k, v := range sameMap {
			if v == 4 {
				Value4 = k
			}
		}
		if playersAllCard[0].Rank == Value4 { //不是43就是34
			MaxCard5[0] = playersAllCard[0]
			MaxCard5[1] = playersAllCard[1]
			MaxCard5[2] = playersAllCard[2]
			MaxCard5[3] = playersAllCard[3]
			MaxCard5[4] = playersAllCard[4]
		} else {
			MaxCard5[0] = playersAllCard[3]
			MaxCard5[1] = playersAllCard[4]
			MaxCard5[2] = playersAllCard[5]
			MaxCard5[3] = playersAllCard[6]
			MaxCard5[4] = playersAllCard[0]
		}
		return Grade, MaxCard5
	case 3: //可能是金刚也可能是葫芦  这种写法不适合后期做娱乐技能判定，标准先这样
		count3 := 0 //有3个相同的牌出现，那一定不是金刚了，值大于2就一定是葫芦（3+3+1）
		count2 := 0 //如果对子出现两次，那一定是葫芦（3+2+2）
		for _, v := range sameMap {
			if v == 3 {
				count3 = count3 + 1
				if count3 == 2 { //只有以下组合（3+3+1）
					Grade = 6
					if playersAllCard[0].Rank == playersAllCard[2].Rank && playersAllCard[3].Rank == playersAllCard[4].Rank { //3+3+1
						MaxCard5[0] = playersAllCard[0]
						MaxCard5[1] = playersAllCard[1]
						MaxCard5[2] = playersAllCard[2]
						MaxCard5[3] = playersAllCard[3]
						MaxCard5[4] = playersAllCard[4]
						return Grade, MaxCard5
					}
					if playersAllCard[0].Rank == playersAllCard[2].Rank && playersAllCard[4].Rank == playersAllCard[6].Rank { //3+1+3
						MaxCard5[0] = playersAllCard[0]
						MaxCard5[1] = playersAllCard[1]
						MaxCard5[2] = playersAllCard[2]
						MaxCard5[3] = playersAllCard[4]
						MaxCard5[4] = playersAllCard[5]
						return Grade, MaxCard5
					} else { //1+3+3
						MaxCard5[0] = playersAllCard[1]
						MaxCard5[1] = playersAllCard[2]
						MaxCard5[2] = playersAllCard[3]
						MaxCard5[3] = playersAllCard[4]
						MaxCard5[4] = playersAllCard[5]
						return Grade, MaxCard5
					}
				}
			}
			if v == 2 {
				count2 = count2 + 1
				if count2 == 2 { //只有以下组合（3+2+2）
					Grade = 6
					if playersAllCard[0].Rank == playersAllCard[2].Rank { //322
						MaxCard5[0] = playersAllCard[0]
						MaxCard5[1] = playersAllCard[1]
						MaxCard5[2] = playersAllCard[2]
						MaxCard5[3] = playersAllCard[3]
						MaxCard5[4] = playersAllCard[4]
						return Grade, MaxCard5
					}
					if playersAllCard[2].Rank == playersAllCard[4].Rank { //232
						MaxCard5[0] = playersAllCard[2]
						MaxCard5[1] = playersAllCard[3]
						MaxCard5[2] = playersAllCard[4]
						MaxCard5[3] = playersAllCard[0]
						MaxCard5[4] = playersAllCard[1]
						return Grade, MaxCard5
					} else { //223  playersAllCard[4].Rank == playersAllCard[6].Rank
						MaxCard5[0] = playersAllCard[4]
						MaxCard5[1] = playersAllCard[5]
						MaxCard5[2] = playersAllCard[6]
						MaxCard5[3] = playersAllCard[0]
						MaxCard5[4] = playersAllCard[1]
						return Grade, MaxCard5
					}
				}

			}
			if v == 4 { //如果有4个相同的牌，那一定是金刚 （4+2+1）
				Grade = 7
				if playersAllCard[0].Rank == playersAllCard[3].Rank { //4+ 2+1或1+2
					MaxCard5[0] = playersAllCard[0]
					MaxCard5[1] = playersAllCard[1]
					MaxCard5[2] = playersAllCard[2]
					MaxCard5[3] = playersAllCard[3]
					MaxCard5[4] = playersAllCard[4]
					return Grade, MaxCard5
				}
				if playersAllCard[1].Rank == playersAllCard[4].Rank { //1+4+2
					MaxCard5[0] = playersAllCard[1]
					MaxCard5[1] = playersAllCard[2]
					MaxCard5[2] = playersAllCard[3]
					MaxCard5[3] = playersAllCard[4]
					MaxCard5[4] = playersAllCard[0]
					return Grade, MaxCard5
				}
				if playersAllCard[2].Rank == playersAllCard[5].Rank { //2+4+1
					MaxCard5[0] = playersAllCard[2]
					MaxCard5[1] = playersAllCard[3]
					MaxCard5[2] = playersAllCard[4]
					MaxCard5[3] = playersAllCard[5]
					MaxCard5[4] = playersAllCard[0]
					return Grade, MaxCard5
				} else { //1+2或2+1 +4
					MaxCard5[0] = playersAllCard[3]
					MaxCard5[1] = playersAllCard[4]
					MaxCard5[2] = playersAllCard[5]
					MaxCard5[3] = playersAllCard[6]
					MaxCard5[4] = playersAllCard[0]
					return Grade, MaxCard5
				}

			}

		}
	case 4: //可能是金刚也可能是葫芦 也可能是两对  这种写法不适合后期做娱乐技能判定，标准先这样
		count2 := 0      //如果对子出现3次，那一定是2对 2+2+2+1
		count2Value := 0 //用于记录葫芦组合的对子的值
		for k, v := range sameMap {
			if v == 4 { // 如果有4个相同的牌，那一定是金刚 （4+1+1+1）
				Grade = 7
				count4Value := 0 //用于记录4个相同的牌  为了提前结束循环
				maxValue := 0    //用于记录最大的牌 	为了提前结束循环
				for i := 0; i < 7; i++ {
					if playersAllCard[i].Rank != k && maxValue == 0 { //不是4条的值就是最大值
						MaxCard5[4] = playersAllCard[i]
						maxValue = playersAllCard[i].Rank
						if count4Value != 0 { //所有值都已确认
							return Grade, MaxCard5
						}
						continue

					}
					if playersAllCard[i].Rank == k { //找到确认金刚的值
						MaxCard5[0] = playersAllCard[i]
						MaxCard5[1] = playersAllCard[i+1]
						MaxCard5[2] = playersAllCard[i+2]
						MaxCard5[3] = playersAllCard[i+3]
						i = i + 3
						count4Value = k
						if maxValue != 0 { //所有值都已确认
							return Grade, MaxCard5
						} else {
							continue
						}
					}
				}
				return Grade, MaxCard5
			}
			if v == 2 {
				count2Value = k //用于记录葫芦中的唯一对子

				count2 = count2 + 1
				if count2 == 3 { //只有以下组合（2+2+2+1）
					Grade = 2
					if playersAllCard[0].Rank == playersAllCard[1].Rank && playersAllCard[2].Rank == playersAllCard[3].Rank { //2+2+2+1
						MaxCard5[0] = playersAllCard[0]
						MaxCard5[1] = playersAllCard[1]
						MaxCard5[2] = playersAllCard[2]
						MaxCard5[3] = playersAllCard[3]
						MaxCard5[4] = playersAllCard[4]
						return Grade, MaxCard5
					}
					if playersAllCard[0].Rank == playersAllCard[1].Rank && playersAllCard[3].Rank == playersAllCard[4].Rank { //2+1+2+2
						MaxCard5[0] = playersAllCard[0]
						MaxCard5[1] = playersAllCard[1]
						MaxCard5[2] = playersAllCard[3]
						MaxCard5[3] = playersAllCard[4]
						MaxCard5[4] = playersAllCard[2]
						return Grade, MaxCard5
					} else { //1+2+2+2  playersAllCard[1].Rank == playersAllCard[2].Rank && playersAllCard[3].Rank == playersAllCard[4].Rank
						MaxCard5[0] = playersAllCard[1]
						MaxCard5[1] = playersAllCard[2]
						MaxCard5[2] = playersAllCard[3]
						MaxCard5[3] = playersAllCard[4]
						MaxCard5[4] = playersAllCard[0]
						return Grade, MaxCard5
					}
				}
			}
			if v == 3 { //只有可能是 3+2+ 1+ 1的葫芦组合
				Grade = 6
				for i := 0; i < len(playersAllCard); i++ {
					if playersAllCard[i].Rank == k { //葫芦中的3已确认
						MaxCard5[0] = playersAllCard[i]
						MaxCard5[1] = playersAllCard[i+1]
						MaxCard5[2] = playersAllCard[i+2]
						break
					}
				}
			}
		}
		for i := 0; i < len(playersAllCard); i++ { //葫芦中三条已在上面的循环中赋值了，还差葫芦中的对子
			if playersAllCard[i].Rank == count2Value {
				MaxCard5[3] = playersAllCard[i]
				MaxCard5[4] = playersAllCard[i+1]
				break
			}
		}
		return Grade, MaxCard5
	case 5: //可能是同花顺、同花、顺子、三条、两对
		straighACEtoFive := false
		straighACEtoFive = containsStraightKeys(sameMap)
		for k, v := range suitMap { //判断是否有同花，可能是同花、同花顺
			if v == 5 { //有同花
				if playersAllCard[0].Rank-playersAllCard[6].Rank == 4 { //同花顺，但不包括5432A的牌型
					Grade = 8
					j := 0 //maxCard5的下标
					for i := 0; i < 7; i++ {
						if playersAllCard[i].Suit == k {
							MaxCard5[j] = playersAllCard[i]
							j++
						}
					}
					return Grade, MaxCard5
				}
				if straighACEtoFive { //同花顺 指定牌型 5432A的牌型
					Grade = 8
					j := 0 //maxCard5的下标
					for i := 0; i < 7; i++ {
						if playersAllCard[i].Suit == k {
							MaxCard5[j] = playersAllCard[i] //顺序还要调整 现在是A5432
							j++
						}
					}
					MaxCard5[0], MaxCard5[1], MaxCard5[2], MaxCard5[3], MaxCard5[4] = MaxCard5[1], MaxCard5[2], MaxCard5[3], MaxCard5[4], MaxCard5[0] //调整顺序
					return Grade, MaxCard5
				} else { //没有同花顺的可能，就是同花
					Grade = 5
					j := 0 //maxCard5的下标
					for i := 0; i < 7; i++ {
						if playersAllCard[i].Suit == k {
							MaxCard5[j] = playersAllCard[i]
							j++
						}
					}
					return Grade, MaxCard5
				}
			}
		}
		if playersAllCard[0].Rank-playersAllCard[6].Rank == 4 { //没有同花的可能，可能是顺子，不包含5432A的牌型
			Grade = 4
			j := 0 //maxCard5的下标
			for i := 0; i < 7; i++ {
				if i == 0 { //第一个直接赋值 仅适用same长度为5的牌型
					MaxCard5[j] = playersAllCard[i]
					j++
					continue
				}
				if MaxCard5[j-1].Rank-playersAllCard[i].Rank == 1 {
					MaxCard5[j] = playersAllCard[i]
					j++
					continue
				}
			}
			return Grade, MaxCard5
		}
		if straighACEtoFive { //ace顺子 5432A
			Grade = 4
			acesign := false
			fiveSign := false
			fourSign := false
			threeSign := false
			twoSign := false
			for i := 0; i < 7; i++ {
				if playersAllCard[i].Rank == 14 && !acesign {
					MaxCard5[4] = playersAllCard[i]
					acesign = true
					continue
				}
				if playersAllCard[i].Rank == 5 && !fiveSign {
					MaxCard5[0] = playersAllCard[i]
					fiveSign = true
					continue
				}
				if playersAllCard[i].Rank == 4 && !fourSign {
					MaxCard5[1] = playersAllCard[i]
					fourSign = true
					continue
				}
				if playersAllCard[i].Rank == 3 && !threeSign {
					MaxCard5[2] = playersAllCard[i]
					threeSign = true
					continue
				}
				if playersAllCard[i].Rank == 2 && !twoSign {
					MaxCard5[3] = playersAllCard[i]
					twoSign = true
					continue
				}
			}
			return Grade, MaxCard5
		} else { //只能是两对或者三条
			pariRank1 := 0
			pariRank2 := 0
			cont3value := 0
			for k, v := range sameMap {
				if v == 2 { //只可能是两对
					Grade = 2
					if pariRank1 == 0 {
						pariRank1 = k
						continue
					}
					pariRank2 = k
					break
				}
				if v == 3 { //只可能是三条
					Grade = 6
					cont3value = k
					break
				}
			}
			if cont3value != 0 { //只能是三条
				Grade = 3
				max1 := 0
				max2 := 0
				for i := 0; i < len(playersAllCard); i++ {
					if playersAllCard[i].Rank == cont3value {
						MaxCard5[0] = playersAllCard[i]
						MaxCard5[1] = playersAllCard[i+1]
						MaxCard5[2] = playersAllCard[i+2]
						i = i + 2
						continue
					}
					if max1 == 0 {
						max1 = playersAllCard[i].Rank
						MaxCard5[3] = playersAllCard[i]
						continue
					}
					if max2 == 0 {
						max2 = playersAllCard[i].Rank
						MaxCard5[4] = playersAllCard[i]
						continue
					}
				}

				return Grade, MaxCard5

			}
			j := 0 //maxCard5的下标
			maxCardSign := false
			for i := 0; i < len(playersAllCard); i++ {
				if playersAllCard[i].Rank == pariRank1 || playersAllCard[i].Rank == pariRank2 {
					MaxCard5[j] = playersAllCard[i]
					j++
					continue
				} else {
					if !maxCardSign {
						MaxCard5[4] = playersAllCard[i]
						maxCardSign = true
					}
					continue
				}
			}
			return Grade, MaxCard5
		}
	case 6: //可能是同花顺、同花、顺子、两对
		straighACEtoFive := false
		straighACEtoFive = containsStraightKeys(sameMap)
		pairRank1 := 0
		var templist []int
		var tempList []int
		for k, v := range sameMap {
			templist = append(templist, k)
			if v == 2 {
				pairRank1 = k
			}
		}
		sortDescending(templist)
		tempList = templist
		if templist[0]-templist[4] == 4 {
			tempList = templist[:5]
		}
		if templist[1]-templist[5] == 4 && templist[0]-templist[4] != 4 {
			tempList = templist[1:]
		}

		for k, v := range suitListMap { //判断是否有同花，可能是同花、同花顺
			if len(v) == 6 { //有同花
				if suitListMap[k][0]-suitListMap[k][4] == 4 { //同花顺，但不包括5432A的牌型
					Grade = 8
					for i := 0; i < 7; i++ {
						if playersAllCard[i].Suit == k && playersAllCard[i].Rank == suitListMap[k][0] {
							MaxCard5[0] = playersAllCard[i]
						}
						if playersAllCard[i].Suit == k && playersAllCard[i].Rank == suitListMap[k][1] {
							MaxCard5[1] = playersAllCard[i]
						}
						if playersAllCard[i].Suit == k && playersAllCard[i].Rank == suitListMap[k][2] {
							MaxCard5[2] = playersAllCard[i]
						}
						if playersAllCard[i].Suit == k && playersAllCard[i].Rank == suitListMap[k][3] {
							MaxCard5[3] = playersAllCard[i]
						}
						if playersAllCard[i].Suit == k && playersAllCard[i].Rank == suitListMap[k][4] {
							MaxCard5[4] = playersAllCard[i]
						}
					}
					return Grade, MaxCard5
				}
				if suitListMap[k][1]-suitListMap[k][5] == 4 { //同花顺，但不包括5432A的牌型
					Grade = 8
					for i := 0; i < 7; i++ {
						if playersAllCard[i].Suit == k && playersAllCard[i].Rank == suitListMap[k][1] {
							MaxCard5[0] = playersAllCard[i]
						}
						if playersAllCard[i].Suit == k && playersAllCard[i].Rank == suitListMap[k][2] {
							MaxCard5[1] = playersAllCard[i]
						}
						if playersAllCard[i].Suit == k && playersAllCard[i].Rank == suitListMap[k][3] {
							MaxCard5[2] = playersAllCard[i]
						}
						if playersAllCard[i].Suit == k && playersAllCard[i].Rank == suitListMap[k][4] {
							MaxCard5[3] = playersAllCard[i]
						}
						if playersAllCard[i].Suit == k && playersAllCard[i].Rank == suitListMap[k][5] {
							MaxCard5[4] = playersAllCard[i]
						}
					}
					return Grade, MaxCard5
				}
				if suitListMap[k][0] == 14 && suitListMap[k][2] == 5 && suitListMap[k][3] == 4 && suitListMap[k][4] == 3 && suitListMap[k][5] == 2 { //同花顺 指定牌型 5432A的牌型

					Grade = 8
					for i := 0; i < 7; i++ {
						if playersAllCard[i].Suit == k && playersAllCard[i].Rank == suitListMap[k][0] {
							MaxCard5[4] = playersAllCard[i]
						}
						if playersAllCard[i].Suit == k && playersAllCard[i].Rank == suitListMap[k][2] {
							MaxCard5[0] = playersAllCard[i]
						}
						if playersAllCard[i].Suit == k && playersAllCard[i].Rank == suitListMap[k][3] {
							MaxCard5[1] = playersAllCard[i]
						}
						if playersAllCard[i].Suit == k && playersAllCard[i].Rank == suitListMap[k][4] {
							MaxCard5[2] = playersAllCard[i]
						}
						if playersAllCard[i].Suit == k && playersAllCard[i].Rank == suitListMap[k][5] {
							MaxCard5[3] = playersAllCard[i]
						}
					}
					return Grade, MaxCard5
				} else { //只是同花
					Grade = 5
					j := 0 //maxCard5的下标
					for i := 0; i < 7; i++ {
						if playersAllCard[i].Suit == k && j < 5 {
							MaxCard5[j] = playersAllCard[i]
							j++
						}
					}
					return Grade, MaxCard5
				}
			}
			if len(v) == 5 { //
				if suitListMap[k][0]-suitListMap[k][4] == 4 { //同花顺，但不包括5432A的牌型
					Grade = 8
					for i := 0; i < 7; i++ {
						if playersAllCard[i].Suit == k && playersAllCard[i].Rank == suitListMap[k][0] {
							MaxCard5[0] = playersAllCard[i]
						}
						if playersAllCard[i].Suit == k && playersAllCard[i].Rank == suitListMap[k][1] {
							MaxCard5[1] = playersAllCard[i]
						}
						if playersAllCard[i].Suit == k && playersAllCard[i].Rank == suitListMap[k][2] {
							MaxCard5[2] = playersAllCard[i]
						}
						if playersAllCard[i].Suit == k && playersAllCard[i].Rank == suitListMap[k][3] {
							MaxCard5[3] = playersAllCard[i]
						}
						if playersAllCard[i].Suit == k && playersAllCard[i].Rank == suitListMap[k][4] {
							MaxCard5[4] = playersAllCard[i]
						}
					}
					return Grade, MaxCard5
				}
				if suitListMap[k][0] == 14 && suitListMap[k][1] == 5 && suitListMap[k][2] == 4 && suitListMap[k][3] == 3 && suitListMap[k][4] == 2 { //同花顺 指定牌型 5432A的牌型
					Grade = 8
					for i := 0; i < 7; i++ {
						if playersAllCard[i].Suit == k && playersAllCard[i].Rank == suitListMap[k][0] {
							MaxCard5[4] = playersAllCard[i]
						}
						if playersAllCard[i].Suit == k && playersAllCard[i].Rank == suitListMap[k][1] {
							MaxCard5[0] = playersAllCard[i]
						}
						if playersAllCard[i].Suit == k && playersAllCard[i].Rank == suitListMap[k][2] {
							MaxCard5[1] = playersAllCard[i]
						}
						if playersAllCard[i].Suit == k && playersAllCard[i].Rank == suitListMap[k][3] {
							MaxCard5[2] = playersAllCard[i]
						}
						if playersAllCard[i].Suit == k && playersAllCard[i].Rank == suitListMap[k][4] {
							MaxCard5[3] = playersAllCard[i]
						}
					}
					return Grade, MaxCard5
				} else { //只是同花
					Grade = 5
					j := 0 //maxCard5的下标
					for i := 0; i < 7; i++ {
						if playersAllCard[i].Suit == k && j < 5 {
							MaxCard5[j] = playersAllCard[i]
							j++
						}
					}
					return Grade, MaxCard5
				}
			}

		}
		if len(tempList) == 5 { //没有同花的可能，只能是顺子
			Grade = 4
			j := 0 //maxCard5的下标
			for i := 0; i < 7; i++ {
				if tempList[j] == playersAllCard[i].Rank { //长度可能超过5 所以还得加上判断
					MaxCard5[j] = playersAllCard[i]
					if j >= 4 {
						break
					}
					j++
					continue
				}
			}
			return Grade, MaxCard5
		}
		if straighACEtoFive { //ace顺子 5432A
			Grade = 4
			acesign := false
			fiveSign := false
			fourSign := false
			threeSign := false
			twoSign := false
			for i := 0; i < 7; i++ {
				if playersAllCard[i].Rank == 14 && !acesign {
					MaxCard5[4] = playersAllCard[i]
					acesign = true
					continue
				}
				if playersAllCard[i].Rank == 5 && !fiveSign {
					MaxCard5[0] = playersAllCard[i]
					fiveSign = true
					continue
				}
				if playersAllCard[i].Rank == 4 && !fourSign {
					MaxCard5[1] = playersAllCard[i]
					fourSign = true
					continue
				}
				if playersAllCard[i].Rank == 3 && !threeSign {
					MaxCard5[2] = playersAllCard[i]
					threeSign = true
					continue
				}
				if playersAllCard[i].Rank == 2 && !twoSign {
					MaxCard5[3] = playersAllCard[i]
					twoSign = true
					continue
				}
			}
			return Grade, MaxCard5
		} else { //只能是对子
			Grade = 1
			j := 2 //maxCard5的下标
			for i := 0; i < 7; i++ {
				if playersAllCard[i].Rank == pairRank1 { //赋值对子
					MaxCard5[0] = playersAllCard[i]
					MaxCard5[1] = playersAllCard[i+1]
					i = i + 1
					continue
				} else { //赋值后三位
					if j < 5 {
						MaxCard5[j] = playersAllCard[i]
						j++
					}
				}
			}
			return Grade, MaxCard5
		}
	case 7: //可能是同花顺、同花、顺子、高牌
		for k, v := range suitMap { //判断是否是同花顺、顺子
			if v >= 5 { //至少是同花
				if playersAllCard[0].Rank-playersAllCard[4].Rank == 4 && playersAllCard[0].Suit == k && playersAllCard[1].Suit == k && playersAllCard[2].Suit == k && playersAllCard[3].Suit == k && playersAllCard[4].Suit == k { //是同花顺 不包含5432A
					Grade = 8
					MaxCard5[0] = playersAllCard[0]
					MaxCard5[1] = playersAllCard[1]
					MaxCard5[2] = playersAllCard[2]
					MaxCard5[3] = playersAllCard[3]
					MaxCard5[4] = playersAllCard[4]
					return Grade, MaxCard5
				}
				if playersAllCard[1].Rank-playersAllCard[5].Rank == 4 && playersAllCard[1].Suit == k && playersAllCard[2].Suit == k && playersAllCard[3].Suit == k && playersAllCard[4].Suit == k && playersAllCard[5].Suit == k { //是同花顺 不包含5432A
					Grade = 8
					MaxCard5[0] = playersAllCard[1]
					MaxCard5[1] = playersAllCard[2]
					MaxCard5[2] = playersAllCard[3]
					MaxCard5[3] = playersAllCard[4]
					MaxCard5[4] = playersAllCard[5]
					return Grade, MaxCard5
				}
				if playersAllCard[2].Rank-playersAllCard[6].Rank == 4 && playersAllCard[2].Suit == k && playersAllCard[3].Suit == k && playersAllCard[4].Suit == k && playersAllCard[5].Suit == k && playersAllCard[6].Suit == k { //是同花顺 不包含5432A
					Grade = 8
					MaxCard5[0] = playersAllCard[2]
					MaxCard5[1] = playersAllCard[3]
					MaxCard5[2] = playersAllCard[4]
					MaxCard5[3] = playersAllCard[5]
					MaxCard5[4] = playersAllCard[6]
					return Grade, MaxCard5
				}
				if playersAllCard[0].Rank == 14 && playersAllCard[3].Rank == 5 && playersAllCard[4].Rank == 4 && playersAllCard[5].Rank == 3 && playersAllCard[6].Rank == 2 && playersAllCard[0].Suit == k && playersAllCard[3].Suit == k && playersAllCard[4].Suit == k && playersAllCard[5].Suit == k && playersAllCard[6].Suit == k { //同花顺 5432A
					Grade = 8
					MaxCard5[4] = playersAllCard[0]
					MaxCard5[0] = playersAllCard[3]
					MaxCard5[1] = playersAllCard[4]
					MaxCard5[2] = playersAllCard[5]
					MaxCard5[3] = playersAllCard[6]
					return Grade, MaxCard5
				} else { // 只能是同花
					Grade = 5
					j := 0 //maxCard5的下标
					for i := 0; i < 7; i++ {
						if playersAllCard[i].Suit == k && j < 5 {
							MaxCard5[j] = playersAllCard[i]
							j++
						}
					}
					return Grade, MaxCard5
				}
			}
		}
		if playersAllCard[0].Rank-playersAllCard[4].Rank == 4 { //顺子 不包含5432A
			Grade = 4
			MaxCard5[0] = playersAllCard[0]
			MaxCard5[1] = playersAllCard[1]
			MaxCard5[2] = playersAllCard[2]
			MaxCard5[3] = playersAllCard[3]
			MaxCard5[4] = playersAllCard[4]
			return Grade, MaxCard5
		}
		if playersAllCard[1].Rank-playersAllCard[5].Rank == 4 { //顺子 不包含5432A
			Grade = 4
			MaxCard5[0] = playersAllCard[1]
			MaxCard5[1] = playersAllCard[2]
			MaxCard5[2] = playersAllCard[3]
			MaxCard5[3] = playersAllCard[4]
			MaxCard5[4] = playersAllCard[5]
			return Grade, MaxCard5
		}
		if playersAllCard[2].Rank-playersAllCard[6].Rank == 4 { //顺子 不包含5432A
			Grade = 4
			MaxCard5[0] = playersAllCard[2]
			MaxCard5[1] = playersAllCard[3]
			MaxCard5[2] = playersAllCard[4]
			MaxCard5[3] = playersAllCard[5]
			MaxCard5[4] = playersAllCard[6]
			return Grade, MaxCard5
		}
		if playersAllCard[0].Rank == 14 && playersAllCard[3].Rank == 5 && playersAllCard[4].Rank == 4 && playersAllCard[5].Rank == 3 && playersAllCard[6].Rank == 2 { //顺子 5432A
			Grade = 4
			MaxCard5[4] = playersAllCard[0]
			MaxCard5[0] = playersAllCard[3]
			MaxCard5[1] = playersAllCard[4]
			MaxCard5[2] = playersAllCard[5]
			MaxCard5[3] = playersAllCard[6]
			return Grade, MaxCard5
		} else { //只能是高牌
			Grade = 0
			MaxCard5[0] = playersAllCard[0]
			MaxCard5[1] = playersAllCard[1]
			MaxCard5[2] = playersAllCard[2]
			MaxCard5[3] = playersAllCard[3]
			MaxCard5[4] = playersAllCard[4]
			return Grade, MaxCard5
		}
	default:
		Grade = 0
		return Grade, MaxCard5
	}

	return 0, MaxCard5
}

// sortDescending 对切片进行降序排序
func sortDescending(arr []int) {
	// 使用 sort.Slice 进行降序排序
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j] // 比较函数，定义降序排序
	})
}

// containsStraightKeys 判断map中是否同时包含14, 2, 3, 4, 5的key
func containsStraightKeys(cards map[int]int) bool {
	requiredKeys := []int{14, 2, 3, 4, 5}

	for _, key := range requiredKeys {
		if _, exists := cards[key]; !exists {
			return false
		}
	}
	return true
}

// CardTranslate 转换卡牌rank的显示值
func (p Card) CardRankTranslate() string {

	ranks := map[int]string{
		14: "A", 13: "K", 12: "Q", 11: "J",
		10: "10", 9: "9", 8: "8", 7: "7",
		6: "6", 5: "5", 4: "4", 3: "3", 2: "2",
	}

	rankSymbol, rankExists := ranks[p.Rank]

	if rankExists {
		return rankSymbol
	}
	return "fuck card"
}
