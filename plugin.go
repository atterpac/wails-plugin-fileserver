package fileserver

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/wailsapp/wails/v3/pkg/application"
)

// ---------------- Plugin Setup ----------------
// This is the main plugin struct. It can be named anything you like.
// It must implement the application.Plugin interface.
// Both the Init() and Shutdown() methods are called synchronously when the app starts and stops.
// Changing the name of this struct will change the name of the plugins class in the frontend
// Bound methods will exist inside frontend/bindings/github.com/user/fileserver under the name of the struct
type FileServer struct{
    // Plugin Options
    app *application.App
}

// Init is called when the app is starting up. You can use this to
// initialise any resources you need. You can also access the application
// instance via the app property.
func (p *FileServer) OnStartup() error {
    p.app = application.Get() // Allows for use of the application instance inside the plugin
	return nil
}

// Shutdown is called when the app is shutting down via runtime.Quit() call
// You can use this to clean up any resources you have allocated
func (p *FileServer) OnShutdown() error {
	return nil
}

// Name returns the name of the plugin.
// You should use the go module format e.g. github.com/myuser/myplugin
func (p *FileServer) Name() string {
	return "github.com/atterpac/wails-plugin-fileserver"
}

// ---------------- Plugin Methods ----------------
// Plugin methods are just normal Go methods. You can add as many as you like.
// The only requirement is that they are exported (start with a capital letter).
// You can also return any type that is JSON serializable.
// See https://golang.org/pkg/encoding/json/#Marshal for more information.

// implement http.Handler
func (p *FileServer) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	 var err error
    requestedFilename := strings.TrimPrefix(req.URL.Path, "/")
    println("Requesting file:", requestedFilename)
    fileData, err := os.ReadFile(requestedFilename)
    if err != nil {
        res.WriteHeader(http.StatusBadRequest)
        res.Write([]byte(fmt.Sprintf("Could not load file %s", requestedFilename)))
    }

    res.Write(fileData)
}
