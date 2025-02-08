package localbasefunc

// 定义花色对应的权重，用于比较相同牌面数字时的大小关系
var suitWeight = map[string]int{
	"黑桃": 4,
	"红桃": 3,
	"方片": 2,
	"梅花": 1,
}

// Card 牌的结构体
type Card struct {
	Suit string `json:"suit"`
	Rank int    `json:"rank"`
}

type HandConfig struct {
	PlayerList  []Players `json:"playerlist,omitempty"`
	RoundNumber int       `json:"roundnumber,omitempty"`
	DebugSwitch bool      `json:"debugswitch,omitempty"`
	PublicCard  []Card    `json:"publiccard,omitempty"`
}

// Players 玩家
type Players struct {
	ID           string  `json:"id"`
	Hand         []Card  `json:"hand"`
	ChipSum      int64   `json:"chipsum,omitempty"`
	ChipBackHand int64   `json:"chipbackhand,omitempty"`
	BankRollSum  int64   `json:"bankrollsum,omitempty"`
	Card7        [7]Card `json:"card7,omitempty"`
	Card5        [5]Card `json:"card5,omitempty"`
	CardAll      []Card  `json:"cardAll,omitempty"` //目前默认长度是7
	Grade        int     `json:"grade"`
	TableNum     int     `json:"tablenum,omitempty"`
	Sitnum       int     `json:"sitnum,omitempty"`
	IsActive     bool    `json:"isactive,omitempty"`
	IsFold       bool    `json:"isfold,omitempty"`
	IsAllIn      bool    `json:"isallin,omitempty"`
	WinCount     int     `json:"winCount"`           //单人训练统计次数用的
	WinRate      float64 `json:"winRate"`            //单人训练统计胜率用的
	Vpip         float64 `json:"vpip,omitempty"`     //Voluntarily Put In Pot主动入池率
	PFR          float64 `json:"pfr,omitempty"`      //Pre-Flop Raise 翻牌前加注概率
	FRR          float64 `json:"frr,omitempty"`      //Flop Raise 翻牌后加注概率
	TRR          float64 `json:"trr,omitempty"`      //Turn Raise 转牌后加注概率
	RRR          float64 `json:"rrr,omitempty"`      //River Raise 河牌后加注概率
	ReRaise      float64 `json:"rerraise,omitempty"` //Re-Raise 加注后加注概率
}
