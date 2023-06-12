// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package wirelesscontrollerhotspot20

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure venue name duple.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/wirelesscontrollerhotspot20"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := wirelesscontrollerhotspot20.NewAnqpvenuename(ctx, "trname", &wirelesscontrollerhotspot20.AnqpvenuenameArgs{
//				ValueLists: wirelesscontrollerhotspot20.AnqpvenuenameValueListArray{
//					&wirelesscontrollerhotspot20.AnqpvenuenameValueListArgs{
//						Index: pulumi.Int(1),
//						Lang:  pulumi.String("CN"),
//						Value: pulumi.String("3"),
//					},
//				},
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
// # WirelessControllerHotspot20 AnqpVenueName can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:wirelesscontrollerhotspot20/anqpvenuename:Anqpvenuename labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:wirelesscontrollerhotspot20/anqpvenuename:Anqpvenuename labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type Anqpvenuename struct {
	pulumi.CustomResourceState

	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Name of venue name duple.
	Name pulumi.StringOutput `pulumi:"name"`
	// Name list. The structure of `valueList` block is documented below.
	ValueLists AnqpvenuenameValueListArrayOutput `pulumi:"valueLists"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewAnqpvenuename registers a new resource with the given unique name, arguments, and options.
func NewAnqpvenuename(ctx *pulumi.Context,
	name string, args *AnqpvenuenameArgs, opts ...pulumi.ResourceOption) (*Anqpvenuename, error) {
	if args == nil {
		args = &AnqpvenuenameArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource Anqpvenuename
	err := ctx.RegisterResource("fortios:wirelesscontrollerhotspot20/anqpvenuename:Anqpvenuename", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAnqpvenuename gets an existing Anqpvenuename resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAnqpvenuename(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AnqpvenuenameState, opts ...pulumi.ResourceOption) (*Anqpvenuename, error) {
	var resource Anqpvenuename
	err := ctx.ReadResource("fortios:wirelesscontrollerhotspot20/anqpvenuename:Anqpvenuename", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Anqpvenuename resources.
type anqpvenuenameState struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Name of venue name duple.
	Name *string `pulumi:"name"`
	// Name list. The structure of `valueList` block is documented below.
	ValueLists []AnqpvenuenameValueList `pulumi:"valueLists"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type AnqpvenuenameState struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Name of venue name duple.
	Name pulumi.StringPtrInput
	// Name list. The structure of `valueList` block is documented below.
	ValueLists AnqpvenuenameValueListArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (AnqpvenuenameState) ElementType() reflect.Type {
	return reflect.TypeOf((*anqpvenuenameState)(nil)).Elem()
}

type anqpvenuenameArgs struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Name of venue name duple.
	Name *string `pulumi:"name"`
	// Name list. The structure of `valueList` block is documented below.
	ValueLists []AnqpvenuenameValueList `pulumi:"valueLists"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Anqpvenuename resource.
type AnqpvenuenameArgs struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Name of venue name duple.
	Name pulumi.StringPtrInput
	// Name list. The structure of `valueList` block is documented below.
	ValueLists AnqpvenuenameValueListArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (AnqpvenuenameArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*anqpvenuenameArgs)(nil)).Elem()
}

type AnqpvenuenameInput interface {
	pulumi.Input

	ToAnqpvenuenameOutput() AnqpvenuenameOutput
	ToAnqpvenuenameOutputWithContext(ctx context.Context) AnqpvenuenameOutput
}

func (*Anqpvenuename) ElementType() reflect.Type {
	return reflect.TypeOf((**Anqpvenuename)(nil)).Elem()
}

func (i *Anqpvenuename) ToAnqpvenuenameOutput() AnqpvenuenameOutput {
	return i.ToAnqpvenuenameOutputWithContext(context.Background())
}

func (i *Anqpvenuename) ToAnqpvenuenameOutputWithContext(ctx context.Context) AnqpvenuenameOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnqpvenuenameOutput)
}

// AnqpvenuenameArrayInput is an input type that accepts AnqpvenuenameArray and AnqpvenuenameArrayOutput values.
// You can construct a concrete instance of `AnqpvenuenameArrayInput` via:
//
//	AnqpvenuenameArray{ AnqpvenuenameArgs{...} }
type AnqpvenuenameArrayInput interface {
	pulumi.Input

	ToAnqpvenuenameArrayOutput() AnqpvenuenameArrayOutput
	ToAnqpvenuenameArrayOutputWithContext(context.Context) AnqpvenuenameArrayOutput
}

type AnqpvenuenameArray []AnqpvenuenameInput

func (AnqpvenuenameArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Anqpvenuename)(nil)).Elem()
}

func (i AnqpvenuenameArray) ToAnqpvenuenameArrayOutput() AnqpvenuenameArrayOutput {
	return i.ToAnqpvenuenameArrayOutputWithContext(context.Background())
}

func (i AnqpvenuenameArray) ToAnqpvenuenameArrayOutputWithContext(ctx context.Context) AnqpvenuenameArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnqpvenuenameArrayOutput)
}

// AnqpvenuenameMapInput is an input type that accepts AnqpvenuenameMap and AnqpvenuenameMapOutput values.
// You can construct a concrete instance of `AnqpvenuenameMapInput` via:
//
//	AnqpvenuenameMap{ "key": AnqpvenuenameArgs{...} }
type AnqpvenuenameMapInput interface {
	pulumi.Input

	ToAnqpvenuenameMapOutput() AnqpvenuenameMapOutput
	ToAnqpvenuenameMapOutputWithContext(context.Context) AnqpvenuenameMapOutput
}

type AnqpvenuenameMap map[string]AnqpvenuenameInput

func (AnqpvenuenameMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Anqpvenuename)(nil)).Elem()
}

func (i AnqpvenuenameMap) ToAnqpvenuenameMapOutput() AnqpvenuenameMapOutput {
	return i.ToAnqpvenuenameMapOutputWithContext(context.Background())
}

func (i AnqpvenuenameMap) ToAnqpvenuenameMapOutputWithContext(ctx context.Context) AnqpvenuenameMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnqpvenuenameMapOutput)
}

type AnqpvenuenameOutput struct{ *pulumi.OutputState }

func (AnqpvenuenameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Anqpvenuename)(nil)).Elem()
}

func (o AnqpvenuenameOutput) ToAnqpvenuenameOutput() AnqpvenuenameOutput {
	return o
}

func (o AnqpvenuenameOutput) ToAnqpvenuenameOutputWithContext(ctx context.Context) AnqpvenuenameOutput {
	return o
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o AnqpvenuenameOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Anqpvenuename) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Name of venue name duple.
func (o AnqpvenuenameOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Anqpvenuename) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Name list. The structure of `valueList` block is documented below.
func (o AnqpvenuenameOutput) ValueLists() AnqpvenuenameValueListArrayOutput {
	return o.ApplyT(func(v *Anqpvenuename) AnqpvenuenameValueListArrayOutput { return v.ValueLists }).(AnqpvenuenameValueListArrayOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o AnqpvenuenameOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Anqpvenuename) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type AnqpvenuenameArrayOutput struct{ *pulumi.OutputState }

func (AnqpvenuenameArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Anqpvenuename)(nil)).Elem()
}

func (o AnqpvenuenameArrayOutput) ToAnqpvenuenameArrayOutput() AnqpvenuenameArrayOutput {
	return o
}

func (o AnqpvenuenameArrayOutput) ToAnqpvenuenameArrayOutputWithContext(ctx context.Context) AnqpvenuenameArrayOutput {
	return o
}

func (o AnqpvenuenameArrayOutput) Index(i pulumi.IntInput) AnqpvenuenameOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Anqpvenuename {
		return vs[0].([]*Anqpvenuename)[vs[1].(int)]
	}).(AnqpvenuenameOutput)
}

type AnqpvenuenameMapOutput struct{ *pulumi.OutputState }

func (AnqpvenuenameMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Anqpvenuename)(nil)).Elem()
}

func (o AnqpvenuenameMapOutput) ToAnqpvenuenameMapOutput() AnqpvenuenameMapOutput {
	return o
}

func (o AnqpvenuenameMapOutput) ToAnqpvenuenameMapOutputWithContext(ctx context.Context) AnqpvenuenameMapOutput {
	return o
}

func (o AnqpvenuenameMapOutput) MapIndex(k pulumi.StringInput) AnqpvenuenameOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Anqpvenuename {
		return vs[0].(map[string]*Anqpvenuename)[vs[1].(string)]
	}).(AnqpvenuenameOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AnqpvenuenameInput)(nil)).Elem(), &Anqpvenuename{})
	pulumi.RegisterInputType(reflect.TypeOf((*AnqpvenuenameArrayInput)(nil)).Elem(), AnqpvenuenameArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AnqpvenuenameMapInput)(nil)).Elem(), AnqpvenuenameMap{})
	pulumi.RegisterOutputType(AnqpvenuenameOutput{})
	pulumi.RegisterOutputType(AnqpvenuenameArrayOutput{})
	pulumi.RegisterOutputType(AnqpvenuenameMapOutput{})
}
