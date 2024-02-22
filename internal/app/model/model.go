package model

import (
	"fmt"
)

type Friend struct {
	FirstName  string
	FamilyName string
	Age        int
}

func (f *Friend) String() string {
	return fmt.Sprintf("%s %s (%d)", f.FirstName, f.FamilyName, f.Age)
}
