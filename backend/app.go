package backend

import (
	"cat-clipboard/backend/config"
	"cat-clipboard/backend/controller"
	"cat-clipboard/backend/service"
	"container/list"
	"context"
	"fmt"
)

// App struct
type App struct {
	ctx                      context.Context
	beans                    []interface{}
	applicationConfigContext *config.ApplicationConfigContext
}

// NewApp creates a new App application struct
func NewApp() *App {
	a := &App{}
	a.applicationConfigContext = config.InitApplicationConfigContext()
	a.beans = initBeans(a)
	return a
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	l := list.New()
	l.PushBack("afe")

	return fmt.Sprintf("Hello %s, It's show time!.... %s", name, l.Front().Value)
}

func (a *App) GetCtx() context.Context {
	return a.ctx
}

func (a *App) GetBeans() []interface{} {
	return a.beans
}

func (a *App) GetApplicationConfigContext() *config.ApplicationConfigContext {
	return a.applicationConfigContext
}

/*
*
初始化bean
*/
func initBeans(a *App) []interface{} {

	configService := service.NewConfigService(a.applicationConfigContext)

	// Create an instance of the app structure
	configController := controller.NewConfigController(a.ctx, configService)
	contentController := controller.NewContentController()
	enumController := controller.NewEnumController()
	return []interface{}{
		a,
		configController,
		contentController,
		enumController,
	}
}
