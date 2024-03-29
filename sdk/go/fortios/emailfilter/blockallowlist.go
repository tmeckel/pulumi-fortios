// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package emailfilter

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure anti-spam block/allow list. Applies to FortiOS Version `>= 7.0.0`.
//
// ## Import
//
// # Emailfilter BlockAllowList can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:emailfilter/blockallowlist:Blockallowlist labelname {{fosid}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:emailfilter/blockallowlist:Blockallowlist labelname {{fosid}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type Blockallowlist struct {
	pulumi.CustomResourceState

	// Optional comments.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Anti-spam block/allow entries. The structure of `entries` block is documented below.
	Entries BlockallowlistEntryArrayOutput `pulumi:"entries"`
	// ID.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// Name of table.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewBlockallowlist registers a new resource with the given unique name, arguments, and options.
func NewBlockallowlist(ctx *pulumi.Context,
	name string, args *BlockallowlistArgs, opts ...pulumi.ResourceOption) (*Blockallowlist, error) {
	if args == nil {
		args = &BlockallowlistArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource Blockallowlist
	err := ctx.RegisterResource("fortios:emailfilter/blockallowlist:Blockallowlist", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBlockallowlist gets an existing Blockallowlist resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBlockallowlist(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BlockallowlistState, opts ...pulumi.ResourceOption) (*Blockallowlist, error) {
	var resource Blockallowlist
	err := ctx.ReadResource("fortios:emailfilter/blockallowlist:Blockallowlist", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Blockallowlist resources.
type blockallowlistState struct {
	// Optional comments.
	Comment *string `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Anti-spam block/allow entries. The structure of `entries` block is documented below.
	Entries []BlockallowlistEntry `pulumi:"entries"`
	// ID.
	Fosid *int `pulumi:"fosid"`
	// Name of table.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type BlockallowlistState struct {
	// Optional comments.
	Comment pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Anti-spam block/allow entries. The structure of `entries` block is documented below.
	Entries BlockallowlistEntryArrayInput
	// ID.
	Fosid pulumi.IntPtrInput
	// Name of table.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (BlockallowlistState) ElementType() reflect.Type {
	return reflect.TypeOf((*blockallowlistState)(nil)).Elem()
}

type blockallowlistArgs struct {
	// Optional comments.
	Comment *string `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Anti-spam block/allow entries. The structure of `entries` block is documented below.
	Entries []BlockallowlistEntry `pulumi:"entries"`
	// ID.
	Fosid *int `pulumi:"fosid"`
	// Name of table.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Blockallowlist resource.
type BlockallowlistArgs struct {
	// Optional comments.
	Comment pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Anti-spam block/allow entries. The structure of `entries` block is documented below.
	Entries BlockallowlistEntryArrayInput
	// ID.
	Fosid pulumi.IntPtrInput
	// Name of table.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (BlockallowlistArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*blockallowlistArgs)(nil)).Elem()
}

type BlockallowlistInput interface {
	pulumi.Input

	ToBlockallowlistOutput() BlockallowlistOutput
	ToBlockallowlistOutputWithContext(ctx context.Context) BlockallowlistOutput
}

func (*Blockallowlist) ElementType() reflect.Type {
	return reflect.TypeOf((**Blockallowlist)(nil)).Elem()
}

func (i *Blockallowlist) ToBlockallowlistOutput() BlockallowlistOutput {
	return i.ToBlockallowlistOutputWithContext(context.Background())
}

func (i *Blockallowlist) ToBlockallowlistOutputWithContext(ctx context.Context) BlockallowlistOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlockallowlistOutput)
}

// BlockallowlistArrayInput is an input type that accepts BlockallowlistArray and BlockallowlistArrayOutput values.
// You can construct a concrete instance of `BlockallowlistArrayInput` via:
//
//	BlockallowlistArray{ BlockallowlistArgs{...} }
type BlockallowlistArrayInput interface {
	pulumi.Input

	ToBlockallowlistArrayOutput() BlockallowlistArrayOutput
	ToBlockallowlistArrayOutputWithContext(context.Context) BlockallowlistArrayOutput
}

type BlockallowlistArray []BlockallowlistInput

func (BlockallowlistArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Blockallowlist)(nil)).Elem()
}

func (i BlockallowlistArray) ToBlockallowlistArrayOutput() BlockallowlistArrayOutput {
	return i.ToBlockallowlistArrayOutputWithContext(context.Background())
}

func (i BlockallowlistArray) ToBlockallowlistArrayOutputWithContext(ctx context.Context) BlockallowlistArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlockallowlistArrayOutput)
}

// BlockallowlistMapInput is an input type that accepts BlockallowlistMap and BlockallowlistMapOutput values.
// You can construct a concrete instance of `BlockallowlistMapInput` via:
//
//	BlockallowlistMap{ "key": BlockallowlistArgs{...} }
type BlockallowlistMapInput interface {
	pulumi.Input

	ToBlockallowlistMapOutput() BlockallowlistMapOutput
	ToBlockallowlistMapOutputWithContext(context.Context) BlockallowlistMapOutput
}

type BlockallowlistMap map[string]BlockallowlistInput

func (BlockallowlistMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Blockallowlist)(nil)).Elem()
}

func (i BlockallowlistMap) ToBlockallowlistMapOutput() BlockallowlistMapOutput {
	return i.ToBlockallowlistMapOutputWithContext(context.Background())
}

func (i BlockallowlistMap) ToBlockallowlistMapOutputWithContext(ctx context.Context) BlockallowlistMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlockallowlistMapOutput)
}

type BlockallowlistOutput struct{ *pulumi.OutputState }

func (BlockallowlistOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Blockallowlist)(nil)).Elem()
}

func (o BlockallowlistOutput) ToBlockallowlistOutput() BlockallowlistOutput {
	return o
}

func (o BlockallowlistOutput) ToBlockallowlistOutputWithContext(ctx context.Context) BlockallowlistOutput {
	return o
}

// Optional comments.
func (o BlockallowlistOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Blockallowlist) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o BlockallowlistOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Blockallowlist) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Anti-spam block/allow entries. The structure of `entries` block is documented below.
func (o BlockallowlistOutput) Entries() BlockallowlistEntryArrayOutput {
	return o.ApplyT(func(v *Blockallowlist) BlockallowlistEntryArrayOutput { return v.Entries }).(BlockallowlistEntryArrayOutput)
}

// ID.
func (o BlockallowlistOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *Blockallowlist) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

// Name of table.
func (o BlockallowlistOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Blockallowlist) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o BlockallowlistOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Blockallowlist) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type BlockallowlistArrayOutput struct{ *pulumi.OutputState }

func (BlockallowlistArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Blockallowlist)(nil)).Elem()
}

func (o BlockallowlistArrayOutput) ToBlockallowlistArrayOutput() BlockallowlistArrayOutput {
	return o
}

func (o BlockallowlistArrayOutput) ToBlockallowlistArrayOutputWithContext(ctx context.Context) BlockallowlistArrayOutput {
	return o
}

func (o BlockallowlistArrayOutput) Index(i pulumi.IntInput) BlockallowlistOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Blockallowlist {
		return vs[0].([]*Blockallowlist)[vs[1].(int)]
	}).(BlockallowlistOutput)
}

type BlockallowlistMapOutput struct{ *pulumi.OutputState }

func (BlockallowlistMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Blockallowlist)(nil)).Elem()
}

func (o BlockallowlistMapOutput) ToBlockallowlistMapOutput() BlockallowlistMapOutput {
	return o
}

func (o BlockallowlistMapOutput) ToBlockallowlistMapOutputWithContext(ctx context.Context) BlockallowlistMapOutput {
	return o
}

func (o BlockallowlistMapOutput) MapIndex(k pulumi.StringInput) BlockallowlistOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Blockallowlist {
		return vs[0].(map[string]*Blockallowlist)[vs[1].(string)]
	}).(BlockallowlistOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BlockallowlistInput)(nil)).Elem(), &Blockallowlist{})
	pulumi.RegisterInputType(reflect.TypeOf((*BlockallowlistArrayInput)(nil)).Elem(), BlockallowlistArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BlockallowlistMapInput)(nil)).Elem(), BlockallowlistMap{})
	pulumi.RegisterOutputType(BlockallowlistOutput{})
	pulumi.RegisterOutputType(BlockallowlistArrayOutput{})
	pulumi.RegisterOutputType(BlockallowlistMapOutput{})
}
