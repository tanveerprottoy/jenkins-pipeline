package server

import (
	"log"
	"net/http"

	resourcecfg "github.com/tanveerprottoy/jenkins-pipeline/service/internal/server/resource/config"
	"github.com/tanveerprottoy/jenkins-pipeline/service/internal/server/router"
	"github.com/tanveerprottoy/jenkins-pipeline/service/pkg/constant"
)

// App struct
type App struct {
	router      *router.Router
	resourceCfg *resourcecfg.Config
}

func NewApp() *App {
	a := new(App)
	a.initComponents()
	return a
}

func (a *App) initModules() {
	a.resourceCfg = resourcecfg.NewConfig()
}

func (a *App) initModuleRouters() {
	router.RegisterResourceRoutes(a.router, constant.V1, a.resourceCfg.Handler)
}

// Init app
func (a *App) initComponents() {
	a.router = router.NewRouter()
	a.initModules()
	a.initModuleRouters()
}

// Run app
func (a *App) Run() {
	err := http.ListenAndServe(
		":8080",
		a.router.Mux,
	)
	if err != nil {
		log.Fatal(err)
	}
}
