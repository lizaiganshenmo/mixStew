package db

func GenArticleId() int64 {
	return SF.NextVal()
}
