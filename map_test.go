package omap_test

import (
	omap "github.com/RealFax/order-map"
	"strconv"
	"testing"
)

var (
	m = omap.New[string, string]()
)

func TestOrderedMap_Load(t *testing.T) {
	m.Store("Hello", "123")
	m.Store("Hello1", "Bonjour")

	t.Log(m.Load("Hello"))
	t.Log(m.Load("Hello1"))
}

func TestOrderedMap_LoadOrStore(t *testing.T) {
	m.Store("Key1", "Value1")
	t.Log(m.LoadOrStore("Key1", "Value2"))
	t.Log(m.LoadOrStore("Key2", "Value3"))
	t.Log(m.Load("Key1"))
	t.Log(m.Load("Key2"))
}

func TestOrderedMap_LoadAndDelete(t *testing.T) {
	m.Store("Key10", "Value10")
	t.Log(m.LoadAndDelete("Key10"))
	t.Log(m.Load("Key10"))
}

func TestOrderedMap_Delete(t *testing.T) {
	m.Store("Key20", "Value20")
	m.Delete("Key20")
	t.Log(m.Load("Key20"))
}

func TestOrderedMap_Swap(t *testing.T) {
	m.Store("Key30", "Value30")
	t.Log(m.Load("Key30"))
	t.Log(m.Swap("Key30", "Value31"))
	t.Log(m.Swap("Key31", "Value32"))
	t.Log(m.Load("Key30"))
	t.Log(m.Load("Key31"))
}

func TestOrderedMap_CompareAndSwap(t *testing.T) {
	m.Store("Key40", "Value40")
	t.Log(m.CompareAndSwap("Key40", "Value40_", "Value41"))
	t.Log(m.Load("Key40"))
	t.Log(m.CompareAndSwap("Key40", "Value40", "Value41"))
	t.Log(m.Load("Key40"))

	t.Log(m.CompareAndSwap("Key41", "", "Value42"))
	t.Log(m.Load("Key41"))

}

func TestOrderedMap_CompareAndDelete(t *testing.T) {
	m.Store("Key50", "Value50")
	m.CompareAndDelete("Key50", "Value50_")
	t.Log(m.Load("Key50"))
	m.CompareAndDelete("Key50", "Value50")
	t.Log(m.Load("Key50"))
}

func TestOrderedMap_Range(t *testing.T) {
	nm := omap.New[int, string]()
	for i := 0; i < 100; i++ {
		nm.Store(i, "VALUE_"+strconv.Itoa(i))
	}

	nm.Range(func(key int, value string) bool {
		// t.Log(key, value)
		return true
	})
}
