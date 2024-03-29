// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package router

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios router prefixlist6
func LookupPrefixlist6(ctx *pulumi.Context, args *LookupPrefixlist6Args, opts ...pulumi.InvokeOption) (*LookupPrefixlist6Result, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupPrefixlist6Result
	err := ctx.Invoke("fortios:router/getPrefixlist6:getPrefixlist6", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPrefixlist6.
type LookupPrefixlist6Args struct {
	// Specify the name of the desired router prefixlist6.
	Name string `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getPrefixlist6.
type LookupPrefixlist6Result struct {
	// Comment.
	Comments string `pulumi:"comments"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Name.
	Name string `pulumi:"name"`
	// IPv6 prefix list rule. The structure of `rule` block is documented below.
	Rules     []GetPrefixlist6Rule `pulumi:"rules"`
	Vdomparam *string              `pulumi:"vdomparam"`
}

func LookupPrefixlist6Output(ctx *pulumi.Context, args LookupPrefixlist6OutputArgs, opts ...pulumi.InvokeOption) LookupPrefixlist6ResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPrefixlist6Result, error) {
			args := v.(LookupPrefixlist6Args)
			r, err := LookupPrefixlist6(ctx, &args, opts...)
			var s LookupPrefixlist6Result
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPrefixlist6ResultOutput)
}

// A collection of arguments for invoking getPrefixlist6.
type LookupPrefixlist6OutputArgs struct {
	// Specify the name of the desired router prefixlist6.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupPrefixlist6OutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrefixlist6Args)(nil)).Elem()
}

// A collection of values returned by getPrefixlist6.
type LookupPrefixlist6ResultOutput struct{ *pulumi.OutputState }

func (LookupPrefixlist6ResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrefixlist6Result)(nil)).Elem()
}

func (o LookupPrefixlist6ResultOutput) ToLookupPrefixlist6ResultOutput() LookupPrefixlist6ResultOutput {
	return o
}

func (o LookupPrefixlist6ResultOutput) ToLookupPrefixlist6ResultOutputWithContext(ctx context.Context) LookupPrefixlist6ResultOutput {
	return o
}

// Comment.
func (o LookupPrefixlist6ResultOutput) Comments() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrefixlist6Result) string { return v.Comments }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupPrefixlist6ResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrefixlist6Result) string { return v.Id }).(pulumi.StringOutput)
}

// Name.
func (o LookupPrefixlist6ResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrefixlist6Result) string { return v.Name }).(pulumi.StringOutput)
}

// IPv6 prefix list rule. The structure of `rule` block is documented below.
func (o LookupPrefixlist6ResultOutput) Rules() GetPrefixlist6RuleArrayOutput {
	return o.ApplyT(func(v LookupPrefixlist6Result) []GetPrefixlist6Rule { return v.Rules }).(GetPrefixlist6RuleArrayOutput)
}

func (o LookupPrefixlist6ResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrefixlist6Result) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPrefixlist6ResultOutput{})
}
