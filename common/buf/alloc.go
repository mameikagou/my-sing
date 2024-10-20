package buf

import (
	"errors"
	"math/bits"
	"sync"
)

// Inspired by https://github.com/xtaci/smux/blob/master/alloc.go

// 待阅读: 字节的sync.Pool, 性能优化的神器

type Allocator interface {
	Get(size int) []byte
	Put(buf []byte) error
}

type defaultAllocator struct {
	buffers [11]sync.Pool
}

// defaultAllocator要先实现Allocator接口的Get和Put方法, 才能赋值给Allocator接口
// 它使用 sync.Pool 来缓存和重用字节切片，从而减少内存分配和垃圾回收的开销
func newDefaultAllocator() Allocator {
	return &defaultAllocator{
		buffers: [...]sync.Pool{ // 64B -> 64K
			{New: func() any { return new([1 << 6]byte) }},
			{New: func() any { return new([1 << 7]byte) }},
			{New: func() any { return new([1 << 8]byte) }},
			{New: func() any { return new([1 << 9]byte) }},
			{New: func() any { return new([1 << 10]byte) }},
			{New: func() any { return new([1 << 11]byte) }},
			{New: func() any { return new([1 << 12]byte) }},
			{New: func() any { return new([1 << 13]byte) }},
			{New: func() any { return new([1 << 14]byte) }},
			{New: func() any { return new([1 << 15]byte) }},
			{New: func() any { return new([1 << 16]byte) }},
		},
	}
}

func (alloc *defaultAllocator) Get(size int) []byte {
	if size <= 0 || size > 65536 { // 64kb
		return nil
	}

	var index uint16
	if size > 64 {
		index = msb(size)     // 找到size的最高有效位
		if size != 1<<index { // 如果size不是2的幂次方
			index += 1 // 则向上取整
		}
		// 这一步的目的是将 index 调整到适合用于 sync.Pool 数组的索引范围。具体来说，sync.Pool 数组的大小范围是从 64 字节到 64KB，对应的索引范围是从 0 到 10。
		// 通过减去 6，可以将 index 映射到这个范围。例如，如果 index 是 7，减去 6 后变为 1，这样可以对应到 sync.Pool 数组中的第二个池（64 字节到 128 字节）。
		index -= 6
	}

	buffer := alloc.buffers[index].Get()
	switch index {
	case 0:
		return buffer.(*[1 << 6]byte)[:size] //(*[1 << 6]byte)是类型断言, 将buffer转换为[1 << 6]byte类型;长度为64字节
	case 1:
		return buffer.(*[1 << 7]byte)[:size]
	case 2:
		return buffer.(*[1 << 8]byte)[:size]
	case 3:
		return buffer.(*[1 << 9]byte)[:size]
	case 4:
		return buffer.(*[1 << 10]byte)[:size]
	case 5:
		return buffer.(*[1 << 11]byte)[:size]
	case 6:
		return buffer.(*[1 << 12]byte)[:size]
	case 7:
		return buffer.(*[1 << 13]byte)[:size]
	case 8:
		return buffer.(*[1 << 14]byte)[:size]
	case 9:
		return buffer.(*[1 << 15]byte)[:size]
	case 10:
		return buffer.(*[1 << 16]byte)[:size]
	default:
		panic("invalid pool index")
	}
}

// Put returns a []byte to pool for future use,
// which the cap must be exactly 2^n
func (alloc *defaultAllocator) Put(buf []byte) error {
	bits := msb(cap(buf))
	if cap(buf) == 0 || cap(buf) > 65536 || cap(buf) != 1<<bits {
		return errors.New("allocator Put() incorrect buffer size")
	}
	bits -= 6
	buf = buf[:cap(buf)]

	//nolint
	//lint:ignore SA6002 ignore temporarily
	switch bits {
	case 0:
		alloc.buffers[bits].Put((*[1 << 6]byte)(buf))
	case 1:
		alloc.buffers[bits].Put((*[1 << 7]byte)(buf))
	case 2:
		alloc.buffers[bits].Put((*[1 << 8]byte)(buf))
	case 3:
		alloc.buffers[bits].Put((*[1 << 9]byte)(buf))
	case 4:
		alloc.buffers[bits].Put((*[1 << 10]byte)(buf))
	case 5:
		alloc.buffers[bits].Put((*[1 << 11]byte)(buf))
	case 6:
		alloc.buffers[bits].Put((*[1 << 12]byte)(buf))
	case 7:
		alloc.buffers[bits].Put((*[1 << 13]byte)(buf))
	case 8:
		alloc.buffers[bits].Put((*[1 << 14]byte)(buf))
	case 9:
		alloc.buffers[bits].Put((*[1 << 15]byte)(buf))
	case 10:
		alloc.buffers[bits].Put((*[1 << 16]byte)(buf))
	default:
		panic("invalid pool index")
	}
	return nil
}

func msb(size int) uint16 {
	return uint16(bits.Len32(uint32(size)) - 1)
}
