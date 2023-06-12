// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package user

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure NAC policy matching pattern to identify matching NAC devices. Applies to FortiOS Version `>= 6.4.0`.
//
// ## Import
//
// # User NacPolicy can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:user/nacpolicy:Nacpolicy labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:user/nacpolicy:Nacpolicy labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type Nacpolicy struct {
	pulumi.CustomResourceState

	// Category of NAC policy.
	Category pulumi.StringOutput `pulumi:"category"`
	// Description for the NAC policy matching pattern.
	Description pulumi.StringOutput `pulumi:"description"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// NAC policy matching EMS tag.
	EmsTag pulumi.StringOutput `pulumi:"emsTag"`
	// NAC policy matching family.
	Family pulumi.StringOutput `pulumi:"family"`
	// Dynamic firewall address to associate MAC which match this policy.
	FirewallAddress pulumi.StringOutput `pulumi:"firewallAddress"`
	// NAC policy matching host.
	Host pulumi.StringOutput `pulumi:"host"`
	// NAC policy matching hardware vendor.
	HwVendor pulumi.StringOutput `pulumi:"hwVendor"`
	// NAC policy matching hardware version.
	HwVersion pulumi.StringOutput `pulumi:"hwVersion"`
	// NAC policy matching MAC address.
	Mac pulumi.StringOutput `pulumi:"mac"`
	// NAC policy name.
	Name pulumi.StringOutput `pulumi:"name"`
	// NAC policy matching operating system.
	Os pulumi.StringOutput `pulumi:"os"`
	// NAC policy matching source.
	Src pulumi.StringOutput `pulumi:"src"`
	// SSID policy to be applied on the matched NAC policy.
	SsidPolicy pulumi.StringOutput `pulumi:"ssidPolicy"`
	// Enable/disable NAC policy. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// NAC policy matching software version.
	SwVersion pulumi.StringOutput `pulumi:"swVersion"`
	// NAC device auto authorization when discovered and nac-policy matched. Valid values: `global`, `disable`, `enable`.
	SwitchAutoAuth pulumi.StringOutput `pulumi:"switchAutoAuth"`
	// FortiLink interface for which this NAC policy belongs to.
	SwitchFortilink pulumi.StringOutput `pulumi:"switchFortilink"`
	// List of managed FortiSwitch groups on which NAC policy can be applied. The structure of `switchGroup` block is documented below.
	SwitchGroups NacpolicySwitchGroupArrayOutput `pulumi:"switchGroups"`
	// switch-mac-policy to be applied on the matched NAC policy.
	SwitchMacPolicy pulumi.StringOutput `pulumi:"switchMacPolicy"`
	// switch-port-policy to be applied on the matched NAC policy.
	SwitchPortPolicy pulumi.StringOutput `pulumi:"switchPortPolicy"`
	// List of managed FortiSwitches on which NAC policy can be applied. The structure of `switchScope` block is documented below.
	SwitchScopes NacpolicySwitchScopeArrayOutput `pulumi:"switchScopes"`
	// NAC policy matching type.
	Type pulumi.StringOutput `pulumi:"type"`
	// NAC policy matching user.
	User pulumi.StringOutput `pulumi:"user"`
	// NAC policy matching user group.
	UserGroup pulumi.StringOutput `pulumi:"userGroup"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewNacpolicy registers a new resource with the given unique name, arguments, and options.
func NewNacpolicy(ctx *pulumi.Context,
	name string, args *NacpolicyArgs, opts ...pulumi.ResourceOption) (*Nacpolicy, error) {
	if args == nil {
		args = &NacpolicyArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource Nacpolicy
	err := ctx.RegisterResource("fortios:user/nacpolicy:Nacpolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNacpolicy gets an existing Nacpolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNacpolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NacpolicyState, opts ...pulumi.ResourceOption) (*Nacpolicy, error) {
	var resource Nacpolicy
	err := ctx.ReadResource("fortios:user/nacpolicy:Nacpolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Nacpolicy resources.
type nacpolicyState struct {
	// Category of NAC policy.
	Category *string `pulumi:"category"`
	// Description for the NAC policy matching pattern.
	Description *string `pulumi:"description"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// NAC policy matching EMS tag.
	EmsTag *string `pulumi:"emsTag"`
	// NAC policy matching family.
	Family *string `pulumi:"family"`
	// Dynamic firewall address to associate MAC which match this policy.
	FirewallAddress *string `pulumi:"firewallAddress"`
	// NAC policy matching host.
	Host *string `pulumi:"host"`
	// NAC policy matching hardware vendor.
	HwVendor *string `pulumi:"hwVendor"`
	// NAC policy matching hardware version.
	HwVersion *string `pulumi:"hwVersion"`
	// NAC policy matching MAC address.
	Mac *string `pulumi:"mac"`
	// NAC policy name.
	Name *string `pulumi:"name"`
	// NAC policy matching operating system.
	Os *string `pulumi:"os"`
	// NAC policy matching source.
	Src *string `pulumi:"src"`
	// SSID policy to be applied on the matched NAC policy.
	SsidPolicy *string `pulumi:"ssidPolicy"`
	// Enable/disable NAC policy. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// NAC policy matching software version.
	SwVersion *string `pulumi:"swVersion"`
	// NAC device auto authorization when discovered and nac-policy matched. Valid values: `global`, `disable`, `enable`.
	SwitchAutoAuth *string `pulumi:"switchAutoAuth"`
	// FortiLink interface for which this NAC policy belongs to.
	SwitchFortilink *string `pulumi:"switchFortilink"`
	// List of managed FortiSwitch groups on which NAC policy can be applied. The structure of `switchGroup` block is documented below.
	SwitchGroups []NacpolicySwitchGroup `pulumi:"switchGroups"`
	// switch-mac-policy to be applied on the matched NAC policy.
	SwitchMacPolicy *string `pulumi:"switchMacPolicy"`
	// switch-port-policy to be applied on the matched NAC policy.
	SwitchPortPolicy *string `pulumi:"switchPortPolicy"`
	// List of managed FortiSwitches on which NAC policy can be applied. The structure of `switchScope` block is documented below.
	SwitchScopes []NacpolicySwitchScope `pulumi:"switchScopes"`
	// NAC policy matching type.
	Type *string `pulumi:"type"`
	// NAC policy matching user.
	User *string `pulumi:"user"`
	// NAC policy matching user group.
	UserGroup *string `pulumi:"userGroup"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type NacpolicyState struct {
	// Category of NAC policy.
	Category pulumi.StringPtrInput
	// Description for the NAC policy matching pattern.
	Description pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// NAC policy matching EMS tag.
	EmsTag pulumi.StringPtrInput
	// NAC policy matching family.
	Family pulumi.StringPtrInput
	// Dynamic firewall address to associate MAC which match this policy.
	FirewallAddress pulumi.StringPtrInput
	// NAC policy matching host.
	Host pulumi.StringPtrInput
	// NAC policy matching hardware vendor.
	HwVendor pulumi.StringPtrInput
	// NAC policy matching hardware version.
	HwVersion pulumi.StringPtrInput
	// NAC policy matching MAC address.
	Mac pulumi.StringPtrInput
	// NAC policy name.
	Name pulumi.StringPtrInput
	// NAC policy matching operating system.
	Os pulumi.StringPtrInput
	// NAC policy matching source.
	Src pulumi.StringPtrInput
	// SSID policy to be applied on the matched NAC policy.
	SsidPolicy pulumi.StringPtrInput
	// Enable/disable NAC policy. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// NAC policy matching software version.
	SwVersion pulumi.StringPtrInput
	// NAC device auto authorization when discovered and nac-policy matched. Valid values: `global`, `disable`, `enable`.
	SwitchAutoAuth pulumi.StringPtrInput
	// FortiLink interface for which this NAC policy belongs to.
	SwitchFortilink pulumi.StringPtrInput
	// List of managed FortiSwitch groups on which NAC policy can be applied. The structure of `switchGroup` block is documented below.
	SwitchGroups NacpolicySwitchGroupArrayInput
	// switch-mac-policy to be applied on the matched NAC policy.
	SwitchMacPolicy pulumi.StringPtrInput
	// switch-port-policy to be applied on the matched NAC policy.
	SwitchPortPolicy pulumi.StringPtrInput
	// List of managed FortiSwitches on which NAC policy can be applied. The structure of `switchScope` block is documented below.
	SwitchScopes NacpolicySwitchScopeArrayInput
	// NAC policy matching type.
	Type pulumi.StringPtrInput
	// NAC policy matching user.
	User pulumi.StringPtrInput
	// NAC policy matching user group.
	UserGroup pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (NacpolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*nacpolicyState)(nil)).Elem()
}

type nacpolicyArgs struct {
	// Category of NAC policy.
	Category *string `pulumi:"category"`
	// Description for the NAC policy matching pattern.
	Description *string `pulumi:"description"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// NAC policy matching EMS tag.
	EmsTag *string `pulumi:"emsTag"`
	// NAC policy matching family.
	Family *string `pulumi:"family"`
	// Dynamic firewall address to associate MAC which match this policy.
	FirewallAddress *string `pulumi:"firewallAddress"`
	// NAC policy matching host.
	Host *string `pulumi:"host"`
	// NAC policy matching hardware vendor.
	HwVendor *string `pulumi:"hwVendor"`
	// NAC policy matching hardware version.
	HwVersion *string `pulumi:"hwVersion"`
	// NAC policy matching MAC address.
	Mac *string `pulumi:"mac"`
	// NAC policy name.
	Name *string `pulumi:"name"`
	// NAC policy matching operating system.
	Os *string `pulumi:"os"`
	// NAC policy matching source.
	Src *string `pulumi:"src"`
	// SSID policy to be applied on the matched NAC policy.
	SsidPolicy *string `pulumi:"ssidPolicy"`
	// Enable/disable NAC policy. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// NAC policy matching software version.
	SwVersion *string `pulumi:"swVersion"`
	// NAC device auto authorization when discovered and nac-policy matched. Valid values: `global`, `disable`, `enable`.
	SwitchAutoAuth *string `pulumi:"switchAutoAuth"`
	// FortiLink interface for which this NAC policy belongs to.
	SwitchFortilink *string `pulumi:"switchFortilink"`
	// List of managed FortiSwitch groups on which NAC policy can be applied. The structure of `switchGroup` block is documented below.
	SwitchGroups []NacpolicySwitchGroup `pulumi:"switchGroups"`
	// switch-mac-policy to be applied on the matched NAC policy.
	SwitchMacPolicy *string `pulumi:"switchMacPolicy"`
	// switch-port-policy to be applied on the matched NAC policy.
	SwitchPortPolicy *string `pulumi:"switchPortPolicy"`
	// List of managed FortiSwitches on which NAC policy can be applied. The structure of `switchScope` block is documented below.
	SwitchScopes []NacpolicySwitchScope `pulumi:"switchScopes"`
	// NAC policy matching type.
	Type *string `pulumi:"type"`
	// NAC policy matching user.
	User *string `pulumi:"user"`
	// NAC policy matching user group.
	UserGroup *string `pulumi:"userGroup"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Nacpolicy resource.
type NacpolicyArgs struct {
	// Category of NAC policy.
	Category pulumi.StringPtrInput
	// Description for the NAC policy matching pattern.
	Description pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// NAC policy matching EMS tag.
	EmsTag pulumi.StringPtrInput
	// NAC policy matching family.
	Family pulumi.StringPtrInput
	// Dynamic firewall address to associate MAC which match this policy.
	FirewallAddress pulumi.StringPtrInput
	// NAC policy matching host.
	Host pulumi.StringPtrInput
	// NAC policy matching hardware vendor.
	HwVendor pulumi.StringPtrInput
	// NAC policy matching hardware version.
	HwVersion pulumi.StringPtrInput
	// NAC policy matching MAC address.
	Mac pulumi.StringPtrInput
	// NAC policy name.
	Name pulumi.StringPtrInput
	// NAC policy matching operating system.
	Os pulumi.StringPtrInput
	// NAC policy matching source.
	Src pulumi.StringPtrInput
	// SSID policy to be applied on the matched NAC policy.
	SsidPolicy pulumi.StringPtrInput
	// Enable/disable NAC policy. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// NAC policy matching software version.
	SwVersion pulumi.StringPtrInput
	// NAC device auto authorization when discovered and nac-policy matched. Valid values: `global`, `disable`, `enable`.
	SwitchAutoAuth pulumi.StringPtrInput
	// FortiLink interface for which this NAC policy belongs to.
	SwitchFortilink pulumi.StringPtrInput
	// List of managed FortiSwitch groups on which NAC policy can be applied. The structure of `switchGroup` block is documented below.
	SwitchGroups NacpolicySwitchGroupArrayInput
	// switch-mac-policy to be applied on the matched NAC policy.
	SwitchMacPolicy pulumi.StringPtrInput
	// switch-port-policy to be applied on the matched NAC policy.
	SwitchPortPolicy pulumi.StringPtrInput
	// List of managed FortiSwitches on which NAC policy can be applied. The structure of `switchScope` block is documented below.
	SwitchScopes NacpolicySwitchScopeArrayInput
	// NAC policy matching type.
	Type pulumi.StringPtrInput
	// NAC policy matching user.
	User pulumi.StringPtrInput
	// NAC policy matching user group.
	UserGroup pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (NacpolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nacpolicyArgs)(nil)).Elem()
}

type NacpolicyInput interface {
	pulumi.Input

	ToNacpolicyOutput() NacpolicyOutput
	ToNacpolicyOutputWithContext(ctx context.Context) NacpolicyOutput
}

func (*Nacpolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**Nacpolicy)(nil)).Elem()
}

func (i *Nacpolicy) ToNacpolicyOutput() NacpolicyOutput {
	return i.ToNacpolicyOutputWithContext(context.Background())
}

func (i *Nacpolicy) ToNacpolicyOutputWithContext(ctx context.Context) NacpolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NacpolicyOutput)
}

// NacpolicyArrayInput is an input type that accepts NacpolicyArray and NacpolicyArrayOutput values.
// You can construct a concrete instance of `NacpolicyArrayInput` via:
//
//	NacpolicyArray{ NacpolicyArgs{...} }
type NacpolicyArrayInput interface {
	pulumi.Input

	ToNacpolicyArrayOutput() NacpolicyArrayOutput
	ToNacpolicyArrayOutputWithContext(context.Context) NacpolicyArrayOutput
}

type NacpolicyArray []NacpolicyInput

func (NacpolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Nacpolicy)(nil)).Elem()
}

func (i NacpolicyArray) ToNacpolicyArrayOutput() NacpolicyArrayOutput {
	return i.ToNacpolicyArrayOutputWithContext(context.Background())
}

func (i NacpolicyArray) ToNacpolicyArrayOutputWithContext(ctx context.Context) NacpolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NacpolicyArrayOutput)
}

// NacpolicyMapInput is an input type that accepts NacpolicyMap and NacpolicyMapOutput values.
// You can construct a concrete instance of `NacpolicyMapInput` via:
//
//	NacpolicyMap{ "key": NacpolicyArgs{...} }
type NacpolicyMapInput interface {
	pulumi.Input

	ToNacpolicyMapOutput() NacpolicyMapOutput
	ToNacpolicyMapOutputWithContext(context.Context) NacpolicyMapOutput
}

type NacpolicyMap map[string]NacpolicyInput

func (NacpolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Nacpolicy)(nil)).Elem()
}

func (i NacpolicyMap) ToNacpolicyMapOutput() NacpolicyMapOutput {
	return i.ToNacpolicyMapOutputWithContext(context.Background())
}

func (i NacpolicyMap) ToNacpolicyMapOutputWithContext(ctx context.Context) NacpolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NacpolicyMapOutput)
}

type NacpolicyOutput struct{ *pulumi.OutputState }

func (NacpolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Nacpolicy)(nil)).Elem()
}

func (o NacpolicyOutput) ToNacpolicyOutput() NacpolicyOutput {
	return o
}

func (o NacpolicyOutput) ToNacpolicyOutputWithContext(ctx context.Context) NacpolicyOutput {
	return o
}

// Category of NAC policy.
func (o NacpolicyOutput) Category() pulumi.StringOutput {
	return o.ApplyT(func(v *Nacpolicy) pulumi.StringOutput { return v.Category }).(pulumi.StringOutput)
}

// Description for the NAC policy matching pattern.
func (o NacpolicyOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Nacpolicy) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o NacpolicyOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Nacpolicy) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// NAC policy matching EMS tag.
func (o NacpolicyOutput) EmsTag() pulumi.StringOutput {
	return o.ApplyT(func(v *Nacpolicy) pulumi.StringOutput { return v.EmsTag }).(pulumi.StringOutput)
}

// NAC policy matching family.
func (o NacpolicyOutput) Family() pulumi.StringOutput {
	return o.ApplyT(func(v *Nacpolicy) pulumi.StringOutput { return v.Family }).(pulumi.StringOutput)
}

// Dynamic firewall address to associate MAC which match this policy.
func (o NacpolicyOutput) FirewallAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *Nacpolicy) pulumi.StringOutput { return v.FirewallAddress }).(pulumi.StringOutput)
}

// NAC policy matching host.
func (o NacpolicyOutput) Host() pulumi.StringOutput {
	return o.ApplyT(func(v *Nacpolicy) pulumi.StringOutput { return v.Host }).(pulumi.StringOutput)
}

// NAC policy matching hardware vendor.
func (o NacpolicyOutput) HwVendor() pulumi.StringOutput {
	return o.ApplyT(func(v *Nacpolicy) pulumi.StringOutput { return v.HwVendor }).(pulumi.StringOutput)
}

// NAC policy matching hardware version.
func (o NacpolicyOutput) HwVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Nacpolicy) pulumi.StringOutput { return v.HwVersion }).(pulumi.StringOutput)
}

// NAC policy matching MAC address.
func (o NacpolicyOutput) Mac() pulumi.StringOutput {
	return o.ApplyT(func(v *Nacpolicy) pulumi.StringOutput { return v.Mac }).(pulumi.StringOutput)
}

// NAC policy name.
func (o NacpolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Nacpolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// NAC policy matching operating system.
func (o NacpolicyOutput) Os() pulumi.StringOutput {
	return o.ApplyT(func(v *Nacpolicy) pulumi.StringOutput { return v.Os }).(pulumi.StringOutput)
}

// NAC policy matching source.
func (o NacpolicyOutput) Src() pulumi.StringOutput {
	return o.ApplyT(func(v *Nacpolicy) pulumi.StringOutput { return v.Src }).(pulumi.StringOutput)
}

// SSID policy to be applied on the matched NAC policy.
func (o NacpolicyOutput) SsidPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v *Nacpolicy) pulumi.StringOutput { return v.SsidPolicy }).(pulumi.StringOutput)
}

// Enable/disable NAC policy. Valid values: `enable`, `disable`.
func (o NacpolicyOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Nacpolicy) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// NAC policy matching software version.
func (o NacpolicyOutput) SwVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Nacpolicy) pulumi.StringOutput { return v.SwVersion }).(pulumi.StringOutput)
}

// NAC device auto authorization when discovered and nac-policy matched. Valid values: `global`, `disable`, `enable`.
func (o NacpolicyOutput) SwitchAutoAuth() pulumi.StringOutput {
	return o.ApplyT(func(v *Nacpolicy) pulumi.StringOutput { return v.SwitchAutoAuth }).(pulumi.StringOutput)
}

// FortiLink interface for which this NAC policy belongs to.
func (o NacpolicyOutput) SwitchFortilink() pulumi.StringOutput {
	return o.ApplyT(func(v *Nacpolicy) pulumi.StringOutput { return v.SwitchFortilink }).(pulumi.StringOutput)
}

// List of managed FortiSwitch groups on which NAC policy can be applied. The structure of `switchGroup` block is documented below.
func (o NacpolicyOutput) SwitchGroups() NacpolicySwitchGroupArrayOutput {
	return o.ApplyT(func(v *Nacpolicy) NacpolicySwitchGroupArrayOutput { return v.SwitchGroups }).(NacpolicySwitchGroupArrayOutput)
}

// switch-mac-policy to be applied on the matched NAC policy.
func (o NacpolicyOutput) SwitchMacPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v *Nacpolicy) pulumi.StringOutput { return v.SwitchMacPolicy }).(pulumi.StringOutput)
}

// switch-port-policy to be applied on the matched NAC policy.
func (o NacpolicyOutput) SwitchPortPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v *Nacpolicy) pulumi.StringOutput { return v.SwitchPortPolicy }).(pulumi.StringOutput)
}

// List of managed FortiSwitches on which NAC policy can be applied. The structure of `switchScope` block is documented below.
func (o NacpolicyOutput) SwitchScopes() NacpolicySwitchScopeArrayOutput {
	return o.ApplyT(func(v *Nacpolicy) NacpolicySwitchScopeArrayOutput { return v.SwitchScopes }).(NacpolicySwitchScopeArrayOutput)
}

// NAC policy matching type.
func (o NacpolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Nacpolicy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// NAC policy matching user.
func (o NacpolicyOutput) User() pulumi.StringOutput {
	return o.ApplyT(func(v *Nacpolicy) pulumi.StringOutput { return v.User }).(pulumi.StringOutput)
}

// NAC policy matching user group.
func (o NacpolicyOutput) UserGroup() pulumi.StringOutput {
	return o.ApplyT(func(v *Nacpolicy) pulumi.StringOutput { return v.UserGroup }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o NacpolicyOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Nacpolicy) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type NacpolicyArrayOutput struct{ *pulumi.OutputState }

func (NacpolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Nacpolicy)(nil)).Elem()
}

func (o NacpolicyArrayOutput) ToNacpolicyArrayOutput() NacpolicyArrayOutput {
	return o
}

func (o NacpolicyArrayOutput) ToNacpolicyArrayOutputWithContext(ctx context.Context) NacpolicyArrayOutput {
	return o
}

func (o NacpolicyArrayOutput) Index(i pulumi.IntInput) NacpolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Nacpolicy {
		return vs[0].([]*Nacpolicy)[vs[1].(int)]
	}).(NacpolicyOutput)
}

type NacpolicyMapOutput struct{ *pulumi.OutputState }

func (NacpolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Nacpolicy)(nil)).Elem()
}

func (o NacpolicyMapOutput) ToNacpolicyMapOutput() NacpolicyMapOutput {
	return o
}

func (o NacpolicyMapOutput) ToNacpolicyMapOutputWithContext(ctx context.Context) NacpolicyMapOutput {
	return o
}

func (o NacpolicyMapOutput) MapIndex(k pulumi.StringInput) NacpolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Nacpolicy {
		return vs[0].(map[string]*Nacpolicy)[vs[1].(string)]
	}).(NacpolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NacpolicyInput)(nil)).Elem(), &Nacpolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*NacpolicyArrayInput)(nil)).Elem(), NacpolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NacpolicyMapInput)(nil)).Elem(), NacpolicyMap{})
	pulumi.RegisterOutputType(NacpolicyOutput{})
	pulumi.RegisterOutputType(NacpolicyArrayOutput{})
	pulumi.RegisterOutputType(NacpolicyMapOutput{})
}
