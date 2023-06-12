// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package logdisk

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type FilterFreeStyle struct {
	// Log category.
	Category *string `pulumi:"category"`
	// Free style filter string.
	Filter *string `pulumi:"filter"`
	// Include/exclude logs that match the filter. Valid values: `include`, `exclude`.
	FilterType *string `pulumi:"filterType"`
	// Entry ID.
	Id *int `pulumi:"id"`
}

// FilterFreeStyleInput is an input type that accepts FilterFreeStyleArgs and FilterFreeStyleOutput values.
// You can construct a concrete instance of `FilterFreeStyleInput` via:
//
//	FilterFreeStyleArgs{...}
type FilterFreeStyleInput interface {
	pulumi.Input

	ToFilterFreeStyleOutput() FilterFreeStyleOutput
	ToFilterFreeStyleOutputWithContext(context.Context) FilterFreeStyleOutput
}

type FilterFreeStyleArgs struct {
	// Log category.
	Category pulumi.StringPtrInput `pulumi:"category"`
	// Free style filter string.
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Include/exclude logs that match the filter. Valid values: `include`, `exclude`.
	FilterType pulumi.StringPtrInput `pulumi:"filterType"`
	// Entry ID.
	Id pulumi.IntPtrInput `pulumi:"id"`
}

func (FilterFreeStyleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FilterFreeStyle)(nil)).Elem()
}

func (i FilterFreeStyleArgs) ToFilterFreeStyleOutput() FilterFreeStyleOutput {
	return i.ToFilterFreeStyleOutputWithContext(context.Background())
}

func (i FilterFreeStyleArgs) ToFilterFreeStyleOutputWithContext(ctx context.Context) FilterFreeStyleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FilterFreeStyleOutput)
}

// FilterFreeStyleArrayInput is an input type that accepts FilterFreeStyleArray and FilterFreeStyleArrayOutput values.
// You can construct a concrete instance of `FilterFreeStyleArrayInput` via:
//
//	FilterFreeStyleArray{ FilterFreeStyleArgs{...} }
type FilterFreeStyleArrayInput interface {
	pulumi.Input

	ToFilterFreeStyleArrayOutput() FilterFreeStyleArrayOutput
	ToFilterFreeStyleArrayOutputWithContext(context.Context) FilterFreeStyleArrayOutput
}

type FilterFreeStyleArray []FilterFreeStyleInput

func (FilterFreeStyleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FilterFreeStyle)(nil)).Elem()
}

func (i FilterFreeStyleArray) ToFilterFreeStyleArrayOutput() FilterFreeStyleArrayOutput {
	return i.ToFilterFreeStyleArrayOutputWithContext(context.Background())
}

func (i FilterFreeStyleArray) ToFilterFreeStyleArrayOutputWithContext(ctx context.Context) FilterFreeStyleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FilterFreeStyleArrayOutput)
}

type FilterFreeStyleOutput struct{ *pulumi.OutputState }

func (FilterFreeStyleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FilterFreeStyle)(nil)).Elem()
}

func (o FilterFreeStyleOutput) ToFilterFreeStyleOutput() FilterFreeStyleOutput {
	return o
}

func (o FilterFreeStyleOutput) ToFilterFreeStyleOutputWithContext(ctx context.Context) FilterFreeStyleOutput {
	return o
}

// Log category.
func (o FilterFreeStyleOutput) Category() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FilterFreeStyle) *string { return v.Category }).(pulumi.StringPtrOutput)
}

// Free style filter string.
func (o FilterFreeStyleOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FilterFreeStyle) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// Include/exclude logs that match the filter. Valid values: `include`, `exclude`.
func (o FilterFreeStyleOutput) FilterType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FilterFreeStyle) *string { return v.FilterType }).(pulumi.StringPtrOutput)
}

// Entry ID.
func (o FilterFreeStyleOutput) Id() pulumi.IntPtrOutput {
	return o.ApplyT(func(v FilterFreeStyle) *int { return v.Id }).(pulumi.IntPtrOutput)
}

type FilterFreeStyleArrayOutput struct{ *pulumi.OutputState }

func (FilterFreeStyleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FilterFreeStyle)(nil)).Elem()
}

func (o FilterFreeStyleArrayOutput) ToFilterFreeStyleArrayOutput() FilterFreeStyleArrayOutput {
	return o
}

func (o FilterFreeStyleArrayOutput) ToFilterFreeStyleArrayOutputWithContext(ctx context.Context) FilterFreeStyleArrayOutput {
	return o
}

func (o FilterFreeStyleArrayOutput) Index(i pulumi.IntInput) FilterFreeStyleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FilterFreeStyle {
		return vs[0].([]FilterFreeStyle)[vs[1].(int)]
	}).(FilterFreeStyleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FilterFreeStyleInput)(nil)).Elem(), FilterFreeStyleArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*FilterFreeStyleArrayInput)(nil)).Elem(), FilterFreeStyleArray{})
	pulumi.RegisterOutputType(FilterFreeStyleOutput{})
	pulumi.RegisterOutputType(FilterFreeStyleArrayOutput{})
}
