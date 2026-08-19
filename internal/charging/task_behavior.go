package charging

import (
	"errors"
	"fmt"
)

var ErrRetryableReminder = errors.New("retryable reminder failure")

func SendReminder() error                { return fmt.Errorf("provider timeout: %w", ErrRetryableReminder) }
func IsRetryableReminder(err error) bool { return err == ErrRetryableReminder }
