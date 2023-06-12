// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package router

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure IPv6 prefix lists.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/router"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := router.NewPrefixlist6(ctx, "trname", nil)
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
// # Router PrefixList6 can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:router/prefixlist6:Prefixlist6 labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:router/prefixlist6:Prefixlist6 labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type Prefixlist6 struct {
	pulumi.CustomResourceState

	// Comment.
	Comments pulumi.StringOutput `pulumi:"comments"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// IPv6 prefix list rule. The structure of `rule` block is documented below.
	Rules Prefixlist6RuleArrayOutput `pulumi:"rules"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewPrefixlist6 registers a new resource with the given unique name, arguments, and options.
func NewPrefixlist6(ctx *pulumi.Context,
	name string, args *Prefixlist6Args, opts ...pulumi.ResourceOption) (*Prefixlist6, error) {
	if args == nil {
		args = &Prefixlist6Args{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource Prefixlist6
	err := ctx.RegisterResource("fortios:router/prefixlist6:Prefixlist6", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrefixlist6 gets an existing Prefixlist6 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrefixlist6(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Prefixlist6State, opts ...pulumi.ResourceOption) (*Prefixlist6, error) {
	var resource Prefixlist6
	err := ctx.ReadResource("fortios:router/prefixlist6:Prefixlist6", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Prefixlist6 resources.
type prefixlist6State struct {
	// Comment.
	Comments *string `pulumi:"comments"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Name.
	Name *string `pulumi:"name"`
	// IPv6 prefix list rule. The structure of `rule` block is documented below.
	Rules []Prefixlist6Rule `pulumi:"rules"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type Prefixlist6State struct {
	// Comment.
	Comments pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Name.
	Name pulumi.StringPtrInput
	// IPv6 prefix list rule. The structure of `rule` block is documented below.
	Rules Prefixlist6RuleArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (Prefixlist6State) ElementType() reflect.Type {
	return reflect.TypeOf((*prefixlist6State)(nil)).Elem()
}

type prefixlist6Args struct {
	// Comment.
	Comments *string `pulumi:"comments"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Name.
	Name *string `pulumi:"name"`
	// IPv6 prefix list rule. The structure of `rule` block is documented below.
	Rules []Prefixlist6Rule `pulumi:"rules"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Prefixlist6 resource.
type Prefixlist6Args struct {
	// Comment.
	Comments pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Name.
	Name pulumi.StringPtrInput
	// IPv6 prefix list rule. The structure of `rule` block is documented below.
	Rules Prefixlist6RuleArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (Prefixlist6Args) ElementType() reflect.Type {
	return reflect.TypeOf((*prefixlist6Args)(nil)).Elem()
}

type Prefixlist6Input interface {
	pulumi.Input

	ToPrefixlist6Output() Prefixlist6Output
	ToPrefixlist6OutputWithContext(ctx context.Context) Prefixlist6Output
}

func (*Prefixlist6) ElementType() reflect.Type {
	return reflect.TypeOf((**Prefixlist6)(nil)).Elem()
}

func (i *Prefixlist6) ToPrefixlist6Output() Prefixlist6Output {
	return i.ToPrefixlist6OutputWithContext(context.Background())
}

func (i *Prefixlist6) ToPrefixlist6OutputWithContext(ctx context.Context) Prefixlist6Output {
	return pulumi.ToOutputWithContext(ctx, i).(Prefixlist6Output)
}

// Prefixlist6ArrayInput is an input type that accepts Prefixlist6Array and Prefixlist6ArrayOutput values.
// You can construct a concrete instance of `Prefixlist6ArrayInput` via:
//
//	Prefixlist6Array{ Prefixlist6Args{...} }
type Prefixlist6ArrayInput interface {
	pulumi.Input

	ToPrefixlist6ArrayOutput() Prefixlist6ArrayOutput
	ToPrefixlist6ArrayOutputWithContext(context.Context) Prefixlist6ArrayOutput
}

type Prefixlist6Array []Prefixlist6Input

func (Prefixlist6Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Prefixlist6)(nil)).Elem()
}

func (i Prefixlist6Array) ToPrefixlist6ArrayOutput() Prefixlist6ArrayOutput {
	return i.ToPrefixlist6ArrayOutputWithContext(context.Background())
}

func (i Prefixlist6Array) ToPrefixlist6ArrayOutputWithContext(ctx context.Context) Prefixlist6ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Prefixlist6ArrayOutput)
}

// Prefixlist6MapInput is an input type that accepts Prefixlist6Map and Prefixlist6MapOutput values.
// You can construct a concrete instance of `Prefixlist6MapInput` via:
//
//	Prefixlist6Map{ "key": Prefixlist6Args{...} }
type Prefixlist6MapInput interface {
	pulumi.Input

	ToPrefixlist6MapOutput() Prefixlist6MapOutput
	ToPrefixlist6MapOutputWithContext(context.Context) Prefixlist6MapOutput
}

type Prefixlist6Map map[string]Prefixlist6Input

func (Prefixlist6Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Prefixlist6)(nil)).Elem()
}

func (i Prefixlist6Map) ToPrefixlist6MapOutput() Prefixlist6MapOutput {
	return i.ToPrefixlist6MapOutputWithContext(context.Background())
}

func (i Prefixlist6Map) ToPrefixlist6MapOutputWithContext(ctx context.Context) Prefixlist6MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Prefixlist6MapOutput)
}

type Prefixlist6Output struct{ *pulumi.OutputState }

func (Prefixlist6Output) ElementType() reflect.Type {
	return reflect.TypeOf((**Prefixlist6)(nil)).Elem()
}

func (o Prefixlist6Output) ToPrefixlist6Output() Prefixlist6Output {
	return o
}

func (o Prefixlist6Output) ToPrefixlist6OutputWithContext(ctx context.Context) Prefixlist6Output {
	return o
}

// Comment.
func (o Prefixlist6Output) Comments() pulumi.StringOutput {
	return o.ApplyT(func(v *Prefixlist6) pulumi.StringOutput { return v.Comments }).(pulumi.StringOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o Prefixlist6Output) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Prefixlist6) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Name.
func (o Prefixlist6Output) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Prefixlist6) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// IPv6 prefix list rule. The structure of `rule` block is documented below.
func (o Prefixlist6Output) Rules() Prefixlist6RuleArrayOutput {
	return o.ApplyT(func(v *Prefixlist6) Prefixlist6RuleArrayOutput { return v.Rules }).(Prefixlist6RuleArrayOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o Prefixlist6Output) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Prefixlist6) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type Prefixlist6ArrayOutput struct{ *pulumi.OutputState }

func (Prefixlist6ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Prefixlist6)(nil)).Elem()
}

func (o Prefixlist6ArrayOutput) ToPrefixlist6ArrayOutput() Prefixlist6ArrayOutput {
	return o
}

func (o Prefixlist6ArrayOutput) ToPrefixlist6ArrayOutputWithContext(ctx context.Context) Prefixlist6ArrayOutput {
	return o
}

func (o Prefixlist6ArrayOutput) Index(i pulumi.IntInput) Prefixlist6Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Prefixlist6 {
		return vs[0].([]*Prefixlist6)[vs[1].(int)]
	}).(Prefixlist6Output)
}

type Prefixlist6MapOutput struct{ *pulumi.OutputState }

func (Prefixlist6MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Prefixlist6)(nil)).Elem()
}

func (o Prefixlist6MapOutput) ToPrefixlist6MapOutput() Prefixlist6MapOutput {
	return o
}

func (o Prefixlist6MapOutput) ToPrefixlist6MapOutputWithContext(ctx context.Context) Prefixlist6MapOutput {
	return o
}

func (o Prefixlist6MapOutput) MapIndex(k pulumi.StringInput) Prefixlist6Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Prefixlist6 {
		return vs[0].(map[string]*Prefixlist6)[vs[1].(string)]
	}).(Prefixlist6Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*Prefixlist6Input)(nil)).Elem(), &Prefixlist6{})
	pulumi.RegisterInputType(reflect.TypeOf((*Prefixlist6ArrayInput)(nil)).Elem(), Prefixlist6Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*Prefixlist6MapInput)(nil)).Elem(), Prefixlist6Map{})
	pulumi.RegisterOutputType(Prefixlist6Output{})
	pulumi.RegisterOutputType(Prefixlist6ArrayOutput{})
	pulumi.RegisterOutputType(Prefixlist6MapOutput{})
}