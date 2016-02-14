package instrmtgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp2/instrument"
)

//NoRelatedSym is a repeating group in InstrmtGrp
type NoRelatedSym struct {
	//Instrument Component
	Instrument instrument.Component
}

//Component is a fix50sp2 InstrmtGrp Component
type Component struct {
	//NoRelatedSym is a non-required field for InstrmtGrp.
	NoRelatedSym []NoRelatedSym `fix:"146,omitempty"`
}

func New() *Component { return new(Component) }