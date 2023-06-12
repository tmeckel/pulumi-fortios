// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package router

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a list of `router.Communitylist`.
func GetCommunitylistlist(ctx *pulumi.Context, args *GetCommunitylistlistArgs, opts ...pulumi.InvokeOption) (*GetCommunitylistlistResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetCommunitylistlistResult
	err := ctx.Invoke("fortios:router/getCommunitylistlist:getCommunitylistlist", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCommunitylistlist.
type GetCommunitylistlistArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getCommunitylistlist.
type GetCommunitylistlistResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the `router.Communitylist`.
	Namelists []string `pulumi:"namelists"`
	Vdomparam *string  `pulumi:"vdomparam"`
}

func GetCommunitylistlistOutput(ctx *pulumi.Context, args GetCommunitylistlistOutputArgs, opts ...pulumi.InvokeOption) GetCommunitylistlistResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetCommunitylistlistResult, error) {
			args := v.(GetCommunitylistlistArgs)
			r, err := GetCommunitylistlist(ctx, &args, opts...)
			var s GetCommunitylistlistResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetCommunitylistlistResultOutput)
}

// A collection of arguments for invoking getCommunitylistlist.
type GetCommunitylistlistOutputArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetCommunitylistlistOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCommunitylistlistArgs)(nil)).Elem()
}

// A collection of values returned by getCommunitylistlist.
type GetCommunitylistlistResultOutput struct{ *pulumi.OutputState }

func (GetCommunitylistlistResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCommunitylistlistResult)(nil)).Elem()
}

func (o GetCommunitylistlistResultOutput) ToGetCommunitylistlistResultOutput() GetCommunitylistlistResultOutput {
	return o
}

func (o GetCommunitylistlistResultOutput) ToGetCommunitylistlistResultOutputWithContext(ctx context.Context) GetCommunitylistlistResultOutput {
	return o
}

func (o GetCommunitylistlistResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCommunitylistlistResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetCommunitylistlistResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetCommunitylistlistResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of the `router.Communitylist`.
func (o GetCommunitylistlistResultOutput) Namelists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetCommunitylistlistResult) []string { return v.Namelists }).(pulumi.StringArrayOutput)
}

func (o GetCommunitylistlistResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCommunitylistlistResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCommunitylistlistResultOutput{})
}
