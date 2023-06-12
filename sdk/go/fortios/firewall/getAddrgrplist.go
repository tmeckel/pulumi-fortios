// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firewall

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a list of `firewall.Addrgrp`.
func GetAddrgrplist(ctx *pulumi.Context, args *GetAddrgrplistArgs, opts ...pulumi.InvokeOption) (*GetAddrgrplistResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetAddrgrplistResult
	err := ctx.Invoke("fortios:firewall/getAddrgrplist:getAddrgrplist", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAddrgrplist.
type GetAddrgrplistArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getAddrgrplist.
type GetAddrgrplistResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the `firewall.Addrgrp`.
	Namelists []string `pulumi:"namelists"`
	Vdomparam *string  `pulumi:"vdomparam"`
}

func GetAddrgrplistOutput(ctx *pulumi.Context, args GetAddrgrplistOutputArgs, opts ...pulumi.InvokeOption) GetAddrgrplistResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAddrgrplistResult, error) {
			args := v.(GetAddrgrplistArgs)
			r, err := GetAddrgrplist(ctx, &args, opts...)
			var s GetAddrgrplistResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAddrgrplistResultOutput)
}

// A collection of arguments for invoking getAddrgrplist.
type GetAddrgrplistOutputArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetAddrgrplistOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAddrgrplistArgs)(nil)).Elem()
}

// A collection of values returned by getAddrgrplist.
type GetAddrgrplistResultOutput struct{ *pulumi.OutputState }

func (GetAddrgrplistResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAddrgrplistResult)(nil)).Elem()
}

func (o GetAddrgrplistResultOutput) ToGetAddrgrplistResultOutput() GetAddrgrplistResultOutput {
	return o
}

func (o GetAddrgrplistResultOutput) ToGetAddrgrplistResultOutputWithContext(ctx context.Context) GetAddrgrplistResultOutput {
	return o
}

func (o GetAddrgrplistResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAddrgrplistResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAddrgrplistResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAddrgrplistResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of the `firewall.Addrgrp`.
func (o GetAddrgrplistResultOutput) Namelists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAddrgrplistResult) []string { return v.Namelists }).(pulumi.StringArrayOutput)
}

func (o GetAddrgrplistResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAddrgrplistResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAddrgrplistResultOutput{})
}