package tui

import "time"

// After returns a command that waits for the specified duration
// then executes the provided command.
// The timer starts when After is called, not when the command executes.
func After(d time.Duration, cmd Cmd) Cmd {
	ch := time.After(d)
	return func() Msg {
		<-ch
		return cmd()
	}
}

// Every returns a command that waits until the next wall-clock aligned
// tick, then executes the provided command. For example, Every(time.Second, cmd)
// fires at the start of each second, and Every(time.Minute, cmd) fires at :00.
func Every(d time.Duration, cmd Cmd) Cmd {
	now := time.Now()
	next := now.Truncate(d).Add(d)
	return After(next.Sub(now), cmd)
}
