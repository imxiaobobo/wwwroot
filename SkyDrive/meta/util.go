package meta

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
	"os"
)

//FileHash 计算文件hash值
func FileHash(file *os.File) string {
	file.Seek(0, 0)
	s1 := sha1.New()
	io.Copy(s1, file)
	return hex.EncodeToString(s1.Sum(nil))
}
