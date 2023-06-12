// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sys

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on fortios system proberesponse
func LookupProberesponse(ctx *pulumi.Context, args *LookupProberesponseArgs, opts ...pulumi.InvokeOption) (*LookupProberesponseResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupProberesponseResult
	err := ctx.Invoke("fortios:sys/getProberesponse:getProberesponse", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProberesponse.
type LookupProberesponseArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getProberesponse.
type LookupProberesponseResult struct {
	// Value to respond to the monitoring server.
	HttpProbeValue string `pulumi:"httpProbeValue"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// SLA response mode.
	Mode string `pulumi:"mode"`
	// Twamp respondor password in authentication mode
	Password string `pulumi:"password"`
	// Port number to response.
	Port int `pulumi:"port"`
	// Twamp respondor security mode.
	SecurityMode string `pulumi:"securityMode"`
	// An inactivity timer for a twamp test session.
	Timeout int `pulumi:"timeout"`
	// Mode for TWAMP packet TTL modification.
	TtlMode   string  `pulumi:"ttlMode"`
	Vdomparam *string `pulumi:"vdomparam"`
}

func LookupProberesponseOutput(ctx *pulumi.Context, args LookupProberesponseOutputArgs, opts ...pulumi.InvokeOption) LookupProberesponseResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupProberesponseResult, error) {
			args := v.(LookupProberesponseArgs)
			r, err := LookupProberesponse(ctx, &args, opts...)
			var s LookupProberesponseResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupProberesponseResultOutput)
}

// A collection of arguments for invoking getProberesponse.
type LookupProberesponseOutputArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupProberesponseOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProberesponseArgs)(nil)).Elem()
}

// A collection of values returned by getProberesponse.
type LookupProberesponseResultOutput struct{ *pulumi.OutputState }

func (LookupProberesponseResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProberesponseResult)(nil)).Elem()
}

func (o LookupProberesponseResultOutput) ToLookupProberesponseResultOutput() LookupProberesponseResultOutput {
	return o
}

func (o LookupProberesponseResultOutput) ToLookupProberesponseResultOutputWithContext(ctx context.Context) LookupProberesponseResultOutput {
	return o
}

// Value to respond to the monitoring server.
func (o LookupProberesponseResultOutput) HttpProbeValue() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProberesponseResult) string { return v.HttpProbeValue }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupProberesponseResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProberesponseResult) string { return v.Id }).(pulumi.StringOutput)
}

// SLA response mode.
func (o LookupProberesponseResultOutput) Mode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProberesponseResult) string { return v.Mode }).(pulumi.StringOutput)
}

// Twamp respondor password in authentication mode
func (o LookupProberesponseResultOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProberesponseResult) string { return v.Password }).(pulumi.StringOutput)
}

// Port number to response.
func (o LookupProberesponseResultOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v LookupProberesponseResult) int { return v.Port }).(pulumi.IntOutput)
}

// Twamp respondor security mode.
func (o LookupProberesponseResultOutput) SecurityMode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProberesponseResult) string { return v.SecurityMode }).(pulumi.StringOutput)
}

// An inactivity timer for a twamp test session.
func (o LookupProberesponseResultOutput) Timeout() pulumi.IntOutput {
	return o.ApplyT(func(v LookupProberesponseResult) int { return v.Timeout }).(pulumi.IntOutput)
}

// Mode for TWAMP packet TTL modification.
func (o LookupProberesponseResultOutput) TtlMode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProberesponseResult) string { return v.TtlMode }).(pulumi.StringOutput)
}

func (o LookupProberesponseResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProberesponseResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProberesponseResultOutput{})
}
