// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sys

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on fortios system emailserver
func LookupEmailserver(ctx *pulumi.Context, args *LookupEmailserverArgs, opts ...pulumi.InvokeOption) (*LookupEmailserverResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupEmailserverResult
	err := ctx.Invoke("fortios:sys/getEmailserver:getEmailserver", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getEmailserver.
type LookupEmailserverArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getEmailserver.
type LookupEmailserverResult struct {
	// Enable/disable authentication.
	Authenticate string `pulumi:"authenticate"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Specify outgoing interface to reach server.
	Interface string `pulumi:"interface"`
	// Specify how to select outgoing interface to reach server.
	InterfaceSelectMethod string `pulumi:"interfaceSelectMethod"`
	// SMTP server user password for authentication.
	Password string `pulumi:"password"`
	// SMTP server port.
	Port int `pulumi:"port"`
	// Reply-To email address.
	ReplyTo string `pulumi:"replyTo"`
	// Connection security used by the email server.
	Security string `pulumi:"security"`
	// SMTP server IP address or hostname.
	Server string `pulumi:"server"`
	// SMTP server IPv4 source IP.
	SourceIp string `pulumi:"sourceIp"`
	// SMTP server IPv6 source IP.
	SourceIp6 string `pulumi:"sourceIp6"`
	// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
	SslMinProtoVersion string `pulumi:"sslMinProtoVersion"`
	// Use FortiGuard Message service or custom email server.
	Type string `pulumi:"type"`
	// SMTP server user name for authentication.
	Username string `pulumi:"username"`
	// Enable/disable validation of server certificate.
	ValidateServer string  `pulumi:"validateServer"`
	Vdomparam      *string `pulumi:"vdomparam"`
}

func LookupEmailserverOutput(ctx *pulumi.Context, args LookupEmailserverOutputArgs, opts ...pulumi.InvokeOption) LookupEmailserverResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEmailserverResult, error) {
			args := v.(LookupEmailserverArgs)
			r, err := LookupEmailserver(ctx, &args, opts...)
			var s LookupEmailserverResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEmailserverResultOutput)
}

// A collection of arguments for invoking getEmailserver.
type LookupEmailserverOutputArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupEmailserverOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEmailserverArgs)(nil)).Elem()
}

// A collection of values returned by getEmailserver.
type LookupEmailserverResultOutput struct{ *pulumi.OutputState }

func (LookupEmailserverResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEmailserverResult)(nil)).Elem()
}

func (o LookupEmailserverResultOutput) ToLookupEmailserverResultOutput() LookupEmailserverResultOutput {
	return o
}

func (o LookupEmailserverResultOutput) ToLookupEmailserverResultOutputWithContext(ctx context.Context) LookupEmailserverResultOutput {
	return o
}

// Enable/disable authentication.
func (o LookupEmailserverResultOutput) Authenticate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEmailserverResult) string { return v.Authenticate }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupEmailserverResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEmailserverResult) string { return v.Id }).(pulumi.StringOutput)
}

// Specify outgoing interface to reach server.
func (o LookupEmailserverResultOutput) Interface() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEmailserverResult) string { return v.Interface }).(pulumi.StringOutput)
}

// Specify how to select outgoing interface to reach server.
func (o LookupEmailserverResultOutput) InterfaceSelectMethod() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEmailserverResult) string { return v.InterfaceSelectMethod }).(pulumi.StringOutput)
}

// SMTP server user password for authentication.
func (o LookupEmailserverResultOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEmailserverResult) string { return v.Password }).(pulumi.StringOutput)
}

// SMTP server port.
func (o LookupEmailserverResultOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v LookupEmailserverResult) int { return v.Port }).(pulumi.IntOutput)
}

// Reply-To email address.
func (o LookupEmailserverResultOutput) ReplyTo() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEmailserverResult) string { return v.ReplyTo }).(pulumi.StringOutput)
}

// Connection security used by the email server.
func (o LookupEmailserverResultOutput) Security() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEmailserverResult) string { return v.Security }).(pulumi.StringOutput)
}

// SMTP server IP address or hostname.
func (o LookupEmailserverResultOutput) Server() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEmailserverResult) string { return v.Server }).(pulumi.StringOutput)
}

// SMTP server IPv4 source IP.
func (o LookupEmailserverResultOutput) SourceIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEmailserverResult) string { return v.SourceIp }).(pulumi.StringOutput)
}

// SMTP server IPv6 source IP.
func (o LookupEmailserverResultOutput) SourceIp6() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEmailserverResult) string { return v.SourceIp6 }).(pulumi.StringOutput)
}

// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
func (o LookupEmailserverResultOutput) SslMinProtoVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEmailserverResult) string { return v.SslMinProtoVersion }).(pulumi.StringOutput)
}

// Use FortiGuard Message service or custom email server.
func (o LookupEmailserverResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEmailserverResult) string { return v.Type }).(pulumi.StringOutput)
}

// SMTP server user name for authentication.
func (o LookupEmailserverResultOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEmailserverResult) string { return v.Username }).(pulumi.StringOutput)
}

// Enable/disable validation of server certificate.
func (o LookupEmailserverResultOutput) ValidateServer() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEmailserverResult) string { return v.ValidateServer }).(pulumi.StringOutput)
}

func (o LookupEmailserverResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEmailserverResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEmailserverResultOutput{})
}
