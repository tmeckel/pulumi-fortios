// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package switchcontrollerinitialconfig

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
	case "fortios:switchcontrollerinitialconfig/template:Template":
		r = &Template{}
	case "fortios:switchcontrollerinitialconfig/vlans:Vlans":
		r = &Vlans{}
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
		"switchcontrollerinitialconfig/template",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"switchcontrollerinitialconfig/vlans",
		&module{version},
	)
}
