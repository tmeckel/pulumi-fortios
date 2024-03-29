// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firewallservice

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios firewallservice group
func LookupGroup(ctx *pulumi.Context, args *LookupGroupArgs, opts ...pulumi.InvokeOption) (*LookupGroupResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupGroupResult
	err := ctx.Invoke("fortios:firewallservice/getGroup:getGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGroup.
type LookupGroupArgs struct {
	// Specify the name of the desired firewallservice group.
	Name string `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getGroup.
type LookupGroupResult struct {
	// Color of icon on the GUI.
	Color int `pulumi:"color"`
	// Comment.
	Comment string `pulumi:"comment"`
	// Security Fabric global object setting.
	FabricObject string `pulumi:"fabricObject"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Service objects contained within the group. The structure of `member` block is documented below.
	Members []GetGroupMember `pulumi:"members"`
	// Address name.
	Name string `pulumi:"name"`
	// Enable/disable web proxy service group.
	Proxy     string  `pulumi:"proxy"`
	Vdomparam *string `pulumi:"vdomparam"`
}

func LookupGroupOutput(ctx *pulumi.Context, args LookupGroupOutputArgs, opts ...pulumi.InvokeOption) LookupGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGroupResult, error) {
			args := v.(LookupGroupArgs)
			r, err := LookupGroup(ctx, &args, opts...)
			var s LookupGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGroupResultOutput)
}

// A collection of arguments for invoking getGroup.
type LookupGroupOutputArgs struct {
	// Specify the name of the desired firewallservice group.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGroupArgs)(nil)).Elem()
}

// A collection of values returned by getGroup.
type LookupGroupResultOutput struct{ *pulumi.OutputState }

func (LookupGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGroupResult)(nil)).Elem()
}

func (o LookupGroupResultOutput) ToLookupGroupResultOutput() LookupGroupResultOutput {
	return o
}

func (o LookupGroupResultOutput) ToLookupGroupResultOutputWithContext(ctx context.Context) LookupGroupResultOutput {
	return o
}

// Color of icon on the GUI.
func (o LookupGroupResultOutput) Color() pulumi.IntOutput {
	return o.ApplyT(func(v LookupGroupResult) int { return v.Color }).(pulumi.IntOutput)
}

// Comment.
func (o LookupGroupResultOutput) Comment() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.Comment }).(pulumi.StringOutput)
}

// Security Fabric global object setting.
func (o LookupGroupResultOutput) FabricObject() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.FabricObject }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

// Service objects contained within the group. The structure of `member` block is documented below.
func (o LookupGroupResultOutput) Members() GetGroupMemberArrayOutput {
	return o.ApplyT(func(v LookupGroupResult) []GetGroupMember { return v.Members }).(GetGroupMemberArrayOutput)
}

// Address name.
func (o LookupGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

// Enable/disable web proxy service group.
func (o LookupGroupResultOutput) Proxy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.Proxy }).(pulumi.StringOutput)
}

func (o LookupGroupResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGroupResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGroupResultOutput{})
}
