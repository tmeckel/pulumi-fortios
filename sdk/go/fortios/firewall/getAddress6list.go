// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firewall

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a list of `firewall.Address6`.
func GetAddress6list(ctx *pulumi.Context, args *GetAddress6listArgs, opts ...pulumi.InvokeOption) (*GetAddress6listResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetAddress6listResult
	err := ctx.Invoke("fortios:firewall/getAddress6list:getAddress6list", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAddress6list.
type GetAddress6listArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getAddress6list.
type GetAddress6listResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the `firewall.Address6`.
	Namelists []string `pulumi:"namelists"`
	Vdomparam *string  `pulumi:"vdomparam"`
}

func GetAddress6listOutput(ctx *pulumi.Context, args GetAddress6listOutputArgs, opts ...pulumi.InvokeOption) GetAddress6listResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAddress6listResult, error) {
			args := v.(GetAddress6listArgs)
			r, err := GetAddress6list(ctx, &args, opts...)
			var s GetAddress6listResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAddress6listResultOutput)
}

// A collection of arguments for invoking getAddress6list.
type GetAddress6listOutputArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetAddress6listOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAddress6listArgs)(nil)).Elem()
}

// A collection of values returned by getAddress6list.
type GetAddress6listResultOutput struct{ *pulumi.OutputState }

func (GetAddress6listResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAddress6listResult)(nil)).Elem()
}

func (o GetAddress6listResultOutput) ToGetAddress6listResultOutput() GetAddress6listResultOutput {
	return o
}

func (o GetAddress6listResultOutput) ToGetAddress6listResultOutputWithContext(ctx context.Context) GetAddress6listResultOutput {
	return o
}

func (o GetAddress6listResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAddress6listResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAddress6listResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAddress6listResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of the `firewall.Address6`.
func (o GetAddress6listResultOutput) Namelists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAddress6listResult) []string { return v.Namelists }).(pulumi.StringArrayOutput)
}

func (o GetAddress6listResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAddress6listResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAddress6listResultOutput{})
}