package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["sabtu/controllers:AyuController"] = append(beego.GlobalControllerRouter["sabtu/controllers:AyuController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["sabtu/controllers:AyuController"] = append(beego.GlobalControllerRouter["sabtu/controllers:AyuController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["sabtu/controllers:AyuController"] = append(beego.GlobalControllerRouter["sabtu/controllers:AyuController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["sabtu/controllers:AyuController"] = append(beego.GlobalControllerRouter["sabtu/controllers:AyuController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["sabtu/controllers:AyuController"] = append(beego.GlobalControllerRouter["sabtu/controllers:AyuController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["sabtu/controllers:PesertaController"] = append(beego.GlobalControllerRouter["sabtu/controllers:PesertaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["sabtu/controllers:PesertaController"] = append(beego.GlobalControllerRouter["sabtu/controllers:PesertaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["sabtu/controllers:PesertaController"] = append(beego.GlobalControllerRouter["sabtu/controllers:PesertaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["sabtu/controllers:PesertaController"] = append(beego.GlobalControllerRouter["sabtu/controllers:PesertaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["sabtu/controllers:PesertaController"] = append(beego.GlobalControllerRouter["sabtu/controllers:PesertaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

}
