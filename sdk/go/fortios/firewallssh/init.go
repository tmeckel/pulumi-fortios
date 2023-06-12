// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firewallssh

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
	case "fortios:firewallssh/hostkey:Hostkey":
		r = &Hostkey{}
	case "fortios:firewallssh/localca:Localca":
		r = &Localca{}
	case "fortios:firewallssh/localkey:Localkey":
		r = &Localkey{}
	case "fortios:firewallssh/setting:Setting":
		r = &Setting{}
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
		"firewallssh/hostkey",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewallssh/localca",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewallssh/localkey",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"firewallssh/setting",
		&module{version},
	)
}