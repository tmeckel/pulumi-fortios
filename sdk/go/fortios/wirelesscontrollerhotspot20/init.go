// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package wirelesscontrollerhotspot20

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "fortios:wirelesscontrollerhotspot20/anqp3gppcellular:Anqp3gppcellular":
		r = &Anqp3gppcellular{}
	case "fortios:wirelesscontrollerhotspot20/anqpipaddresstype:Anqpipaddresstype":
		r = &Anqpipaddresstype{}
	case "fortios:wirelesscontrollerhotspot20/anqpnairealm:Anqpnairealm":
		r = &Anqpnairealm{}
	case "fortios:wirelesscontrollerhotspot20/anqpnetworkauthtype:Anqpnetworkauthtype":
		r = &Anqpnetworkauthtype{}
	case "fortios:wirelesscontrollerhotspot20/anqproamingconsortium:Anqproamingconsortium":
		r = &Anqproamingconsortium{}
	case "fortios:wirelesscontrollerhotspot20/anqpvenuename:Anqpvenuename":
		r = &Anqpvenuename{}
	case "fortios:wirelesscontrollerhotspot20/anqpvenueurl:Anqpvenueurl":
		r = &Anqpvenueurl{}
	case "fortios:wirelesscontrollerhotspot20/h2qpadviceofcharge:H2qpadviceofcharge":
		r = &H2qpadviceofcharge{}
	case "fortios:wirelesscontrollerhotspot20/h2qpconncapability:H2qpconncapability":
		r = &H2qpconncapability{}
	case "fortios:wirelesscontrollerhotspot20/h2qpoperatorname:H2qpoperatorname":
		r = &H2qpoperatorname{}
	case "fortios:wirelesscontrollerhotspot20/h2qposuprovider:H2qposuprovider":
		r = &H2qposuprovider{}
	case "fortios:wirelesscontrollerhotspot20/h2qposuprovidernai:H2qposuprovidernai":
		r = &H2qposuprovidernai{}
	case "fortios:wirelesscontrollerhotspot20/h2qptermsandconditions:H2qptermsandconditions":
		r = &H2qptermsandconditions{}
	case "fortios:wirelesscontrollerhotspot20/h2qpwanmetric:H2qpwanmetric":
		r = &H2qpwanmetric{}
	case "fortios:wirelesscontrollerhotspot20/hsprofile:Hsprofile":
		r = &Hsprofile{}
	case "fortios:wirelesscontrollerhotspot20/icon:Icon":
		r = &Icon{}
	case "fortios:wirelesscontrollerhotspot20/qosmap:Qosmap":
		r = &Qosmap{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := fortios.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"fortios",
		"wirelesscontrollerhotspot20/anqp3gppcellular",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"wirelesscontrollerhotspot20/anqpipaddresstype",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"wirelesscontrollerhotspot20/anqpnairealm",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"wirelesscontrollerhotspot20/anqpnetworkauthtype",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"wirelesscontrollerhotspot20/anqproamingconsortium",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"wirelesscontrollerhotspot20/anqpvenuename",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"wirelesscontrollerhotspot20/anqpvenueurl",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"wirelesscontrollerhotspot20/h2qpadviceofcharge",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"wirelesscontrollerhotspot20/h2qpconncapability",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"wirelesscontrollerhotspot20/h2qpoperatorname",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"wirelesscontrollerhotspot20/h2qposuprovider",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"wirelesscontrollerhotspot20/h2qposuprovidernai",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"wirelesscontrollerhotspot20/h2qptermsandconditions",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"wirelesscontrollerhotspot20/h2qpwanmetric",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"wirelesscontrollerhotspot20/hsprofile",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"wirelesscontrollerhotspot20/icon",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"wirelesscontrollerhotspot20/qosmap",
		&module{version},
	)
}
