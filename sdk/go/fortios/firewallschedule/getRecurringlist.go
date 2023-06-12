// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firewallschedule

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a list of `firewallschedule.Recurring`.
func GetRecurringlist(ctx *pulumi.Context, args *GetRecurringlistArgs, opts ...pulumi.InvokeOption) (*GetRecurringlistResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetRecurringlistResult
	err := ctx.Invoke("fortios:firewallschedule/getRecurringlist:getRecurringlist", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRecurringlist.
type GetRecurringlistArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getRecurringlist.
type GetRecurringlistResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the `firewallschedule.Recurring`.
	Namelists []string `pulumi:"namelists"`
	Vdomparam *string  `pulumi:"vdomparam"`
}

func GetRecurringlistOutput(ctx *pulumi.Context, args GetRecurringlistOutputArgs, opts ...pulumi.InvokeOption) GetRecurringlistResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetRecurringlistResult, error) {
			args := v.(GetRecurringlistArgs)
			r, err := GetRecurringlist(ctx, &args, opts...)
			var s GetRecurringlistResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetRecurringlistResultOutput)
}

// A collection of arguments for invoking getRecurringlist.
type GetRecurringlistOutputArgs struct {
	// A filter used to scope the list. See Filter results of datasource.
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetRecurringlistOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRecurringlistArgs)(nil)).Elem()
}

// A collection of values returned by getRecurringlist.
type GetRecurringlistResultOutput struct{ *pulumi.OutputState }

func (GetRecurringlistResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRecurringlistResult)(nil)).Elem()
}

func (o GetRecurringlistResultOutput) ToGetRecurringlistResultOutput() GetRecurringlistResultOutput {
	return o
}

func (o GetRecurringlistResultOutput) ToGetRecurringlistResultOutputWithContext(ctx context.Context) GetRecurringlistResultOutput {
	return o
}

func (o GetRecurringlistResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRecurringlistResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetRecurringlistResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetRecurringlistResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of the `firewallschedule.Recurring`.
func (o GetRecurringlistResultOutput) Namelists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetRecurringlistResult) []string { return v.Namelists }).(pulumi.StringArrayOutput)
}

func (o GetRecurringlistResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRecurringlistResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetRecurringlistResultOutput{})
}