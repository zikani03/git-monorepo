package main

import (
	"fmt"
	"os"

	"github.com/alecthomas/kong"
	gitMonorepo "github.com/zikani03/git-monorepo"
)

var version string = "0.0.0"

type Globals struct {
	Config  string      `help:"Location of configuration file" default:"monorepo.toml" type:"path"`
	Debug   bool        `help:"Enable debug mode"`
	Version VersionFlag `name:"version" help:"Show version and quit"`
}

type VersionFlag string

func (v VersionFlag) Decode(ctx *kong.DecodeContext) error { return nil }
func (v VersionFlag) IsBool() bool                         { return true }
func (v VersionFlag) BeforeApply(app *kong.Kong, vars kong.Vars) error {
	fmt.Println(vars["version"])
	app.Exit(0)
	return nil
}

type InitCmd struct {
	Mangle          bool     `help:"Combine files from repos in one directory (not recommended!)"`
	PreserveHistory bool     `help:"Preserve history from the repos"`
	InMemory        bool     `help:"Clone repos in-memory instead of a temporary directory on file system"`
	TargetDir       string   `name:"target" help:"The target directory to create repo in. Must not exist" type:"path"`
	Sources         []string `help:"Source repositories with support for 'git-down' shortcuts"`
}

// CheckIfError should be used to naively panics if an error is not nil.
func CheckIfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}

func (r *InitCmd) Run(globals *Globals) error {
	repo := gitMonorepo.NewMonorepoFromSources(r.Sources)
	if r.InMemory {
		repo = gitMonorepo.NewMonorepoFromSourcesInMemory(r.Sources)
	}
	err := repo.Init(r.TargetDir)
	CheckIfError(err)

	return nil
}

// MergeCmd allows users to merge other repos into existing repositories
type MergeCmd struct {
	PreserveHistory bool     `help:"Preserve history from the repos"`
	TargetDir       string   `name:"target" type:"path" help:"The target repository. Must be a valid repo"`
	Sources         []string `help:"Source repositories with support for 'git-down' shortcuts"`
}

func (r *MergeCmd) Run(globals *Globals) error {
	repo := gitMonorepo.NewMonorepoFromSources(r.Sources)
	err := repo.MergeSourcesInto(r.TargetDir)
	CheckIfError(err)
	return nil
}

type CLI struct {
	Globals

	Init  InitCmd  `cmd:"" help:"Initialize a monorepo from source repos"`
	Merge MergeCmd `cmd:"" help:"Merge existing repos into another to form a monorepo"`
}

func main() {
	cli := CLI{
		Globals: Globals{
			Version: VersionFlag(version),
		},
	}

	ctx := kong.Parse(&cli,
		kong.Name("git-monorepo"),
		kong.Description("Make monorepos from existing repositories"),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
		}),
		kong.Vars{
			"version": version,
		})

	err := ctx.Run(&cli.Globals)
	ctx.FatalIfErrorf(err)
}
