package httpMagic

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExpBackOff_Execute(t *testing.T) {
	f := func(req HttpRequest) (string, error) {
		return "success", nil
	}

	eb := &ExpBackOff{}

	result, error := eb.Execute(f, HttpRequest{})

	assert.Nil(t, error)
	assert.Equal(t, "success", result)
}

func TestExpBackOff_Execute_Should_Retry(t *testing.T) {
	f := func(req HttpRequest) (string, error) {
		return "", errors.New("failed")
	}

	//todo - could use a timer to prove that it actually retried N times
	eb := &ExpBackOff{}

	result, error := eb.Execute(f, HttpRequest{})

	assert.Empty(t, result)
	assert.Equal(t, errors.New("failed"), error)
}
