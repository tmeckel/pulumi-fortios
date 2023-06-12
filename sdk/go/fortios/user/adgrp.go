// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package user

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure FSSO groups.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/user"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			trname1, err := user.NewFsso(ctx, "trname1", &user.FssoArgs{
//				Port:      pulumi.Int(32381),
//				Port2:     pulumi.Int(8000),
//				Port3:     pulumi.Int(8000),
//				Port4:     pulumi.Int(8000),
//				Port5:     pulumi.Int(8000),
//				Server:    pulumi.String("1.1.1.1"),
//				SourceIp:  pulumi.String("0.0.0.0"),
//				SourceIp6: pulumi.String("::"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = user.NewAdgrp(ctx, "trname", &user.AdgrpArgs{
//				ServerName: trname1.Name,
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
// # User Adgrp can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:user/adgrp:Adgrp labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:user/adgrp:Adgrp labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type Adgrp struct {
	pulumi.CustomResourceState

	// FSSO connector source.
	ConnectorSource pulumi.StringOutput `pulumi:"connectorSource"`
	// Group ID.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// FSSO agent name.
	ServerName pulumi.StringOutput `pulumi:"serverName"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewAdgrp registers a new resource with the given unique name, arguments, and options.
func NewAdgrp(ctx *pulumi.Context,
	name string, args *AdgrpArgs, opts ...pulumi.ResourceOption) (*Adgrp, error) {
	if args == nil {
		args = &AdgrpArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource Adgrp
	err := ctx.RegisterResource("fortios:user/adgrp:Adgrp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAdgrp gets an existing Adgrp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAdgrp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AdgrpState, opts ...pulumi.ResourceOption) (*Adgrp, error) {
	var resource Adgrp
	err := ctx.ReadResource("fortios:user/adgrp:Adgrp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Adgrp resources.
type adgrpState struct {
	// FSSO connector source.
	ConnectorSource *string `pulumi:"connectorSource"`
	// Group ID.
	Fosid *int `pulumi:"fosid"`
	// Name.
	Name *string `pulumi:"name"`
	// FSSO agent name.
	ServerName *string `pulumi:"serverName"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type AdgrpState struct {
	// FSSO connector source.
	ConnectorSource pulumi.StringPtrInput
	// Group ID.
	Fosid pulumi.IntPtrInput
	// Name.
	Name pulumi.StringPtrInput
	// FSSO agent name.
	ServerName pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (AdgrpState) ElementType() reflect.Type {
	return reflect.TypeOf((*adgrpState)(nil)).Elem()
}

type adgrpArgs struct {
	// FSSO connector source.
	ConnectorSource *string `pulumi:"connectorSource"`
	// Group ID.
	Fosid *int `pulumi:"fosid"`
	// Name.
	Name *string `pulumi:"name"`
	// FSSO agent name.
	ServerName *string `pulumi:"serverName"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Adgrp resource.
type AdgrpArgs struct {
	// FSSO connector source.
	ConnectorSource pulumi.StringPtrInput
	// Group ID.
	Fosid pulumi.IntPtrInput
	// Name.
	Name pulumi.StringPtrInput
	// FSSO agent name.
	ServerName pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (AdgrpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*adgrpArgs)(nil)).Elem()
}

type AdgrpInput interface {
	pulumi.Input

	ToAdgrpOutput() AdgrpOutput
	ToAdgrpOutputWithContext(ctx context.Context) AdgrpOutput
}

func (*Adgrp) ElementType() reflect.Type {
	return reflect.TypeOf((**Adgrp)(nil)).Elem()
}

func (i *Adgrp) ToAdgrpOutput() AdgrpOutput {
	return i.ToAdgrpOutputWithContext(context.Background())
}

func (i *Adgrp) ToAdgrpOutputWithContext(ctx context.Context) AdgrpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdgrpOutput)
}

// AdgrpArrayInput is an input type that accepts AdgrpArray and AdgrpArrayOutput values.
// You can construct a concrete instance of `AdgrpArrayInput` via:
//
//	AdgrpArray{ AdgrpArgs{...} }
type AdgrpArrayInput interface {
	pulumi.Input

	ToAdgrpArrayOutput() AdgrpArrayOutput
	ToAdgrpArrayOutputWithContext(context.Context) AdgrpArrayOutput
}

type AdgrpArray []AdgrpInput

func (AdgrpArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Adgrp)(nil)).Elem()
}

func (i AdgrpArray) ToAdgrpArrayOutput() AdgrpArrayOutput {
	return i.ToAdgrpArrayOutputWithContext(context.Background())
}

func (i AdgrpArray) ToAdgrpArrayOutputWithContext(ctx context.Context) AdgrpArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdgrpArrayOutput)
}

// AdgrpMapInput is an input type that accepts AdgrpMap and AdgrpMapOutput values.
// You can construct a concrete instance of `AdgrpMapInput` via:
//
//	AdgrpMap{ "key": AdgrpArgs{...} }
type AdgrpMapInput interface {
	pulumi.Input

	ToAdgrpMapOutput() AdgrpMapOutput
	ToAdgrpMapOutputWithContext(context.Context) AdgrpMapOutput
}

type AdgrpMap map[string]AdgrpInput

func (AdgrpMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Adgrp)(nil)).Elem()
}

func (i AdgrpMap) ToAdgrpMapOutput() AdgrpMapOutput {
	return i.ToAdgrpMapOutputWithContext(context.Background())
}

func (i AdgrpMap) ToAdgrpMapOutputWithContext(ctx context.Context) AdgrpMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdgrpMapOutput)
}

type AdgrpOutput struct{ *pulumi.OutputState }

func (AdgrpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Adgrp)(nil)).Elem()
}

func (o AdgrpOutput) ToAdgrpOutput() AdgrpOutput {
	return o
}

func (o AdgrpOutput) ToAdgrpOutputWithContext(ctx context.Context) AdgrpOutput {
	return o
}

// FSSO connector source.
func (o AdgrpOutput) ConnectorSource() pulumi.StringOutput {
	return o.ApplyT(func(v *Adgrp) pulumi.StringOutput { return v.ConnectorSource }).(pulumi.StringOutput)
}

// Group ID.
func (o AdgrpOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *Adgrp) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

// Name.
func (o AdgrpOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Adgrp) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// FSSO agent name.
func (o AdgrpOutput) ServerName() pulumi.StringOutput {
	return o.ApplyT(func(v *Adgrp) pulumi.StringOutput { return v.ServerName }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o AdgrpOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Adgrp) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type AdgrpArrayOutput struct{ *pulumi.OutputState }

func (AdgrpArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Adgrp)(nil)).Elem()
}

func (o AdgrpArrayOutput) ToAdgrpArrayOutput() AdgrpArrayOutput {
	return o
}

func (o AdgrpArrayOutput) ToAdgrpArrayOutputWithContext(ctx context.Context) AdgrpArrayOutput {
	return o
}

func (o AdgrpArrayOutput) Index(i pulumi.IntInput) AdgrpOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Adgrp {
		return vs[0].([]*Adgrp)[vs[1].(int)]
	}).(AdgrpOutput)
}

type AdgrpMapOutput struct{ *pulumi.OutputState }

func (AdgrpMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Adgrp)(nil)).Elem()
}

func (o AdgrpMapOutput) ToAdgrpMapOutput() AdgrpMapOutput {
	return o
}

func (o AdgrpMapOutput) ToAdgrpMapOutputWithContext(ctx context.Context) AdgrpMapOutput {
	return o
}

func (o AdgrpMapOutput) MapIndex(k pulumi.StringInput) AdgrpOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Adgrp {
		return vs[0].(map[string]*Adgrp)[vs[1].(string)]
	}).(AdgrpOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AdgrpInput)(nil)).Elem(), &Adgrp{})
	pulumi.RegisterInputType(reflect.TypeOf((*AdgrpArrayInput)(nil)).Elem(), AdgrpArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AdgrpMapInput)(nil)).Elem(), AdgrpMap{})
	pulumi.RegisterOutputType(AdgrpOutput{})
	pulumi.RegisterOutputType(AdgrpArrayOutput{})
	pulumi.RegisterOutputType(AdgrpMapOutput{})
}