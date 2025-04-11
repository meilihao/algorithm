package list

import (
	"fmt"
	"reflect"
	"unsafe"
)

// [double circular linked list, 双向循环链表](https://my.oschina.net/u/4857646/blog/5036189)
type DCL struct { // DoubleCircularLink
	Prev *DCL
	Next *DCL
}

// 初始化链表时，只有一个链表节点，这个节点的前驱和后继节点都是自身
func DCLListInit(list *DCL) {
	list.Prev = list
	list.Next = list
}

func DCLListIsEmpty(list *DCL) bool {
	return list.Next == list
}

func DCLListAdd(list, n *DCL) {
	n.Next = list.Next
	n.Prev = list

	list.Next.Prev = n
	list.Next = n
}

// 在list.Next后面插入n
func DCLListTailInsert(list, n *DCL) {
	DCLListAdd(list.Next, n)
}

func DCLListDelete(n *DCL) {
	n.Next.Prev = n.Prev
	n.Prev.Next = n.Next

	n.Prev = nil
	n.Next = nil
}

// 将链表节点n从所在的双向链表中删除, 并把删除后的节点重新初始化为一个新的双向链表
func DCLListDeleteInit(n *DCL) {
	n.Next.Prev = n.Prev
	n.Prev.Next = n.Next

	DCLListInit(n)
}

func DCLListForEach(item, list *DCL) {
	item = list.Next

	for item != list {
		item = item.Next
	}
}

// 这个参数表示遍历到的双向链表节点的下一个节点. 用于安全删除，如果删除遍历到的item, 不影响继续遍历
func DCLListForEachSafe(item, next, list *DCL) {
	item = list.Next
	next = item.Next

	for item != list { // 删除item, `nil != list` is true
		item = next
		next = next.Next
	}
}

// 获取双向链表节点所在的业务结构体的内存地址
func DCLListEntry(item *DCL, name string) interface{} {
	if item == nil || name == "" {
		return nil
	}

	offset := StructFieldOffest(item, name)
	if offset == -1 {
		return nil
	}

	return uintptr(unsafe.Pointer(item)) - uintptr(offset)
}

func DCLListForEachEntry(item, list *DCL, name string) {
	if item == nil || name == "" {
		return
	}

	offset := StructFieldOffest(item, name)
	if offset == -1 {
		return
	}

	item = list.Next

	for item != list {
		item = item.Next

		// v := (*[2]byte)(unsafe.Pointer(uintptr(unsafe.Pointer(item)) - uintptr(offset)))
	}
}

func DCLListForEachEntrySafe(item, next, list *DCL, name string) {
	if item == nil || name == "" {
		return
	}

	offset := StructFieldOffest(item, name)
	if offset == -1 {
		return
	}

	item = list.Next
	next = item.Next

	for item != list {
		item = next
		next = next.Next

		// v := (*[2]byte)(unsafe.Pointer(uintptr(unsafe.Pointer(item)) - uintptr(offset)))
	}
}

func DCLListForEachEntryHook(item, list *DCL, name string, hook interface{}) {
	if item == nil || name == "" {
		return
	}

	offset := StructFieldOffest(item, name)
	if offset == -1 {
		return
	}
	typ := reflect.TypeOf(hook)
	if typ.Kind() != reflect.Func {
		panic(fmt.Errorf("%v isn't func", hook))
	}

	item = list.Next

	for item != list {
		item = item.Next

		// v := (*[2]byte)(unsafe.Pointer(uintptr(unsafe.Pointer(item)) - uintptr(offset)))
		// hook
	}
}

func StructFieldOffest(v interface{}, fieldName string) int {
	typ := reflect.TypeOf(v)
	if typ.Kind() != reflect.Struct {
		panic(fmt.Errorf("%v isn't struct", v))
	}

	if field, ok := typ.FieldByName(fieldName); ok {
		return int(field.Offset)
	}

	return -1
}
