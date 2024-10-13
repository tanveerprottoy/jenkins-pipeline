package server

import (
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/tanveerprottoy/basic-go-server/internal/pkg/constant"
	"github.com/tanveerprottoy/basic-go-server/internal/pkg/router"
	resourcecfg "github.com/tanveerprottoy/basic-go-server/internal/server/resource/config"
	"github.com/tanveerprottoy/basic-go-server/pkg/data/sqlxpkg"
	"github.com/tanveerprottoy/basic-go-server/pkg/validatorpkg"
)

// App struct
type App struct {
	DBClient    *sqlxpkg.Client
	router      *router.Router
	resourceCfg *resourcecfg.Config
	Validate    *validator.Validate
}

func NewApp() *App {
	a := new(App)
	a.initComponents()
	return a
}

func (a *App) initDB() {
	a.DBClient = sqlxpkg.GetInstance()
}

func (a *App) initModules() {
	a.resourceCfg = resourcecfg.NewConfig(a.DBClient.DB, a.Validate)
}

func (a *App) initModuleRouters() {
	router.RegisterUserRoutes(a.router, constant.V1, a.resourceCfg.Handler)
}

func (a *App) initValidators() {
	a.Validate = validator.New()
	_ = a.Validate.RegisterValidation("notempty", validatorpkg.NotEmpty)
}

// Init app
func (a *App) initComponents() {
	a.initDB()
	a.router = router.NewRouter()
	a.initModules()
	a.initModuleRouters()
	a.initValidators()
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
