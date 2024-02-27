package processors

import (
	"errors"

	"github.com/Parallels/prl-devops-service/basecontext"
	"github.com/Parallels/prl-devops-service/pdfile/diagnostics"
	"github.com/Parallels/prl-devops-service/pdfile/models"
)

type LocalPathCommandProcessor struct{}

func (p LocalPathCommandProcessor) Process(ctx basecontext.ApiContext, line string, dest *models.PDFile) (bool, *diagnostics.PDFileDiagnostics) {
	diag := diagnostics.NewPDFileDiagnostics()
	command := getCommand(line)
	if command == nil {
		return false, diag
	}
	if command.Command != "LOCAL_PATH" {
		return false, diag
	}
	if command.Argument == "" {
		diag.AddError(errors.New("run command is missing argument"))
	}

	dest.LocalPath = command.Argument
	ctx.LogDebugf("Processed by LocalPathCommandProcessor, line %v", line)
	return true, diag
}
