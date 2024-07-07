package app

import "time"

type Trigger interface {
	On() <-chan time.Time
	Stop()
}

type TickerTrigger struct {
	ticker *time.Ticker
}

func NewTickerTrigger(period time.Duration) *TickerTrigger {
	return &TickerTrigger{
		ticker: time.NewTicker(period),
	}
}

func (x *TickerTrigger) On() <-chan time.Time {
	return x.ticker.C
}

func (x *TickerTrigger) Stop() {
	x.ticker.Stop()
}
