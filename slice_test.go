package main

import (
	"fmt"
	"reflect"
	"testing"
)

// 切片大小
func TestInfo(t *testing.T) {
	sli := []int{1, 2, 3}
	fmt.Printf("%T\n", sli) //类型
	printInfo(sli)

	sli = append(sli, 1)
	printInfo1(sli)
}

func printInfo(sli []int) {
	fmt.Println(cap(sli), len(sli)) //切片底层数组的大小, 切片大小
}

// 通过interface{} 和 reflect 实现泛型的功能
func printInfo1(sli interface{}) {
	if reflect.TypeOf(sli).Kind() != reflect.Slice {
		fmt.Printf("not slice but %T", sli)
	}
	sl := reflect.ValueOf(sli)
	fmt.Println(sl.Cap(), sl.Len()) //切片底层数组的大小, 切片大小
}

// 遍历
func TestTraverse(t *testing.T) {
	sli := []int{1, 2, 3}
	for i, v := range sli {
		fmt.Println(i, v)
	}
}

// 按索引取值
func TestIndex(t *testing.T) {
	sli := []int{1, 2, 3}
	v := sli[0]
	fmt.Println(v)
	v = sli[len(sli)-1]
	fmt.Println(v)
}

// 子切片
func TestSlice(t *testing.T) {
	sli := []int{1, 2, 3, 4}
	sliSub := sli[0:3]

	fmt.Printf("%p %p\n", sli, sliSub)   // 切片引用数组的第一个元素的内存地址（实际地址）
	fmt.Printf("%p %p\n", &sli, &sliSub) // 切片是引用，引用取地址是指针（指针地址）

	fmt.Println(sliSub)

	// 切片共享底层数组，子切片共享的数组片段从开头算起
	fmt.Println(cap(sli), cap(sliSub))
	fmt.Println(len(sli), len(sliSub))

	// 因为切片是对同一个底层数据的引用，所以修改子切片会影响到原切片
	sliSub[0] = 10
	fmt.Println(sli)
}

// 按索引更新
func TestUpdate(t *testing.T) {
	sli := []int{1, 2, 3}
	sli[1] = 10
	fmt.Println(sli)
}

// 末尾追加
func TestAppend(t *testing.T) {
	sli := []int{1, 2, 3}
	sli = append(sli, 4)
	fmt.Println(sli)
	sli = append(sli, []int{5, 6}...)
	fmt.Println(sli)
}

// 按索引添加
func TestInsert(t *testing.T) {
	sli := []int{1, 2, 3}

	rear := append([]int{}, sli[2:]...)
	sli = append(sli[:2], 10)
	sli = append(sli, rear...)
	fmt.Println(sli)
}

// 复制切片
func TestCopy(t *testing.T) {
	sli := []int{1, 2, 3, 4, 5}
	// 新切片要有足够的大小（len），而不是容量（cap）。len不够会丢失后面的数据
	sliCopy := make([]int, 5)
	copy(sliCopy, sli)
	fmt.Println(sli, sliCopy)
	fmt.Printf("%p, %p\n", sli, sliCopy)
}

// 按索引删除
func TestDelete(t *testing.T) {
	sli := []int{1, 2, 3, 4, 5}
	index := 2
	sli = append(sli[:index], sli[index+1:]...)
	fmt.Println(sli)
}

// 清空 方法1
func TestClear(t *testing.T) {
	sli := []int{1, 2, 3, 4, 5}
	fmt.Printf("%p, %p\n", sli, &sli)
	fmt.Printf("%v, %v\n", len(sli), cap(sli)) // 5, 5
	sli = sli[0:0]                             //切片的len变为0，相当于清空了切片
	fmt.Printf("%p, %p\n", sli, &sli)          //数组地址 和 指针地址都没有变
	fmt.Printf("%v, %v\n", len(sli), cap(sli)) // 0, 5
}

// 清空 方法2
func TestClear1(t *testing.T) {
	sli := []int{1, 2, 3, 4, 5}
	fmt.Printf("%p, %p\n", sli, &sli)
	fmt.Printf("%v, %v\n", len(sli), cap(sli)) // 5, 5
	sli = append([]int{})                      // 指针指向一个新分配的空数组
	fmt.Printf("%p, %p\n", sli, &sli)          //数组地址变了 指针地址没有变
	fmt.Printf("%v, %v\n", len(sli), cap(sli)) // 0, 0  切片的len变为0，相当于清空了切片
}

func TestMemory(t *testing.T) {
	//a := struct {
	//	Name string
	//}{Name: "cm"}
	//b := &a
	a := make(chan int)
	b := a

	// 引用类型的 实际地址 指针地址
	fmt.Printf("%p %p\n", a, &a)
	fmt.Printf("%p %p\n", b, &b)
}

func TestNilSlice(t *testing.T) {
	var sli []string
	fmt.Println(sli, sli == nil)
	_ = append(sli, "a")
	fmt.Println(sli)
	sli = append(sli, "b")
	fmt.Println(sli)
	_ = append(sli, "c")
	fmt.Println(sli)
}

func TestSliceCopy(t *testing.T) {
	type P struct {
		A map[string]string
	}
	sli := []P{{A: map[string]string{"a": "b"}}}
	sli2 := sli
	fmt.Printf("%p %p %p\n", sli, &sli, sli[0].A)
	fmt.Printf("%p %p %p\n", sli2, &sli2, sli2[0].A)
}

func TestSliceNil(t *testing.T) {
	var a []string
	fmt.Printf("%p %p\n", a, &a)
}

func TestSliceInsert(t *testing.T) {
	a := []int{1, 2, 3}
	a = append([]int{4, 5}, a...)

	fmt.Println(a)
}

func TestSliceDelete(t *testing.T) {
	data := DeleteSlice2([]int{1, 2, 3, -1, 0, 2, -3})
	fmt.Println(data)
}

func DeleteSlice2(a []int) []int {
	j := 0
	for _, val := range a {
		if val > 0 {
			a[j] = val
			j++
		}
	}
	return a[:j]
}

func TestNilSliceRange(t *testing.T) {
	var a []int
	for i, v := range a {
		fmt.Println(i, v)
	}
}

func TestSlice2(t *testing.T) {
	a := []int{1, 2, 3}
	fmt.Println(a[:0])
}

func TestSlice3(t *testing.T) {
	a := []int{1, 2, 3}
	b := make([]*int, 0)
	for _, v := range a {
		if v == 1 {
			b = append(b, &v) // b存储的是局部变量v的内存地址，最终v存储的是3
		}
	}
	for _, v := range b {
		fmt.Println(*v)
	}
}

func TestSlice4(t *testing.T) {
	a := []int{1, 2, 3}
	b := make([]int, 0)
	for _, v := range a {
		if v == 1 {
			b = append(b, v) // b存储的是局部变量v的拷贝，最终v存储的是3，和拷贝无关
		}
	}
	for _, v := range b {
		fmt.Println(v)
	}
}

func TestSlice5(t *testing.T) {
	a := []int{1, 2, 3}
	b := make([]*int, 0)
	for _, v := range a {
		if v == 1 {
			v1 := v            // 局部变量v的拷贝
			b = append(b, &v1) // b存储的是新局部变量v1的内存地址，每次都会创建
		}
	}
	for _, v := range b {
		fmt.Println(*v)
	}
}

func TestSlice6(t *testing.T) {
	fmt.Printf("%v", []int{1, 2})
}
