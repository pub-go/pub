package assert

import (
	"reflect"
	"strings"
	"testing"

	"code.gopub.tech/pub/util/reflects"
)

func True(t *testing.T, got bool, msg ...string) {
	t.Helper()
	if !got {
		t.Errorf("[%v] assert True failed: %v\n got=false, want=true",
			t.Name(), strings.Join(msg, ". "))
	}
}

func False(t *testing.T, got bool, msg ...string) {
	t.Helper()
	if got {
		t.Errorf("[%v] assert False failed: %v\n got=true, want=false",
			t.Name(), strings.Join(msg, ". "))
	}
}

func Nil(t *testing.T, got interface{}, msg ...string) {
	t.Helper()
	if !reflects.IsNil(got) {
		t.Errorf("[%v] assert Nil failed: %v\n got=%#v, want=nil",
			t.Name(), strings.Join(msg, ". "), got)
	}
}

func NotNil(t *testing.T, got interface{}, msg ...string) {
	t.Helper()
	if reflects.IsNil(got) {
		t.Errorf("[%v] assert NotNil failed: %v\n got=nil, want=not-nil",
			t.Name(), strings.Join(msg, ". "))
	}
}

func Equal(t *testing.T, got, want interface{}, msg ...string) {
	t.Helper()
	if got != want {
		t.Errorf("[%v] assert Equal failed: %v\n got=%#v\nwant=%#v",
			t.Name(), strings.Join(msg, ". "), got, want)
	}
}

func DeepEqual(t *testing.T, got, want interface{}, msg ...string) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("[%v] assert DeepEqual failed: %v\n got=%#v\nwant=%#v",
			t.Name(), strings.Join(msg, ". "), got, want)
	}
}

func ShouldPanic(t *testing.T, fn func()) {
	t.Helper()
	defer func() {
		if recover() == nil {
			t.Errorf("expected panic")
		}
	}()
	fn()
}

func ShouldNotPanic(t *testing.T, fn func()) {
	t.Helper()
	defer func() {
		if recover() != nil {
			t.Errorf("unexpected panic")
		}
	}()
	fn()
}
