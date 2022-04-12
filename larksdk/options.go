package larksdk

import "time"

type Options struct {
	TotalTimeout time.Duration
	MaxRedirects int
}

func (opt *Options) init() {
	if opt.MaxRedirects < 0 {
		opt.MaxRedirects = 0
	} else if opt.MaxRedirects == 0 {
		opt.MaxRedirects = 3
	}
	if opt.TotalTimeout == 0 {
		opt.TotalTimeout = 3 * time.Second
	}
}
