package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/xxxcrel/kafka-console/pkg/kconsole"
)

//go:embed all:frontend
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "kafka-kconsole",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
		EnumBind: []interface{}{
			kconsole.AllSchemaType,
			kconsole.AllACLOperation,
			kconsole.AllACLPermissionType,
			kconsole.AllCompatibilityLevel,
			kconsole.AllACLResourceType,
			kconsole.AllMode,
			kconsole.AllSchemaRuleMode,
			kconsole.AllSchemaRuleKind,
			kconsole.AllACLResourcePatternType,
			kconsole.AllFrontendFormat,
			kconsole.AllConfigType,
			kconsole.AllIncrementalAlterConfigOp,
			kconsole.AllConfigResourceType,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
