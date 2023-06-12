// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firewallshaper

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a list of `firewallshaper.Peripshaper`.
func GetPeripshaperlist(ctx *pulumi.Context, args *GetPeripshaperlistArgs, opts ...pulumi.InvokeOption) (*GetPeripshaperlistResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetPeripshaperlistResult
	err := ctx.Invoke("fortios:firewallshaper/getPeripshaperlist:getPeripshaperlist", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPeripshaperlist.
type GetPeripshaperlistArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getPeripshaperlist.
type GetPeripshaperlistResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the `firewallshaper.Peripshaper`.
	Namelists []string `pulumi:"namelists"`
	Vdomparam *string  `pulumi:"vdomparam"`
}

func GetPeripshaperlistOutput(ctx *pulumi.Context, args GetPeripshaperlistOutputArgs, opts ...pulumi.InvokeOption) GetPeripshaperlistResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetPeripshaperlistResult, error) {
			args := v.(GetPeripshaperlistArgs)
			r, err := GetPeripshaperlist(ctx, &args, opts...)
			var s GetPeripshaperlistResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetPeripshaperlistResultOutput)
}

// A collection of arguments for invoking getPeripshaperlist.
type GetPeripshaperlistOutputArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetPeripshaperlistOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPeripshaperlistArgs)(nil)).Elem()
}

// A collection of values returned by getPeripshaperlist.
type GetPeripshaperlistResultOutput struct{ *pulumi.OutputState }

func (GetPeripshaperlistResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPeripshaperlistResult)(nil)).Elem()
}

func (o GetPeripshaperlistResultOutput) ToGetPeripshaperlistResultOutput() GetPeripshaperlistResultOutput {
	return o
}

func (o GetPeripshaperlistResultOutput) ToGetPeripshaperlistResultOutputWithContext(ctx context.Context) GetPeripshaperlistResultOutput {
	return o
}

func (o GetPeripshaperlistResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPeripshaperlistResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetPeripshaperlistResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetPeripshaperlistResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of the `firewallshaper.Peripshaper`.
func (o GetPeripshaperlistResultOutput) Namelists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetPeripshaperlistResult) []string { return v.Namelists }).(pulumi.StringArrayOutput)
}

func (o GetPeripshaperlistResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPeripshaperlistResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetPeripshaperlistResultOutput{})
}
