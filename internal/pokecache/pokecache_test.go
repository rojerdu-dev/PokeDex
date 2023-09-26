package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	t.Parallel()
	c := NewCache(time.Millisecond)
	if c.cache == nil {
		t.Errorf("Expected cache to be initialized")
	}
}

func TestAddGetCache(t *testing.T) {
	t.Parallel()
	c := NewCache(time.Millisecond)

	cases := []struct {
		inputKey string
		inputVal []byte
	}{
		{
			inputKey: "key1",
			inputVal: []byte("val1"),
		},
		{
			inputKey: "key2",
			inputVal: []byte("val2"),
		},
		{
			inputKey: "key3",
			inputVal: []byte("val3"),
		},
	}

	for _, cs := range cases {
		c.Add(cs.inputKey, cs.inputVal)
		got, ok := c.Get(cs.inputKey)
		if !ok {
			t.Errorf("Expected key %s to be in cache", cs.inputKey)
		}
		if string(got) != string(cs.inputVal) {
			t.Errorf("Expected value %s, got %s", cs.inputVal, got)
		}
	}

}

func TestReap(t *testing.T) {
	t.Parallel()
	interval := time.Millisecond * 10
	c := NewCache(interval)

	keyOne := "key1"
	c.Add(keyOne, []byte("val1"))

	time.Sleep(interval + time.Millisecond)

	_, ok := c.Get(keyOne)
	if ok {
		t.Errorf("%s should have been reaped", keyOne)
	}
}

func TestReapFail(t *testing.T) {
	t.Parallel()
	interval := time.Millisecond * 10
	c := NewCache(interval)

	keyOne := "key1"
	c.Add(keyOne, []byte("val1"))

	time.Sleep(interval / 2)

	_, ok := c.Get(keyOne)
	if !ok {
		t.Errorf("%s should not have been reaped", keyOne)
	}
}
