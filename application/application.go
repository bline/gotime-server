package application

import (
	"github.com/carbocation/interpose"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/gorilla/sessions"
	"github.com/spf13/viper"
	"net/http"

	"github.com/bline/gotime-server/middleware"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/bline/gotime-server/api/proto"
	"github.com/bline/gotime-server/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"github.com/bline/gotime-server/handler"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"context"
	"github.com/gorilla/mux"
)

// New is the constructor for Application struct.
func New(config *viper.Viper) (*Application, error) {
	dsn := config.Get("dsn").(string)

	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	cookieStoreSecret1 := config.GetString("cookie_secret1")
	cookieStoreSecret2 := config.GetString("cookie_secret2")

	app := &Application{}
	app.config = config
	app.dsn = dsn
	app.db = db
	app.sessionStore = sessions.NewCookieStore([]byte(cookieStoreSecret1), []byte(cookieStoreSecret2))
	app.mux = app.Mux()
	return app, nil
}

// Application is the application object that runs HTTP server.
type Application struct {
	config       *viper.Viper
	dsn          string
	db           *gorm.DB
	mux          *gin.Engine
}

func (app *Application) Mux() *gin.Engine {
	return gin.Default()
}

func (app *Application) MiddlewareStruct() (*interpose.Middleware, error) {
	r := app.mux
	r.Use(middleware.SetDB(app.db))
	r.Use(middleware.Session(r, app.config))
	r.POST("/idtoken", handler.PostIdToken)
	r.Use(middleware.MustLogin)
	r.Use(GrpcHandlerFunc(app))
	return r, nil
}

func GrpcAuthInterceptor(c context.Context, r interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(c)

	if ok {
		token = md[]
		user, err = service.UserFromIDToken()
	}
}

func GrpcHandlerFunc(app *Application) gin.HandlerFunc {
	var opts []grpc.ServerOption
	opts = append(opts, grpc.UnaryInterceptor(func(c context.Context, r interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		newCtx := context.WithValue(c, "config", app.config)
		newCtx = context.WithValue(newCtx, "db", app.db)

	}))
	defaultMux := app.mux
	server := grpc.NewServer()
	api.RegisterTimeSheetServer(server, service.NewTimeSheetService(app.db, app.config))
	api.RegisterAccountsServer(server, service.NewAccountsService(app.db, app.config))

	wrapper := grpcweb.WrapServer(server)

	return gin.HandlerFunc(func(ctx *gin.Context) {
		w, r := ctx.Writer, ctx.Request
		if wrapper.IsGrpcWebRequest(r) {
			wrapper.ServeHTTP(w, r)
		} else {
			defaultMux.ServeHTTP(w, r)
		}
	})
}
