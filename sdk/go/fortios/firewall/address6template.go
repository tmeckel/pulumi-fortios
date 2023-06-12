// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firewall

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure IPv6 address templates.
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
//			_, err := firewall.NewAddress6template(ctx, "trname", &firewall.Address6templateArgs{
//				Ip6: pulumi.String("2001:db8:0:b::/64"),
//				SubnetSegments: firewall.Address6templateSubnetSegmentArray{
//					&firewall.Address6templateSubnetSegmentArgs{
//						Bits:      pulumi.Int(4),
//						Exclusive: pulumi.String("disable"),
//						Id:        pulumi.Int(1),
//						Name:      pulumi.String("country"),
//					},
//					&firewall.Address6templateSubnetSegmentArgs{
//						Bits:      pulumi.Int(4),
//						Exclusive: pulumi.String("disable"),
//						Id:        pulumi.Int(2),
//						Name:      pulumi.String("state"),
//					},
//				},
//				SubnetSegmentCount: pulumi.Int(2),
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
// # Firewall Address6Template can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:firewall/address6template:Address6template labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:firewall/address6template:Address6template labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type Address6template struct {
	pulumi.CustomResourceState

	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Security Fabric global object setting. Valid values: `enable`, `disable`.
	FabricObject pulumi.StringOutput `pulumi:"fabricObject"`
	// IPv6 address prefix.
	Ip6 pulumi.StringOutput `pulumi:"ip6"`
	// IPv6 address template name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Number of IPv6 subnet segments.
	SubnetSegmentCount pulumi.IntOutput `pulumi:"subnetSegmentCount"`
	// IPv6 subnet segments. The structure of `subnetSegment` block is documented below.
	SubnetSegments Address6templateSubnetSegmentArrayOutput `pulumi:"subnetSegments"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewAddress6template registers a new resource with the given unique name, arguments, and options.
func NewAddress6template(ctx *pulumi.Context,
	name string, args *Address6templateArgs, opts ...pulumi.ResourceOption) (*Address6template, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Ip6 == nil {
		return nil, errors.New("invalid value for required argument 'Ip6'")
	}
	if args.SubnetSegmentCount == nil {
		return nil, errors.New("invalid value for required argument 'SubnetSegmentCount'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Address6template
	err := ctx.RegisterResource("fortios:firewall/address6template:Address6template", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAddress6template gets an existing Address6template resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAddress6template(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Address6templateState, opts ...pulumi.ResourceOption) (*Address6template, error) {
	var resource Address6template
	err := ctx.ReadResource("fortios:firewall/address6template:Address6template", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Address6template resources.
type address6templateState struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Security Fabric global object setting. Valid values: `enable`, `disable`.
	FabricObject *string `pulumi:"fabricObject"`
	// IPv6 address prefix.
	Ip6 *string `pulumi:"ip6"`
	// IPv6 address template name.
	Name *string `pulumi:"name"`
	// Number of IPv6 subnet segments.
	SubnetSegmentCount *int `pulumi:"subnetSegmentCount"`
	// IPv6 subnet segments. The structure of `subnetSegment` block is documented below.
	SubnetSegments []Address6templateSubnetSegment `pulumi:"subnetSegments"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type Address6templateState struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Security Fabric global object setting. Valid values: `enable`, `disable`.
	FabricObject pulumi.StringPtrInput
	// IPv6 address prefix.
	Ip6 pulumi.StringPtrInput
	// IPv6 address template name.
	Name pulumi.StringPtrInput
	// Number of IPv6 subnet segments.
	SubnetSegmentCount pulumi.IntPtrInput
	// IPv6 subnet segments. The structure of `subnetSegment` block is documented below.
	SubnetSegments Address6templateSubnetSegmentArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (Address6templateState) ElementType() reflect.Type {
	return reflect.TypeOf((*address6templateState)(nil)).Elem()
}

type address6templateArgs struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Security Fabric global object setting. Valid values: `enable`, `disable`.
	FabricObject *string `pulumi:"fabricObject"`
	// IPv6 address prefix.
	Ip6 string `pulumi:"ip6"`
	// IPv6 address template name.
	Name *string `pulumi:"name"`
	// Number of IPv6 subnet segments.
	SubnetSegmentCount int `pulumi:"subnetSegmentCount"`
	// IPv6 subnet segments. The structure of `subnetSegment` block is documented below.
	SubnetSegments []Address6templateSubnetSegment `pulumi:"subnetSegments"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Address6template resource.
type Address6templateArgs struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Security Fabric global object setting. Valid values: `enable`, `disable`.
	FabricObject pulumi.StringPtrInput
	// IPv6 address prefix.
	Ip6 pulumi.StringInput
	// IPv6 address template name.
	Name pulumi.StringPtrInput
	// Number of IPv6 subnet segments.
	SubnetSegmentCount pulumi.IntInput
	// IPv6 subnet segments. The structure of `subnetSegment` block is documented below.
	SubnetSegments Address6templateSubnetSegmentArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (Address6templateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*address6templateArgs)(nil)).Elem()
}

type Address6templateInput interface {
	pulumi.Input

	ToAddress6templateOutput() Address6templateOutput
	ToAddress6templateOutputWithContext(ctx context.Context) Address6templateOutput
}

func (*Address6template) ElementType() reflect.Type {
	return reflect.TypeOf((**Address6template)(nil)).Elem()
}

func (i *Address6template) ToAddress6templateOutput() Address6templateOutput {
	return i.ToAddress6templateOutputWithContext(context.Background())
}

func (i *Address6template) ToAddress6templateOutputWithContext(ctx context.Context) Address6templateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Address6templateOutput)
}

// Address6templateArrayInput is an input type that accepts Address6templateArray and Address6templateArrayOutput values.
// You can construct a concrete instance of `Address6templateArrayInput` via:
//
//	Address6templateArray{ Address6templateArgs{...} }
type Address6templateArrayInput interface {
	pulumi.Input

	ToAddress6templateArrayOutput() Address6templateArrayOutput
	ToAddress6templateArrayOutputWithContext(context.Context) Address6templateArrayOutput
}

type Address6templateArray []Address6templateInput

func (Address6templateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Address6template)(nil)).Elem()
}

func (i Address6templateArray) ToAddress6templateArrayOutput() Address6templateArrayOutput {
	return i.ToAddress6templateArrayOutputWithContext(context.Background())
}

func (i Address6templateArray) ToAddress6templateArrayOutputWithContext(ctx context.Context) Address6templateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Address6templateArrayOutput)
}

// Address6templateMapInput is an input type that accepts Address6templateMap and Address6templateMapOutput values.
// You can construct a concrete instance of `Address6templateMapInput` via:
//
//	Address6templateMap{ "key": Address6templateArgs{...} }
type Address6templateMapInput interface {
	pulumi.Input

	ToAddress6templateMapOutput() Address6templateMapOutput
	ToAddress6templateMapOutputWithContext(context.Context) Address6templateMapOutput
}

type Address6templateMap map[string]Address6templateInput

func (Address6templateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Address6template)(nil)).Elem()
}

func (i Address6templateMap) ToAddress6templateMapOutput() Address6templateMapOutput {
	return i.ToAddress6templateMapOutputWithContext(context.Background())
}

func (i Address6templateMap) ToAddress6templateMapOutputWithContext(ctx context.Context) Address6templateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Address6templateMapOutput)
}

type Address6templateOutput struct{ *pulumi.OutputState }

func (Address6templateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Address6template)(nil)).Elem()
}

func (o Address6templateOutput) ToAddress6templateOutput() Address6templateOutput {
	return o
}

func (o Address6templateOutput) ToAddress6templateOutputWithContext(ctx context.Context) Address6templateOutput {
	return o
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o Address6templateOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Address6template) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Security Fabric global object setting. Valid values: `enable`, `disable`.
func (o Address6templateOutput) FabricObject() pulumi.StringOutput {
	return o.ApplyT(func(v *Address6template) pulumi.StringOutput { return v.FabricObject }).(pulumi.StringOutput)
}

// IPv6 address prefix.
func (o Address6templateOutput) Ip6() pulumi.StringOutput {
	return o.ApplyT(func(v *Address6template) pulumi.StringOutput { return v.Ip6 }).(pulumi.StringOutput)
}

// IPv6 address template name.
func (o Address6templateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Address6template) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Number of IPv6 subnet segments.
func (o Address6templateOutput) SubnetSegmentCount() pulumi.IntOutput {
	return o.ApplyT(func(v *Address6template) pulumi.IntOutput { return v.SubnetSegmentCount }).(pulumi.IntOutput)
}

// IPv6 subnet segments. The structure of `subnetSegment` block is documented below.
func (o Address6templateOutput) SubnetSegments() Address6templateSubnetSegmentArrayOutput {
	return o.ApplyT(func(v *Address6template) Address6templateSubnetSegmentArrayOutput { return v.SubnetSegments }).(Address6templateSubnetSegmentArrayOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o Address6templateOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Address6template) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type Address6templateArrayOutput struct{ *pulumi.OutputState }

func (Address6templateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Address6template)(nil)).Elem()
}

func (o Address6templateArrayOutput) ToAddress6templateArrayOutput() Address6templateArrayOutput {
	return o
}

func (o Address6templateArrayOutput) ToAddress6templateArrayOutputWithContext(ctx context.Context) Address6templateArrayOutput {
	return o
}

func (o Address6templateArrayOutput) Index(i pulumi.IntInput) Address6templateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Address6template {
		return vs[0].([]*Address6template)[vs[1].(int)]
	}).(Address6templateOutput)
}

type Address6templateMapOutput struct{ *pulumi.OutputState }

func (Address6templateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Address6template)(nil)).Elem()
}

func (o Address6templateMapOutput) ToAddress6templateMapOutput() Address6templateMapOutput {
	return o
}

func (o Address6templateMapOutput) ToAddress6templateMapOutputWithContext(ctx context.Context) Address6templateMapOutput {
	return o
}

func (o Address6templateMapOutput) MapIndex(k pulumi.StringInput) Address6templateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Address6template {
		return vs[0].(map[string]*Address6template)[vs[1].(string)]
	}).(Address6templateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*Address6templateInput)(nil)).Elem(), &Address6template{})
	pulumi.RegisterInputType(reflect.TypeOf((*Address6templateArrayInput)(nil)).Elem(), Address6templateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*Address6templateMapInput)(nil)).Elem(), Address6templateMap{})
	pulumi.RegisterOutputType(Address6templateOutput{})
	pulumi.RegisterOutputType(Address6templateArrayOutput{})
	pulumi.RegisterOutputType(Address6templateMapOutput{})
}
