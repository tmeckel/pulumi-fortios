// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package router

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios router accesslist6
func LookupAccesslist6(ctx *pulumi.Context, args *LookupAccesslist6Args, opts ...pulumi.InvokeOption) (*LookupAccesslist6Result, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupAccesslist6Result
	err := ctx.Invoke("fortios:router/getAccesslist6:getAccesslist6", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAccesslist6.
type LookupAccesslist6Args struct {
	// Specify the name of the desired router accesslist6.
	Name string `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getAccesslist6.
type LookupAccesslist6Result struct {
	// Comment.
	Comments string `pulumi:"comments"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Name.
	Name string `pulumi:"name"`
	// Rule. The structure of `rule` block is documented below.
	Rules     []GetAccesslist6Rule `pulumi:"rules"`
	Vdomparam *string              `pulumi:"vdomparam"`
}

func LookupAccesslist6Output(ctx *pulumi.Context, args LookupAccesslist6OutputArgs, opts ...pulumi.InvokeOption) LookupAccesslist6ResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAccesslist6Result, error) {
			args := v.(LookupAccesslist6Args)
			r, err := LookupAccesslist6(ctx, &args, opts...)
			var s LookupAccesslist6Result
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAccesslist6ResultOutput)
}

// A collection of arguments for invoking getAccesslist6.
type LookupAccesslist6OutputArgs struct {
	// Specify the name of the desired router accesslist6.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupAccesslist6OutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccesslist6Args)(nil)).Elem()
}

// A collection of values returned by getAccesslist6.
type LookupAccesslist6ResultOutput struct{ *pulumi.OutputState }

func (LookupAccesslist6ResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccesslist6Result)(nil)).Elem()
}

func (o LookupAccesslist6ResultOutput) ToLookupAccesslist6ResultOutput() LookupAccesslist6ResultOutput {
	return o
}

func (o LookupAccesslist6ResultOutput) ToLookupAccesslist6ResultOutputWithContext(ctx context.Context) LookupAccesslist6ResultOutput {
	return o
}

// Comment.
func (o LookupAccesslist6ResultOutput) Comments() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccesslist6Result) string { return v.Comments }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupAccesslist6ResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccesslist6Result) string { return v.Id }).(pulumi.StringOutput)
}

// Name.
func (o LookupAccesslist6ResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccesslist6Result) string { return v.Name }).(pulumi.StringOutput)
}

// Rule. The structure of `rule` block is documented below.
func (o LookupAccesslist6ResultOutput) Rules() GetAccesslist6RuleArrayOutput {
	return o.ApplyT(func(v LookupAccesslist6Result) []GetAccesslist6Rule { return v.Rules }).(GetAccesslist6RuleArrayOutput)
}

func (o LookupAccesslist6ResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccesslist6Result) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAccesslist6ResultOutput{})
}
