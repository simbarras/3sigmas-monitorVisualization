package storer

import (
	"github.com/simbarras/3sigmas-monitorVisualization/pkg/data"
	"sync"
	"testing"
)

func TestNewInfluxStorer(t *testing.T) {
	env := data.Env{
		InfluxUrl:   "http://localhost:8086",
		InfluxToken: "token",
	}
	s := NewInfluxStorer(env)
	if s != nil {
		t.Errorf("No connection should be established with dummy configuration")
	}
}

func TestInfluxStorer_Store(t *testing.T) {
	s := InfluxStorer{
		mu:           sync.Mutex{},
		bucketPrefix: "prefix",
	}
	_, err := s.Store("project", "source", nil)
	if err == nil {
		t.Errorf("No store should be done with dummy configuration")
	}
}

func TestInfluxStorer_setBucket(t *testing.T) {
	s := InfluxStorer{
		mu: sync.Mutex{},
	}
	bucket := s.setBucket("test")
	if bucket != nil {
		t.Errorf("No bucket should be created with dummy configuration")
	}
}
