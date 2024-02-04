package console

import (
	"testing"

	consolemocks "github.com/renan5g/goframework/mocks/console"
	"github.com/renan5g/goframework/support/file"

	"github.com/stretchr/testify/assert"
)

func TestRuleMakeCommand(t *testing.T) {
	requestMakeCommand := &RuleMakeCommand{}
	mockContext := &consolemocks.Context{}
	mockContext.On("Argument", 0).Return("").Once()
	err := requestMakeCommand.Handle(mockContext)
	assert.EqualError(t, err, "Not enough arguments (missing: name) ")

	mockContext.On("Argument", 0).Return("Uppercase").Once()
	err = requestMakeCommand.Handle(mockContext)
	assert.Nil(t, err)
	assert.True(t, file.Exists("internal/application/rules/uppercase.go"))

	mockContext.On("Argument", 0).Return("User/Phone").Once()
	err = requestMakeCommand.Handle(mockContext)
	assert.Nil(t, err)
	assert.True(t, file.Exists("internal/application/rules/user/phone.go"))
	assert.True(t, file.Contain("internal/application/rules/user/phone.go", "package user"))
	assert.True(t, file.Contain("internal/application/rules/user/phone.go", "type Phone struct"))
	assert.Nil(t, file.Remove("internal"))
}
