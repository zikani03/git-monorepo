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

func TestSourceRepoName(t *testing.T) {
	var testCases = []struct {
		want string
		repo *sourceRepo
	}{
		{want: "example", repo: &sourceRepo{url: "https://github.com/user/example.git"}},
		{want: "example", repo: &sourceRepo{url: "https://github.com/user/example"}},
		{want: "example", repo: &sourceRepo{url: "gh:user/example"}},
		{want: "example", repo: &sourceRepo{url: "bb:user/example"}},
		{want: "example", repo: &sourceRepo{url: "gl:user/example"}},
	}

	for _, tt := range testCases {
		assert.Equal(t, tt.want, tt.repo.Name())
	}
}

func TestTempDirs(t *testing.T) {
	t.Helper()

	tempDirs := make(TempDirs, 0)
	// TODO: implement test for tempDirs
	// Schedule cleanup for later.
	t.Cleanup(func() {
		tempDirs.Clean()
	})
}
