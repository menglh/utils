package stringsutil

import (
	"fmt"
	"io"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

// https://www.dotnetperls.com/between-before-after-go

// Between extracts the string between a and b
// returns value as is and error if a or b are not found
func Between(value string, a string, b string) (string, error) {
	after, err := After(value, a)
	if err != nil {
		return value, err
	}
	final, err := Before(after, b)
	if err != nil {
		return value, err
	}
	return final, nil
}

// Before extracts the string before a from value
// returns value as is and error if a is not found
func Before(value string, a string) (string, error) {
	pos := strings.Index(value, a)
	if pos == -1 {
		return value, fmt.Errorf("%s not found in %s", a, value)
	}
	return value[0:pos], nil
}

// After extracts the string after a from value
// returns value as is and error if a is not found
func After(value string, a string) (string, error) {
	pos := strings.Index(value, a)
	if pos == -1 {
		return value, fmt.Errorf("%s not found in %s", a, value)
	}
	adjustedPos := pos + len(a)
	if adjustedPos >= len(value) {
		return value, fmt.Errorf("After: %s is not long enough to contain %s", value, a)
	}
	return value[adjustedPos:], nil
}

// HasPrefixAny checks if the string starts with any specified prefix
func HasPrefixAny(s string, prefixes ...string) bool {
	for _, prefix := range prefixes {
		if strings.HasPrefix(s, prefix) {
			return true
		}
	}
	return false
}

// HasSuffixAny checks if the string ends with any specified suffix
func HasSuffixAny(s string, suffixes ...string) bool {
	for _, suffix := range suffixes {
		if strings.HasSuffix(s, suffix) {
			return true
		}
	}
	return false
}

// TrimPrefixAny trims all prefixes from string in order
func TrimPrefixAny(s string, prefixes ...string) string {
	for _, prefix := range prefixes {
		s = strings.TrimPrefix(s, prefix)
	}
	return s
}

// TrimSuffixAny trims all suffixes from string in order
func TrimSuffixAny(s string, suffixes ...string) string {
	for _, suffix := range suffixes {
		s = strings.TrimSuffix(s, suffix)
	}
	return s
}

// Join concatenates the elements of its first argument to create a single string. The separator
// string sep is placed between elements in the resulting string.
func Join(elems []interface{}, sep string) string {
	switch len(elems) {
	case 0:
		return ""
	case 1:
		return fmt.Sprint(elems[0])
	}
	n := len(sep) * (len(elems) - 1)
	for i := 0; i < len(elems); i++ {
		n += len(fmt.Sprint(elems[i]))
	}

	var b strings.Builder
	b.Grow(n)
	b.WriteString(fmt.Sprint(elems[0]))
	for _, s := range elems[1:] {
		b.WriteString(sep)
		b.WriteString(fmt.Sprint(s))
	}
	return b.String()
}

// HasPrefixI is case insensitive HasPrefix
func HasPrefixI(s, prefix string) bool {
	return strings.HasPrefix(strings.ToLower(s), strings.ToLower(prefix))
}

// HasSuffixI is case insensitive HasSuffix
func HasSuffixI(s, suffix string) bool {
	return strings.HasSuffix(strings.ToLower(s), strings.ToLower(suffix))
}

// Reverse the string
func Reverse(s string) string {
	n := 0
	rune := make([]rune, len(s))
	for _, r := range s {
		rune[n] = r
		n++
	}
	rune = rune[0:n]
	for i := 0; i < n/2; i++ {
		rune[i], rune[n-1-i] = rune[n-1-i], rune[i]
	}
	return string(rune)
}

// ContainsAny returns true is s contains any specified substring
func ContainsAny(s string, ss ...string) bool {
	for _, sss := range ss {
		if strings.Contains(s, sss) {
			return true
		}
	}
	return false
}

// EqualFoldAny returns true if s is equal to any specified substring
func EqualFoldAny(s string, ss ...string) bool {
	for _, sss := range ss {
		if strings.EqualFold(s, sss) {
			return true
		}
	}
	return false
}

// IndexAt look for a substring starting at position x
func IndexAt(s, sep string, n int) int {
	idx := strings.Index(s[n:], sep)
	if idx > -1 {
		idx += n
	}
	return idx
}

// SplitAny string by a list of separators
func SplitAny(s string, seps ...string) []string {
	sepsStr := strings.Join(seps, "")
	splitter := func(r rune) bool {
		return strings.ContainsRune(sepsStr, r)
	}
	return strings.FieldsFunc(s, splitter)
}

// SlideWithLength returns all the strings of the specified length while moving forward the extraction window
func SlideWithLength(s string, l int) chan string {
	out := make(chan string)

	go func(s string, l int) {
		defer close(out)

		if len(s) < l {
			out <- s
			return
		}

		for i := 0; i < len(s); i++ {
			if i+l <= len(s) {
				out <- s[i : i+l]
			} else {
				out <- s[i:]
				break
			}
		}
	}(s, l)

	return out
}

// ReplaceAll returns a copy of the string s with all
// instances of old incrementally replaced by new.
func ReplaceAll(s, new string, olds ...string) string {
	for _, old := range olds {
		s = strings.ReplaceAll(s, old, new)
	}
	return s
}

type LongestSequence struct {
	Sequence string
	Count    int
}

// LongestRepeatingSequence finds the longest repeating non-overlapping sequence in a string
func LongestRepeatingSequence(s string) LongestSequence {
	res := ""
	resLength := 0
	n := len(s)
	lcsre := make([][]int, n+1)

	for i := range lcsre {
		lcsre[i] = make([]int, n+1)
	}

	idx := 0
	for i := 1; i <= n; i++ {
		for j := i + 1; j <= n; j++ {
			if s[i-1] == s[j-1] && lcsre[i-1][j-1] < (j-i) {
				lcsre[i][j] = lcsre[i-1][j-1] + 1
				if lcsre[i][j] > resLength {
					resLength = lcsre[i][j]
					if i > idx {
						idx = i
					}
				}
			} else {
				lcsre[i][j] = 0
			}
		}
	}
	if resLength > 0 {
		for i := idx - resLength + 1; i <= idx; i++ {
			res += string(s[i-1])
		}
	}
	resCount := 0
	if res != "" {
		resCount = strings.Count(s, res)
	}
	return LongestSequence{Sequence: res, Count: resCount}
}

// Basename 返回某个文件路径的文件名，suffix如果提供了则去掉改后缀
func Basename(filename string, suffix ...string) string {
	name := filepath.Base(filename)
	for _, sf := range suffix {
		name = strings.TrimRight(name, sf)
	}
	return name
}

// Chown 改变文件所属者
func Chown(filename, username string) bool {
	ex := FileExists(filename)
	if !ex {
		return false
	}
	// windows下无意义
	if GetSystemOS() == "linux" {
		sysUser, err := user.Lookup(username)
		if err != nil {
			return false
		}
		uid, err := strconv.Atoi(sysUser.Uid)
		if err != nil {
			return false
		}
		gid, err := strconv.Atoi(sysUser.Gid)
		if err != nil {
			return false
		}
		err = os.Chown(filename, uid, gid)
		if err != nil {
			return false
		}
	}
	return true

}

// FileExists 检查文件是否存在
func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

// GetSystemOS 获取当前运行的操作系统类型
func GetSystemOS() string {
	return runtime.GOOS
}

// Copy 复制文件
func Copy(src, dest string) bool {
	if !FileExists(src) {
		return false
	}
	fileSrc, err := os.Open(src)
	if err != nil {
		return false
	}
	defer fileSrc.Close()
	fileDest, err := os.Create(dest)
	if err != nil {
		return false
	}
	defer fileDest.Close()
	_, err = io.Copy(fileDest, fileSrc)
	if err != nil {
		return false
	}
	return true
}

// UnLink 删除文件
func UnLink(filename string) bool {
	return os.Remove(filename) == nil
}

// Dirname 返回路径中的目录部分
func Dirname(path string) string {
	return filepath.Dir(path)
}

// FileGetContents 获取文件内容
func FileGetContents(filename string) []byte {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil
	}
	return data
}

// FilePutContents 写入文件内容
func FilePutContents(filename string, content []byte, fileMode os.FileMode) (int, error) {
	handle, err := os.OpenFile(filename, os.O_WRONLY, fileMode)
	if err != nil {
		return 0, err
	}
	defer handle.Close()
	n, err := handle.Write(content)
	if err != nil {
		return n, err
	}
	return n, nil
}

// FileATime 获取文件上次访问时间
// 对应windows下使用：
// fileSys := sys.(*syscall.Win32FileAttributeData)
// second := fileSys.LastAccessTime.Nanoseconds() / 1e9
// 对于linux下使用
// fileSys := sys.(*syscall.Stat_t)
// second := fileSys.Atim.Sec
func FileATime(filename string) (sys interface{}) {
	info, _ := os.Stat(filename)
	sys = info.Sys()
	return
}

// FileMTime 获取文件修改时间
func FileMTime(filename string) int64 {
	info, _ := os.Stat(filename)
	return info.ModTime().Unix()
}

// FileSize 获取文件大小
func FileSize(filename string) int64 {
	info, _ := os.Stat(filename)
	return info.Size()
}

// IsDir 判断是否为文件夹
func IsDir(filename string) bool {
	info, err := os.Stat(filename)
	if err != nil {
		return false
	}
	return info.IsDir()
}

// IsFile 判断是否为文件
func IsFile(filename string) bool {
	info, err := os.Stat(filename)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

// IsReadable 判断是否可读
func IsReadable(filename string) bool {
	h, err := os.Open(filename)
	if err != nil {
		return false
	}
	defer h.Close()
	return true
}

// IsWritable 判断文件是否可写
func IsWritable(filename string) bool {
	h, err := os.OpenFile(filename, os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return false
	}
	_ = h.Close()
	return true
}

// MkDir 创建文件夹
func MkDir(pathname string, fileMode os.FileMode, recursive bool) bool {
	if recursive {
		return os.MkdirAll(pathname, fileMode) == nil
	}
	return os.Mkdir(pathname, fileMode) == nil
}

// Rename 重命名
func Rename(src, dest string) error {
	return os.Rename(src, dest)
}

// ChMod 改变文件权限
func ChMod(filename string, mode os.FileMode) bool {
	return os.Chmod(filename, mode) == nil
}

// ChDir 改变当前目录
func ChDir(dir string) bool {
	return os.Chdir(dir) == nil
}

// GetCWD 取得当前工作目录
func GetCWD() string {
	dir, _ := os.Getwd()
	return dir
}

// ReadDir 读取目录，golang里面不用opendir
func ReadDir(dir string) ([]os.DirEntry, error) {
	return os.ReadDir(dir)
}
