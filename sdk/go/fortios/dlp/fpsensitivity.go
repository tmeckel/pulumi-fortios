// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dlp

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Create self-explanatory DLP sensitivity levels to be used when setting sensitivity under config fp-doc-source. Applies to FortiOS Version `<= 6.2.0`.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/dlp"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := dlp.NewFpsensitivity(ctx, "trname", nil)
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
// # Dlp FpSensitivity can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:dlp/fpsensitivity:Fpsensitivity labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:dlp/fpsensitivity:Fpsensitivity labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type Fpsensitivity struct {
	pulumi.CustomResourceState

	// DLP Sensitivity Levels.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewFpsensitivity registers a new resource with the given unique name, arguments, and options.
func NewFpsensitivity(ctx *pulumi.Context,
	name string, args *FpsensitivityArgs, opts ...pulumi.ResourceOption) (*Fpsensitivity, error) {
	if args == nil {
		args = &FpsensitivityArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource Fpsensitivity
	err := ctx.RegisterResource("fortios:dlp/fpsensitivity:Fpsensitivity", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFpsensitivity gets an existing Fpsensitivity resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFpsensitivity(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FpsensitivityState, opts ...pulumi.ResourceOption) (*Fpsensitivity, error) {
	var resource Fpsensitivity
	err := ctx.ReadResource("fortios:dlp/fpsensitivity:Fpsensitivity", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Fpsensitivity resources.
type fpsensitivityState struct {
	// DLP Sensitivity Levels.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type FpsensitivityState struct {
	// DLP Sensitivity Levels.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FpsensitivityState) ElementType() reflect.Type {
	return reflect.TypeOf((*fpsensitivityState)(nil)).Elem()
}

type fpsensitivityArgs struct {
	// DLP Sensitivity Levels.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Fpsensitivity resource.
type FpsensitivityArgs struct {
	// DLP Sensitivity Levels.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FpsensitivityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fpsensitivityArgs)(nil)).Elem()
}

type FpsensitivityInput interface {
	pulumi.Input

	ToFpsensitivityOutput() FpsensitivityOutput
	ToFpsensitivityOutputWithContext(ctx context.Context) FpsensitivityOutput
}

func (*Fpsensitivity) ElementType() reflect.Type {
	return reflect.TypeOf((**Fpsensitivity)(nil)).Elem()
}

func (i *Fpsensitivity) ToFpsensitivityOutput() FpsensitivityOutput {
	return i.ToFpsensitivityOutputWithContext(context.Background())
}

func (i *Fpsensitivity) ToFpsensitivityOutputWithContext(ctx context.Context) FpsensitivityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FpsensitivityOutput)
}

// FpsensitivityArrayInput is an input type that accepts FpsensitivityArray and FpsensitivityArrayOutput values.
// You can construct a concrete instance of `FpsensitivityArrayInput` via:
//
//	FpsensitivityArray{ FpsensitivityArgs{...} }
type FpsensitivityArrayInput interface {
	pulumi.Input

	ToFpsensitivityArrayOutput() FpsensitivityArrayOutput
	ToFpsensitivityArrayOutputWithContext(context.Context) FpsensitivityArrayOutput
}

type FpsensitivityArray []FpsensitivityInput

func (FpsensitivityArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Fpsensitivity)(nil)).Elem()
}

func (i FpsensitivityArray) ToFpsensitivityArrayOutput() FpsensitivityArrayOutput {
	return i.ToFpsensitivityArrayOutputWithContext(context.Background())
}

func (i FpsensitivityArray) ToFpsensitivityArrayOutputWithContext(ctx context.Context) FpsensitivityArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FpsensitivityArrayOutput)
}

// FpsensitivityMapInput is an input type that accepts FpsensitivityMap and FpsensitivityMapOutput values.
// You can construct a concrete instance of `FpsensitivityMapInput` via:
//
//	FpsensitivityMap{ "key": FpsensitivityArgs{...} }
type FpsensitivityMapInput interface {
	pulumi.Input

	ToFpsensitivityMapOutput() FpsensitivityMapOutput
	ToFpsensitivityMapOutputWithContext(context.Context) FpsensitivityMapOutput
}

type FpsensitivityMap map[string]FpsensitivityInput

func (FpsensitivityMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Fpsensitivity)(nil)).Elem()
}

func (i FpsensitivityMap) ToFpsensitivityMapOutput() FpsensitivityMapOutput {
	return i.ToFpsensitivityMapOutputWithContext(context.Background())
}

func (i FpsensitivityMap) ToFpsensitivityMapOutputWithContext(ctx context.Context) FpsensitivityMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FpsensitivityMapOutput)
}

type FpsensitivityOutput struct{ *pulumi.OutputState }

func (FpsensitivityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Fpsensitivity)(nil)).Elem()
}

func (o FpsensitivityOutput) ToFpsensitivityOutput() FpsensitivityOutput {
	return o
}

func (o FpsensitivityOutput) ToFpsensitivityOutputWithContext(ctx context.Context) FpsensitivityOutput {
	return o
}

// DLP Sensitivity Levels.
func (o FpsensitivityOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Fpsensitivity) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o FpsensitivityOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Fpsensitivity) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type FpsensitivityArrayOutput struct{ *pulumi.OutputState }

func (FpsensitivityArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Fpsensitivity)(nil)).Elem()
}

func (o FpsensitivityArrayOutput) ToFpsensitivityArrayOutput() FpsensitivityArrayOutput {
	return o
}

func (o FpsensitivityArrayOutput) ToFpsensitivityArrayOutputWithContext(ctx context.Context) FpsensitivityArrayOutput {
	return o
}

func (o FpsensitivityArrayOutput) Index(i pulumi.IntInput) FpsensitivityOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Fpsensitivity {
		return vs[0].([]*Fpsensitivity)[vs[1].(int)]
	}).(FpsensitivityOutput)
}

type FpsensitivityMapOutput struct{ *pulumi.OutputState }

func (FpsensitivityMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Fpsensitivity)(nil)).Elem()
}

func (o FpsensitivityMapOutput) ToFpsensitivityMapOutput() FpsensitivityMapOutput {
	return o
}

func (o FpsensitivityMapOutput) ToFpsensitivityMapOutputWithContext(ctx context.Context) FpsensitivityMapOutput {
	return o
}

func (o FpsensitivityMapOutput) MapIndex(k pulumi.StringInput) FpsensitivityOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Fpsensitivity {
		return vs[0].(map[string]*Fpsensitivity)[vs[1].(string)]
	}).(FpsensitivityOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FpsensitivityInput)(nil)).Elem(), &Fpsensitivity{})
	pulumi.RegisterInputType(reflect.TypeOf((*FpsensitivityArrayInput)(nil)).Elem(), FpsensitivityArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FpsensitivityMapInput)(nil)).Elem(), FpsensitivityMap{})
	pulumi.RegisterOutputType(FpsensitivityOutput{})
	pulumi.RegisterOutputType(FpsensitivityArrayOutput{})
	pulumi.RegisterOutputType(FpsensitivityMapOutput{})
}