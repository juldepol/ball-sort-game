package main

import (
	. "github.com/stretchr/testify/assert"
	"testing" 
)

func TestIsCorrectInput(t *testing.T)  {
	is1Between0And2:= isCorrectInput(1,2)
	Equal(t, true, is1Between0And2)

	is0Between0And2:= isCorrectInput(0,2)
	Equal(t, false, is0Between0And2)

	is3Between0And2:= isCorrectInput(3,2)
	Equal(t, false, is3Between0And2)

	is2Between0And2:= isCorrectInput(2,2)
	Equal(t, false, is2Between0And2)

	isMinus1Between0And2:= isCorrectInput(-1,2)
	Equal(t, false, isMinus1Between0And2)
}