// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firewallschedule

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios firewallschedule onetime
func LookupOnetime(ctx *pulumi.Context, args *LookupOnetimeArgs, opts ...pulumi.InvokeOption) (*LookupOnetimeResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupOnetimeResult
	err := ctx.Invoke("fortios:firewallschedule/getOnetime:getOnetime", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOnetime.
type LookupOnetimeArgs struct {
	// Specify the name of the desired firewallschedule onetime.
	Name string `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getOnetime.
type LookupOnetimeResult struct {
	// Color of icon on the GUI.
	Color int `pulumi:"color"`
	// Schedule end date and time, format hh:mm yyyy/mm/dd.
	End string `pulumi:"end"`
	// Write an event log message this many days before the schedule expires.
	ExpirationDays int `pulumi:"expirationDays"`
	// Security Fabric global object setting.
	FabricObject string `pulumi:"fabricObject"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Onetime schedule name.
	Name string `pulumi:"name"`
	// Schedule start date and time, format hh:mm yyyy/mm/dd.
	Start     string  `pulumi:"start"`
	Vdomparam *string `pulumi:"vdomparam"`
}

func LookupOnetimeOutput(ctx *pulumi.Context, args LookupOnetimeOutputArgs, opts ...pulumi.InvokeOption) LookupOnetimeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOnetimeResult, error) {
			args := v.(LookupOnetimeArgs)
			r, err := LookupOnetime(ctx, &args, opts...)
			var s LookupOnetimeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupOnetimeResultOutput)
}

// A collection of arguments for invoking getOnetime.
type LookupOnetimeOutputArgs struct {
	// Specify the name of the desired firewallschedule onetime.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupOnetimeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOnetimeArgs)(nil)).Elem()
}

// A collection of values returned by getOnetime.
type LookupOnetimeResultOutput struct{ *pulumi.OutputState }

func (LookupOnetimeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOnetimeResult)(nil)).Elem()
}

func (o LookupOnetimeResultOutput) ToLookupOnetimeResultOutput() LookupOnetimeResultOutput {
	return o
}

func (o LookupOnetimeResultOutput) ToLookupOnetimeResultOutputWithContext(ctx context.Context) LookupOnetimeResultOutput {
	return o
}

// Color of icon on the GUI.
func (o LookupOnetimeResultOutput) Color() pulumi.IntOutput {
	return o.ApplyT(func(v LookupOnetimeResult) int { return v.Color }).(pulumi.IntOutput)
}

// Schedule end date and time, format hh:mm yyyy/mm/dd.
func (o LookupOnetimeResultOutput) End() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOnetimeResult) string { return v.End }).(pulumi.StringOutput)
}

// Write an event log message this many days before the schedule expires.
func (o LookupOnetimeResultOutput) ExpirationDays() pulumi.IntOutput {
	return o.ApplyT(func(v LookupOnetimeResult) int { return v.ExpirationDays }).(pulumi.IntOutput)
}

// Security Fabric global object setting.
func (o LookupOnetimeResultOutput) FabricObject() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOnetimeResult) string { return v.FabricObject }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupOnetimeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOnetimeResult) string { return v.Id }).(pulumi.StringOutput)
}

// Onetime schedule name.
func (o LookupOnetimeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOnetimeResult) string { return v.Name }).(pulumi.StringOutput)
}

// Schedule start date and time, format hh:mm yyyy/mm/dd.
func (o LookupOnetimeResultOutput) Start() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOnetimeResult) string { return v.Start }).(pulumi.StringOutput)
}

func (o LookupOnetimeResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupOnetimeResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOnetimeResultOutput{})
}
