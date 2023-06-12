// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package spamfilter

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
	case "fortios:spamfilter/bwl:Bwl":
		r = &Bwl{}
	case "fortios:spamfilter/bword:Bword":
		r = &Bword{}
	case "fortios:spamfilter/dnsbl:Dnsbl":
		r = &Dnsbl{}
	case "fortios:spamfilter/fortishield:Fortishield":
		r = &Fortishield{}
	case "fortios:spamfilter/iptrust:Iptrust":
		r = &Iptrust{}
	case "fortios:spamfilter/mheader:Mheader":
		r = &Mheader{}
	case "fortios:spamfilter/options:Options":
		r = &Options{}
	case "fortios:spamfilter/profile:Profile":
		r = &Profile{}
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
		"spamfilter/bwl",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"spamfilter/bword",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"spamfilter/dnsbl",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"spamfilter/fortishield",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"spamfilter/iptrust",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"spamfilter/mheader",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"spamfilter/options",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"spamfilter/profile",
		&module{version},
	)
}