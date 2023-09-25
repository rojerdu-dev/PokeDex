package pokecache

import "testing"

func TestCreateCache(t *testing.T) {
	c := NewCache()
	if c.cache == nil {
		t.Errorf("Expected cache to be initialized")
	}
}

func TestAddGetCache(t *testing.T) {
	c := NewCache()

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
