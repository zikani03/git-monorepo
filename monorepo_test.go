package monorepo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeUrlFromSpec(t *testing.T) {

	assert := assert.New(t)

	assert.Equal("https://github.com/zikani03/git-monorepo", makeFullUrl("gh:zikani03/git-monorepo"))
	assert.Equal("https://bitbucket.org/zikani03/git-monorepo", makeFullUrl("bb:zikani03/git-monorepo"))
	assert.Equal("https://gitlab.com/zikani03/git-monorepo", makeFullUrl("gl:zikani03/git-monorepo"))

	assert.Equal("somethingelse:zikani03/git-monorepo", makeFullUrl("somethingelse:zikani03/git-monorepo"))
}
