package plugin

import (
	"fmt"
	"os"
	goPlugin "plugin"
	"strings"

	"github.com/SharkEzz/provisiond/pkg/context"
	"github.com/SharkEzz/provisiond/pkg/plugin/internal"
)

// The Plugin interface define one method,
// which execute the content of the associated content in the deployment file.
// It return the command stdout output.
type Plugin interface {
	Execute(data any, ctx *context.JobContext) (string, error)
}

// The Plugins map contain all the registered plugins,
// whoses names must be the same as the ones used in the deployment file.
var Plugins = map[string]Plugin{
	"shell": &internal.Shell{},
}

func init() {
	pluginsDir, err := os.ReadDir("./plugins")
	if err != nil {
		panic(err)
	}

	for _, pluginItem := range pluginsDir {
		pluginName := strings.Split(pluginItem.Name(), ".")[0]
		loadedPlugin, err := loadPlugin(fmt.Sprintf("./plugins/%s", pluginItem.Name()))
		if err != nil {
			panic(err)
		}
		Plugins[pluginName] = loadedPlugin
	}

	fmt.Printf("Loaded %d plugins\n", len(pluginsDir))
}

func loadPlugin(path string) (Plugin, error) {
	pluginFile, err := goPlugin.Open(path)
	if err != nil {
		return nil, err
	}

	GetPlugin, err := pluginFile.Lookup("GetPlugin")
	if err != nil {
		return nil, err
	}

	loadedPlugin := GetPlugin.(func() any)()

	return loadedPlugin.(Plugin), nil
}
