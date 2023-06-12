// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package user

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure device groups. Applies to FortiOS Version `<= 6.2.0`.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/user"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			trnames12, err := user.NewDevice(ctx, "trnames12", &user.DeviceArgs{
//				Alias:    pulumi.String("user_devices2"),
//				Category: pulumi.String("amazon-device"),
//				Mac:      pulumi.String("08:00:20:0a:1c:1d"),
//				Type:     pulumi.String("unknown"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = user.NewDevicegroup(ctx, "trname", &user.DevicegroupArgs{
//				Members: user.DevicegroupMemberArray{
//					&user.DevicegroupMemberArgs{
//						Name: trnames12.Alias,
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
// # User DeviceGroup can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:user/devicegroup:Devicegroup labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:user/devicegroup:Devicegroup labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type Devicegroup struct {
	pulumi.CustomResourceState

	// Comment.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Device group member. The structure of `member` block is documented below.
	Members DevicegroupMemberArrayOutput `pulumi:"members"`
	// Device group name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Config object tagging. The structure of `tagging` block is documented below.
	Taggings DevicegroupTaggingArrayOutput `pulumi:"taggings"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewDevicegroup registers a new resource with the given unique name, arguments, and options.
func NewDevicegroup(ctx *pulumi.Context,
	name string, args *DevicegroupArgs, opts ...pulumi.ResourceOption) (*Devicegroup, error) {
	if args == nil {
		args = &DevicegroupArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource Devicegroup
	err := ctx.RegisterResource("fortios:user/devicegroup:Devicegroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDevicegroup gets an existing Devicegroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDevicegroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DevicegroupState, opts ...pulumi.ResourceOption) (*Devicegroup, error) {
	var resource Devicegroup
	err := ctx.ReadResource("fortios:user/devicegroup:Devicegroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Devicegroup resources.
type devicegroupState struct {
	// Comment.
	Comment *string `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Device group member. The structure of `member` block is documented below.
	Members []DevicegroupMember `pulumi:"members"`
	// Device group name.
	Name *string `pulumi:"name"`
	// Config object tagging. The structure of `tagging` block is documented below.
	Taggings []DevicegroupTagging `pulumi:"taggings"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type DevicegroupState struct {
	// Comment.
	Comment pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Device group member. The structure of `member` block is documented below.
	Members DevicegroupMemberArrayInput
	// Device group name.
	Name pulumi.StringPtrInput
	// Config object tagging. The structure of `tagging` block is documented below.
	Taggings DevicegroupTaggingArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (DevicegroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*devicegroupState)(nil)).Elem()
}

type devicegroupArgs struct {
	// Comment.
	Comment *string `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Device group member. The structure of `member` block is documented below.
	Members []DevicegroupMember `pulumi:"members"`
	// Device group name.
	Name *string `pulumi:"name"`
	// Config object tagging. The structure of `tagging` block is documented below.
	Taggings []DevicegroupTagging `pulumi:"taggings"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Devicegroup resource.
type DevicegroupArgs struct {
	// Comment.
	Comment pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Device group member. The structure of `member` block is documented below.
	Members DevicegroupMemberArrayInput
	// Device group name.
	Name pulumi.StringPtrInput
	// Config object tagging. The structure of `tagging` block is documented below.
	Taggings DevicegroupTaggingArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (DevicegroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*devicegroupArgs)(nil)).Elem()
}

type DevicegroupInput interface {
	pulumi.Input

	ToDevicegroupOutput() DevicegroupOutput
	ToDevicegroupOutputWithContext(ctx context.Context) DevicegroupOutput
}

func (*Devicegroup) ElementType() reflect.Type {
	return reflect.TypeOf((**Devicegroup)(nil)).Elem()
}

func (i *Devicegroup) ToDevicegroupOutput() DevicegroupOutput {
	return i.ToDevicegroupOutputWithContext(context.Background())
}

func (i *Devicegroup) ToDevicegroupOutputWithContext(ctx context.Context) DevicegroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DevicegroupOutput)
}

// DevicegroupArrayInput is an input type that accepts DevicegroupArray and DevicegroupArrayOutput values.
// You can construct a concrete instance of `DevicegroupArrayInput` via:
//
//	DevicegroupArray{ DevicegroupArgs{...} }
type DevicegroupArrayInput interface {
	pulumi.Input

	ToDevicegroupArrayOutput() DevicegroupArrayOutput
	ToDevicegroupArrayOutputWithContext(context.Context) DevicegroupArrayOutput
}

type DevicegroupArray []DevicegroupInput

func (DevicegroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Devicegroup)(nil)).Elem()
}

func (i DevicegroupArray) ToDevicegroupArrayOutput() DevicegroupArrayOutput {
	return i.ToDevicegroupArrayOutputWithContext(context.Background())
}

func (i DevicegroupArray) ToDevicegroupArrayOutputWithContext(ctx context.Context) DevicegroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DevicegroupArrayOutput)
}

// DevicegroupMapInput is an input type that accepts DevicegroupMap and DevicegroupMapOutput values.
// You can construct a concrete instance of `DevicegroupMapInput` via:
//
//	DevicegroupMap{ "key": DevicegroupArgs{...} }
type DevicegroupMapInput interface {
	pulumi.Input

	ToDevicegroupMapOutput() DevicegroupMapOutput
	ToDevicegroupMapOutputWithContext(context.Context) DevicegroupMapOutput
}

type DevicegroupMap map[string]DevicegroupInput

func (DevicegroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Devicegroup)(nil)).Elem()
}

func (i DevicegroupMap) ToDevicegroupMapOutput() DevicegroupMapOutput {
	return i.ToDevicegroupMapOutputWithContext(context.Background())
}

func (i DevicegroupMap) ToDevicegroupMapOutputWithContext(ctx context.Context) DevicegroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DevicegroupMapOutput)
}

type DevicegroupOutput struct{ *pulumi.OutputState }

func (DevicegroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Devicegroup)(nil)).Elem()
}

func (o DevicegroupOutput) ToDevicegroupOutput() DevicegroupOutput {
	return o
}

func (o DevicegroupOutput) ToDevicegroupOutputWithContext(ctx context.Context) DevicegroupOutput {
	return o
}

// Comment.
func (o DevicegroupOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Devicegroup) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o DevicegroupOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Devicegroup) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Device group member. The structure of `member` block is documented below.
func (o DevicegroupOutput) Members() DevicegroupMemberArrayOutput {
	return o.ApplyT(func(v *Devicegroup) DevicegroupMemberArrayOutput { return v.Members }).(DevicegroupMemberArrayOutput)
}

// Device group name.
func (o DevicegroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Devicegroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Config object tagging. The structure of `tagging` block is documented below.
func (o DevicegroupOutput) Taggings() DevicegroupTaggingArrayOutput {
	return o.ApplyT(func(v *Devicegroup) DevicegroupTaggingArrayOutput { return v.Taggings }).(DevicegroupTaggingArrayOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o DevicegroupOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Devicegroup) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type DevicegroupArrayOutput struct{ *pulumi.OutputState }

func (DevicegroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Devicegroup)(nil)).Elem()
}

func (o DevicegroupArrayOutput) ToDevicegroupArrayOutput() DevicegroupArrayOutput {
	return o
}

func (o DevicegroupArrayOutput) ToDevicegroupArrayOutputWithContext(ctx context.Context) DevicegroupArrayOutput {
	return o
}

func (o DevicegroupArrayOutput) Index(i pulumi.IntInput) DevicegroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Devicegroup {
		return vs[0].([]*Devicegroup)[vs[1].(int)]
	}).(DevicegroupOutput)
}

type DevicegroupMapOutput struct{ *pulumi.OutputState }

func (DevicegroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Devicegroup)(nil)).Elem()
}

func (o DevicegroupMapOutput) ToDevicegroupMapOutput() DevicegroupMapOutput {
	return o
}

func (o DevicegroupMapOutput) ToDevicegroupMapOutputWithContext(ctx context.Context) DevicegroupMapOutput {
	return o
}

func (o DevicegroupMapOutput) MapIndex(k pulumi.StringInput) DevicegroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Devicegroup {
		return vs[0].(map[string]*Devicegroup)[vs[1].(string)]
	}).(DevicegroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DevicegroupInput)(nil)).Elem(), &Devicegroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*DevicegroupArrayInput)(nil)).Elem(), DevicegroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DevicegroupMapInput)(nil)).Elem(), DevicegroupMap{})
	pulumi.RegisterOutputType(DevicegroupOutput{})
	pulumi.RegisterOutputType(DevicegroupArrayOutput{})
	pulumi.RegisterOutputType(DevicegroupMapOutput{})
}