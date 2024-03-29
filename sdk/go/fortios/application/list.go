// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package application

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure application control lists.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/application"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := application.NewList(ctx, "trname", &application.ListArgs{
//				AppReplacemsg:            pulumi.String("enable"),
//				DeepAppInspection:        pulumi.String("enable"),
//				EnforceDefaultAppPort:    pulumi.String("disable"),
//				ExtendedLog:              pulumi.String("disable"),
//				Options:                  pulumi.String("allow-dns"),
//				OtherApplicationAction:   pulumi.String("pass"),
//				OtherApplicationLog:      pulumi.String("disable"),
//				UnknownApplicationAction: pulumi.String("pass"),
//				UnknownApplicationLog:    pulumi.String("disable"),
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
// # Application List can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:application/list:List labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:application/list:List labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type List struct {
	pulumi.CustomResourceState

	// Enable/disable replacement messages for blocked applications. Valid values: `disable`, `enable`.
	AppReplacemsg pulumi.StringOutput `pulumi:"appReplacemsg"`
	// comments
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Enable/disable enforcement of protocols over selected ports. Valid values: `disable`, `enable`.
	ControlDefaultNetworkServices pulumi.StringOutput `pulumi:"controlDefaultNetworkServices"`
	// Enable/disable deep application inspection. Valid values: `disable`, `enable`.
	DeepAppInspection pulumi.StringOutput `pulumi:"deepAppInspection"`
	// Default network service entries. The structure of `defaultNetworkServices` block is documented below.
	DefaultNetworkServices ListDefaultNetworkServiceArrayOutput `pulumi:"defaultNetworkServices"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Enable/disable default application port enforcement for allowed applications. Valid values: `disable`, `enable`.
	EnforceDefaultAppPort pulumi.StringOutput `pulumi:"enforceDefaultAppPort"`
	// Application list entries. The structure of `entries` block is documented below.
	Entries ListEntryArrayOutput `pulumi:"entries"`
	// Enable/disable extended logging. Valid values: `enable`, `disable`.
	ExtendedLog pulumi.StringOutput `pulumi:"extendedLog"`
	// Enable/disable forced inclusion of SSL deep inspection signatures. Valid values: `disable`, `enable`.
	ForceInclusionSslDiSigs pulumi.StringOutput `pulumi:"forceInclusionSslDiSigs"`
	// List name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Basic application protocol signatures allowed by default. Valid values: `allow-dns`, `allow-icmp`, `allow-http`, `allow-ssl`, `allow-quic`.
	Options pulumi.StringOutput `pulumi:"options"`
	// Action for other applications. Valid values: `pass`, `block`.
	OtherApplicationAction pulumi.StringOutput `pulumi:"otherApplicationAction"`
	// Enable/disable logging for other applications. Valid values: `disable`, `enable`.
	OtherApplicationLog pulumi.StringOutput `pulumi:"otherApplicationLog"`
	// P2P applications to be black listed. Valid values: `skype`, `edonkey`, `bittorrent`.
	P2pBlackList pulumi.StringOutput `pulumi:"p2pBlackList"`
	// P2P applications to be blocklisted. Valid values: `skype`, `edonkey`, `bittorrent`.
	P2pBlockList pulumi.StringOutput `pulumi:"p2pBlockList"`
	// Replacement message group.
	ReplacemsgGroup pulumi.StringOutput `pulumi:"replacemsgGroup"`
	// Pass or block traffic from unknown applications. Valid values: `pass`, `block`.
	UnknownApplicationAction pulumi.StringOutput `pulumi:"unknownApplicationAction"`
	// Enable/disable logging for unknown applications. Valid values: `disable`, `enable`.
	UnknownApplicationLog pulumi.StringOutput `pulumi:"unknownApplicationLog"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewList registers a new resource with the given unique name, arguments, and options.
func NewList(ctx *pulumi.Context,
	name string, args *ListArgs, opts ...pulumi.ResourceOption) (*List, error) {
	if args == nil {
		args = &ListArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource List
	err := ctx.RegisterResource("fortios:application/list:List", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetList gets an existing List resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ListState, opts ...pulumi.ResourceOption) (*List, error) {
	var resource List
	err := ctx.ReadResource("fortios:application/list:List", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering List resources.
type listState struct {
	// Enable/disable replacement messages for blocked applications. Valid values: `disable`, `enable`.
	AppReplacemsg *string `pulumi:"appReplacemsg"`
	// comments
	Comment *string `pulumi:"comment"`
	// Enable/disable enforcement of protocols over selected ports. Valid values: `disable`, `enable`.
	ControlDefaultNetworkServices *string `pulumi:"controlDefaultNetworkServices"`
	// Enable/disable deep application inspection. Valid values: `disable`, `enable`.
	DeepAppInspection *string `pulumi:"deepAppInspection"`
	// Default network service entries. The structure of `defaultNetworkServices` block is documented below.
	DefaultNetworkServices []ListDefaultNetworkService `pulumi:"defaultNetworkServices"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Enable/disable default application port enforcement for allowed applications. Valid values: `disable`, `enable`.
	EnforceDefaultAppPort *string `pulumi:"enforceDefaultAppPort"`
	// Application list entries. The structure of `entries` block is documented below.
	Entries []ListEntry `pulumi:"entries"`
	// Enable/disable extended logging. Valid values: `enable`, `disable`.
	ExtendedLog *string `pulumi:"extendedLog"`
	// Enable/disable forced inclusion of SSL deep inspection signatures. Valid values: `disable`, `enable`.
	ForceInclusionSslDiSigs *string `pulumi:"forceInclusionSslDiSigs"`
	// List name.
	Name *string `pulumi:"name"`
	// Basic application protocol signatures allowed by default. Valid values: `allow-dns`, `allow-icmp`, `allow-http`, `allow-ssl`, `allow-quic`.
	Options *string `pulumi:"options"`
	// Action for other applications. Valid values: `pass`, `block`.
	OtherApplicationAction *string `pulumi:"otherApplicationAction"`
	// Enable/disable logging for other applications. Valid values: `disable`, `enable`.
	OtherApplicationLog *string `pulumi:"otherApplicationLog"`
	// P2P applications to be black listed. Valid values: `skype`, `edonkey`, `bittorrent`.
	P2pBlackList *string `pulumi:"p2pBlackList"`
	// P2P applications to be blocklisted. Valid values: `skype`, `edonkey`, `bittorrent`.
	P2pBlockList *string `pulumi:"p2pBlockList"`
	// Replacement message group.
	ReplacemsgGroup *string `pulumi:"replacemsgGroup"`
	// Pass or block traffic from unknown applications. Valid values: `pass`, `block`.
	UnknownApplicationAction *string `pulumi:"unknownApplicationAction"`
	// Enable/disable logging for unknown applications. Valid values: `disable`, `enable`.
	UnknownApplicationLog *string `pulumi:"unknownApplicationLog"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type ListState struct {
	// Enable/disable replacement messages for blocked applications. Valid values: `disable`, `enable`.
	AppReplacemsg pulumi.StringPtrInput
	// comments
	Comment pulumi.StringPtrInput
	// Enable/disable enforcement of protocols over selected ports. Valid values: `disable`, `enable`.
	ControlDefaultNetworkServices pulumi.StringPtrInput
	// Enable/disable deep application inspection. Valid values: `disable`, `enable`.
	DeepAppInspection pulumi.StringPtrInput
	// Default network service entries. The structure of `defaultNetworkServices` block is documented below.
	DefaultNetworkServices ListDefaultNetworkServiceArrayInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Enable/disable default application port enforcement for allowed applications. Valid values: `disable`, `enable`.
	EnforceDefaultAppPort pulumi.StringPtrInput
	// Application list entries. The structure of `entries` block is documented below.
	Entries ListEntryArrayInput
	// Enable/disable extended logging. Valid values: `enable`, `disable`.
	ExtendedLog pulumi.StringPtrInput
	// Enable/disable forced inclusion of SSL deep inspection signatures. Valid values: `disable`, `enable`.
	ForceInclusionSslDiSigs pulumi.StringPtrInput
	// List name.
	Name pulumi.StringPtrInput
	// Basic application protocol signatures allowed by default. Valid values: `allow-dns`, `allow-icmp`, `allow-http`, `allow-ssl`, `allow-quic`.
	Options pulumi.StringPtrInput
	// Action for other applications. Valid values: `pass`, `block`.
	OtherApplicationAction pulumi.StringPtrInput
	// Enable/disable logging for other applications. Valid values: `disable`, `enable`.
	OtherApplicationLog pulumi.StringPtrInput
	// P2P applications to be black listed. Valid values: `skype`, `edonkey`, `bittorrent`.
	P2pBlackList pulumi.StringPtrInput
	// P2P applications to be blocklisted. Valid values: `skype`, `edonkey`, `bittorrent`.
	P2pBlockList pulumi.StringPtrInput
	// Replacement message group.
	ReplacemsgGroup pulumi.StringPtrInput
	// Pass or block traffic from unknown applications. Valid values: `pass`, `block`.
	UnknownApplicationAction pulumi.StringPtrInput
	// Enable/disable logging for unknown applications. Valid values: `disable`, `enable`.
	UnknownApplicationLog pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (ListState) ElementType() reflect.Type {
	return reflect.TypeOf((*listState)(nil)).Elem()
}

type listArgs struct {
	// Enable/disable replacement messages for blocked applications. Valid values: `disable`, `enable`.
	AppReplacemsg *string `pulumi:"appReplacemsg"`
	// comments
	Comment *string `pulumi:"comment"`
	// Enable/disable enforcement of protocols over selected ports. Valid values: `disable`, `enable`.
	ControlDefaultNetworkServices *string `pulumi:"controlDefaultNetworkServices"`
	// Enable/disable deep application inspection. Valid values: `disable`, `enable`.
	DeepAppInspection *string `pulumi:"deepAppInspection"`
	// Default network service entries. The structure of `defaultNetworkServices` block is documented below.
	DefaultNetworkServices []ListDefaultNetworkService `pulumi:"defaultNetworkServices"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Enable/disable default application port enforcement for allowed applications. Valid values: `disable`, `enable`.
	EnforceDefaultAppPort *string `pulumi:"enforceDefaultAppPort"`
	// Application list entries. The structure of `entries` block is documented below.
	Entries []ListEntry `pulumi:"entries"`
	// Enable/disable extended logging. Valid values: `enable`, `disable`.
	ExtendedLog *string `pulumi:"extendedLog"`
	// Enable/disable forced inclusion of SSL deep inspection signatures. Valid values: `disable`, `enable`.
	ForceInclusionSslDiSigs *string `pulumi:"forceInclusionSslDiSigs"`
	// List name.
	Name *string `pulumi:"name"`
	// Basic application protocol signatures allowed by default. Valid values: `allow-dns`, `allow-icmp`, `allow-http`, `allow-ssl`, `allow-quic`.
	Options *string `pulumi:"options"`
	// Action for other applications. Valid values: `pass`, `block`.
	OtherApplicationAction *string `pulumi:"otherApplicationAction"`
	// Enable/disable logging for other applications. Valid values: `disable`, `enable`.
	OtherApplicationLog *string `pulumi:"otherApplicationLog"`
	// P2P applications to be black listed. Valid values: `skype`, `edonkey`, `bittorrent`.
	P2pBlackList *string `pulumi:"p2pBlackList"`
	// P2P applications to be blocklisted. Valid values: `skype`, `edonkey`, `bittorrent`.
	P2pBlockList *string `pulumi:"p2pBlockList"`
	// Replacement message group.
	ReplacemsgGroup *string `pulumi:"replacemsgGroup"`
	// Pass or block traffic from unknown applications. Valid values: `pass`, `block`.
	UnknownApplicationAction *string `pulumi:"unknownApplicationAction"`
	// Enable/disable logging for unknown applications. Valid values: `disable`, `enable`.
	UnknownApplicationLog *string `pulumi:"unknownApplicationLog"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a List resource.
type ListArgs struct {
	// Enable/disable replacement messages for blocked applications. Valid values: `disable`, `enable`.
	AppReplacemsg pulumi.StringPtrInput
	// comments
	Comment pulumi.StringPtrInput
	// Enable/disable enforcement of protocols over selected ports. Valid values: `disable`, `enable`.
	ControlDefaultNetworkServices pulumi.StringPtrInput
	// Enable/disable deep application inspection. Valid values: `disable`, `enable`.
	DeepAppInspection pulumi.StringPtrInput
	// Default network service entries. The structure of `defaultNetworkServices` block is documented below.
	DefaultNetworkServices ListDefaultNetworkServiceArrayInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Enable/disable default application port enforcement for allowed applications. Valid values: `disable`, `enable`.
	EnforceDefaultAppPort pulumi.StringPtrInput
	// Application list entries. The structure of `entries` block is documented below.
	Entries ListEntryArrayInput
	// Enable/disable extended logging. Valid values: `enable`, `disable`.
	ExtendedLog pulumi.StringPtrInput
	// Enable/disable forced inclusion of SSL deep inspection signatures. Valid values: `disable`, `enable`.
	ForceInclusionSslDiSigs pulumi.StringPtrInput
	// List name.
	Name pulumi.StringPtrInput
	// Basic application protocol signatures allowed by default. Valid values: `allow-dns`, `allow-icmp`, `allow-http`, `allow-ssl`, `allow-quic`.
	Options pulumi.StringPtrInput
	// Action for other applications. Valid values: `pass`, `block`.
	OtherApplicationAction pulumi.StringPtrInput
	// Enable/disable logging for other applications. Valid values: `disable`, `enable`.
	OtherApplicationLog pulumi.StringPtrInput
	// P2P applications to be black listed. Valid values: `skype`, `edonkey`, `bittorrent`.
	P2pBlackList pulumi.StringPtrInput
	// P2P applications to be blocklisted. Valid values: `skype`, `edonkey`, `bittorrent`.
	P2pBlockList pulumi.StringPtrInput
	// Replacement message group.
	ReplacemsgGroup pulumi.StringPtrInput
	// Pass or block traffic from unknown applications. Valid values: `pass`, `block`.
	UnknownApplicationAction pulumi.StringPtrInput
	// Enable/disable logging for unknown applications. Valid values: `disable`, `enable`.
	UnknownApplicationLog pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (ListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*listArgs)(nil)).Elem()
}

type ListInput interface {
	pulumi.Input

	ToListOutput() ListOutput
	ToListOutputWithContext(ctx context.Context) ListOutput
}

func (*List) ElementType() reflect.Type {
	return reflect.TypeOf((**List)(nil)).Elem()
}

func (i *List) ToListOutput() ListOutput {
	return i.ToListOutputWithContext(context.Background())
}

func (i *List) ToListOutputWithContext(ctx context.Context) ListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ListOutput)
}

// ListArrayInput is an input type that accepts ListArray and ListArrayOutput values.
// You can construct a concrete instance of `ListArrayInput` via:
//
//	ListArray{ ListArgs{...} }
type ListArrayInput interface {
	pulumi.Input

	ToListArrayOutput() ListArrayOutput
	ToListArrayOutputWithContext(context.Context) ListArrayOutput
}

type ListArray []ListInput

func (ListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*List)(nil)).Elem()
}

func (i ListArray) ToListArrayOutput() ListArrayOutput {
	return i.ToListArrayOutputWithContext(context.Background())
}

func (i ListArray) ToListArrayOutputWithContext(ctx context.Context) ListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ListArrayOutput)
}

// ListMapInput is an input type that accepts ListMap and ListMapOutput values.
// You can construct a concrete instance of `ListMapInput` via:
//
//	ListMap{ "key": ListArgs{...} }
type ListMapInput interface {
	pulumi.Input

	ToListMapOutput() ListMapOutput
	ToListMapOutputWithContext(context.Context) ListMapOutput
}

type ListMap map[string]ListInput

func (ListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*List)(nil)).Elem()
}

func (i ListMap) ToListMapOutput() ListMapOutput {
	return i.ToListMapOutputWithContext(context.Background())
}

func (i ListMap) ToListMapOutputWithContext(ctx context.Context) ListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ListMapOutput)
}

type ListOutput struct{ *pulumi.OutputState }

func (ListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**List)(nil)).Elem()
}

func (o ListOutput) ToListOutput() ListOutput {
	return o
}

func (o ListOutput) ToListOutputWithContext(ctx context.Context) ListOutput {
	return o
}

// Enable/disable replacement messages for blocked applications. Valid values: `disable`, `enable`.
func (o ListOutput) AppReplacemsg() pulumi.StringOutput {
	return o.ApplyT(func(v *List) pulumi.StringOutput { return v.AppReplacemsg }).(pulumi.StringOutput)
}

// comments
func (o ListOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *List) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// Enable/disable enforcement of protocols over selected ports. Valid values: `disable`, `enable`.
func (o ListOutput) ControlDefaultNetworkServices() pulumi.StringOutput {
	return o.ApplyT(func(v *List) pulumi.StringOutput { return v.ControlDefaultNetworkServices }).(pulumi.StringOutput)
}

// Enable/disable deep application inspection. Valid values: `disable`, `enable`.
func (o ListOutput) DeepAppInspection() pulumi.StringOutput {
	return o.ApplyT(func(v *List) pulumi.StringOutput { return v.DeepAppInspection }).(pulumi.StringOutput)
}

// Default network service entries. The structure of `defaultNetworkServices` block is documented below.
func (o ListOutput) DefaultNetworkServices() ListDefaultNetworkServiceArrayOutput {
	return o.ApplyT(func(v *List) ListDefaultNetworkServiceArrayOutput { return v.DefaultNetworkServices }).(ListDefaultNetworkServiceArrayOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o ListOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *List) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Enable/disable default application port enforcement for allowed applications. Valid values: `disable`, `enable`.
func (o ListOutput) EnforceDefaultAppPort() pulumi.StringOutput {
	return o.ApplyT(func(v *List) pulumi.StringOutput { return v.EnforceDefaultAppPort }).(pulumi.StringOutput)
}

// Application list entries. The structure of `entries` block is documented below.
func (o ListOutput) Entries() ListEntryArrayOutput {
	return o.ApplyT(func(v *List) ListEntryArrayOutput { return v.Entries }).(ListEntryArrayOutput)
}

// Enable/disable extended logging. Valid values: `enable`, `disable`.
func (o ListOutput) ExtendedLog() pulumi.StringOutput {
	return o.ApplyT(func(v *List) pulumi.StringOutput { return v.ExtendedLog }).(pulumi.StringOutput)
}

// Enable/disable forced inclusion of SSL deep inspection signatures. Valid values: `disable`, `enable`.
func (o ListOutput) ForceInclusionSslDiSigs() pulumi.StringOutput {
	return o.ApplyT(func(v *List) pulumi.StringOutput { return v.ForceInclusionSslDiSigs }).(pulumi.StringOutput)
}

// List name.
func (o ListOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *List) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Basic application protocol signatures allowed by default. Valid values: `allow-dns`, `allow-icmp`, `allow-http`, `allow-ssl`, `allow-quic`.
func (o ListOutput) Options() pulumi.StringOutput {
	return o.ApplyT(func(v *List) pulumi.StringOutput { return v.Options }).(pulumi.StringOutput)
}

// Action for other applications. Valid values: `pass`, `block`.
func (o ListOutput) OtherApplicationAction() pulumi.StringOutput {
	return o.ApplyT(func(v *List) pulumi.StringOutput { return v.OtherApplicationAction }).(pulumi.StringOutput)
}

// Enable/disable logging for other applications. Valid values: `disable`, `enable`.
func (o ListOutput) OtherApplicationLog() pulumi.StringOutput {
	return o.ApplyT(func(v *List) pulumi.StringOutput { return v.OtherApplicationLog }).(pulumi.StringOutput)
}

// P2P applications to be black listed. Valid values: `skype`, `edonkey`, `bittorrent`.
func (o ListOutput) P2pBlackList() pulumi.StringOutput {
	return o.ApplyT(func(v *List) pulumi.StringOutput { return v.P2pBlackList }).(pulumi.StringOutput)
}

// P2P applications to be blocklisted. Valid values: `skype`, `edonkey`, `bittorrent`.
func (o ListOutput) P2pBlockList() pulumi.StringOutput {
	return o.ApplyT(func(v *List) pulumi.StringOutput { return v.P2pBlockList }).(pulumi.StringOutput)
}

// Replacement message group.
func (o ListOutput) ReplacemsgGroup() pulumi.StringOutput {
	return o.ApplyT(func(v *List) pulumi.StringOutput { return v.ReplacemsgGroup }).(pulumi.StringOutput)
}

// Pass or block traffic from unknown applications. Valid values: `pass`, `block`.
func (o ListOutput) UnknownApplicationAction() pulumi.StringOutput {
	return o.ApplyT(func(v *List) pulumi.StringOutput { return v.UnknownApplicationAction }).(pulumi.StringOutput)
}

// Enable/disable logging for unknown applications. Valid values: `disable`, `enable`.
func (o ListOutput) UnknownApplicationLog() pulumi.StringOutput {
	return o.ApplyT(func(v *List) pulumi.StringOutput { return v.UnknownApplicationLog }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o ListOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *List) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type ListArrayOutput struct{ *pulumi.OutputState }

func (ListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*List)(nil)).Elem()
}

func (o ListArrayOutput) ToListArrayOutput() ListArrayOutput {
	return o
}

func (o ListArrayOutput) ToListArrayOutputWithContext(ctx context.Context) ListArrayOutput {
	return o
}

func (o ListArrayOutput) Index(i pulumi.IntInput) ListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *List {
		return vs[0].([]*List)[vs[1].(int)]
	}).(ListOutput)
}

type ListMapOutput struct{ *pulumi.OutputState }

func (ListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*List)(nil)).Elem()
}

func (o ListMapOutput) ToListMapOutput() ListMapOutput {
	return o
}

func (o ListMapOutput) ToListMapOutputWithContext(ctx context.Context) ListMapOutput {
	return o
}

func (o ListMapOutput) MapIndex(k pulumi.StringInput) ListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *List {
		return vs[0].(map[string]*List)[vs[1].(string)]
	}).(ListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ListInput)(nil)).Elem(), &List{})
	pulumi.RegisterInputType(reflect.TypeOf((*ListArrayInput)(nil)).Elem(), ListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ListMapInput)(nil)).Elem(), ListMap{})
	pulumi.RegisterOutputType(ListOutput{})
	pulumi.RegisterOutputType(ListArrayOutput{})
	pulumi.RegisterOutputType(ListMapOutput{})
}
