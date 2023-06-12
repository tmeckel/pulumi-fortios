// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package switchcontrollerautoconfig

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure FortiSwitch Auto-Config QoS policy.
//
// ## Import
//
// # SwitchControllerAutoConfig Policy can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:switchcontrollerautoconfig/policy:Policy labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:switchcontrollerautoconfig/policy:Policy labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type Policy struct {
	pulumi.CustomResourceState

	// Enable/disable IGMP flood report. Valid values: `enable`, `disable`.
	IgmpFloodReport pulumi.StringOutput `pulumi:"igmpFloodReport"`
	// Enable/disable IGMP flood traffic. Valid values: `enable`, `disable`.
	IgmpFloodTraffic pulumi.StringOutput `pulumi:"igmpFloodTraffic"`
	// Auto-Config QoS policy name
	Name pulumi.StringOutput `pulumi:"name"`
	// Enable/disable PoE status. Valid values: `enable`, `disable`.
	PoeStatus pulumi.StringOutput `pulumi:"poeStatus"`
	// Auto-Config QoS policy.
	QosPolicy pulumi.StringOutput `pulumi:"qosPolicy"`
	// Auto-Config storm control policy.
	StormControlPolicy pulumi.StringOutput `pulumi:"stormControlPolicy"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewPolicy registers a new resource with the given unique name, arguments, and options.
func NewPolicy(ctx *pulumi.Context,
	name string, args *PolicyArgs, opts ...pulumi.ResourceOption) (*Policy, error) {
	if args == nil {
		args = &PolicyArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource Policy
	err := ctx.RegisterResource("fortios:switchcontrollerautoconfig/policy:Policy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicy gets an existing Policy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyState, opts ...pulumi.ResourceOption) (*Policy, error) {
	var resource Policy
	err := ctx.ReadResource("fortios:switchcontrollerautoconfig/policy:Policy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Policy resources.
type policyState struct {
	// Enable/disable IGMP flood report. Valid values: `enable`, `disable`.
	IgmpFloodReport *string `pulumi:"igmpFloodReport"`
	// Enable/disable IGMP flood traffic. Valid values: `enable`, `disable`.
	IgmpFloodTraffic *string `pulumi:"igmpFloodTraffic"`
	// Auto-Config QoS policy name
	Name *string `pulumi:"name"`
	// Enable/disable PoE status. Valid values: `enable`, `disable`.
	PoeStatus *string `pulumi:"poeStatus"`
	// Auto-Config QoS policy.
	QosPolicy *string `pulumi:"qosPolicy"`
	// Auto-Config storm control policy.
	StormControlPolicy *string `pulumi:"stormControlPolicy"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type PolicyState struct {
	// Enable/disable IGMP flood report. Valid values: `enable`, `disable`.
	IgmpFloodReport pulumi.StringPtrInput
	// Enable/disable IGMP flood traffic. Valid values: `enable`, `disable`.
	IgmpFloodTraffic pulumi.StringPtrInput
	// Auto-Config QoS policy name
	Name pulumi.StringPtrInput
	// Enable/disable PoE status. Valid values: `enable`, `disable`.
	PoeStatus pulumi.StringPtrInput
	// Auto-Config QoS policy.
	QosPolicy pulumi.StringPtrInput
	// Auto-Config storm control policy.
	StormControlPolicy pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (PolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyState)(nil)).Elem()
}

type policyArgs struct {
	// Enable/disable IGMP flood report. Valid values: `enable`, `disable`.
	IgmpFloodReport *string `pulumi:"igmpFloodReport"`
	// Enable/disable IGMP flood traffic. Valid values: `enable`, `disable`.
	IgmpFloodTraffic *string `pulumi:"igmpFloodTraffic"`
	// Auto-Config QoS policy name
	Name *string `pulumi:"name"`
	// Enable/disable PoE status. Valid values: `enable`, `disable`.
	PoeStatus *string `pulumi:"poeStatus"`
	// Auto-Config QoS policy.
	QosPolicy *string `pulumi:"qosPolicy"`
	// Auto-Config storm control policy.
	StormControlPolicy *string `pulumi:"stormControlPolicy"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Policy resource.
type PolicyArgs struct {
	// Enable/disable IGMP flood report. Valid values: `enable`, `disable`.
	IgmpFloodReport pulumi.StringPtrInput
	// Enable/disable IGMP flood traffic. Valid values: `enable`, `disable`.
	IgmpFloodTraffic pulumi.StringPtrInput
	// Auto-Config QoS policy name
	Name pulumi.StringPtrInput
	// Enable/disable PoE status. Valid values: `enable`, `disable`.
	PoeStatus pulumi.StringPtrInput
	// Auto-Config QoS policy.
	QosPolicy pulumi.StringPtrInput
	// Auto-Config storm control policy.
	StormControlPolicy pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (PolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyArgs)(nil)).Elem()
}

type PolicyInput interface {
	pulumi.Input

	ToPolicyOutput() PolicyOutput
	ToPolicyOutputWithContext(ctx context.Context) PolicyOutput
}

func (*Policy) ElementType() reflect.Type {
	return reflect.TypeOf((**Policy)(nil)).Elem()
}

func (i *Policy) ToPolicyOutput() PolicyOutput {
	return i.ToPolicyOutputWithContext(context.Background())
}

func (i *Policy) ToPolicyOutputWithContext(ctx context.Context) PolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyOutput)
}

// PolicyArrayInput is an input type that accepts PolicyArray and PolicyArrayOutput values.
// You can construct a concrete instance of `PolicyArrayInput` via:
//
//	PolicyArray{ PolicyArgs{...} }
type PolicyArrayInput interface {
	pulumi.Input

	ToPolicyArrayOutput() PolicyArrayOutput
	ToPolicyArrayOutputWithContext(context.Context) PolicyArrayOutput
}

type PolicyArray []PolicyInput

func (PolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Policy)(nil)).Elem()
}

func (i PolicyArray) ToPolicyArrayOutput() PolicyArrayOutput {
	return i.ToPolicyArrayOutputWithContext(context.Background())
}

func (i PolicyArray) ToPolicyArrayOutputWithContext(ctx context.Context) PolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyArrayOutput)
}

// PolicyMapInput is an input type that accepts PolicyMap and PolicyMapOutput values.
// You can construct a concrete instance of `PolicyMapInput` via:
//
//	PolicyMap{ "key": PolicyArgs{...} }
type PolicyMapInput interface {
	pulumi.Input

	ToPolicyMapOutput() PolicyMapOutput
	ToPolicyMapOutputWithContext(context.Context) PolicyMapOutput
}

type PolicyMap map[string]PolicyInput

func (PolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Policy)(nil)).Elem()
}

func (i PolicyMap) ToPolicyMapOutput() PolicyMapOutput {
	return i.ToPolicyMapOutputWithContext(context.Background())
}

func (i PolicyMap) ToPolicyMapOutputWithContext(ctx context.Context) PolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyMapOutput)
}

type PolicyOutput struct{ *pulumi.OutputState }

func (PolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Policy)(nil)).Elem()
}

func (o PolicyOutput) ToPolicyOutput() PolicyOutput {
	return o
}

func (o PolicyOutput) ToPolicyOutputWithContext(ctx context.Context) PolicyOutput {
	return o
}

// Enable/disable IGMP flood report. Valid values: `enable`, `disable`.
func (o PolicyOutput) IgmpFloodReport() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringOutput { return v.IgmpFloodReport }).(pulumi.StringOutput)
}

// Enable/disable IGMP flood traffic. Valid values: `enable`, `disable`.
func (o PolicyOutput) IgmpFloodTraffic() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringOutput { return v.IgmpFloodTraffic }).(pulumi.StringOutput)
}

// Auto-Config QoS policy name
func (o PolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Enable/disable PoE status. Valid values: `enable`, `disable`.
func (o PolicyOutput) PoeStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringOutput { return v.PoeStatus }).(pulumi.StringOutput)
}

// Auto-Config QoS policy.
func (o PolicyOutput) QosPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringOutput { return v.QosPolicy }).(pulumi.StringOutput)
}

// Auto-Config storm control policy.
func (o PolicyOutput) StormControlPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringOutput { return v.StormControlPolicy }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o PolicyOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type PolicyArrayOutput struct{ *pulumi.OutputState }

func (PolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Policy)(nil)).Elem()
}

func (o PolicyArrayOutput) ToPolicyArrayOutput() PolicyArrayOutput {
	return o
}

func (o PolicyArrayOutput) ToPolicyArrayOutputWithContext(ctx context.Context) PolicyArrayOutput {
	return o
}

func (o PolicyArrayOutput) Index(i pulumi.IntInput) PolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Policy {
		return vs[0].([]*Policy)[vs[1].(int)]
	}).(PolicyOutput)
}

type PolicyMapOutput struct{ *pulumi.OutputState }

func (PolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Policy)(nil)).Elem()
}

func (o PolicyMapOutput) ToPolicyMapOutput() PolicyMapOutput {
	return o
}

func (o PolicyMapOutput) ToPolicyMapOutputWithContext(ctx context.Context) PolicyMapOutput {
	return o
}

func (o PolicyMapOutput) MapIndex(k pulumi.StringInput) PolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Policy {
		return vs[0].(map[string]*Policy)[vs[1].(string)]
	}).(PolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyInput)(nil)).Elem(), &Policy{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyArrayInput)(nil)).Elem(), PolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyMapInput)(nil)).Elem(), PolicyMap{})
	pulumi.RegisterOutputType(PolicyOutput{})
	pulumi.RegisterOutputType(PolicyArrayOutput{})
	pulumi.RegisterOutputType(PolicyMapOutput{})
}
