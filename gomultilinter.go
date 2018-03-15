package main

import (
	"context"
	"fmt"
	"go/token"

	"github.com/liut0/gomultilinter/api"
)

type malignedLinter struct {
}

var LinterFactory api.LinterFactory = &malignedLinter{}

func (l *malignedLinter) NewLinterConfig() api.LinterConfig {
	return &malignedLinter{}
}

func (l *malignedLinter) NewLinter() (api.Linter, error) {
	return l, nil
}

func (*malignedLinter) Name() string {
	return "maligned"
}

func (l *malignedLinter) LintFile(ctx context.Context, file *api.File, reporter api.IssueReporter) error {
	malignFile(file.ASTFile, file.FSet, file.PkgInfo, func(pos token.Position, size, optimal int64) {
		reporter.Report(&api.Issue{
			Position: pos,
			Severity: api.SeverityWarning,
			Message:  fmt.Sprintf("struct of size %d could be %d", size, optimal),
			Category: "maligned",
		})
	})
	return nil
}
