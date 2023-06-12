// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firewall

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure IPv6 to IPv4 virtual IPs. Applies to FortiOS Version `<= 7.0.0`.
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
//			_, err := firewall.NewVip64(ctx, "trname", &firewall.Vip64Args{
//				ArpReply:    pulumi.String("enable"),
//				Color:       pulumi.Int(0),
//				Extip:       pulumi.String("2001:db8:99:203::22"),
//				Extport:     pulumi.String("0-65535"),
//				Fosid:       pulumi.Int(0),
//				LdbMethod:   pulumi.String("static"),
//				Mappedip:    pulumi.String("1.1.1.1"),
//				Mappedport:  pulumi.String("0-65535"),
//				Portforward: pulumi.String("disable"),
//				Protocol:    pulumi.String("tcp"),
//				Type:        pulumi.String("static-nat"),
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
// # Firewall Vip64 can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:firewall/vip64:Vip64 labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:firewall/vip64:Vip64 labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type Vip64 struct {
	pulumi.CustomResourceState

	// Enable ARP reply. Valid values: `disable`, `enable`.
	ArpReply pulumi.StringOutput `pulumi:"arpReply"`
	// Color of icon on the GUI.
	Color pulumi.IntOutput `pulumi:"color"`
	// Comment.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Start-external-IP [-end-external-IP].
	Extip pulumi.StringOutput `pulumi:"extip"`
	// External service port.
	Extport pulumi.StringOutput `pulumi:"extport"`
	// Custom defined id.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// Load balance method. Valid values: `static`, `round-robin`, `weighted`, `least-session`, `least-rtt`, `first-alive`.
	LdbMethod pulumi.StringOutput `pulumi:"ldbMethod"`
	// Start-mapped-IP [-end-mapped-IP].
	Mappedip pulumi.StringOutput `pulumi:"mappedip"`
	// Mapped service port.
	Mappedport pulumi.StringOutput `pulumi:"mappedport"`
	// Health monitors. The structure of `monitor` block is documented below.
	Monitors Vip64MonitorArrayOutput `pulumi:"monitors"`
	// VIP64 name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Enable port forwarding. Valid values: `disable`, `enable`.
	Portforward pulumi.StringOutput `pulumi:"portforward"`
	// Mapped port protocol. Valid values: `tcp`, `udp`.
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// Real servers. The structure of `realservers` block is documented below.
	Realservers Vip64RealserverArrayOutput `pulumi:"realservers"`
	// Server type. Valid values: `http`, `tcp`, `udp`, `ip`.
	ServerType pulumi.StringOutput `pulumi:"serverType"`
	// Source IP6 filter (x:x:x:x:x:x:x:x/x). The structure of `srcFilter` block is documented below.
	SrcFilters Vip64SrcFilterArrayOutput `pulumi:"srcFilters"`
	// VIP type: static NAT or server load balance. Valid values: `static-nat`, `server-load-balance`.
	Type pulumi.StringOutput `pulumi:"type"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid pulumi.StringOutput `pulumi:"uuid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewVip64 registers a new resource with the given unique name, arguments, and options.
func NewVip64(ctx *pulumi.Context,
	name string, args *Vip64Args, opts ...pulumi.ResourceOption) (*Vip64, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Extip == nil {
		return nil, errors.New("invalid value for required argument 'Extip'")
	}
	if args.Mappedip == nil {
		return nil, errors.New("invalid value for required argument 'Mappedip'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Vip64
	err := ctx.RegisterResource("fortios:firewall/vip64:Vip64", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVip64 gets an existing Vip64 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVip64(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Vip64State, opts ...pulumi.ResourceOption) (*Vip64, error) {
	var resource Vip64
	err := ctx.ReadResource("fortios:firewall/vip64:Vip64", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Vip64 resources.
type vip64State struct {
	// Enable ARP reply. Valid values: `disable`, `enable`.
	ArpReply *string `pulumi:"arpReply"`
	// Color of icon on the GUI.
	Color *int `pulumi:"color"`
	// Comment.
	Comment *string `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Start-external-IP [-end-external-IP].
	Extip *string `pulumi:"extip"`
	// External service port.
	Extport *string `pulumi:"extport"`
	// Custom defined id.
	Fosid *int `pulumi:"fosid"`
	// Load balance method. Valid values: `static`, `round-robin`, `weighted`, `least-session`, `least-rtt`, `first-alive`.
	LdbMethod *string `pulumi:"ldbMethod"`
	// Start-mapped-IP [-end-mapped-IP].
	Mappedip *string `pulumi:"mappedip"`
	// Mapped service port.
	Mappedport *string `pulumi:"mappedport"`
	// Health monitors. The structure of `monitor` block is documented below.
	Monitors []Vip64Monitor `pulumi:"monitors"`
	// VIP64 name.
	Name *string `pulumi:"name"`
	// Enable port forwarding. Valid values: `disable`, `enable`.
	Portforward *string `pulumi:"portforward"`
	// Mapped port protocol. Valid values: `tcp`, `udp`.
	Protocol *string `pulumi:"protocol"`
	// Real servers. The structure of `realservers` block is documented below.
	Realservers []Vip64Realserver `pulumi:"realservers"`
	// Server type. Valid values: `http`, `tcp`, `udp`, `ip`.
	ServerType *string `pulumi:"serverType"`
	// Source IP6 filter (x:x:x:x:x:x:x:x/x). The structure of `srcFilter` block is documented below.
	SrcFilters []Vip64SrcFilter `pulumi:"srcFilters"`
	// VIP type: static NAT or server load balance. Valid values: `static-nat`, `server-load-balance`.
	Type *string `pulumi:"type"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid *string `pulumi:"uuid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type Vip64State struct {
	// Enable ARP reply. Valid values: `disable`, `enable`.
	ArpReply pulumi.StringPtrInput
	// Color of icon on the GUI.
	Color pulumi.IntPtrInput
	// Comment.
	Comment pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Start-external-IP [-end-external-IP].
	Extip pulumi.StringPtrInput
	// External service port.
	Extport pulumi.StringPtrInput
	// Custom defined id.
	Fosid pulumi.IntPtrInput
	// Load balance method. Valid values: `static`, `round-robin`, `weighted`, `least-session`, `least-rtt`, `first-alive`.
	LdbMethod pulumi.StringPtrInput
	// Start-mapped-IP [-end-mapped-IP].
	Mappedip pulumi.StringPtrInput
	// Mapped service port.
	Mappedport pulumi.StringPtrInput
	// Health monitors. The structure of `monitor` block is documented below.
	Monitors Vip64MonitorArrayInput
	// VIP64 name.
	Name pulumi.StringPtrInput
	// Enable port forwarding. Valid values: `disable`, `enable`.
	Portforward pulumi.StringPtrInput
	// Mapped port protocol. Valid values: `tcp`, `udp`.
	Protocol pulumi.StringPtrInput
	// Real servers. The structure of `realservers` block is documented below.
	Realservers Vip64RealserverArrayInput
	// Server type. Valid values: `http`, `tcp`, `udp`, `ip`.
	ServerType pulumi.StringPtrInput
	// Source IP6 filter (x:x:x:x:x:x:x:x/x). The structure of `srcFilter` block is documented below.
	SrcFilters Vip64SrcFilterArrayInput
	// VIP type: static NAT or server load balance. Valid values: `static-nat`, `server-load-balance`.
	Type pulumi.StringPtrInput
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (Vip64State) ElementType() reflect.Type {
	return reflect.TypeOf((*vip64State)(nil)).Elem()
}

type vip64Args struct {
	// Enable ARP reply. Valid values: `disable`, `enable`.
	ArpReply *string `pulumi:"arpReply"`
	// Color of icon on the GUI.
	Color *int `pulumi:"color"`
	// Comment.
	Comment *string `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Start-external-IP [-end-external-IP].
	Extip string `pulumi:"extip"`
	// External service port.
	Extport *string `pulumi:"extport"`
	// Custom defined id.
	Fosid *int `pulumi:"fosid"`
	// Load balance method. Valid values: `static`, `round-robin`, `weighted`, `least-session`, `least-rtt`, `first-alive`.
	LdbMethod *string `pulumi:"ldbMethod"`
	// Start-mapped-IP [-end-mapped-IP].
	Mappedip string `pulumi:"mappedip"`
	// Mapped service port.
	Mappedport *string `pulumi:"mappedport"`
	// Health monitors. The structure of `monitor` block is documented below.
	Monitors []Vip64Monitor `pulumi:"monitors"`
	// VIP64 name.
	Name *string `pulumi:"name"`
	// Enable port forwarding. Valid values: `disable`, `enable`.
	Portforward *string `pulumi:"portforward"`
	// Mapped port protocol. Valid values: `tcp`, `udp`.
	Protocol *string `pulumi:"protocol"`
	// Real servers. The structure of `realservers` block is documented below.
	Realservers []Vip64Realserver `pulumi:"realservers"`
	// Server type. Valid values: `http`, `tcp`, `udp`, `ip`.
	ServerType *string `pulumi:"serverType"`
	// Source IP6 filter (x:x:x:x:x:x:x:x/x). The structure of `srcFilter` block is documented below.
	SrcFilters []Vip64SrcFilter `pulumi:"srcFilters"`
	// VIP type: static NAT or server load balance. Valid values: `static-nat`, `server-load-balance`.
	Type *string `pulumi:"type"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid *string `pulumi:"uuid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Vip64 resource.
type Vip64Args struct {
	// Enable ARP reply. Valid values: `disable`, `enable`.
	ArpReply pulumi.StringPtrInput
	// Color of icon on the GUI.
	Color pulumi.IntPtrInput
	// Comment.
	Comment pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Start-external-IP [-end-external-IP].
	Extip pulumi.StringInput
	// External service port.
	Extport pulumi.StringPtrInput
	// Custom defined id.
	Fosid pulumi.IntPtrInput
	// Load balance method. Valid values: `static`, `round-robin`, `weighted`, `least-session`, `least-rtt`, `first-alive`.
	LdbMethod pulumi.StringPtrInput
	// Start-mapped-IP [-end-mapped-IP].
	Mappedip pulumi.StringInput
	// Mapped service port.
	Mappedport pulumi.StringPtrInput
	// Health monitors. The structure of `monitor` block is documented below.
	Monitors Vip64MonitorArrayInput
	// VIP64 name.
	Name pulumi.StringPtrInput
	// Enable port forwarding. Valid values: `disable`, `enable`.
	Portforward pulumi.StringPtrInput
	// Mapped port protocol. Valid values: `tcp`, `udp`.
	Protocol pulumi.StringPtrInput
	// Real servers. The structure of `realservers` block is documented below.
	Realservers Vip64RealserverArrayInput
	// Server type. Valid values: `http`, `tcp`, `udp`, `ip`.
	ServerType pulumi.StringPtrInput
	// Source IP6 filter (x:x:x:x:x:x:x:x/x). The structure of `srcFilter` block is documented below.
	SrcFilters Vip64SrcFilterArrayInput
	// VIP type: static NAT or server load balance. Valid values: `static-nat`, `server-load-balance`.
	Type pulumi.StringPtrInput
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (Vip64Args) ElementType() reflect.Type {
	return reflect.TypeOf((*vip64Args)(nil)).Elem()
}

type Vip64Input interface {
	pulumi.Input

	ToVip64Output() Vip64Output
	ToVip64OutputWithContext(ctx context.Context) Vip64Output
}

func (*Vip64) ElementType() reflect.Type {
	return reflect.TypeOf((**Vip64)(nil)).Elem()
}

func (i *Vip64) ToVip64Output() Vip64Output {
	return i.ToVip64OutputWithContext(context.Background())
}

func (i *Vip64) ToVip64OutputWithContext(ctx context.Context) Vip64Output {
	return pulumi.ToOutputWithContext(ctx, i).(Vip64Output)
}

// Vip64ArrayInput is an input type that accepts Vip64Array and Vip64ArrayOutput values.
// You can construct a concrete instance of `Vip64ArrayInput` via:
//
//	Vip64Array{ Vip64Args{...} }
type Vip64ArrayInput interface {
	pulumi.Input

	ToVip64ArrayOutput() Vip64ArrayOutput
	ToVip64ArrayOutputWithContext(context.Context) Vip64ArrayOutput
}

type Vip64Array []Vip64Input

func (Vip64Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Vip64)(nil)).Elem()
}

func (i Vip64Array) ToVip64ArrayOutput() Vip64ArrayOutput {
	return i.ToVip64ArrayOutputWithContext(context.Background())
}

func (i Vip64Array) ToVip64ArrayOutputWithContext(ctx context.Context) Vip64ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Vip64ArrayOutput)
}

// Vip64MapInput is an input type that accepts Vip64Map and Vip64MapOutput values.
// You can construct a concrete instance of `Vip64MapInput` via:
//
//	Vip64Map{ "key": Vip64Args{...} }
type Vip64MapInput interface {
	pulumi.Input

	ToVip64MapOutput() Vip64MapOutput
	ToVip64MapOutputWithContext(context.Context) Vip64MapOutput
}

type Vip64Map map[string]Vip64Input

func (Vip64Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Vip64)(nil)).Elem()
}

func (i Vip64Map) ToVip64MapOutput() Vip64MapOutput {
	return i.ToVip64MapOutputWithContext(context.Background())
}

func (i Vip64Map) ToVip64MapOutputWithContext(ctx context.Context) Vip64MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Vip64MapOutput)
}

type Vip64Output struct{ *pulumi.OutputState }

func (Vip64Output) ElementType() reflect.Type {
	return reflect.TypeOf((**Vip64)(nil)).Elem()
}

func (o Vip64Output) ToVip64Output() Vip64Output {
	return o
}

func (o Vip64Output) ToVip64OutputWithContext(ctx context.Context) Vip64Output {
	return o
}

// Enable ARP reply. Valid values: `disable`, `enable`.
func (o Vip64Output) ArpReply() pulumi.StringOutput {
	return o.ApplyT(func(v *Vip64) pulumi.StringOutput { return v.ArpReply }).(pulumi.StringOutput)
}

// Color of icon on the GUI.
func (o Vip64Output) Color() pulumi.IntOutput {
	return o.ApplyT(func(v *Vip64) pulumi.IntOutput { return v.Color }).(pulumi.IntOutput)
}

// Comment.
func (o Vip64Output) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Vip64) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o Vip64Output) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Vip64) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Start-external-IP [-end-external-IP].
func (o Vip64Output) Extip() pulumi.StringOutput {
	return o.ApplyT(func(v *Vip64) pulumi.StringOutput { return v.Extip }).(pulumi.StringOutput)
}

// External service port.
func (o Vip64Output) Extport() pulumi.StringOutput {
	return o.ApplyT(func(v *Vip64) pulumi.StringOutput { return v.Extport }).(pulumi.StringOutput)
}

// Custom defined id.
func (o Vip64Output) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *Vip64) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

// Load balance method. Valid values: `static`, `round-robin`, `weighted`, `least-session`, `least-rtt`, `first-alive`.
func (o Vip64Output) LdbMethod() pulumi.StringOutput {
	return o.ApplyT(func(v *Vip64) pulumi.StringOutput { return v.LdbMethod }).(pulumi.StringOutput)
}

// Start-mapped-IP [-end-mapped-IP].
func (o Vip64Output) Mappedip() pulumi.StringOutput {
	return o.ApplyT(func(v *Vip64) pulumi.StringOutput { return v.Mappedip }).(pulumi.StringOutput)
}

// Mapped service port.
func (o Vip64Output) Mappedport() pulumi.StringOutput {
	return o.ApplyT(func(v *Vip64) pulumi.StringOutput { return v.Mappedport }).(pulumi.StringOutput)
}

// Health monitors. The structure of `monitor` block is documented below.
func (o Vip64Output) Monitors() Vip64MonitorArrayOutput {
	return o.ApplyT(func(v *Vip64) Vip64MonitorArrayOutput { return v.Monitors }).(Vip64MonitorArrayOutput)
}

// VIP64 name.
func (o Vip64Output) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Vip64) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Enable port forwarding. Valid values: `disable`, `enable`.
func (o Vip64Output) Portforward() pulumi.StringOutput {
	return o.ApplyT(func(v *Vip64) pulumi.StringOutput { return v.Portforward }).(pulumi.StringOutput)
}

// Mapped port protocol. Valid values: `tcp`, `udp`.
func (o Vip64Output) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v *Vip64) pulumi.StringOutput { return v.Protocol }).(pulumi.StringOutput)
}

// Real servers. The structure of `realservers` block is documented below.
func (o Vip64Output) Realservers() Vip64RealserverArrayOutput {
	return o.ApplyT(func(v *Vip64) Vip64RealserverArrayOutput { return v.Realservers }).(Vip64RealserverArrayOutput)
}

// Server type. Valid values: `http`, `tcp`, `udp`, `ip`.
func (o Vip64Output) ServerType() pulumi.StringOutput {
	return o.ApplyT(func(v *Vip64) pulumi.StringOutput { return v.ServerType }).(pulumi.StringOutput)
}

// Source IP6 filter (x:x:x:x:x:x:x:x/x). The structure of `srcFilter` block is documented below.
func (o Vip64Output) SrcFilters() Vip64SrcFilterArrayOutput {
	return o.ApplyT(func(v *Vip64) Vip64SrcFilterArrayOutput { return v.SrcFilters }).(Vip64SrcFilterArrayOutput)
}

// VIP type: static NAT or server load balance. Valid values: `static-nat`, `server-load-balance`.
func (o Vip64Output) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Vip64) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
func (o Vip64Output) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v *Vip64) pulumi.StringOutput { return v.Uuid }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o Vip64Output) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Vip64) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type Vip64ArrayOutput struct{ *pulumi.OutputState }

func (Vip64ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Vip64)(nil)).Elem()
}

func (o Vip64ArrayOutput) ToVip64ArrayOutput() Vip64ArrayOutput {
	return o
}

func (o Vip64ArrayOutput) ToVip64ArrayOutputWithContext(ctx context.Context) Vip64ArrayOutput {
	return o
}

func (o Vip64ArrayOutput) Index(i pulumi.IntInput) Vip64Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Vip64 {
		return vs[0].([]*Vip64)[vs[1].(int)]
	}).(Vip64Output)
}

type Vip64MapOutput struct{ *pulumi.OutputState }

func (Vip64MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Vip64)(nil)).Elem()
}

func (o Vip64MapOutput) ToVip64MapOutput() Vip64MapOutput {
	return o
}

func (o Vip64MapOutput) ToVip64MapOutputWithContext(ctx context.Context) Vip64MapOutput {
	return o
}

func (o Vip64MapOutput) MapIndex(k pulumi.StringInput) Vip64Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Vip64 {
		return vs[0].(map[string]*Vip64)[vs[1].(string)]
	}).(Vip64Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*Vip64Input)(nil)).Elem(), &Vip64{})
	pulumi.RegisterInputType(reflect.TypeOf((*Vip64ArrayInput)(nil)).Elem(), Vip64Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*Vip64MapInput)(nil)).Elem(), Vip64Map{})
	pulumi.RegisterOutputType(Vip64Output{})
	pulumi.RegisterOutputType(Vip64ArrayOutput{})
	pulumi.RegisterOutputType(Vip64MapOutput{})
}
