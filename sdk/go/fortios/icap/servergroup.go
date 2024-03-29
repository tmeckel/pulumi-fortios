// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package icap

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure an ICAP server group consisting of multiple forward servers. Supports failover and load balancing. Applies to FortiOS Version `>= 7.2.0`.
//
// ## Import
//
// # Icap ServerGroup can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:icap/servergroup:Servergroup labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:icap/servergroup:Servergroup labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type Servergroup struct {
	pulumi.CustomResourceState

	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Load balance method. Valid values: `weighted`, `least-session`, `active-passive`.
	LdbMethod pulumi.StringOutput `pulumi:"ldbMethod"`
	// Configure an ICAP server group consisting one or multiple forward servers. Supports failover and load balancing.
	Name pulumi.StringOutput `pulumi:"name"`
	// Add ICAP servers to a list to form a server group. Optionally assign weights to each server. The structure of `serverList` block is documented below.
	ServerLists ServergroupServerListArrayOutput `pulumi:"serverLists"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewServergroup registers a new resource with the given unique name, arguments, and options.
func NewServergroup(ctx *pulumi.Context,
	name string, args *ServergroupArgs, opts ...pulumi.ResourceOption) (*Servergroup, error) {
	if args == nil {
		args = &ServergroupArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource Servergroup
	err := ctx.RegisterResource("fortios:icap/servergroup:Servergroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServergroup gets an existing Servergroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServergroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServergroupState, opts ...pulumi.ResourceOption) (*Servergroup, error) {
	var resource Servergroup
	err := ctx.ReadResource("fortios:icap/servergroup:Servergroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Servergroup resources.
type servergroupState struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Load balance method. Valid values: `weighted`, `least-session`, `active-passive`.
	LdbMethod *string `pulumi:"ldbMethod"`
	// Configure an ICAP server group consisting one or multiple forward servers. Supports failover and load balancing.
	Name *string `pulumi:"name"`
	// Add ICAP servers to a list to form a server group. Optionally assign weights to each server. The structure of `serverList` block is documented below.
	ServerLists []ServergroupServerList `pulumi:"serverLists"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type ServergroupState struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Load balance method. Valid values: `weighted`, `least-session`, `active-passive`.
	LdbMethod pulumi.StringPtrInput
	// Configure an ICAP server group consisting one or multiple forward servers. Supports failover and load balancing.
	Name pulumi.StringPtrInput
	// Add ICAP servers to a list to form a server group. Optionally assign weights to each server. The structure of `serverList` block is documented below.
	ServerLists ServergroupServerListArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (ServergroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*servergroupState)(nil)).Elem()
}

type servergroupArgs struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Load balance method. Valid values: `weighted`, `least-session`, `active-passive`.
	LdbMethod *string `pulumi:"ldbMethod"`
	// Configure an ICAP server group consisting one or multiple forward servers. Supports failover and load balancing.
	Name *string `pulumi:"name"`
	// Add ICAP servers to a list to form a server group. Optionally assign weights to each server. The structure of `serverList` block is documented below.
	ServerLists []ServergroupServerList `pulumi:"serverLists"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Servergroup resource.
type ServergroupArgs struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Load balance method. Valid values: `weighted`, `least-session`, `active-passive`.
	LdbMethod pulumi.StringPtrInput
	// Configure an ICAP server group consisting one or multiple forward servers. Supports failover and load balancing.
	Name pulumi.StringPtrInput
	// Add ICAP servers to a list to form a server group. Optionally assign weights to each server. The structure of `serverList` block is documented below.
	ServerLists ServergroupServerListArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (ServergroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*servergroupArgs)(nil)).Elem()
}

type ServergroupInput interface {
	pulumi.Input

	ToServergroupOutput() ServergroupOutput
	ToServergroupOutputWithContext(ctx context.Context) ServergroupOutput
}

func (*Servergroup) ElementType() reflect.Type {
	return reflect.TypeOf((**Servergroup)(nil)).Elem()
}

func (i *Servergroup) ToServergroupOutput() ServergroupOutput {
	return i.ToServergroupOutputWithContext(context.Background())
}

func (i *Servergroup) ToServergroupOutputWithContext(ctx context.Context) ServergroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServergroupOutput)
}

// ServergroupArrayInput is an input type that accepts ServergroupArray and ServergroupArrayOutput values.
// You can construct a concrete instance of `ServergroupArrayInput` via:
//
//	ServergroupArray{ ServergroupArgs{...} }
type ServergroupArrayInput interface {
	pulumi.Input

	ToServergroupArrayOutput() ServergroupArrayOutput
	ToServergroupArrayOutputWithContext(context.Context) ServergroupArrayOutput
}

type ServergroupArray []ServergroupInput

func (ServergroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Servergroup)(nil)).Elem()
}

func (i ServergroupArray) ToServergroupArrayOutput() ServergroupArrayOutput {
	return i.ToServergroupArrayOutputWithContext(context.Background())
}

func (i ServergroupArray) ToServergroupArrayOutputWithContext(ctx context.Context) ServergroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServergroupArrayOutput)
}

// ServergroupMapInput is an input type that accepts ServergroupMap and ServergroupMapOutput values.
// You can construct a concrete instance of `ServergroupMapInput` via:
//
//	ServergroupMap{ "key": ServergroupArgs{...} }
type ServergroupMapInput interface {
	pulumi.Input

	ToServergroupMapOutput() ServergroupMapOutput
	ToServergroupMapOutputWithContext(context.Context) ServergroupMapOutput
}

type ServergroupMap map[string]ServergroupInput

func (ServergroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Servergroup)(nil)).Elem()
}

func (i ServergroupMap) ToServergroupMapOutput() ServergroupMapOutput {
	return i.ToServergroupMapOutputWithContext(context.Background())
}

func (i ServergroupMap) ToServergroupMapOutputWithContext(ctx context.Context) ServergroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServergroupMapOutput)
}

type ServergroupOutput struct{ *pulumi.OutputState }

func (ServergroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Servergroup)(nil)).Elem()
}

func (o ServergroupOutput) ToServergroupOutput() ServergroupOutput {
	return o
}

func (o ServergroupOutput) ToServergroupOutputWithContext(ctx context.Context) ServergroupOutput {
	return o
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o ServergroupOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Servergroup) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Load balance method. Valid values: `weighted`, `least-session`, `active-passive`.
func (o ServergroupOutput) LdbMethod() pulumi.StringOutput {
	return o.ApplyT(func(v *Servergroup) pulumi.StringOutput { return v.LdbMethod }).(pulumi.StringOutput)
}

// Configure an ICAP server group consisting one or multiple forward servers. Supports failover and load balancing.
func (o ServergroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Servergroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Add ICAP servers to a list to form a server group. Optionally assign weights to each server. The structure of `serverList` block is documented below.
func (o ServergroupOutput) ServerLists() ServergroupServerListArrayOutput {
	return o.ApplyT(func(v *Servergroup) ServergroupServerListArrayOutput { return v.ServerLists }).(ServergroupServerListArrayOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o ServergroupOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Servergroup) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type ServergroupArrayOutput struct{ *pulumi.OutputState }

func (ServergroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Servergroup)(nil)).Elem()
}

func (o ServergroupArrayOutput) ToServergroupArrayOutput() ServergroupArrayOutput {
	return o
}

func (o ServergroupArrayOutput) ToServergroupArrayOutputWithContext(ctx context.Context) ServergroupArrayOutput {
	return o
}

func (o ServergroupArrayOutput) Index(i pulumi.IntInput) ServergroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Servergroup {
		return vs[0].([]*Servergroup)[vs[1].(int)]
	}).(ServergroupOutput)
}

type ServergroupMapOutput struct{ *pulumi.OutputState }

func (ServergroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Servergroup)(nil)).Elem()
}

func (o ServergroupMapOutput) ToServergroupMapOutput() ServergroupMapOutput {
	return o
}

func (o ServergroupMapOutput) ToServergroupMapOutputWithContext(ctx context.Context) ServergroupMapOutput {
	return o
}

func (o ServergroupMapOutput) MapIndex(k pulumi.StringInput) ServergroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Servergroup {
		return vs[0].(map[string]*Servergroup)[vs[1].(string)]
	}).(ServergroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServergroupInput)(nil)).Elem(), &Servergroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServergroupArrayInput)(nil)).Elem(), ServergroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServergroupMapInput)(nil)).Elem(), ServergroupMap{})
	pulumi.RegisterOutputType(ServergroupOutput{})
	pulumi.RegisterOutputType(ServergroupArrayOutput{})
	pulumi.RegisterOutputType(ServergroupMapOutput{})
}
