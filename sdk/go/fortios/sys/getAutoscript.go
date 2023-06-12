// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sys

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios system autoscript
func LookupAutoscript(ctx *pulumi.Context, args *LookupAutoscriptArgs, opts ...pulumi.InvokeOption) (*LookupAutoscriptResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupAutoscriptResult
	err := ctx.Invoke("fortios:sys/getAutoscript:getAutoscript", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAutoscript.
type LookupAutoscriptArgs struct {
	// Specify the name of the desired system autoscript.
	Name string `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getAutoscript.
type LookupAutoscriptResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Repeat interval in seconds.
	Interval int `pulumi:"interval"`
	// Auto script name.
	Name string `pulumi:"name"`
	// Number of megabytes to limit script output to (10 - 1024, default = 10).
	OutputSize int `pulumi:"outputSize"`
	// Number of times to repeat this script (0 = infinite).
	Repeat int `pulumi:"repeat"`
	// List of FortiOS CLI commands to repeat.
	Script string `pulumi:"script"`
	// Script starting mode.
	Start string `pulumi:"start"`
	// Maximum running time for this script in seconds (0 = no timeout).
	Timeout   int     `pulumi:"timeout"`
	Vdomparam *string `pulumi:"vdomparam"`
}

func LookupAutoscriptOutput(ctx *pulumi.Context, args LookupAutoscriptOutputArgs, opts ...pulumi.InvokeOption) LookupAutoscriptResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAutoscriptResult, error) {
			args := v.(LookupAutoscriptArgs)
			r, err := LookupAutoscript(ctx, &args, opts...)
			var s LookupAutoscriptResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAutoscriptResultOutput)
}

// A collection of arguments for invoking getAutoscript.
type LookupAutoscriptOutputArgs struct {
	// Specify the name of the desired system autoscript.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupAutoscriptOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAutoscriptArgs)(nil)).Elem()
}

// A collection of values returned by getAutoscript.
type LookupAutoscriptResultOutput struct{ *pulumi.OutputState }

func (LookupAutoscriptResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAutoscriptResult)(nil)).Elem()
}

func (o LookupAutoscriptResultOutput) ToLookupAutoscriptResultOutput() LookupAutoscriptResultOutput {
	return o
}

func (o LookupAutoscriptResultOutput) ToLookupAutoscriptResultOutputWithContext(ctx context.Context) LookupAutoscriptResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupAutoscriptResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAutoscriptResult) string { return v.Id }).(pulumi.StringOutput)
}

// Repeat interval in seconds.
func (o LookupAutoscriptResultOutput) Interval() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAutoscriptResult) int { return v.Interval }).(pulumi.IntOutput)
}

// Auto script name.
func (o LookupAutoscriptResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAutoscriptResult) string { return v.Name }).(pulumi.StringOutput)
}

// Number of megabytes to limit script output to (10 - 1024, default = 10).
func (o LookupAutoscriptResultOutput) OutputSize() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAutoscriptResult) int { return v.OutputSize }).(pulumi.IntOutput)
}

// Number of times to repeat this script (0 = infinite).
func (o LookupAutoscriptResultOutput) Repeat() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAutoscriptResult) int { return v.Repeat }).(pulumi.IntOutput)
}

// List of FortiOS CLI commands to repeat.
func (o LookupAutoscriptResultOutput) Script() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAutoscriptResult) string { return v.Script }).(pulumi.StringOutput)
}

// Script starting mode.
func (o LookupAutoscriptResultOutput) Start() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAutoscriptResult) string { return v.Start }).(pulumi.StringOutput)
}

// Maximum running time for this script in seconds (0 = no timeout).
func (o LookupAutoscriptResultOutput) Timeout() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAutoscriptResult) int { return v.Timeout }).(pulumi.IntOutput)
}

func (o LookupAutoscriptResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAutoscriptResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAutoscriptResultOutput{})
}