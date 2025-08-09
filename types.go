package main

var commitTypes = []string{
	"feat",
	"fix",
	"docs",
	"style",
	"refactor",
	"perf",
	"test",
	"build",
	"ci",
	"chore",
}

var commitTypeDescriptions = map[string]string{
	"feat":     "A new feature",
	"fix":      "A bug fix",
	"docs":     "Documentation only changes",
	"style":    "Changes that do not affect the meaning of the code",
	"refactor": "A code change that neither fixes a bug nor adds a feature",
	"perf":     "A code change that improves performance",
	"test":     "Adding missing tests or correcting existing tests",
	"build":    "Changes that affect the build system or external dependencies",
	"ci":       "Changes to CI configuration files and scripts",
	"chore":    "Other changes that don't modify src or test files",
}
