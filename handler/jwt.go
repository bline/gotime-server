package handler

import (
    "github.com/bline/gotime-server/service"
    "github.com/bline/gotime-server/libhttp"
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/sessions"
    "net/http"
    "github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)


func PostIdToken(c *gin.Context) {
	r, w := c.Request, c.Writer
	r.ParseForm()

    db := r.Context().Value("db").(*gorm.DB)
	config := r.Context().Value("config").(*viper.Viper)
	session := sessions.Default()

    idToken := r.FormValue("idtoken")
    user, err := service.UserFromIDToken(db, config, idToken)
    if err != nil {
        libhttp.HandleErrorJson(w, err)
		return
    }
    session.Set("user", user)

	err = session.Save()
	if err != nil {
		libhttp.HandleErrorJson(w, err)
		return
	}
	libhttp.HandleSuccessJson(w, "login success")
}

func GetLogout(w http.ResponseWriter, r *http.Request) {

    sessionStore := r.Context().Value( "sessionStore").(sessions.Store)

    session, _ := sessionStore.Get(r, "gotime-server-session")

    delete(session.Values, "user")
    session.Save(r, w)

	libhttp.HandleSuccessJson(w, "logout success")
}
