## 包含包
* arrutil
  1. ArrayInt2ArrayString 将int slice转为string slice
  2. ArrayReverse 数组反转，需要传指针进来，仅支持slice，array请转为slice
  3. Join 将一个int slice转为sep分割的字符串
  4. ArrayWalk 遍历某一个数组切片，callback返回false则停止遍历 
  5. ArrayProduct 计算数组的各元素的乘积
  6. ArraySearch 搜索arrayData里面是否有item，有返回对应的index，无返回-1，只返回首次
  7. ArrayMerge 数组合并
  8. ArrayIntersect 模拟PHP array_intersect函数 计算交集
  9. ArrayDiff 模拟PHP array_diff函数 计算差集
  10. ArrayChunk 整数版本的数组切割。将arrayData按照每个长度为length切割为子数组
  11. ArraySum 计算数组之和
  12. ArrayUnique 切片去重，目前仅支持int和string两种类型
  13. InArray PHP对应的in_array函数
  14. Merge and dedupe multiple items into a
  15. Merge and dedupe multiple items
  16. Diff calculates the extra elements between two sequences
  17. ContainsItems checks if s1 contains s2
  18. PickRandom item from a slice of elements

* cdncheck
* cryptoutil 
* executil 
* fileutil 
* folderutil 
* fsync 
* goevent 
* iputil 
* mapsutil 
* permission 
* ratelimit 
* sizewaitgroup
* stringsutil

## 常用包
* go验证器 github.com/asaskevich/govalidator
* cidr切分 github.com/menglh/mapcidr
* 自动增加当前程序的最大文件描述符数量 github.com/projectdiscovery/fdmax
* go调用lua github.com/yuin/gopher-lua
* go运行时 github.com/arl/statsviz
* 日志库 github.com/menglh/golog

