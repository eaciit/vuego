package main

import (
	"os"
	"path/filepath"
	"strings"

	knot "github.com/eaciit/knot/knot.v1"
	"github.com/eaciit/toolkit"

	"github.com/eaciit/clit"
)

var (
	err error
	log *toolkit.LogEngine
)

func main() {
	log, _ = toolkit.NewLog(true, false, "", "", "")

	clit.SetFlag("host", "0.0.0.0:9000", "host address of server")
	clit.LoadConfigFromFlag("", "", "../config/app.json")
	if err = clit.Commit(); err != nil {
		log.Error(err.Error())
		os.Exit(0)
	}

	host := clit.Value("host", "default", "")
	log.Infof("Serving on host: %s", host)

	functions := map[string]knot.FnContent{
		"/": VuePage("", "vuepage.html"),
	}
	knot.StartAppWithFn(
		App("demo", clit.Config("", "workingdir", clit.ExeDir()).(string)),
		host, functions)
}

func App(name, wd string, controllers ...interface{}) *knot.App {
	app := knot.NewApp(name)
	app.Static("asset", filepath.Join(wd, "asset"))
	app.Static("dist", filepath.Join(wd, "dist"))
	app.Static("view", filepath.Join(wd, "view"))
	app.ViewsPath = filepath.Join(wd, "view")

	app.SetValidation(true, func(k *knot.WebContext) bool {
		session := k.Session("ecappsessionid")
		if session == nil {
			k.SetData("validationresult", "nosession")
			return false
		} else if session == "" {
			k.SetData("validationresult", "noaccess")
			return false
		}
		return true
	}, func(k *knot.WebContext) string {
		vr := k.Data("validationresult", "").(string)
		if vr == "nosession" {
			return "/login"
		} else {
			return "/noaccess"
		}
	})

	for _, c := range controllers {
		if err := app.Register(c); err != nil {
			log.Errorf("controller registration error. %s", err.Error())
			os.Exit(0)
		}
	}

	return app
}

func VuePage(layout, view string) knot.FnContent {
	return func(c *knot.WebContext) interface{} {
		uri := c.Request.URL
		module := strings.Trim(uri.Path, " ")
		if module == "" || module == "/" {
			module = "home"
		}

		c.Config.OutputType = knot.OutputTemplate
		c.Config.ViewName = view
		c.Config.LayoutTemplate = layout

		data := toolkit.M{}
		data.Set("module", module)
		return data
	}
}
