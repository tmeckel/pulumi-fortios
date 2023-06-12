// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sys

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure ACME client. Applies to FortiOS Version `>= 7.0.1`.
//
// ## Import
//
// # System Acme can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:sys/acme:Acme labelname SystemAcme
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:sys/acme:Acme labelname SystemAcme
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type Acme struct {
	pulumi.CustomResourceState

	// ACME accounts list. The structure of `accounts` block is documented below.
	Accounts AcmeAccountArrayOutput `pulumi:"accounts"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Interface(s) on which the ACME client will listen for challenges. The structure of `interface` block is documented below.
	Interfaces AcmeInterfaceArrayOutput `pulumi:"interfaces"`
	// Source IPv4 address used to connect to the ACME server.
	SourceIp pulumi.StringOutput `pulumi:"sourceIp"`
	// Source IPv6 address used to connect to the ACME server.
	SourceIp6 pulumi.StringOutput `pulumi:"sourceIp6"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewAcme registers a new resource with the given unique name, arguments, and options.
func NewAcme(ctx *pulumi.Context,
	name string, args *AcmeArgs, opts ...pulumi.ResourceOption) (*Acme, error) {
	if args == nil {
		args = &AcmeArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource Acme
	err := ctx.RegisterResource("fortios:sys/acme:Acme", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAcme gets an existing Acme resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAcme(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AcmeState, opts ...pulumi.ResourceOption) (*Acme, error) {
	var resource Acme
	err := ctx.ReadResource("fortios:sys/acme:Acme", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Acme resources.
type acmeState struct {
	// ACME accounts list. The structure of `accounts` block is documented below.
	Accounts []AcmeAccount `pulumi:"accounts"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Interface(s) on which the ACME client will listen for challenges. The structure of `interface` block is documented below.
	Interfaces []AcmeInterface `pulumi:"interfaces"`
	// Source IPv4 address used to connect to the ACME server.
	SourceIp *string `pulumi:"sourceIp"`
	// Source IPv6 address used to connect to the ACME server.
	SourceIp6 *string `pulumi:"sourceIp6"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type AcmeState struct {
	// ACME accounts list. The structure of `accounts` block is documented below.
	Accounts AcmeAccountArrayInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Interface(s) on which the ACME client will listen for challenges. The structure of `interface` block is documented below.
	Interfaces AcmeInterfaceArrayInput
	// Source IPv4 address used to connect to the ACME server.
	SourceIp pulumi.StringPtrInput
	// Source IPv6 address used to connect to the ACME server.
	SourceIp6 pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (AcmeState) ElementType() reflect.Type {
	return reflect.TypeOf((*acmeState)(nil)).Elem()
}

type acmeArgs struct {
	// ACME accounts list. The structure of `accounts` block is documented below.
	Accounts []AcmeAccount `pulumi:"accounts"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Interface(s) on which the ACME client will listen for challenges. The structure of `interface` block is documented below.
	Interfaces []AcmeInterface `pulumi:"interfaces"`
	// Source IPv4 address used to connect to the ACME server.
	SourceIp *string `pulumi:"sourceIp"`
	// Source IPv6 address used to connect to the ACME server.
	SourceIp6 *string `pulumi:"sourceIp6"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Acme resource.
type AcmeArgs struct {
	// ACME accounts list. The structure of `accounts` block is documented below.
	Accounts AcmeAccountArrayInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Interface(s) on which the ACME client will listen for challenges. The structure of `interface` block is documented below.
	Interfaces AcmeInterfaceArrayInput
	// Source IPv4 address used to connect to the ACME server.
	SourceIp pulumi.StringPtrInput
	// Source IPv6 address used to connect to the ACME server.
	SourceIp6 pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (AcmeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*acmeArgs)(nil)).Elem()
}

type AcmeInput interface {
	pulumi.Input

	ToAcmeOutput() AcmeOutput
	ToAcmeOutputWithContext(ctx context.Context) AcmeOutput
}

func (*Acme) ElementType() reflect.Type {
	return reflect.TypeOf((**Acme)(nil)).Elem()
}

func (i *Acme) ToAcmeOutput() AcmeOutput {
	return i.ToAcmeOutputWithContext(context.Background())
}

func (i *Acme) ToAcmeOutputWithContext(ctx context.Context) AcmeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AcmeOutput)
}

// AcmeArrayInput is an input type that accepts AcmeArray and AcmeArrayOutput values.
// You can construct a concrete instance of `AcmeArrayInput` via:
//
//	AcmeArray{ AcmeArgs{...} }
type AcmeArrayInput interface {
	pulumi.Input

	ToAcmeArrayOutput() AcmeArrayOutput
	ToAcmeArrayOutputWithContext(context.Context) AcmeArrayOutput
}

type AcmeArray []AcmeInput

func (AcmeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Acme)(nil)).Elem()
}

func (i AcmeArray) ToAcmeArrayOutput() AcmeArrayOutput {
	return i.ToAcmeArrayOutputWithContext(context.Background())
}

func (i AcmeArray) ToAcmeArrayOutputWithContext(ctx context.Context) AcmeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AcmeArrayOutput)
}

// AcmeMapInput is an input type that accepts AcmeMap and AcmeMapOutput values.
// You can construct a concrete instance of `AcmeMapInput` via:
//
//	AcmeMap{ "key": AcmeArgs{...} }
type AcmeMapInput interface {
	pulumi.Input

	ToAcmeMapOutput() AcmeMapOutput
	ToAcmeMapOutputWithContext(context.Context) AcmeMapOutput
}

type AcmeMap map[string]AcmeInput

func (AcmeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Acme)(nil)).Elem()
}

func (i AcmeMap) ToAcmeMapOutput() AcmeMapOutput {
	return i.ToAcmeMapOutputWithContext(context.Background())
}

func (i AcmeMap) ToAcmeMapOutputWithContext(ctx context.Context) AcmeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AcmeMapOutput)
}

type AcmeOutput struct{ *pulumi.OutputState }

func (AcmeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Acme)(nil)).Elem()
}

func (o AcmeOutput) ToAcmeOutput() AcmeOutput {
	return o
}

func (o AcmeOutput) ToAcmeOutputWithContext(ctx context.Context) AcmeOutput {
	return o
}

// ACME accounts list. The structure of `accounts` block is documented below.
func (o AcmeOutput) Accounts() AcmeAccountArrayOutput {
	return o.ApplyT(func(v *Acme) AcmeAccountArrayOutput { return v.Accounts }).(AcmeAccountArrayOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o AcmeOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Acme) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Interface(s) on which the ACME client will listen for challenges. The structure of `interface` block is documented below.
func (o AcmeOutput) Interfaces() AcmeInterfaceArrayOutput {
	return o.ApplyT(func(v *Acme) AcmeInterfaceArrayOutput { return v.Interfaces }).(AcmeInterfaceArrayOutput)
}

// Source IPv4 address used to connect to the ACME server.
func (o AcmeOutput) SourceIp() pulumi.StringOutput {
	return o.ApplyT(func(v *Acme) pulumi.StringOutput { return v.SourceIp }).(pulumi.StringOutput)
}

// Source IPv6 address used to connect to the ACME server.
func (o AcmeOutput) SourceIp6() pulumi.StringOutput {
	return o.ApplyT(func(v *Acme) pulumi.StringOutput { return v.SourceIp6 }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o AcmeOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Acme) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type AcmeArrayOutput struct{ *pulumi.OutputState }

func (AcmeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Acme)(nil)).Elem()
}

func (o AcmeArrayOutput) ToAcmeArrayOutput() AcmeArrayOutput {
	return o
}

func (o AcmeArrayOutput) ToAcmeArrayOutputWithContext(ctx context.Context) AcmeArrayOutput {
	return o
}

func (o AcmeArrayOutput) Index(i pulumi.IntInput) AcmeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Acme {
		return vs[0].([]*Acme)[vs[1].(int)]
	}).(AcmeOutput)
}

type AcmeMapOutput struct{ *pulumi.OutputState }

func (AcmeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Acme)(nil)).Elem()
}

func (o AcmeMapOutput) ToAcmeMapOutput() AcmeMapOutput {
	return o
}

func (o AcmeMapOutput) ToAcmeMapOutputWithContext(ctx context.Context) AcmeMapOutput {
	return o
}

func (o AcmeMapOutput) MapIndex(k pulumi.StringInput) AcmeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Acme {
		return vs[0].(map[string]*Acme)[vs[1].(string)]
	}).(AcmeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AcmeInput)(nil)).Elem(), &Acme{})
	pulumi.RegisterInputType(reflect.TypeOf((*AcmeArrayInput)(nil)).Elem(), AcmeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AcmeMapInput)(nil)).Elem(), AcmeMap{})
	pulumi.RegisterOutputType(AcmeOutput{})
	pulumi.RegisterOutputType(AcmeArrayOutput{})
	pulumi.RegisterOutputType(AcmeMapOutput{})
}