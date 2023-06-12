// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpnssl

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// client Applies to FortiOS Version `>= 7.0.1`.
//
// ## Import
//
// # VpnSsl Client can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:vpnssl/client:Client labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:vpnssl/client:Client labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type Client struct {
	pulumi.CustomResourceState

	// Certificate to offer to SSL-VPN server if it requests one.
	Certificate pulumi.StringOutput `pulumi:"certificate"`
	// Traffic class ID.
	ClassId pulumi.IntOutput `pulumi:"classId"`
	// Comment.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Distance for routes added by SSL-VPN (1 - 255).
	Distance pulumi.IntOutput `pulumi:"distance"`
	// SSL interface to send/receive traffic over.
	Interface pulumi.StringOutput `pulumi:"interface"`
	// SSL-VPN tunnel name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Authenticate peer's certificate with the peer/peergrp.
	Peer pulumi.StringOutput `pulumi:"peer"`
	// SSL-VPN server port.
	Port pulumi.IntOutput `pulumi:"port"`
	// Priority for routes added by SSL-VPN (0 - 4294967295).
	Priority pulumi.IntOutput `pulumi:"priority"`
	// Pre-shared secret to authenticate with the server (ASCII string or hexadecimal encoded with a leading 0x).
	Psk pulumi.StringPtrOutput `pulumi:"psk"`
	// Realm name configured on SSL-VPN server.
	Realm pulumi.StringOutput `pulumi:"realm"`
	// IPv4, IPv6 or DNS address of the SSL-VPN server.
	Server pulumi.StringOutput `pulumi:"server"`
	// IPv4 or IPv6 address to use as a source for the SSL-VPN connection to the server.
	SourceIp pulumi.StringOutput `pulumi:"sourceIp"`
	// Enable/disable this SSL-VPN client configuration. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Username to offer to the peer to authenticate the client.
	User pulumi.StringOutput `pulumi:"user"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewClient registers a new resource with the given unique name, arguments, and options.
func NewClient(ctx *pulumi.Context,
	name string, args *ClientArgs, opts ...pulumi.ResourceOption) (*Client, error) {
	if args == nil {
		args = &ClientArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource Client
	err := ctx.RegisterResource("fortios:vpnssl/client:Client", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClient gets an existing Client resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClient(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClientState, opts ...pulumi.ResourceOption) (*Client, error) {
	var resource Client
	err := ctx.ReadResource("fortios:vpnssl/client:Client", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Client resources.
type clientState struct {
	// Certificate to offer to SSL-VPN server if it requests one.
	Certificate *string `pulumi:"certificate"`
	// Traffic class ID.
	ClassId *int `pulumi:"classId"`
	// Comment.
	Comment *string `pulumi:"comment"`
	// Distance for routes added by SSL-VPN (1 - 255).
	Distance *int `pulumi:"distance"`
	// SSL interface to send/receive traffic over.
	Interface *string `pulumi:"interface"`
	// SSL-VPN tunnel name.
	Name *string `pulumi:"name"`
	// Authenticate peer's certificate with the peer/peergrp.
	Peer *string `pulumi:"peer"`
	// SSL-VPN server port.
	Port *int `pulumi:"port"`
	// Priority for routes added by SSL-VPN (0 - 4294967295).
	Priority *int `pulumi:"priority"`
	// Pre-shared secret to authenticate with the server (ASCII string or hexadecimal encoded with a leading 0x).
	Psk *string `pulumi:"psk"`
	// Realm name configured on SSL-VPN server.
	Realm *string `pulumi:"realm"`
	// IPv4, IPv6 or DNS address of the SSL-VPN server.
	Server *string `pulumi:"server"`
	// IPv4 or IPv6 address to use as a source for the SSL-VPN connection to the server.
	SourceIp *string `pulumi:"sourceIp"`
	// Enable/disable this SSL-VPN client configuration. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Username to offer to the peer to authenticate the client.
	User *string `pulumi:"user"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type ClientState struct {
	// Certificate to offer to SSL-VPN server if it requests one.
	Certificate pulumi.StringPtrInput
	// Traffic class ID.
	ClassId pulumi.IntPtrInput
	// Comment.
	Comment pulumi.StringPtrInput
	// Distance for routes added by SSL-VPN (1 - 255).
	Distance pulumi.IntPtrInput
	// SSL interface to send/receive traffic over.
	Interface pulumi.StringPtrInput
	// SSL-VPN tunnel name.
	Name pulumi.StringPtrInput
	// Authenticate peer's certificate with the peer/peergrp.
	Peer pulumi.StringPtrInput
	// SSL-VPN server port.
	Port pulumi.IntPtrInput
	// Priority for routes added by SSL-VPN (0 - 4294967295).
	Priority pulumi.IntPtrInput
	// Pre-shared secret to authenticate with the server (ASCII string or hexadecimal encoded with a leading 0x).
	Psk pulumi.StringPtrInput
	// Realm name configured on SSL-VPN server.
	Realm pulumi.StringPtrInput
	// IPv4, IPv6 or DNS address of the SSL-VPN server.
	Server pulumi.StringPtrInput
	// IPv4 or IPv6 address to use as a source for the SSL-VPN connection to the server.
	SourceIp pulumi.StringPtrInput
	// Enable/disable this SSL-VPN client configuration. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Username to offer to the peer to authenticate the client.
	User pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (ClientState) ElementType() reflect.Type {
	return reflect.TypeOf((*clientState)(nil)).Elem()
}

type clientArgs struct {
	// Certificate to offer to SSL-VPN server if it requests one.
	Certificate *string `pulumi:"certificate"`
	// Traffic class ID.
	ClassId *int `pulumi:"classId"`
	// Comment.
	Comment *string `pulumi:"comment"`
	// Distance for routes added by SSL-VPN (1 - 255).
	Distance *int `pulumi:"distance"`
	// SSL interface to send/receive traffic over.
	Interface *string `pulumi:"interface"`
	// SSL-VPN tunnel name.
	Name *string `pulumi:"name"`
	// Authenticate peer's certificate with the peer/peergrp.
	Peer *string `pulumi:"peer"`
	// SSL-VPN server port.
	Port *int `pulumi:"port"`
	// Priority for routes added by SSL-VPN (0 - 4294967295).
	Priority *int `pulumi:"priority"`
	// Pre-shared secret to authenticate with the server (ASCII string or hexadecimal encoded with a leading 0x).
	Psk *string `pulumi:"psk"`
	// Realm name configured on SSL-VPN server.
	Realm *string `pulumi:"realm"`
	// IPv4, IPv6 or DNS address of the SSL-VPN server.
	Server *string `pulumi:"server"`
	// IPv4 or IPv6 address to use as a source for the SSL-VPN connection to the server.
	SourceIp *string `pulumi:"sourceIp"`
	// Enable/disable this SSL-VPN client configuration. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Username to offer to the peer to authenticate the client.
	User *string `pulumi:"user"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Client resource.
type ClientArgs struct {
	// Certificate to offer to SSL-VPN server if it requests one.
	Certificate pulumi.StringPtrInput
	// Traffic class ID.
	ClassId pulumi.IntPtrInput
	// Comment.
	Comment pulumi.StringPtrInput
	// Distance for routes added by SSL-VPN (1 - 255).
	Distance pulumi.IntPtrInput
	// SSL interface to send/receive traffic over.
	Interface pulumi.StringPtrInput
	// SSL-VPN tunnel name.
	Name pulumi.StringPtrInput
	// Authenticate peer's certificate with the peer/peergrp.
	Peer pulumi.StringPtrInput
	// SSL-VPN server port.
	Port pulumi.IntPtrInput
	// Priority for routes added by SSL-VPN (0 - 4294967295).
	Priority pulumi.IntPtrInput
	// Pre-shared secret to authenticate with the server (ASCII string or hexadecimal encoded with a leading 0x).
	Psk pulumi.StringPtrInput
	// Realm name configured on SSL-VPN server.
	Realm pulumi.StringPtrInput
	// IPv4, IPv6 or DNS address of the SSL-VPN server.
	Server pulumi.StringPtrInput
	// IPv4 or IPv6 address to use as a source for the SSL-VPN connection to the server.
	SourceIp pulumi.StringPtrInput
	// Enable/disable this SSL-VPN client configuration. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Username to offer to the peer to authenticate the client.
	User pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (ClientArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clientArgs)(nil)).Elem()
}

type ClientInput interface {
	pulumi.Input

	ToClientOutput() ClientOutput
	ToClientOutputWithContext(ctx context.Context) ClientOutput
}

func (*Client) ElementType() reflect.Type {
	return reflect.TypeOf((**Client)(nil)).Elem()
}

func (i *Client) ToClientOutput() ClientOutput {
	return i.ToClientOutputWithContext(context.Background())
}

func (i *Client) ToClientOutputWithContext(ctx context.Context) ClientOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientOutput)
}

// ClientArrayInput is an input type that accepts ClientArray and ClientArrayOutput values.
// You can construct a concrete instance of `ClientArrayInput` via:
//
//	ClientArray{ ClientArgs{...} }
type ClientArrayInput interface {
	pulumi.Input

	ToClientArrayOutput() ClientArrayOutput
	ToClientArrayOutputWithContext(context.Context) ClientArrayOutput
}

type ClientArray []ClientInput

func (ClientArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Client)(nil)).Elem()
}

func (i ClientArray) ToClientArrayOutput() ClientArrayOutput {
	return i.ToClientArrayOutputWithContext(context.Background())
}

func (i ClientArray) ToClientArrayOutputWithContext(ctx context.Context) ClientArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientArrayOutput)
}

// ClientMapInput is an input type that accepts ClientMap and ClientMapOutput values.
// You can construct a concrete instance of `ClientMapInput` via:
//
//	ClientMap{ "key": ClientArgs{...} }
type ClientMapInput interface {
	pulumi.Input

	ToClientMapOutput() ClientMapOutput
	ToClientMapOutputWithContext(context.Context) ClientMapOutput
}

type ClientMap map[string]ClientInput

func (ClientMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Client)(nil)).Elem()
}

func (i ClientMap) ToClientMapOutput() ClientMapOutput {
	return i.ToClientMapOutputWithContext(context.Background())
}

func (i ClientMap) ToClientMapOutputWithContext(ctx context.Context) ClientMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientMapOutput)
}

type ClientOutput struct{ *pulumi.OutputState }

func (ClientOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Client)(nil)).Elem()
}

func (o ClientOutput) ToClientOutput() ClientOutput {
	return o
}

func (o ClientOutput) ToClientOutputWithContext(ctx context.Context) ClientOutput {
	return o
}

// Certificate to offer to SSL-VPN server if it requests one.
func (o ClientOutput) Certificate() pulumi.StringOutput {
	return o.ApplyT(func(v *Client) pulumi.StringOutput { return v.Certificate }).(pulumi.StringOutput)
}

// Traffic class ID.
func (o ClientOutput) ClassId() pulumi.IntOutput {
	return o.ApplyT(func(v *Client) pulumi.IntOutput { return v.ClassId }).(pulumi.IntOutput)
}

// Comment.
func (o ClientOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Client) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// Distance for routes added by SSL-VPN (1 - 255).
func (o ClientOutput) Distance() pulumi.IntOutput {
	return o.ApplyT(func(v *Client) pulumi.IntOutput { return v.Distance }).(pulumi.IntOutput)
}

// SSL interface to send/receive traffic over.
func (o ClientOutput) Interface() pulumi.StringOutput {
	return o.ApplyT(func(v *Client) pulumi.StringOutput { return v.Interface }).(pulumi.StringOutput)
}

// SSL-VPN tunnel name.
func (o ClientOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Client) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Authenticate peer's certificate with the peer/peergrp.
func (o ClientOutput) Peer() pulumi.StringOutput {
	return o.ApplyT(func(v *Client) pulumi.StringOutput { return v.Peer }).(pulumi.StringOutput)
}

// SSL-VPN server port.
func (o ClientOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v *Client) pulumi.IntOutput { return v.Port }).(pulumi.IntOutput)
}

// Priority for routes added by SSL-VPN (0 - 4294967295).
func (o ClientOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v *Client) pulumi.IntOutput { return v.Priority }).(pulumi.IntOutput)
}

// Pre-shared secret to authenticate with the server (ASCII string or hexadecimal encoded with a leading 0x).
func (o ClientOutput) Psk() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Client) pulumi.StringPtrOutput { return v.Psk }).(pulumi.StringPtrOutput)
}

// Realm name configured on SSL-VPN server.
func (o ClientOutput) Realm() pulumi.StringOutput {
	return o.ApplyT(func(v *Client) pulumi.StringOutput { return v.Realm }).(pulumi.StringOutput)
}

// IPv4, IPv6 or DNS address of the SSL-VPN server.
func (o ClientOutput) Server() pulumi.StringOutput {
	return o.ApplyT(func(v *Client) pulumi.StringOutput { return v.Server }).(pulumi.StringOutput)
}

// IPv4 or IPv6 address to use as a source for the SSL-VPN connection to the server.
func (o ClientOutput) SourceIp() pulumi.StringOutput {
	return o.ApplyT(func(v *Client) pulumi.StringOutput { return v.SourceIp }).(pulumi.StringOutput)
}

// Enable/disable this SSL-VPN client configuration. Valid values: `enable`, `disable`.
func (o ClientOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Client) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Username to offer to the peer to authenticate the client.
func (o ClientOutput) User() pulumi.StringOutput {
	return o.ApplyT(func(v *Client) pulumi.StringOutput { return v.User }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o ClientOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Client) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type ClientArrayOutput struct{ *pulumi.OutputState }

func (ClientArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Client)(nil)).Elem()
}

func (o ClientArrayOutput) ToClientArrayOutput() ClientArrayOutput {
	return o
}

func (o ClientArrayOutput) ToClientArrayOutputWithContext(ctx context.Context) ClientArrayOutput {
	return o
}

func (o ClientArrayOutput) Index(i pulumi.IntInput) ClientOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Client {
		return vs[0].([]*Client)[vs[1].(int)]
	}).(ClientOutput)
}

type ClientMapOutput struct{ *pulumi.OutputState }

func (ClientMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Client)(nil)).Elem()
}

func (o ClientMapOutput) ToClientMapOutput() ClientMapOutput {
	return o
}

func (o ClientMapOutput) ToClientMapOutputWithContext(ctx context.Context) ClientMapOutput {
	return o
}

func (o ClientMapOutput) MapIndex(k pulumi.StringInput) ClientOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Client {
		return vs[0].(map[string]*Client)[vs[1].(string)]
	}).(ClientOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClientInput)(nil)).Elem(), &Client{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClientArrayInput)(nil)).Elem(), ClientArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClientMapInput)(nil)).Elem(), ClientMap{})
	pulumi.RegisterOutputType(ClientOutput{})
	pulumi.RegisterOutputType(ClientArrayOutput{})
	pulumi.RegisterOutputType(ClientMapOutput{})
}
