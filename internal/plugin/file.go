package plugin

import (
	"fmt"
	"os"

	"github.com/SharkEzz/provisiond/pkg/context"
	"github.com/SharkEzz/provisiond/pkg/utils"
)

type fileStruct struct {
	Action  string
	Path    string
	Content string
}

// The File plugin allow easy interactions with files (create, exist, delete).
type File struct{}

func (s *File) Execute(data any, ctx *context.JobContext) (string, error) {
	fileData, err := utils.ConvertToType[fileStruct](data)
	if err != nil {
		return "", err
	}

	switch fileData.Action {
	case "create":
		err := os.WriteFile(fileData.Path, []byte(fileData.Content), 0644)
		if err != nil {
			return "", err
		}
	case "exist":
		if _, err := os.Stat(fileData.Path); os.IsNotExist(err) {
			return "", fmt.Errorf("error: file %s does not exist, aborting", fileData.Path)
		}
	case "delete":
		err := os.Remove(fileData.Path)
		if err != nil {
			return "", err
		}
	default:
		return "", fmt.Errorf("invalid file action: '%s' expected one of: 'create' | 'exist' | 'delete'", fileData.Action)
	}

	return "", nil
}
