package generator

import (
	"time"
)

type Event struct {
	Title    string
	Type     string
	DateTime time.Time

	Summary   string
	ImageLink string
	PreLink   string
	PostLink  string
}
