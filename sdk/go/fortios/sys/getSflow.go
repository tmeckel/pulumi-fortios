// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sys

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on fortios system sflow
func LookupSflow(ctx *pulumi.Context, args *LookupSflowArgs, opts ...pulumi.InvokeOption) (*LookupSflowResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupSflowResult
	err := ctx.Invoke("fortios:sys/getSflow:getSflow", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSflow.
type LookupSflowArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getSflow.
type LookupSflowResult struct {
	// IP address of the sFlow collector that sFlow agents added to interfaces in this VDOM send sFlow datagrams to (default = 0.0.0.0).
	CollectorIp string `pulumi:"collectorIp"`
	// UDP port number used for sending sFlow datagrams (configure only if required by your sFlow collector or your network configuration) (0 - 65535, default = 6343).
	CollectorPort int `pulumi:"collectorPort"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Specify outgoing interface to reach server.
	Interface string `pulumi:"interface"`
	// Specify how to select outgoing interface to reach server.
	InterfaceSelectMethod string `pulumi:"interfaceSelectMethod"`
	// Source IP address for sFlow agent.
	SourceIp  string  `pulumi:"sourceIp"`
	Vdomparam *string `pulumi:"vdomparam"`
}

func LookupSflowOutput(ctx *pulumi.Context, args LookupSflowOutputArgs, opts ...pulumi.InvokeOption) LookupSflowResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSflowResult, error) {
			args := v.(LookupSflowArgs)
			r, err := LookupSflow(ctx, &args, opts...)
			var s LookupSflowResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSflowResultOutput)
}

// A collection of arguments for invoking getSflow.
type LookupSflowOutputArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupSflowOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSflowArgs)(nil)).Elem()
}

// A collection of values returned by getSflow.
type LookupSflowResultOutput struct{ *pulumi.OutputState }

func (LookupSflowResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSflowResult)(nil)).Elem()
}

func (o LookupSflowResultOutput) ToLookupSflowResultOutput() LookupSflowResultOutput {
	return o
}

func (o LookupSflowResultOutput) ToLookupSflowResultOutputWithContext(ctx context.Context) LookupSflowResultOutput {
	return o
}

// IP address of the sFlow collector that sFlow agents added to interfaces in this VDOM send sFlow datagrams to (default = 0.0.0.0).
func (o LookupSflowResultOutput) CollectorIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSflowResult) string { return v.CollectorIp }).(pulumi.StringOutput)
}

// UDP port number used for sending sFlow datagrams (configure only if required by your sFlow collector or your network configuration) (0 - 65535, default = 6343).
func (o LookupSflowResultOutput) CollectorPort() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSflowResult) int { return v.CollectorPort }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSflowResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSflowResult) string { return v.Id }).(pulumi.StringOutput)
}

// Specify outgoing interface to reach server.
func (o LookupSflowResultOutput) Interface() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSflowResult) string { return v.Interface }).(pulumi.StringOutput)
}

// Specify how to select outgoing interface to reach server.
func (o LookupSflowResultOutput) InterfaceSelectMethod() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSflowResult) string { return v.InterfaceSelectMethod }).(pulumi.StringOutput)
}

// Source IP address for sFlow agent.
func (o LookupSflowResultOutput) SourceIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSflowResult) string { return v.SourceIp }).(pulumi.StringOutput)
}

func (o LookupSflowResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSflowResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSflowResultOutput{})
}
