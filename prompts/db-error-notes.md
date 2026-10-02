# Custom error (ручна частина)

DatabaseError має Query і OriginalErr.
FetchUser обгортає його через fmt.Errorf("%w").
У cmd/fetchuser/main.go перевірка errors.As і друк query.
