// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package switchcontroller

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure VLANs for switch controller. Applies to FortiOS Version `<= 6.2.0`.
//
// ## Import
//
// # SwitchController Vlan can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:switchcontroller/vlan:Vlan labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:switchcontroller/vlan:Vlan labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type Vlan struct {
	pulumi.CustomResourceState

	// Authentication. Valid values: `radius`, `usergroup`.
	Auth pulumi.StringOutput `pulumi:"auth"`
	// Color of icon on the GUI.
	Color pulumi.IntOutput `pulumi:"color"`
	// Comment.
	Comments pulumi.StringOutput `pulumi:"comments"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Switch VLAN name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specify captive portal replacement message override group.
	PortalMessageOverrideGroup pulumi.StringOutput `pulumi:"portalMessageOverrideGroup"`
	// Individual message overrides. The structure of `portalMessageOverrides` block is documented below.
	PortalMessageOverrides VlanPortalMessageOverridesOutput `pulumi:"portalMessageOverrides"`
	// Authentication radius server.
	RadiusServer pulumi.StringOutput `pulumi:"radiusServer"`
	// Security. Valid values: `open`, `captive-portal`, `8021x`.
	Security pulumi.StringOutput `pulumi:"security"`
	// Selected user group. The structure of `selectedUsergroups` block is documented below.
	SelectedUsergroups VlanSelectedUsergroupArrayOutput `pulumi:"selectedUsergroups"`
	// Authentication usergroup.
	Usergroup pulumi.StringOutput `pulumi:"usergroup"`
	// Virtual domain,
	Vdom pulumi.StringOutput `pulumi:"vdom"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// VLAN ID.
	Vlanid pulumi.IntOutput `pulumi:"vlanid"`
}

// NewVlan registers a new resource with the given unique name, arguments, and options.
func NewVlan(ctx *pulumi.Context,
	name string, args *VlanArgs, opts ...pulumi.ResourceOption) (*Vlan, error) {
	if args == nil {
		args = &VlanArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource Vlan
	err := ctx.RegisterResource("fortios:switchcontroller/vlan:Vlan", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVlan gets an existing Vlan resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVlan(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VlanState, opts ...pulumi.ResourceOption) (*Vlan, error) {
	var resource Vlan
	err := ctx.ReadResource("fortios:switchcontroller/vlan:Vlan", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Vlan resources.
type vlanState struct {
	// Authentication. Valid values: `radius`, `usergroup`.
	Auth *string `pulumi:"auth"`
	// Color of icon on the GUI.
	Color *int `pulumi:"color"`
	// Comment.
	Comments *string `pulumi:"comments"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Switch VLAN name.
	Name *string `pulumi:"name"`
	// Specify captive portal replacement message override group.
	PortalMessageOverrideGroup *string `pulumi:"portalMessageOverrideGroup"`
	// Individual message overrides. The structure of `portalMessageOverrides` block is documented below.
	PortalMessageOverrides *VlanPortalMessageOverrides `pulumi:"portalMessageOverrides"`
	// Authentication radius server.
	RadiusServer *string `pulumi:"radiusServer"`
	// Security. Valid values: `open`, `captive-portal`, `8021x`.
	Security *string `pulumi:"security"`
	// Selected user group. The structure of `selectedUsergroups` block is documented below.
	SelectedUsergroups []VlanSelectedUsergroup `pulumi:"selectedUsergroups"`
	// Authentication usergroup.
	Usergroup *string `pulumi:"usergroup"`
	// Virtual domain,
	Vdom *string `pulumi:"vdom"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// VLAN ID.
	Vlanid *int `pulumi:"vlanid"`
}

type VlanState struct {
	// Authentication. Valid values: `radius`, `usergroup`.
	Auth pulumi.StringPtrInput
	// Color of icon on the GUI.
	Color pulumi.IntPtrInput
	// Comment.
	Comments pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Switch VLAN name.
	Name pulumi.StringPtrInput
	// Specify captive portal replacement message override group.
	PortalMessageOverrideGroup pulumi.StringPtrInput
	// Individual message overrides. The structure of `portalMessageOverrides` block is documented below.
	PortalMessageOverrides VlanPortalMessageOverridesPtrInput
	// Authentication radius server.
	RadiusServer pulumi.StringPtrInput
	// Security. Valid values: `open`, `captive-portal`, `8021x`.
	Security pulumi.StringPtrInput
	// Selected user group. The structure of `selectedUsergroups` block is documented below.
	SelectedUsergroups VlanSelectedUsergroupArrayInput
	// Authentication usergroup.
	Usergroup pulumi.StringPtrInput
	// Virtual domain,
	Vdom pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// VLAN ID.
	Vlanid pulumi.IntPtrInput
}

func (VlanState) ElementType() reflect.Type {
	return reflect.TypeOf((*vlanState)(nil)).Elem()
}

type vlanArgs struct {
	// Authentication. Valid values: `radius`, `usergroup`.
	Auth *string `pulumi:"auth"`
	// Color of icon on the GUI.
	Color *int `pulumi:"color"`
	// Comment.
	Comments *string `pulumi:"comments"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Switch VLAN name.
	Name *string `pulumi:"name"`
	// Specify captive portal replacement message override group.
	PortalMessageOverrideGroup *string `pulumi:"portalMessageOverrideGroup"`
	// Individual message overrides. The structure of `portalMessageOverrides` block is documented below.
	PortalMessageOverrides *VlanPortalMessageOverrides `pulumi:"portalMessageOverrides"`
	// Authentication radius server.
	RadiusServer *string `pulumi:"radiusServer"`
	// Security. Valid values: `open`, `captive-portal`, `8021x`.
	Security *string `pulumi:"security"`
	// Selected user group. The structure of `selectedUsergroups` block is documented below.
	SelectedUsergroups []VlanSelectedUsergroup `pulumi:"selectedUsergroups"`
	// Authentication usergroup.
	Usergroup *string `pulumi:"usergroup"`
	// Virtual domain,
	Vdom *string `pulumi:"vdom"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// VLAN ID.
	Vlanid *int `pulumi:"vlanid"`
}

// The set of arguments for constructing a Vlan resource.
type VlanArgs struct {
	// Authentication. Valid values: `radius`, `usergroup`.
	Auth pulumi.StringPtrInput
	// Color of icon on the GUI.
	Color pulumi.IntPtrInput
	// Comment.
	Comments pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Switch VLAN name.
	Name pulumi.StringPtrInput
	// Specify captive portal replacement message override group.
	PortalMessageOverrideGroup pulumi.StringPtrInput
	// Individual message overrides. The structure of `portalMessageOverrides` block is documented below.
	PortalMessageOverrides VlanPortalMessageOverridesPtrInput
	// Authentication radius server.
	RadiusServer pulumi.StringPtrInput
	// Security. Valid values: `open`, `captive-portal`, `8021x`.
	Security pulumi.StringPtrInput
	// Selected user group. The structure of `selectedUsergroups` block is documented below.
	SelectedUsergroups VlanSelectedUsergroupArrayInput
	// Authentication usergroup.
	Usergroup pulumi.StringPtrInput
	// Virtual domain,
	Vdom pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// VLAN ID.
	Vlanid pulumi.IntPtrInput
}

func (VlanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vlanArgs)(nil)).Elem()
}

type VlanInput interface {
	pulumi.Input

	ToVlanOutput() VlanOutput
	ToVlanOutputWithContext(ctx context.Context) VlanOutput
}

func (*Vlan) ElementType() reflect.Type {
	return reflect.TypeOf((**Vlan)(nil)).Elem()
}

func (i *Vlan) ToVlanOutput() VlanOutput {
	return i.ToVlanOutputWithContext(context.Background())
}

func (i *Vlan) ToVlanOutputWithContext(ctx context.Context) VlanOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VlanOutput)
}

// VlanArrayInput is an input type that accepts VlanArray and VlanArrayOutput values.
// You can construct a concrete instance of `VlanArrayInput` via:
//
//	VlanArray{ VlanArgs{...} }
type VlanArrayInput interface {
	pulumi.Input

	ToVlanArrayOutput() VlanArrayOutput
	ToVlanArrayOutputWithContext(context.Context) VlanArrayOutput
}

type VlanArray []VlanInput

func (VlanArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Vlan)(nil)).Elem()
}

func (i VlanArray) ToVlanArrayOutput() VlanArrayOutput {
	return i.ToVlanArrayOutputWithContext(context.Background())
}

func (i VlanArray) ToVlanArrayOutputWithContext(ctx context.Context) VlanArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VlanArrayOutput)
}

// VlanMapInput is an input type that accepts VlanMap and VlanMapOutput values.
// You can construct a concrete instance of `VlanMapInput` via:
//
//	VlanMap{ "key": VlanArgs{...} }
type VlanMapInput interface {
	pulumi.Input

	ToVlanMapOutput() VlanMapOutput
	ToVlanMapOutputWithContext(context.Context) VlanMapOutput
}

type VlanMap map[string]VlanInput

func (VlanMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Vlan)(nil)).Elem()
}

func (i VlanMap) ToVlanMapOutput() VlanMapOutput {
	return i.ToVlanMapOutputWithContext(context.Background())
}

func (i VlanMap) ToVlanMapOutputWithContext(ctx context.Context) VlanMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VlanMapOutput)
}

type VlanOutput struct{ *pulumi.OutputState }

func (VlanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Vlan)(nil)).Elem()
}

func (o VlanOutput) ToVlanOutput() VlanOutput {
	return o
}

func (o VlanOutput) ToVlanOutputWithContext(ctx context.Context) VlanOutput {
	return o
}

// Authentication. Valid values: `radius`, `usergroup`.
func (o VlanOutput) Auth() pulumi.StringOutput {
	return o.ApplyT(func(v *Vlan) pulumi.StringOutput { return v.Auth }).(pulumi.StringOutput)
}

// Color of icon on the GUI.
func (o VlanOutput) Color() pulumi.IntOutput {
	return o.ApplyT(func(v *Vlan) pulumi.IntOutput { return v.Color }).(pulumi.IntOutput)
}

// Comment.
func (o VlanOutput) Comments() pulumi.StringOutput {
	return o.ApplyT(func(v *Vlan) pulumi.StringOutput { return v.Comments }).(pulumi.StringOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o VlanOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Vlan) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Switch VLAN name.
func (o VlanOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Vlan) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specify captive portal replacement message override group.
func (o VlanOutput) PortalMessageOverrideGroup() pulumi.StringOutput {
	return o.ApplyT(func(v *Vlan) pulumi.StringOutput { return v.PortalMessageOverrideGroup }).(pulumi.StringOutput)
}

// Individual message overrides. The structure of `portalMessageOverrides` block is documented below.
func (o VlanOutput) PortalMessageOverrides() VlanPortalMessageOverridesOutput {
	return o.ApplyT(func(v *Vlan) VlanPortalMessageOverridesOutput { return v.PortalMessageOverrides }).(VlanPortalMessageOverridesOutput)
}

// Authentication radius server.
func (o VlanOutput) RadiusServer() pulumi.StringOutput {
	return o.ApplyT(func(v *Vlan) pulumi.StringOutput { return v.RadiusServer }).(pulumi.StringOutput)
}

// Security. Valid values: `open`, `captive-portal`, `8021x`.
func (o VlanOutput) Security() pulumi.StringOutput {
	return o.ApplyT(func(v *Vlan) pulumi.StringOutput { return v.Security }).(pulumi.StringOutput)
}

// Selected user group. The structure of `selectedUsergroups` block is documented below.
func (o VlanOutput) SelectedUsergroups() VlanSelectedUsergroupArrayOutput {
	return o.ApplyT(func(v *Vlan) VlanSelectedUsergroupArrayOutput { return v.SelectedUsergroups }).(VlanSelectedUsergroupArrayOutput)
}

// Authentication usergroup.
func (o VlanOutput) Usergroup() pulumi.StringOutput {
	return o.ApplyT(func(v *Vlan) pulumi.StringOutput { return v.Usergroup }).(pulumi.StringOutput)
}

// Virtual domain,
func (o VlanOutput) Vdom() pulumi.StringOutput {
	return o.ApplyT(func(v *Vlan) pulumi.StringOutput { return v.Vdom }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o VlanOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Vlan) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// VLAN ID.
func (o VlanOutput) Vlanid() pulumi.IntOutput {
	return o.ApplyT(func(v *Vlan) pulumi.IntOutput { return v.Vlanid }).(pulumi.IntOutput)
}

type VlanArrayOutput struct{ *pulumi.OutputState }

func (VlanArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Vlan)(nil)).Elem()
}

func (o VlanArrayOutput) ToVlanArrayOutput() VlanArrayOutput {
	return o
}

func (o VlanArrayOutput) ToVlanArrayOutputWithContext(ctx context.Context) VlanArrayOutput {
	return o
}

func (o VlanArrayOutput) Index(i pulumi.IntInput) VlanOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Vlan {
		return vs[0].([]*Vlan)[vs[1].(int)]
	}).(VlanOutput)
}

type VlanMapOutput struct{ *pulumi.OutputState }

func (VlanMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Vlan)(nil)).Elem()
}

func (o VlanMapOutput) ToVlanMapOutput() VlanMapOutput {
	return o
}

func (o VlanMapOutput) ToVlanMapOutputWithContext(ctx context.Context) VlanMapOutput {
	return o
}

func (o VlanMapOutput) MapIndex(k pulumi.StringInput) VlanOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Vlan {
		return vs[0].(map[string]*Vlan)[vs[1].(string)]
	}).(VlanOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VlanInput)(nil)).Elem(), &Vlan{})
	pulumi.RegisterInputType(reflect.TypeOf((*VlanArrayInput)(nil)).Elem(), VlanArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VlanMapInput)(nil)).Elem(), VlanMap{})
	pulumi.RegisterOutputType(VlanOutput{})
	pulumi.RegisterOutputType(VlanArrayOutput{})
	pulumi.RegisterOutputType(VlanMapOutput{})
}
