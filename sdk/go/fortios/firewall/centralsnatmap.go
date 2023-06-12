// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firewall

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure central SNAT policies.
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
//			_, err := firewall.NewCentralsnatmap(ctx, "trname", &firewall.CentralsnatmapArgs{
//				DstAddrs: firewall.CentralsnatmapDstAddrArray{
//					&firewall.CentralsnatmapDstAddrArgs{
//						Name: pulumi.String("all"),
//					},
//				},
//				Dstintfs: firewall.CentralsnatmapDstintfArray{
//					&firewall.CentralsnatmapDstintfArgs{
//						Name: pulumi.String("port3"),
//					},
//				},
//				Nat:     pulumi.String("enable"),
//				NatPort: pulumi.String("0"),
//				OrigAddrs: firewall.CentralsnatmapOrigAddrArray{
//					&firewall.CentralsnatmapOrigAddrArgs{
//						Name: pulumi.String("all"),
//					},
//				},
//				OrigPort: pulumi.String("0"),
//				Policyid: pulumi.Int(1),
//				Protocol: pulumi.Int(33),
//				Srcintfs: firewall.CentralsnatmapSrcintfArray{
//					&firewall.CentralsnatmapSrcintfArgs{
//						Name: pulumi.String("port1"),
//					},
//				},
//				Status: pulumi.String("enable"),
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
// # Firewall CentralSnatMap can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:firewall/centralsnatmap:Centralsnatmap labelname {{policyid}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:firewall/centralsnatmap:Centralsnatmap labelname {{policyid}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type Centralsnatmap struct {
	pulumi.CustomResourceState

	// Comment.
	Comments pulumi.StringPtrOutput `pulumi:"comments"`
	// IPv6 Destination address. The structure of `dstAddr6` block is documented below.
	DstAddr6s CentralsnatmapDstAddr6ArrayOutput `pulumi:"dstAddr6s"`
	// Destination address name from available addresses. The structure of `dstAddr` block is documented below.
	DstAddrs CentralsnatmapDstAddrArrayOutput `pulumi:"dstAddrs"`
	// Destination interface name from available interfaces. The structure of `dstintf` block is documented below.
	Dstintfs CentralsnatmapDstintfArrayOutput `pulumi:"dstintfs"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Enable/disable source NAT. Valid values: `disable`, `enable`.
	Nat pulumi.StringOutput `pulumi:"nat"`
	// Enable/disable NAT46. Valid values: `enable`, `disable`.
	Nat46 pulumi.StringOutput `pulumi:"nat46"`
	// Enable/disable NAT64. Valid values: `enable`, `disable`.
	Nat64 pulumi.StringOutput `pulumi:"nat64"`
	// IPv6 pools to be used for source NAT. The structure of `natIppool6` block is documented below.
	NatIppool6s CentralsnatmapNatIppool6ArrayOutput `pulumi:"natIppool6s"`
	// Name of the IP pools to be used to translate addresses from available IP Pools. The structure of `natIppool` block is documented below.
	NatIppools CentralsnatmapNatIppoolArrayOutput `pulumi:"natIppools"`
	// Translated port or port range (0 to 65535).
	NatPort pulumi.StringOutput `pulumi:"natPort"`
	// IPv6 Original address. The structure of `origAddr6` block is documented below.
	OrigAddr6s CentralsnatmapOrigAddr6ArrayOutput `pulumi:"origAddr6s"`
	// Original address. The structure of `origAddr` block is documented below.
	OrigAddrs CentralsnatmapOrigAddrArrayOutput `pulumi:"origAddrs"`
	// Original TCP port (0 to 65535).
	OrigPort pulumi.StringOutput `pulumi:"origPort"`
	// Policy ID.
	Policyid pulumi.IntOutput `pulumi:"policyid"`
	// Integer value for the protocol type (0 - 255).
	Protocol pulumi.IntOutput `pulumi:"protocol"`
	// Source interface name from available interfaces. The structure of `srcintf` block is documented below.
	Srcintfs CentralsnatmapSrcintfArrayOutput `pulumi:"srcintfs"`
	// Enable/disable the active status of this policy. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// IPv4/IPv6 source NAT. Valid values: `ipv4`, `ipv6`.
	Type pulumi.StringOutput `pulumi:"type"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid pulumi.StringOutput `pulumi:"uuid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewCentralsnatmap registers a new resource with the given unique name, arguments, and options.
func NewCentralsnatmap(ctx *pulumi.Context,
	name string, args *CentralsnatmapArgs, opts ...pulumi.ResourceOption) (*Centralsnatmap, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DstAddrs == nil {
		return nil, errors.New("invalid value for required argument 'DstAddrs'")
	}
	if args.Dstintfs == nil {
		return nil, errors.New("invalid value for required argument 'Dstintfs'")
	}
	if args.Nat == nil {
		return nil, errors.New("invalid value for required argument 'Nat'")
	}
	if args.OrigAddrs == nil {
		return nil, errors.New("invalid value for required argument 'OrigAddrs'")
	}
	if args.OrigPort == nil {
		return nil, errors.New("invalid value for required argument 'OrigPort'")
	}
	if args.Protocol == nil {
		return nil, errors.New("invalid value for required argument 'Protocol'")
	}
	if args.Srcintfs == nil {
		return nil, errors.New("invalid value for required argument 'Srcintfs'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Centralsnatmap
	err := ctx.RegisterResource("fortios:firewall/centralsnatmap:Centralsnatmap", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCentralsnatmap gets an existing Centralsnatmap resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCentralsnatmap(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CentralsnatmapState, opts ...pulumi.ResourceOption) (*Centralsnatmap, error) {
	var resource Centralsnatmap
	err := ctx.ReadResource("fortios:firewall/centralsnatmap:Centralsnatmap", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Centralsnatmap resources.
type centralsnatmapState struct {
	// Comment.
	Comments *string `pulumi:"comments"`
	// IPv6 Destination address. The structure of `dstAddr6` block is documented below.
	DstAddr6s []CentralsnatmapDstAddr6 `pulumi:"dstAddr6s"`
	// Destination address name from available addresses. The structure of `dstAddr` block is documented below.
	DstAddrs []CentralsnatmapDstAddr `pulumi:"dstAddrs"`
	// Destination interface name from available interfaces. The structure of `dstintf` block is documented below.
	Dstintfs []CentralsnatmapDstintf `pulumi:"dstintfs"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Enable/disable source NAT. Valid values: `disable`, `enable`.
	Nat *string `pulumi:"nat"`
	// Enable/disable NAT46. Valid values: `enable`, `disable`.
	Nat46 *string `pulumi:"nat46"`
	// Enable/disable NAT64. Valid values: `enable`, `disable`.
	Nat64 *string `pulumi:"nat64"`
	// IPv6 pools to be used for source NAT. The structure of `natIppool6` block is documented below.
	NatIppool6s []CentralsnatmapNatIppool6 `pulumi:"natIppool6s"`
	// Name of the IP pools to be used to translate addresses from available IP Pools. The structure of `natIppool` block is documented below.
	NatIppools []CentralsnatmapNatIppool `pulumi:"natIppools"`
	// Translated port or port range (0 to 65535).
	NatPort *string `pulumi:"natPort"`
	// IPv6 Original address. The structure of `origAddr6` block is documented below.
	OrigAddr6s []CentralsnatmapOrigAddr6 `pulumi:"origAddr6s"`
	// Original address. The structure of `origAddr` block is documented below.
	OrigAddrs []CentralsnatmapOrigAddr `pulumi:"origAddrs"`
	// Original TCP port (0 to 65535).
	OrigPort *string `pulumi:"origPort"`
	// Policy ID.
	Policyid *int `pulumi:"policyid"`
	// Integer value for the protocol type (0 - 255).
	Protocol *int `pulumi:"protocol"`
	// Source interface name from available interfaces. The structure of `srcintf` block is documented below.
	Srcintfs []CentralsnatmapSrcintf `pulumi:"srcintfs"`
	// Enable/disable the active status of this policy. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// IPv4/IPv6 source NAT. Valid values: `ipv4`, `ipv6`.
	Type *string `pulumi:"type"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid *string `pulumi:"uuid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type CentralsnatmapState struct {
	// Comment.
	Comments pulumi.StringPtrInput
	// IPv6 Destination address. The structure of `dstAddr6` block is documented below.
	DstAddr6s CentralsnatmapDstAddr6ArrayInput
	// Destination address name from available addresses. The structure of `dstAddr` block is documented below.
	DstAddrs CentralsnatmapDstAddrArrayInput
	// Destination interface name from available interfaces. The structure of `dstintf` block is documented below.
	Dstintfs CentralsnatmapDstintfArrayInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Enable/disable source NAT. Valid values: `disable`, `enable`.
	Nat pulumi.StringPtrInput
	// Enable/disable NAT46. Valid values: `enable`, `disable`.
	Nat46 pulumi.StringPtrInput
	// Enable/disable NAT64. Valid values: `enable`, `disable`.
	Nat64 pulumi.StringPtrInput
	// IPv6 pools to be used for source NAT. The structure of `natIppool6` block is documented below.
	NatIppool6s CentralsnatmapNatIppool6ArrayInput
	// Name of the IP pools to be used to translate addresses from available IP Pools. The structure of `natIppool` block is documented below.
	NatIppools CentralsnatmapNatIppoolArrayInput
	// Translated port or port range (0 to 65535).
	NatPort pulumi.StringPtrInput
	// IPv6 Original address. The structure of `origAddr6` block is documented below.
	OrigAddr6s CentralsnatmapOrigAddr6ArrayInput
	// Original address. The structure of `origAddr` block is documented below.
	OrigAddrs CentralsnatmapOrigAddrArrayInput
	// Original TCP port (0 to 65535).
	OrigPort pulumi.StringPtrInput
	// Policy ID.
	Policyid pulumi.IntPtrInput
	// Integer value for the protocol type (0 - 255).
	Protocol pulumi.IntPtrInput
	// Source interface name from available interfaces. The structure of `srcintf` block is documented below.
	Srcintfs CentralsnatmapSrcintfArrayInput
	// Enable/disable the active status of this policy. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// IPv4/IPv6 source NAT. Valid values: `ipv4`, `ipv6`.
	Type pulumi.StringPtrInput
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (CentralsnatmapState) ElementType() reflect.Type {
	return reflect.TypeOf((*centralsnatmapState)(nil)).Elem()
}

type centralsnatmapArgs struct {
	// Comment.
	Comments *string `pulumi:"comments"`
	// IPv6 Destination address. The structure of `dstAddr6` block is documented below.
	DstAddr6s []CentralsnatmapDstAddr6 `pulumi:"dstAddr6s"`
	// Destination address name from available addresses. The structure of `dstAddr` block is documented below.
	DstAddrs []CentralsnatmapDstAddr `pulumi:"dstAddrs"`
	// Destination interface name from available interfaces. The structure of `dstintf` block is documented below.
	Dstintfs []CentralsnatmapDstintf `pulumi:"dstintfs"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Enable/disable source NAT. Valid values: `disable`, `enable`.
	Nat string `pulumi:"nat"`
	// Enable/disable NAT46. Valid values: `enable`, `disable`.
	Nat46 *string `pulumi:"nat46"`
	// Enable/disable NAT64. Valid values: `enable`, `disable`.
	Nat64 *string `pulumi:"nat64"`
	// IPv6 pools to be used for source NAT. The structure of `natIppool6` block is documented below.
	NatIppool6s []CentralsnatmapNatIppool6 `pulumi:"natIppool6s"`
	// Name of the IP pools to be used to translate addresses from available IP Pools. The structure of `natIppool` block is documented below.
	NatIppools []CentralsnatmapNatIppool `pulumi:"natIppools"`
	// Translated port or port range (0 to 65535).
	NatPort *string `pulumi:"natPort"`
	// IPv6 Original address. The structure of `origAddr6` block is documented below.
	OrigAddr6s []CentralsnatmapOrigAddr6 `pulumi:"origAddr6s"`
	// Original address. The structure of `origAddr` block is documented below.
	OrigAddrs []CentralsnatmapOrigAddr `pulumi:"origAddrs"`
	// Original TCP port (0 to 65535).
	OrigPort string `pulumi:"origPort"`
	// Policy ID.
	Policyid *int `pulumi:"policyid"`
	// Integer value for the protocol type (0 - 255).
	Protocol int `pulumi:"protocol"`
	// Source interface name from available interfaces. The structure of `srcintf` block is documented below.
	Srcintfs []CentralsnatmapSrcintf `pulumi:"srcintfs"`
	// Enable/disable the active status of this policy. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// IPv4/IPv6 source NAT. Valid values: `ipv4`, `ipv6`.
	Type *string `pulumi:"type"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid *string `pulumi:"uuid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Centralsnatmap resource.
type CentralsnatmapArgs struct {
	// Comment.
	Comments pulumi.StringPtrInput
	// IPv6 Destination address. The structure of `dstAddr6` block is documented below.
	DstAddr6s CentralsnatmapDstAddr6ArrayInput
	// Destination address name from available addresses. The structure of `dstAddr` block is documented below.
	DstAddrs CentralsnatmapDstAddrArrayInput
	// Destination interface name from available interfaces. The structure of `dstintf` block is documented below.
	Dstintfs CentralsnatmapDstintfArrayInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Enable/disable source NAT. Valid values: `disable`, `enable`.
	Nat pulumi.StringInput
	// Enable/disable NAT46. Valid values: `enable`, `disable`.
	Nat46 pulumi.StringPtrInput
	// Enable/disable NAT64. Valid values: `enable`, `disable`.
	Nat64 pulumi.StringPtrInput
	// IPv6 pools to be used for source NAT. The structure of `natIppool6` block is documented below.
	NatIppool6s CentralsnatmapNatIppool6ArrayInput
	// Name of the IP pools to be used to translate addresses from available IP Pools. The structure of `natIppool` block is documented below.
	NatIppools CentralsnatmapNatIppoolArrayInput
	// Translated port or port range (0 to 65535).
	NatPort pulumi.StringPtrInput
	// IPv6 Original address. The structure of `origAddr6` block is documented below.
	OrigAddr6s CentralsnatmapOrigAddr6ArrayInput
	// Original address. The structure of `origAddr` block is documented below.
	OrigAddrs CentralsnatmapOrigAddrArrayInput
	// Original TCP port (0 to 65535).
	OrigPort pulumi.StringInput
	// Policy ID.
	Policyid pulumi.IntPtrInput
	// Integer value for the protocol type (0 - 255).
	Protocol pulumi.IntInput
	// Source interface name from available interfaces. The structure of `srcintf` block is documented below.
	Srcintfs CentralsnatmapSrcintfArrayInput
	// Enable/disable the active status of this policy. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// IPv4/IPv6 source NAT. Valid values: `ipv4`, `ipv6`.
	Type pulumi.StringPtrInput
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (CentralsnatmapArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*centralsnatmapArgs)(nil)).Elem()
}

type CentralsnatmapInput interface {
	pulumi.Input

	ToCentralsnatmapOutput() CentralsnatmapOutput
	ToCentralsnatmapOutputWithContext(ctx context.Context) CentralsnatmapOutput
}

func (*Centralsnatmap) ElementType() reflect.Type {
	return reflect.TypeOf((**Centralsnatmap)(nil)).Elem()
}

func (i *Centralsnatmap) ToCentralsnatmapOutput() CentralsnatmapOutput {
	return i.ToCentralsnatmapOutputWithContext(context.Background())
}

func (i *Centralsnatmap) ToCentralsnatmapOutputWithContext(ctx context.Context) CentralsnatmapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CentralsnatmapOutput)
}

// CentralsnatmapArrayInput is an input type that accepts CentralsnatmapArray and CentralsnatmapArrayOutput values.
// You can construct a concrete instance of `CentralsnatmapArrayInput` via:
//
//	CentralsnatmapArray{ CentralsnatmapArgs{...} }
type CentralsnatmapArrayInput interface {
	pulumi.Input

	ToCentralsnatmapArrayOutput() CentralsnatmapArrayOutput
	ToCentralsnatmapArrayOutputWithContext(context.Context) CentralsnatmapArrayOutput
}

type CentralsnatmapArray []CentralsnatmapInput

func (CentralsnatmapArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Centralsnatmap)(nil)).Elem()
}

func (i CentralsnatmapArray) ToCentralsnatmapArrayOutput() CentralsnatmapArrayOutput {
	return i.ToCentralsnatmapArrayOutputWithContext(context.Background())
}

func (i CentralsnatmapArray) ToCentralsnatmapArrayOutputWithContext(ctx context.Context) CentralsnatmapArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CentralsnatmapArrayOutput)
}

// CentralsnatmapMapInput is an input type that accepts CentralsnatmapMap and CentralsnatmapMapOutput values.
// You can construct a concrete instance of `CentralsnatmapMapInput` via:
//
//	CentralsnatmapMap{ "key": CentralsnatmapArgs{...} }
type CentralsnatmapMapInput interface {
	pulumi.Input

	ToCentralsnatmapMapOutput() CentralsnatmapMapOutput
	ToCentralsnatmapMapOutputWithContext(context.Context) CentralsnatmapMapOutput
}

type CentralsnatmapMap map[string]CentralsnatmapInput

func (CentralsnatmapMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Centralsnatmap)(nil)).Elem()
}

func (i CentralsnatmapMap) ToCentralsnatmapMapOutput() CentralsnatmapMapOutput {
	return i.ToCentralsnatmapMapOutputWithContext(context.Background())
}

func (i CentralsnatmapMap) ToCentralsnatmapMapOutputWithContext(ctx context.Context) CentralsnatmapMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CentralsnatmapMapOutput)
}

type CentralsnatmapOutput struct{ *pulumi.OutputState }

func (CentralsnatmapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Centralsnatmap)(nil)).Elem()
}

func (o CentralsnatmapOutput) ToCentralsnatmapOutput() CentralsnatmapOutput {
	return o
}

func (o CentralsnatmapOutput) ToCentralsnatmapOutputWithContext(ctx context.Context) CentralsnatmapOutput {
	return o
}

// Comment.
func (o CentralsnatmapOutput) Comments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Centralsnatmap) pulumi.StringPtrOutput { return v.Comments }).(pulumi.StringPtrOutput)
}

// IPv6 Destination address. The structure of `dstAddr6` block is documented below.
func (o CentralsnatmapOutput) DstAddr6s() CentralsnatmapDstAddr6ArrayOutput {
	return o.ApplyT(func(v *Centralsnatmap) CentralsnatmapDstAddr6ArrayOutput { return v.DstAddr6s }).(CentralsnatmapDstAddr6ArrayOutput)
}

// Destination address name from available addresses. The structure of `dstAddr` block is documented below.
func (o CentralsnatmapOutput) DstAddrs() CentralsnatmapDstAddrArrayOutput {
	return o.ApplyT(func(v *Centralsnatmap) CentralsnatmapDstAddrArrayOutput { return v.DstAddrs }).(CentralsnatmapDstAddrArrayOutput)
}

// Destination interface name from available interfaces. The structure of `dstintf` block is documented below.
func (o CentralsnatmapOutput) Dstintfs() CentralsnatmapDstintfArrayOutput {
	return o.ApplyT(func(v *Centralsnatmap) CentralsnatmapDstintfArrayOutput { return v.Dstintfs }).(CentralsnatmapDstintfArrayOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o CentralsnatmapOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Centralsnatmap) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Enable/disable source NAT. Valid values: `disable`, `enable`.
func (o CentralsnatmapOutput) Nat() pulumi.StringOutput {
	return o.ApplyT(func(v *Centralsnatmap) pulumi.StringOutput { return v.Nat }).(pulumi.StringOutput)
}

// Enable/disable NAT46. Valid values: `enable`, `disable`.
func (o CentralsnatmapOutput) Nat46() pulumi.StringOutput {
	return o.ApplyT(func(v *Centralsnatmap) pulumi.StringOutput { return v.Nat46 }).(pulumi.StringOutput)
}

// Enable/disable NAT64. Valid values: `enable`, `disable`.
func (o CentralsnatmapOutput) Nat64() pulumi.StringOutput {
	return o.ApplyT(func(v *Centralsnatmap) pulumi.StringOutput { return v.Nat64 }).(pulumi.StringOutput)
}

// IPv6 pools to be used for source NAT. The structure of `natIppool6` block is documented below.
func (o CentralsnatmapOutput) NatIppool6s() CentralsnatmapNatIppool6ArrayOutput {
	return o.ApplyT(func(v *Centralsnatmap) CentralsnatmapNatIppool6ArrayOutput { return v.NatIppool6s }).(CentralsnatmapNatIppool6ArrayOutput)
}

// Name of the IP pools to be used to translate addresses from available IP Pools. The structure of `natIppool` block is documented below.
func (o CentralsnatmapOutput) NatIppools() CentralsnatmapNatIppoolArrayOutput {
	return o.ApplyT(func(v *Centralsnatmap) CentralsnatmapNatIppoolArrayOutput { return v.NatIppools }).(CentralsnatmapNatIppoolArrayOutput)
}

// Translated port or port range (0 to 65535).
func (o CentralsnatmapOutput) NatPort() pulumi.StringOutput {
	return o.ApplyT(func(v *Centralsnatmap) pulumi.StringOutput { return v.NatPort }).(pulumi.StringOutput)
}

// IPv6 Original address. The structure of `origAddr6` block is documented below.
func (o CentralsnatmapOutput) OrigAddr6s() CentralsnatmapOrigAddr6ArrayOutput {
	return o.ApplyT(func(v *Centralsnatmap) CentralsnatmapOrigAddr6ArrayOutput { return v.OrigAddr6s }).(CentralsnatmapOrigAddr6ArrayOutput)
}

// Original address. The structure of `origAddr` block is documented below.
func (o CentralsnatmapOutput) OrigAddrs() CentralsnatmapOrigAddrArrayOutput {
	return o.ApplyT(func(v *Centralsnatmap) CentralsnatmapOrigAddrArrayOutput { return v.OrigAddrs }).(CentralsnatmapOrigAddrArrayOutput)
}

// Original TCP port (0 to 65535).
func (o CentralsnatmapOutput) OrigPort() pulumi.StringOutput {
	return o.ApplyT(func(v *Centralsnatmap) pulumi.StringOutput { return v.OrigPort }).(pulumi.StringOutput)
}

// Policy ID.
func (o CentralsnatmapOutput) Policyid() pulumi.IntOutput {
	return o.ApplyT(func(v *Centralsnatmap) pulumi.IntOutput { return v.Policyid }).(pulumi.IntOutput)
}

// Integer value for the protocol type (0 - 255).
func (o CentralsnatmapOutput) Protocol() pulumi.IntOutput {
	return o.ApplyT(func(v *Centralsnatmap) pulumi.IntOutput { return v.Protocol }).(pulumi.IntOutput)
}

// Source interface name from available interfaces. The structure of `srcintf` block is documented below.
func (o CentralsnatmapOutput) Srcintfs() CentralsnatmapSrcintfArrayOutput {
	return o.ApplyT(func(v *Centralsnatmap) CentralsnatmapSrcintfArrayOutput { return v.Srcintfs }).(CentralsnatmapSrcintfArrayOutput)
}

// Enable/disable the active status of this policy. Valid values: `enable`, `disable`.
func (o CentralsnatmapOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Centralsnatmap) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// IPv4/IPv6 source NAT. Valid values: `ipv4`, `ipv6`.
func (o CentralsnatmapOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Centralsnatmap) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
func (o CentralsnatmapOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v *Centralsnatmap) pulumi.StringOutput { return v.Uuid }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o CentralsnatmapOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Centralsnatmap) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type CentralsnatmapArrayOutput struct{ *pulumi.OutputState }

func (CentralsnatmapArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Centralsnatmap)(nil)).Elem()
}

func (o CentralsnatmapArrayOutput) ToCentralsnatmapArrayOutput() CentralsnatmapArrayOutput {
	return o
}

func (o CentralsnatmapArrayOutput) ToCentralsnatmapArrayOutputWithContext(ctx context.Context) CentralsnatmapArrayOutput {
	return o
}

func (o CentralsnatmapArrayOutput) Index(i pulumi.IntInput) CentralsnatmapOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Centralsnatmap {
		return vs[0].([]*Centralsnatmap)[vs[1].(int)]
	}).(CentralsnatmapOutput)
}

type CentralsnatmapMapOutput struct{ *pulumi.OutputState }

func (CentralsnatmapMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Centralsnatmap)(nil)).Elem()
}

func (o CentralsnatmapMapOutput) ToCentralsnatmapMapOutput() CentralsnatmapMapOutput {
	return o
}

func (o CentralsnatmapMapOutput) ToCentralsnatmapMapOutputWithContext(ctx context.Context) CentralsnatmapMapOutput {
	return o
}

func (o CentralsnatmapMapOutput) MapIndex(k pulumi.StringInput) CentralsnatmapOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Centralsnatmap {
		return vs[0].(map[string]*Centralsnatmap)[vs[1].(string)]
	}).(CentralsnatmapOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CentralsnatmapInput)(nil)).Elem(), &Centralsnatmap{})
	pulumi.RegisterInputType(reflect.TypeOf((*CentralsnatmapArrayInput)(nil)).Elem(), CentralsnatmapArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CentralsnatmapMapInput)(nil)).Elem(), CentralsnatmapMap{})
	pulumi.RegisterOutputType(CentralsnatmapOutput{})
	pulumi.RegisterOutputType(CentralsnatmapArrayOutput{})
	pulumi.RegisterOutputType(CentralsnatmapMapOutput{})
}
