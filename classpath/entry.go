package classpath

import (
	"os"
	"strings"
)

const pathListSeparator = string(os.PathListSeparator)

// 类路径项
type Entry interface {
	readClass(className string) ([]byte, Entry, error) // 寻找和加载class
	String() string                                    // toString
}


func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		// 递归
		return newCompositeEntry(path)
	}

	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}

	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}

	return newDirEntry(path)
}
