package db

func GenCommentId() int64 {
	return SF.NextVal()
}
