// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package router

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure IPv6 BFD.
//
// ## Import
//
// # Router Bfd6 can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:router/bfd6:Bfd6 labelname RouterBfd6
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:router/bfd6:Bfd6 labelname RouterBfd6
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type Bfd6 struct {
	pulumi.CustomResourceState

	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// BFD IPv6 multi-hop template table. The structure of `multihopTemplate` block is documented below.
	MultihopTemplates Bfd6MultihopTemplateArrayOutput `pulumi:"multihopTemplates"`
	// Configure neighbor of IPv6 BFD. The structure of `neighbor` block is documented below.
	Neighbors Bfd6NeighborArrayOutput `pulumi:"neighbors"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewBfd6 registers a new resource with the given unique name, arguments, and options.
func NewBfd6(ctx *pulumi.Context,
	name string, args *Bfd6Args, opts ...pulumi.ResourceOption) (*Bfd6, error) {
	if args == nil {
		args = &Bfd6Args{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource Bfd6
	err := ctx.RegisterResource("fortios:router/bfd6:Bfd6", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBfd6 gets an existing Bfd6 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBfd6(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Bfd6State, opts ...pulumi.ResourceOption) (*Bfd6, error) {
	var resource Bfd6
	err := ctx.ReadResource("fortios:router/bfd6:Bfd6", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Bfd6 resources.
type bfd6State struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// BFD IPv6 multi-hop template table. The structure of `multihopTemplate` block is documented below.
	MultihopTemplates []Bfd6MultihopTemplate `pulumi:"multihopTemplates"`
	// Configure neighbor of IPv6 BFD. The structure of `neighbor` block is documented below.
	Neighbors []Bfd6Neighbor `pulumi:"neighbors"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type Bfd6State struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// BFD IPv6 multi-hop template table. The structure of `multihopTemplate` block is documented below.
	MultihopTemplates Bfd6MultihopTemplateArrayInput
	// Configure neighbor of IPv6 BFD. The structure of `neighbor` block is documented below.
	Neighbors Bfd6NeighborArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (Bfd6State) ElementType() reflect.Type {
	return reflect.TypeOf((*bfd6State)(nil)).Elem()
}

type bfd6Args struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// BFD IPv6 multi-hop template table. The structure of `multihopTemplate` block is documented below.
	MultihopTemplates []Bfd6MultihopTemplate `pulumi:"multihopTemplates"`
	// Configure neighbor of IPv6 BFD. The structure of `neighbor` block is documented below.
	Neighbors []Bfd6Neighbor `pulumi:"neighbors"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Bfd6 resource.
type Bfd6Args struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// BFD IPv6 multi-hop template table. The structure of `multihopTemplate` block is documented below.
	MultihopTemplates Bfd6MultihopTemplateArrayInput
	// Configure neighbor of IPv6 BFD. The structure of `neighbor` block is documented below.
	Neighbors Bfd6NeighborArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (Bfd6Args) ElementType() reflect.Type {
	return reflect.TypeOf((*bfd6Args)(nil)).Elem()
}

type Bfd6Input interface {
	pulumi.Input

	ToBfd6Output() Bfd6Output
	ToBfd6OutputWithContext(ctx context.Context) Bfd6Output
}

func (*Bfd6) ElementType() reflect.Type {
	return reflect.TypeOf((**Bfd6)(nil)).Elem()
}

func (i *Bfd6) ToBfd6Output() Bfd6Output {
	return i.ToBfd6OutputWithContext(context.Background())
}

func (i *Bfd6) ToBfd6OutputWithContext(ctx context.Context) Bfd6Output {
	return pulumi.ToOutputWithContext(ctx, i).(Bfd6Output)
}

// Bfd6ArrayInput is an input type that accepts Bfd6Array and Bfd6ArrayOutput values.
// You can construct a concrete instance of `Bfd6ArrayInput` via:
//
//	Bfd6Array{ Bfd6Args{...} }
type Bfd6ArrayInput interface {
	pulumi.Input

	ToBfd6ArrayOutput() Bfd6ArrayOutput
	ToBfd6ArrayOutputWithContext(context.Context) Bfd6ArrayOutput
}

type Bfd6Array []Bfd6Input

func (Bfd6Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Bfd6)(nil)).Elem()
}

func (i Bfd6Array) ToBfd6ArrayOutput() Bfd6ArrayOutput {
	return i.ToBfd6ArrayOutputWithContext(context.Background())
}

func (i Bfd6Array) ToBfd6ArrayOutputWithContext(ctx context.Context) Bfd6ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Bfd6ArrayOutput)
}

// Bfd6MapInput is an input type that accepts Bfd6Map and Bfd6MapOutput values.
// You can construct a concrete instance of `Bfd6MapInput` via:
//
//	Bfd6Map{ "key": Bfd6Args{...} }
type Bfd6MapInput interface {
	pulumi.Input

	ToBfd6MapOutput() Bfd6MapOutput
	ToBfd6MapOutputWithContext(context.Context) Bfd6MapOutput
}

type Bfd6Map map[string]Bfd6Input

func (Bfd6Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Bfd6)(nil)).Elem()
}

func (i Bfd6Map) ToBfd6MapOutput() Bfd6MapOutput {
	return i.ToBfd6MapOutputWithContext(context.Background())
}

func (i Bfd6Map) ToBfd6MapOutputWithContext(ctx context.Context) Bfd6MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Bfd6MapOutput)
}

type Bfd6Output struct{ *pulumi.OutputState }

func (Bfd6Output) ElementType() reflect.Type {
	return reflect.TypeOf((**Bfd6)(nil)).Elem()
}

func (o Bfd6Output) ToBfd6Output() Bfd6Output {
	return o
}

func (o Bfd6Output) ToBfd6OutputWithContext(ctx context.Context) Bfd6Output {
	return o
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o Bfd6Output) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Bfd6) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// BFD IPv6 multi-hop template table. The structure of `multihopTemplate` block is documented below.
func (o Bfd6Output) MultihopTemplates() Bfd6MultihopTemplateArrayOutput {
	return o.ApplyT(func(v *Bfd6) Bfd6MultihopTemplateArrayOutput { return v.MultihopTemplates }).(Bfd6MultihopTemplateArrayOutput)
}

// Configure neighbor of IPv6 BFD. The structure of `neighbor` block is documented below.
func (o Bfd6Output) Neighbors() Bfd6NeighborArrayOutput {
	return o.ApplyT(func(v *Bfd6) Bfd6NeighborArrayOutput { return v.Neighbors }).(Bfd6NeighborArrayOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o Bfd6Output) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Bfd6) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type Bfd6ArrayOutput struct{ *pulumi.OutputState }

func (Bfd6ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Bfd6)(nil)).Elem()
}

func (o Bfd6ArrayOutput) ToBfd6ArrayOutput() Bfd6ArrayOutput {
	return o
}

func (o Bfd6ArrayOutput) ToBfd6ArrayOutputWithContext(ctx context.Context) Bfd6ArrayOutput {
	return o
}

func (o Bfd6ArrayOutput) Index(i pulumi.IntInput) Bfd6Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Bfd6 {
		return vs[0].([]*Bfd6)[vs[1].(int)]
	}).(Bfd6Output)
}

type Bfd6MapOutput struct{ *pulumi.OutputState }

func (Bfd6MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Bfd6)(nil)).Elem()
}

func (o Bfd6MapOutput) ToBfd6MapOutput() Bfd6MapOutput {
	return o
}

func (o Bfd6MapOutput) ToBfd6MapOutputWithContext(ctx context.Context) Bfd6MapOutput {
	return o
}

func (o Bfd6MapOutput) MapIndex(k pulumi.StringInput) Bfd6Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Bfd6 {
		return vs[0].(map[string]*Bfd6)[vs[1].(string)]
	}).(Bfd6Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*Bfd6Input)(nil)).Elem(), &Bfd6{})
	pulumi.RegisterInputType(reflect.TypeOf((*Bfd6ArrayInput)(nil)).Elem(), Bfd6Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*Bfd6MapInput)(nil)).Elem(), Bfd6Map{})
	pulumi.RegisterOutputType(Bfd6Output{})
	pulumi.RegisterOutputType(Bfd6ArrayOutput{})
	pulumi.RegisterOutputType(Bfd6MapOutput{})
}