// Package textanalyzer реалізує прості функції аналізу тексту для
// Завдання 0 (об'єднання text analyzer із Заняття 2 в multi-package проєкт).
package textanalyzer

import (
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"
)

// ErrEmptyText — sentinel error для порожнього вхідного тексту.
var ErrEmptyText = errors.New("textanalyzer: empty text")

func WordCount(text string) (int, error) {
	trimmed := strings.TrimSpace(text)
	if trimmed == "" {
		return 0, fmt.Errorf("textanalyzer: word count: %w", ErrEmptyText)
	}
	return len(strings.Fields(trimmed)), nil
}

func CharCount(text string) int {
	return utf8.RuneCountInString(strings.TrimSpace(text))
}
