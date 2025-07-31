package stock

import "time"

type Item struct {
	Ticker     string    `json:"ticker"`
	TargetFrom string    `json:"target_from"`
	TargetTo   string    `json:"target_to"`
	Company    string    `json:"company"`
	Action     string    `json:"action"`
	Brokerage  string    `json:"brokerage"`
	RatingFrom string    `json:"rating_from"`
	RatingTo   string    `json:"rating_to"`
	Time       time.Time `json:"time"`
}

type listResponse struct {
	Items    []Item `json:"items"`
	NextPage string `json:"next_page"`
}

type loginResponse struct {
	Token   string `json:"auth_token"`
	Message string `json:"message"`
}
