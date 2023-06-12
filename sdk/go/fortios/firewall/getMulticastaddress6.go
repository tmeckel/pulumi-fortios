// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firewall

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios firewall multicastaddress6
func LookupMulticastaddress6(ctx *pulumi.Context, args *LookupMulticastaddress6Args, opts ...pulumi.InvokeOption) (*LookupMulticastaddress6Result, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupMulticastaddress6Result
	err := ctx.Invoke("fortios:firewall/getMulticastaddress6:getMulticastaddress6", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getMulticastaddress6.
type LookupMulticastaddress6Args struct {
	// Specify the name of the desired firewall multicastaddress6.
	Name string `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getMulticastaddress6.
type LookupMulticastaddress6Result struct {
	// Color of icon on the GUI.
	Color int `pulumi:"color"`
	// Comment.
	Comment string `pulumi:"comment"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// IPv6 address prefix (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx).
	Ip6 string `pulumi:"ip6"`
	// Tag name.
	Name string `pulumi:"name"`
	// Config object tagging. The structure of `tagging` block is documented below.
	Taggings  []GetMulticastaddress6Tagging `pulumi:"taggings"`
	Vdomparam *string                       `pulumi:"vdomparam"`
	// Enable/disable visibility of the IPv6 multicast address on the GUI.
	Visibility string `pulumi:"visibility"`
}

func LookupMulticastaddress6Output(ctx *pulumi.Context, args LookupMulticastaddress6OutputArgs, opts ...pulumi.InvokeOption) LookupMulticastaddress6ResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMulticastaddress6Result, error) {
			args := v.(LookupMulticastaddress6Args)
			r, err := LookupMulticastaddress6(ctx, &args, opts...)
			var s LookupMulticastaddress6Result
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMulticastaddress6ResultOutput)
}

// A collection of arguments for invoking getMulticastaddress6.
type LookupMulticastaddress6OutputArgs struct {
	// Specify the name of the desired firewall multicastaddress6.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupMulticastaddress6OutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMulticastaddress6Args)(nil)).Elem()
}

// A collection of values returned by getMulticastaddress6.
type LookupMulticastaddress6ResultOutput struct{ *pulumi.OutputState }

func (LookupMulticastaddress6ResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMulticastaddress6Result)(nil)).Elem()
}

func (o LookupMulticastaddress6ResultOutput) ToLookupMulticastaddress6ResultOutput() LookupMulticastaddress6ResultOutput {
	return o
}

func (o LookupMulticastaddress6ResultOutput) ToLookupMulticastaddress6ResultOutputWithContext(ctx context.Context) LookupMulticastaddress6ResultOutput {
	return o
}

// Color of icon on the GUI.
func (o LookupMulticastaddress6ResultOutput) Color() pulumi.IntOutput {
	return o.ApplyT(func(v LookupMulticastaddress6Result) int { return v.Color }).(pulumi.IntOutput)
}

// Comment.
func (o LookupMulticastaddress6ResultOutput) Comment() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMulticastaddress6Result) string { return v.Comment }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupMulticastaddress6ResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMulticastaddress6Result) string { return v.Id }).(pulumi.StringOutput)
}

// IPv6 address prefix (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx).
func (o LookupMulticastaddress6ResultOutput) Ip6() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMulticastaddress6Result) string { return v.Ip6 }).(pulumi.StringOutput)
}

// Tag name.
func (o LookupMulticastaddress6ResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMulticastaddress6Result) string { return v.Name }).(pulumi.StringOutput)
}

// Config object tagging. The structure of `tagging` block is documented below.
func (o LookupMulticastaddress6ResultOutput) Taggings() GetMulticastaddress6TaggingArrayOutput {
	return o.ApplyT(func(v LookupMulticastaddress6Result) []GetMulticastaddress6Tagging { return v.Taggings }).(GetMulticastaddress6TaggingArrayOutput)
}

func (o LookupMulticastaddress6ResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMulticastaddress6Result) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// Enable/disable visibility of the IPv6 multicast address on the GUI.
func (o LookupMulticastaddress6ResultOutput) Visibility() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMulticastaddress6Result) string { return v.Visibility }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMulticastaddress6ResultOutput{})
}
