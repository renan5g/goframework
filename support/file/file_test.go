package file

import (
	"os"
	"testing"

	"github.com/renan5g/goframework/testing/file"

	"github.com/stretchr/testify/assert"
)

func TestClientOriginalExtension(t *testing.T) {
	assert.Equal(t, ClientOriginalExtension("logo.png"), "png")
}

func TestContain(t *testing.T) {
	assert.True(t, Contain("../constant.go", "const Version"))
}

func TestCreate(t *testing.T) {
	pwd, _ := os.Getwd()
	path := pwd + "/report/test.txt"
	assert.Nil(t, Create(path, `report`))
	assert.Equal(t, 1, file.GetLineNum(path))
	assert.True(t, Exists(path))
	assert.Nil(t, Remove(path))
	assert.Nil(t, Remove(pwd+"/report"))
}

func TestExists(t *testing.T) {
	assert.True(t, Exists("file.go"))
}

func TestExtension(t *testing.T) {
	extension, err := Extension("file.go")
	assert.Nil(t, err)
	assert.Equal(t, "txt", extension)
}

func TestLastModified(t *testing.T) {
	ti, err := LastModified("../../logo.png", "UTC")
	assert.Nil(t, err)
	assert.NotNil(t, ti)
}

func TestMimeType(t *testing.T) {
	mimeType, err := MimeType("../../logo.png")
	assert.Nil(t, err)
	assert.Equal(t, "image/png", mimeType)
}

func TestRemove(t *testing.T) {
	pwd, _ := os.Getwd()
	path := pwd + "/report/report.txt"
	assert.Nil(t, Create(path, `report`))

	assert.Nil(t, Remove(path))
	assert.Nil(t, Remove(pwd+"/report"))
}

func TestSize(t *testing.T) {
	size, err := Size("../../logo.png")
	assert.Nil(t, err)
	assert.Equal(t, int64(16438), size)
}
