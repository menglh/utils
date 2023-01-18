
#### fsync包保持同步文件和目录
```go
s := NewSyncer()
s.SyncTo("dst", "src/a", "src/c")
```

