// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package routerospf

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// OSPF neighbor configuration are used when OSPF runs on non-broadcast media
//
// > The provider supports the definition of Neighbor in Router Ospf `router.Ospf`, and also allows the definition of separate Neighbor resources `routerospf.Neighbor`, but do not use a `router.Ospf` with in-line Neighbor in conjunction with any `routerospf.Neighbor` resources, otherwise conflicts and overwrite will occur.
//
// ## Import
//
// # Routerospf Neighbor can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:routerospf/neighbor:Neighbor labelname {{fosid}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:routerospf/neighbor:Neighbor labelname {{fosid}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type Neighbor struct {
	pulumi.CustomResourceState

	// Cost of the interface, value range from 0 to 65535, 0 means auto-cost.
	Cost pulumi.IntOutput `pulumi:"cost"`
	// Neighbor entry ID.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// Interface IP address of the neighbor.
	Ip pulumi.StringOutput `pulumi:"ip"`
	// Poll interval time in seconds.
	PollInterval pulumi.IntOutput `pulumi:"pollInterval"`
	// Priority.
	Priority pulumi.IntOutput `pulumi:"priority"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewNeighbor registers a new resource with the given unique name, arguments, and options.
func NewNeighbor(ctx *pulumi.Context,
	name string, args *NeighborArgs, opts ...pulumi.ResourceOption) (*Neighbor, error) {
	if args == nil {
		args = &NeighborArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource Neighbor
	err := ctx.RegisterResource("fortios:routerospf/neighbor:Neighbor", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNeighbor gets an existing Neighbor resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNeighbor(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NeighborState, opts ...pulumi.ResourceOption) (*Neighbor, error) {
	var resource Neighbor
	err := ctx.ReadResource("fortios:routerospf/neighbor:Neighbor", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Neighbor resources.
type neighborState struct {
	// Cost of the interface, value range from 0 to 65535, 0 means auto-cost.
	Cost *int `pulumi:"cost"`
	// Neighbor entry ID.
	Fosid *int `pulumi:"fosid"`
	// Interface IP address of the neighbor.
	Ip *string `pulumi:"ip"`
	// Poll interval time in seconds.
	PollInterval *int `pulumi:"pollInterval"`
	// Priority.
	Priority *int `pulumi:"priority"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type NeighborState struct {
	// Cost of the interface, value range from 0 to 65535, 0 means auto-cost.
	Cost pulumi.IntPtrInput
	// Neighbor entry ID.
	Fosid pulumi.IntPtrInput
	// Interface IP address of the neighbor.
	Ip pulumi.StringPtrInput
	// Poll interval time in seconds.
	PollInterval pulumi.IntPtrInput
	// Priority.
	Priority pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (NeighborState) ElementType() reflect.Type {
	return reflect.TypeOf((*neighborState)(nil)).Elem()
}

type neighborArgs struct {
	// Cost of the interface, value range from 0 to 65535, 0 means auto-cost.
	Cost *int `pulumi:"cost"`
	// Neighbor entry ID.
	Fosid *int `pulumi:"fosid"`
	// Interface IP address of the neighbor.
	Ip *string `pulumi:"ip"`
	// Poll interval time in seconds.
	PollInterval *int `pulumi:"pollInterval"`
	// Priority.
	Priority *int `pulumi:"priority"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Neighbor resource.
type NeighborArgs struct {
	// Cost of the interface, value range from 0 to 65535, 0 means auto-cost.
	Cost pulumi.IntPtrInput
	// Neighbor entry ID.
	Fosid pulumi.IntPtrInput
	// Interface IP address of the neighbor.
	Ip pulumi.StringPtrInput
	// Poll interval time in seconds.
	PollInterval pulumi.IntPtrInput
	// Priority.
	Priority pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (NeighborArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*neighborArgs)(nil)).Elem()
}

type NeighborInput interface {
	pulumi.Input

	ToNeighborOutput() NeighborOutput
	ToNeighborOutputWithContext(ctx context.Context) NeighborOutput
}

func (*Neighbor) ElementType() reflect.Type {
	return reflect.TypeOf((**Neighbor)(nil)).Elem()
}

func (i *Neighbor) ToNeighborOutput() NeighborOutput {
	return i.ToNeighborOutputWithContext(context.Background())
}

func (i *Neighbor) ToNeighborOutputWithContext(ctx context.Context) NeighborOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NeighborOutput)
}

// NeighborArrayInput is an input type that accepts NeighborArray and NeighborArrayOutput values.
// You can construct a concrete instance of `NeighborArrayInput` via:
//
//	NeighborArray{ NeighborArgs{...} }
type NeighborArrayInput interface {
	pulumi.Input

	ToNeighborArrayOutput() NeighborArrayOutput
	ToNeighborArrayOutputWithContext(context.Context) NeighborArrayOutput
}

type NeighborArray []NeighborInput

func (NeighborArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Neighbor)(nil)).Elem()
}

func (i NeighborArray) ToNeighborArrayOutput() NeighborArrayOutput {
	return i.ToNeighborArrayOutputWithContext(context.Background())
}

func (i NeighborArray) ToNeighborArrayOutputWithContext(ctx context.Context) NeighborArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NeighborArrayOutput)
}

// NeighborMapInput is an input type that accepts NeighborMap and NeighborMapOutput values.
// You can construct a concrete instance of `NeighborMapInput` via:
//
//	NeighborMap{ "key": NeighborArgs{...} }
type NeighborMapInput interface {
	pulumi.Input

	ToNeighborMapOutput() NeighborMapOutput
	ToNeighborMapOutputWithContext(context.Context) NeighborMapOutput
}

type NeighborMap map[string]NeighborInput

func (NeighborMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Neighbor)(nil)).Elem()
}

func (i NeighborMap) ToNeighborMapOutput() NeighborMapOutput {
	return i.ToNeighborMapOutputWithContext(context.Background())
}

func (i NeighborMap) ToNeighborMapOutputWithContext(ctx context.Context) NeighborMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NeighborMapOutput)
}

type NeighborOutput struct{ *pulumi.OutputState }

func (NeighborOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Neighbor)(nil)).Elem()
}

func (o NeighborOutput) ToNeighborOutput() NeighborOutput {
	return o
}

func (o NeighborOutput) ToNeighborOutputWithContext(ctx context.Context) NeighborOutput {
	return o
}

// Cost of the interface, value range from 0 to 65535, 0 means auto-cost.
func (o NeighborOutput) Cost() pulumi.IntOutput {
	return o.ApplyT(func(v *Neighbor) pulumi.IntOutput { return v.Cost }).(pulumi.IntOutput)
}

// Neighbor entry ID.
func (o NeighborOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *Neighbor) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

// Interface IP address of the neighbor.
func (o NeighborOutput) Ip() pulumi.StringOutput {
	return o.ApplyT(func(v *Neighbor) pulumi.StringOutput { return v.Ip }).(pulumi.StringOutput)
}

// Poll interval time in seconds.
func (o NeighborOutput) PollInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *Neighbor) pulumi.IntOutput { return v.PollInterval }).(pulumi.IntOutput)
}

// Priority.
func (o NeighborOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v *Neighbor) pulumi.IntOutput { return v.Priority }).(pulumi.IntOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o NeighborOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Neighbor) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type NeighborArrayOutput struct{ *pulumi.OutputState }

func (NeighborArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Neighbor)(nil)).Elem()
}

func (o NeighborArrayOutput) ToNeighborArrayOutput() NeighborArrayOutput {
	return o
}

func (o NeighborArrayOutput) ToNeighborArrayOutputWithContext(ctx context.Context) NeighborArrayOutput {
	return o
}

func (o NeighborArrayOutput) Index(i pulumi.IntInput) NeighborOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Neighbor {
		return vs[0].([]*Neighbor)[vs[1].(int)]
	}).(NeighborOutput)
}

type NeighborMapOutput struct{ *pulumi.OutputState }

func (NeighborMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Neighbor)(nil)).Elem()
}

func (o NeighborMapOutput) ToNeighborMapOutput() NeighborMapOutput {
	return o
}

func (o NeighborMapOutput) ToNeighborMapOutputWithContext(ctx context.Context) NeighborMapOutput {
	return o
}

func (o NeighborMapOutput) MapIndex(k pulumi.StringInput) NeighborOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Neighbor {
		return vs[0].(map[string]*Neighbor)[vs[1].(string)]
	}).(NeighborOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NeighborInput)(nil)).Elem(), &Neighbor{})
	pulumi.RegisterInputType(reflect.TypeOf((*NeighborArrayInput)(nil)).Elem(), NeighborArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NeighborMapInput)(nil)).Elem(), NeighborMap{})
	pulumi.RegisterOutputType(NeighborOutput{})
	pulumi.RegisterOutputType(NeighborArrayOutput{})
	pulumi.RegisterOutputType(NeighborMapOutput{})
}
