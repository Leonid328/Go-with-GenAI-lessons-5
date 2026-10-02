// Package retry містить завдання 2 домашньої роботи: retry-обгортку
// навколо "нестабільної" операції, що використовує sentinel error
// ErrTemporary для розпізнавання відновлюваних відмов.
package retry

import (
	"errors"
	"fmt"
	"time"
)

var ErrTemporary = errors.New("retry: temporary failure")

type Operation func() (string, error)

func Do(op Operation, maxAttempts int, backoff time.Duration) (string, error) {
	if maxAttempts < 1 {
		return "", fmt.Errorf("retry: maxAttempts must be >= 1")
	}

	var last error
	for i := 1; i <= maxAttempts; i++ {
		result, err := op()
		if err == nil {
			return result, nil
		}
		last = err
		if !errors.Is(err, ErrTemporary) {
			return "", fmt.Errorf("retry: permanent failure: %w", err)
		}
		if i == maxAttempts {
			break
		}
		time.Sleep(backoff)
	}
	return "", fmt.Errorf("retry: exhausted %d attempts: %w", maxAttempts, last)
}

func NewFlakyOperation(failuresBeforeSuccess int, successValue string) Operation {
	attempt := 0
	return func() (string, error) {
		attempt++
		if attempt <= failuresBeforeSuccess {
			return "", fmt.Errorf("flaky op: attempt %d failed: %w", attempt, ErrTemporary)
		}
		return successValue, nil
	}
}
