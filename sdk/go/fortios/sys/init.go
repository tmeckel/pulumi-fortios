// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sys

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
	case "fortios:sys/accprofile:Accprofile":
		r = &Accprofile{}
	case "fortios:sys/acme:Acme":
		r = &Acme{}
	case "fortios:sys/admin:Admin":
		r = &Admin{}
	case "fortios:sys/adminAdministrator:AdminAdministrator":
		r = &AdminAdministrator{}
	case "fortios:sys/adminProfiles:AdminProfiles":
		r = &AdminProfiles{}
	case "fortios:sys/affinityinterrupt:Affinityinterrupt":
		r = &Affinityinterrupt{}
	case "fortios:sys/affinitypacketredistribution:Affinitypacketredistribution":
		r = &Affinitypacketredistribution{}
	case "fortios:sys/alarm:Alarm":
		r = &Alarm{}
	case "fortios:sys/alias:Alias":
		r = &Alias{}
	case "fortios:sys/apiuser:Apiuser":
		r = &Apiuser{}
	case "fortios:sys/apiuserSetting:ApiuserSetting":
		r = &ApiuserSetting{}
	case "fortios:sys/arptable:Arptable":
		r = &Arptable{}
	case "fortios:sys/autoinstall:Autoinstall":
		r = &Autoinstall{}
	case "fortios:sys/automationaction:Automationaction":
		r = &Automationaction{}
	case "fortios:sys/automationdestination:Automationdestination":
		r = &Automationdestination{}
	case "fortios:sys/automationstitch:Automationstitch":
		r = &Automationstitch{}
	case "fortios:sys/automationtrigger:Automationtrigger":
		r = &Automationtrigger{}
	case "fortios:sys/autoscript:Autoscript":
		r = &Autoscript{}
	case "fortios:sys/centralmanagement:Centralmanagement":
		r = &Centralmanagement{}
	case "fortios:sys/clustersync:Clustersync":
		r = &Clustersync{}
	case "fortios:sys/console:Console":
		r = &Console{}
	case "fortios:sys/csf:Csf":
		r = &Csf{}
	case "fortios:sys/customlanguage:Customlanguage":
		r = &Customlanguage{}
	case "fortios:sys/ddns:Ddns":
		r = &Ddns{}
	case "fortios:sys/dedicatedmgmt:Dedicatedmgmt":
		r = &Dedicatedmgmt{}
	case "fortios:sys/dns64:Dns64":
		r = &Dns64{}
	case "fortios:sys/dns:Dns":
		r = &Dns{}
	case "fortios:sys/dnsdatabase:Dnsdatabase":
		r = &Dnsdatabase{}
	case "fortios:sys/dnsserver:Dnsserver":
		r = &Dnsserver{}
	case "fortios:sys/dscpbasedpriority:Dscpbasedpriority":
		r = &Dscpbasedpriority{}
	case "fortios:sys/emailserver:Emailserver":
		r = &Emailserver{}
	case "fortios:sys/externalresource:Externalresource":
		r = &Externalresource{}
	case "fortios:sys/federatedupgrade:Federatedupgrade":
		r = &Federatedupgrade{}
	case "fortios:sys/fipscc:Fipscc":
		r = &Fipscc{}
	case "fortios:sys/fm:Fm":
		r = &Fm{}
	case "fortios:sys/fortiai:Fortiai":
		r = &Fortiai{}
	case "fortios:sys/fortiguard:Fortiguard":
		r = &Fortiguard{}
	case "fortios:sys/fortimanager:Fortimanager":
		r = &Fortimanager{}
	case "fortios:sys/fortindr:Fortindr":
		r = &Fortindr{}
	case "fortios:sys/fortisandbox:Fortisandbox":
		r = &Fortisandbox{}
	case "fortios:sys/fssopolling:Fssopolling":
		r = &Fssopolling{}
	case "fortios:sys/ftmpush:Ftmpush":
		r = &Ftmpush{}
	case "fortios:sys/geneve:Geneve":
		r = &Geneve{}
	case "fortios:sys/geoipcountry:Geoipcountry":
		r = &Geoipcountry{}
	case "fortios:sys/geoipoverride:Geoipoverride":
		r = &Geoipoverride{}
	case "fortios:sys/global:Global":
		r = &Global{}
	case "fortios:sys/gretunnel:Gretunnel":
		r = &Gretunnel{}
	case "fortios:sys/ha:Ha":
		r = &Ha{}
	case "fortios:sys/hamonitor:Hamonitor":
		r = &Hamonitor{}
	case "fortios:sys/ike:Ike":
		r = &Ike{}
	case "fortios:sys/interface:Interface":
		r = &Interface{}
	case "fortios:sys/ipam:Ipam":
		r = &Ipam{}
	case "fortios:sys/ipiptunnel:Ipiptunnel":
		r = &Ipiptunnel{}
	case "fortios:sys/ips:Ips":
		r = &Ips{}
	case "fortios:sys/ipsecaggregate:Ipsecaggregate":
		r = &Ipsecaggregate{}
	case "fortios:sys/ipsurlfilterdns6:Ipsurlfilterdns6":
		r = &Ipsurlfilterdns6{}
	case "fortios:sys/ipsurlfilterdns:Ipsurlfilterdns":
		r = &Ipsurlfilterdns{}
	case "fortios:sys/ipv6neighborcache:Ipv6neighborcache":
		r = &Ipv6neighborcache{}
	case "fortios:sys/ipv6tunnel:Ipv6tunnel":
		r = &Ipv6tunnel{}
	case "fortios:sys/licenseForticare:LicenseForticare":
		r = &LicenseForticare{}
	case "fortios:sys/licenseVdom:LicenseVdom":
		r = &LicenseVdom{}
	case "fortios:sys/licenseVm:LicenseVm":
		r = &LicenseVm{}
	case "fortios:sys/linkmonitor:Linkmonitor":
		r = &Linkmonitor{}
	case "fortios:sys/ltemodem:Ltemodem":
		r = &Ltemodem{}
	case "fortios:sys/macaddresstable:Macaddresstable":
		r = &Macaddresstable{}
	case "fortios:sys/managementtunnel:Managementtunnel":
		r = &Managementtunnel{}
	case "fortios:sys/mobiletunnel:Mobiletunnel":
		r = &Mobiletunnel{}
	case "fortios:sys/modem:Modem":
		r = &Modem{}
	case "fortios:sys/nat64:Nat64":
		r = &Nat64{}
	case "fortios:sys/ndproxy:Ndproxy":
		r = &Ndproxy{}
	case "fortios:sys/netflow:Netflow":
		r = &Netflow{}
	case "fortios:sys/networkvisibility:Networkvisibility":
		r = &Networkvisibility{}
	case "fortios:sys/npu:Npu":
		r = &Npu{}
	case "fortios:sys/ntp:Ntp":
		r = &Ntp{}
	case "fortios:sys/objecttagging:Objecttagging":
		r = &Objecttagging{}
	case "fortios:sys/passwordpolicy:Passwordpolicy":
		r = &Passwordpolicy{}
	case "fortios:sys/passwordpolicyguestadmin:Passwordpolicyguestadmin":
		r = &Passwordpolicyguestadmin{}
	case "fortios:sys/physicalswitch:Physicalswitch":
		r = &Physicalswitch{}
	case "fortios:sys/pppoeinterface:Pppoeinterface":
		r = &Pppoeinterface{}
	case "fortios:sys/proberesponse:Proberesponse":
		r = &Proberesponse{}
	case "fortios:sys/proxyarp:Proxyarp":
		r = &Proxyarp{}
	case "fortios:sys/ptp:Ptp":
		r = &Ptp{}
	case "fortios:sys/replacemsggroup:Replacemsggroup":
		r = &Replacemsggroup{}
	case "fortios:sys/replacemsgimage:Replacemsgimage":
		r = &Replacemsgimage{}
	case "fortios:sys/resourcelimits:Resourcelimits":
		r = &Resourcelimits{}
	case "fortios:sys/saml:Saml":
		r = &Saml{}
	case "fortios:sys/sdnconnector:Sdnconnector":
		r = &Sdnconnector{}
	case "fortios:sys/sdwan:Sdwan":
		r = &Sdwan{}
	case "fortios:sys/sessionhelper:Sessionhelper":
		r = &Sessionhelper{}
	case "fortios:sys/sessionttl:Sessionttl":
		r = &Sessionttl{}
	case "fortios:sys/settingDns:SettingDns":
		r = &SettingDns{}
	case "fortios:sys/settingGlobal:SettingGlobal":
		r = &SettingGlobal{}
	case "fortios:sys/settingNtp:SettingNtp":
		r = &SettingNtp{}
	case "fortios:sys/settings:Settings":
		r = &Settings{}
	case "fortios:sys/sflow:Sflow":
		r = &Sflow{}
	case "fortios:sys/sittunnel:Sittunnel":
		r = &Sittunnel{}
	case "fortios:sys/smsserver:Smsserver":
		r = &Smsserver{}
	case "fortios:sys/speedtestschedule:Speedtestschedule":
		r = &Speedtestschedule{}
	case "fortios:sys/speedtestserver:Speedtestserver":
		r = &Speedtestserver{}
	case "fortios:sys/ssoadmin:Ssoadmin":
		r = &Ssoadmin{}
	case "fortios:sys/ssoforticloudadmin:Ssoforticloudadmin":
		r = &Ssoforticloudadmin{}
	case "fortios:sys/standalonecluster:Standalonecluster":
		r = &Standalonecluster{}
	case "fortios:sys/storage:Storage":
		r = &Storage{}
	case "fortios:sys/stp:Stp":
		r = &Stp{}
	case "fortios:sys/switchinterface:Switchinterface":
		r = &Switchinterface{}
	case "fortios:sys/tosbasedpriority:Tosbasedpriority":
		r = &Tosbasedpriority{}
	case "fortios:sys/vdom:Vdom":
		r = &Vdom{}
	case "fortios:sys/vdomSetting:VdomSetting":
		r = &VdomSetting{}
	case "fortios:sys/vdomdns:Vdomdns":
		r = &Vdomdns{}
	case "fortios:sys/vdomexception:Vdomexception":
		r = &Vdomexception{}
	case "fortios:sys/vdomlink:Vdomlink":
		r = &Vdomlink{}
	case "fortios:sys/vdomnetflow:Vdomnetflow":
		r = &Vdomnetflow{}
	case "fortios:sys/vdomproperty:Vdomproperty":
		r = &Vdomproperty{}
	case "fortios:sys/vdomradiusserver:Vdomradiusserver":
		r = &Vdomradiusserver{}
	case "fortios:sys/vdomsflow:Vdomsflow":
		r = &Vdomsflow{}
	case "fortios:sys/virtualswitch:Virtualswitch":
		r = &Virtualswitch{}
	case "fortios:sys/virtualwanlink:Virtualwanlink":
		r = &Virtualwanlink{}
	case "fortios:sys/virtualwirepair:Virtualwirepair":
		r = &Virtualwirepair{}
	case "fortios:sys/vnetunnel:Vnetunnel":
		r = &Vnetunnel{}
	case "fortios:sys/vxlan:Vxlan":
		r = &Vxlan{}
	case "fortios:sys/wccp:Wccp":
		r = &Wccp{}
	case "fortios:sys/zone:Zone":
		r = &Zone{}
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
		"sys/accprofile",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/acme",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/admin",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/adminAdministrator",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/adminProfiles",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/affinityinterrupt",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/affinitypacketredistribution",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/alarm",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/alias",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/apiuser",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/apiuserSetting",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/arptable",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/autoinstall",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/automationaction",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/automationdestination",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/automationstitch",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/automationtrigger",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/autoscript",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/centralmanagement",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/clustersync",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/console",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/csf",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/customlanguage",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/ddns",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/dedicatedmgmt",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/dns",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/dns64",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/dnsdatabase",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/dnsserver",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/dscpbasedpriority",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/emailserver",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/externalresource",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/federatedupgrade",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/fipscc",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/fm",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/fortiai",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/fortiguard",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/fortimanager",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/fortindr",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/fortisandbox",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/fssopolling",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/ftmpush",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/geneve",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/geoipcountry",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/geoipoverride",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/global",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/gretunnel",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/ha",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/hamonitor",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/ike",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/interface",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/ipam",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/ipiptunnel",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/ips",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/ipsecaggregate",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/ipsurlfilterdns",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/ipsurlfilterdns6",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/ipv6neighborcache",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/ipv6tunnel",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/licenseForticare",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/licenseVdom",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/licenseVm",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/linkmonitor",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/ltemodem",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/macaddresstable",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/managementtunnel",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/mobiletunnel",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/modem",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/nat64",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/ndproxy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/netflow",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/networkvisibility",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/npu",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/ntp",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/objecttagging",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/passwordpolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/passwordpolicyguestadmin",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/physicalswitch",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/pppoeinterface",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/proberesponse",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/proxyarp",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/ptp",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/replacemsggroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/replacemsgimage",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/resourcelimits",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/saml",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/sdnconnector",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/sdwan",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/sessionhelper",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/sessionttl",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/settingDns",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/settingGlobal",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/settingNtp",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/settings",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/sflow",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/sittunnel",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/smsserver",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/speedtestschedule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/speedtestserver",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/ssoadmin",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/ssoforticloudadmin",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/standalonecluster",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/storage",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/stp",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/switchinterface",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/tosbasedpriority",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/vdom",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/vdomSetting",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/vdomdns",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/vdomexception",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/vdomlink",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/vdomnetflow",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/vdomproperty",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/vdomradiusserver",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/vdomsflow",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/virtualswitch",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/virtualwanlink",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/virtualwirepair",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/vnetunnel",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/vxlan",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/wccp",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"fortios",
		"sys/zone",
		&module{version},
	)
}
