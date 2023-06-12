// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package switchcontroller

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure port policy to be applied on the managed FortiSwitch ports through NAC device. Applies to FortiOS Version `6.4.0,6.4.1,6.4.2,6.4.10,7.0.0`.
//
// ## Import
//
// # SwitchController PortPolicy can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:switchcontroller/portpolicy:Portpolicy labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:switchcontroller/portpolicy:Portpolicy labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type Portpolicy struct {
	pulumi.CustomResourceState

	// Enable/disable bouncing (administratively bring the link down, up) of a switch port where this port policy is applied. Helps to clear and reassign VLAN from lldp-profile. Valid values: `disable`, `enable`.
	BouncePortLink pulumi.StringOutput `pulumi:"bouncePortLink"`
	// Description for the port policy.
	Description pulumi.StringOutput `pulumi:"description"`
	// FortiLink interface for which this port policy belongs to.
	Fortilink pulumi.StringOutput `pulumi:"fortilink"`
	// LLDP profile to be applied when using this port-policy.
	LldpProfile pulumi.StringOutput `pulumi:"lldpProfile"`
	// 802.1x security policy to be applied when using this port-policy.
	N8021x pulumi.StringOutput `pulumi:"n8021x"`
	// Port policy name.
	Name pulumi.StringOutput `pulumi:"name"`
	// QoS policy to be applied when using this port-policy.
	QosPolicy pulumi.StringOutput `pulumi:"qosPolicy"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// VLAN policy to be applied when using this port-policy.
	VlanPolicy pulumi.StringOutput `pulumi:"vlanPolicy"`
}

// NewPortpolicy registers a new resource with the given unique name, arguments, and options.
func NewPortpolicy(ctx *pulumi.Context,
	name string, args *PortpolicyArgs, opts ...pulumi.ResourceOption) (*Portpolicy, error) {
	if args == nil {
		args = &PortpolicyArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource Portpolicy
	err := ctx.RegisterResource("fortios:switchcontroller/portpolicy:Portpolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPortpolicy gets an existing Portpolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPortpolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PortpolicyState, opts ...pulumi.ResourceOption) (*Portpolicy, error) {
	var resource Portpolicy
	err := ctx.ReadResource("fortios:switchcontroller/portpolicy:Portpolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Portpolicy resources.
type portpolicyState struct {
	// Enable/disable bouncing (administratively bring the link down, up) of a switch port where this port policy is applied. Helps to clear and reassign VLAN from lldp-profile. Valid values: `disable`, `enable`.
	BouncePortLink *string `pulumi:"bouncePortLink"`
	// Description for the port policy.
	Description *string `pulumi:"description"`
	// FortiLink interface for which this port policy belongs to.
	Fortilink *string `pulumi:"fortilink"`
	// LLDP profile to be applied when using this port-policy.
	LldpProfile *string `pulumi:"lldpProfile"`
	// 802.1x security policy to be applied when using this port-policy.
	N8021x *string `pulumi:"n8021x"`
	// Port policy name.
	Name *string `pulumi:"name"`
	// QoS policy to be applied when using this port-policy.
	QosPolicy *string `pulumi:"qosPolicy"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// VLAN policy to be applied when using this port-policy.
	VlanPolicy *string `pulumi:"vlanPolicy"`
}

type PortpolicyState struct {
	// Enable/disable bouncing (administratively bring the link down, up) of a switch port where this port policy is applied. Helps to clear and reassign VLAN from lldp-profile. Valid values: `disable`, `enable`.
	BouncePortLink pulumi.StringPtrInput
	// Description for the port policy.
	Description pulumi.StringPtrInput
	// FortiLink interface for which this port policy belongs to.
	Fortilink pulumi.StringPtrInput
	// LLDP profile to be applied when using this port-policy.
	LldpProfile pulumi.StringPtrInput
	// 802.1x security policy to be applied when using this port-policy.
	N8021x pulumi.StringPtrInput
	// Port policy name.
	Name pulumi.StringPtrInput
	// QoS policy to be applied when using this port-policy.
	QosPolicy pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// VLAN policy to be applied when using this port-policy.
	VlanPolicy pulumi.StringPtrInput
}

func (PortpolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*portpolicyState)(nil)).Elem()
}

type portpolicyArgs struct {
	// Enable/disable bouncing (administratively bring the link down, up) of a switch port where this port policy is applied. Helps to clear and reassign VLAN from lldp-profile. Valid values: `disable`, `enable`.
	BouncePortLink *string `pulumi:"bouncePortLink"`
	// Description for the port policy.
	Description *string `pulumi:"description"`
	// FortiLink interface for which this port policy belongs to.
	Fortilink *string `pulumi:"fortilink"`
	// LLDP profile to be applied when using this port-policy.
	LldpProfile *string `pulumi:"lldpProfile"`
	// 802.1x security policy to be applied when using this port-policy.
	N8021x *string `pulumi:"n8021x"`
	// Port policy name.
	Name *string `pulumi:"name"`
	// QoS policy to be applied when using this port-policy.
	QosPolicy *string `pulumi:"qosPolicy"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// VLAN policy to be applied when using this port-policy.
	VlanPolicy *string `pulumi:"vlanPolicy"`
}

// The set of arguments for constructing a Portpolicy resource.
type PortpolicyArgs struct {
	// Enable/disable bouncing (administratively bring the link down, up) of a switch port where this port policy is applied. Helps to clear and reassign VLAN from lldp-profile. Valid values: `disable`, `enable`.
	BouncePortLink pulumi.StringPtrInput
	// Description for the port policy.
	Description pulumi.StringPtrInput
	// FortiLink interface for which this port policy belongs to.
	Fortilink pulumi.StringPtrInput
	// LLDP profile to be applied when using this port-policy.
	LldpProfile pulumi.StringPtrInput
	// 802.1x security policy to be applied when using this port-policy.
	N8021x pulumi.StringPtrInput
	// Port policy name.
	Name pulumi.StringPtrInput
	// QoS policy to be applied when using this port-policy.
	QosPolicy pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// VLAN policy to be applied when using this port-policy.
	VlanPolicy pulumi.StringPtrInput
}

func (PortpolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*portpolicyArgs)(nil)).Elem()
}

type PortpolicyInput interface {
	pulumi.Input

	ToPortpolicyOutput() PortpolicyOutput
	ToPortpolicyOutputWithContext(ctx context.Context) PortpolicyOutput
}

func (*Portpolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**Portpolicy)(nil)).Elem()
}

func (i *Portpolicy) ToPortpolicyOutput() PortpolicyOutput {
	return i.ToPortpolicyOutputWithContext(context.Background())
}

func (i *Portpolicy) ToPortpolicyOutputWithContext(ctx context.Context) PortpolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PortpolicyOutput)
}

// PortpolicyArrayInput is an input type that accepts PortpolicyArray and PortpolicyArrayOutput values.
// You can construct a concrete instance of `PortpolicyArrayInput` via:
//
//	PortpolicyArray{ PortpolicyArgs{...} }
type PortpolicyArrayInput interface {
	pulumi.Input

	ToPortpolicyArrayOutput() PortpolicyArrayOutput
	ToPortpolicyArrayOutputWithContext(context.Context) PortpolicyArrayOutput
}

type PortpolicyArray []PortpolicyInput

func (PortpolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Portpolicy)(nil)).Elem()
}

func (i PortpolicyArray) ToPortpolicyArrayOutput() PortpolicyArrayOutput {
	return i.ToPortpolicyArrayOutputWithContext(context.Background())
}

func (i PortpolicyArray) ToPortpolicyArrayOutputWithContext(ctx context.Context) PortpolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PortpolicyArrayOutput)
}

// PortpolicyMapInput is an input type that accepts PortpolicyMap and PortpolicyMapOutput values.
// You can construct a concrete instance of `PortpolicyMapInput` via:
//
//	PortpolicyMap{ "key": PortpolicyArgs{...} }
type PortpolicyMapInput interface {
	pulumi.Input

	ToPortpolicyMapOutput() PortpolicyMapOutput
	ToPortpolicyMapOutputWithContext(context.Context) PortpolicyMapOutput
}

type PortpolicyMap map[string]PortpolicyInput

func (PortpolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Portpolicy)(nil)).Elem()
}

func (i PortpolicyMap) ToPortpolicyMapOutput() PortpolicyMapOutput {
	return i.ToPortpolicyMapOutputWithContext(context.Background())
}

func (i PortpolicyMap) ToPortpolicyMapOutputWithContext(ctx context.Context) PortpolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PortpolicyMapOutput)
}

type PortpolicyOutput struct{ *pulumi.OutputState }

func (PortpolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Portpolicy)(nil)).Elem()
}

func (o PortpolicyOutput) ToPortpolicyOutput() PortpolicyOutput {
	return o
}

func (o PortpolicyOutput) ToPortpolicyOutputWithContext(ctx context.Context) PortpolicyOutput {
	return o
}

// Enable/disable bouncing (administratively bring the link down, up) of a switch port where this port policy is applied. Helps to clear and reassign VLAN from lldp-profile. Valid values: `disable`, `enable`.
func (o PortpolicyOutput) BouncePortLink() pulumi.StringOutput {
	return o.ApplyT(func(v *Portpolicy) pulumi.StringOutput { return v.BouncePortLink }).(pulumi.StringOutput)
}

// Description for the port policy.
func (o PortpolicyOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Portpolicy) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// FortiLink interface for which this port policy belongs to.
func (o PortpolicyOutput) Fortilink() pulumi.StringOutput {
	return o.ApplyT(func(v *Portpolicy) pulumi.StringOutput { return v.Fortilink }).(pulumi.StringOutput)
}

// LLDP profile to be applied when using this port-policy.
func (o PortpolicyOutput) LldpProfile() pulumi.StringOutput {
	return o.ApplyT(func(v *Portpolicy) pulumi.StringOutput { return v.LldpProfile }).(pulumi.StringOutput)
}

// 802.1x security policy to be applied when using this port-policy.
func (o PortpolicyOutput) N8021x() pulumi.StringOutput {
	return o.ApplyT(func(v *Portpolicy) pulumi.StringOutput { return v.N8021x }).(pulumi.StringOutput)
}

// Port policy name.
func (o PortpolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Portpolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// QoS policy to be applied when using this port-policy.
func (o PortpolicyOutput) QosPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v *Portpolicy) pulumi.StringOutput { return v.QosPolicy }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o PortpolicyOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Portpolicy) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// VLAN policy to be applied when using this port-policy.
func (o PortpolicyOutput) VlanPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v *Portpolicy) pulumi.StringOutput { return v.VlanPolicy }).(pulumi.StringOutput)
}

type PortpolicyArrayOutput struct{ *pulumi.OutputState }

func (PortpolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Portpolicy)(nil)).Elem()
}

func (o PortpolicyArrayOutput) ToPortpolicyArrayOutput() PortpolicyArrayOutput {
	return o
}

func (o PortpolicyArrayOutput) ToPortpolicyArrayOutputWithContext(ctx context.Context) PortpolicyArrayOutput {
	return o
}

func (o PortpolicyArrayOutput) Index(i pulumi.IntInput) PortpolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Portpolicy {
		return vs[0].([]*Portpolicy)[vs[1].(int)]
	}).(PortpolicyOutput)
}

type PortpolicyMapOutput struct{ *pulumi.OutputState }

func (PortpolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Portpolicy)(nil)).Elem()
}

func (o PortpolicyMapOutput) ToPortpolicyMapOutput() PortpolicyMapOutput {
	return o
}

func (o PortpolicyMapOutput) ToPortpolicyMapOutputWithContext(ctx context.Context) PortpolicyMapOutput {
	return o
}

func (o PortpolicyMapOutput) MapIndex(k pulumi.StringInput) PortpolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Portpolicy {
		return vs[0].(map[string]*Portpolicy)[vs[1].(string)]
	}).(PortpolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PortpolicyInput)(nil)).Elem(), &Portpolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*PortpolicyArrayInput)(nil)).Elem(), PortpolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PortpolicyMapInput)(nil)).Elem(), PortpolicyMap{})
	pulumi.RegisterOutputType(PortpolicyOutput{})
	pulumi.RegisterOutputType(PortpolicyArrayOutput{})
	pulumi.RegisterOutputType(PortpolicyMapOutput{})
}