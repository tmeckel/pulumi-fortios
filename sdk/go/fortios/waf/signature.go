// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package waf

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Hidden table for datasource.
//
// ## Import
//
// # Waf Signature can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:waf/signature:Signature labelname {{fosid}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:waf/signature:Signature labelname {{fosid}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type Signature struct {
	pulumi.CustomResourceState

	// Signature description.
	Desc pulumi.StringOutput `pulumi:"desc"`
	// Signature ID.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSignature registers a new resource with the given unique name, arguments, and options.
func NewSignature(ctx *pulumi.Context,
	name string, args *SignatureArgs, opts ...pulumi.ResourceOption) (*Signature, error) {
	if args == nil {
		args = &SignatureArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource Signature
	err := ctx.RegisterResource("fortios:waf/signature:Signature", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSignature gets an existing Signature resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSignature(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SignatureState, opts ...pulumi.ResourceOption) (*Signature, error) {
	var resource Signature
	err := ctx.ReadResource("fortios:waf/signature:Signature", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Signature resources.
type signatureState struct {
	// Signature description.
	Desc *string `pulumi:"desc"`
	// Signature ID.
	Fosid *int `pulumi:"fosid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SignatureState struct {
	// Signature description.
	Desc pulumi.StringPtrInput
	// Signature ID.
	Fosid pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SignatureState) ElementType() reflect.Type {
	return reflect.TypeOf((*signatureState)(nil)).Elem()
}

type signatureArgs struct {
	// Signature description.
	Desc *string `pulumi:"desc"`
	// Signature ID.
	Fosid *int `pulumi:"fosid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Signature resource.
type SignatureArgs struct {
	// Signature description.
	Desc pulumi.StringPtrInput
	// Signature ID.
	Fosid pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SignatureArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*signatureArgs)(nil)).Elem()
}

type SignatureInput interface {
	pulumi.Input

	ToSignatureOutput() SignatureOutput
	ToSignatureOutputWithContext(ctx context.Context) SignatureOutput
}

func (*Signature) ElementType() reflect.Type {
	return reflect.TypeOf((**Signature)(nil)).Elem()
}

func (i *Signature) ToSignatureOutput() SignatureOutput {
	return i.ToSignatureOutputWithContext(context.Background())
}

func (i *Signature) ToSignatureOutputWithContext(ctx context.Context) SignatureOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SignatureOutput)
}

// SignatureArrayInput is an input type that accepts SignatureArray and SignatureArrayOutput values.
// You can construct a concrete instance of `SignatureArrayInput` via:
//
//	SignatureArray{ SignatureArgs{...} }
type SignatureArrayInput interface {
	pulumi.Input

	ToSignatureArrayOutput() SignatureArrayOutput
	ToSignatureArrayOutputWithContext(context.Context) SignatureArrayOutput
}

type SignatureArray []SignatureInput

func (SignatureArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Signature)(nil)).Elem()
}

func (i SignatureArray) ToSignatureArrayOutput() SignatureArrayOutput {
	return i.ToSignatureArrayOutputWithContext(context.Background())
}

func (i SignatureArray) ToSignatureArrayOutputWithContext(ctx context.Context) SignatureArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SignatureArrayOutput)
}

// SignatureMapInput is an input type that accepts SignatureMap and SignatureMapOutput values.
// You can construct a concrete instance of `SignatureMapInput` via:
//
//	SignatureMap{ "key": SignatureArgs{...} }
type SignatureMapInput interface {
	pulumi.Input

	ToSignatureMapOutput() SignatureMapOutput
	ToSignatureMapOutputWithContext(context.Context) SignatureMapOutput
}

type SignatureMap map[string]SignatureInput

func (SignatureMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Signature)(nil)).Elem()
}

func (i SignatureMap) ToSignatureMapOutput() SignatureMapOutput {
	return i.ToSignatureMapOutputWithContext(context.Background())
}

func (i SignatureMap) ToSignatureMapOutputWithContext(ctx context.Context) SignatureMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SignatureMapOutput)
}

type SignatureOutput struct{ *pulumi.OutputState }

func (SignatureOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Signature)(nil)).Elem()
}

func (o SignatureOutput) ToSignatureOutput() SignatureOutput {
	return o
}

func (o SignatureOutput) ToSignatureOutputWithContext(ctx context.Context) SignatureOutput {
	return o
}

// Signature description.
func (o SignatureOutput) Desc() pulumi.StringOutput {
	return o.ApplyT(func(v *Signature) pulumi.StringOutput { return v.Desc }).(pulumi.StringOutput)
}

// Signature ID.
func (o SignatureOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *Signature) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SignatureOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Signature) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SignatureArrayOutput struct{ *pulumi.OutputState }

func (SignatureArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Signature)(nil)).Elem()
}

func (o SignatureArrayOutput) ToSignatureArrayOutput() SignatureArrayOutput {
	return o
}

func (o SignatureArrayOutput) ToSignatureArrayOutputWithContext(ctx context.Context) SignatureArrayOutput {
	return o
}

func (o SignatureArrayOutput) Index(i pulumi.IntInput) SignatureOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Signature {
		return vs[0].([]*Signature)[vs[1].(int)]
	}).(SignatureOutput)
}

type SignatureMapOutput struct{ *pulumi.OutputState }

func (SignatureMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Signature)(nil)).Elem()
}

func (o SignatureMapOutput) ToSignatureMapOutput() SignatureMapOutput {
	return o
}

func (o SignatureMapOutput) ToSignatureMapOutputWithContext(ctx context.Context) SignatureMapOutput {
	return o
}

func (o SignatureMapOutput) MapIndex(k pulumi.StringInput) SignatureOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Signature {
		return vs[0].(map[string]*Signature)[vs[1].(string)]
	}).(SignatureOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SignatureInput)(nil)).Elem(), &Signature{})
	pulumi.RegisterInputType(reflect.TypeOf((*SignatureArrayInput)(nil)).Elem(), SignatureArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SignatureMapInput)(nil)).Elem(), SignatureMap{})
	pulumi.RegisterOutputType(SignatureOutput{})
	pulumi.RegisterOutputType(SignatureArrayOutput{})
	pulumi.RegisterOutputType(SignatureMapOutput{})
}
