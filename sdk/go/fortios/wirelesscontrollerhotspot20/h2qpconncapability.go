// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package wirelesscontrollerhotspot20

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure connection capability.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/wirelesscontrollerhotspot20"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := wirelesscontrollerhotspot20.NewH2qpconncapability(ctx, "trname", &wirelesscontrollerhotspot20.H2qpconncapabilityArgs{
//				EspPort:     pulumi.String("unknown"),
//				FtpPort:     pulumi.String("unknown"),
//				HttpPort:    pulumi.String("unknown"),
//				IcmpPort:    pulumi.String("closed"),
//				Ikev2Port:   pulumi.String("unknown"),
//				Ikev2XxPort: pulumi.String("unknown"),
//				PptpVpnPort: pulumi.String("unknown"),
//				SshPort:     pulumi.String("unknown"),
//				TlsPort:     pulumi.String("unknown"),
//				VoipTcpPort: pulumi.String("unknown"),
//				VoipUdpPort: pulumi.String("unknown"),
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
// # WirelessControllerHotspot20 H2QpConnCapability can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:wirelesscontrollerhotspot20/h2qpconncapability:H2qpconncapability labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:wirelesscontrollerhotspot20/h2qpconncapability:H2qpconncapability labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type H2qpconncapability struct {
	pulumi.CustomResourceState

	// Set ESP port service (used by IPsec VPNs) status. Valid values: `closed`, `open`, `unknown`.
	EspPort pulumi.StringOutput `pulumi:"espPort"`
	// Set FTP port service status. Valid values: `closed`, `open`, `unknown`.
	FtpPort pulumi.StringOutput `pulumi:"ftpPort"`
	// Set HTTP port service status. Valid values: `closed`, `open`, `unknown`.
	HttpPort pulumi.StringOutput `pulumi:"httpPort"`
	// Set ICMP port service status. Valid values: `closed`, `open`, `unknown`.
	IcmpPort pulumi.StringOutput `pulumi:"icmpPort"`
	// Set IKEv2 port service for IPsec VPN status. Valid values: `closed`, `open`, `unknown`.
	Ikev2Port pulumi.StringOutput `pulumi:"ikev2Port"`
	// Set UDP port 4500 (which may be used by IKEv2 for IPsec VPN) service status. Valid values: `closed`, `open`, `unknown`.
	Ikev2XxPort pulumi.StringOutput `pulumi:"ikev2XxPort"`
	// Connection capability name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Set Point to Point Tunneling Protocol (PPTP) VPN port service status. Valid values: `closed`, `open`, `unknown`.
	PptpVpnPort pulumi.StringOutput `pulumi:"pptpVpnPort"`
	// Set SSH port service status. Valid values: `closed`, `open`, `unknown`.
	SshPort pulumi.StringOutput `pulumi:"sshPort"`
	// Set TLS VPN (HTTPS) port service status. Valid values: `closed`, `open`, `unknown`.
	TlsPort pulumi.StringOutput `pulumi:"tlsPort"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// Set VoIP TCP port service status. Valid values: `closed`, `open`, `unknown`.
	VoipTcpPort pulumi.StringOutput `pulumi:"voipTcpPort"`
	// Set VoIP UDP port service status. Valid values: `closed`, `open`, `unknown`.
	VoipUdpPort pulumi.StringOutput `pulumi:"voipUdpPort"`
}

// NewH2qpconncapability registers a new resource with the given unique name, arguments, and options.
func NewH2qpconncapability(ctx *pulumi.Context,
	name string, args *H2qpconncapabilityArgs, opts ...pulumi.ResourceOption) (*H2qpconncapability, error) {
	if args == nil {
		args = &H2qpconncapabilityArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource H2qpconncapability
	err := ctx.RegisterResource("fortios:wirelesscontrollerhotspot20/h2qpconncapability:H2qpconncapability", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetH2qpconncapability gets an existing H2qpconncapability resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetH2qpconncapability(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *H2qpconncapabilityState, opts ...pulumi.ResourceOption) (*H2qpconncapability, error) {
	var resource H2qpconncapability
	err := ctx.ReadResource("fortios:wirelesscontrollerhotspot20/h2qpconncapability:H2qpconncapability", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering H2qpconncapability resources.
type h2qpconncapabilityState struct {
	// Set ESP port service (used by IPsec VPNs) status. Valid values: `closed`, `open`, `unknown`.
	EspPort *string `pulumi:"espPort"`
	// Set FTP port service status. Valid values: `closed`, `open`, `unknown`.
	FtpPort *string `pulumi:"ftpPort"`
	// Set HTTP port service status. Valid values: `closed`, `open`, `unknown`.
	HttpPort *string `pulumi:"httpPort"`
	// Set ICMP port service status. Valid values: `closed`, `open`, `unknown`.
	IcmpPort *string `pulumi:"icmpPort"`
	// Set IKEv2 port service for IPsec VPN status. Valid values: `closed`, `open`, `unknown`.
	Ikev2Port *string `pulumi:"ikev2Port"`
	// Set UDP port 4500 (which may be used by IKEv2 for IPsec VPN) service status. Valid values: `closed`, `open`, `unknown`.
	Ikev2XxPort *string `pulumi:"ikev2XxPort"`
	// Connection capability name.
	Name *string `pulumi:"name"`
	// Set Point to Point Tunneling Protocol (PPTP) VPN port service status. Valid values: `closed`, `open`, `unknown`.
	PptpVpnPort *string `pulumi:"pptpVpnPort"`
	// Set SSH port service status. Valid values: `closed`, `open`, `unknown`.
	SshPort *string `pulumi:"sshPort"`
	// Set TLS VPN (HTTPS) port service status. Valid values: `closed`, `open`, `unknown`.
	TlsPort *string `pulumi:"tlsPort"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Set VoIP TCP port service status. Valid values: `closed`, `open`, `unknown`.
	VoipTcpPort *string `pulumi:"voipTcpPort"`
	// Set VoIP UDP port service status. Valid values: `closed`, `open`, `unknown`.
	VoipUdpPort *string `pulumi:"voipUdpPort"`
}

type H2qpconncapabilityState struct {
	// Set ESP port service (used by IPsec VPNs) status. Valid values: `closed`, `open`, `unknown`.
	EspPort pulumi.StringPtrInput
	// Set FTP port service status. Valid values: `closed`, `open`, `unknown`.
	FtpPort pulumi.StringPtrInput
	// Set HTTP port service status. Valid values: `closed`, `open`, `unknown`.
	HttpPort pulumi.StringPtrInput
	// Set ICMP port service status. Valid values: `closed`, `open`, `unknown`.
	IcmpPort pulumi.StringPtrInput
	// Set IKEv2 port service for IPsec VPN status. Valid values: `closed`, `open`, `unknown`.
	Ikev2Port pulumi.StringPtrInput
	// Set UDP port 4500 (which may be used by IKEv2 for IPsec VPN) service status. Valid values: `closed`, `open`, `unknown`.
	Ikev2XxPort pulumi.StringPtrInput
	// Connection capability name.
	Name pulumi.StringPtrInput
	// Set Point to Point Tunneling Protocol (PPTP) VPN port service status. Valid values: `closed`, `open`, `unknown`.
	PptpVpnPort pulumi.StringPtrInput
	// Set SSH port service status. Valid values: `closed`, `open`, `unknown`.
	SshPort pulumi.StringPtrInput
	// Set TLS VPN (HTTPS) port service status. Valid values: `closed`, `open`, `unknown`.
	TlsPort pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Set VoIP TCP port service status. Valid values: `closed`, `open`, `unknown`.
	VoipTcpPort pulumi.StringPtrInput
	// Set VoIP UDP port service status. Valid values: `closed`, `open`, `unknown`.
	VoipUdpPort pulumi.StringPtrInput
}

func (H2qpconncapabilityState) ElementType() reflect.Type {
	return reflect.TypeOf((*h2qpconncapabilityState)(nil)).Elem()
}

type h2qpconncapabilityArgs struct {
	// Set ESP port service (used by IPsec VPNs) status. Valid values: `closed`, `open`, `unknown`.
	EspPort *string `pulumi:"espPort"`
	// Set FTP port service status. Valid values: `closed`, `open`, `unknown`.
	FtpPort *string `pulumi:"ftpPort"`
	// Set HTTP port service status. Valid values: `closed`, `open`, `unknown`.
	HttpPort *string `pulumi:"httpPort"`
	// Set ICMP port service status. Valid values: `closed`, `open`, `unknown`.
	IcmpPort *string `pulumi:"icmpPort"`
	// Set IKEv2 port service for IPsec VPN status. Valid values: `closed`, `open`, `unknown`.
	Ikev2Port *string `pulumi:"ikev2Port"`
	// Set UDP port 4500 (which may be used by IKEv2 for IPsec VPN) service status. Valid values: `closed`, `open`, `unknown`.
	Ikev2XxPort *string `pulumi:"ikev2XxPort"`
	// Connection capability name.
	Name *string `pulumi:"name"`
	// Set Point to Point Tunneling Protocol (PPTP) VPN port service status. Valid values: `closed`, `open`, `unknown`.
	PptpVpnPort *string `pulumi:"pptpVpnPort"`
	// Set SSH port service status. Valid values: `closed`, `open`, `unknown`.
	SshPort *string `pulumi:"sshPort"`
	// Set TLS VPN (HTTPS) port service status. Valid values: `closed`, `open`, `unknown`.
	TlsPort *string `pulumi:"tlsPort"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Set VoIP TCP port service status. Valid values: `closed`, `open`, `unknown`.
	VoipTcpPort *string `pulumi:"voipTcpPort"`
	// Set VoIP UDP port service status. Valid values: `closed`, `open`, `unknown`.
	VoipUdpPort *string `pulumi:"voipUdpPort"`
}

// The set of arguments for constructing a H2qpconncapability resource.
type H2qpconncapabilityArgs struct {
	// Set ESP port service (used by IPsec VPNs) status. Valid values: `closed`, `open`, `unknown`.
	EspPort pulumi.StringPtrInput
	// Set FTP port service status. Valid values: `closed`, `open`, `unknown`.
	FtpPort pulumi.StringPtrInput
	// Set HTTP port service status. Valid values: `closed`, `open`, `unknown`.
	HttpPort pulumi.StringPtrInput
	// Set ICMP port service status. Valid values: `closed`, `open`, `unknown`.
	IcmpPort pulumi.StringPtrInput
	// Set IKEv2 port service for IPsec VPN status. Valid values: `closed`, `open`, `unknown`.
	Ikev2Port pulumi.StringPtrInput
	// Set UDP port 4500 (which may be used by IKEv2 for IPsec VPN) service status. Valid values: `closed`, `open`, `unknown`.
	Ikev2XxPort pulumi.StringPtrInput
	// Connection capability name.
	Name pulumi.StringPtrInput
	// Set Point to Point Tunneling Protocol (PPTP) VPN port service status. Valid values: `closed`, `open`, `unknown`.
	PptpVpnPort pulumi.StringPtrInput
	// Set SSH port service status. Valid values: `closed`, `open`, `unknown`.
	SshPort pulumi.StringPtrInput
	// Set TLS VPN (HTTPS) port service status. Valid values: `closed`, `open`, `unknown`.
	TlsPort pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Set VoIP TCP port service status. Valid values: `closed`, `open`, `unknown`.
	VoipTcpPort pulumi.StringPtrInput
	// Set VoIP UDP port service status. Valid values: `closed`, `open`, `unknown`.
	VoipUdpPort pulumi.StringPtrInput
}

func (H2qpconncapabilityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*h2qpconncapabilityArgs)(nil)).Elem()
}

type H2qpconncapabilityInput interface {
	pulumi.Input

	ToH2qpconncapabilityOutput() H2qpconncapabilityOutput
	ToH2qpconncapabilityOutputWithContext(ctx context.Context) H2qpconncapabilityOutput
}

func (*H2qpconncapability) ElementType() reflect.Type {
	return reflect.TypeOf((**H2qpconncapability)(nil)).Elem()
}

func (i *H2qpconncapability) ToH2qpconncapabilityOutput() H2qpconncapabilityOutput {
	return i.ToH2qpconncapabilityOutputWithContext(context.Background())
}

func (i *H2qpconncapability) ToH2qpconncapabilityOutputWithContext(ctx context.Context) H2qpconncapabilityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(H2qpconncapabilityOutput)
}

// H2qpconncapabilityArrayInput is an input type that accepts H2qpconncapabilityArray and H2qpconncapabilityArrayOutput values.
// You can construct a concrete instance of `H2qpconncapabilityArrayInput` via:
//
//	H2qpconncapabilityArray{ H2qpconncapabilityArgs{...} }
type H2qpconncapabilityArrayInput interface {
	pulumi.Input

	ToH2qpconncapabilityArrayOutput() H2qpconncapabilityArrayOutput
	ToH2qpconncapabilityArrayOutputWithContext(context.Context) H2qpconncapabilityArrayOutput
}

type H2qpconncapabilityArray []H2qpconncapabilityInput

func (H2qpconncapabilityArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*H2qpconncapability)(nil)).Elem()
}

func (i H2qpconncapabilityArray) ToH2qpconncapabilityArrayOutput() H2qpconncapabilityArrayOutput {
	return i.ToH2qpconncapabilityArrayOutputWithContext(context.Background())
}

func (i H2qpconncapabilityArray) ToH2qpconncapabilityArrayOutputWithContext(ctx context.Context) H2qpconncapabilityArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(H2qpconncapabilityArrayOutput)
}

// H2qpconncapabilityMapInput is an input type that accepts H2qpconncapabilityMap and H2qpconncapabilityMapOutput values.
// You can construct a concrete instance of `H2qpconncapabilityMapInput` via:
//
//	H2qpconncapabilityMap{ "key": H2qpconncapabilityArgs{...} }
type H2qpconncapabilityMapInput interface {
	pulumi.Input

	ToH2qpconncapabilityMapOutput() H2qpconncapabilityMapOutput
	ToH2qpconncapabilityMapOutputWithContext(context.Context) H2qpconncapabilityMapOutput
}

type H2qpconncapabilityMap map[string]H2qpconncapabilityInput

func (H2qpconncapabilityMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*H2qpconncapability)(nil)).Elem()
}

func (i H2qpconncapabilityMap) ToH2qpconncapabilityMapOutput() H2qpconncapabilityMapOutput {
	return i.ToH2qpconncapabilityMapOutputWithContext(context.Background())
}

func (i H2qpconncapabilityMap) ToH2qpconncapabilityMapOutputWithContext(ctx context.Context) H2qpconncapabilityMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(H2qpconncapabilityMapOutput)
}

type H2qpconncapabilityOutput struct{ *pulumi.OutputState }

func (H2qpconncapabilityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**H2qpconncapability)(nil)).Elem()
}

func (o H2qpconncapabilityOutput) ToH2qpconncapabilityOutput() H2qpconncapabilityOutput {
	return o
}

func (o H2qpconncapabilityOutput) ToH2qpconncapabilityOutputWithContext(ctx context.Context) H2qpconncapabilityOutput {
	return o
}

// Set ESP port service (used by IPsec VPNs) status. Valid values: `closed`, `open`, `unknown`.
func (o H2qpconncapabilityOutput) EspPort() pulumi.StringOutput {
	return o.ApplyT(func(v *H2qpconncapability) pulumi.StringOutput { return v.EspPort }).(pulumi.StringOutput)
}

// Set FTP port service status. Valid values: `closed`, `open`, `unknown`.
func (o H2qpconncapabilityOutput) FtpPort() pulumi.StringOutput {
	return o.ApplyT(func(v *H2qpconncapability) pulumi.StringOutput { return v.FtpPort }).(pulumi.StringOutput)
}

// Set HTTP port service status. Valid values: `closed`, `open`, `unknown`.
func (o H2qpconncapabilityOutput) HttpPort() pulumi.StringOutput {
	return o.ApplyT(func(v *H2qpconncapability) pulumi.StringOutput { return v.HttpPort }).(pulumi.StringOutput)
}

// Set ICMP port service status. Valid values: `closed`, `open`, `unknown`.
func (o H2qpconncapabilityOutput) IcmpPort() pulumi.StringOutput {
	return o.ApplyT(func(v *H2qpconncapability) pulumi.StringOutput { return v.IcmpPort }).(pulumi.StringOutput)
}

// Set IKEv2 port service for IPsec VPN status. Valid values: `closed`, `open`, `unknown`.
func (o H2qpconncapabilityOutput) Ikev2Port() pulumi.StringOutput {
	return o.ApplyT(func(v *H2qpconncapability) pulumi.StringOutput { return v.Ikev2Port }).(pulumi.StringOutput)
}

// Set UDP port 4500 (which may be used by IKEv2 for IPsec VPN) service status. Valid values: `closed`, `open`, `unknown`.
func (o H2qpconncapabilityOutput) Ikev2XxPort() pulumi.StringOutput {
	return o.ApplyT(func(v *H2qpconncapability) pulumi.StringOutput { return v.Ikev2XxPort }).(pulumi.StringOutput)
}

// Connection capability name.
func (o H2qpconncapabilityOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *H2qpconncapability) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Set Point to Point Tunneling Protocol (PPTP) VPN port service status. Valid values: `closed`, `open`, `unknown`.
func (o H2qpconncapabilityOutput) PptpVpnPort() pulumi.StringOutput {
	return o.ApplyT(func(v *H2qpconncapability) pulumi.StringOutput { return v.PptpVpnPort }).(pulumi.StringOutput)
}

// Set SSH port service status. Valid values: `closed`, `open`, `unknown`.
func (o H2qpconncapabilityOutput) SshPort() pulumi.StringOutput {
	return o.ApplyT(func(v *H2qpconncapability) pulumi.StringOutput { return v.SshPort }).(pulumi.StringOutput)
}

// Set TLS VPN (HTTPS) port service status. Valid values: `closed`, `open`, `unknown`.
func (o H2qpconncapabilityOutput) TlsPort() pulumi.StringOutput {
	return o.ApplyT(func(v *H2qpconncapability) pulumi.StringOutput { return v.TlsPort }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o H2qpconncapabilityOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *H2qpconncapability) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// Set VoIP TCP port service status. Valid values: `closed`, `open`, `unknown`.
func (o H2qpconncapabilityOutput) VoipTcpPort() pulumi.StringOutput {
	return o.ApplyT(func(v *H2qpconncapability) pulumi.StringOutput { return v.VoipTcpPort }).(pulumi.StringOutput)
}

// Set VoIP UDP port service status. Valid values: `closed`, `open`, `unknown`.
func (o H2qpconncapabilityOutput) VoipUdpPort() pulumi.StringOutput {
	return o.ApplyT(func(v *H2qpconncapability) pulumi.StringOutput { return v.VoipUdpPort }).(pulumi.StringOutput)
}

type H2qpconncapabilityArrayOutput struct{ *pulumi.OutputState }

func (H2qpconncapabilityArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*H2qpconncapability)(nil)).Elem()
}

func (o H2qpconncapabilityArrayOutput) ToH2qpconncapabilityArrayOutput() H2qpconncapabilityArrayOutput {
	return o
}

func (o H2qpconncapabilityArrayOutput) ToH2qpconncapabilityArrayOutputWithContext(ctx context.Context) H2qpconncapabilityArrayOutput {
	return o
}

func (o H2qpconncapabilityArrayOutput) Index(i pulumi.IntInput) H2qpconncapabilityOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *H2qpconncapability {
		return vs[0].([]*H2qpconncapability)[vs[1].(int)]
	}).(H2qpconncapabilityOutput)
}

type H2qpconncapabilityMapOutput struct{ *pulumi.OutputState }

func (H2qpconncapabilityMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*H2qpconncapability)(nil)).Elem()
}

func (o H2qpconncapabilityMapOutput) ToH2qpconncapabilityMapOutput() H2qpconncapabilityMapOutput {
	return o
}

func (o H2qpconncapabilityMapOutput) ToH2qpconncapabilityMapOutputWithContext(ctx context.Context) H2qpconncapabilityMapOutput {
	return o
}

func (o H2qpconncapabilityMapOutput) MapIndex(k pulumi.StringInput) H2qpconncapabilityOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *H2qpconncapability {
		return vs[0].(map[string]*H2qpconncapability)[vs[1].(string)]
	}).(H2qpconncapabilityOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*H2qpconncapabilityInput)(nil)).Elem(), &H2qpconncapability{})
	pulumi.RegisterInputType(reflect.TypeOf((*H2qpconncapabilityArrayInput)(nil)).Elem(), H2qpconncapabilityArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*H2qpconncapabilityMapInput)(nil)).Elem(), H2qpconncapabilityMap{})
	pulumi.RegisterOutputType(H2qpconncapabilityOutput{})
	pulumi.RegisterOutputType(H2qpconncapabilityArrayOutput{})
	pulumi.RegisterOutputType(H2qpconncapabilityMapOutput{})
}
