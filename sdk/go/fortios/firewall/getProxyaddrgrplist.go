// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firewall

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a list of `firewall.Proxyaddrgrp`.
func GetProxyaddrgrplist(ctx *pulumi.Context, args *GetProxyaddrgrplistArgs, opts ...pulumi.InvokeOption) (*GetProxyaddrgrplistResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetProxyaddrgrplistResult
	err := ctx.Invoke("fortios:firewall/getProxyaddrgrplist:getProxyaddrgrplist", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProxyaddrgrplist.
type GetProxyaddrgrplistArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getProxyaddrgrplist.
type GetProxyaddrgrplistResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the `firewall.Proxyaddrgrp`.
	Namelists []string `pulumi:"namelists"`
	Vdomparam *string  `pulumi:"vdomparam"`
}

func GetProxyaddrgrplistOutput(ctx *pulumi.Context, args GetProxyaddrgrplistOutputArgs, opts ...pulumi.InvokeOption) GetProxyaddrgrplistResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetProxyaddrgrplistResult, error) {
			args := v.(GetProxyaddrgrplistArgs)
			r, err := GetProxyaddrgrplist(ctx, &args, opts...)
			var s GetProxyaddrgrplistResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetProxyaddrgrplistResultOutput)
}

// A collection of arguments for invoking getProxyaddrgrplist.
type GetProxyaddrgrplistOutputArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetProxyaddrgrplistOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProxyaddrgrplistArgs)(nil)).Elem()
}

// A collection of values returned by getProxyaddrgrplist.
type GetProxyaddrgrplistResultOutput struct{ *pulumi.OutputState }

func (GetProxyaddrgrplistResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProxyaddrgrplistResult)(nil)).Elem()
}

func (o GetProxyaddrgrplistResultOutput) ToGetProxyaddrgrplistResultOutput() GetProxyaddrgrplistResultOutput {
	return o
}

func (o GetProxyaddrgrplistResultOutput) ToGetProxyaddrgrplistResultOutputWithContext(ctx context.Context) GetProxyaddrgrplistResultOutput {
	return o
}

func (o GetProxyaddrgrplistResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProxyaddrgrplistResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetProxyaddrgrplistResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetProxyaddrgrplistResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of the `firewall.Proxyaddrgrp`.
func (o GetProxyaddrgrplistResultOutput) Namelists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetProxyaddrgrplistResult) []string { return v.Namelists }).(pulumi.StringArrayOutput)
}

func (o GetProxyaddrgrplistResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProxyaddrgrplistResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetProxyaddrgrplistResultOutput{})
}
