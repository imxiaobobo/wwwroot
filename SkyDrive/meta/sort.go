package meta

import "time"

/**
  自定义排序,需要实现sort中三个方法
  len：返回需要排序的长度
*/
type ByUploadTime []FileMeta

const baseTime = "2006-01-02 15:04:05"

func (b ByUploadTime) Len() int {
	return len(b)
}

func (b ByUploadTime) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b ByUploadTime) Less(i, j int) bool {
	iTime, _ := time.Parse(baseTime, b[i].UploadAt)
	jTime, _ := time.Parse(baseTime, b[j].UploadAt)
	return iTime.UnixNano() > jTime.UnixNano()
}
