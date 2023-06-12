// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sys

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure access profiles for system administrators.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/sys"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := sys.NewAccprofile(ctx, "test12", &sys.AccprofileArgs{
//				Admintimeout:         pulumi.Int(10),
//				AdmintimeoutOverride: pulumi.String("disable"),
//				Authgrp:              pulumi.String("read-write"),
//				Ftviewgrp:            pulumi.String("read-write"),
//				Fwgrp:                pulumi.String("custom"),
//				FwgrpPermission: &sys.AccprofileFwgrpPermissionArgs{
//					Address:  pulumi.String("read-write"),
//					Policy:   pulumi.String("read-write"),
//					Schedule: pulumi.String("none"),
//					Service:  pulumi.String("none"),
//				},
//				Loggrp: pulumi.String("read-write"),
//				LoggrpPermission: &sys.AccprofileLoggrpPermissionArgs{
//					Config:       pulumi.String("none"),
//					DataAccess:   pulumi.String("none"),
//					ReportAccess: pulumi.String("none"),
//					ThreatWeight: pulumi.String("none"),
//				},
//				Netgrp: pulumi.String("read-write"),
//				NetgrpPermission: &sys.AccprofileNetgrpPermissionArgs{
//					Cfg:           pulumi.String("none"),
//					PacketCapture: pulumi.String("none"),
//					RouteCfg:      pulumi.String("none"),
//				},
//				Scope:     pulumi.String("vdom"),
//				Secfabgrp: pulumi.String("read-write"),
//				Sysgrp:    pulumi.String("read-write"),
//				SysgrpPermission: &sys.AccprofileSysgrpPermissionArgs{
//					Admin: pulumi.String("none"),
//					Cfg:   pulumi.String("none"),
//					Mnt:   pulumi.String("none"),
//					Upd:   pulumi.String("none"),
//				},
//				Utmgrp: pulumi.String("custom"),
//				UtmgrpPermission: &sys.AccprofileUtmgrpPermissionArgs{
//					Antivirus:          pulumi.String("read-write"),
//					ApplicationControl: pulumi.String("none"),
//					DataLossPrevention: pulumi.String("none"),
//					Dnsfilter:          pulumi.String("none"),
//					EndpointControl:    pulumi.String("none"),
//					Icap:               pulumi.String("none"),
//					Ips:                pulumi.String("read-write"),
//					Voip:               pulumi.String("none"),
//					Waf:                pulumi.String("none"),
//					Webfilter:          pulumi.String("none"),
//				},
//				Vpngrp:    pulumi.String("read-write"),
//				Wanoptgrp: pulumi.String("read-write"),
//				Wifi:      pulumi.String("read-write"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// # System Accprofile can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:sys/accprofile:Accprofile labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:sys/accprofile:Accprofile labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type Accprofile struct {
	pulumi.CustomResourceState

	// Administrator timeout for this access profile (0 - 480 min, default = 10, 0 means never timeout).
	Admintimeout pulumi.IntOutput `pulumi:"admintimeout"`
	// Enable/disable overriding the global administrator idle timeout. Valid values: `enable`, `disable`.
	AdmintimeoutOverride pulumi.StringOutput `pulumi:"admintimeoutOverride"`
	// Administrator access to Users and Devices. Valid values: `none`, `read`, `read-write`.
	Authgrp pulumi.StringOutput `pulumi:"authgrp"`
	// Comment.
	Comments pulumi.StringPtrOutput `pulumi:"comments"`
	// FortiView. Valid values: `none`, `read`, `read-write`.
	Ftviewgrp pulumi.StringOutput `pulumi:"ftviewgrp"`
	// Administrator access to the Firewall configuration. Valid values: `none`, `read`, `read-write`, `custom`.
	Fwgrp pulumi.StringOutput `pulumi:"fwgrp"`
	// Custom firewall permission. The structure of `fwgrpPermission` block is documented below.
	FwgrpPermission AccprofileFwgrpPermissionOutput `pulumi:"fwgrpPermission"`
	// Administrator access to Logging and Reporting including viewing log messages. Valid values: `none`, `read`, `read-write`, `custom`.
	Loggrp pulumi.StringOutput `pulumi:"loggrp"`
	// Custom Log & Report permission. The structure of `loggrpPermission` block is documented below.
	LoggrpPermission AccprofileLoggrpPermissionOutput `pulumi:"loggrpPermission"`
	// Profile name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Network Configuration. Valid values: `none`, `read`, `read-write`, `custom`.
	Netgrp pulumi.StringOutput `pulumi:"netgrp"`
	// Custom network permission. The structure of `netgrpPermission` block is documented below.
	NetgrpPermission AccprofileNetgrpPermissionOutput `pulumi:"netgrpPermission"`
	// Scope of admin access: global or specific VDOM(s). Valid values: `vdom`, `global`.
	Scope pulumi.StringOutput `pulumi:"scope"`
	// Security Fabric. Valid values: `none`, `read`, `read-write`.
	Secfabgrp pulumi.StringOutput `pulumi:"secfabgrp"`
	// System Configuration. Valid values: `none`, `read`, `read-write`, `custom`.
	Sysgrp pulumi.StringOutput `pulumi:"sysgrp"`
	// Custom system permission. The structure of `sysgrpPermission` block is documented below.
	SysgrpPermission AccprofileSysgrpPermissionOutput `pulumi:"sysgrpPermission"`
	// Enable/disable permission to run system diagnostic commands. Valid values: `enable`, `disable`.
	SystemDiagnostics pulumi.StringOutput `pulumi:"systemDiagnostics"`
	// Enable/disable permission to execute SSH commands. Valid values: `enable`, `disable`.
	SystemExecuteSsh pulumi.StringOutput `pulumi:"systemExecuteSsh"`
	// Enable/disable permission to execute TELNET commands. Valid values: `enable`, `disable`.
	SystemExecuteTelnet pulumi.StringOutput `pulumi:"systemExecuteTelnet"`
	// Administrator access to Security Profiles. Valid values: `none`, `read`, `read-write`, `custom`.
	Utmgrp pulumi.StringOutput `pulumi:"utmgrp"`
	// Custom Security Profile permissions. The structure of `utmgrpPermission` block is documented below.
	UtmgrpPermission AccprofileUtmgrpPermissionOutput `pulumi:"utmgrpPermission"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// Administrator access to IPsec, SSL, PPTP, and L2TP VPN. Valid values: `none`, `read`, `read-write`.
	Vpngrp pulumi.StringOutput `pulumi:"vpngrp"`
	// Administrator access to WAN Opt & Cache. Valid values: `none`, `read`, `read-write`.
	Wanoptgrp pulumi.StringOutput `pulumi:"wanoptgrp"`
	// Administrator access to the WiFi controller and Switch controller. Valid values: `none`, `read`, `read-write`.
	Wifi pulumi.StringOutput `pulumi:"wifi"`
}

// NewAccprofile registers a new resource with the given unique name, arguments, and options.
func NewAccprofile(ctx *pulumi.Context,
	name string, args *AccprofileArgs, opts ...pulumi.ResourceOption) (*Accprofile, error) {
	if args == nil {
		args = &AccprofileArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource Accprofile
	err := ctx.RegisterResource("fortios:sys/accprofile:Accprofile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccprofile gets an existing Accprofile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccprofile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccprofileState, opts ...pulumi.ResourceOption) (*Accprofile, error) {
	var resource Accprofile
	err := ctx.ReadResource("fortios:sys/accprofile:Accprofile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Accprofile resources.
type accprofileState struct {
	// Administrator timeout for this access profile (0 - 480 min, default = 10, 0 means never timeout).
	Admintimeout *int `pulumi:"admintimeout"`
	// Enable/disable overriding the global administrator idle timeout. Valid values: `enable`, `disable`.
	AdmintimeoutOverride *string `pulumi:"admintimeoutOverride"`
	// Administrator access to Users and Devices. Valid values: `none`, `read`, `read-write`.
	Authgrp *string `pulumi:"authgrp"`
	// Comment.
	Comments *string `pulumi:"comments"`
	// FortiView. Valid values: `none`, `read`, `read-write`.
	Ftviewgrp *string `pulumi:"ftviewgrp"`
	// Administrator access to the Firewall configuration. Valid values: `none`, `read`, `read-write`, `custom`.
	Fwgrp *string `pulumi:"fwgrp"`
	// Custom firewall permission. The structure of `fwgrpPermission` block is documented below.
	FwgrpPermission *AccprofileFwgrpPermission `pulumi:"fwgrpPermission"`
	// Administrator access to Logging and Reporting including viewing log messages. Valid values: `none`, `read`, `read-write`, `custom`.
	Loggrp *string `pulumi:"loggrp"`
	// Custom Log & Report permission. The structure of `loggrpPermission` block is documented below.
	LoggrpPermission *AccprofileLoggrpPermission `pulumi:"loggrpPermission"`
	// Profile name.
	Name *string `pulumi:"name"`
	// Network Configuration. Valid values: `none`, `read`, `read-write`, `custom`.
	Netgrp *string `pulumi:"netgrp"`
	// Custom network permission. The structure of `netgrpPermission` block is documented below.
	NetgrpPermission *AccprofileNetgrpPermission `pulumi:"netgrpPermission"`
	// Scope of admin access: global or specific VDOM(s). Valid values: `vdom`, `global`.
	Scope *string `pulumi:"scope"`
	// Security Fabric. Valid values: `none`, `read`, `read-write`.
	Secfabgrp *string `pulumi:"secfabgrp"`
	// System Configuration. Valid values: `none`, `read`, `read-write`, `custom`.
	Sysgrp *string `pulumi:"sysgrp"`
	// Custom system permission. The structure of `sysgrpPermission` block is documented below.
	SysgrpPermission *AccprofileSysgrpPermission `pulumi:"sysgrpPermission"`
	// Enable/disable permission to run system diagnostic commands. Valid values: `enable`, `disable`.
	SystemDiagnostics *string `pulumi:"systemDiagnostics"`
	// Enable/disable permission to execute SSH commands. Valid values: `enable`, `disable`.
	SystemExecuteSsh *string `pulumi:"systemExecuteSsh"`
	// Enable/disable permission to execute TELNET commands. Valid values: `enable`, `disable`.
	SystemExecuteTelnet *string `pulumi:"systemExecuteTelnet"`
	// Administrator access to Security Profiles. Valid values: `none`, `read`, `read-write`, `custom`.
	Utmgrp *string `pulumi:"utmgrp"`
	// Custom Security Profile permissions. The structure of `utmgrpPermission` block is documented below.
	UtmgrpPermission *AccprofileUtmgrpPermission `pulumi:"utmgrpPermission"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Administrator access to IPsec, SSL, PPTP, and L2TP VPN. Valid values: `none`, `read`, `read-write`.
	Vpngrp *string `pulumi:"vpngrp"`
	// Administrator access to WAN Opt & Cache. Valid values: `none`, `read`, `read-write`.
	Wanoptgrp *string `pulumi:"wanoptgrp"`
	// Administrator access to the WiFi controller and Switch controller. Valid values: `none`, `read`, `read-write`.
	Wifi *string `pulumi:"wifi"`
}

type AccprofileState struct {
	// Administrator timeout for this access profile (0 - 480 min, default = 10, 0 means never timeout).
	Admintimeout pulumi.IntPtrInput
	// Enable/disable overriding the global administrator idle timeout. Valid values: `enable`, `disable`.
	AdmintimeoutOverride pulumi.StringPtrInput
	// Administrator access to Users and Devices. Valid values: `none`, `read`, `read-write`.
	Authgrp pulumi.StringPtrInput
	// Comment.
	Comments pulumi.StringPtrInput
	// FortiView. Valid values: `none`, `read`, `read-write`.
	Ftviewgrp pulumi.StringPtrInput
	// Administrator access to the Firewall configuration. Valid values: `none`, `read`, `read-write`, `custom`.
	Fwgrp pulumi.StringPtrInput
	// Custom firewall permission. The structure of `fwgrpPermission` block is documented below.
	FwgrpPermission AccprofileFwgrpPermissionPtrInput
	// Administrator access to Logging and Reporting including viewing log messages. Valid values: `none`, `read`, `read-write`, `custom`.
	Loggrp pulumi.StringPtrInput
	// Custom Log & Report permission. The structure of `loggrpPermission` block is documented below.
	LoggrpPermission AccprofileLoggrpPermissionPtrInput
	// Profile name.
	Name pulumi.StringPtrInput
	// Network Configuration. Valid values: `none`, `read`, `read-write`, `custom`.
	Netgrp pulumi.StringPtrInput
	// Custom network permission. The structure of `netgrpPermission` block is documented below.
	NetgrpPermission AccprofileNetgrpPermissionPtrInput
	// Scope of admin access: global or specific VDOM(s). Valid values: `vdom`, `global`.
	Scope pulumi.StringPtrInput
	// Security Fabric. Valid values: `none`, `read`, `read-write`.
	Secfabgrp pulumi.StringPtrInput
	// System Configuration. Valid values: `none`, `read`, `read-write`, `custom`.
	Sysgrp pulumi.StringPtrInput
	// Custom system permission. The structure of `sysgrpPermission` block is documented below.
	SysgrpPermission AccprofileSysgrpPermissionPtrInput
	// Enable/disable permission to run system diagnostic commands. Valid values: `enable`, `disable`.
	SystemDiagnostics pulumi.StringPtrInput
	// Enable/disable permission to execute SSH commands. Valid values: `enable`, `disable`.
	SystemExecuteSsh pulumi.StringPtrInput
	// Enable/disable permission to execute TELNET commands. Valid values: `enable`, `disable`.
	SystemExecuteTelnet pulumi.StringPtrInput
	// Administrator access to Security Profiles. Valid values: `none`, `read`, `read-write`, `custom`.
	Utmgrp pulumi.StringPtrInput
	// Custom Security Profile permissions. The structure of `utmgrpPermission` block is documented below.
	UtmgrpPermission AccprofileUtmgrpPermissionPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Administrator access to IPsec, SSL, PPTP, and L2TP VPN. Valid values: `none`, `read`, `read-write`.
	Vpngrp pulumi.StringPtrInput
	// Administrator access to WAN Opt & Cache. Valid values: `none`, `read`, `read-write`.
	Wanoptgrp pulumi.StringPtrInput
	// Administrator access to the WiFi controller and Switch controller. Valid values: `none`, `read`, `read-write`.
	Wifi pulumi.StringPtrInput
}

func (AccprofileState) ElementType() reflect.Type {
	return reflect.TypeOf((*accprofileState)(nil)).Elem()
}

type accprofileArgs struct {
	// Administrator timeout for this access profile (0 - 480 min, default = 10, 0 means never timeout).
	Admintimeout *int `pulumi:"admintimeout"`
	// Enable/disable overriding the global administrator idle timeout. Valid values: `enable`, `disable`.
	AdmintimeoutOverride *string `pulumi:"admintimeoutOverride"`
	// Administrator access to Users and Devices. Valid values: `none`, `read`, `read-write`.
	Authgrp *string `pulumi:"authgrp"`
	// Comment.
	Comments *string `pulumi:"comments"`
	// FortiView. Valid values: `none`, `read`, `read-write`.
	Ftviewgrp *string `pulumi:"ftviewgrp"`
	// Administrator access to the Firewall configuration. Valid values: `none`, `read`, `read-write`, `custom`.
	Fwgrp *string `pulumi:"fwgrp"`
	// Custom firewall permission. The structure of `fwgrpPermission` block is documented below.
	FwgrpPermission *AccprofileFwgrpPermission `pulumi:"fwgrpPermission"`
	// Administrator access to Logging and Reporting including viewing log messages. Valid values: `none`, `read`, `read-write`, `custom`.
	Loggrp *string `pulumi:"loggrp"`
	// Custom Log & Report permission. The structure of `loggrpPermission` block is documented below.
	LoggrpPermission *AccprofileLoggrpPermission `pulumi:"loggrpPermission"`
	// Profile name.
	Name *string `pulumi:"name"`
	// Network Configuration. Valid values: `none`, `read`, `read-write`, `custom`.
	Netgrp *string `pulumi:"netgrp"`
	// Custom network permission. The structure of `netgrpPermission` block is documented below.
	NetgrpPermission *AccprofileNetgrpPermission `pulumi:"netgrpPermission"`
	// Scope of admin access: global or specific VDOM(s). Valid values: `vdom`, `global`.
	Scope *string `pulumi:"scope"`
	// Security Fabric. Valid values: `none`, `read`, `read-write`.
	Secfabgrp *string `pulumi:"secfabgrp"`
	// System Configuration. Valid values: `none`, `read`, `read-write`, `custom`.
	Sysgrp *string `pulumi:"sysgrp"`
	// Custom system permission. The structure of `sysgrpPermission` block is documented below.
	SysgrpPermission *AccprofileSysgrpPermission `pulumi:"sysgrpPermission"`
	// Enable/disable permission to run system diagnostic commands. Valid values: `enable`, `disable`.
	SystemDiagnostics *string `pulumi:"systemDiagnostics"`
	// Enable/disable permission to execute SSH commands. Valid values: `enable`, `disable`.
	SystemExecuteSsh *string `pulumi:"systemExecuteSsh"`
	// Enable/disable permission to execute TELNET commands. Valid values: `enable`, `disable`.
	SystemExecuteTelnet *string `pulumi:"systemExecuteTelnet"`
	// Administrator access to Security Profiles. Valid values: `none`, `read`, `read-write`, `custom`.
	Utmgrp *string `pulumi:"utmgrp"`
	// Custom Security Profile permissions. The structure of `utmgrpPermission` block is documented below.
	UtmgrpPermission *AccprofileUtmgrpPermission `pulumi:"utmgrpPermission"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Administrator access to IPsec, SSL, PPTP, and L2TP VPN. Valid values: `none`, `read`, `read-write`.
	Vpngrp *string `pulumi:"vpngrp"`
	// Administrator access to WAN Opt & Cache. Valid values: `none`, `read`, `read-write`.
	Wanoptgrp *string `pulumi:"wanoptgrp"`
	// Administrator access to the WiFi controller and Switch controller. Valid values: `none`, `read`, `read-write`.
	Wifi *string `pulumi:"wifi"`
}

// The set of arguments for constructing a Accprofile resource.
type AccprofileArgs struct {
	// Administrator timeout for this access profile (0 - 480 min, default = 10, 0 means never timeout).
	Admintimeout pulumi.IntPtrInput
	// Enable/disable overriding the global administrator idle timeout. Valid values: `enable`, `disable`.
	AdmintimeoutOverride pulumi.StringPtrInput
	// Administrator access to Users and Devices. Valid values: `none`, `read`, `read-write`.
	Authgrp pulumi.StringPtrInput
	// Comment.
	Comments pulumi.StringPtrInput
	// FortiView. Valid values: `none`, `read`, `read-write`.
	Ftviewgrp pulumi.StringPtrInput
	// Administrator access to the Firewall configuration. Valid values: `none`, `read`, `read-write`, `custom`.
	Fwgrp pulumi.StringPtrInput
	// Custom firewall permission. The structure of `fwgrpPermission` block is documented below.
	FwgrpPermission AccprofileFwgrpPermissionPtrInput
	// Administrator access to Logging and Reporting including viewing log messages. Valid values: `none`, `read`, `read-write`, `custom`.
	Loggrp pulumi.StringPtrInput
	// Custom Log & Report permission. The structure of `loggrpPermission` block is documented below.
	LoggrpPermission AccprofileLoggrpPermissionPtrInput
	// Profile name.
	Name pulumi.StringPtrInput
	// Network Configuration. Valid values: `none`, `read`, `read-write`, `custom`.
	Netgrp pulumi.StringPtrInput
	// Custom network permission. The structure of `netgrpPermission` block is documented below.
	NetgrpPermission AccprofileNetgrpPermissionPtrInput
	// Scope of admin access: global or specific VDOM(s). Valid values: `vdom`, `global`.
	Scope pulumi.StringPtrInput
	// Security Fabric. Valid values: `none`, `read`, `read-write`.
	Secfabgrp pulumi.StringPtrInput
	// System Configuration. Valid values: `none`, `read`, `read-write`, `custom`.
	Sysgrp pulumi.StringPtrInput
	// Custom system permission. The structure of `sysgrpPermission` block is documented below.
	SysgrpPermission AccprofileSysgrpPermissionPtrInput
	// Enable/disable permission to run system diagnostic commands. Valid values: `enable`, `disable`.
	SystemDiagnostics pulumi.StringPtrInput
	// Enable/disable permission to execute SSH commands. Valid values: `enable`, `disable`.
	SystemExecuteSsh pulumi.StringPtrInput
	// Enable/disable permission to execute TELNET commands. Valid values: `enable`, `disable`.
	SystemExecuteTelnet pulumi.StringPtrInput
	// Administrator access to Security Profiles. Valid values: `none`, `read`, `read-write`, `custom`.
	Utmgrp pulumi.StringPtrInput
	// Custom Security Profile permissions. The structure of `utmgrpPermission` block is documented below.
	UtmgrpPermission AccprofileUtmgrpPermissionPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Administrator access to IPsec, SSL, PPTP, and L2TP VPN. Valid values: `none`, `read`, `read-write`.
	Vpngrp pulumi.StringPtrInput
	// Administrator access to WAN Opt & Cache. Valid values: `none`, `read`, `read-write`.
	Wanoptgrp pulumi.StringPtrInput
	// Administrator access to the WiFi controller and Switch controller. Valid values: `none`, `read`, `read-write`.
	Wifi pulumi.StringPtrInput
}

func (AccprofileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accprofileArgs)(nil)).Elem()
}

type AccprofileInput interface {
	pulumi.Input

	ToAccprofileOutput() AccprofileOutput
	ToAccprofileOutputWithContext(ctx context.Context) AccprofileOutput
}

func (*Accprofile) ElementType() reflect.Type {
	return reflect.TypeOf((**Accprofile)(nil)).Elem()
}

func (i *Accprofile) ToAccprofileOutput() AccprofileOutput {
	return i.ToAccprofileOutputWithContext(context.Background())
}

func (i *Accprofile) ToAccprofileOutputWithContext(ctx context.Context) AccprofileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccprofileOutput)
}

// AccprofileArrayInput is an input type that accepts AccprofileArray and AccprofileArrayOutput values.
// You can construct a concrete instance of `AccprofileArrayInput` via:
//
//	AccprofileArray{ AccprofileArgs{...} }
type AccprofileArrayInput interface {
	pulumi.Input

	ToAccprofileArrayOutput() AccprofileArrayOutput
	ToAccprofileArrayOutputWithContext(context.Context) AccprofileArrayOutput
}

type AccprofileArray []AccprofileInput

func (AccprofileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Accprofile)(nil)).Elem()
}

func (i AccprofileArray) ToAccprofileArrayOutput() AccprofileArrayOutput {
	return i.ToAccprofileArrayOutputWithContext(context.Background())
}

func (i AccprofileArray) ToAccprofileArrayOutputWithContext(ctx context.Context) AccprofileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccprofileArrayOutput)
}

// AccprofileMapInput is an input type that accepts AccprofileMap and AccprofileMapOutput values.
// You can construct a concrete instance of `AccprofileMapInput` via:
//
//	AccprofileMap{ "key": AccprofileArgs{...} }
type AccprofileMapInput interface {
	pulumi.Input

	ToAccprofileMapOutput() AccprofileMapOutput
	ToAccprofileMapOutputWithContext(context.Context) AccprofileMapOutput
}

type AccprofileMap map[string]AccprofileInput

func (AccprofileMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Accprofile)(nil)).Elem()
}

func (i AccprofileMap) ToAccprofileMapOutput() AccprofileMapOutput {
	return i.ToAccprofileMapOutputWithContext(context.Background())
}

func (i AccprofileMap) ToAccprofileMapOutputWithContext(ctx context.Context) AccprofileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccprofileMapOutput)
}

type AccprofileOutput struct{ *pulumi.OutputState }

func (AccprofileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Accprofile)(nil)).Elem()
}

func (o AccprofileOutput) ToAccprofileOutput() AccprofileOutput {
	return o
}

func (o AccprofileOutput) ToAccprofileOutputWithContext(ctx context.Context) AccprofileOutput {
	return o
}

// Administrator timeout for this access profile (0 - 480 min, default = 10, 0 means never timeout).
func (o AccprofileOutput) Admintimeout() pulumi.IntOutput {
	return o.ApplyT(func(v *Accprofile) pulumi.IntOutput { return v.Admintimeout }).(pulumi.IntOutput)
}

// Enable/disable overriding the global administrator idle timeout. Valid values: `enable`, `disable`.
func (o AccprofileOutput) AdmintimeoutOverride() pulumi.StringOutput {
	return o.ApplyT(func(v *Accprofile) pulumi.StringOutput { return v.AdmintimeoutOverride }).(pulumi.StringOutput)
}

// Administrator access to Users and Devices. Valid values: `none`, `read`, `read-write`.
func (o AccprofileOutput) Authgrp() pulumi.StringOutput {
	return o.ApplyT(func(v *Accprofile) pulumi.StringOutput { return v.Authgrp }).(pulumi.StringOutput)
}

// Comment.
func (o AccprofileOutput) Comments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Accprofile) pulumi.StringPtrOutput { return v.Comments }).(pulumi.StringPtrOutput)
}

// FortiView. Valid values: `none`, `read`, `read-write`.
func (o AccprofileOutput) Ftviewgrp() pulumi.StringOutput {
	return o.ApplyT(func(v *Accprofile) pulumi.StringOutput { return v.Ftviewgrp }).(pulumi.StringOutput)
}

// Administrator access to the Firewall configuration. Valid values: `none`, `read`, `read-write`, `custom`.
func (o AccprofileOutput) Fwgrp() pulumi.StringOutput {
	return o.ApplyT(func(v *Accprofile) pulumi.StringOutput { return v.Fwgrp }).(pulumi.StringOutput)
}

// Custom firewall permission. The structure of `fwgrpPermission` block is documented below.
func (o AccprofileOutput) FwgrpPermission() AccprofileFwgrpPermissionOutput {
	return o.ApplyT(func(v *Accprofile) AccprofileFwgrpPermissionOutput { return v.FwgrpPermission }).(AccprofileFwgrpPermissionOutput)
}

// Administrator access to Logging and Reporting including viewing log messages. Valid values: `none`, `read`, `read-write`, `custom`.
func (o AccprofileOutput) Loggrp() pulumi.StringOutput {
	return o.ApplyT(func(v *Accprofile) pulumi.StringOutput { return v.Loggrp }).(pulumi.StringOutput)
}

// Custom Log & Report permission. The structure of `loggrpPermission` block is documented below.
func (o AccprofileOutput) LoggrpPermission() AccprofileLoggrpPermissionOutput {
	return o.ApplyT(func(v *Accprofile) AccprofileLoggrpPermissionOutput { return v.LoggrpPermission }).(AccprofileLoggrpPermissionOutput)
}

// Profile name.
func (o AccprofileOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Accprofile) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Network Configuration. Valid values: `none`, `read`, `read-write`, `custom`.
func (o AccprofileOutput) Netgrp() pulumi.StringOutput {
	return o.ApplyT(func(v *Accprofile) pulumi.StringOutput { return v.Netgrp }).(pulumi.StringOutput)
}

// Custom network permission. The structure of `netgrpPermission` block is documented below.
func (o AccprofileOutput) NetgrpPermission() AccprofileNetgrpPermissionOutput {
	return o.ApplyT(func(v *Accprofile) AccprofileNetgrpPermissionOutput { return v.NetgrpPermission }).(AccprofileNetgrpPermissionOutput)
}

// Scope of admin access: global or specific VDOM(s). Valid values: `vdom`, `global`.
func (o AccprofileOutput) Scope() pulumi.StringOutput {
	return o.ApplyT(func(v *Accprofile) pulumi.StringOutput { return v.Scope }).(pulumi.StringOutput)
}

// Security Fabric. Valid values: `none`, `read`, `read-write`.
func (o AccprofileOutput) Secfabgrp() pulumi.StringOutput {
	return o.ApplyT(func(v *Accprofile) pulumi.StringOutput { return v.Secfabgrp }).(pulumi.StringOutput)
}

// System Configuration. Valid values: `none`, `read`, `read-write`, `custom`.
func (o AccprofileOutput) Sysgrp() pulumi.StringOutput {
	return o.ApplyT(func(v *Accprofile) pulumi.StringOutput { return v.Sysgrp }).(pulumi.StringOutput)
}

// Custom system permission. The structure of `sysgrpPermission` block is documented below.
func (o AccprofileOutput) SysgrpPermission() AccprofileSysgrpPermissionOutput {
	return o.ApplyT(func(v *Accprofile) AccprofileSysgrpPermissionOutput { return v.SysgrpPermission }).(AccprofileSysgrpPermissionOutput)
}

// Enable/disable permission to run system diagnostic commands. Valid values: `enable`, `disable`.
func (o AccprofileOutput) SystemDiagnostics() pulumi.StringOutput {
	return o.ApplyT(func(v *Accprofile) pulumi.StringOutput { return v.SystemDiagnostics }).(pulumi.StringOutput)
}

// Enable/disable permission to execute SSH commands. Valid values: `enable`, `disable`.
func (o AccprofileOutput) SystemExecuteSsh() pulumi.StringOutput {
	return o.ApplyT(func(v *Accprofile) pulumi.StringOutput { return v.SystemExecuteSsh }).(pulumi.StringOutput)
}

// Enable/disable permission to execute TELNET commands. Valid values: `enable`, `disable`.
func (o AccprofileOutput) SystemExecuteTelnet() pulumi.StringOutput {
	return o.ApplyT(func(v *Accprofile) pulumi.StringOutput { return v.SystemExecuteTelnet }).(pulumi.StringOutput)
}

// Administrator access to Security Profiles. Valid values: `none`, `read`, `read-write`, `custom`.
func (o AccprofileOutput) Utmgrp() pulumi.StringOutput {
	return o.ApplyT(func(v *Accprofile) pulumi.StringOutput { return v.Utmgrp }).(pulumi.StringOutput)
}

// Custom Security Profile permissions. The structure of `utmgrpPermission` block is documented below.
func (o AccprofileOutput) UtmgrpPermission() AccprofileUtmgrpPermissionOutput {
	return o.ApplyT(func(v *Accprofile) AccprofileUtmgrpPermissionOutput { return v.UtmgrpPermission }).(AccprofileUtmgrpPermissionOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o AccprofileOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Accprofile) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// Administrator access to IPsec, SSL, PPTP, and L2TP VPN. Valid values: `none`, `read`, `read-write`.
func (o AccprofileOutput) Vpngrp() pulumi.StringOutput {
	return o.ApplyT(func(v *Accprofile) pulumi.StringOutput { return v.Vpngrp }).(pulumi.StringOutput)
}

// Administrator access to WAN Opt & Cache. Valid values: `none`, `read`, `read-write`.
func (o AccprofileOutput) Wanoptgrp() pulumi.StringOutput {
	return o.ApplyT(func(v *Accprofile) pulumi.StringOutput { return v.Wanoptgrp }).(pulumi.StringOutput)
}

// Administrator access to the WiFi controller and Switch controller. Valid values: `none`, `read`, `read-write`.
func (o AccprofileOutput) Wifi() pulumi.StringOutput {
	return o.ApplyT(func(v *Accprofile) pulumi.StringOutput { return v.Wifi }).(pulumi.StringOutput)
}

type AccprofileArrayOutput struct{ *pulumi.OutputState }

func (AccprofileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Accprofile)(nil)).Elem()
}

func (o AccprofileArrayOutput) ToAccprofileArrayOutput() AccprofileArrayOutput {
	return o
}

func (o AccprofileArrayOutput) ToAccprofileArrayOutputWithContext(ctx context.Context) AccprofileArrayOutput {
	return o
}

func (o AccprofileArrayOutput) Index(i pulumi.IntInput) AccprofileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Accprofile {
		return vs[0].([]*Accprofile)[vs[1].(int)]
	}).(AccprofileOutput)
}

type AccprofileMapOutput struct{ *pulumi.OutputState }

func (AccprofileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Accprofile)(nil)).Elem()
}

func (o AccprofileMapOutput) ToAccprofileMapOutput() AccprofileMapOutput {
	return o
}

func (o AccprofileMapOutput) ToAccprofileMapOutputWithContext(ctx context.Context) AccprofileMapOutput {
	return o
}

func (o AccprofileMapOutput) MapIndex(k pulumi.StringInput) AccprofileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Accprofile {
		return vs[0].(map[string]*Accprofile)[vs[1].(string)]
	}).(AccprofileOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AccprofileInput)(nil)).Elem(), &Accprofile{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccprofileArrayInput)(nil)).Elem(), AccprofileArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccprofileMapInput)(nil)).Elem(), AccprofileMap{})
	pulumi.RegisterOutputType(AccprofileOutput{})
	pulumi.RegisterOutputType(AccprofileArrayOutput{})
	pulumi.RegisterOutputType(AccprofileMapOutput{})
}