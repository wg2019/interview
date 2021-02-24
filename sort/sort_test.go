package sort

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

var testCases = []struct {
	test   []int
	expect []int
}{
	{
		test:   []int{1, 3, 2, 4, 5, 7, 2, 8, 123, 64, 23, 58, 32, 12},
		expect: []int{1, 2, 2, 3, 4, 5, 7, 8, 12, 23, 32, 58, 64, 123},
	},
	{
		test:   []int{2, 3, 4, 5, 1, 2, 52, 234, 5, 123, 5, -1, 2, 1, 4, -4, 234, -13},
		expect: []int{-13, -4, -1, 1, 1, 2, 2, 2, 3, 4, 4, 5, 5, 5, 52, 123, 234, 234},
	},
}
var testList []int

func TestMain(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	testList = make([]int, 0)
	for i := 0; i < 10*0000; i++ {
		testList = append(testList, rand.Intn(10000))
	}
}
func TestBubbleSort(t *testing.T) {
	for _, tC := range testCases {
		BubbleSort(tC.test)
		if !reflect.DeepEqual(tC.test, tC.expect) {
			t.Fatalf("test: %+v, expect: %+v", tC.test, tC.expect)
		}
	}
}

func TestSelectSort(t *testing.T) {
	for _, tC := range testCases {
		SelectSort(tC.test)
		if !reflect.DeepEqual(tC.test, tC.expect) {
			t.Fatalf("test: %+v, expect: %+v", tC.test, tC.expect)
		}
	}
}

func TestInsertSort(t *testing.T) {
	for _, tC := range testCases {
		InsertSort(tC.test)
		if !reflect.DeepEqual(tC.test, tC.expect) {
			t.Fatalf("test: %+v, expect: %+v", tC.test, tC.expect)
		}
	}
}

func TestShellSort(t *testing.T) {
	for _, tC := range testCases {
		ShellSort(tC.test)
		if !reflect.DeepEqual(tC.test, tC.expect) {
			t.Fatalf("test: %+v, expect: %+v", tC.test, tC.expect)
		}
	}
}
func TestMergeSort(t *testing.T) {
	for _, tC := range testCases {
		result := MergeSort(tC.test)
		if !reflect.DeepEqual(result, tC.expect) {
			t.Fatalf("test: %+v, expect: %+v", result, tC.expect)
		}
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BubbleSort(testList)
	}
}

func BenchmarkSelectSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SelectSort(testList)
	}
}

func BenchmarkInsertSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertSort(testList)
	}
}

func BenchmarkShellSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ShellSort(testList)
	}
}

func BenchmarkMergeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MergeSort(testList)
	}
}
