package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrite(t *testing.T) {
	source := strings.NewReader("Lorem Ipsum is simply dummy text of the printing and typesetting industry.")
	n := source.Len()

	t.Logf("length %v", n)

	target := bytes.NewBuffer(make([]byte, 0))
	err := WriteBuf(source, target)

	assert.Equal(t, err, nil)
	assert.Equal(t, target.Len(), n)
}
