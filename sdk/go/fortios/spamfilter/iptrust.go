// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package spamfilter

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure AntiSpam IP trust. Applies to FortiOS Version `<= 6.2.0`.
//
// ## Import
//
// # Spamfilter Iptrust can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:spamfilter/iptrust:Iptrust labelname {{fosid}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:spamfilter/iptrust:Iptrust labelname {{fosid}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type Iptrust struct {
	pulumi.CustomResourceState

	// Optional comments.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Spam filter trusted IP addresses. The structure of `entries` block is documented below.
	Entries IptrustEntryArrayOutput `pulumi:"entries"`
	// ID.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// Name of table.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewIptrust registers a new resource with the given unique name, arguments, and options.
func NewIptrust(ctx *pulumi.Context,
	name string, args *IptrustArgs, opts ...pulumi.ResourceOption) (*Iptrust, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Fosid == nil {
		return nil, errors.New("invalid value for required argument 'Fosid'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Iptrust
	err := ctx.RegisterResource("fortios:spamfilter/iptrust:Iptrust", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIptrust gets an existing Iptrust resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIptrust(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IptrustState, opts ...pulumi.ResourceOption) (*Iptrust, error) {
	var resource Iptrust
	err := ctx.ReadResource("fortios:spamfilter/iptrust:Iptrust", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Iptrust resources.
type iptrustState struct {
	// Optional comments.
	Comment *string `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Spam filter trusted IP addresses. The structure of `entries` block is documented below.
	Entries []IptrustEntry `pulumi:"entries"`
	// ID.
	Fosid *int `pulumi:"fosid"`
	// Name of table.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type IptrustState struct {
	// Optional comments.
	Comment pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Spam filter trusted IP addresses. The structure of `entries` block is documented below.
	Entries IptrustEntryArrayInput
	// ID.
	Fosid pulumi.IntPtrInput
	// Name of table.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (IptrustState) ElementType() reflect.Type {
	return reflect.TypeOf((*iptrustState)(nil)).Elem()
}

type iptrustArgs struct {
	// Optional comments.
	Comment *string `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Spam filter trusted IP addresses. The structure of `entries` block is documented below.
	Entries []IptrustEntry `pulumi:"entries"`
	// ID.
	Fosid int `pulumi:"fosid"`
	// Name of table.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Iptrust resource.
type IptrustArgs struct {
	// Optional comments.
	Comment pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Spam filter trusted IP addresses. The structure of `entries` block is documented below.
	Entries IptrustEntryArrayInput
	// ID.
	Fosid pulumi.IntInput
	// Name of table.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (IptrustArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iptrustArgs)(nil)).Elem()
}

type IptrustInput interface {
	pulumi.Input

	ToIptrustOutput() IptrustOutput
	ToIptrustOutputWithContext(ctx context.Context) IptrustOutput
}

func (*Iptrust) ElementType() reflect.Type {
	return reflect.TypeOf((**Iptrust)(nil)).Elem()
}

func (i *Iptrust) ToIptrustOutput() IptrustOutput {
	return i.ToIptrustOutputWithContext(context.Background())
}

func (i *Iptrust) ToIptrustOutputWithContext(ctx context.Context) IptrustOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IptrustOutput)
}

// IptrustArrayInput is an input type that accepts IptrustArray and IptrustArrayOutput values.
// You can construct a concrete instance of `IptrustArrayInput` via:
//
//	IptrustArray{ IptrustArgs{...} }
type IptrustArrayInput interface {
	pulumi.Input

	ToIptrustArrayOutput() IptrustArrayOutput
	ToIptrustArrayOutputWithContext(context.Context) IptrustArrayOutput
}

type IptrustArray []IptrustInput

func (IptrustArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Iptrust)(nil)).Elem()
}

func (i IptrustArray) ToIptrustArrayOutput() IptrustArrayOutput {
	return i.ToIptrustArrayOutputWithContext(context.Background())
}

func (i IptrustArray) ToIptrustArrayOutputWithContext(ctx context.Context) IptrustArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IptrustArrayOutput)
}

// IptrustMapInput is an input type that accepts IptrustMap and IptrustMapOutput values.
// You can construct a concrete instance of `IptrustMapInput` via:
//
//	IptrustMap{ "key": IptrustArgs{...} }
type IptrustMapInput interface {
	pulumi.Input

	ToIptrustMapOutput() IptrustMapOutput
	ToIptrustMapOutputWithContext(context.Context) IptrustMapOutput
}

type IptrustMap map[string]IptrustInput

func (IptrustMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Iptrust)(nil)).Elem()
}

func (i IptrustMap) ToIptrustMapOutput() IptrustMapOutput {
	return i.ToIptrustMapOutputWithContext(context.Background())
}

func (i IptrustMap) ToIptrustMapOutputWithContext(ctx context.Context) IptrustMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IptrustMapOutput)
}

type IptrustOutput struct{ *pulumi.OutputState }

func (IptrustOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Iptrust)(nil)).Elem()
}

func (o IptrustOutput) ToIptrustOutput() IptrustOutput {
	return o
}

func (o IptrustOutput) ToIptrustOutputWithContext(ctx context.Context) IptrustOutput {
	return o
}

// Optional comments.
func (o IptrustOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Iptrust) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o IptrustOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Iptrust) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Spam filter trusted IP addresses. The structure of `entries` block is documented below.
func (o IptrustOutput) Entries() IptrustEntryArrayOutput {
	return o.ApplyT(func(v *Iptrust) IptrustEntryArrayOutput { return v.Entries }).(IptrustEntryArrayOutput)
}

// ID.
func (o IptrustOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *Iptrust) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

// Name of table.
func (o IptrustOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Iptrust) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o IptrustOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Iptrust) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type IptrustArrayOutput struct{ *pulumi.OutputState }

func (IptrustArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Iptrust)(nil)).Elem()
}

func (o IptrustArrayOutput) ToIptrustArrayOutput() IptrustArrayOutput {
	return o
}

func (o IptrustArrayOutput) ToIptrustArrayOutputWithContext(ctx context.Context) IptrustArrayOutput {
	return o
}

func (o IptrustArrayOutput) Index(i pulumi.IntInput) IptrustOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Iptrust {
		return vs[0].([]*Iptrust)[vs[1].(int)]
	}).(IptrustOutput)
}

type IptrustMapOutput struct{ *pulumi.OutputState }

func (IptrustMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Iptrust)(nil)).Elem()
}

func (o IptrustMapOutput) ToIptrustMapOutput() IptrustMapOutput {
	return o
}

func (o IptrustMapOutput) ToIptrustMapOutputWithContext(ctx context.Context) IptrustMapOutput {
	return o
}

func (o IptrustMapOutput) MapIndex(k pulumi.StringInput) IptrustOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Iptrust {
		return vs[0].(map[string]*Iptrust)[vs[1].(string)]
	}).(IptrustOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IptrustInput)(nil)).Elem(), &Iptrust{})
	pulumi.RegisterInputType(reflect.TypeOf((*IptrustArrayInput)(nil)).Elem(), IptrustArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IptrustMapInput)(nil)).Elem(), IptrustMap{})
	pulumi.RegisterOutputType(IptrustOutput{})
	pulumi.RegisterOutputType(IptrustArrayOutput{})
	pulumi.RegisterOutputType(IptrustMapOutput{})
}