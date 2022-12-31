package main

import (
	"embed"
	"log"
	"sync"
	"time"

	"github.com/sasha-s/go-deadlock"
	"github.com/spf13/viper"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	"mindmachine/messaging/eventcatcher"
	"mindmachine/mindmachine"
	"mindmachine/scumclass/eventbucket"
)

//go:embed frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

func main() {
	deadlock.Opts.DisableLockOrderDetection = true
	deadlock.Opts.DeadlockTimeout = time.Millisecond * 60000
	afterGui := beforeGui()
	// Create an instance of the app structure
	app := NewApp()
	eventBucket := &eventbucket.EventBucket{}

	// Create application with options
	err := wails.Run(&options.App{
		Title:             "Best Nostr Evuh",
		Width:             1366,
		Height:            768,
		DisableResize:     false,
		Fullscreen:        false,
		Frameless:         true,
		StartHidden:       false,
		HideWindowOnClose: false,
		BackgroundColour:  &options.RGBA{R: 255, G: 255, B: 255, A: 255},
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Menu:             nil,
		Logger:           nil,
		LogLevel:         logger.DEBUG,
		OnStartup:        app.startup,
		OnDomReady:       app.domReady,
		OnBeforeClose:    app.beforeClose,
		OnShutdown:       app.shutdown,
		CSSDragProperty:  "--wails-draggable",
		CSSDragValue:     "drag",
		WindowStartState: options.Normal,
		Bind: []interface{}{
			app,
			eventBucket,
		},
		// Windows platform specific options
		Windows: &windows.Options{
			WebviewIsTransparent: false,
			WindowIsTranslucent:  false,
			DisableWindowIcon:    false,
			// DisableFramelessWindowDecorations: false,
			WebviewUserDataPath: "",
		},
		// Mac platform specific options
		Mac: &mac.Options{
			TitleBar: &mac.TitleBar{
				TitlebarAppearsTransparent: false,
				HideTitle:                  false,
				HideTitleBar:               false,
				FullSizeContent:            false,
				UseToolbar:                 false,
				HideToolbarSeparator:       true,
			},
			Appearance:           mac.NSAppearanceNameDarkAqua,
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			About: &mac.AboutInfo{
				Title:   "Best Nostr",
				Message: "",
				Icon:    icon,
			},
		},
	})

	afterGui()

	if err != nil {
		log.Fatal(err)
	}

}

func beforeGui() func() {
	terminator := make(chan struct{})
	wg := &sync.WaitGroup{}
	// Various aspect of this application require global and local settings. To keep things
	// clean and tidy we put these settings in a Viper configuration.
	conf := viper.New()

	// Now we initialise this configuration with basic settings that are required on startup.
	mindmachine.InitConfig(conf)

	// make the config accessible globally
	mindmachine.SetConfig(conf)
	eventbucket.StartDb(terminator, wg)
	go eventcatcher.SubscribeToAllEvents(terminator)

	return func() {
		err := mindmachine.MakeOrGetConfig().WriteConfig()
		if err != nil {
			mindmachine.LogCLI(err.Error(), 3)
		}
		mindmachine.LogCLI("exiting", 3)
		close(terminator)
		wg.Wait()
		mindmachine.LogCLI("exited", 3)
	}
}