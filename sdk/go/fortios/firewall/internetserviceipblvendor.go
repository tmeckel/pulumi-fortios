// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firewall

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// IP blacklist vendor. Applies to FortiOS Version `>= 6.2.4`.
//
// ## Import
//
// # Firewall InternetServiceIpblVendor can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:firewall/internetserviceipblvendor:Internetserviceipblvendor labelname {{fosid}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:firewall/internetserviceipblvendor:Internetserviceipblvendor labelname {{fosid}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type Internetserviceipblvendor struct {
	pulumi.CustomResourceState

	// IP blacklist vendor ID.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// IP blacklist vendor name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewInternetserviceipblvendor registers a new resource with the given unique name, arguments, and options.
func NewInternetserviceipblvendor(ctx *pulumi.Context,
	name string, args *InternetserviceipblvendorArgs, opts ...pulumi.ResourceOption) (*Internetserviceipblvendor, error) {
	if args == nil {
		args = &InternetserviceipblvendorArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource Internetserviceipblvendor
	err := ctx.RegisterResource("fortios:firewall/internetserviceipblvendor:Internetserviceipblvendor", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInternetserviceipblvendor gets an existing Internetserviceipblvendor resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInternetserviceipblvendor(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InternetserviceipblvendorState, opts ...pulumi.ResourceOption) (*Internetserviceipblvendor, error) {
	var resource Internetserviceipblvendor
	err := ctx.ReadResource("fortios:firewall/internetserviceipblvendor:Internetserviceipblvendor", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Internetserviceipblvendor resources.
type internetserviceipblvendorState struct {
	// IP blacklist vendor ID.
	Fosid *int `pulumi:"fosid"`
	// IP blacklist vendor name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type InternetserviceipblvendorState struct {
	// IP blacklist vendor ID.
	Fosid pulumi.IntPtrInput
	// IP blacklist vendor name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (InternetserviceipblvendorState) ElementType() reflect.Type {
	return reflect.TypeOf((*internetserviceipblvendorState)(nil)).Elem()
}

type internetserviceipblvendorArgs struct {
	// IP blacklist vendor ID.
	Fosid *int `pulumi:"fosid"`
	// IP blacklist vendor name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Internetserviceipblvendor resource.
type InternetserviceipblvendorArgs struct {
	// IP blacklist vendor ID.
	Fosid pulumi.IntPtrInput
	// IP blacklist vendor name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (InternetserviceipblvendorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*internetserviceipblvendorArgs)(nil)).Elem()
}

type InternetserviceipblvendorInput interface {
	pulumi.Input

	ToInternetserviceipblvendorOutput() InternetserviceipblvendorOutput
	ToInternetserviceipblvendorOutputWithContext(ctx context.Context) InternetserviceipblvendorOutput
}

func (*Internetserviceipblvendor) ElementType() reflect.Type {
	return reflect.TypeOf((**Internetserviceipblvendor)(nil)).Elem()
}

func (i *Internetserviceipblvendor) ToInternetserviceipblvendorOutput() InternetserviceipblvendorOutput {
	return i.ToInternetserviceipblvendorOutputWithContext(context.Background())
}

func (i *Internetserviceipblvendor) ToInternetserviceipblvendorOutputWithContext(ctx context.Context) InternetserviceipblvendorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InternetserviceipblvendorOutput)
}

// InternetserviceipblvendorArrayInput is an input type that accepts InternetserviceipblvendorArray and InternetserviceipblvendorArrayOutput values.
// You can construct a concrete instance of `InternetserviceipblvendorArrayInput` via:
//
//	InternetserviceipblvendorArray{ InternetserviceipblvendorArgs{...} }
type InternetserviceipblvendorArrayInput interface {
	pulumi.Input

	ToInternetserviceipblvendorArrayOutput() InternetserviceipblvendorArrayOutput
	ToInternetserviceipblvendorArrayOutputWithContext(context.Context) InternetserviceipblvendorArrayOutput
}

type InternetserviceipblvendorArray []InternetserviceipblvendorInput

func (InternetserviceipblvendorArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Internetserviceipblvendor)(nil)).Elem()
}

func (i InternetserviceipblvendorArray) ToInternetserviceipblvendorArrayOutput() InternetserviceipblvendorArrayOutput {
	return i.ToInternetserviceipblvendorArrayOutputWithContext(context.Background())
}

func (i InternetserviceipblvendorArray) ToInternetserviceipblvendorArrayOutputWithContext(ctx context.Context) InternetserviceipblvendorArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InternetserviceipblvendorArrayOutput)
}

// InternetserviceipblvendorMapInput is an input type that accepts InternetserviceipblvendorMap and InternetserviceipblvendorMapOutput values.
// You can construct a concrete instance of `InternetserviceipblvendorMapInput` via:
//
//	InternetserviceipblvendorMap{ "key": InternetserviceipblvendorArgs{...} }
type InternetserviceipblvendorMapInput interface {
	pulumi.Input

	ToInternetserviceipblvendorMapOutput() InternetserviceipblvendorMapOutput
	ToInternetserviceipblvendorMapOutputWithContext(context.Context) InternetserviceipblvendorMapOutput
}

type InternetserviceipblvendorMap map[string]InternetserviceipblvendorInput

func (InternetserviceipblvendorMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Internetserviceipblvendor)(nil)).Elem()
}

func (i InternetserviceipblvendorMap) ToInternetserviceipblvendorMapOutput() InternetserviceipblvendorMapOutput {
	return i.ToInternetserviceipblvendorMapOutputWithContext(context.Background())
}

func (i InternetserviceipblvendorMap) ToInternetserviceipblvendorMapOutputWithContext(ctx context.Context) InternetserviceipblvendorMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InternetserviceipblvendorMapOutput)
}

type InternetserviceipblvendorOutput struct{ *pulumi.OutputState }

func (InternetserviceipblvendorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Internetserviceipblvendor)(nil)).Elem()
}

func (o InternetserviceipblvendorOutput) ToInternetserviceipblvendorOutput() InternetserviceipblvendorOutput {
	return o
}

func (o InternetserviceipblvendorOutput) ToInternetserviceipblvendorOutputWithContext(ctx context.Context) InternetserviceipblvendorOutput {
	return o
}

// IP blacklist vendor ID.
func (o InternetserviceipblvendorOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *Internetserviceipblvendor) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

// IP blacklist vendor name.
func (o InternetserviceipblvendorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Internetserviceipblvendor) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o InternetserviceipblvendorOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Internetserviceipblvendor) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type InternetserviceipblvendorArrayOutput struct{ *pulumi.OutputState }

func (InternetserviceipblvendorArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Internetserviceipblvendor)(nil)).Elem()
}

func (o InternetserviceipblvendorArrayOutput) ToInternetserviceipblvendorArrayOutput() InternetserviceipblvendorArrayOutput {
	return o
}

func (o InternetserviceipblvendorArrayOutput) ToInternetserviceipblvendorArrayOutputWithContext(ctx context.Context) InternetserviceipblvendorArrayOutput {
	return o
}

func (o InternetserviceipblvendorArrayOutput) Index(i pulumi.IntInput) InternetserviceipblvendorOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Internetserviceipblvendor {
		return vs[0].([]*Internetserviceipblvendor)[vs[1].(int)]
	}).(InternetserviceipblvendorOutput)
}

type InternetserviceipblvendorMapOutput struct{ *pulumi.OutputState }

func (InternetserviceipblvendorMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Internetserviceipblvendor)(nil)).Elem()
}

func (o InternetserviceipblvendorMapOutput) ToInternetserviceipblvendorMapOutput() InternetserviceipblvendorMapOutput {
	return o
}

func (o InternetserviceipblvendorMapOutput) ToInternetserviceipblvendorMapOutputWithContext(ctx context.Context) InternetserviceipblvendorMapOutput {
	return o
}

func (o InternetserviceipblvendorMapOutput) MapIndex(k pulumi.StringInput) InternetserviceipblvendorOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Internetserviceipblvendor {
		return vs[0].(map[string]*Internetserviceipblvendor)[vs[1].(string)]
	}).(InternetserviceipblvendorOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InternetserviceipblvendorInput)(nil)).Elem(), &Internetserviceipblvendor{})
	pulumi.RegisterInputType(reflect.TypeOf((*InternetserviceipblvendorArrayInput)(nil)).Elem(), InternetserviceipblvendorArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InternetserviceipblvendorMapInput)(nil)).Elem(), InternetserviceipblvendorMap{})
	pulumi.RegisterOutputType(InternetserviceipblvendorOutput{})
	pulumi.RegisterOutputType(InternetserviceipblvendorArrayOutput{})
	pulumi.RegisterOutputType(InternetserviceipblvendorMapOutput{})
}
