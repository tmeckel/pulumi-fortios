// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package user

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure device categories. Applies to FortiOS Version `<= 6.2.0`.
//
// ## Import
//
// # User DeviceCategory can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:user/devicecategory:Devicecategory labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:user/devicecategory:Devicecategory labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type Devicecategory struct {
	pulumi.CustomResourceState

	// Comment.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Device category description.
	Desc pulumi.StringPtrOutput `pulumi:"desc"`
	// Device category name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewDevicecategory registers a new resource with the given unique name, arguments, and options.
func NewDevicecategory(ctx *pulumi.Context,
	name string, args *DevicecategoryArgs, opts ...pulumi.ResourceOption) (*Devicecategory, error) {
	if args == nil {
		args = &DevicecategoryArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource Devicecategory
	err := ctx.RegisterResource("fortios:user/devicecategory:Devicecategory", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDevicecategory gets an existing Devicecategory resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDevicecategory(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DevicecategoryState, opts ...pulumi.ResourceOption) (*Devicecategory, error) {
	var resource Devicecategory
	err := ctx.ReadResource("fortios:user/devicecategory:Devicecategory", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Devicecategory resources.
type devicecategoryState struct {
	// Comment.
	Comment *string `pulumi:"comment"`
	// Device category description.
	Desc *string `pulumi:"desc"`
	// Device category name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type DevicecategoryState struct {
	// Comment.
	Comment pulumi.StringPtrInput
	// Device category description.
	Desc pulumi.StringPtrInput
	// Device category name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (DevicecategoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*devicecategoryState)(nil)).Elem()
}

type devicecategoryArgs struct {
	// Comment.
	Comment *string `pulumi:"comment"`
	// Device category description.
	Desc *string `pulumi:"desc"`
	// Device category name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Devicecategory resource.
type DevicecategoryArgs struct {
	// Comment.
	Comment pulumi.StringPtrInput
	// Device category description.
	Desc pulumi.StringPtrInput
	// Device category name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (DevicecategoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*devicecategoryArgs)(nil)).Elem()
}

type DevicecategoryInput interface {
	pulumi.Input

	ToDevicecategoryOutput() DevicecategoryOutput
	ToDevicecategoryOutputWithContext(ctx context.Context) DevicecategoryOutput
}

func (*Devicecategory) ElementType() reflect.Type {
	return reflect.TypeOf((**Devicecategory)(nil)).Elem()
}

func (i *Devicecategory) ToDevicecategoryOutput() DevicecategoryOutput {
	return i.ToDevicecategoryOutputWithContext(context.Background())
}

func (i *Devicecategory) ToDevicecategoryOutputWithContext(ctx context.Context) DevicecategoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DevicecategoryOutput)
}

// DevicecategoryArrayInput is an input type that accepts DevicecategoryArray and DevicecategoryArrayOutput values.
// You can construct a concrete instance of `DevicecategoryArrayInput` via:
//
//	DevicecategoryArray{ DevicecategoryArgs{...} }
type DevicecategoryArrayInput interface {
	pulumi.Input

	ToDevicecategoryArrayOutput() DevicecategoryArrayOutput
	ToDevicecategoryArrayOutputWithContext(context.Context) DevicecategoryArrayOutput
}

type DevicecategoryArray []DevicecategoryInput

func (DevicecategoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Devicecategory)(nil)).Elem()
}

func (i DevicecategoryArray) ToDevicecategoryArrayOutput() DevicecategoryArrayOutput {
	return i.ToDevicecategoryArrayOutputWithContext(context.Background())
}

func (i DevicecategoryArray) ToDevicecategoryArrayOutputWithContext(ctx context.Context) DevicecategoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DevicecategoryArrayOutput)
}

// DevicecategoryMapInput is an input type that accepts DevicecategoryMap and DevicecategoryMapOutput values.
// You can construct a concrete instance of `DevicecategoryMapInput` via:
//
//	DevicecategoryMap{ "key": DevicecategoryArgs{...} }
type DevicecategoryMapInput interface {
	pulumi.Input

	ToDevicecategoryMapOutput() DevicecategoryMapOutput
	ToDevicecategoryMapOutputWithContext(context.Context) DevicecategoryMapOutput
}

type DevicecategoryMap map[string]DevicecategoryInput

func (DevicecategoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Devicecategory)(nil)).Elem()
}

func (i DevicecategoryMap) ToDevicecategoryMapOutput() DevicecategoryMapOutput {
	return i.ToDevicecategoryMapOutputWithContext(context.Background())
}

func (i DevicecategoryMap) ToDevicecategoryMapOutputWithContext(ctx context.Context) DevicecategoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DevicecategoryMapOutput)
}

type DevicecategoryOutput struct{ *pulumi.OutputState }

func (DevicecategoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Devicecategory)(nil)).Elem()
}

func (o DevicecategoryOutput) ToDevicecategoryOutput() DevicecategoryOutput {
	return o
}

func (o DevicecategoryOutput) ToDevicecategoryOutputWithContext(ctx context.Context) DevicecategoryOutput {
	return o
}

// Comment.
func (o DevicecategoryOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Devicecategory) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// Device category description.
func (o DevicecategoryOutput) Desc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Devicecategory) pulumi.StringPtrOutput { return v.Desc }).(pulumi.StringPtrOutput)
}

// Device category name.
func (o DevicecategoryOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Devicecategory) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o DevicecategoryOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Devicecategory) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type DevicecategoryArrayOutput struct{ *pulumi.OutputState }

func (DevicecategoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Devicecategory)(nil)).Elem()
}

func (o DevicecategoryArrayOutput) ToDevicecategoryArrayOutput() DevicecategoryArrayOutput {
	return o
}

func (o DevicecategoryArrayOutput) ToDevicecategoryArrayOutputWithContext(ctx context.Context) DevicecategoryArrayOutput {
	return o
}

func (o DevicecategoryArrayOutput) Index(i pulumi.IntInput) DevicecategoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Devicecategory {
		return vs[0].([]*Devicecategory)[vs[1].(int)]
	}).(DevicecategoryOutput)
}

type DevicecategoryMapOutput struct{ *pulumi.OutputState }

func (DevicecategoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Devicecategory)(nil)).Elem()
}

func (o DevicecategoryMapOutput) ToDevicecategoryMapOutput() DevicecategoryMapOutput {
	return o
}

func (o DevicecategoryMapOutput) ToDevicecategoryMapOutputWithContext(ctx context.Context) DevicecategoryMapOutput {
	return o
}

func (o DevicecategoryMapOutput) MapIndex(k pulumi.StringInput) DevicecategoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Devicecategory {
		return vs[0].(map[string]*Devicecategory)[vs[1].(string)]
	}).(DevicecategoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DevicecategoryInput)(nil)).Elem(), &Devicecategory{})
	pulumi.RegisterInputType(reflect.TypeOf((*DevicecategoryArrayInput)(nil)).Elem(), DevicecategoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DevicecategoryMapInput)(nil)).Elem(), DevicecategoryMap{})
	pulumi.RegisterOutputType(DevicecategoryOutput{})
	pulumi.RegisterOutputType(DevicecategoryArrayOutput{})
	pulumi.RegisterOutputType(DevicecategoryMapOutput{})
}