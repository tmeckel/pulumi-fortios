// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firewall

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a list of `firewall.Address`.
func GetAddresslist(ctx *pulumi.Context, args *GetAddresslistArgs, opts ...pulumi.InvokeOption) (*GetAddresslistResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetAddresslistResult
	err := ctx.Invoke("fortios:firewall/getAddresslist:getAddresslist", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAddresslist.
type GetAddresslistArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getAddresslist.
type GetAddresslistResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the `firewall.Address`.
	Namelists []string `pulumi:"namelists"`
	Vdomparam *string  `pulumi:"vdomparam"`
}

func GetAddresslistOutput(ctx *pulumi.Context, args GetAddresslistOutputArgs, opts ...pulumi.InvokeOption) GetAddresslistResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAddresslistResult, error) {
			args := v.(GetAddresslistArgs)
			r, err := GetAddresslist(ctx, &args, opts...)
			var s GetAddresslistResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAddresslistResultOutput)
}

// A collection of arguments for invoking getAddresslist.
type GetAddresslistOutputArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetAddresslistOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAddresslistArgs)(nil)).Elem()
}

// A collection of values returned by getAddresslist.
type GetAddresslistResultOutput struct{ *pulumi.OutputState }

func (GetAddresslistResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAddresslistResult)(nil)).Elem()
}

func (o GetAddresslistResultOutput) ToGetAddresslistResultOutput() GetAddresslistResultOutput {
	return o
}

func (o GetAddresslistResultOutput) ToGetAddresslistResultOutputWithContext(ctx context.Context) GetAddresslistResultOutput {
	return o
}

func (o GetAddresslistResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAddresslistResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAddresslistResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAddresslistResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of the `firewall.Address`.
func (o GetAddresslistResultOutput) Namelists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAddresslistResult) []string { return v.Namelists }).(pulumi.StringArrayOutput)
}

func (o GetAddresslistResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAddresslistResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAddresslistResultOutput{})
}
