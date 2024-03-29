// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package router

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on fortios router ripng
func LookupRipng(ctx *pulumi.Context, args *LookupRipngArgs, opts ...pulumi.InvokeOption) (*LookupRipngResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupRipngResult
	err := ctx.Invoke("fortios:router/getRipng:getRipng", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRipng.
type LookupRipngArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getRipng.
type LookupRipngResult struct {
	// Aggregate address. The structure of `aggregateAddress` block is documented below.
	AggregateAddresses []GetRipngAggregateAddress `pulumi:"aggregateAddresses"`
	// Enable/disable generation of default route.
	DefaultInformationOriginate string `pulumi:"defaultInformationOriginate"`
	// Default metric.
	DefaultMetric int `pulumi:"defaultMetric"`
	// Distance (1 - 255).
	Distances []GetRipngDistance `pulumi:"distances"`
	// Distribute list. The structure of `distributeList` block is documented below.
	DistributeLists []GetRipngDistributeList `pulumi:"distributeLists"`
	// Garbage timer.
	GarbageTimer int `pulumi:"garbageTimer"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Interface name.
	Interfaces []GetRipngInterface `pulumi:"interfaces"`
	// Maximum metric allowed to output(0 means 'not set').
	MaxOutMetric int `pulumi:"maxOutMetric"`
	// neighbor The structure of `neighbor` block is documented below.
	Neighbors []GetRipngNeighbor `pulumi:"neighbors"`
	// Network. The structure of `network` block is documented below.
	Networks []GetRipngNetwork `pulumi:"networks"`
	// Offset list. The structure of `offsetList` block is documented below.
	OffsetLists []GetRipngOffsetList `pulumi:"offsetLists"`
	// Passive interface configuration. The structure of `passiveInterface` block is documented below.
	PassiveInterfaces []GetRipngPassiveInterface `pulumi:"passiveInterfaces"`
	// Redistribute configuration. The structure of `redistribute` block is documented below.
	Redistributes []GetRipngRedistribute `pulumi:"redistributes"`
	// Timeout timer.
	TimeoutTimer int `pulumi:"timeoutTimer"`
	// Update timer.
	UpdateTimer int     `pulumi:"updateTimer"`
	Vdomparam   *string `pulumi:"vdomparam"`
}

func LookupRipngOutput(ctx *pulumi.Context, args LookupRipngOutputArgs, opts ...pulumi.InvokeOption) LookupRipngResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRipngResult, error) {
			args := v.(LookupRipngArgs)
			r, err := LookupRipng(ctx, &args, opts...)
			var s LookupRipngResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRipngResultOutput)
}

// A collection of arguments for invoking getRipng.
type LookupRipngOutputArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupRipngOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRipngArgs)(nil)).Elem()
}

// A collection of values returned by getRipng.
type LookupRipngResultOutput struct{ *pulumi.OutputState }

func (LookupRipngResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRipngResult)(nil)).Elem()
}

func (o LookupRipngResultOutput) ToLookupRipngResultOutput() LookupRipngResultOutput {
	return o
}

func (o LookupRipngResultOutput) ToLookupRipngResultOutputWithContext(ctx context.Context) LookupRipngResultOutput {
	return o
}

// Aggregate address. The structure of `aggregateAddress` block is documented below.
func (o LookupRipngResultOutput) AggregateAddresses() GetRipngAggregateAddressArrayOutput {
	return o.ApplyT(func(v LookupRipngResult) []GetRipngAggregateAddress { return v.AggregateAddresses }).(GetRipngAggregateAddressArrayOutput)
}

// Enable/disable generation of default route.
func (o LookupRipngResultOutput) DefaultInformationOriginate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRipngResult) string { return v.DefaultInformationOriginate }).(pulumi.StringOutput)
}

// Default metric.
func (o LookupRipngResultOutput) DefaultMetric() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRipngResult) int { return v.DefaultMetric }).(pulumi.IntOutput)
}

// Distance (1 - 255).
func (o LookupRipngResultOutput) Distances() GetRipngDistanceArrayOutput {
	return o.ApplyT(func(v LookupRipngResult) []GetRipngDistance { return v.Distances }).(GetRipngDistanceArrayOutput)
}

// Distribute list. The structure of `distributeList` block is documented below.
func (o LookupRipngResultOutput) DistributeLists() GetRipngDistributeListArrayOutput {
	return o.ApplyT(func(v LookupRipngResult) []GetRipngDistributeList { return v.DistributeLists }).(GetRipngDistributeListArrayOutput)
}

// Garbage timer.
func (o LookupRipngResultOutput) GarbageTimer() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRipngResult) int { return v.GarbageTimer }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRipngResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRipngResult) string { return v.Id }).(pulumi.StringOutput)
}

// Interface name.
func (o LookupRipngResultOutput) Interfaces() GetRipngInterfaceArrayOutput {
	return o.ApplyT(func(v LookupRipngResult) []GetRipngInterface { return v.Interfaces }).(GetRipngInterfaceArrayOutput)
}

// Maximum metric allowed to output(0 means 'not set').
func (o LookupRipngResultOutput) MaxOutMetric() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRipngResult) int { return v.MaxOutMetric }).(pulumi.IntOutput)
}

// neighbor The structure of `neighbor` block is documented below.
func (o LookupRipngResultOutput) Neighbors() GetRipngNeighborArrayOutput {
	return o.ApplyT(func(v LookupRipngResult) []GetRipngNeighbor { return v.Neighbors }).(GetRipngNeighborArrayOutput)
}

// Network. The structure of `network` block is documented below.
func (o LookupRipngResultOutput) Networks() GetRipngNetworkArrayOutput {
	return o.ApplyT(func(v LookupRipngResult) []GetRipngNetwork { return v.Networks }).(GetRipngNetworkArrayOutput)
}

// Offset list. The structure of `offsetList` block is documented below.
func (o LookupRipngResultOutput) OffsetLists() GetRipngOffsetListArrayOutput {
	return o.ApplyT(func(v LookupRipngResult) []GetRipngOffsetList { return v.OffsetLists }).(GetRipngOffsetListArrayOutput)
}

// Passive interface configuration. The structure of `passiveInterface` block is documented below.
func (o LookupRipngResultOutput) PassiveInterfaces() GetRipngPassiveInterfaceArrayOutput {
	return o.ApplyT(func(v LookupRipngResult) []GetRipngPassiveInterface { return v.PassiveInterfaces }).(GetRipngPassiveInterfaceArrayOutput)
}

// Redistribute configuration. The structure of `redistribute` block is documented below.
func (o LookupRipngResultOutput) Redistributes() GetRipngRedistributeArrayOutput {
	return o.ApplyT(func(v LookupRipngResult) []GetRipngRedistribute { return v.Redistributes }).(GetRipngRedistributeArrayOutput)
}

// Timeout timer.
func (o LookupRipngResultOutput) TimeoutTimer() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRipngResult) int { return v.TimeoutTimer }).(pulumi.IntOutput)
}

// Update timer.
func (o LookupRipngResultOutput) UpdateTimer() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRipngResult) int { return v.UpdateTimer }).(pulumi.IntOutput)
}

func (o LookupRipngResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRipngResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRipngResultOutput{})
}
