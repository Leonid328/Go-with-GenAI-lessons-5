# Project structure (AI agent)

## Мій промпт

Є один великий main.go: FetchUser, DatabaseError і друк помилки в одному файлі.
Розклади по стандартній структурі Go:

- cmd/fetchuser/main.go тільки wiring
- internal/models
- internal/repository (FetchUser, DatabaseError)
- internal/service
- internal/handler

Залежності тільки в один бік: handler -> service -> repository.
errors.As має лишитися в main, щоб надрукувати Query.

## Що вийшло

Так і зробив. main тонкий, DatabaseError живе в repository.
