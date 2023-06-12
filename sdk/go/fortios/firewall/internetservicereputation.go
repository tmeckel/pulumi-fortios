// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firewall

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Show Internet Service reputation. Applies to FortiOS Version `>= 6.2.4`.
//
// ## Import
//
// # Firewall InternetServiceReputation can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:firewall/internetservicereputation:Internetservicereputation labelname {{fosid}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:firewall/internetservicereputation:Internetservicereputation labelname {{fosid}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type Internetservicereputation struct {
	pulumi.CustomResourceState

	// Description.
	Description pulumi.StringOutput `pulumi:"description"`
	// Internet Service Reputation ID.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewInternetservicereputation registers a new resource with the given unique name, arguments, and options.
func NewInternetservicereputation(ctx *pulumi.Context,
	name string, args *InternetservicereputationArgs, opts ...pulumi.ResourceOption) (*Internetservicereputation, error) {
	if args == nil {
		args = &InternetservicereputationArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource Internetservicereputation
	err := ctx.RegisterResource("fortios:firewall/internetservicereputation:Internetservicereputation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInternetservicereputation gets an existing Internetservicereputation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInternetservicereputation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InternetservicereputationState, opts ...pulumi.ResourceOption) (*Internetservicereputation, error) {
	var resource Internetservicereputation
	err := ctx.ReadResource("fortios:firewall/internetservicereputation:Internetservicereputation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Internetservicereputation resources.
type internetservicereputationState struct {
	// Description.
	Description *string `pulumi:"description"`
	// Internet Service Reputation ID.
	Fosid *int `pulumi:"fosid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type InternetservicereputationState struct {
	// Description.
	Description pulumi.StringPtrInput
	// Internet Service Reputation ID.
	Fosid pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (InternetservicereputationState) ElementType() reflect.Type {
	return reflect.TypeOf((*internetservicereputationState)(nil)).Elem()
}

type internetservicereputationArgs struct {
	// Description.
	Description *string `pulumi:"description"`
	// Internet Service Reputation ID.
	Fosid *int `pulumi:"fosid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Internetservicereputation resource.
type InternetservicereputationArgs struct {
	// Description.
	Description pulumi.StringPtrInput
	// Internet Service Reputation ID.
	Fosid pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (InternetservicereputationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*internetservicereputationArgs)(nil)).Elem()
}

type InternetservicereputationInput interface {
	pulumi.Input

	ToInternetservicereputationOutput() InternetservicereputationOutput
	ToInternetservicereputationOutputWithContext(ctx context.Context) InternetservicereputationOutput
}

func (*Internetservicereputation) ElementType() reflect.Type {
	return reflect.TypeOf((**Internetservicereputation)(nil)).Elem()
}

func (i *Internetservicereputation) ToInternetservicereputationOutput() InternetservicereputationOutput {
	return i.ToInternetservicereputationOutputWithContext(context.Background())
}

func (i *Internetservicereputation) ToInternetservicereputationOutputWithContext(ctx context.Context) InternetservicereputationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InternetservicereputationOutput)
}

// InternetservicereputationArrayInput is an input type that accepts InternetservicereputationArray and InternetservicereputationArrayOutput values.
// You can construct a concrete instance of `InternetservicereputationArrayInput` via:
//
//	InternetservicereputationArray{ InternetservicereputationArgs{...} }
type InternetservicereputationArrayInput interface {
	pulumi.Input

	ToInternetservicereputationArrayOutput() InternetservicereputationArrayOutput
	ToInternetservicereputationArrayOutputWithContext(context.Context) InternetservicereputationArrayOutput
}

type InternetservicereputationArray []InternetservicereputationInput

func (InternetservicereputationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Internetservicereputation)(nil)).Elem()
}

func (i InternetservicereputationArray) ToInternetservicereputationArrayOutput() InternetservicereputationArrayOutput {
	return i.ToInternetservicereputationArrayOutputWithContext(context.Background())
}

func (i InternetservicereputationArray) ToInternetservicereputationArrayOutputWithContext(ctx context.Context) InternetservicereputationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InternetservicereputationArrayOutput)
}

// InternetservicereputationMapInput is an input type that accepts InternetservicereputationMap and InternetservicereputationMapOutput values.
// You can construct a concrete instance of `InternetservicereputationMapInput` via:
//
//	InternetservicereputationMap{ "key": InternetservicereputationArgs{...} }
type InternetservicereputationMapInput interface {
	pulumi.Input

	ToInternetservicereputationMapOutput() InternetservicereputationMapOutput
	ToInternetservicereputationMapOutputWithContext(context.Context) InternetservicereputationMapOutput
}

type InternetservicereputationMap map[string]InternetservicereputationInput

func (InternetservicereputationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Internetservicereputation)(nil)).Elem()
}

func (i InternetservicereputationMap) ToInternetservicereputationMapOutput() InternetservicereputationMapOutput {
	return i.ToInternetservicereputationMapOutputWithContext(context.Background())
}

func (i InternetservicereputationMap) ToInternetservicereputationMapOutputWithContext(ctx context.Context) InternetservicereputationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InternetservicereputationMapOutput)
}

type InternetservicereputationOutput struct{ *pulumi.OutputState }

func (InternetservicereputationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Internetservicereputation)(nil)).Elem()
}

func (o InternetservicereputationOutput) ToInternetservicereputationOutput() InternetservicereputationOutput {
	return o
}

func (o InternetservicereputationOutput) ToInternetservicereputationOutputWithContext(ctx context.Context) InternetservicereputationOutput {
	return o
}

// Description.
func (o InternetservicereputationOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Internetservicereputation) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Internet Service Reputation ID.
func (o InternetservicereputationOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *Internetservicereputation) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o InternetservicereputationOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Internetservicereputation) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type InternetservicereputationArrayOutput struct{ *pulumi.OutputState }

func (InternetservicereputationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Internetservicereputation)(nil)).Elem()
}

func (o InternetservicereputationArrayOutput) ToInternetservicereputationArrayOutput() InternetservicereputationArrayOutput {
	return o
}

func (o InternetservicereputationArrayOutput) ToInternetservicereputationArrayOutputWithContext(ctx context.Context) InternetservicereputationArrayOutput {
	return o
}

func (o InternetservicereputationArrayOutput) Index(i pulumi.IntInput) InternetservicereputationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Internetservicereputation {
		return vs[0].([]*Internetservicereputation)[vs[1].(int)]
	}).(InternetservicereputationOutput)
}

type InternetservicereputationMapOutput struct{ *pulumi.OutputState }

func (InternetservicereputationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Internetservicereputation)(nil)).Elem()
}

func (o InternetservicereputationMapOutput) ToInternetservicereputationMapOutput() InternetservicereputationMapOutput {
	return o
}

func (o InternetservicereputationMapOutput) ToInternetservicereputationMapOutputWithContext(ctx context.Context) InternetservicereputationMapOutput {
	return o
}

func (o InternetservicereputationMapOutput) MapIndex(k pulumi.StringInput) InternetservicereputationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Internetservicereputation {
		return vs[0].(map[string]*Internetservicereputation)[vs[1].(string)]
	}).(InternetservicereputationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InternetservicereputationInput)(nil)).Elem(), &Internetservicereputation{})
	pulumi.RegisterInputType(reflect.TypeOf((*InternetservicereputationArrayInput)(nil)).Elem(), InternetservicereputationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InternetservicereputationMapInput)(nil)).Elem(), InternetservicereputationMap{})
	pulumi.RegisterOutputType(InternetservicereputationOutput{})
	pulumi.RegisterOutputType(InternetservicereputationArrayOutput{})
	pulumi.RegisterOutputType(InternetservicereputationMapOutput{})
}
