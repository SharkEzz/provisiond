package plugin

import (
	"fmt"
	"os"
	goPlugin "plugin"
	"strings"

	"github.com/SharkEzz/provisiond/internal/plugin"
	"github.com/SharkEzz/provisiond/pkg/context"
	"github.com/SharkEzz/provisiond/pkg/logging"
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
	"shell": &plugin.Shell{},
}

// Load all the plugins in ./plugins (relative to the current executable directory)
func init() {
	// Do not attempt to load plugins if the plugins folder does not exist
	if _, err := os.Stat("./plugins"); os.IsNotExist(err) {
		return
	}

	pluginsDir, err := os.ReadDir("./plugins")
	if err != nil {
		panic(err)
	}

	loadCount := 0

	for _, pluginItem := range pluginsDir {
		pluginName := strings.Split(pluginItem.Name(), ".")[0]

		if !strings.HasSuffix(pluginItem.Name(), ".so") {
			continue
		}

		loadedPlugin, err := loadPlugin(fmt.Sprintf("./plugins/%s", pluginItem.Name()))
		if err != nil {
			panic(err)
		}
		Plugins[pluginName] = loadedPlugin
		loadCount++
	}

	logging.LogOut(fmt.Sprintf("Loaded %d external plugins, %d plugins in total", loadCount, len(Plugins)))
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
