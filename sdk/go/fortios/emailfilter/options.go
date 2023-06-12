// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package emailfilter

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure AntiSpam options. Applies to FortiOS Version `>= 6.2.4`.
//
// ## Import
//
// # Emailfilter Options can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:emailfilter/options:Options labelname EmailfilterOptions
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:emailfilter/options:Options labelname EmailfilterOptions
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type Options struct {
	pulumi.CustomResourceState

	// DNS query time out (1 - 30 sec).
	DnsTimeout pulumi.IntOutput `pulumi:"dnsTimeout"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewOptions registers a new resource with the given unique name, arguments, and options.
func NewOptions(ctx *pulumi.Context,
	name string, args *OptionsArgs, opts ...pulumi.ResourceOption) (*Options, error) {
	if args == nil {
		args = &OptionsArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource Options
	err := ctx.RegisterResource("fortios:emailfilter/options:Options", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOptions gets an existing Options resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOptions(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OptionsState, opts ...pulumi.ResourceOption) (*Options, error) {
	var resource Options
	err := ctx.ReadResource("fortios:emailfilter/options:Options", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Options resources.
type optionsState struct {
	// DNS query time out (1 - 30 sec).
	DnsTimeout *int `pulumi:"dnsTimeout"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type OptionsState struct {
	// DNS query time out (1 - 30 sec).
	DnsTimeout pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (OptionsState) ElementType() reflect.Type {
	return reflect.TypeOf((*optionsState)(nil)).Elem()
}

type optionsArgs struct {
	// DNS query time out (1 - 30 sec).
	DnsTimeout *int `pulumi:"dnsTimeout"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Options resource.
type OptionsArgs struct {
	// DNS query time out (1 - 30 sec).
	DnsTimeout pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (OptionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*optionsArgs)(nil)).Elem()
}

type OptionsInput interface {
	pulumi.Input

	ToOptionsOutput() OptionsOutput
	ToOptionsOutputWithContext(ctx context.Context) OptionsOutput
}

func (*Options) ElementType() reflect.Type {
	return reflect.TypeOf((**Options)(nil)).Elem()
}

func (i *Options) ToOptionsOutput() OptionsOutput {
	return i.ToOptionsOutputWithContext(context.Background())
}

func (i *Options) ToOptionsOutputWithContext(ctx context.Context) OptionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OptionsOutput)
}

// OptionsArrayInput is an input type that accepts OptionsArray and OptionsArrayOutput values.
// You can construct a concrete instance of `OptionsArrayInput` via:
//
//	OptionsArray{ OptionsArgs{...} }
type OptionsArrayInput interface {
	pulumi.Input

	ToOptionsArrayOutput() OptionsArrayOutput
	ToOptionsArrayOutputWithContext(context.Context) OptionsArrayOutput
}

type OptionsArray []OptionsInput

func (OptionsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Options)(nil)).Elem()
}

func (i OptionsArray) ToOptionsArrayOutput() OptionsArrayOutput {
	return i.ToOptionsArrayOutputWithContext(context.Background())
}

func (i OptionsArray) ToOptionsArrayOutputWithContext(ctx context.Context) OptionsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OptionsArrayOutput)
}

// OptionsMapInput is an input type that accepts OptionsMap and OptionsMapOutput values.
// You can construct a concrete instance of `OptionsMapInput` via:
//
//	OptionsMap{ "key": OptionsArgs{...} }
type OptionsMapInput interface {
	pulumi.Input

	ToOptionsMapOutput() OptionsMapOutput
	ToOptionsMapOutputWithContext(context.Context) OptionsMapOutput
}

type OptionsMap map[string]OptionsInput

func (OptionsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Options)(nil)).Elem()
}

func (i OptionsMap) ToOptionsMapOutput() OptionsMapOutput {
	return i.ToOptionsMapOutputWithContext(context.Background())
}

func (i OptionsMap) ToOptionsMapOutputWithContext(ctx context.Context) OptionsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OptionsMapOutput)
}

type OptionsOutput struct{ *pulumi.OutputState }

func (OptionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Options)(nil)).Elem()
}

func (o OptionsOutput) ToOptionsOutput() OptionsOutput {
	return o
}

func (o OptionsOutput) ToOptionsOutputWithContext(ctx context.Context) OptionsOutput {
	return o
}

// DNS query time out (1 - 30 sec).
func (o OptionsOutput) DnsTimeout() pulumi.IntOutput {
	return o.ApplyT(func(v *Options) pulumi.IntOutput { return v.DnsTimeout }).(pulumi.IntOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o OptionsOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Options) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type OptionsArrayOutput struct{ *pulumi.OutputState }

func (OptionsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Options)(nil)).Elem()
}

func (o OptionsArrayOutput) ToOptionsArrayOutput() OptionsArrayOutput {
	return o
}

func (o OptionsArrayOutput) ToOptionsArrayOutputWithContext(ctx context.Context) OptionsArrayOutput {
	return o
}

func (o OptionsArrayOutput) Index(i pulumi.IntInput) OptionsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Options {
		return vs[0].([]*Options)[vs[1].(int)]
	}).(OptionsOutput)
}

type OptionsMapOutput struct{ *pulumi.OutputState }

func (OptionsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Options)(nil)).Elem()
}

func (o OptionsMapOutput) ToOptionsMapOutput() OptionsMapOutput {
	return o
}

func (o OptionsMapOutput) ToOptionsMapOutputWithContext(ctx context.Context) OptionsMapOutput {
	return o
}

func (o OptionsMapOutput) MapIndex(k pulumi.StringInput) OptionsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Options {
		return vs[0].(map[string]*Options)[vs[1].(string)]
	}).(OptionsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OptionsInput)(nil)).Elem(), &Options{})
	pulumi.RegisterInputType(reflect.TypeOf((*OptionsArrayInput)(nil)).Elem(), OptionsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OptionsMapInput)(nil)).Elem(), OptionsMap{})
	pulumi.RegisterOutputType(OptionsOutput{})
	pulumi.RegisterOutputType(OptionsArrayOutput{})
	pulumi.RegisterOutputType(OptionsMapOutput{})
}
