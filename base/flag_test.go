package base_test

import (
	"testing"

	"github.com/ebiiim/fantasy/base"
)

func TestAnyOf(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name  string
		input []base.Flag
		want  base.Flag
	}{
		{"0010", []base.Flag{0b0000, 0b0010}, 0b0010},
		{"1110", []base.Flag{0b1010, 0b1110}, 0b1110},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			got := base.AnyOf(c.input)
			if got != c.want {
				t.Errorf("want=%b but got=%b", c.want, got)
			}
		})
	}
}

func TestAllOf(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name  string
		input []base.Flag
		want  base.Flag
	}{
		{"0010", []base.Flag{0b1110, 0b0010}, 0b0010},
		{"0000", []base.Flag{0b0001, 0b1110}, 0b0000},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			got := base.AllOf(c.input)
			if got != c.want {
				t.Errorf("want=%b but got=%b", c.want, got)
			}
		})
	}
}

func TestHas(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name string
		f0   base.Flag
		has  []base.Flag
		want bool
	}{
		{"case1", 0b101010, []base.Flag{0b100000, 0b001000}, true},
		{"case2", 0b101010, []base.Flag{0b100000, 0b001000, 0b000010}, true},
		{"case3", 0b101010, []base.Flag{0b100000, 0b001000, 0b000001}, false},
		{"case4", 0b101010, []base.Flag{0b100000000000}, false},
		{"case5", 0b101010, []base.Flag{0b100000000010}, false},
		{"case6", 0b10000000101010, []base.Flag{0b000001}, false},
		{"case7", 0b10000000101010, []base.Flag{0b000001, 0b000100, 0b010000}, false},
		{"case8", 0b10000000101010, []base.Flag{0b000010}, true},
		{"case9", 0b10000000101010, []base.Flag{0b000010, 0b001000, 0b100000}, true},
		{"case10", 0b10000000101010, []base.Flag{0b000010, 0b001000, 0b100000, 0b1}, false},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			got := c.f0.Has(c.has...)
			if got != c.want {
				t.Errorf("f0=%b AnyOf(has)=%b", c.f0, base.AnyOf(c.has))
			}
		})
	}
}

func TestExcepts(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name string
		f0   base.Flag
		has  []base.Flag
		want bool
	}{
		{"case1", 0b101010, []base.Flag{0b010000, 0b000100}, true},
		{"case2", 0b101010, []base.Flag{0b010000, 0b000100, 0b000001}, true},
		{"case3", 0b101010, []base.Flag{0b100000}, false},
		{"case4", 0b101010, []base.Flag{0b100000000000}, true},
		{"case5", 0b101010, []base.Flag{0b100000000010}, false},
		{"case6", 0b10000000101010, []base.Flag{0b000001}, true},
		{"case7", 0b10000000101010, []base.Flag{0b000001, 0b000100, 0b010000}, true},
		{"case8", 0b10000000101010, []base.Flag{0b000001, 0b000100, 0b010000, 0b10}, false},
		{"case9", 0b10000000101010, []base.Flag{0b000010}, false},
		{"case10", 0b10000000101010, []base.Flag{0b000010, 0b001000, 0b100000}, false},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			got := c.f0.Excepts(c.has...)
			if got != c.want {
				t.Errorf("f0=%b ^AnyOf(has)=%b", c.f0, ^base.AnyOf(c.has))
			}
		})
	}
}
