package ttlmap

import (
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	m := NewTTLMap(5 * time.Second)

	if m == nil {
		t.Error("Expected a map and got nil")
	}
}

func TestSetAndGet(t *testing.T) {
	m := NewTTLMap(5 * time.Second)

	expected := "Simon"
	m.Set("name", expected)

	v, _ := m.Get("name")

	if v != expected {
		t.Errorf("Expected %s, got %s", expected, v)
	}
}
func TestExpiry(t *testing.T) {
	m := NewTTLMap(1 * time.Second)

	expected := "Simon"
	m.Set("name", expected)

	v, _ := m.Get("name")
	if v != expected {
		t.Errorf("Expected %s, got %s", expected, v)
	}

	time.Sleep(1 * time.Second)
	v, _ = m.Get("name")
	if v != nil {
		t.Errorf("Expected nil, got %s", v)
	}
}

func TestGetBeforeSet(t *testing.T) {
	m := NewTTLMap(1 * time.Second)

	v, _ := m.Get("name")
	if v != nil {
		t.Errorf("Expected nil, got %s", v)
	}
}
