// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sys

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios system sessionhelper
func LookupSessionhelper(ctx *pulumi.Context, args *LookupSessionhelperArgs, opts ...pulumi.InvokeOption) (*LookupSessionhelperResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupSessionhelperResult
	err := ctx.Invoke("fortios:sys/getSessionhelper:getSessionhelper", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSessionhelper.
type LookupSessionhelperArgs struct {
	// Specify the fosid of the desired system sessionhelper.
	Fosid int `pulumi:"fosid"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getSessionhelper.
type LookupSessionhelperResult struct {
	// Session helper ID.
	Fosid int `pulumi:"fosid"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Helper name.
	Name string `pulumi:"name"`
	// Protocol port.
	Port int `pulumi:"port"`
	// Protocol number.
	Protocol  int     `pulumi:"protocol"`
	Vdomparam *string `pulumi:"vdomparam"`
}

func LookupSessionhelperOutput(ctx *pulumi.Context, args LookupSessionhelperOutputArgs, opts ...pulumi.InvokeOption) LookupSessionhelperResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSessionhelperResult, error) {
			args := v.(LookupSessionhelperArgs)
			r, err := LookupSessionhelper(ctx, &args, opts...)
			var s LookupSessionhelperResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSessionhelperResultOutput)
}

// A collection of arguments for invoking getSessionhelper.
type LookupSessionhelperOutputArgs struct {
	// Specify the fosid of the desired system sessionhelper.
	Fosid pulumi.IntInput `pulumi:"fosid"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupSessionhelperOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSessionhelperArgs)(nil)).Elem()
}

// A collection of values returned by getSessionhelper.
type LookupSessionhelperResultOutput struct{ *pulumi.OutputState }

func (LookupSessionhelperResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSessionhelperResult)(nil)).Elem()
}

func (o LookupSessionhelperResultOutput) ToLookupSessionhelperResultOutput() LookupSessionhelperResultOutput {
	return o
}

func (o LookupSessionhelperResultOutput) ToLookupSessionhelperResultOutputWithContext(ctx context.Context) LookupSessionhelperResultOutput {
	return o
}

// Session helper ID.
func (o LookupSessionhelperResultOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSessionhelperResult) int { return v.Fosid }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSessionhelperResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSessionhelperResult) string { return v.Id }).(pulumi.StringOutput)
}

// Helper name.
func (o LookupSessionhelperResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSessionhelperResult) string { return v.Name }).(pulumi.StringOutput)
}

// Protocol port.
func (o LookupSessionhelperResultOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSessionhelperResult) int { return v.Port }).(pulumi.IntOutput)
}

// Protocol number.
func (o LookupSessionhelperResultOutput) Protocol() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSessionhelperResult) int { return v.Protocol }).(pulumi.IntOutput)
}

func (o LookupSessionhelperResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSessionhelperResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSessionhelperResultOutput{})
}