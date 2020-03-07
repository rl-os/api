package config

import (
	"fmt"
	"reflect"
	"time"
)

// Load will unmarshal configurations to struct from files that you provide
func (c *Config) Load(files ...string) (err error) {
	config := c

	defaultValue := reflect.Indirect(reflect.ValueOf(config))
	if !defaultValue.CanAddr() {
		return fmt.Errorf("config %v should be addressable", config)
	}
	err, _ = c.load(config, false, files...)

	if c.AutoReload {
		go func() {
			timer := time.NewTimer(c.AutoReloadInterval)
			for range timer.C {
				reflectPtr := reflect.New(reflect.ValueOf(config).Elem().Type())
				reflectPtr.Elem().Set(defaultValue)

				var changed bool
				if err, changed = c.load(reflectPtr.Interface(), true, files...); err == nil && changed {
					reflect.ValueOf(config).Elem().Set(reflectPtr.Elem())
					if c.AutoReloadCallback != nil {
						c.AutoReloadCallback(config)
					}
				} else if err != nil {
					fmt.Printf("Failed to reload configuration from %v, got error %v\n", files, err)
				}
				timer.Reset(c.AutoReloadInterval)
			}
		}()
	}

	return
}