// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package antivirus

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure global heuristic options. Applies to FortiOS Version `<= 7.0.0`.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/antivirus"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := antivirus.NewHeuristic(ctx, "trname", &antivirus.HeuristicArgs{
//				Mode: pulumi.String("disable"),
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
// # Antivirus Heuristic can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:antivirus/heuristic:Heuristic labelname AntivirusHeuristic
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:antivirus/heuristic:Heuristic labelname AntivirusHeuristic
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type Heuristic struct {
	pulumi.CustomResourceState

	// Enable/disable heuristics and determine how the system behaves if heuristics detects a problem. Valid values: `pass`, `block`, `disable`.
	Mode pulumi.StringOutput `pulumi:"mode"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewHeuristic registers a new resource with the given unique name, arguments, and options.
func NewHeuristic(ctx *pulumi.Context,
	name string, args *HeuristicArgs, opts ...pulumi.ResourceOption) (*Heuristic, error) {
	if args == nil {
		args = &HeuristicArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource Heuristic
	err := ctx.RegisterResource("fortios:antivirus/heuristic:Heuristic", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHeuristic gets an existing Heuristic resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHeuristic(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HeuristicState, opts ...pulumi.ResourceOption) (*Heuristic, error) {
	var resource Heuristic
	err := ctx.ReadResource("fortios:antivirus/heuristic:Heuristic", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Heuristic resources.
type heuristicState struct {
	// Enable/disable heuristics and determine how the system behaves if heuristics detects a problem. Valid values: `pass`, `block`, `disable`.
	Mode *string `pulumi:"mode"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type HeuristicState struct {
	// Enable/disable heuristics and determine how the system behaves if heuristics detects a problem. Valid values: `pass`, `block`, `disable`.
	Mode pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (HeuristicState) ElementType() reflect.Type {
	return reflect.TypeOf((*heuristicState)(nil)).Elem()
}

type heuristicArgs struct {
	// Enable/disable heuristics and determine how the system behaves if heuristics detects a problem. Valid values: `pass`, `block`, `disable`.
	Mode *string `pulumi:"mode"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Heuristic resource.
type HeuristicArgs struct {
	// Enable/disable heuristics and determine how the system behaves if heuristics detects a problem. Valid values: `pass`, `block`, `disable`.
	Mode pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (HeuristicArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*heuristicArgs)(nil)).Elem()
}

type HeuristicInput interface {
	pulumi.Input

	ToHeuristicOutput() HeuristicOutput
	ToHeuristicOutputWithContext(ctx context.Context) HeuristicOutput
}

func (*Heuristic) ElementType() reflect.Type {
	return reflect.TypeOf((**Heuristic)(nil)).Elem()
}

func (i *Heuristic) ToHeuristicOutput() HeuristicOutput {
	return i.ToHeuristicOutputWithContext(context.Background())
}

func (i *Heuristic) ToHeuristicOutputWithContext(ctx context.Context) HeuristicOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HeuristicOutput)
}

// HeuristicArrayInput is an input type that accepts HeuristicArray and HeuristicArrayOutput values.
// You can construct a concrete instance of `HeuristicArrayInput` via:
//
//	HeuristicArray{ HeuristicArgs{...} }
type HeuristicArrayInput interface {
	pulumi.Input

	ToHeuristicArrayOutput() HeuristicArrayOutput
	ToHeuristicArrayOutputWithContext(context.Context) HeuristicArrayOutput
}

type HeuristicArray []HeuristicInput

func (HeuristicArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Heuristic)(nil)).Elem()
}

func (i HeuristicArray) ToHeuristicArrayOutput() HeuristicArrayOutput {
	return i.ToHeuristicArrayOutputWithContext(context.Background())
}

func (i HeuristicArray) ToHeuristicArrayOutputWithContext(ctx context.Context) HeuristicArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HeuristicArrayOutput)
}

// HeuristicMapInput is an input type that accepts HeuristicMap and HeuristicMapOutput values.
// You can construct a concrete instance of `HeuristicMapInput` via:
//
//	HeuristicMap{ "key": HeuristicArgs{...} }
type HeuristicMapInput interface {
	pulumi.Input

	ToHeuristicMapOutput() HeuristicMapOutput
	ToHeuristicMapOutputWithContext(context.Context) HeuristicMapOutput
}

type HeuristicMap map[string]HeuristicInput

func (HeuristicMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Heuristic)(nil)).Elem()
}

func (i HeuristicMap) ToHeuristicMapOutput() HeuristicMapOutput {
	return i.ToHeuristicMapOutputWithContext(context.Background())
}

func (i HeuristicMap) ToHeuristicMapOutputWithContext(ctx context.Context) HeuristicMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HeuristicMapOutput)
}

type HeuristicOutput struct{ *pulumi.OutputState }

func (HeuristicOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Heuristic)(nil)).Elem()
}

func (o HeuristicOutput) ToHeuristicOutput() HeuristicOutput {
	return o
}

func (o HeuristicOutput) ToHeuristicOutputWithContext(ctx context.Context) HeuristicOutput {
	return o
}

// Enable/disable heuristics and determine how the system behaves if heuristics detects a problem. Valid values: `pass`, `block`, `disable`.
func (o HeuristicOutput) Mode() pulumi.StringOutput {
	return o.ApplyT(func(v *Heuristic) pulumi.StringOutput { return v.Mode }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o HeuristicOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Heuristic) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type HeuristicArrayOutput struct{ *pulumi.OutputState }

func (HeuristicArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Heuristic)(nil)).Elem()
}

func (o HeuristicArrayOutput) ToHeuristicArrayOutput() HeuristicArrayOutput {
	return o
}

func (o HeuristicArrayOutput) ToHeuristicArrayOutputWithContext(ctx context.Context) HeuristicArrayOutput {
	return o
}

func (o HeuristicArrayOutput) Index(i pulumi.IntInput) HeuristicOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Heuristic {
		return vs[0].([]*Heuristic)[vs[1].(int)]
	}).(HeuristicOutput)
}

type HeuristicMapOutput struct{ *pulumi.OutputState }

func (HeuristicMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Heuristic)(nil)).Elem()
}

func (o HeuristicMapOutput) ToHeuristicMapOutput() HeuristicMapOutput {
	return o
}

func (o HeuristicMapOutput) ToHeuristicMapOutputWithContext(ctx context.Context) HeuristicMapOutput {
	return o
}

func (o HeuristicMapOutput) MapIndex(k pulumi.StringInput) HeuristicOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Heuristic {
		return vs[0].(map[string]*Heuristic)[vs[1].(string)]
	}).(HeuristicOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HeuristicInput)(nil)).Elem(), &Heuristic{})
	pulumi.RegisterInputType(reflect.TypeOf((*HeuristicArrayInput)(nil)).Elem(), HeuristicArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*HeuristicMapInput)(nil)).Elem(), HeuristicMap{})
	pulumi.RegisterOutputType(HeuristicOutput{})
	pulumi.RegisterOutputType(HeuristicArrayOutput{})
	pulumi.RegisterOutputType(HeuristicMapOutput{})
}
