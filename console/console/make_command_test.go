package console

import (
	"testing"

	"github.com/stretchr/testify/assert"

	consolemocks "github.com/renan5g/goframework/mocks/console"
	"github.com/renan5g/goframework/support/file"
)

func TestMakeCommand(t *testing.T) {
	makeCommand := &MakeCommand{}
	mockContext := &consolemocks.Context{}
	mockContext.On("Argument", 0).Return("CleanCache").Once()
	assert.Nil(t, makeCommand.Handle(mockContext))
	assert.True(t, file.Exists("internal/console/commands/clean_cache.go"))

	mockContext.On("Argument", 0).Return("Goframework/CleanCache").Once()
	assert.Nil(t, makeCommand.Handle(mockContext))
	assert.True(t, file.Exists("internal/console/commands/goframework/clean_cache.go"))
	assert.True(t, file.Contain("internal/console/commands/goframework/clean_cache.go", "package goframework"))
	assert.True(t, file.Contain("internal/console/commands/goframework/clean_cache.go", "type CleanCache struct"))

	assert.Nil(t, file.Remove("internal"))
}
