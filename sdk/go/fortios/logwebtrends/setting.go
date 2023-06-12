// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package logwebtrends

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Settings for WebTrends.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/logwebtrends"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := logwebtrends.NewSetting(ctx, "trname", &logwebtrends.SettingArgs{
//				Status: pulumi.String("disable"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// # LogWebtrends Setting can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:logwebtrends/setting:Setting labelname LogWebtrendsSetting
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:logwebtrends/setting:Setting labelname LogWebtrendsSetting
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type Setting struct {
	pulumi.CustomResourceState

	// Address of the remote WebTrends server.
	Server pulumi.StringOutput `pulumi:"server"`
	// Enable/disable logging to WebTrends. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSetting registers a new resource with the given unique name, arguments, and options.
func NewSetting(ctx *pulumi.Context,
	name string, args *SettingArgs, opts ...pulumi.ResourceOption) (*Setting, error) {
	if args == nil {
		args = &SettingArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource Setting
	err := ctx.RegisterResource("fortios:logwebtrends/setting:Setting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSetting gets an existing Setting resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SettingState, opts ...pulumi.ResourceOption) (*Setting, error) {
	var resource Setting
	err := ctx.ReadResource("fortios:logwebtrends/setting:Setting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Setting resources.
type settingState struct {
	// Address of the remote WebTrends server.
	Server *string `pulumi:"server"`
	// Enable/disable logging to WebTrends. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SettingState struct {
	// Address of the remote WebTrends server.
	Server pulumi.StringPtrInput
	// Enable/disable logging to WebTrends. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*settingState)(nil)).Elem()
}

type settingArgs struct {
	// Address of the remote WebTrends server.
	Server *string `pulumi:"server"`
	// Enable/disable logging to WebTrends. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Setting resource.
type SettingArgs struct {
	// Address of the remote WebTrends server.
	Server pulumi.StringPtrInput
	// Enable/disable logging to WebTrends. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*settingArgs)(nil)).Elem()
}

type SettingInput interface {
	pulumi.Input

	ToSettingOutput() SettingOutput
	ToSettingOutputWithContext(ctx context.Context) SettingOutput
}

func (*Setting) ElementType() reflect.Type {
	return reflect.TypeOf((**Setting)(nil)).Elem()
}

func (i *Setting) ToSettingOutput() SettingOutput {
	return i.ToSettingOutputWithContext(context.Background())
}

func (i *Setting) ToSettingOutputWithContext(ctx context.Context) SettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SettingOutput)
}

// SettingArrayInput is an input type that accepts SettingArray and SettingArrayOutput values.
// You can construct a concrete instance of `SettingArrayInput` via:
//
//	SettingArray{ SettingArgs{...} }
type SettingArrayInput interface {
	pulumi.Input

	ToSettingArrayOutput() SettingArrayOutput
	ToSettingArrayOutputWithContext(context.Context) SettingArrayOutput
}

type SettingArray []SettingInput

func (SettingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Setting)(nil)).Elem()
}

func (i SettingArray) ToSettingArrayOutput() SettingArrayOutput {
	return i.ToSettingArrayOutputWithContext(context.Background())
}

func (i SettingArray) ToSettingArrayOutputWithContext(ctx context.Context) SettingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SettingArrayOutput)
}

// SettingMapInput is an input type that accepts SettingMap and SettingMapOutput values.
// You can construct a concrete instance of `SettingMapInput` via:
//
//	SettingMap{ "key": SettingArgs{...} }
type SettingMapInput interface {
	pulumi.Input

	ToSettingMapOutput() SettingMapOutput
	ToSettingMapOutputWithContext(context.Context) SettingMapOutput
}

type SettingMap map[string]SettingInput

func (SettingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Setting)(nil)).Elem()
}

func (i SettingMap) ToSettingMapOutput() SettingMapOutput {
	return i.ToSettingMapOutputWithContext(context.Background())
}

func (i SettingMap) ToSettingMapOutputWithContext(ctx context.Context) SettingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SettingMapOutput)
}

type SettingOutput struct{ *pulumi.OutputState }

func (SettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Setting)(nil)).Elem()
}

func (o SettingOutput) ToSettingOutput() SettingOutput {
	return o
}

func (o SettingOutput) ToSettingOutputWithContext(ctx context.Context) SettingOutput {
	return o
}

// Address of the remote WebTrends server.
func (o SettingOutput) Server() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.Server }).(pulumi.StringOutput)
}

// Enable/disable logging to WebTrends. Valid values: `enable`, `disable`.
func (o SettingOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SettingOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SettingArrayOutput struct{ *pulumi.OutputState }

func (SettingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Setting)(nil)).Elem()
}

func (o SettingArrayOutput) ToSettingArrayOutput() SettingArrayOutput {
	return o
}

func (o SettingArrayOutput) ToSettingArrayOutputWithContext(ctx context.Context) SettingArrayOutput {
	return o
}

func (o SettingArrayOutput) Index(i pulumi.IntInput) SettingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Setting {
		return vs[0].([]*Setting)[vs[1].(int)]
	}).(SettingOutput)
}

type SettingMapOutput struct{ *pulumi.OutputState }

func (SettingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Setting)(nil)).Elem()
}

func (o SettingMapOutput) ToSettingMapOutput() SettingMapOutput {
	return o
}

func (o SettingMapOutput) ToSettingMapOutputWithContext(ctx context.Context) SettingMapOutput {
	return o
}

func (o SettingMapOutput) MapIndex(k pulumi.StringInput) SettingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Setting {
		return vs[0].(map[string]*Setting)[vs[1].(string)]
	}).(SettingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SettingInput)(nil)).Elem(), &Setting{})
	pulumi.RegisterInputType(reflect.TypeOf((*SettingArrayInput)(nil)).Elem(), SettingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SettingMapInput)(nil)).Elem(), SettingMap{})
	pulumi.RegisterOutputType(SettingOutput{})
	pulumi.RegisterOutputType(SettingArrayOutput{})
	pulumi.RegisterOutputType(SettingMapOutput{})
}
