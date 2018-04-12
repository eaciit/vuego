package main

import (
	"os"
	"path/filepath"
	"strings"

	knot "github.com/eaciit/knot/knot.v1"
	"github.com/eaciit/toolkit"

	controller "eaciit/vuego/demoapp/controller"

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
		"/": VuePage("_simple.html", "vuepage.html"),
	}
	knot.StartAppWithFn(
		App("demo",
			strings.Replace(clit.Config("", "workingdir", clit.ExeDir()).(string), "{exeDir}", clit.ExeDir(), -1),
			new(controller.ACL)),
		host, functions)
}

func App(name, wd string, controllers ...interface{}) *knot.App {

	// App initialization
	app := knot.NewApp(name)
	app.Static("asset", filepath.Join(wd, "asset"))
	app.Static("dist", filepath.Join(wd, "dist"))
	app.Static("view", filepath.Join(wd, "view"))
	app.ViewsPath = filepath.Join(wd, "view")

	// Change below to true to use validation, and ensure login and noaccess has no-validation
	app.SetValidation(false, func(k *knot.WebContext) bool {
		url := k.Request.URL
		if url.Path == "/login" || url.Path == "/noaccess" {
			return true
		}

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

	// controller registration
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
		module := strings.Trim(uri.Path, " ")[1:]
		if module == "" {
			module = "home"
		}

		c.Config.OutputType = knot.OutputTemplate
		c.Config.ViewName = view
		c.Config.LayoutTemplate = layout

		data := toolkit.M{}
		data.Set("module", module+"-main.js?"+toolkit.RandomString(32))
		return data
	}
}
