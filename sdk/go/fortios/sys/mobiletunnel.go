// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sys

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure Mobile tunnels, an implementation of Network Mobility (NEMO) extensions for Mobile IPv4 RFC5177.
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
//			_, err := sys.NewMobiletunnel(ctx, "trname", &sys.MobiletunnelArgs{
//				HashAlgorithm:    pulumi.String("hmac-md5"),
//				HomeAddress:      pulumi.String("0.0.0.0"),
//				HomeAgent:        pulumi.String("1.1.1.1"),
//				Lifetime:         pulumi.Int(65535),
//				NMhaeKey:         pulumi.String("'ENC M2wyM3DcnUhqgich7vsLk5oVuPAI9LTkcFNt0c3jI1ujC6w1XBot7gsRAf2S8X5dagfUnJGhZ5LrQxw21e4y8oXuCOLp8MmaRZbCkxYCAl1wm/wVY3aNzVk2+jE='"),
//				NMhaeKeyType:     pulumi.String("ascii"),
//				NMhaeSpi:         pulumi.Int(256),
//				RegInterval:      pulumi.Int(5),
//				RegRetry:         pulumi.Int(3),
//				RenewInterval:    pulumi.Int(60),
//				RoamingInterface: pulumi.String("port3"),
//				Status:           pulumi.String("disable"),
//				TunnelMode:       pulumi.String("gre"),
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
// # System MobileTunnel can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:sys/mobiletunnel:Mobiletunnel labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:sys/mobiletunnel:Mobiletunnel labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type Mobiletunnel struct {
	pulumi.CustomResourceState

	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Hash Algorithm (Keyed MD5). Valid values: `hmac-md5`.
	HashAlgorithm pulumi.StringOutput `pulumi:"hashAlgorithm"`
	// Home IP address (Format: xxx.xxx.xxx.xxx).
	HomeAddress pulumi.StringOutput `pulumi:"homeAddress"`
	// IPv4 address of the NEMO HA (Format: xxx.xxx.xxx.xxx).
	HomeAgent pulumi.StringOutput `pulumi:"homeAgent"`
	// NMMO HA registration request lifetime (180 - 65535 sec, default = 65535).
	Lifetime pulumi.IntOutput `pulumi:"lifetime"`
	// NEMO authentication key.
	NMhaeKey pulumi.StringOutput `pulumi:"nMhaeKey"`
	// NEMO authentication key type (ascii or base64). Valid values: `ascii`, `base64`.
	NMhaeKeyType pulumi.StringOutput `pulumi:"nMhaeKeyType"`
	// NEMO authentication SPI (default: 256).
	NMhaeSpi pulumi.IntOutput `pulumi:"nMhaeSpi"`
	// Tunnel name.
	Name pulumi.StringOutput `pulumi:"name"`
	// NEMO network configuration. The structure of `network` block is documented below.
	Networks MobiletunnelNetworkArrayOutput `pulumi:"networks"`
	// NMMO HA registration interval (5 - 300, default = 5).
	RegInterval pulumi.IntOutput `pulumi:"regInterval"`
	// Maximum number of NMMO HA registration retries (1 to 30, default = 3).
	RegRetry pulumi.IntOutput `pulumi:"regRetry"`
	// Time before lifetime expiraton to send NMMO HA re-registration (5 - 60, default = 60).
	RenewInterval pulumi.IntOutput `pulumi:"renewInterval"`
	// Select the associated interface name from available options.
	RoamingInterface pulumi.StringOutput `pulumi:"roamingInterface"`
	// Enable/disable this mobile tunnel. Valid values: `disable`, `enable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// NEMO tunnnel mode (GRE tunnel). Valid values: `gre`.
	TunnelMode pulumi.StringOutput `pulumi:"tunnelMode"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewMobiletunnel registers a new resource with the given unique name, arguments, and options.
func NewMobiletunnel(ctx *pulumi.Context,
	name string, args *MobiletunnelArgs, opts ...pulumi.ResourceOption) (*Mobiletunnel, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.HashAlgorithm == nil {
		return nil, errors.New("invalid value for required argument 'HashAlgorithm'")
	}
	if args.HomeAgent == nil {
		return nil, errors.New("invalid value for required argument 'HomeAgent'")
	}
	if args.Lifetime == nil {
		return nil, errors.New("invalid value for required argument 'Lifetime'")
	}
	if args.NMhaeKeyType == nil {
		return nil, errors.New("invalid value for required argument 'NMhaeKeyType'")
	}
	if args.NMhaeSpi == nil {
		return nil, errors.New("invalid value for required argument 'NMhaeSpi'")
	}
	if args.RegInterval == nil {
		return nil, errors.New("invalid value for required argument 'RegInterval'")
	}
	if args.RegRetry == nil {
		return nil, errors.New("invalid value for required argument 'RegRetry'")
	}
	if args.RenewInterval == nil {
		return nil, errors.New("invalid value for required argument 'RenewInterval'")
	}
	if args.RoamingInterface == nil {
		return nil, errors.New("invalid value for required argument 'RoamingInterface'")
	}
	if args.TunnelMode == nil {
		return nil, errors.New("invalid value for required argument 'TunnelMode'")
	}
	if args.NMhaeKey != nil {
		args.NMhaeKey = pulumi.ToSecret(args.NMhaeKey).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"nMhaeKey",
	})
	opts = append(opts, secrets)
	opts = pkgResourceDefaultOpts(opts)
	var resource Mobiletunnel
	err := ctx.RegisterResource("fortios:sys/mobiletunnel:Mobiletunnel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMobiletunnel gets an existing Mobiletunnel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMobiletunnel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MobiletunnelState, opts ...pulumi.ResourceOption) (*Mobiletunnel, error) {
	var resource Mobiletunnel
	err := ctx.ReadResource("fortios:sys/mobiletunnel:Mobiletunnel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Mobiletunnel resources.
type mobiletunnelState struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Hash Algorithm (Keyed MD5). Valid values: `hmac-md5`.
	HashAlgorithm *string `pulumi:"hashAlgorithm"`
	// Home IP address (Format: xxx.xxx.xxx.xxx).
	HomeAddress *string `pulumi:"homeAddress"`
	// IPv4 address of the NEMO HA (Format: xxx.xxx.xxx.xxx).
	HomeAgent *string `pulumi:"homeAgent"`
	// NMMO HA registration request lifetime (180 - 65535 sec, default = 65535).
	Lifetime *int `pulumi:"lifetime"`
	// NEMO authentication key.
	NMhaeKey *string `pulumi:"nMhaeKey"`
	// NEMO authentication key type (ascii or base64). Valid values: `ascii`, `base64`.
	NMhaeKeyType *string `pulumi:"nMhaeKeyType"`
	// NEMO authentication SPI (default: 256).
	NMhaeSpi *int `pulumi:"nMhaeSpi"`
	// Tunnel name.
	Name *string `pulumi:"name"`
	// NEMO network configuration. The structure of `network` block is documented below.
	Networks []MobiletunnelNetwork `pulumi:"networks"`
	// NMMO HA registration interval (5 - 300, default = 5).
	RegInterval *int `pulumi:"regInterval"`
	// Maximum number of NMMO HA registration retries (1 to 30, default = 3).
	RegRetry *int `pulumi:"regRetry"`
	// Time before lifetime expiraton to send NMMO HA re-registration (5 - 60, default = 60).
	RenewInterval *int `pulumi:"renewInterval"`
	// Select the associated interface name from available options.
	RoamingInterface *string `pulumi:"roamingInterface"`
	// Enable/disable this mobile tunnel. Valid values: `disable`, `enable`.
	Status *string `pulumi:"status"`
	// NEMO tunnnel mode (GRE tunnel). Valid values: `gre`.
	TunnelMode *string `pulumi:"tunnelMode"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type MobiletunnelState struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Hash Algorithm (Keyed MD5). Valid values: `hmac-md5`.
	HashAlgorithm pulumi.StringPtrInput
	// Home IP address (Format: xxx.xxx.xxx.xxx).
	HomeAddress pulumi.StringPtrInput
	// IPv4 address of the NEMO HA (Format: xxx.xxx.xxx.xxx).
	HomeAgent pulumi.StringPtrInput
	// NMMO HA registration request lifetime (180 - 65535 sec, default = 65535).
	Lifetime pulumi.IntPtrInput
	// NEMO authentication key.
	NMhaeKey pulumi.StringPtrInput
	// NEMO authentication key type (ascii or base64). Valid values: `ascii`, `base64`.
	NMhaeKeyType pulumi.StringPtrInput
	// NEMO authentication SPI (default: 256).
	NMhaeSpi pulumi.IntPtrInput
	// Tunnel name.
	Name pulumi.StringPtrInput
	// NEMO network configuration. The structure of `network` block is documented below.
	Networks MobiletunnelNetworkArrayInput
	// NMMO HA registration interval (5 - 300, default = 5).
	RegInterval pulumi.IntPtrInput
	// Maximum number of NMMO HA registration retries (1 to 30, default = 3).
	RegRetry pulumi.IntPtrInput
	// Time before lifetime expiraton to send NMMO HA re-registration (5 - 60, default = 60).
	RenewInterval pulumi.IntPtrInput
	// Select the associated interface name from available options.
	RoamingInterface pulumi.StringPtrInput
	// Enable/disable this mobile tunnel. Valid values: `disable`, `enable`.
	Status pulumi.StringPtrInput
	// NEMO tunnnel mode (GRE tunnel). Valid values: `gre`.
	TunnelMode pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (MobiletunnelState) ElementType() reflect.Type {
	return reflect.TypeOf((*mobiletunnelState)(nil)).Elem()
}

type mobiletunnelArgs struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Hash Algorithm (Keyed MD5). Valid values: `hmac-md5`.
	HashAlgorithm string `pulumi:"hashAlgorithm"`
	// Home IP address (Format: xxx.xxx.xxx.xxx).
	HomeAddress *string `pulumi:"homeAddress"`
	// IPv4 address of the NEMO HA (Format: xxx.xxx.xxx.xxx).
	HomeAgent string `pulumi:"homeAgent"`
	// NMMO HA registration request lifetime (180 - 65535 sec, default = 65535).
	Lifetime int `pulumi:"lifetime"`
	// NEMO authentication key.
	NMhaeKey *string `pulumi:"nMhaeKey"`
	// NEMO authentication key type (ascii or base64). Valid values: `ascii`, `base64`.
	NMhaeKeyType string `pulumi:"nMhaeKeyType"`
	// NEMO authentication SPI (default: 256).
	NMhaeSpi int `pulumi:"nMhaeSpi"`
	// Tunnel name.
	Name *string `pulumi:"name"`
	// NEMO network configuration. The structure of `network` block is documented below.
	Networks []MobiletunnelNetwork `pulumi:"networks"`
	// NMMO HA registration interval (5 - 300, default = 5).
	RegInterval int `pulumi:"regInterval"`
	// Maximum number of NMMO HA registration retries (1 to 30, default = 3).
	RegRetry int `pulumi:"regRetry"`
	// Time before lifetime expiraton to send NMMO HA re-registration (5 - 60, default = 60).
	RenewInterval int `pulumi:"renewInterval"`
	// Select the associated interface name from available options.
	RoamingInterface string `pulumi:"roamingInterface"`
	// Enable/disable this mobile tunnel. Valid values: `disable`, `enable`.
	Status *string `pulumi:"status"`
	// NEMO tunnnel mode (GRE tunnel). Valid values: `gre`.
	TunnelMode string `pulumi:"tunnelMode"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Mobiletunnel resource.
type MobiletunnelArgs struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Hash Algorithm (Keyed MD5). Valid values: `hmac-md5`.
	HashAlgorithm pulumi.StringInput
	// Home IP address (Format: xxx.xxx.xxx.xxx).
	HomeAddress pulumi.StringPtrInput
	// IPv4 address of the NEMO HA (Format: xxx.xxx.xxx.xxx).
	HomeAgent pulumi.StringInput
	// NMMO HA registration request lifetime (180 - 65535 sec, default = 65535).
	Lifetime pulumi.IntInput
	// NEMO authentication key.
	NMhaeKey pulumi.StringPtrInput
	// NEMO authentication key type (ascii or base64). Valid values: `ascii`, `base64`.
	NMhaeKeyType pulumi.StringInput
	// NEMO authentication SPI (default: 256).
	NMhaeSpi pulumi.IntInput
	// Tunnel name.
	Name pulumi.StringPtrInput
	// NEMO network configuration. The structure of `network` block is documented below.
	Networks MobiletunnelNetworkArrayInput
	// NMMO HA registration interval (5 - 300, default = 5).
	RegInterval pulumi.IntInput
	// Maximum number of NMMO HA registration retries (1 to 30, default = 3).
	RegRetry pulumi.IntInput
	// Time before lifetime expiraton to send NMMO HA re-registration (5 - 60, default = 60).
	RenewInterval pulumi.IntInput
	// Select the associated interface name from available options.
	RoamingInterface pulumi.StringInput
	// Enable/disable this mobile tunnel. Valid values: `disable`, `enable`.
	Status pulumi.StringPtrInput
	// NEMO tunnnel mode (GRE tunnel). Valid values: `gre`.
	TunnelMode pulumi.StringInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (MobiletunnelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mobiletunnelArgs)(nil)).Elem()
}

type MobiletunnelInput interface {
	pulumi.Input

	ToMobiletunnelOutput() MobiletunnelOutput
	ToMobiletunnelOutputWithContext(ctx context.Context) MobiletunnelOutput
}

func (*Mobiletunnel) ElementType() reflect.Type {
	return reflect.TypeOf((**Mobiletunnel)(nil)).Elem()
}

func (i *Mobiletunnel) ToMobiletunnelOutput() MobiletunnelOutput {
	return i.ToMobiletunnelOutputWithContext(context.Background())
}

func (i *Mobiletunnel) ToMobiletunnelOutputWithContext(ctx context.Context) MobiletunnelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MobiletunnelOutput)
}

// MobiletunnelArrayInput is an input type that accepts MobiletunnelArray and MobiletunnelArrayOutput values.
// You can construct a concrete instance of `MobiletunnelArrayInput` via:
//
//	MobiletunnelArray{ MobiletunnelArgs{...} }
type MobiletunnelArrayInput interface {
	pulumi.Input

	ToMobiletunnelArrayOutput() MobiletunnelArrayOutput
	ToMobiletunnelArrayOutputWithContext(context.Context) MobiletunnelArrayOutput
}

type MobiletunnelArray []MobiletunnelInput

func (MobiletunnelArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Mobiletunnel)(nil)).Elem()
}

func (i MobiletunnelArray) ToMobiletunnelArrayOutput() MobiletunnelArrayOutput {
	return i.ToMobiletunnelArrayOutputWithContext(context.Background())
}

func (i MobiletunnelArray) ToMobiletunnelArrayOutputWithContext(ctx context.Context) MobiletunnelArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MobiletunnelArrayOutput)
}

// MobiletunnelMapInput is an input type that accepts MobiletunnelMap and MobiletunnelMapOutput values.
// You can construct a concrete instance of `MobiletunnelMapInput` via:
//
//	MobiletunnelMap{ "key": MobiletunnelArgs{...} }
type MobiletunnelMapInput interface {
	pulumi.Input

	ToMobiletunnelMapOutput() MobiletunnelMapOutput
	ToMobiletunnelMapOutputWithContext(context.Context) MobiletunnelMapOutput
}

type MobiletunnelMap map[string]MobiletunnelInput

func (MobiletunnelMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Mobiletunnel)(nil)).Elem()
}

func (i MobiletunnelMap) ToMobiletunnelMapOutput() MobiletunnelMapOutput {
	return i.ToMobiletunnelMapOutputWithContext(context.Background())
}

func (i MobiletunnelMap) ToMobiletunnelMapOutputWithContext(ctx context.Context) MobiletunnelMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MobiletunnelMapOutput)
}

type MobiletunnelOutput struct{ *pulumi.OutputState }

func (MobiletunnelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Mobiletunnel)(nil)).Elem()
}

func (o MobiletunnelOutput) ToMobiletunnelOutput() MobiletunnelOutput {
	return o
}

func (o MobiletunnelOutput) ToMobiletunnelOutputWithContext(ctx context.Context) MobiletunnelOutput {
	return o
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o MobiletunnelOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Mobiletunnel) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Hash Algorithm (Keyed MD5). Valid values: `hmac-md5`.
func (o MobiletunnelOutput) HashAlgorithm() pulumi.StringOutput {
	return o.ApplyT(func(v *Mobiletunnel) pulumi.StringOutput { return v.HashAlgorithm }).(pulumi.StringOutput)
}

// Home IP address (Format: xxx.xxx.xxx.xxx).
func (o MobiletunnelOutput) HomeAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *Mobiletunnel) pulumi.StringOutput { return v.HomeAddress }).(pulumi.StringOutput)
}

// IPv4 address of the NEMO HA (Format: xxx.xxx.xxx.xxx).
func (o MobiletunnelOutput) HomeAgent() pulumi.StringOutput {
	return o.ApplyT(func(v *Mobiletunnel) pulumi.StringOutput { return v.HomeAgent }).(pulumi.StringOutput)
}

// NMMO HA registration request lifetime (180 - 65535 sec, default = 65535).
func (o MobiletunnelOutput) Lifetime() pulumi.IntOutput {
	return o.ApplyT(func(v *Mobiletunnel) pulumi.IntOutput { return v.Lifetime }).(pulumi.IntOutput)
}

// NEMO authentication key.
func (o MobiletunnelOutput) NMhaeKey() pulumi.StringOutput {
	return o.ApplyT(func(v *Mobiletunnel) pulumi.StringOutput { return v.NMhaeKey }).(pulumi.StringOutput)
}

// NEMO authentication key type (ascii or base64). Valid values: `ascii`, `base64`.
func (o MobiletunnelOutput) NMhaeKeyType() pulumi.StringOutput {
	return o.ApplyT(func(v *Mobiletunnel) pulumi.StringOutput { return v.NMhaeKeyType }).(pulumi.StringOutput)
}

// NEMO authentication SPI (default: 256).
func (o MobiletunnelOutput) NMhaeSpi() pulumi.IntOutput {
	return o.ApplyT(func(v *Mobiletunnel) pulumi.IntOutput { return v.NMhaeSpi }).(pulumi.IntOutput)
}

// Tunnel name.
func (o MobiletunnelOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Mobiletunnel) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// NEMO network configuration. The structure of `network` block is documented below.
func (o MobiletunnelOutput) Networks() MobiletunnelNetworkArrayOutput {
	return o.ApplyT(func(v *Mobiletunnel) MobiletunnelNetworkArrayOutput { return v.Networks }).(MobiletunnelNetworkArrayOutput)
}

// NMMO HA registration interval (5 - 300, default = 5).
func (o MobiletunnelOutput) RegInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *Mobiletunnel) pulumi.IntOutput { return v.RegInterval }).(pulumi.IntOutput)
}

// Maximum number of NMMO HA registration retries (1 to 30, default = 3).
func (o MobiletunnelOutput) RegRetry() pulumi.IntOutput {
	return o.ApplyT(func(v *Mobiletunnel) pulumi.IntOutput { return v.RegRetry }).(pulumi.IntOutput)
}

// Time before lifetime expiraton to send NMMO HA re-registration (5 - 60, default = 60).
func (o MobiletunnelOutput) RenewInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *Mobiletunnel) pulumi.IntOutput { return v.RenewInterval }).(pulumi.IntOutput)
}

// Select the associated interface name from available options.
func (o MobiletunnelOutput) RoamingInterface() pulumi.StringOutput {
	return o.ApplyT(func(v *Mobiletunnel) pulumi.StringOutput { return v.RoamingInterface }).(pulumi.StringOutput)
}

// Enable/disable this mobile tunnel. Valid values: `disable`, `enable`.
func (o MobiletunnelOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Mobiletunnel) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// NEMO tunnnel mode (GRE tunnel). Valid values: `gre`.
func (o MobiletunnelOutput) TunnelMode() pulumi.StringOutput {
	return o.ApplyT(func(v *Mobiletunnel) pulumi.StringOutput { return v.TunnelMode }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o MobiletunnelOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Mobiletunnel) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type MobiletunnelArrayOutput struct{ *pulumi.OutputState }

func (MobiletunnelArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Mobiletunnel)(nil)).Elem()
}

func (o MobiletunnelArrayOutput) ToMobiletunnelArrayOutput() MobiletunnelArrayOutput {
	return o
}

func (o MobiletunnelArrayOutput) ToMobiletunnelArrayOutputWithContext(ctx context.Context) MobiletunnelArrayOutput {
	return o
}

func (o MobiletunnelArrayOutput) Index(i pulumi.IntInput) MobiletunnelOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Mobiletunnel {
		return vs[0].([]*Mobiletunnel)[vs[1].(int)]
	}).(MobiletunnelOutput)
}

type MobiletunnelMapOutput struct{ *pulumi.OutputState }

func (MobiletunnelMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Mobiletunnel)(nil)).Elem()
}

func (o MobiletunnelMapOutput) ToMobiletunnelMapOutput() MobiletunnelMapOutput {
	return o
}

func (o MobiletunnelMapOutput) ToMobiletunnelMapOutputWithContext(ctx context.Context) MobiletunnelMapOutput {
	return o
}

func (o MobiletunnelMapOutput) MapIndex(k pulumi.StringInput) MobiletunnelOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Mobiletunnel {
		return vs[0].(map[string]*Mobiletunnel)[vs[1].(string)]
	}).(MobiletunnelOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MobiletunnelInput)(nil)).Elem(), &Mobiletunnel{})
	pulumi.RegisterInputType(reflect.TypeOf((*MobiletunnelArrayInput)(nil)).Elem(), MobiletunnelArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MobiletunnelMapInput)(nil)).Elem(), MobiletunnelMap{})
	pulumi.RegisterOutputType(MobiletunnelOutput{})
	pulumi.RegisterOutputType(MobiletunnelArrayOutput{})
	pulumi.RegisterOutputType(MobiletunnelMapOutput{})
}