package sensitiveword

// 判断word中是否包含敏感词
func HasSensitiveWord(word string) bool {
	return SensitiveTrie.Contains(word)
}
