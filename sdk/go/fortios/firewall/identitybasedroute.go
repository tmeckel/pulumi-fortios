// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firewall

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure identity based routing.
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
//			_, err := firewall.NewIdentitybasedroute(ctx, "trname", &firewall.IdentitybasedrouteArgs{
//				Comments: pulumi.String("test"),
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
// # Firewall IdentityBasedRoute can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:firewall/identitybasedroute:Identitybasedroute labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:firewall/identitybasedroute:Identitybasedroute labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type Identitybasedroute struct {
	pulumi.CustomResourceState

	// Comments.
	Comments pulumi.StringOutput `pulumi:"comments"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Rule. The structure of `rule` block is documented below.
	Rules IdentitybasedrouteRuleArrayOutput `pulumi:"rules"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewIdentitybasedroute registers a new resource with the given unique name, arguments, and options.
func NewIdentitybasedroute(ctx *pulumi.Context,
	name string, args *IdentitybasedrouteArgs, opts ...pulumi.ResourceOption) (*Identitybasedroute, error) {
	if args == nil {
		args = &IdentitybasedrouteArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource Identitybasedroute
	err := ctx.RegisterResource("fortios:firewall/identitybasedroute:Identitybasedroute", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIdentitybasedroute gets an existing Identitybasedroute resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIdentitybasedroute(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IdentitybasedrouteState, opts ...pulumi.ResourceOption) (*Identitybasedroute, error) {
	var resource Identitybasedroute
	err := ctx.ReadResource("fortios:firewall/identitybasedroute:Identitybasedroute", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Identitybasedroute resources.
type identitybasedrouteState struct {
	// Comments.
	Comments *string `pulumi:"comments"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Name.
	Name *string `pulumi:"name"`
	// Rule. The structure of `rule` block is documented below.
	Rules []IdentitybasedrouteRule `pulumi:"rules"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type IdentitybasedrouteState struct {
	// Comments.
	Comments pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Name.
	Name pulumi.StringPtrInput
	// Rule. The structure of `rule` block is documented below.
	Rules IdentitybasedrouteRuleArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (IdentitybasedrouteState) ElementType() reflect.Type {
	return reflect.TypeOf((*identitybasedrouteState)(nil)).Elem()
}

type identitybasedrouteArgs struct {
	// Comments.
	Comments *string `pulumi:"comments"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Name.
	Name *string `pulumi:"name"`
	// Rule. The structure of `rule` block is documented below.
	Rules []IdentitybasedrouteRule `pulumi:"rules"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Identitybasedroute resource.
type IdentitybasedrouteArgs struct {
	// Comments.
	Comments pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Name.
	Name pulumi.StringPtrInput
	// Rule. The structure of `rule` block is documented below.
	Rules IdentitybasedrouteRuleArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (IdentitybasedrouteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*identitybasedrouteArgs)(nil)).Elem()
}

type IdentitybasedrouteInput interface {
	pulumi.Input

	ToIdentitybasedrouteOutput() IdentitybasedrouteOutput
	ToIdentitybasedrouteOutputWithContext(ctx context.Context) IdentitybasedrouteOutput
}

func (*Identitybasedroute) ElementType() reflect.Type {
	return reflect.TypeOf((**Identitybasedroute)(nil)).Elem()
}

func (i *Identitybasedroute) ToIdentitybasedrouteOutput() IdentitybasedrouteOutput {
	return i.ToIdentitybasedrouteOutputWithContext(context.Background())
}

func (i *Identitybasedroute) ToIdentitybasedrouteOutputWithContext(ctx context.Context) IdentitybasedrouteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentitybasedrouteOutput)
}

// IdentitybasedrouteArrayInput is an input type that accepts IdentitybasedrouteArray and IdentitybasedrouteArrayOutput values.
// You can construct a concrete instance of `IdentitybasedrouteArrayInput` via:
//
//	IdentitybasedrouteArray{ IdentitybasedrouteArgs{...} }
type IdentitybasedrouteArrayInput interface {
	pulumi.Input

	ToIdentitybasedrouteArrayOutput() IdentitybasedrouteArrayOutput
	ToIdentitybasedrouteArrayOutputWithContext(context.Context) IdentitybasedrouteArrayOutput
}

type IdentitybasedrouteArray []IdentitybasedrouteInput

func (IdentitybasedrouteArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Identitybasedroute)(nil)).Elem()
}

func (i IdentitybasedrouteArray) ToIdentitybasedrouteArrayOutput() IdentitybasedrouteArrayOutput {
	return i.ToIdentitybasedrouteArrayOutputWithContext(context.Background())
}

func (i IdentitybasedrouteArray) ToIdentitybasedrouteArrayOutputWithContext(ctx context.Context) IdentitybasedrouteArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentitybasedrouteArrayOutput)
}

// IdentitybasedrouteMapInput is an input type that accepts IdentitybasedrouteMap and IdentitybasedrouteMapOutput values.
// You can construct a concrete instance of `IdentitybasedrouteMapInput` via:
//
//	IdentitybasedrouteMap{ "key": IdentitybasedrouteArgs{...} }
type IdentitybasedrouteMapInput interface {
	pulumi.Input

	ToIdentitybasedrouteMapOutput() IdentitybasedrouteMapOutput
	ToIdentitybasedrouteMapOutputWithContext(context.Context) IdentitybasedrouteMapOutput
}

type IdentitybasedrouteMap map[string]IdentitybasedrouteInput

func (IdentitybasedrouteMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Identitybasedroute)(nil)).Elem()
}

func (i IdentitybasedrouteMap) ToIdentitybasedrouteMapOutput() IdentitybasedrouteMapOutput {
	return i.ToIdentitybasedrouteMapOutputWithContext(context.Background())
}

func (i IdentitybasedrouteMap) ToIdentitybasedrouteMapOutputWithContext(ctx context.Context) IdentitybasedrouteMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentitybasedrouteMapOutput)
}

type IdentitybasedrouteOutput struct{ *pulumi.OutputState }

func (IdentitybasedrouteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Identitybasedroute)(nil)).Elem()
}

func (o IdentitybasedrouteOutput) ToIdentitybasedrouteOutput() IdentitybasedrouteOutput {
	return o
}

func (o IdentitybasedrouteOutput) ToIdentitybasedrouteOutputWithContext(ctx context.Context) IdentitybasedrouteOutput {
	return o
}

// Comments.
func (o IdentitybasedrouteOutput) Comments() pulumi.StringOutput {
	return o.ApplyT(func(v *Identitybasedroute) pulumi.StringOutput { return v.Comments }).(pulumi.StringOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o IdentitybasedrouteOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Identitybasedroute) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Name.
func (o IdentitybasedrouteOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Identitybasedroute) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Rule. The structure of `rule` block is documented below.
func (o IdentitybasedrouteOutput) Rules() IdentitybasedrouteRuleArrayOutput {
	return o.ApplyT(func(v *Identitybasedroute) IdentitybasedrouteRuleArrayOutput { return v.Rules }).(IdentitybasedrouteRuleArrayOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o IdentitybasedrouteOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Identitybasedroute) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type IdentitybasedrouteArrayOutput struct{ *pulumi.OutputState }

func (IdentitybasedrouteArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Identitybasedroute)(nil)).Elem()
}

func (o IdentitybasedrouteArrayOutput) ToIdentitybasedrouteArrayOutput() IdentitybasedrouteArrayOutput {
	return o
}

func (o IdentitybasedrouteArrayOutput) ToIdentitybasedrouteArrayOutputWithContext(ctx context.Context) IdentitybasedrouteArrayOutput {
	return o
}

func (o IdentitybasedrouteArrayOutput) Index(i pulumi.IntInput) IdentitybasedrouteOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Identitybasedroute {
		return vs[0].([]*Identitybasedroute)[vs[1].(int)]
	}).(IdentitybasedrouteOutput)
}

type IdentitybasedrouteMapOutput struct{ *pulumi.OutputState }

func (IdentitybasedrouteMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Identitybasedroute)(nil)).Elem()
}

func (o IdentitybasedrouteMapOutput) ToIdentitybasedrouteMapOutput() IdentitybasedrouteMapOutput {
	return o
}

func (o IdentitybasedrouteMapOutput) ToIdentitybasedrouteMapOutputWithContext(ctx context.Context) IdentitybasedrouteMapOutput {
	return o
}

func (o IdentitybasedrouteMapOutput) MapIndex(k pulumi.StringInput) IdentitybasedrouteOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Identitybasedroute {
		return vs[0].(map[string]*Identitybasedroute)[vs[1].(string)]
	}).(IdentitybasedrouteOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IdentitybasedrouteInput)(nil)).Elem(), &Identitybasedroute{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdentitybasedrouteArrayInput)(nil)).Elem(), IdentitybasedrouteArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdentitybasedrouteMapInput)(nil)).Elem(), IdentitybasedrouteMap{})
	pulumi.RegisterOutputType(IdentitybasedrouteOutput{})
	pulumi.RegisterOutputType(IdentitybasedrouteArrayOutput{})
	pulumi.RegisterOutputType(IdentitybasedrouteMapOutput{})
}
