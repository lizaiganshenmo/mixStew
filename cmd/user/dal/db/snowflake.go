package db

func GenUID() int64 {
	return SF.NextVal()
}
