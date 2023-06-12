// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The provider type for the fortios package. By default, resources use package-wide configuration
// settings, however an explicit `Provider` instance may be created and passed during resource
// construction to achieve fine-grained programmatic control over provider settings. See the
// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
type Provider struct {
	pulumi.ProviderResourceState

	// CA Bundle file content
	Cabundlecontent pulumi.StringPtrOutput `pulumi:"cabundlecontent"`
	// CA Bundle file
	Cabundlefile pulumi.StringPtrOutput `pulumi:"cabundlefile"`
	// CA certtificate(Optional)
	Cacert pulumi.StringPtrOutput `pulumi:"cacert"`
	// User certificate
	Clientcert pulumi.StringPtrOutput `pulumi:"clientcert"`
	// User private key
	Clientkey pulumi.StringPtrOutput `pulumi:"clientkey"`
	// CA Bundle file
	FmgCabundlefile pulumi.StringPtrOutput `pulumi:"fmgCabundlefile"`
	// Hostname/IP address of the FortiManager to connect to
	FmgHostname pulumi.StringPtrOutput `pulumi:"fmgHostname"`
	FmgPasswd   pulumi.StringPtrOutput `pulumi:"fmgPasswd"`
	FmgUsername pulumi.StringPtrOutput `pulumi:"fmgUsername"`
	// The hostname/IP address of the FortiOS to be connected
	Hostname pulumi.StringPtrOutput `pulumi:"hostname"`
	// HTTP proxy address
	HttpProxy pulumi.StringPtrOutput `pulumi:"httpProxy"`
	// Enable/disable peer authentication, can be 'enable' or 'disable'
	Peerauth pulumi.StringPtrOutput `pulumi:"peerauth"`
	Token    pulumi.StringPtrOutput `pulumi:"token"`
	Vdom     pulumi.StringPtrOutput `pulumi:"vdom"`
}

// NewProvider registers a new resource with the given unique name, arguments, and options.
func NewProvider(ctx *pulumi.Context,
	name string, args *ProviderArgs, opts ...pulumi.ResourceOption) (*Provider, error) {
	if args == nil {
		args = &ProviderArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource Provider
	err := ctx.RegisterResource("pulumi:providers:fortios", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type providerArgs struct {
	// CA Bundle file content
	Cabundlecontent *string `pulumi:"cabundlecontent"`
	// CA Bundle file
	Cabundlefile *string `pulumi:"cabundlefile"`
	// CA certtificate(Optional)
	Cacert *string `pulumi:"cacert"`
	// User certificate
	Clientcert *string `pulumi:"clientcert"`
	// User private key
	Clientkey *string `pulumi:"clientkey"`
	// CA Bundle file
	FmgCabundlefile *string `pulumi:"fmgCabundlefile"`
	// Hostname/IP address of the FortiManager to connect to
	FmgHostname *string `pulumi:"fmgHostname"`
	FmgInsecure *bool   `pulumi:"fmgInsecure"`
	FmgPasswd   *string `pulumi:"fmgPasswd"`
	FmgUsername *string `pulumi:"fmgUsername"`
	// The hostname/IP address of the FortiOS to be connected
	Hostname *string `pulumi:"hostname"`
	// HTTP proxy address
	HttpProxy *string `pulumi:"httpProxy"`
	Insecure  *bool   `pulumi:"insecure"`
	// Enable/disable peer authentication, can be 'enable' or 'disable'
	Peerauth *string `pulumi:"peerauth"`
	Token    *string `pulumi:"token"`
	Vdom     *string `pulumi:"vdom"`
}

// The set of arguments for constructing a Provider resource.
type ProviderArgs struct {
	// CA Bundle file content
	Cabundlecontent pulumi.StringPtrInput
	// CA Bundle file
	Cabundlefile pulumi.StringPtrInput
	// CA certtificate(Optional)
	Cacert pulumi.StringPtrInput
	// User certificate
	Clientcert pulumi.StringPtrInput
	// User private key
	Clientkey pulumi.StringPtrInput
	// CA Bundle file
	FmgCabundlefile pulumi.StringPtrInput
	// Hostname/IP address of the FortiManager to connect to
	FmgHostname pulumi.StringPtrInput
	FmgInsecure pulumi.BoolPtrInput
	FmgPasswd   pulumi.StringPtrInput
	FmgUsername pulumi.StringPtrInput
	// The hostname/IP address of the FortiOS to be connected
	Hostname pulumi.StringPtrInput
	// HTTP proxy address
	HttpProxy pulumi.StringPtrInput
	Insecure  pulumi.BoolPtrInput
	// Enable/disable peer authentication, can be 'enable' or 'disable'
	Peerauth pulumi.StringPtrInput
	Token    pulumi.StringPtrInput
	Vdom     pulumi.StringPtrInput
}

func (ProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerArgs)(nil)).Elem()
}

type ProviderInput interface {
	pulumi.Input

	ToProviderOutput() ProviderOutput
	ToProviderOutputWithContext(ctx context.Context) ProviderOutput
}

func (*Provider) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (i *Provider) ToProviderOutput() ProviderOutput {
	return i.ToProviderOutputWithContext(context.Background())
}

func (i *Provider) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderOutput)
}

type ProviderOutput struct{ *pulumi.OutputState }

func (ProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (o ProviderOutput) ToProviderOutput() ProviderOutput {
	return o
}

func (o ProviderOutput) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return o
}

// CA Bundle file content
func (o ProviderOutput) Cabundlecontent() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.Cabundlecontent }).(pulumi.StringPtrOutput)
}

// CA Bundle file
func (o ProviderOutput) Cabundlefile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.Cabundlefile }).(pulumi.StringPtrOutput)
}

// CA certtificate(Optional)
func (o ProviderOutput) Cacert() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.Cacert }).(pulumi.StringPtrOutput)
}

// User certificate
func (o ProviderOutput) Clientcert() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.Clientcert }).(pulumi.StringPtrOutput)
}

// User private key
func (o ProviderOutput) Clientkey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.Clientkey }).(pulumi.StringPtrOutput)
}

// CA Bundle file
func (o ProviderOutput) FmgCabundlefile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.FmgCabundlefile }).(pulumi.StringPtrOutput)
}

// Hostname/IP address of the FortiManager to connect to
func (o ProviderOutput) FmgHostname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.FmgHostname }).(pulumi.StringPtrOutput)
}

func (o ProviderOutput) FmgPasswd() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.FmgPasswd }).(pulumi.StringPtrOutput)
}

func (o ProviderOutput) FmgUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.FmgUsername }).(pulumi.StringPtrOutput)
}

// The hostname/IP address of the FortiOS to be connected
func (o ProviderOutput) Hostname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.Hostname }).(pulumi.StringPtrOutput)
}

// HTTP proxy address
func (o ProviderOutput) HttpProxy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.HttpProxy }).(pulumi.StringPtrOutput)
}

// Enable/disable peer authentication, can be 'enable' or 'disable'
func (o ProviderOutput) Peerauth() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.Peerauth }).(pulumi.StringPtrOutput)
}

func (o ProviderOutput) Token() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.Token }).(pulumi.StringPtrOutput)
}

func (o ProviderOutput) Vdom() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.Vdom }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderInput)(nil)).Elem(), &Provider{})
	pulumi.RegisterOutputType(ProviderOutput{})
}
