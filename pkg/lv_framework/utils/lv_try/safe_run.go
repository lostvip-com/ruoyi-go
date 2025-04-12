package lv_try

import "fmt"

// Catch catches a panic and returns an error.
func Catch(f func()) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Recovered from panic: %v", r)
		}
	}()

	// Call the provided function.
	f()

	// If we reach here, no panic occurred, so return nil error.
	return nil
}
