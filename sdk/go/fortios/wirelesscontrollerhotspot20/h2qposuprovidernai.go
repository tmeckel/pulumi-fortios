// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package wirelesscontrollerhotspot20

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure online sign up (OSU) provider NAI list. Applies to FortiOS Version `>= 7.0.2`.
//
// ## Import
//
// # WirelessControllerHotspot20 H2QpOsuProviderNai can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:wirelesscontrollerhotspot20/h2qposuprovidernai:H2qposuprovidernai labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:wirelesscontrollerhotspot20/h2qposuprovidernai:H2qposuprovidernai labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type H2qposuprovidernai struct {
	pulumi.CustomResourceState

	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// OSU NAI list. The structure of `naiList` block is documented below.
	NaiLists H2qposuprovidernaiNaiListArrayOutput `pulumi:"naiLists"`
	// OSU provider NAI ID.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewH2qposuprovidernai registers a new resource with the given unique name, arguments, and options.
func NewH2qposuprovidernai(ctx *pulumi.Context,
	name string, args *H2qposuprovidernaiArgs, opts ...pulumi.ResourceOption) (*H2qposuprovidernai, error) {
	if args == nil {
		args = &H2qposuprovidernaiArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource H2qposuprovidernai
	err := ctx.RegisterResource("fortios:wirelesscontrollerhotspot20/h2qposuprovidernai:H2qposuprovidernai", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetH2qposuprovidernai gets an existing H2qposuprovidernai resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetH2qposuprovidernai(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *H2qposuprovidernaiState, opts ...pulumi.ResourceOption) (*H2qposuprovidernai, error) {
	var resource H2qposuprovidernai
	err := ctx.ReadResource("fortios:wirelesscontrollerhotspot20/h2qposuprovidernai:H2qposuprovidernai", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering H2qposuprovidernai resources.
type h2qposuprovidernaiState struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// OSU NAI list. The structure of `naiList` block is documented below.
	NaiLists []H2qposuprovidernaiNaiList `pulumi:"naiLists"`
	// OSU provider NAI ID.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type H2qposuprovidernaiState struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// OSU NAI list. The structure of `naiList` block is documented below.
	NaiLists H2qposuprovidernaiNaiListArrayInput
	// OSU provider NAI ID.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (H2qposuprovidernaiState) ElementType() reflect.Type {
	return reflect.TypeOf((*h2qposuprovidernaiState)(nil)).Elem()
}

type h2qposuprovidernaiArgs struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// OSU NAI list. The structure of `naiList` block is documented below.
	NaiLists []H2qposuprovidernaiNaiList `pulumi:"naiLists"`
	// OSU provider NAI ID.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a H2qposuprovidernai resource.
type H2qposuprovidernaiArgs struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// OSU NAI list. The structure of `naiList` block is documented below.
	NaiLists H2qposuprovidernaiNaiListArrayInput
	// OSU provider NAI ID.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (H2qposuprovidernaiArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*h2qposuprovidernaiArgs)(nil)).Elem()
}

type H2qposuprovidernaiInput interface {
	pulumi.Input

	ToH2qposuprovidernaiOutput() H2qposuprovidernaiOutput
	ToH2qposuprovidernaiOutputWithContext(ctx context.Context) H2qposuprovidernaiOutput
}

func (*H2qposuprovidernai) ElementType() reflect.Type {
	return reflect.TypeOf((**H2qposuprovidernai)(nil)).Elem()
}

func (i *H2qposuprovidernai) ToH2qposuprovidernaiOutput() H2qposuprovidernaiOutput {
	return i.ToH2qposuprovidernaiOutputWithContext(context.Background())
}

func (i *H2qposuprovidernai) ToH2qposuprovidernaiOutputWithContext(ctx context.Context) H2qposuprovidernaiOutput {
	return pulumi.ToOutputWithContext(ctx, i).(H2qposuprovidernaiOutput)
}

// H2qposuprovidernaiArrayInput is an input type that accepts H2qposuprovidernaiArray and H2qposuprovidernaiArrayOutput values.
// You can construct a concrete instance of `H2qposuprovidernaiArrayInput` via:
//
//	H2qposuprovidernaiArray{ H2qposuprovidernaiArgs{...} }
type H2qposuprovidernaiArrayInput interface {
	pulumi.Input

	ToH2qposuprovidernaiArrayOutput() H2qposuprovidernaiArrayOutput
	ToH2qposuprovidernaiArrayOutputWithContext(context.Context) H2qposuprovidernaiArrayOutput
}

type H2qposuprovidernaiArray []H2qposuprovidernaiInput

func (H2qposuprovidernaiArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*H2qposuprovidernai)(nil)).Elem()
}

func (i H2qposuprovidernaiArray) ToH2qposuprovidernaiArrayOutput() H2qposuprovidernaiArrayOutput {
	return i.ToH2qposuprovidernaiArrayOutputWithContext(context.Background())
}

func (i H2qposuprovidernaiArray) ToH2qposuprovidernaiArrayOutputWithContext(ctx context.Context) H2qposuprovidernaiArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(H2qposuprovidernaiArrayOutput)
}

// H2qposuprovidernaiMapInput is an input type that accepts H2qposuprovidernaiMap and H2qposuprovidernaiMapOutput values.
// You can construct a concrete instance of `H2qposuprovidernaiMapInput` via:
//
//	H2qposuprovidernaiMap{ "key": H2qposuprovidernaiArgs{...} }
type H2qposuprovidernaiMapInput interface {
	pulumi.Input

	ToH2qposuprovidernaiMapOutput() H2qposuprovidernaiMapOutput
	ToH2qposuprovidernaiMapOutputWithContext(context.Context) H2qposuprovidernaiMapOutput
}

type H2qposuprovidernaiMap map[string]H2qposuprovidernaiInput

func (H2qposuprovidernaiMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*H2qposuprovidernai)(nil)).Elem()
}

func (i H2qposuprovidernaiMap) ToH2qposuprovidernaiMapOutput() H2qposuprovidernaiMapOutput {
	return i.ToH2qposuprovidernaiMapOutputWithContext(context.Background())
}

func (i H2qposuprovidernaiMap) ToH2qposuprovidernaiMapOutputWithContext(ctx context.Context) H2qposuprovidernaiMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(H2qposuprovidernaiMapOutput)
}

type H2qposuprovidernaiOutput struct{ *pulumi.OutputState }

func (H2qposuprovidernaiOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**H2qposuprovidernai)(nil)).Elem()
}

func (o H2qposuprovidernaiOutput) ToH2qposuprovidernaiOutput() H2qposuprovidernaiOutput {
	return o
}

func (o H2qposuprovidernaiOutput) ToH2qposuprovidernaiOutputWithContext(ctx context.Context) H2qposuprovidernaiOutput {
	return o
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o H2qposuprovidernaiOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *H2qposuprovidernai) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// OSU NAI list. The structure of `naiList` block is documented below.
func (o H2qposuprovidernaiOutput) NaiLists() H2qposuprovidernaiNaiListArrayOutput {
	return o.ApplyT(func(v *H2qposuprovidernai) H2qposuprovidernaiNaiListArrayOutput { return v.NaiLists }).(H2qposuprovidernaiNaiListArrayOutput)
}

// OSU provider NAI ID.
func (o H2qposuprovidernaiOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *H2qposuprovidernai) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o H2qposuprovidernaiOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *H2qposuprovidernai) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type H2qposuprovidernaiArrayOutput struct{ *pulumi.OutputState }

func (H2qposuprovidernaiArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*H2qposuprovidernai)(nil)).Elem()
}

func (o H2qposuprovidernaiArrayOutput) ToH2qposuprovidernaiArrayOutput() H2qposuprovidernaiArrayOutput {
	return o
}

func (o H2qposuprovidernaiArrayOutput) ToH2qposuprovidernaiArrayOutputWithContext(ctx context.Context) H2qposuprovidernaiArrayOutput {
	return o
}

func (o H2qposuprovidernaiArrayOutput) Index(i pulumi.IntInput) H2qposuprovidernaiOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *H2qposuprovidernai {
		return vs[0].([]*H2qposuprovidernai)[vs[1].(int)]
	}).(H2qposuprovidernaiOutput)
}

type H2qposuprovidernaiMapOutput struct{ *pulumi.OutputState }

func (H2qposuprovidernaiMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*H2qposuprovidernai)(nil)).Elem()
}

func (o H2qposuprovidernaiMapOutput) ToH2qposuprovidernaiMapOutput() H2qposuprovidernaiMapOutput {
	return o
}

func (o H2qposuprovidernaiMapOutput) ToH2qposuprovidernaiMapOutputWithContext(ctx context.Context) H2qposuprovidernaiMapOutput {
	return o
}

func (o H2qposuprovidernaiMapOutput) MapIndex(k pulumi.StringInput) H2qposuprovidernaiOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *H2qposuprovidernai {
		return vs[0].(map[string]*H2qposuprovidernai)[vs[1].(string)]
	}).(H2qposuprovidernaiOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*H2qposuprovidernaiInput)(nil)).Elem(), &H2qposuprovidernai{})
	pulumi.RegisterInputType(reflect.TypeOf((*H2qposuprovidernaiArrayInput)(nil)).Elem(), H2qposuprovidernaiArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*H2qposuprovidernaiMapInput)(nil)).Elem(), H2qposuprovidernaiMap{})
	pulumi.RegisterOutputType(H2qposuprovidernaiOutput{})
	pulumi.RegisterOutputType(H2qposuprovidernaiArrayOutput{})
	pulumi.RegisterOutputType(H2qposuprovidernaiMapOutput{})
}