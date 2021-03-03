package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

/**
 * 通配符
 * WildcardEntry与CompositeEntry结构一样
 * 通配符类路径跳过子目录
 */
func newWildcardEntry(path string) CompositeEntry {
	baseDir := path[:len(path)-1]
	compositeEntry := []Entry{}
	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}
		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
			jarEntry := newZipEntry(path)
			compositeEntry = append(compositeEntry, jarEntry)
		}
		return nil
	}
	// 遍历指定目录(包括子目录)，对遍历的项目用walkFn函数进行处理
	// WalkFunc 返回 nil，则 Walk 函数继续遍历
	// 如果返回 SkipDir，则 Walk 函数会跳过当前目录
	filepath.Walk(baseDir, walkFn)
	return compositeEntry

}
