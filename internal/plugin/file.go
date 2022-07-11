package plugin

import (
	"fmt"

	"github.com/SharkEzz/provisiond/pkg/deployment"
	"github.com/SharkEzz/provisiond/pkg/utils"
)

type fileStruct struct {
	Action  string
	Path    string
	Content string
}

// The File plugin allow easy interactions with files (create, is file exist, delete).
type File struct{}

func (s *File) Execute(ctx *deployment.JobContext, data any) (string, error) {
	fileData, err := utils.ConvertToType[fileStruct](data)
	if err != nil {
		return "", err
	}

	switch fileData.Action {
	case "create":
		_, err := ctx.ExecuteCommand(fmt.Sprintf("echo %s > %s", fileData.Content, fileData.Path))
		if err != nil {
			return "", err
		}
	case "exist":
		_, err := ctx.ExecuteCommand(fmt.Sprintf("test -e %s", fileData.Path))
		if err != nil {
			return "", fmt.Errorf("error: file %s does not exist", fileData.Path)
		}
	case "delete":
		_, err := ctx.ExecuteCommand(fmt.Sprintf("rm -f %s", fileData.Path))
		if err != nil {
			return "", err
		}
	default:
		return "", fmt.Errorf("invalid file action: '%s' expected one of: 'create' | 'exist' | 'delete'", fileData.Action)
	}

	return "", nil
}
