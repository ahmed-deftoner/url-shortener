package routes

import (
	"time"
)

type request struct {
	URL         string        `json:"url"`
	CustomShort string        `json:"short"`
	Expiry      time.Duration `json:"expiry"`
}

type respomse struct {
	URL           string        `json:"url"`
	CustomShort   string        `json:"short"`
	Expiry        time.Duration `json:"expiry"`
	RateRemaining int           `json:"rate_limit"`
	RateLimitRest time.Duration `json:"rate_limit_remaining"`
}
