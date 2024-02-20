package utils

import (
	"regexp"
	"strings"
)

type Trie struct {
	// Val      byte
	PassNum  int
	EndNum   int
	Children map[rune]*Trie
}

func NewPrefixTree() *Trie {
	return &Trie{
		Children: make(map[rune]*Trie, 0),
	}
}

func (this *Trie) Insert(word string) {
	// n := len(word)
	cur := this
	for _, v := range word {
		cur.PassNum++

		if tmp := cur.Children[v]; tmp != nil {
			cur = tmp
		} else {
			tmp = &Trie{
				PassNum:  1,
				Children: make(map[rune]*Trie, 0),
			}
			cur.Children[v] = tmp
			cur = tmp
		}

	}
	cur.EndNum++

}

func (this *Trie) Search(word string) bool {
	cur := this
	for _, v := range word {
		if tmp := cur.Children[v]; tmp != nil {
			cur = tmp
		} else {
			return false
		}

	}

	if cur.EndNum == 0 {
		return false
	}

	return true

}

func (this *Trie) StartsWith(prefix string) bool {
	cur := this
	for _, v := range prefix {
		if tmp := cur.Children[v]; tmp != nil {
			cur = tmp
		} else {
			return false
		}

	}

	return true

}

func (this *Trie) FilterSpecialChar(text string) string {
	text = strings.ToLower(text)
	text = strings.ReplaceAll(text, " ", "") // 去除空格

	// 过滤除中英文及数字以外的其他字符
	otherCharReg := regexp.MustCompile("[^\u4e00-\u9fa5a-zA-Z0-9]")
	text = otherCharReg.ReplaceAllString(text, "")
	return text
}

// word中是否包含有敏感词
func (this *Trie) Contains(word string) bool {
	if len(word) == 0 {
		return false
	}

	// 过滤特殊字符, 避免出现 "沙 比" "沙 , 比" 等有害评论
	text := this.FilterSpecialChar(word)
	textChars := []rune(text)
	n := len(textChars)
	for i := range textChars {
		node := this.getChildNode(textChars[i])
		if node == nil {
			continue
		}

		// 前缀树中存在该字符为前缀的, 继续遍历寻找
		j := i + 1
		for ; j < n; j++ {
			if node.EndNum > 0 {
				return true
			}

			node = node.getChildNode(textChars[j])

		}

		if j == n && node != nil && node.EndNum > 0 {
			return true
		}

	}

	return false

}

func (this *Trie) getChildNode(c rune) *Trie {
	if this.Children == nil {
		return nil
	}

	return this.Children[c]
}
