// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package switchcontroller

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
	case "fortios:switchcontroller/customcommand:Customcommand":
		r = &Customcommand{}
	case "fortios:switchcontroller/dynamicportpolicy:Dynamicportpolicy":
		r = &Dynamicportpolicy{}
	case "fortios:switchcontroller/flowtracking:Flowtracking":
		r = &Flowtracking{}
	case "fortios:switchcontroller/fortilinksettings:Fortilinksettings":
		r = &Fortilinksettings{}
	case "fortios:switchcontroller/global:Global":
		r = &Global{}
	case "fortios:switchcontroller/igmpsnooping:Igmpsnooping":
		r = &Igmpsnooping{}
	case "fortios:switchcontroller/lldpprofile:Lldpprofile":
		r = &Lldpprofile{}
	case "fortios:switchcontroller/lldpsettings:Lldpsettings":
		r = &Lldpsettings{}
	case "fortios:switchcontroller/location:Location":
		r = &Location{}
	case "fortios:switchcontroller/macsyncsettings:Macsyncsettings":
		r = &Macsyncsettings{}
	case "fortios:switchcontroller/managedswitch:Managedswitch":
		r = &Managedswitch{}
	case "fortios:switchcontroller/nacdevice:Nacdevice":
		r = &Nacdevice{}
	case "fortios:switchcontroller/nacsettings:Nacsettings":
		r = &Nacsettings{}
	case "fortios:switchcontroller/networkmonitorsettings:Networkmonitorsettings":
		r = &Networkmonitorsettings{}
	case "fortios:switchcontroller/portpolicy:Portpolicy":
		r = &Portpolicy{}
	case "fortios:switchcontroller/quarantine:Quarantine":
		r = &Quarantine{}
	case "fortios:switchcontroller/remotelog:Remotelog":
		r = &Remotelog{}
	case "fortios:switchcontroller/settings8021X:Settings8021X":
		r = &Settings8021X{}
	case "fortios:switchcontroller/sflow:Sflow":
		r = &Sflow{}
	case "fortios:switchcontroller/snmpcommunity:Snmpcommunity":
		r = &Snmpcommunity{}
	case "fortios:switchcontroller/snmpsysinfo:Snmpsysinfo":
		r = &Snmpsysinfo{}
	case "fortios:switchcontroller/snmptrapthreshold:Snmptrapthreshold":
		r = &Snmptrapthreshold{}
	case "fortios:switchcontroller/snmpuser:Snmpuser":
		r = &Snmpuser{}
	case "fortios:switchcontroller/stormcontrol:Stormcontrol":
		r = &Stormcontrol{}
	case "fortios:switchcontroller/stormcontrolpolicy:Stormcontrolpolicy":
		r = &Stormcontrolpolicy{}
	case "fortios:switchcontroller/stpinstance:Stpinstance":
		r = &Stpinstance{}
	case "fortios:switchcontroller/stpsettings:Stpsettings":
		r = &Stpsettings{}
	case "fortios:switchcontroller/switchgroup:Switchgroup":
		r = &Switchgroup{}
	case "fortios:switchcontroller/switchinterfacetag:Switchinterfacetag":
		r = &Switchinterfacetag{}
	case "fortios:switchcontroller/switchlog:Switchlog":
		r = &Switchlog{}
	case "fortios:switchcontroller/switchprofile:Switchprofile":
		r = &Switchprofile{}
	case "fortios:switchcontroller/system:System":
		r = &System{}
	case "fortios:switchcontroller/trafficpolicy:Trafficpolicy":
		r = &Trafficpolicy{}
	case "fortios:switchcontroller/trafficsniffer:Trafficsniffer":
		r = &Trafficsniffer{}
	case "fortios:switchcontroller/virtualportpool:Virtualportpool":
		r = &Virtualportpool{}
	case "fortios:switchcontroller/vlan:Vlan":
		r = &Vlan{}
	case "fortios:switchcontroller/vlanpolicy:Vlanpolicy":
		r = &Vlanpolicy{}
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
		"switchcontroller/customcommand",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"switchcontroller/dynamicportpolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"switchcontroller/flowtracking",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"switchcontroller/fortilinksettings",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"switchcontroller/global",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"switchcontroller/igmpsnooping",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"switchcontroller/lldpprofile",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"switchcontroller/lldpsettings",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"switchcontroller/location",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"switchcontroller/macsyncsettings",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"switchcontroller/managedswitch",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"switchcontroller/nacdevice",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"switchcontroller/nacsettings",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"switchcontroller/networkmonitorsettings",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"switchcontroller/portpolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"switchcontroller/quarantine",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"switchcontroller/remotelog",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"switchcontroller/settings8021X",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"switchcontroller/sflow",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"switchcontroller/snmpcommunity",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"switchcontroller/snmpsysinfo",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"switchcontroller/snmptrapthreshold",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"switchcontroller/snmpuser",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"switchcontroller/stormcontrol",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"switchcontroller/stormcontrolpolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"switchcontroller/stpinstance",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"switchcontroller/stpsettings",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"switchcontroller/switchgroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"switchcontroller/switchinterfacetag",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"switchcontroller/switchlog",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"switchcontroller/switchprofile",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"switchcontroller/system",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"switchcontroller/trafficpolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"switchcontroller/trafficsniffer",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"switchcontroller/virtualportpool",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"switchcontroller/vlan",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"switchcontroller/vlanpolicy",
		&module{version},
	)
}
