// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpn

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure L2TP.
//
// ## Import
//
// # Vpn L2Tp can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:vpn/l2tp:L2tp labelname VpnL2Tp
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:vpn/l2tp:L2tp labelname VpnL2Tp
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type L2tp struct {
	pulumi.CustomResourceState

	// Enable/disable data compression. Valid values: `enable`, `disable`.
	Compress pulumi.StringOutput `pulumi:"compress"`
	// End IP.
	Eip pulumi.StringOutput `pulumi:"eip"`
	// Enable/disable IPsec enforcement. Valid values: `enable`, `disable`.
	EnforceIpsec pulumi.StringOutput `pulumi:"enforceIpsec"`
	// L2TP hello message interval in seconds (0 - 3600 sec, default = 60).
	HelloInterval pulumi.IntOutput `pulumi:"helloInterval"`
	// Time in seconds between PPPoE Link Control Protocol (LCP) echo requests.
	LcpEchoInterval pulumi.IntOutput `pulumi:"lcpEchoInterval"`
	// Maximum number of missed LCP echo messages before disconnect.
	LcpMaxEchoFails pulumi.IntOutput `pulumi:"lcpMaxEchoFails"`
	// Start IP.
	Sip pulumi.StringOutput `pulumi:"sip"`
	// Enable/disable FortiGate as a L2TP gateway. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// User group.
	Usrgrp pulumi.StringOutput `pulumi:"usrgrp"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewL2tp registers a new resource with the given unique name, arguments, and options.
func NewL2tp(ctx *pulumi.Context,
	name string, args *L2tpArgs, opts ...pulumi.ResourceOption) (*L2tp, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Status == nil {
		return nil, errors.New("invalid value for required argument 'Status'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource L2tp
	err := ctx.RegisterResource("fortios:vpn/l2tp:L2tp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetL2tp gets an existing L2tp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetL2tp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *L2tpState, opts ...pulumi.ResourceOption) (*L2tp, error) {
	var resource L2tp
	err := ctx.ReadResource("fortios:vpn/l2tp:L2tp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering L2tp resources.
type l2tpState struct {
	// Enable/disable data compression. Valid values: `enable`, `disable`.
	Compress *string `pulumi:"compress"`
	// End IP.
	Eip *string `pulumi:"eip"`
	// Enable/disable IPsec enforcement. Valid values: `enable`, `disable`.
	EnforceIpsec *string `pulumi:"enforceIpsec"`
	// L2TP hello message interval in seconds (0 - 3600 sec, default = 60).
	HelloInterval *int `pulumi:"helloInterval"`
	// Time in seconds between PPPoE Link Control Protocol (LCP) echo requests.
	LcpEchoInterval *int `pulumi:"lcpEchoInterval"`
	// Maximum number of missed LCP echo messages before disconnect.
	LcpMaxEchoFails *int `pulumi:"lcpMaxEchoFails"`
	// Start IP.
	Sip *string `pulumi:"sip"`
	// Enable/disable FortiGate as a L2TP gateway. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// User group.
	Usrgrp *string `pulumi:"usrgrp"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type L2tpState struct {
	// Enable/disable data compression. Valid values: `enable`, `disable`.
	Compress pulumi.StringPtrInput
	// End IP.
	Eip pulumi.StringPtrInput
	// Enable/disable IPsec enforcement. Valid values: `enable`, `disable`.
	EnforceIpsec pulumi.StringPtrInput
	// L2TP hello message interval in seconds (0 - 3600 sec, default = 60).
	HelloInterval pulumi.IntPtrInput
	// Time in seconds between PPPoE Link Control Protocol (LCP) echo requests.
	LcpEchoInterval pulumi.IntPtrInput
	// Maximum number of missed LCP echo messages before disconnect.
	LcpMaxEchoFails pulumi.IntPtrInput
	// Start IP.
	Sip pulumi.StringPtrInput
	// Enable/disable FortiGate as a L2TP gateway. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// User group.
	Usrgrp pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (L2tpState) ElementType() reflect.Type {
	return reflect.TypeOf((*l2tpState)(nil)).Elem()
}

type l2tpArgs struct {
	// Enable/disable data compression. Valid values: `enable`, `disable`.
	Compress *string `pulumi:"compress"`
	// End IP.
	Eip *string `pulumi:"eip"`
	// Enable/disable IPsec enforcement. Valid values: `enable`, `disable`.
	EnforceIpsec *string `pulumi:"enforceIpsec"`
	// L2TP hello message interval in seconds (0 - 3600 sec, default = 60).
	HelloInterval *int `pulumi:"helloInterval"`
	// Time in seconds between PPPoE Link Control Protocol (LCP) echo requests.
	LcpEchoInterval *int `pulumi:"lcpEchoInterval"`
	// Maximum number of missed LCP echo messages before disconnect.
	LcpMaxEchoFails *int `pulumi:"lcpMaxEchoFails"`
	// Start IP.
	Sip *string `pulumi:"sip"`
	// Enable/disable FortiGate as a L2TP gateway. Valid values: `enable`, `disable`.
	Status string `pulumi:"status"`
	// User group.
	Usrgrp *string `pulumi:"usrgrp"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a L2tp resource.
type L2tpArgs struct {
	// Enable/disable data compression. Valid values: `enable`, `disable`.
	Compress pulumi.StringPtrInput
	// End IP.
	Eip pulumi.StringPtrInput
	// Enable/disable IPsec enforcement. Valid values: `enable`, `disable`.
	EnforceIpsec pulumi.StringPtrInput
	// L2TP hello message interval in seconds (0 - 3600 sec, default = 60).
	HelloInterval pulumi.IntPtrInput
	// Time in seconds between PPPoE Link Control Protocol (LCP) echo requests.
	LcpEchoInterval pulumi.IntPtrInput
	// Maximum number of missed LCP echo messages before disconnect.
	LcpMaxEchoFails pulumi.IntPtrInput
	// Start IP.
	Sip pulumi.StringPtrInput
	// Enable/disable FortiGate as a L2TP gateway. Valid values: `enable`, `disable`.
	Status pulumi.StringInput
	// User group.
	Usrgrp pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (L2tpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*l2tpArgs)(nil)).Elem()
}

type L2tpInput interface {
	pulumi.Input

	ToL2tpOutput() L2tpOutput
	ToL2tpOutputWithContext(ctx context.Context) L2tpOutput
}

func (*L2tp) ElementType() reflect.Type {
	return reflect.TypeOf((**L2tp)(nil)).Elem()
}

func (i *L2tp) ToL2tpOutput() L2tpOutput {
	return i.ToL2tpOutputWithContext(context.Background())
}

func (i *L2tp) ToL2tpOutputWithContext(ctx context.Context) L2tpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(L2tpOutput)
}

// L2tpArrayInput is an input type that accepts L2tpArray and L2tpArrayOutput values.
// You can construct a concrete instance of `L2tpArrayInput` via:
//
//	L2tpArray{ L2tpArgs{...} }
type L2tpArrayInput interface {
	pulumi.Input

	ToL2tpArrayOutput() L2tpArrayOutput
	ToL2tpArrayOutputWithContext(context.Context) L2tpArrayOutput
}

type L2tpArray []L2tpInput

func (L2tpArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*L2tp)(nil)).Elem()
}

func (i L2tpArray) ToL2tpArrayOutput() L2tpArrayOutput {
	return i.ToL2tpArrayOutputWithContext(context.Background())
}

func (i L2tpArray) ToL2tpArrayOutputWithContext(ctx context.Context) L2tpArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(L2tpArrayOutput)
}

// L2tpMapInput is an input type that accepts L2tpMap and L2tpMapOutput values.
// You can construct a concrete instance of `L2tpMapInput` via:
//
//	L2tpMap{ "key": L2tpArgs{...} }
type L2tpMapInput interface {
	pulumi.Input

	ToL2tpMapOutput() L2tpMapOutput
	ToL2tpMapOutputWithContext(context.Context) L2tpMapOutput
}

type L2tpMap map[string]L2tpInput

func (L2tpMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*L2tp)(nil)).Elem()
}

func (i L2tpMap) ToL2tpMapOutput() L2tpMapOutput {
	return i.ToL2tpMapOutputWithContext(context.Background())
}

func (i L2tpMap) ToL2tpMapOutputWithContext(ctx context.Context) L2tpMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(L2tpMapOutput)
}

type L2tpOutput struct{ *pulumi.OutputState }

func (L2tpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**L2tp)(nil)).Elem()
}

func (o L2tpOutput) ToL2tpOutput() L2tpOutput {
	return o
}

func (o L2tpOutput) ToL2tpOutputWithContext(ctx context.Context) L2tpOutput {
	return o
}

// Enable/disable data compression. Valid values: `enable`, `disable`.
func (o L2tpOutput) Compress() pulumi.StringOutput {
	return o.ApplyT(func(v *L2tp) pulumi.StringOutput { return v.Compress }).(pulumi.StringOutput)
}

// End IP.
func (o L2tpOutput) Eip() pulumi.StringOutput {
	return o.ApplyT(func(v *L2tp) pulumi.StringOutput { return v.Eip }).(pulumi.StringOutput)
}

// Enable/disable IPsec enforcement. Valid values: `enable`, `disable`.
func (o L2tpOutput) EnforceIpsec() pulumi.StringOutput {
	return o.ApplyT(func(v *L2tp) pulumi.StringOutput { return v.EnforceIpsec }).(pulumi.StringOutput)
}

// L2TP hello message interval in seconds (0 - 3600 sec, default = 60).
func (o L2tpOutput) HelloInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *L2tp) pulumi.IntOutput { return v.HelloInterval }).(pulumi.IntOutput)
}

// Time in seconds between PPPoE Link Control Protocol (LCP) echo requests.
func (o L2tpOutput) LcpEchoInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *L2tp) pulumi.IntOutput { return v.LcpEchoInterval }).(pulumi.IntOutput)
}

// Maximum number of missed LCP echo messages before disconnect.
func (o L2tpOutput) LcpMaxEchoFails() pulumi.IntOutput {
	return o.ApplyT(func(v *L2tp) pulumi.IntOutput { return v.LcpMaxEchoFails }).(pulumi.IntOutput)
}

// Start IP.
func (o L2tpOutput) Sip() pulumi.StringOutput {
	return o.ApplyT(func(v *L2tp) pulumi.StringOutput { return v.Sip }).(pulumi.StringOutput)
}

// Enable/disable FortiGate as a L2TP gateway. Valid values: `enable`, `disable`.
func (o L2tpOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *L2tp) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// User group.
func (o L2tpOutput) Usrgrp() pulumi.StringOutput {
	return o.ApplyT(func(v *L2tp) pulumi.StringOutput { return v.Usrgrp }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o L2tpOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *L2tp) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type L2tpArrayOutput struct{ *pulumi.OutputState }

func (L2tpArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*L2tp)(nil)).Elem()
}

func (o L2tpArrayOutput) ToL2tpArrayOutput() L2tpArrayOutput {
	return o
}

func (o L2tpArrayOutput) ToL2tpArrayOutputWithContext(ctx context.Context) L2tpArrayOutput {
	return o
}

func (o L2tpArrayOutput) Index(i pulumi.IntInput) L2tpOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *L2tp {
		return vs[0].([]*L2tp)[vs[1].(int)]
	}).(L2tpOutput)
}

type L2tpMapOutput struct{ *pulumi.OutputState }

func (L2tpMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*L2tp)(nil)).Elem()
}

func (o L2tpMapOutput) ToL2tpMapOutput() L2tpMapOutput {
	return o
}

func (o L2tpMapOutput) ToL2tpMapOutputWithContext(ctx context.Context) L2tpMapOutput {
	return o
}

func (o L2tpMapOutput) MapIndex(k pulumi.StringInput) L2tpOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *L2tp {
		return vs[0].(map[string]*L2tp)[vs[1].(string)]
	}).(L2tpOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*L2tpInput)(nil)).Elem(), &L2tp{})
	pulumi.RegisterInputType(reflect.TypeOf((*L2tpArrayInput)(nil)).Elem(), L2tpArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*L2tpMapInput)(nil)).Elem(), L2tpMap{})
	pulumi.RegisterOutputType(L2tpOutput{})
	pulumi.RegisterOutputType(L2tpArrayOutput{})
	pulumi.RegisterOutputType(L2tpMapOutput{})
}
