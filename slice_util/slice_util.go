package slice_util

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func init() {
	// initialize seed once (ref. https://stackoverflow.com/a/12321192/8155097)
	rand.Seed(time.Now().UnixNano())
}

// PruneEmptyStrings from the slice
func PruneEmptyStrings(v []string) []string {
	return PruneEqual(v, "")
}

// PruneEqual removes items from the slice equal to the specified value
func PruneEqual[T comparable](inputSlice []T, equalTo T) (r []T) {
	for i := range inputSlice {
		if inputSlice[i] != equalTo {
			r = append(r, inputSlice[i])
		}
	}

	return
}

// Dedupe removes duplicates from a slice of elements preserving the order
func Dedupe[T comparable](inputSlice []T) (result []T) {
	seen := make(map[T]struct{})
	for _, inputValue := range inputSlice {
		if _, ok := seen[inputValue]; !ok {
			seen[inputValue] = struct{}{}
			result = append(result, inputValue)
		}
	}

	return
}

// PickRandom item from a slice of elements
func PickRandom[T any](v []T) T {
	return v[rand.Intn(len(v))]
}

// Contains if a slice contains an element
func Contains[T comparable](inputSlice []T, element T) bool {
	for _, inputValue := range inputSlice {
		if inputValue == element {
			return true
		}
	}

	return false
}

// ContainsItems checks if s1 contains s2
func ContainsItems[T comparable](s1 []T, s2 []T) bool {
	for _, e := range s2 {
		if !Contains(s1, e) {
			return false
		}
	}

	return true
}

// ToInt converts a slice of strings to a slice of ints
func ToInt(s []string) ([]int, error) {
	var ns []int
	for _, ss := range s {
		n, err := strconv.Atoi(ss)
		if err != nil {
			return nil, err
		}
		ns = append(ns, n)
	}

	return ns, nil
}

// Equal checks if the items of two slices are equal respecting the order
func Equal[T comparable](s1, s2 []T) bool {
	if len(s1) != len(s2) {
		return false
	}

	for idx := range s1 {
		if s1[idx] != s2[idx] {
			return false
		}
	}

	return true
}

// IsEmpty checks if the slice has length zero
func IsEmpty[V comparable](s []V) bool {
	return len(s) == 0
}

// ElementsMatch asserts that the specified listA(array, slice...) is equal to specified
// listB(array, slice...) ignoring the order of the elements. If there are duplicate elements,
// the number of appearances of each of them in both lists should match.
func ElementsMatch[V comparable](s1, s2 []V) bool {
	if IsEmpty(s1) && IsEmpty(s2) {
		return true
	}

	extraS1, extrS2 := Diff(s1, s2)
	return IsEmpty(extraS1) && IsEmpty(extrS2)
}

// Diff calculates the extra elements between two sequences
func Diff[V comparable](s1, s2 []V) (extraS1, extraS2 []V) {
	s1Len := len(s1)
	s2Len := len(s2)

	visited := make([]bool, s2Len)
	for i := 0; i < s1Len; i++ {
		element := s1[i]
		found := false
		for j := 0; j < s2Len; j++ {
			if visited[j] {
				continue
			}
			if s2[j] == element {
				visited[j] = true
				found = true
				break
			}
		}
		if !found {
			extraS1 = append(extraS1, element)
		}
	}

	for j := 0; j < s2Len; j++ {
		if visited[j] {
			continue
		}
		extraS2 = append(extraS2, s2[j])
	}

	return
}

// Merge and dedupe multiple items
func Merge[V comparable](ss ...[]V) []V {
	var final []V
	for _, s := range ss {
		final = append(final, s...)
	}
	return Dedupe(final)
}

// Merge and dedupe multiple items into a
func MergeItems[V comparable](items ...V) []V {
	return Dedupe(items)
}

// InArray PHP对应的in_array函数
func InArray[T comparable](item T, arrayData []T) bool {
	dataLen := len(arrayData)
	if dataLen == 0 {
		return false
	}
	for i := 0; i < dataLen; i++ {
		if item == arrayData[i] {
			return true
		}
	}
	return false
}

// ArrayUnique 切片去重，目前仅支持int和string两种类型
func ArrayUnique[T comparable](arrayData []T) []T {
	dataLen := len(arrayData)
	if dataLen == 0 {
		return nil
	}
	results := make([]T, 0, dataLen)
	m := make(map[T]struct{}, dataLen)
	for i := 0; i < dataLen; i++ {
		if _, ok := m[arrayData[i]]; !ok {
			results = append(results, arrayData[i])
			m[arrayData[i]] = struct{}{}
		}
	}
	return results
}

// ArraySum 计算数组之和
func ArraySum[T constraints.Float | constraints.Integer](data []T) T {
	var total T
	for _, num := range data {
		total += num
	}
	return total
}

// ArrayChunk 整数版本的数组切割。将arrayData按照每个长度为length切割为子数组
func ArrayChunk[T any](arrayData []T, length int) [][]T {
	dataLen := len(arrayData)
	if length == 0 {
		return nil
	}
	var result [][]T
	for i := 0; i < dataLen; i += length {
		if length+i > dataLen {
			result = append(result, arrayData[i:])
		} else {
			result = append(result, arrayData[i:i+length])
		}
	}
	return result
}

// ArrayDiff 模拟PHP array_diff函数 计算差集
func ArrayDiff[T comparable](base []T, othersParams ...[]T) []T {
	if len(base) == 0 {
		return nil
	}
	if len(base) > 0 && len(othersParams) == 0 {
		return base
	}
	var tmp = make(map[T]int, len(base))
	for _, v := range base {
		tmp[v] = 1
	}
	for _, param := range othersParams {
		for _, arg := range param {
			if tmp[arg] != 0 {
				tmp[arg]++
			}
		}
	}
	var res = make([]T, 0, len(tmp))
	for k, v := range tmp {
		if v == 1 {
			res = append(res, k)
		}
	}
	return res
}

// ArrayIntersect 模拟PHP array_intersect函数 计算交集
func ArrayIntersect[T comparable](base []T, othersParams ...[]T) []T {
	if len(base) == 0 {
		return nil
	}
	if len(base) > 0 && len(othersParams) == 0 {
		return base
	}
	var tmp = make(map[T]int, len(base))
	for _, v := range base {
		tmp[v] = 1
	}
	for _, param := range othersParams {
		for _, arg := range param {
			if tmp[arg] != 0 {
				tmp[arg]++
			}
		}
	}
	min := len(othersParams) + 1
	var res = make([]T, 0, len(tmp))
	for k, v := range tmp {
		if v >= min {
			res = append(res, k)
		}
	}
	return res
}

// ArrayMerge PHP array_merge函数，数组合并
func ArrayMerge[T any](datas ...[]T) []T {
	dataLen := len(datas)
	if dataLen == 0 {
		return nil
	}
	total := 0
	for i := 0; i < dataLen; i++ {
		total += len(datas[i])
	}
	results := make([]T, 0, total)
	for i := 0; i < dataLen; i++ {
		results = append(results, datas[i]...)
	}
	return results
}

// ArraySearch 搜索arrayData里面是否有item，有返回对应的index，无返回-1，只返回首次
func ArraySearch[T comparable](item T, arrayData []T) int {
	dataLen := len(arrayData)
	if dataLen == 0 {
		return -1
	}
	for i := 0; i < dataLen; i++ {
		if arrayData[i] == item {
			return i
		}
	}
	return -1
}

// ArrayProduct 计算数组的各元素的乘积
func ArrayProduct[T constraints.Float | constraints.Integer](data []T) T {
	var result T = 1
	for _, num := range data {
		result *= num
	}
	return result
}

// ArrayWalk 遍历某一个数组切片，callback返回false则停止遍历
func ArrayWalk[T any](data []T, callback func(item T, index int) bool) {
	for i := 0; i < len(data); i++ {
		if !callback(data[i], i) {
			break
		}
	}
}

// Join 将一个int slice转为sep分割的字符串
func Join[T constraints.Float | constraints.Integer | ~string](a []T, sep string) string {
	var s strings.Builder
	sLen := len(a)
	for i, item := range a {
		s.WriteString(fmt.Sprintf("%v", item))
		if i < sLen-1 {
			s.WriteString(sep)
		}
	}
	return s.String()
}

// ArrayInt2ArrayString 将int slice转为string slice
func ArrayInt2ArrayString(in []int) []string {
	length := len(in)
	data := make([]string, length)
	for i := 0; i < length; i++ {
		data[i] = strconv.Itoa(in[i])
	}
	return data
}

// ArrayReverse 数组反转，需要传指针进来，仅支持slice，array请转为slice
func ArrayReverse[T any](array []T) {
	if len(array) == 0 {
		return
	}
	for i, j := 0, len(array)-1; i < j; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}
}
