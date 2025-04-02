package watcher

import (
	"sync"
	"time"
)

type Debouncer struct {
	interval time.Duration
	lastSent time.Time
	mu       sync.Mutex
	timer    *time.Timer
	pending  string
	sendFunc func(msg string)
}

func NewDebouncer(interval time.Duration, sendFunc func(msg string)) *Debouncer {
	return &Debouncer{
		interval: interval,
		sendFunc: sendFunc,
	}
}

func (d *Debouncer) Trigger(newLine string) {
	d.mu.Lock()
	defer d.mu.Unlock()

	now := time.Now()
	if now.Sub(d.lastSent) >= d.interval {
		// Send immediately
		d.sendFunc(newLine)
		d.lastSent = now
	} else {
		// Store pending and schedule if not already
		d.pending = newLine
		if d.timer == nil {
			delay := d.interval - now.Sub(d.lastSent)
			d.timer = time.AfterFunc(delay, d.flush)
		}
	}
}

func (d *Debouncer) flush() {
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.pending != "" {
		d.sendFunc(d.pending)
		d.lastSent = time.Now()
		d.pending = ""
		d.timer = nil
	}
}
