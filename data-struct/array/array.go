package array

import (
	"fmt"

	"github.com/pkg/errors"
)

type Array struct {
	data   []int
	length uint
}

// New 初始化数组
func New(capacity uint) *Array {
	if capacity == 0 {
		return nil
	}

	return &Array{
		data:   make([]int, capacity, capacity),
		length: 0,
	}
}

// Len 返回数组的长度
func (this *Array) Len() uint {
	return this.length
}

func (this *Array) isIndexOutOfRange(index uint) bool {
	if index >= this.length {
		return true
	}
	return false
}

// Find 查找数组的值 index为索引
func (this *Array) Find(index uint) (int, error) {
	if this.isIndexOutOfRange(index) {
		return 0, errors.New("out of index range")
	}
	return this.data[index], nil
}

// Insert 数组中插入值
func (this *Array) Insert(index uint, v int) error {

	if int(this.Len()) == cap(this.data) {
		return errors.New("full array")
	}

	if index != this.length && this.isIndexOutOfRange(index) {
		return errors.New("out of range")
	}

	//移动位置
	for i := this.length; i > index; i-- {
		this.data[i+1] = this.data[i]
	}

	this.data[index] = v
	this.length++

	return nil
}

// InsertToTail 插入到结尾
func (this *Array) InsertToTail(v int) error {
	return this.Insert(this.Len(), v)
}

// Delete 删除
func (this *Array) Delete(index uint) (int, error) {
	if this.isIndexOutOfRange(index) {
		return 0, errors.New("out of index range")
	}

	v := this.data[index]

	for i := index; i < this.length-1; i++ {
		this.data[index] = this.data[index+1]
	}
	this.length--
	return v, nil
}

// Print 打印数组
func (this *Array) Print() {
	var format string

	for i := uint(0); i < this.length; i++ {
		format += fmt.Sprintf("|%+v", this.data[i])
	}

	fmt.Println(format)

}
