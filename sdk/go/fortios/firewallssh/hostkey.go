// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firewallssh

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// SSH proxy host public keys.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/firewallssh"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := firewallssh.NewHostkey(ctx, "trname", &firewallssh.HostkeyArgs{
//				Hostname: pulumi.String("testmachine"),
//				Ip:       pulumi.String("1.1.1.1"),
//				Port:     pulumi.Int(22),
//				Status:   pulumi.String("trusted"),
//				Type:     pulumi.String("RSA"),
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
// # FirewallSsh HostKey can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:firewallssh/hostkey:Hostkey labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:firewallssh/hostkey:Hostkey labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type Hostkey struct {
	pulumi.CustomResourceState

	// Hostname of the SSH server.
	Hostname pulumi.StringOutput `pulumi:"hostname"`
	// IP address of the SSH server.
	Ip pulumi.StringOutput `pulumi:"ip"`
	// SSH public key name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Set the nid of the ECDSA key. Valid values: `256`, `384`, `521`.
	Nid pulumi.StringOutput `pulumi:"nid"`
	// Port of the SSH server.
	Port pulumi.IntOutput `pulumi:"port"`
	// SSH public key.
	PublicKey pulumi.StringPtrOutput `pulumi:"publicKey"`
	// Set the trust status of the public key. Valid values: `trusted`, `revoked`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Set the type of the public key. Valid values: `RSA`, `DSA`, `ECDSA`, `ED25519`, `RSA-CA`, `DSA-CA`, `ECDSA-CA`, `ED25519-CA`.
	Type pulumi.StringOutput `pulumi:"type"`
	// Usage for this public key. Valid values: `transparent-proxy`, `access-proxy`.
	Usage pulumi.StringOutput `pulumi:"usage"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewHostkey registers a new resource with the given unique name, arguments, and options.
func NewHostkey(ctx *pulumi.Context,
	name string, args *HostkeyArgs, opts ...pulumi.ResourceOption) (*Hostkey, error) {
	if args == nil {
		args = &HostkeyArgs{}
	}

	if args.PublicKey != nil {
		args.PublicKey = pulumi.ToSecret(args.PublicKey).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"publicKey",
	})
	opts = append(opts, secrets)
	opts = pkgResourceDefaultOpts(opts)
	var resource Hostkey
	err := ctx.RegisterResource("fortios:firewallssh/hostkey:Hostkey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHostkey gets an existing Hostkey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHostkey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HostkeyState, opts ...pulumi.ResourceOption) (*Hostkey, error) {
	var resource Hostkey
	err := ctx.ReadResource("fortios:firewallssh/hostkey:Hostkey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Hostkey resources.
type hostkeyState struct {
	// Hostname of the SSH server.
	Hostname *string `pulumi:"hostname"`
	// IP address of the SSH server.
	Ip *string `pulumi:"ip"`
	// SSH public key name.
	Name *string `pulumi:"name"`
	// Set the nid of the ECDSA key. Valid values: `256`, `384`, `521`.
	Nid *string `pulumi:"nid"`
	// Port of the SSH server.
	Port *int `pulumi:"port"`
	// SSH public key.
	PublicKey *string `pulumi:"publicKey"`
	// Set the trust status of the public key. Valid values: `trusted`, `revoked`.
	Status *string `pulumi:"status"`
	// Set the type of the public key. Valid values: `RSA`, `DSA`, `ECDSA`, `ED25519`, `RSA-CA`, `DSA-CA`, `ECDSA-CA`, `ED25519-CA`.
	Type *string `pulumi:"type"`
	// Usage for this public key. Valid values: `transparent-proxy`, `access-proxy`.
	Usage *string `pulumi:"usage"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type HostkeyState struct {
	// Hostname of the SSH server.
	Hostname pulumi.StringPtrInput
	// IP address of the SSH server.
	Ip pulumi.StringPtrInput
	// SSH public key name.
	Name pulumi.StringPtrInput
	// Set the nid of the ECDSA key. Valid values: `256`, `384`, `521`.
	Nid pulumi.StringPtrInput
	// Port of the SSH server.
	Port pulumi.IntPtrInput
	// SSH public key.
	PublicKey pulumi.StringPtrInput
	// Set the trust status of the public key. Valid values: `trusted`, `revoked`.
	Status pulumi.StringPtrInput
	// Set the type of the public key. Valid values: `RSA`, `DSA`, `ECDSA`, `ED25519`, `RSA-CA`, `DSA-CA`, `ECDSA-CA`, `ED25519-CA`.
	Type pulumi.StringPtrInput
	// Usage for this public key. Valid values: `transparent-proxy`, `access-proxy`.
	Usage pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (HostkeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*hostkeyState)(nil)).Elem()
}

type hostkeyArgs struct {
	// Hostname of the SSH server.
	Hostname *string `pulumi:"hostname"`
	// IP address of the SSH server.
	Ip *string `pulumi:"ip"`
	// SSH public key name.
	Name *string `pulumi:"name"`
	// Set the nid of the ECDSA key. Valid values: `256`, `384`, `521`.
	Nid *string `pulumi:"nid"`
	// Port of the SSH server.
	Port *int `pulumi:"port"`
	// SSH public key.
	PublicKey *string `pulumi:"publicKey"`
	// Set the trust status of the public key. Valid values: `trusted`, `revoked`.
	Status *string `pulumi:"status"`
	// Set the type of the public key. Valid values: `RSA`, `DSA`, `ECDSA`, `ED25519`, `RSA-CA`, `DSA-CA`, `ECDSA-CA`, `ED25519-CA`.
	Type *string `pulumi:"type"`
	// Usage for this public key. Valid values: `transparent-proxy`, `access-proxy`.
	Usage *string `pulumi:"usage"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Hostkey resource.
type HostkeyArgs struct {
	// Hostname of the SSH server.
	Hostname pulumi.StringPtrInput
	// IP address of the SSH server.
	Ip pulumi.StringPtrInput
	// SSH public key name.
	Name pulumi.StringPtrInput
	// Set the nid of the ECDSA key. Valid values: `256`, `384`, `521`.
	Nid pulumi.StringPtrInput
	// Port of the SSH server.
	Port pulumi.IntPtrInput
	// SSH public key.
	PublicKey pulumi.StringPtrInput
	// Set the trust status of the public key. Valid values: `trusted`, `revoked`.
	Status pulumi.StringPtrInput
	// Set the type of the public key. Valid values: `RSA`, `DSA`, `ECDSA`, `ED25519`, `RSA-CA`, `DSA-CA`, `ECDSA-CA`, `ED25519-CA`.
	Type pulumi.StringPtrInput
	// Usage for this public key. Valid values: `transparent-proxy`, `access-proxy`.
	Usage pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (HostkeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hostkeyArgs)(nil)).Elem()
}

type HostkeyInput interface {
	pulumi.Input

	ToHostkeyOutput() HostkeyOutput
	ToHostkeyOutputWithContext(ctx context.Context) HostkeyOutput
}

func (*Hostkey) ElementType() reflect.Type {
	return reflect.TypeOf((**Hostkey)(nil)).Elem()
}

func (i *Hostkey) ToHostkeyOutput() HostkeyOutput {
	return i.ToHostkeyOutputWithContext(context.Background())
}

func (i *Hostkey) ToHostkeyOutputWithContext(ctx context.Context) HostkeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostkeyOutput)
}

// HostkeyArrayInput is an input type that accepts HostkeyArray and HostkeyArrayOutput values.
// You can construct a concrete instance of `HostkeyArrayInput` via:
//
//	HostkeyArray{ HostkeyArgs{...} }
type HostkeyArrayInput interface {
	pulumi.Input

	ToHostkeyArrayOutput() HostkeyArrayOutput
	ToHostkeyArrayOutputWithContext(context.Context) HostkeyArrayOutput
}

type HostkeyArray []HostkeyInput

func (HostkeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Hostkey)(nil)).Elem()
}

func (i HostkeyArray) ToHostkeyArrayOutput() HostkeyArrayOutput {
	return i.ToHostkeyArrayOutputWithContext(context.Background())
}

func (i HostkeyArray) ToHostkeyArrayOutputWithContext(ctx context.Context) HostkeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostkeyArrayOutput)
}

// HostkeyMapInput is an input type that accepts HostkeyMap and HostkeyMapOutput values.
// You can construct a concrete instance of `HostkeyMapInput` via:
//
//	HostkeyMap{ "key": HostkeyArgs{...} }
type HostkeyMapInput interface {
	pulumi.Input

	ToHostkeyMapOutput() HostkeyMapOutput
	ToHostkeyMapOutputWithContext(context.Context) HostkeyMapOutput
}

type HostkeyMap map[string]HostkeyInput

func (HostkeyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Hostkey)(nil)).Elem()
}

func (i HostkeyMap) ToHostkeyMapOutput() HostkeyMapOutput {
	return i.ToHostkeyMapOutputWithContext(context.Background())
}

func (i HostkeyMap) ToHostkeyMapOutputWithContext(ctx context.Context) HostkeyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostkeyMapOutput)
}

type HostkeyOutput struct{ *pulumi.OutputState }

func (HostkeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Hostkey)(nil)).Elem()
}

func (o HostkeyOutput) ToHostkeyOutput() HostkeyOutput {
	return o
}

func (o HostkeyOutput) ToHostkeyOutputWithContext(ctx context.Context) HostkeyOutput {
	return o
}

// Hostname of the SSH server.
func (o HostkeyOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v *Hostkey) pulumi.StringOutput { return v.Hostname }).(pulumi.StringOutput)
}

// IP address of the SSH server.
func (o HostkeyOutput) Ip() pulumi.StringOutput {
	return o.ApplyT(func(v *Hostkey) pulumi.StringOutput { return v.Ip }).(pulumi.StringOutput)
}

// SSH public key name.
func (o HostkeyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Hostkey) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Set the nid of the ECDSA key. Valid values: `256`, `384`, `521`.
func (o HostkeyOutput) Nid() pulumi.StringOutput {
	return o.ApplyT(func(v *Hostkey) pulumi.StringOutput { return v.Nid }).(pulumi.StringOutput)
}

// Port of the SSH server.
func (o HostkeyOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v *Hostkey) pulumi.IntOutput { return v.Port }).(pulumi.IntOutput)
}

// SSH public key.
func (o HostkeyOutput) PublicKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Hostkey) pulumi.StringPtrOutput { return v.PublicKey }).(pulumi.StringPtrOutput)
}

// Set the trust status of the public key. Valid values: `trusted`, `revoked`.
func (o HostkeyOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Hostkey) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Set the type of the public key. Valid values: `RSA`, `DSA`, `ECDSA`, `ED25519`, `RSA-CA`, `DSA-CA`, `ECDSA-CA`, `ED25519-CA`.
func (o HostkeyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Hostkey) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Usage for this public key. Valid values: `transparent-proxy`, `access-proxy`.
func (o HostkeyOutput) Usage() pulumi.StringOutput {
	return o.ApplyT(func(v *Hostkey) pulumi.StringOutput { return v.Usage }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o HostkeyOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Hostkey) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type HostkeyArrayOutput struct{ *pulumi.OutputState }

func (HostkeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Hostkey)(nil)).Elem()
}

func (o HostkeyArrayOutput) ToHostkeyArrayOutput() HostkeyArrayOutput {
	return o
}

func (o HostkeyArrayOutput) ToHostkeyArrayOutputWithContext(ctx context.Context) HostkeyArrayOutput {
	return o
}

func (o HostkeyArrayOutput) Index(i pulumi.IntInput) HostkeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Hostkey {
		return vs[0].([]*Hostkey)[vs[1].(int)]
	}).(HostkeyOutput)
}

type HostkeyMapOutput struct{ *pulumi.OutputState }

func (HostkeyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Hostkey)(nil)).Elem()
}

func (o HostkeyMapOutput) ToHostkeyMapOutput() HostkeyMapOutput {
	return o
}

func (o HostkeyMapOutput) ToHostkeyMapOutputWithContext(ctx context.Context) HostkeyMapOutput {
	return o
}

func (o HostkeyMapOutput) MapIndex(k pulumi.StringInput) HostkeyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Hostkey {
		return vs[0].(map[string]*Hostkey)[vs[1].(string)]
	}).(HostkeyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HostkeyInput)(nil)).Elem(), &Hostkey{})
	pulumi.RegisterInputType(reflect.TypeOf((*HostkeyArrayInput)(nil)).Elem(), HostkeyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*HostkeyMapInput)(nil)).Elem(), HostkeyMap{})
	pulumi.RegisterOutputType(HostkeyOutput{})
	pulumi.RegisterOutputType(HostkeyArrayOutput{})
	pulumi.RegisterOutputType(HostkeyMapOutput{})
}