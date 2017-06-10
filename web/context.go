package web

import (
	"html/template"

	"github.com/evepraisal/go-evepraisal"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

type Context struct {
	app         *evepraisal.App
	baseURL     string
	extraJS     string
	adBlock     string
	templates   map[string]*template.Template
	cookieStore *sessions.CookieStore
	etags       map[string]string
}

func NewContext(app *evepraisal.App, baseURL string, extraJS string, adBlock string) *Context {
	ctx := &Context{
		app:         app,
		baseURL:     baseURL,
		extraJS:     extraJS,
		adBlock:     adBlock,
		cookieStore: sessions.NewCookieStore(securecookie.GenerateRandomKey(32)),
	}
	ctx.GenerateStaticEtags()
	return ctx
}
