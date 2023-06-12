// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firewall

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure IPv4 virtual IP groups.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/firewall"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			trname1, err := firewall.NewVip(ctx, "trname1", &firewall.VipArgs{
//				Extintf: pulumi.String("any"),
//				Extport: pulumi.String("0-65535"),
//				Extip:   pulumi.String("2.0.0.1-2.0.0.4"),
//				Mappedips: firewall.VipMappedipArray{
//					&firewall.VipMappedipArgs{
//						Range: pulumi.String("3.0.0.0-3.0.0.3"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = firewall.NewVipgrp(ctx, "trname", &firewall.VipgrpArgs{
//				Color:     pulumi.Int(0),
//				Interface: pulumi.String("any"),
//				Members: firewall.VipgrpMemberArray{
//					&firewall.VipgrpMemberArgs{
//						Name: trname1.Name,
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
// # Firewall Vipgrp can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:firewall/vipgrp:Vipgrp labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:firewall/vipgrp:Vipgrp labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type Vipgrp struct {
	pulumi.CustomResourceState

	// Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
	Color pulumi.IntOutput `pulumi:"color"`
	// Comment.
	Comments pulumi.StringPtrOutput `pulumi:"comments"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// interface
	Interface pulumi.StringOutput `pulumi:"interface"`
	// Member VIP objects of the group (Separate multiple objects with a space). The structure of `member` block is documented below.
	Members VipgrpMemberArrayOutput `pulumi:"members"`
	// VIP group name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid pulumi.StringOutput `pulumi:"uuid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewVipgrp registers a new resource with the given unique name, arguments, and options.
func NewVipgrp(ctx *pulumi.Context,
	name string, args *VipgrpArgs, opts ...pulumi.ResourceOption) (*Vipgrp, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Interface == nil {
		return nil, errors.New("invalid value for required argument 'Interface'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Vipgrp
	err := ctx.RegisterResource("fortios:firewall/vipgrp:Vipgrp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVipgrp gets an existing Vipgrp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVipgrp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VipgrpState, opts ...pulumi.ResourceOption) (*Vipgrp, error) {
	var resource Vipgrp
	err := ctx.ReadResource("fortios:firewall/vipgrp:Vipgrp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Vipgrp resources.
type vipgrpState struct {
	// Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
	Color *int `pulumi:"color"`
	// Comment.
	Comments *string `pulumi:"comments"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// interface
	Interface *string `pulumi:"interface"`
	// Member VIP objects of the group (Separate multiple objects with a space). The structure of `member` block is documented below.
	Members []VipgrpMember `pulumi:"members"`
	// VIP group name.
	Name *string `pulumi:"name"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid *string `pulumi:"uuid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type VipgrpState struct {
	// Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
	Color pulumi.IntPtrInput
	// Comment.
	Comments pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// interface
	Interface pulumi.StringPtrInput
	// Member VIP objects of the group (Separate multiple objects with a space). The structure of `member` block is documented below.
	Members VipgrpMemberArrayInput
	// VIP group name.
	Name pulumi.StringPtrInput
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (VipgrpState) ElementType() reflect.Type {
	return reflect.TypeOf((*vipgrpState)(nil)).Elem()
}

type vipgrpArgs struct {
	// Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
	Color *int `pulumi:"color"`
	// Comment.
	Comments *string `pulumi:"comments"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// interface
	Interface string `pulumi:"interface"`
	// Member VIP objects of the group (Separate multiple objects with a space). The structure of `member` block is documented below.
	Members []VipgrpMember `pulumi:"members"`
	// VIP group name.
	Name *string `pulumi:"name"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid *string `pulumi:"uuid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Vipgrp resource.
type VipgrpArgs struct {
	// Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
	Color pulumi.IntPtrInput
	// Comment.
	Comments pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// interface
	Interface pulumi.StringInput
	// Member VIP objects of the group (Separate multiple objects with a space). The structure of `member` block is documented below.
	Members VipgrpMemberArrayInput
	// VIP group name.
	Name pulumi.StringPtrInput
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (VipgrpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vipgrpArgs)(nil)).Elem()
}

type VipgrpInput interface {
	pulumi.Input

	ToVipgrpOutput() VipgrpOutput
	ToVipgrpOutputWithContext(ctx context.Context) VipgrpOutput
}

func (*Vipgrp) ElementType() reflect.Type {
	return reflect.TypeOf((**Vipgrp)(nil)).Elem()
}

func (i *Vipgrp) ToVipgrpOutput() VipgrpOutput {
	return i.ToVipgrpOutputWithContext(context.Background())
}

func (i *Vipgrp) ToVipgrpOutputWithContext(ctx context.Context) VipgrpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VipgrpOutput)
}

// VipgrpArrayInput is an input type that accepts VipgrpArray and VipgrpArrayOutput values.
// You can construct a concrete instance of `VipgrpArrayInput` via:
//
//	VipgrpArray{ VipgrpArgs{...} }
type VipgrpArrayInput interface {
	pulumi.Input

	ToVipgrpArrayOutput() VipgrpArrayOutput
	ToVipgrpArrayOutputWithContext(context.Context) VipgrpArrayOutput
}

type VipgrpArray []VipgrpInput

func (VipgrpArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Vipgrp)(nil)).Elem()
}

func (i VipgrpArray) ToVipgrpArrayOutput() VipgrpArrayOutput {
	return i.ToVipgrpArrayOutputWithContext(context.Background())
}

func (i VipgrpArray) ToVipgrpArrayOutputWithContext(ctx context.Context) VipgrpArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VipgrpArrayOutput)
}

// VipgrpMapInput is an input type that accepts VipgrpMap and VipgrpMapOutput values.
// You can construct a concrete instance of `VipgrpMapInput` via:
//
//	VipgrpMap{ "key": VipgrpArgs{...} }
type VipgrpMapInput interface {
	pulumi.Input

	ToVipgrpMapOutput() VipgrpMapOutput
	ToVipgrpMapOutputWithContext(context.Context) VipgrpMapOutput
}

type VipgrpMap map[string]VipgrpInput

func (VipgrpMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Vipgrp)(nil)).Elem()
}

func (i VipgrpMap) ToVipgrpMapOutput() VipgrpMapOutput {
	return i.ToVipgrpMapOutputWithContext(context.Background())
}

func (i VipgrpMap) ToVipgrpMapOutputWithContext(ctx context.Context) VipgrpMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VipgrpMapOutput)
}

type VipgrpOutput struct{ *pulumi.OutputState }

func (VipgrpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Vipgrp)(nil)).Elem()
}

func (o VipgrpOutput) ToVipgrpOutput() VipgrpOutput {
	return o
}

func (o VipgrpOutput) ToVipgrpOutputWithContext(ctx context.Context) VipgrpOutput {
	return o
}

// Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
func (o VipgrpOutput) Color() pulumi.IntOutput {
	return o.ApplyT(func(v *Vipgrp) pulumi.IntOutput { return v.Color }).(pulumi.IntOutput)
}

// Comment.
func (o VipgrpOutput) Comments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Vipgrp) pulumi.StringPtrOutput { return v.Comments }).(pulumi.StringPtrOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o VipgrpOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Vipgrp) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// interface
func (o VipgrpOutput) Interface() pulumi.StringOutput {
	return o.ApplyT(func(v *Vipgrp) pulumi.StringOutput { return v.Interface }).(pulumi.StringOutput)
}

// Member VIP objects of the group (Separate multiple objects with a space). The structure of `member` block is documented below.
func (o VipgrpOutput) Members() VipgrpMemberArrayOutput {
	return o.ApplyT(func(v *Vipgrp) VipgrpMemberArrayOutput { return v.Members }).(VipgrpMemberArrayOutput)
}

// VIP group name.
func (o VipgrpOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Vipgrp) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
func (o VipgrpOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v *Vipgrp) pulumi.StringOutput { return v.Uuid }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o VipgrpOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Vipgrp) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type VipgrpArrayOutput struct{ *pulumi.OutputState }

func (VipgrpArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Vipgrp)(nil)).Elem()
}

func (o VipgrpArrayOutput) ToVipgrpArrayOutput() VipgrpArrayOutput {
	return o
}

func (o VipgrpArrayOutput) ToVipgrpArrayOutputWithContext(ctx context.Context) VipgrpArrayOutput {
	return o
}

func (o VipgrpArrayOutput) Index(i pulumi.IntInput) VipgrpOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Vipgrp {
		return vs[0].([]*Vipgrp)[vs[1].(int)]
	}).(VipgrpOutput)
}

type VipgrpMapOutput struct{ *pulumi.OutputState }

func (VipgrpMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Vipgrp)(nil)).Elem()
}

func (o VipgrpMapOutput) ToVipgrpMapOutput() VipgrpMapOutput {
	return o
}

func (o VipgrpMapOutput) ToVipgrpMapOutputWithContext(ctx context.Context) VipgrpMapOutput {
	return o
}

func (o VipgrpMapOutput) MapIndex(k pulumi.StringInput) VipgrpOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Vipgrp {
		return vs[0].(map[string]*Vipgrp)[vs[1].(string)]
	}).(VipgrpOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VipgrpInput)(nil)).Elem(), &Vipgrp{})
	pulumi.RegisterInputType(reflect.TypeOf((*VipgrpArrayInput)(nil)).Elem(), VipgrpArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VipgrpMapInput)(nil)).Elem(), VipgrpMap{})
	pulumi.RegisterOutputType(VipgrpOutput{})
	pulumi.RegisterOutputType(VipgrpArrayOutput{})
	pulumi.RegisterOutputType(VipgrpMapOutput{})
}
