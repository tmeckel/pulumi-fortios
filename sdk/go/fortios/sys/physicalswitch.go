// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sys

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure physical switches. Applies to FortiOS Version `7.0.4`.
//
// ## Import
//
// # System PhysicalSwitch can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:sys/physicalswitch:Physicalswitch labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:sys/physicalswitch:Physicalswitch labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type Physicalswitch struct {
	pulumi.CustomResourceState

	// Enable/disable layer 2 age timer. Valid values: `enable`, `disable`.
	AgeEnable pulumi.StringOutput `pulumi:"ageEnable"`
	// Layer 2 table age timer value.
	AgeVal pulumi.IntOutput `pulumi:"ageVal"`
	// Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewPhysicalswitch registers a new resource with the given unique name, arguments, and options.
func NewPhysicalswitch(ctx *pulumi.Context,
	name string, args *PhysicalswitchArgs, opts ...pulumi.ResourceOption) (*Physicalswitch, error) {
	if args == nil {
		args = &PhysicalswitchArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource Physicalswitch
	err := ctx.RegisterResource("fortios:sys/physicalswitch:Physicalswitch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPhysicalswitch gets an existing Physicalswitch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPhysicalswitch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PhysicalswitchState, opts ...pulumi.ResourceOption) (*Physicalswitch, error) {
	var resource Physicalswitch
	err := ctx.ReadResource("fortios:sys/physicalswitch:Physicalswitch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Physicalswitch resources.
type physicalswitchState struct {
	// Enable/disable layer 2 age timer. Valid values: `enable`, `disable`.
	AgeEnable *string `pulumi:"ageEnable"`
	// Layer 2 table age timer value.
	AgeVal *int `pulumi:"ageVal"`
	// Name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type PhysicalswitchState struct {
	// Enable/disable layer 2 age timer. Valid values: `enable`, `disable`.
	AgeEnable pulumi.StringPtrInput
	// Layer 2 table age timer value.
	AgeVal pulumi.IntPtrInput
	// Name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (PhysicalswitchState) ElementType() reflect.Type {
	return reflect.TypeOf((*physicalswitchState)(nil)).Elem()
}

type physicalswitchArgs struct {
	// Enable/disable layer 2 age timer. Valid values: `enable`, `disable`.
	AgeEnable *string `pulumi:"ageEnable"`
	// Layer 2 table age timer value.
	AgeVal *int `pulumi:"ageVal"`
	// Name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Physicalswitch resource.
type PhysicalswitchArgs struct {
	// Enable/disable layer 2 age timer. Valid values: `enable`, `disable`.
	AgeEnable pulumi.StringPtrInput
	// Layer 2 table age timer value.
	AgeVal pulumi.IntPtrInput
	// Name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (PhysicalswitchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*physicalswitchArgs)(nil)).Elem()
}

type PhysicalswitchInput interface {
	pulumi.Input

	ToPhysicalswitchOutput() PhysicalswitchOutput
	ToPhysicalswitchOutputWithContext(ctx context.Context) PhysicalswitchOutput
}

func (*Physicalswitch) ElementType() reflect.Type {
	return reflect.TypeOf((**Physicalswitch)(nil)).Elem()
}

func (i *Physicalswitch) ToPhysicalswitchOutput() PhysicalswitchOutput {
	return i.ToPhysicalswitchOutputWithContext(context.Background())
}

func (i *Physicalswitch) ToPhysicalswitchOutputWithContext(ctx context.Context) PhysicalswitchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PhysicalswitchOutput)
}

// PhysicalswitchArrayInput is an input type that accepts PhysicalswitchArray and PhysicalswitchArrayOutput values.
// You can construct a concrete instance of `PhysicalswitchArrayInput` via:
//
//	PhysicalswitchArray{ PhysicalswitchArgs{...} }
type PhysicalswitchArrayInput interface {
	pulumi.Input

	ToPhysicalswitchArrayOutput() PhysicalswitchArrayOutput
	ToPhysicalswitchArrayOutputWithContext(context.Context) PhysicalswitchArrayOutput
}

type PhysicalswitchArray []PhysicalswitchInput

func (PhysicalswitchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Physicalswitch)(nil)).Elem()
}

func (i PhysicalswitchArray) ToPhysicalswitchArrayOutput() PhysicalswitchArrayOutput {
	return i.ToPhysicalswitchArrayOutputWithContext(context.Background())
}

func (i PhysicalswitchArray) ToPhysicalswitchArrayOutputWithContext(ctx context.Context) PhysicalswitchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PhysicalswitchArrayOutput)
}

// PhysicalswitchMapInput is an input type that accepts PhysicalswitchMap and PhysicalswitchMapOutput values.
// You can construct a concrete instance of `PhysicalswitchMapInput` via:
//
//	PhysicalswitchMap{ "key": PhysicalswitchArgs{...} }
type PhysicalswitchMapInput interface {
	pulumi.Input

	ToPhysicalswitchMapOutput() PhysicalswitchMapOutput
	ToPhysicalswitchMapOutputWithContext(context.Context) PhysicalswitchMapOutput
}

type PhysicalswitchMap map[string]PhysicalswitchInput

func (PhysicalswitchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Physicalswitch)(nil)).Elem()
}

func (i PhysicalswitchMap) ToPhysicalswitchMapOutput() PhysicalswitchMapOutput {
	return i.ToPhysicalswitchMapOutputWithContext(context.Background())
}

func (i PhysicalswitchMap) ToPhysicalswitchMapOutputWithContext(ctx context.Context) PhysicalswitchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PhysicalswitchMapOutput)
}

type PhysicalswitchOutput struct{ *pulumi.OutputState }

func (PhysicalswitchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Physicalswitch)(nil)).Elem()
}

func (o PhysicalswitchOutput) ToPhysicalswitchOutput() PhysicalswitchOutput {
	return o
}

func (o PhysicalswitchOutput) ToPhysicalswitchOutputWithContext(ctx context.Context) PhysicalswitchOutput {
	return o
}

// Enable/disable layer 2 age timer. Valid values: `enable`, `disable`.
func (o PhysicalswitchOutput) AgeEnable() pulumi.StringOutput {
	return o.ApplyT(func(v *Physicalswitch) pulumi.StringOutput { return v.AgeEnable }).(pulumi.StringOutput)
}

// Layer 2 table age timer value.
func (o PhysicalswitchOutput) AgeVal() pulumi.IntOutput {
	return o.ApplyT(func(v *Physicalswitch) pulumi.IntOutput { return v.AgeVal }).(pulumi.IntOutput)
}

// Name.
func (o PhysicalswitchOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Physicalswitch) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o PhysicalswitchOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Physicalswitch) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type PhysicalswitchArrayOutput struct{ *pulumi.OutputState }

func (PhysicalswitchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Physicalswitch)(nil)).Elem()
}

func (o PhysicalswitchArrayOutput) ToPhysicalswitchArrayOutput() PhysicalswitchArrayOutput {
	return o
}

func (o PhysicalswitchArrayOutput) ToPhysicalswitchArrayOutputWithContext(ctx context.Context) PhysicalswitchArrayOutput {
	return o
}

func (o PhysicalswitchArrayOutput) Index(i pulumi.IntInput) PhysicalswitchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Physicalswitch {
		return vs[0].([]*Physicalswitch)[vs[1].(int)]
	}).(PhysicalswitchOutput)
}

type PhysicalswitchMapOutput struct{ *pulumi.OutputState }

func (PhysicalswitchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Physicalswitch)(nil)).Elem()
}

func (o PhysicalswitchMapOutput) ToPhysicalswitchMapOutput() PhysicalswitchMapOutput {
	return o
}

func (o PhysicalswitchMapOutput) ToPhysicalswitchMapOutputWithContext(ctx context.Context) PhysicalswitchMapOutput {
	return o
}

func (o PhysicalswitchMapOutput) MapIndex(k pulumi.StringInput) PhysicalswitchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Physicalswitch {
		return vs[0].(map[string]*Physicalswitch)[vs[1].(string)]
	}).(PhysicalswitchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PhysicalswitchInput)(nil)).Elem(), &Physicalswitch{})
	pulumi.RegisterInputType(reflect.TypeOf((*PhysicalswitchArrayInput)(nil)).Elem(), PhysicalswitchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PhysicalswitchMapInput)(nil)).Elem(), PhysicalswitchMap{})
	pulumi.RegisterOutputType(PhysicalswitchOutput{})
	pulumi.RegisterOutputType(PhysicalswitchArrayOutput{})
	pulumi.RegisterOutputType(PhysicalswitchMapOutput{})
}
