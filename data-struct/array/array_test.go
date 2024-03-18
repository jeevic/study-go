package array

import "testing"

func TestArray_Insert(t *testing.T) {
	capacity := uint(10)

	arr := New(capacity)

	for i := uint(0); i < capacity-5; i++ {
		err := arr.Insert(i, int(i+2))

		if nil != err {
			t.Fatal(err.Error())
		}
	}

	arr.Print()

	t.Log(arr.Len())

	err := arr.Insert(uint(5), 999)

	if nil != err {
		t.Fatal(err.Error())
	}
	arr.Print()

	arr.InsertToTail(666)

	arr.Print()

	arr.InsertToTail(777)

	arr.Print()

	arr.InsertToTail(777)

	arr.Print()
}

// 删除
func TestArray_Delete(t *testing.T) {
	capacity := uint(10)

	arr := New(capacity)

	for i := uint(0); i < capacity; i++ {
		err := arr.Insert(i, int(i+2))

		if nil != err {
			t.Fatal(err.Error())
		}
	}
	arr.Print()

	for i := 9; i >= 0; i-- {
		_, err := arr.Delete(uint(i))

		if nil != err {
			t.Fatal(err.Error())
		}
		arr.Print()
	}
}

// find
func TestArray_Find(t *testing.T) {
	capacity := uint(10)
	arr := New(capacity)

	for i := uint(0); i < capacity; i++ {
		err := arr.Insert(i, int(i+2))

		if nil != err {
			t.Fatal(err.Error())
		}
	}
	arr.Print()

	t.Log(arr.Find(1))
	t.Log(arr.Find(9))
	t.Log(arr.Find(11))

}
