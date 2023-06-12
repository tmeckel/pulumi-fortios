// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sys

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Coordinate federated upgrades within the Security Fabric. Applies to FortiOS Version `>= 7.0.0`.
//
// ## Import
//
// # System FederatedUpgrade can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:sys/federatedupgrade:Federatedupgrade labelname SystemFederatedUpgrade
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:sys/federatedupgrade:Federatedupgrade labelname SystemFederatedUpgrade
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type Federatedupgrade struct {
	pulumi.CustomResourceState

	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Serial number of the node to include.
	FailureDevice pulumi.StringOutput `pulumi:"failureDevice"`
	// Reason for upgrade failure. Valid values: `none`, `internal`, `timeout`, `device-type-unsupported`, `download-failed`, `device-missing`, `version-unavailable`, `staging-failed`, `reboot-failed`, `device-not-reconnected`, `node-not-ready`, `no-final-confirmation`, `no-confirmation-query`.
	FailureReason pulumi.StringOutput `pulumi:"failureReason"`
	// The index of the next image to upgrade to.
	NextPathIndex pulumi.IntOutput `pulumi:"nextPathIndex"`
	// Nodes which will be included in the upgrade. The structure of `nodeList` block is documented below.
	NodeLists FederatedupgradeNodeListArrayOutput `pulumi:"nodeLists"`
	// Current status of the upgrade.
	Status pulumi.StringOutput `pulumi:"status"`
	// Unique identifier for this upgrade.
	UpgradeId pulumi.IntOutput `pulumi:"upgradeId"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewFederatedupgrade registers a new resource with the given unique name, arguments, and options.
func NewFederatedupgrade(ctx *pulumi.Context,
	name string, args *FederatedupgradeArgs, opts ...pulumi.ResourceOption) (*Federatedupgrade, error) {
	if args == nil {
		args = &FederatedupgradeArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource Federatedupgrade
	err := ctx.RegisterResource("fortios:sys/federatedupgrade:Federatedupgrade", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFederatedupgrade gets an existing Federatedupgrade resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFederatedupgrade(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FederatedupgradeState, opts ...pulumi.ResourceOption) (*Federatedupgrade, error) {
	var resource Federatedupgrade
	err := ctx.ReadResource("fortios:sys/federatedupgrade:Federatedupgrade", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Federatedupgrade resources.
type federatedupgradeState struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Serial number of the node to include.
	FailureDevice *string `pulumi:"failureDevice"`
	// Reason for upgrade failure. Valid values: `none`, `internal`, `timeout`, `device-type-unsupported`, `download-failed`, `device-missing`, `version-unavailable`, `staging-failed`, `reboot-failed`, `device-not-reconnected`, `node-not-ready`, `no-final-confirmation`, `no-confirmation-query`.
	FailureReason *string `pulumi:"failureReason"`
	// The index of the next image to upgrade to.
	NextPathIndex *int `pulumi:"nextPathIndex"`
	// Nodes which will be included in the upgrade. The structure of `nodeList` block is documented below.
	NodeLists []FederatedupgradeNodeList `pulumi:"nodeLists"`
	// Current status of the upgrade.
	Status *string `pulumi:"status"`
	// Unique identifier for this upgrade.
	UpgradeId *int `pulumi:"upgradeId"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type FederatedupgradeState struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Serial number of the node to include.
	FailureDevice pulumi.StringPtrInput
	// Reason for upgrade failure. Valid values: `none`, `internal`, `timeout`, `device-type-unsupported`, `download-failed`, `device-missing`, `version-unavailable`, `staging-failed`, `reboot-failed`, `device-not-reconnected`, `node-not-ready`, `no-final-confirmation`, `no-confirmation-query`.
	FailureReason pulumi.StringPtrInput
	// The index of the next image to upgrade to.
	NextPathIndex pulumi.IntPtrInput
	// Nodes which will be included in the upgrade. The structure of `nodeList` block is documented below.
	NodeLists FederatedupgradeNodeListArrayInput
	// Current status of the upgrade.
	Status pulumi.StringPtrInput
	// Unique identifier for this upgrade.
	UpgradeId pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FederatedupgradeState) ElementType() reflect.Type {
	return reflect.TypeOf((*federatedupgradeState)(nil)).Elem()
}

type federatedupgradeArgs struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Serial number of the node to include.
	FailureDevice *string `pulumi:"failureDevice"`
	// Reason for upgrade failure. Valid values: `none`, `internal`, `timeout`, `device-type-unsupported`, `download-failed`, `device-missing`, `version-unavailable`, `staging-failed`, `reboot-failed`, `device-not-reconnected`, `node-not-ready`, `no-final-confirmation`, `no-confirmation-query`.
	FailureReason *string `pulumi:"failureReason"`
	// The index of the next image to upgrade to.
	NextPathIndex *int `pulumi:"nextPathIndex"`
	// Nodes which will be included in the upgrade. The structure of `nodeList` block is documented below.
	NodeLists []FederatedupgradeNodeList `pulumi:"nodeLists"`
	// Current status of the upgrade.
	Status *string `pulumi:"status"`
	// Unique identifier for this upgrade.
	UpgradeId *int `pulumi:"upgradeId"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Federatedupgrade resource.
type FederatedupgradeArgs struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Serial number of the node to include.
	FailureDevice pulumi.StringPtrInput
	// Reason for upgrade failure. Valid values: `none`, `internal`, `timeout`, `device-type-unsupported`, `download-failed`, `device-missing`, `version-unavailable`, `staging-failed`, `reboot-failed`, `device-not-reconnected`, `node-not-ready`, `no-final-confirmation`, `no-confirmation-query`.
	FailureReason pulumi.StringPtrInput
	// The index of the next image to upgrade to.
	NextPathIndex pulumi.IntPtrInput
	// Nodes which will be included in the upgrade. The structure of `nodeList` block is documented below.
	NodeLists FederatedupgradeNodeListArrayInput
	// Current status of the upgrade.
	Status pulumi.StringPtrInput
	// Unique identifier for this upgrade.
	UpgradeId pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FederatedupgradeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*federatedupgradeArgs)(nil)).Elem()
}

type FederatedupgradeInput interface {
	pulumi.Input

	ToFederatedupgradeOutput() FederatedupgradeOutput
	ToFederatedupgradeOutputWithContext(ctx context.Context) FederatedupgradeOutput
}

func (*Federatedupgrade) ElementType() reflect.Type {
	return reflect.TypeOf((**Federatedupgrade)(nil)).Elem()
}

func (i *Federatedupgrade) ToFederatedupgradeOutput() FederatedupgradeOutput {
	return i.ToFederatedupgradeOutputWithContext(context.Background())
}

func (i *Federatedupgrade) ToFederatedupgradeOutputWithContext(ctx context.Context) FederatedupgradeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedupgradeOutput)
}

// FederatedupgradeArrayInput is an input type that accepts FederatedupgradeArray and FederatedupgradeArrayOutput values.
// You can construct a concrete instance of `FederatedupgradeArrayInput` via:
//
//	FederatedupgradeArray{ FederatedupgradeArgs{...} }
type FederatedupgradeArrayInput interface {
	pulumi.Input

	ToFederatedupgradeArrayOutput() FederatedupgradeArrayOutput
	ToFederatedupgradeArrayOutputWithContext(context.Context) FederatedupgradeArrayOutput
}

type FederatedupgradeArray []FederatedupgradeInput

func (FederatedupgradeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Federatedupgrade)(nil)).Elem()
}

func (i FederatedupgradeArray) ToFederatedupgradeArrayOutput() FederatedupgradeArrayOutput {
	return i.ToFederatedupgradeArrayOutputWithContext(context.Background())
}

func (i FederatedupgradeArray) ToFederatedupgradeArrayOutputWithContext(ctx context.Context) FederatedupgradeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedupgradeArrayOutput)
}

// FederatedupgradeMapInput is an input type that accepts FederatedupgradeMap and FederatedupgradeMapOutput values.
// You can construct a concrete instance of `FederatedupgradeMapInput` via:
//
//	FederatedupgradeMap{ "key": FederatedupgradeArgs{...} }
type FederatedupgradeMapInput interface {
	pulumi.Input

	ToFederatedupgradeMapOutput() FederatedupgradeMapOutput
	ToFederatedupgradeMapOutputWithContext(context.Context) FederatedupgradeMapOutput
}

type FederatedupgradeMap map[string]FederatedupgradeInput

func (FederatedupgradeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Federatedupgrade)(nil)).Elem()
}

func (i FederatedupgradeMap) ToFederatedupgradeMapOutput() FederatedupgradeMapOutput {
	return i.ToFederatedupgradeMapOutputWithContext(context.Background())
}

func (i FederatedupgradeMap) ToFederatedupgradeMapOutputWithContext(ctx context.Context) FederatedupgradeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedupgradeMapOutput)
}

type FederatedupgradeOutput struct{ *pulumi.OutputState }

func (FederatedupgradeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Federatedupgrade)(nil)).Elem()
}

func (o FederatedupgradeOutput) ToFederatedupgradeOutput() FederatedupgradeOutput {
	return o
}

func (o FederatedupgradeOutput) ToFederatedupgradeOutputWithContext(ctx context.Context) FederatedupgradeOutput {
	return o
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o FederatedupgradeOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Federatedupgrade) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Serial number of the node to include.
func (o FederatedupgradeOutput) FailureDevice() pulumi.StringOutput {
	return o.ApplyT(func(v *Federatedupgrade) pulumi.StringOutput { return v.FailureDevice }).(pulumi.StringOutput)
}

// Reason for upgrade failure. Valid values: `none`, `internal`, `timeout`, `device-type-unsupported`, `download-failed`, `device-missing`, `version-unavailable`, `staging-failed`, `reboot-failed`, `device-not-reconnected`, `node-not-ready`, `no-final-confirmation`, `no-confirmation-query`.
func (o FederatedupgradeOutput) FailureReason() pulumi.StringOutput {
	return o.ApplyT(func(v *Federatedupgrade) pulumi.StringOutput { return v.FailureReason }).(pulumi.StringOutput)
}

// The index of the next image to upgrade to.
func (o FederatedupgradeOutput) NextPathIndex() pulumi.IntOutput {
	return o.ApplyT(func(v *Federatedupgrade) pulumi.IntOutput { return v.NextPathIndex }).(pulumi.IntOutput)
}

// Nodes which will be included in the upgrade. The structure of `nodeList` block is documented below.
func (o FederatedupgradeOutput) NodeLists() FederatedupgradeNodeListArrayOutput {
	return o.ApplyT(func(v *Federatedupgrade) FederatedupgradeNodeListArrayOutput { return v.NodeLists }).(FederatedupgradeNodeListArrayOutput)
}

// Current status of the upgrade.
func (o FederatedupgradeOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Federatedupgrade) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Unique identifier for this upgrade.
func (o FederatedupgradeOutput) UpgradeId() pulumi.IntOutput {
	return o.ApplyT(func(v *Federatedupgrade) pulumi.IntOutput { return v.UpgradeId }).(pulumi.IntOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o FederatedupgradeOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Federatedupgrade) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type FederatedupgradeArrayOutput struct{ *pulumi.OutputState }

func (FederatedupgradeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Federatedupgrade)(nil)).Elem()
}

func (o FederatedupgradeArrayOutput) ToFederatedupgradeArrayOutput() FederatedupgradeArrayOutput {
	return o
}

func (o FederatedupgradeArrayOutput) ToFederatedupgradeArrayOutputWithContext(ctx context.Context) FederatedupgradeArrayOutput {
	return o
}

func (o FederatedupgradeArrayOutput) Index(i pulumi.IntInput) FederatedupgradeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Federatedupgrade {
		return vs[0].([]*Federatedupgrade)[vs[1].(int)]
	}).(FederatedupgradeOutput)
}

type FederatedupgradeMapOutput struct{ *pulumi.OutputState }

func (FederatedupgradeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Federatedupgrade)(nil)).Elem()
}

func (o FederatedupgradeMapOutput) ToFederatedupgradeMapOutput() FederatedupgradeMapOutput {
	return o
}

func (o FederatedupgradeMapOutput) ToFederatedupgradeMapOutputWithContext(ctx context.Context) FederatedupgradeMapOutput {
	return o
}

func (o FederatedupgradeMapOutput) MapIndex(k pulumi.StringInput) FederatedupgradeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Federatedupgrade {
		return vs[0].(map[string]*Federatedupgrade)[vs[1].(string)]
	}).(FederatedupgradeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedupgradeInput)(nil)).Elem(), &Federatedupgrade{})
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedupgradeArrayInput)(nil)).Elem(), FederatedupgradeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedupgradeMapInput)(nil)).Elem(), FederatedupgradeMap{})
	pulumi.RegisterOutputType(FederatedupgradeOutput{})
	pulumi.RegisterOutputType(FederatedupgradeArrayOutput{})
	pulumi.RegisterOutputType(FederatedupgradeMapOutput{})
}