package models

type Team struct {
	Id          int
	Name        string
	TotalClicks int64
}

type Score struct {
	WhiteScore    int64   `json:"white_score"`
	WhitePercents float32 `json:"white_percents"`
	BlackScore    int64   `json:"black_score"`
	BlackPercents float32 `json:"black_percents"`
}
