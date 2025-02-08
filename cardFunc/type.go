package cardFunc

// Card 牌的结构体
type Card struct {
	Suit string `json:"suit"`
	Rank int    `json:"rank"`
}

// HanCard 手牌
type HandCard struct {
	HandCard [2]Card `json:"handCard"`
	PlayerID string  `json:"playerid,omitempty"`
}

type HandConfig struct {
	PlayerNumber int        `json:"playernumber,omitempty"`
	HandCardList []HandCard `json:"handcardlist,omitempty"`
	RoundNumber  int        `json:"roundnumber,omitempty"`
	DebugSwitch  bool       `json:"debugswitch,omitempty"`
	PublicCard   []Card     `json:"publiccard,omitempty"`
}

// Players 玩家
type Players struct {
	ID           string   `json:"id"`
	Hand         HandCard `json:"hand"`
	ChipSum      int64    `json:"chipsum,omitempty"`
	ChipBackHand int64    `json:"chipbackhand,omitempty"`
	BankRollSum  int64    `json:"bankrollsum,omitempty"`
	Card7        [7]Card  `json:"card7,omitempty"`
	Card5        [5]Card  `json:"card5,omitempty"`
	CardAll      []Card   `json:"cardAll,omitempty"` //目前默认长度是7
	Grade        int      `json:"grade"`
	TableNum     int      `json:"tablenum,omitempty"`
	Sitnum       int      `json:"sitnum,omitempty"`
	IsActive     bool     `json:"isactive,omitempty"`
	IsFold       bool     `json:"isfold,omitempty"`
	IsAllIn      bool     `json:"isallin,omitempty"`
	WinCount     int      `json:"winCount"`           //单人训练统计次数用的
	WinRate      float64  `json:"winRate"`            //单人训练统计胜率用的
	Vpip         float64  `json:"vpip,omitempty"`     //Voluntarily Put In Pot主动入池率
	PFR          float64  `json:"pfr,omitempty"`      //Pre-Flop Raise 翻牌前加注概率
	FRR          float64  `json:"frr,omitempty"`      //Flop Raise 翻牌后加注概率
	TRR          float64  `json:"trr,omitempty"`      //Turn Raise 转牌后加注概率
	RRR          float64  `json:"rrr,omitempty"`      //River Raise 河牌后加注概率
	ReRaise      float64  `json:"rerraise,omitempty"` //Re-Raise 加注后加注概率
}

type PracticeResDemo02 struct {
	PlayerWinCount []PlayersRes   `json:"playerwincount"`       //统计玩家的获胜次数  按座次排序
	WinGradeList   []WinGradeList `json:"wingradelist"`         //获胜的成牌牌力分布统计  按出现次数排序
	DrawCount      int            `json:"drawcount"`            //平局次数
	So169ComboList []So169Combo   `json:"so169combo,omitempty"` //169组so组合的胜率统计  按胜率排序
	RoundNumber    int            `json:"roundnumber,omitempty"`
}
type PlayersRes struct {
	PlayerID string  `json:"playerid"`
	WinCount int     `json:"wincount"`
	WinRate  float64 `json:"winrate"`
}

type So169Combo struct {
	WinRateRank int     `json:"winraterank"`
	So169       string  `json:"so169"`
	WinRate     float64 `json:"winrate"`
	ExistCount  int     `json:"existcount"`
	WinCount    int     `json:"wincount"`
}
type WinGradeList struct {
	Grade     int    `json:"grade"`
	GradeName string `json:"gradename"`
	WinCount  int    `json:"wincount"`
}
