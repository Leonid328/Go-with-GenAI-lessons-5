# Identifying Ignored Errors with AI

## Bad snippet (before)

```go
package main

import (
	"os"
	"strconv"
)

func loadAge(path string) int {
	b, _ := os.ReadFile(path)
	n, err := strconv.Atoi(string(b))
	_ = err
	if n < 0 {
		panic("age must be >= 0")
	}
	return n
}

func main() {
	fmt.Println(loadAge("age.txt"))
}
```

## Мій промпт

Review this Go code for missing or incomplete error handling. List every location where an error is ignored, mishandled, or under-wrapped, and propose a fix for each using best practices like fmt.Errorf with %w.

## Що ШІ знайшов

1. `b, _ := os.ReadFile(path)` — помилка читання файлу викинута.
2. `_ = err` після `strconv.Atoi` — parse error ігнорується.
3. `panic("age must be >= 0")` — валідація через panic, не через error.
4. Немає wrapping, викликач не бачить path і причину.

## Після фіксу

```go
func loadAge(path string) (int, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return 0, fmt.Errorf("load age from %s: %w", path, err)
	}
	n, err := strconv.Atoi(strings.TrimSpace(string(b)))
	if err != nil {
		return 0, fmt.Errorf("parse age from %s: %w", path, err)
	}
	if n < 0 {
		return 0, fmt.Errorf("load age from %s: age must be >= 0", path)
	}
	return n, nil
}
```
