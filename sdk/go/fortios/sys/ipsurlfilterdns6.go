// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sys

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure IPS URL filter IPv6 DNS servers.
//
// ## Import
//
// # System IpsUrlfilterDns6 can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:sys/ipsurlfilterdns6:Ipsurlfilterdns6 labelname {{address6}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:sys/ipsurlfilterdns6:Ipsurlfilterdns6 labelname {{address6}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type Ipsurlfilterdns6 struct {
	pulumi.CustomResourceState

	// IPv6 address of DNS server.
	Address6 pulumi.StringOutput `pulumi:"address6"`
	// Enable/disable this server for IPv6 DNS queries. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewIpsurlfilterdns6 registers a new resource with the given unique name, arguments, and options.
func NewIpsurlfilterdns6(ctx *pulumi.Context,
	name string, args *Ipsurlfilterdns6Args, opts ...pulumi.ResourceOption) (*Ipsurlfilterdns6, error) {
	if args == nil {
		args = &Ipsurlfilterdns6Args{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource Ipsurlfilterdns6
	err := ctx.RegisterResource("fortios:sys/ipsurlfilterdns6:Ipsurlfilterdns6", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIpsurlfilterdns6 gets an existing Ipsurlfilterdns6 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIpsurlfilterdns6(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Ipsurlfilterdns6State, opts ...pulumi.ResourceOption) (*Ipsurlfilterdns6, error) {
	var resource Ipsurlfilterdns6
	err := ctx.ReadResource("fortios:sys/ipsurlfilterdns6:Ipsurlfilterdns6", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Ipsurlfilterdns6 resources.
type ipsurlfilterdns6State struct {
	// IPv6 address of DNS server.
	Address6 *string `pulumi:"address6"`
	// Enable/disable this server for IPv6 DNS queries. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type Ipsurlfilterdns6State struct {
	// IPv6 address of DNS server.
	Address6 pulumi.StringPtrInput
	// Enable/disable this server for IPv6 DNS queries. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (Ipsurlfilterdns6State) ElementType() reflect.Type {
	return reflect.TypeOf((*ipsurlfilterdns6State)(nil)).Elem()
}

type ipsurlfilterdns6Args struct {
	// IPv6 address of DNS server.
	Address6 *string `pulumi:"address6"`
	// Enable/disable this server for IPv6 DNS queries. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Ipsurlfilterdns6 resource.
type Ipsurlfilterdns6Args struct {
	// IPv6 address of DNS server.
	Address6 pulumi.StringPtrInput
	// Enable/disable this server for IPv6 DNS queries. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (Ipsurlfilterdns6Args) ElementType() reflect.Type {
	return reflect.TypeOf((*ipsurlfilterdns6Args)(nil)).Elem()
}

type Ipsurlfilterdns6Input interface {
	pulumi.Input

	ToIpsurlfilterdns6Output() Ipsurlfilterdns6Output
	ToIpsurlfilterdns6OutputWithContext(ctx context.Context) Ipsurlfilterdns6Output
}

func (*Ipsurlfilterdns6) ElementType() reflect.Type {
	return reflect.TypeOf((**Ipsurlfilterdns6)(nil)).Elem()
}

func (i *Ipsurlfilterdns6) ToIpsurlfilterdns6Output() Ipsurlfilterdns6Output {
	return i.ToIpsurlfilterdns6OutputWithContext(context.Background())
}

func (i *Ipsurlfilterdns6) ToIpsurlfilterdns6OutputWithContext(ctx context.Context) Ipsurlfilterdns6Output {
	return pulumi.ToOutputWithContext(ctx, i).(Ipsurlfilterdns6Output)
}

// Ipsurlfilterdns6ArrayInput is an input type that accepts Ipsurlfilterdns6Array and Ipsurlfilterdns6ArrayOutput values.
// You can construct a concrete instance of `Ipsurlfilterdns6ArrayInput` via:
//
//	Ipsurlfilterdns6Array{ Ipsurlfilterdns6Args{...} }
type Ipsurlfilterdns6ArrayInput interface {
	pulumi.Input

	ToIpsurlfilterdns6ArrayOutput() Ipsurlfilterdns6ArrayOutput
	ToIpsurlfilterdns6ArrayOutputWithContext(context.Context) Ipsurlfilterdns6ArrayOutput
}

type Ipsurlfilterdns6Array []Ipsurlfilterdns6Input

func (Ipsurlfilterdns6Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ipsurlfilterdns6)(nil)).Elem()
}

func (i Ipsurlfilterdns6Array) ToIpsurlfilterdns6ArrayOutput() Ipsurlfilterdns6ArrayOutput {
	return i.ToIpsurlfilterdns6ArrayOutputWithContext(context.Background())
}

func (i Ipsurlfilterdns6Array) ToIpsurlfilterdns6ArrayOutputWithContext(ctx context.Context) Ipsurlfilterdns6ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ipsurlfilterdns6ArrayOutput)
}

// Ipsurlfilterdns6MapInput is an input type that accepts Ipsurlfilterdns6Map and Ipsurlfilterdns6MapOutput values.
// You can construct a concrete instance of `Ipsurlfilterdns6MapInput` via:
//
//	Ipsurlfilterdns6Map{ "key": Ipsurlfilterdns6Args{...} }
type Ipsurlfilterdns6MapInput interface {
	pulumi.Input

	ToIpsurlfilterdns6MapOutput() Ipsurlfilterdns6MapOutput
	ToIpsurlfilterdns6MapOutputWithContext(context.Context) Ipsurlfilterdns6MapOutput
}

type Ipsurlfilterdns6Map map[string]Ipsurlfilterdns6Input

func (Ipsurlfilterdns6Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ipsurlfilterdns6)(nil)).Elem()
}

func (i Ipsurlfilterdns6Map) ToIpsurlfilterdns6MapOutput() Ipsurlfilterdns6MapOutput {
	return i.ToIpsurlfilterdns6MapOutputWithContext(context.Background())
}

func (i Ipsurlfilterdns6Map) ToIpsurlfilterdns6MapOutputWithContext(ctx context.Context) Ipsurlfilterdns6MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ipsurlfilterdns6MapOutput)
}

type Ipsurlfilterdns6Output struct{ *pulumi.OutputState }

func (Ipsurlfilterdns6Output) ElementType() reflect.Type {
	return reflect.TypeOf((**Ipsurlfilterdns6)(nil)).Elem()
}

func (o Ipsurlfilterdns6Output) ToIpsurlfilterdns6Output() Ipsurlfilterdns6Output {
	return o
}

func (o Ipsurlfilterdns6Output) ToIpsurlfilterdns6OutputWithContext(ctx context.Context) Ipsurlfilterdns6Output {
	return o
}

// IPv6 address of DNS server.
func (o Ipsurlfilterdns6Output) Address6() pulumi.StringOutput {
	return o.ApplyT(func(v *Ipsurlfilterdns6) pulumi.StringOutput { return v.Address6 }).(pulumi.StringOutput)
}

// Enable/disable this server for IPv6 DNS queries. Valid values: `enable`, `disable`.
func (o Ipsurlfilterdns6Output) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Ipsurlfilterdns6) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o Ipsurlfilterdns6Output) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ipsurlfilterdns6) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type Ipsurlfilterdns6ArrayOutput struct{ *pulumi.OutputState }

func (Ipsurlfilterdns6ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ipsurlfilterdns6)(nil)).Elem()
}

func (o Ipsurlfilterdns6ArrayOutput) ToIpsurlfilterdns6ArrayOutput() Ipsurlfilterdns6ArrayOutput {
	return o
}

func (o Ipsurlfilterdns6ArrayOutput) ToIpsurlfilterdns6ArrayOutputWithContext(ctx context.Context) Ipsurlfilterdns6ArrayOutput {
	return o
}

func (o Ipsurlfilterdns6ArrayOutput) Index(i pulumi.IntInput) Ipsurlfilterdns6Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Ipsurlfilterdns6 {
		return vs[0].([]*Ipsurlfilterdns6)[vs[1].(int)]
	}).(Ipsurlfilterdns6Output)
}

type Ipsurlfilterdns6MapOutput struct{ *pulumi.OutputState }

func (Ipsurlfilterdns6MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ipsurlfilterdns6)(nil)).Elem()
}

func (o Ipsurlfilterdns6MapOutput) ToIpsurlfilterdns6MapOutput() Ipsurlfilterdns6MapOutput {
	return o
}

func (o Ipsurlfilterdns6MapOutput) ToIpsurlfilterdns6MapOutputWithContext(ctx context.Context) Ipsurlfilterdns6MapOutput {
	return o
}

func (o Ipsurlfilterdns6MapOutput) MapIndex(k pulumi.StringInput) Ipsurlfilterdns6Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Ipsurlfilterdns6 {
		return vs[0].(map[string]*Ipsurlfilterdns6)[vs[1].(string)]
	}).(Ipsurlfilterdns6Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*Ipsurlfilterdns6Input)(nil)).Elem(), &Ipsurlfilterdns6{})
	pulumi.RegisterInputType(reflect.TypeOf((*Ipsurlfilterdns6ArrayInput)(nil)).Elem(), Ipsurlfilterdns6Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*Ipsurlfilterdns6MapInput)(nil)).Elem(), Ipsurlfilterdns6Map{})
	pulumi.RegisterOutputType(Ipsurlfilterdns6Output{})
	pulumi.RegisterOutputType(Ipsurlfilterdns6ArrayOutput{})
	pulumi.RegisterOutputType(Ipsurlfilterdns6MapOutput{})
}
