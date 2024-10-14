package resource

import (
	"context"
	"testing"
)

func TestService(t *testing.T) {
	// test service
	s := NewService()
	e, err := s.GetData(context.Background())
	if err != nil {
		t.Fatalf("DoSomething() returned error: %s", err)
	}
	if &e == nil {
		t.Logf("expected val, got nil: %d", &e)
		return
	}
	t.Logf("val returned, val: %d", e)
}
