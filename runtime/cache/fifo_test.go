package cache

import (
	"testing"
)

type testFifoObject struct {
	name string
	val  interface{}
}

func testFifoObjectKeyFunc(obj interface{}) (string, error) {
	return obj.(testFifoObject).name, nil
}

func TestFIFO_requeueOnPop(t *testing.T) {
	f := NewFIFO(testFifoObjectKeyFunc)
	var testData = []testFifoObject{
		{name: "test", val: 10},
		{name: "test2", val: 10},
		{name: "test3", val: 10},
		{name: "tes4", val: 10},
	}
	for _, test := range testData {
		f.Add(test)
	}
	for i := 0; i < len(testData); i++ {
		t.Logf("keys : %v", f.ListKeys())
		data := Pop(f)
		t.Logf("pop from fifo : %v\n", data)
	}
	item, ok, err := f.Get(testFifoObject{name: "test", val: 10})
	t.Log(item, ok, err)
}
