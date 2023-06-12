// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firewall

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios firewall internetservicegroup
func LookupInternetservicegroup(ctx *pulumi.Context, args *LookupInternetservicegroupArgs, opts ...pulumi.InvokeOption) (*LookupInternetservicegroupResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupInternetservicegroupResult
	err := ctx.Invoke("fortios:firewall/getInternetservicegroup:getInternetservicegroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInternetservicegroup.
type LookupInternetservicegroupArgs struct {
	// Specify the name of the desired firewall internetservicegroup.
	Name string `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getInternetservicegroup.
type LookupInternetservicegroupResult struct {
	// Comment.
	Comment string `pulumi:"comment"`
	// How this service may be used (source, destination or both).
	Direction string `pulumi:"direction"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Internet Service group member. The structure of `member` block is documented below.
	Members []GetInternetservicegroupMember `pulumi:"members"`
	// Internet Service name.
	Name      string  `pulumi:"name"`
	Vdomparam *string `pulumi:"vdomparam"`
}

func LookupInternetservicegroupOutput(ctx *pulumi.Context, args LookupInternetservicegroupOutputArgs, opts ...pulumi.InvokeOption) LookupInternetservicegroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupInternetservicegroupResult, error) {
			args := v.(LookupInternetservicegroupArgs)
			r, err := LookupInternetservicegroup(ctx, &args, opts...)
			var s LookupInternetservicegroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupInternetservicegroupResultOutput)
}

// A collection of arguments for invoking getInternetservicegroup.
type LookupInternetservicegroupOutputArgs struct {
	// Specify the name of the desired firewall internetservicegroup.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupInternetservicegroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInternetservicegroupArgs)(nil)).Elem()
}

// A collection of values returned by getInternetservicegroup.
type LookupInternetservicegroupResultOutput struct{ *pulumi.OutputState }

func (LookupInternetservicegroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInternetservicegroupResult)(nil)).Elem()
}

func (o LookupInternetservicegroupResultOutput) ToLookupInternetservicegroupResultOutput() LookupInternetservicegroupResultOutput {
	return o
}

func (o LookupInternetservicegroupResultOutput) ToLookupInternetservicegroupResultOutputWithContext(ctx context.Context) LookupInternetservicegroupResultOutput {
	return o
}

// Comment.
func (o LookupInternetservicegroupResultOutput) Comment() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInternetservicegroupResult) string { return v.Comment }).(pulumi.StringOutput)
}

// How this service may be used (source, destination or both).
func (o LookupInternetservicegroupResultOutput) Direction() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInternetservicegroupResult) string { return v.Direction }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupInternetservicegroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInternetservicegroupResult) string { return v.Id }).(pulumi.StringOutput)
}

// Internet Service group member. The structure of `member` block is documented below.
func (o LookupInternetservicegroupResultOutput) Members() GetInternetservicegroupMemberArrayOutput {
	return o.ApplyT(func(v LookupInternetservicegroupResult) []GetInternetservicegroupMember { return v.Members }).(GetInternetservicegroupMemberArrayOutput)
}

// Internet Service name.
func (o LookupInternetservicegroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInternetservicegroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupInternetservicegroupResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInternetservicegroupResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInternetservicegroupResultOutput{})
}
