// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package wirelesscontroller

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure FortiAP regions (for floor plans and maps).
//
// ## Import
//
// # WirelessController Region can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:wirelesscontroller/region:Region labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:wirelesscontroller/region:Region labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type Region struct {
	pulumi.CustomResourceState

	// Comments.
	Comments pulumi.StringOutput `pulumi:"comments"`
	// Region image grayscale. Valid values: `enable`, `disable`.
	Grayscale pulumi.StringOutput `pulumi:"grayscale"`
	// FortiAP region image type (png|jpeg|gif). Valid values: `png`, `jpeg`, `gif`.
	ImageType pulumi.StringOutput `pulumi:"imageType"`
	// FortiAP region name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Region image opacity (0 - 100).
	Opacity pulumi.IntOutput `pulumi:"opacity"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewRegion registers a new resource with the given unique name, arguments, and options.
func NewRegion(ctx *pulumi.Context,
	name string, args *RegionArgs, opts ...pulumi.ResourceOption) (*Region, error) {
	if args == nil {
		args = &RegionArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource Region
	err := ctx.RegisterResource("fortios:wirelesscontroller/region:Region", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegion gets an existing Region resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegionState, opts ...pulumi.ResourceOption) (*Region, error) {
	var resource Region
	err := ctx.ReadResource("fortios:wirelesscontroller/region:Region", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Region resources.
type regionState struct {
	// Comments.
	Comments *string `pulumi:"comments"`
	// Region image grayscale. Valid values: `enable`, `disable`.
	Grayscale *string `pulumi:"grayscale"`
	// FortiAP region image type (png|jpeg|gif). Valid values: `png`, `jpeg`, `gif`.
	ImageType *string `pulumi:"imageType"`
	// FortiAP region name.
	Name *string `pulumi:"name"`
	// Region image opacity (0 - 100).
	Opacity *int `pulumi:"opacity"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type RegionState struct {
	// Comments.
	Comments pulumi.StringPtrInput
	// Region image grayscale. Valid values: `enable`, `disable`.
	Grayscale pulumi.StringPtrInput
	// FortiAP region image type (png|jpeg|gif). Valid values: `png`, `jpeg`, `gif`.
	ImageType pulumi.StringPtrInput
	// FortiAP region name.
	Name pulumi.StringPtrInput
	// Region image opacity (0 - 100).
	Opacity pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (RegionState) ElementType() reflect.Type {
	return reflect.TypeOf((*regionState)(nil)).Elem()
}

type regionArgs struct {
	// Comments.
	Comments *string `pulumi:"comments"`
	// Region image grayscale. Valid values: `enable`, `disable`.
	Grayscale *string `pulumi:"grayscale"`
	// FortiAP region image type (png|jpeg|gif). Valid values: `png`, `jpeg`, `gif`.
	ImageType *string `pulumi:"imageType"`
	// FortiAP region name.
	Name *string `pulumi:"name"`
	// Region image opacity (0 - 100).
	Opacity *int `pulumi:"opacity"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Region resource.
type RegionArgs struct {
	// Comments.
	Comments pulumi.StringPtrInput
	// Region image grayscale. Valid values: `enable`, `disable`.
	Grayscale pulumi.StringPtrInput
	// FortiAP region image type (png|jpeg|gif). Valid values: `png`, `jpeg`, `gif`.
	ImageType pulumi.StringPtrInput
	// FortiAP region name.
	Name pulumi.StringPtrInput
	// Region image opacity (0 - 100).
	Opacity pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (RegionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*regionArgs)(nil)).Elem()
}

type RegionInput interface {
	pulumi.Input

	ToRegionOutput() RegionOutput
	ToRegionOutputWithContext(ctx context.Context) RegionOutput
}

func (*Region) ElementType() reflect.Type {
	return reflect.TypeOf((**Region)(nil)).Elem()
}

func (i *Region) ToRegionOutput() RegionOutput {
	return i.ToRegionOutputWithContext(context.Background())
}

func (i *Region) ToRegionOutputWithContext(ctx context.Context) RegionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionOutput)
}

// RegionArrayInput is an input type that accepts RegionArray and RegionArrayOutput values.
// You can construct a concrete instance of `RegionArrayInput` via:
//
//	RegionArray{ RegionArgs{...} }
type RegionArrayInput interface {
	pulumi.Input

	ToRegionArrayOutput() RegionArrayOutput
	ToRegionArrayOutputWithContext(context.Context) RegionArrayOutput
}

type RegionArray []RegionInput

func (RegionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Region)(nil)).Elem()
}

func (i RegionArray) ToRegionArrayOutput() RegionArrayOutput {
	return i.ToRegionArrayOutputWithContext(context.Background())
}

func (i RegionArray) ToRegionArrayOutputWithContext(ctx context.Context) RegionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionArrayOutput)
}

// RegionMapInput is an input type that accepts RegionMap and RegionMapOutput values.
// You can construct a concrete instance of `RegionMapInput` via:
//
//	RegionMap{ "key": RegionArgs{...} }
type RegionMapInput interface {
	pulumi.Input

	ToRegionMapOutput() RegionMapOutput
	ToRegionMapOutputWithContext(context.Context) RegionMapOutput
}

type RegionMap map[string]RegionInput

func (RegionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Region)(nil)).Elem()
}

func (i RegionMap) ToRegionMapOutput() RegionMapOutput {
	return i.ToRegionMapOutputWithContext(context.Background())
}

func (i RegionMap) ToRegionMapOutputWithContext(ctx context.Context) RegionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionMapOutput)
}

type RegionOutput struct{ *pulumi.OutputState }

func (RegionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Region)(nil)).Elem()
}

func (o RegionOutput) ToRegionOutput() RegionOutput {
	return o
}

func (o RegionOutput) ToRegionOutputWithContext(ctx context.Context) RegionOutput {
	return o
}

// Comments.
func (o RegionOutput) Comments() pulumi.StringOutput {
	return o.ApplyT(func(v *Region) pulumi.StringOutput { return v.Comments }).(pulumi.StringOutput)
}

// Region image grayscale. Valid values: `enable`, `disable`.
func (o RegionOutput) Grayscale() pulumi.StringOutput {
	return o.ApplyT(func(v *Region) pulumi.StringOutput { return v.Grayscale }).(pulumi.StringOutput)
}

// FortiAP region image type (png|jpeg|gif). Valid values: `png`, `jpeg`, `gif`.
func (o RegionOutput) ImageType() pulumi.StringOutput {
	return o.ApplyT(func(v *Region) pulumi.StringOutput { return v.ImageType }).(pulumi.StringOutput)
}

// FortiAP region name.
func (o RegionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Region) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Region image opacity (0 - 100).
func (o RegionOutput) Opacity() pulumi.IntOutput {
	return o.ApplyT(func(v *Region) pulumi.IntOutput { return v.Opacity }).(pulumi.IntOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o RegionOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Region) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type RegionArrayOutput struct{ *pulumi.OutputState }

func (RegionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Region)(nil)).Elem()
}

func (o RegionArrayOutput) ToRegionArrayOutput() RegionArrayOutput {
	return o
}

func (o RegionArrayOutput) ToRegionArrayOutputWithContext(ctx context.Context) RegionArrayOutput {
	return o
}

func (o RegionArrayOutput) Index(i pulumi.IntInput) RegionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Region {
		return vs[0].([]*Region)[vs[1].(int)]
	}).(RegionOutput)
}

type RegionMapOutput struct{ *pulumi.OutputState }

func (RegionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Region)(nil)).Elem()
}

func (o RegionMapOutput) ToRegionMapOutput() RegionMapOutput {
	return o
}

func (o RegionMapOutput) ToRegionMapOutputWithContext(ctx context.Context) RegionMapOutput {
	return o
}

func (o RegionMapOutput) MapIndex(k pulumi.StringInput) RegionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Region {
		return vs[0].(map[string]*Region)[vs[1].(string)]
	}).(RegionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RegionInput)(nil)).Elem(), &Region{})
	pulumi.RegisterInputType(reflect.TypeOf((*RegionArrayInput)(nil)).Elem(), RegionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RegionMapInput)(nil)).Elem(), RegionMap{})
	pulumi.RegisterOutputType(RegionOutput{})
	pulumi.RegisterOutputType(RegionArrayOutput{})
	pulumi.RegisterOutputType(RegionMapOutput{})
}