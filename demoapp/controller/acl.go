package controller

import (
	"github.com/eaciit/knot/knot.v1"
	"github.com/eaciit/toolkit"
)

type ACL struct{}

func (o *ACL) Auth(wc *knot.WebContext) interface{} {
	wc.Config.OutputType = knot.OutputJson
	return toolkit.NewResult().SetData("This REST has been call")
}
