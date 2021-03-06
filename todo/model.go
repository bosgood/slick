package todo

import (
	"strings"
	"time"
)

type Todo []*Task

type Task struct {
	ID          string
	CreatedAt   time.Time
	User        string
	Text        []string
	Closed      bool
	ClosingNote string
	ClosedAt    time.Time
}

func (t *Task) String() string {
	out := "`" + t.ID + "` "
	text := strings.Join(t.Text, " // ")

	if t.Closed {
		out += "~" + text + "~"
	} else {
		out += text
	}

	if t.ClosingNote != "" {
		out += " _" + t.ClosingNote + "_"
	}
	return out
}
