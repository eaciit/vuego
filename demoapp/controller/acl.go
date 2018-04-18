package controller

import (
	"github.com/eaciit/knot/knot.v1"
	"github.com/eaciit/toolkit"
)

type ACL struct{}

func (o *ACL) Auth(wc *knot.WebContext) interface{} {
	wc.Config.OutputType = knot.OutputJson
	wc.SetSession("ecappsessionid", toolkit.RandomString(32))
	return toolkit.NewResult().SetData("This REST has been call")
}

func (o *ACL) Logout(wc *knot.WebContext) interface{} {
	wc.Config.OutputType = knot.OutputJson
	wc.SetSession("ecappsessionid", "")
	return nil
}
