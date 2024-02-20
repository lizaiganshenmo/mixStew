package sensitiveword

import (
	"bufio"
	"io"
	"os"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/lizaiganshenmo/mixStew/library/utils"
)

var SensitiveTrie *utils.Trie

func Init(path string) {
	SensitiveTrie = utils.NewPrefixTree()

	fileHandle, err := os.OpenFile(path+"sensitive_word.txt", os.O_RDONLY, 0666)
	if err != nil {
		klog.Warn(err)
		return
	}
	defer fileHandle.Close()
	reader := bufio.NewReader(fileHandle)

	// 按行处理txt
	for {
		word, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if len(word) == 0 {
			continue
		}
		SensitiveTrie.Insert(string(word))
	}

}
