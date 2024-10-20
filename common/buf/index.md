buf 包主要用于管理字节切片的分配和释放，以优化内存使用和减少垃圾回收的开销。它通过实现一个自定义的分配器 defaultAllocator，提供了一种高效的方式来获取和释放不同大小的字节切片。

 Inspired by https://github.com/xtaci/smux/blob/master/alloc.go and 