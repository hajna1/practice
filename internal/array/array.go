package array

import "fmt"

type Array struct {
	data []int
}

func (a *Array) At(index int) (int, error) {
	if index >= len(a.data) {
		return 0, fmt.Errorf("out of range")
	}
	return a.data[index], nil
}

func (a *Array) Insert(index int, values ...int) {
	temp := a.data[:index]
	temp = append(temp, values...)
	temp = append(temp, a.data[index:len(a.data)]...)
	a.data = temp
}

func (a *Array) Size() int {
	return len(a.data)
}

func (a *Array) Capacity() int {
	return cap(a.data)
}

func (a *Array) IsEmpty() bool {
	return len(a.data) == 0
}

func (a *Array) Reserve(capacity int) {
	temp := make([]int, 0, capacity)
	a.data = append(temp, a.data...)
}

func (a *Array) Push(items ...int) {
	a.data = append(a.data, items...)
}

func (a *Array) Remove(item int) {
	pos := 0
	for i := 0; i< len(a.data); i++ {
		a.data[pos] = a.data[i]
		if a.data[i] != item {
			pos+=1
		}
	}
	a.data = a.data[:pos]
}

func (a *Array) Delete(start, end int) error{
	size := end - start
	if size < 0 {
		return fmt.Errorf("bad args")
	}
	if size == 0 {
		return nil
	}
	for i := start; i < len(a.data)-size; i++ {
		a.data[i] = a.data[i + size]
	}
	a.data = a.data[:len(a.data)-size]
	return nil
}
