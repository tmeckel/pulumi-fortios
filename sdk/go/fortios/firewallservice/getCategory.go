// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firewallservice

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios firewallservice category
func LookupCategory(ctx *pulumi.Context, args *LookupCategoryArgs, opts ...pulumi.InvokeOption) (*LookupCategoryResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupCategoryResult
	err := ctx.Invoke("fortios:firewallservice/getCategory:getCategory", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCategory.
type LookupCategoryArgs struct {
	// Specify the name of the desired firewallservice category.
	Name string `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getCategory.
type LookupCategoryResult struct {
	// Comment.
	Comment string `pulumi:"comment"`
	// Security Fabric global object setting.
	FabricObject string `pulumi:"fabricObject"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Service category name.
	Name      string  `pulumi:"name"`
	Vdomparam *string `pulumi:"vdomparam"`
}

func LookupCategoryOutput(ctx *pulumi.Context, args LookupCategoryOutputArgs, opts ...pulumi.InvokeOption) LookupCategoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCategoryResult, error) {
			args := v.(LookupCategoryArgs)
			r, err := LookupCategory(ctx, &args, opts...)
			var s LookupCategoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCategoryResultOutput)
}

// A collection of arguments for invoking getCategory.
type LookupCategoryOutputArgs struct {
	// Specify the name of the desired firewallservice category.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupCategoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCategoryArgs)(nil)).Elem()
}

// A collection of values returned by getCategory.
type LookupCategoryResultOutput struct{ *pulumi.OutputState }

func (LookupCategoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCategoryResult)(nil)).Elem()
}

func (o LookupCategoryResultOutput) ToLookupCategoryResultOutput() LookupCategoryResultOutput {
	return o
}

func (o LookupCategoryResultOutput) ToLookupCategoryResultOutputWithContext(ctx context.Context) LookupCategoryResultOutput {
	return o
}

// Comment.
func (o LookupCategoryResultOutput) Comment() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCategoryResult) string { return v.Comment }).(pulumi.StringOutput)
}

// Security Fabric global object setting.
func (o LookupCategoryResultOutput) FabricObject() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCategoryResult) string { return v.FabricObject }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupCategoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCategoryResult) string { return v.Id }).(pulumi.StringOutput)
}

// Service category name.
func (o LookupCategoryResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCategoryResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupCategoryResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCategoryResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCategoryResultOutput{})
}
