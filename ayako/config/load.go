package config

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"reflect"
	"time"
)

// Load will unmarshal configurations to struct from files that you provide
func (c *Config) Load(files ...string) (err error) {
	defaultValue := reflect.Indirect(reflect.ValueOf(c))
	if !defaultValue.CanAddr() {
		return fmt.Errorf("config %v should be addressable", c)
	}
	err, _ = c.load(c, false, files...)

	if c.AutoReload {
		go func() {
			timer := time.NewTimer(c.AutoReloadInterval)
			for range timer.C {
				reflectPtr := reflect.New(reflect.ValueOf(c).Elem().Type())
				reflectPtr.Elem().Set(defaultValue)

				var changed bool
				if err, changed = c.load(reflectPtr.Interface().(*Config), true, files...); err == nil && changed {
					reflect.ValueOf(c).Elem().Set(reflectPtr.Elem())
					if c.AutoReloadCallback != nil {
						c.AutoReloadCallback(c)
					}
				} else if err != nil {
					log.Error().Msgf("Failed to reload configuration from %v, got errors %v", files, err)
				}
				timer.Reset(c.AutoReloadInterval)
			}
		}()
	}

	return
}
