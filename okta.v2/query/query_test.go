package query

import "testing"

func TestEmpty(t *testing.T) {
	p := NewQueryParams()
	qs := p.String()
	if qs != "" {
		t.Fatalf("expected empty string got '%s'", qs)
	}
}

func TestInt64(t *testing.T) {
	p := NewQueryParams(WithLimit(100))
	qs := p.String()
	if qs != `?limit=100` {
		t.Fatalf("expected '?limit=100' got '%s'", qs)
	}
}

func TestBool(t *testing.T) {
	p := NewQueryParams(WithActivate(false))
	qs := p.String()
	if qs != `?activate=false` {
		t.Fatalf("expected '?activate=false' got '%s'", qs)
	}

	p = NewQueryParams(WithActivate(true))
	qs = p.String()
	if qs != `?activate=true` {
		t.Fatalf("expected '?activate=true' got '%s'", qs)
	}
}

func TestString(t *testing.T) {
	p := NewQueryParams(WithQ(`x`))
	qs := p.String()
	if qs != "?q=x" {
		t.Fatalf("expected '?q=x' got '%s'", qs)
	}

	// Testing if we are applying URL encoding to a reserved char
	p = NewQueryParams(WithQ(`x=/`))
	qs = p.String()
	if qs != `?q=x%3D%2F` {
		t.Fatalf(`expected '%s' got '%s'`, `?q=x%3D%2F`, qs)
	}
}

func TestMultiple(t *testing.T) {
	p := NewQueryParams(WithQ(`x`), WithLimit(100), WithActivate(true))
	qs := p.String()
	if qs != "?activate=true&limit=100&q=x" {
		t.Fatalf("expected '?activate=true&limit=100&q=x' got '%s'", qs)
	}
}
