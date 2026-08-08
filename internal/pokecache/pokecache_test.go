package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	c := NewCache(time.Second * 10)

	if c.cache == nil {
		t.Error("Cache is nil")
	}
}

func TestAddAndGetCache(t *testing.T) {
	c := NewCache(time.Second * 10)
	inputKey := "key"
	inputVal := []byte("value")
	c.Add(inputKey, inputVal)
	actual, ok := c.Get(inputKey)

	if !ok {
		t.Errorf("%s not found", inputKey)
	}
	if string(actual) != string(inputVal) {
		t.Errorf("%v doesn't match %v", actual, inputVal)
	}
}

func TestReapLoop(t *testing.T) {
	interval := time.Duration(time.Millisecond * 10)
	c := NewCache(interval)
	inputKey := "key"
	inputVal := []byte("value")
	c.Add(inputKey, inputVal)

	time.Sleep(time.Duration(time.Millisecond + interval))

	_, ok := c.Get(inputKey)
	if ok {
		t.Errorf("%s should have been reaped", inputKey)
	}
}
