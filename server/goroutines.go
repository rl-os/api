package server

import (
	"github.com/rs/zerolog/log"
	"reflect"
	"sync/atomic"
)

// Go creates a goroutine, but maintains a record of it to ensure that execution completes before
// the server is shutdown.
func (s *Server) Go(f func()) {
	atomic.AddInt32(&s.GoroutineCount, 1)
	log.Debug().
		Str("func_path", reflect.TypeOf(f).PkgPath()).
		Msg("Starting new goroutine")

	go func() {
		f()

		atomic.AddInt32(&s.GoroutineCount, -1)
		select {
		case s.GoroutineExitSignal <- struct{}{}:
		default:
		}
	}()
}

// WaitForGoroutines blocks until all goroutines created by App.Go exit.
func (s *Server) WaitForGoroutines() {
	log.Debug().
		Int32("total", s.GoroutineCount).
		Msg("waiting goroutines before exit app")
	for atomic.LoadInt32(&s.GoroutineCount) != 0 {
		<-s.GoroutineExitSignal
	}
}
