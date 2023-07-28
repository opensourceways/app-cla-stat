package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/opensourceways/app-cla-stat/controllers:AuthController"] = append(beego.GlobalControllerRouter["github.com/opensourceways/app-cla-stat/controllers:AuthController"],
        beego.ControllerComments{
            Method: "Auth",
            Router: `/:platform`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/opensourceways/app-cla-stat/controllers:CorporationSigningController"] = append(beego.GlobalControllerRouter["github.com/opensourceways/app-cla-stat/controllers:CorporationSigningController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/:link_id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/opensourceways/app-cla-stat/controllers:EmployeeSigningController"] = append(beego.GlobalControllerRouter["github.com/opensourceways/app-cla-stat/controllers:EmployeeSigningController"],
        beego.ControllerComments{
            Method: "List",
            Router: `/:link_id/:signing_id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/opensourceways/app-cla-stat/controllers:IndividualSigningController"] = append(beego.GlobalControllerRouter["github.com/opensourceways/app-cla-stat/controllers:IndividualSigningController"],
        beego.ControllerComments{
            Method: "List",
            Router: `/:link_id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/opensourceways/app-cla-stat/controllers:LinkController"] = append(beego.GlobalControllerRouter["github.com/opensourceways/app-cla-stat/controllers:LinkController"],
        beego.ControllerComments{
            Method: "ListLinks",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
