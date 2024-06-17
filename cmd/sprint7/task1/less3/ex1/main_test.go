package ex1_test

import (
	"strings"
	"testing"
)

type Modifier interface {
	Modify() string
}

type Original struct {
	Value string
}

func (o *Original) Modify() string {
	return o.Value
}

// Upper возвращает строку в верхнем регистре.
type Upper struct {
	modifier Modifier
}

func (u *Upper) Modify() string {
	toModifyValue := u.modifier.Modify()
	return strings.ToUpper(toModifyValue)
}

// Replace заменяет строки old на new.
type Replace struct {
	modifier Modifier
	old      string
	new      string
}

func (r *Replace) Modify() string {
	toModifyValue := r.modifier.Modify()
	return strings.ReplaceAll(toModifyValue, r.old, r.new)
}

// добавьте метод Modify для *Replace
// он должен заменять old на new
// ...

func TestModifier(t *testing.T) {
	original := &Original{Value: "Привет, гофер!"}
	replace := &Replace{
		modifier: original,
		old:      "гофер",
		new:      "мир",
	}
	upper := &Upper{
		modifier: replace,
	}
	if upper.Modify() != "ПРИВЕТ, МИР!" {
		t.Errorf(`get %s`, upper.Modify())
	}
}
