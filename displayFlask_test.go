package main

import (
	. "github.com/stretchr/testify/assert"
	"testing" 
)

func TestSpace(t *testing.T)  {
	zeroSpaces:= space(0)
	Equal(t, 0, len(zeroSpaces))
	Equal(t, "", zeroSpaces)
	
	fiveSpaces:= space(5)
	Equal(t, 5, len(fiveSpaces))
	Equal(t, "     ", fiveSpaces)
}