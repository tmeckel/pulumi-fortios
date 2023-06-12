// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sys

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure USB auto installation.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/sys"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := sys.NewAutoinstall(ctx, "trname", &sys.AutoinstallArgs{
//				AutoInstallConfig: pulumi.String("enable"),
//				AutoInstallImage:  pulumi.String("enable"),
//				DefaultConfigFile: pulumi.String("fgt_system.conf"),
//				DefaultImageFile:  pulumi.String("image.out"),
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
// # System AutoInstall can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:sys/autoinstall:Autoinstall labelname SystemAutoInstall
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:sys/autoinstall:Autoinstall labelname SystemAutoInstall
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type Autoinstall struct {
	pulumi.CustomResourceState

	// Enable/disable auto install the config in USB disk. Valid values: `enable`, `disable`.
	AutoInstallConfig pulumi.StringOutput `pulumi:"autoInstallConfig"`
	// Enable/disable auto install the image in USB disk. Valid values: `enable`, `disable`.
	AutoInstallImage pulumi.StringOutput `pulumi:"autoInstallImage"`
	// Default config file name in USB disk.
	DefaultConfigFile pulumi.StringOutput `pulumi:"defaultConfigFile"`
	// Default image file name in USB disk.
	DefaultImageFile pulumi.StringOutput `pulumi:"defaultImageFile"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewAutoinstall registers a new resource with the given unique name, arguments, and options.
func NewAutoinstall(ctx *pulumi.Context,
	name string, args *AutoinstallArgs, opts ...pulumi.ResourceOption) (*Autoinstall, error) {
	if args == nil {
		args = &AutoinstallArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource Autoinstall
	err := ctx.RegisterResource("fortios:sys/autoinstall:Autoinstall", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAutoinstall gets an existing Autoinstall resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAutoinstall(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AutoinstallState, opts ...pulumi.ResourceOption) (*Autoinstall, error) {
	var resource Autoinstall
	err := ctx.ReadResource("fortios:sys/autoinstall:Autoinstall", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Autoinstall resources.
type autoinstallState struct {
	// Enable/disable auto install the config in USB disk. Valid values: `enable`, `disable`.
	AutoInstallConfig *string `pulumi:"autoInstallConfig"`
	// Enable/disable auto install the image in USB disk. Valid values: `enable`, `disable`.
	AutoInstallImage *string `pulumi:"autoInstallImage"`
	// Default config file name in USB disk.
	DefaultConfigFile *string `pulumi:"defaultConfigFile"`
	// Default image file name in USB disk.
	DefaultImageFile *string `pulumi:"defaultImageFile"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type AutoinstallState struct {
	// Enable/disable auto install the config in USB disk. Valid values: `enable`, `disable`.
	AutoInstallConfig pulumi.StringPtrInput
	// Enable/disable auto install the image in USB disk. Valid values: `enable`, `disable`.
	AutoInstallImage pulumi.StringPtrInput
	// Default config file name in USB disk.
	DefaultConfigFile pulumi.StringPtrInput
	// Default image file name in USB disk.
	DefaultImageFile pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (AutoinstallState) ElementType() reflect.Type {
	return reflect.TypeOf((*autoinstallState)(nil)).Elem()
}

type autoinstallArgs struct {
	// Enable/disable auto install the config in USB disk. Valid values: `enable`, `disable`.
	AutoInstallConfig *string `pulumi:"autoInstallConfig"`
	// Enable/disable auto install the image in USB disk. Valid values: `enable`, `disable`.
	AutoInstallImage *string `pulumi:"autoInstallImage"`
	// Default config file name in USB disk.
	DefaultConfigFile *string `pulumi:"defaultConfigFile"`
	// Default image file name in USB disk.
	DefaultImageFile *string `pulumi:"defaultImageFile"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Autoinstall resource.
type AutoinstallArgs struct {
	// Enable/disable auto install the config in USB disk. Valid values: `enable`, `disable`.
	AutoInstallConfig pulumi.StringPtrInput
	// Enable/disable auto install the image in USB disk. Valid values: `enable`, `disable`.
	AutoInstallImage pulumi.StringPtrInput
	// Default config file name in USB disk.
	DefaultConfigFile pulumi.StringPtrInput
	// Default image file name in USB disk.
	DefaultImageFile pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (AutoinstallArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*autoinstallArgs)(nil)).Elem()
}

type AutoinstallInput interface {
	pulumi.Input

	ToAutoinstallOutput() AutoinstallOutput
	ToAutoinstallOutputWithContext(ctx context.Context) AutoinstallOutput
}

func (*Autoinstall) ElementType() reflect.Type {
	return reflect.TypeOf((**Autoinstall)(nil)).Elem()
}

func (i *Autoinstall) ToAutoinstallOutput() AutoinstallOutput {
	return i.ToAutoinstallOutputWithContext(context.Background())
}

func (i *Autoinstall) ToAutoinstallOutputWithContext(ctx context.Context) AutoinstallOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoinstallOutput)
}

// AutoinstallArrayInput is an input type that accepts AutoinstallArray and AutoinstallArrayOutput values.
// You can construct a concrete instance of `AutoinstallArrayInput` via:
//
//	AutoinstallArray{ AutoinstallArgs{...} }
type AutoinstallArrayInput interface {
	pulumi.Input

	ToAutoinstallArrayOutput() AutoinstallArrayOutput
	ToAutoinstallArrayOutputWithContext(context.Context) AutoinstallArrayOutput
}

type AutoinstallArray []AutoinstallInput

func (AutoinstallArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Autoinstall)(nil)).Elem()
}

func (i AutoinstallArray) ToAutoinstallArrayOutput() AutoinstallArrayOutput {
	return i.ToAutoinstallArrayOutputWithContext(context.Background())
}

func (i AutoinstallArray) ToAutoinstallArrayOutputWithContext(ctx context.Context) AutoinstallArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoinstallArrayOutput)
}

// AutoinstallMapInput is an input type that accepts AutoinstallMap and AutoinstallMapOutput values.
// You can construct a concrete instance of `AutoinstallMapInput` via:
//
//	AutoinstallMap{ "key": AutoinstallArgs{...} }
type AutoinstallMapInput interface {
	pulumi.Input

	ToAutoinstallMapOutput() AutoinstallMapOutput
	ToAutoinstallMapOutputWithContext(context.Context) AutoinstallMapOutput
}

type AutoinstallMap map[string]AutoinstallInput

func (AutoinstallMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Autoinstall)(nil)).Elem()
}

func (i AutoinstallMap) ToAutoinstallMapOutput() AutoinstallMapOutput {
	return i.ToAutoinstallMapOutputWithContext(context.Background())
}

func (i AutoinstallMap) ToAutoinstallMapOutputWithContext(ctx context.Context) AutoinstallMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoinstallMapOutput)
}

type AutoinstallOutput struct{ *pulumi.OutputState }

func (AutoinstallOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Autoinstall)(nil)).Elem()
}

func (o AutoinstallOutput) ToAutoinstallOutput() AutoinstallOutput {
	return o
}

func (o AutoinstallOutput) ToAutoinstallOutputWithContext(ctx context.Context) AutoinstallOutput {
	return o
}

// Enable/disable auto install the config in USB disk. Valid values: `enable`, `disable`.
func (o AutoinstallOutput) AutoInstallConfig() pulumi.StringOutput {
	return o.ApplyT(func(v *Autoinstall) pulumi.StringOutput { return v.AutoInstallConfig }).(pulumi.StringOutput)
}

// Enable/disable auto install the image in USB disk. Valid values: `enable`, `disable`.
func (o AutoinstallOutput) AutoInstallImage() pulumi.StringOutput {
	return o.ApplyT(func(v *Autoinstall) pulumi.StringOutput { return v.AutoInstallImage }).(pulumi.StringOutput)
}

// Default config file name in USB disk.
func (o AutoinstallOutput) DefaultConfigFile() pulumi.StringOutput {
	return o.ApplyT(func(v *Autoinstall) pulumi.StringOutput { return v.DefaultConfigFile }).(pulumi.StringOutput)
}

// Default image file name in USB disk.
func (o AutoinstallOutput) DefaultImageFile() pulumi.StringOutput {
	return o.ApplyT(func(v *Autoinstall) pulumi.StringOutput { return v.DefaultImageFile }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o AutoinstallOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Autoinstall) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type AutoinstallArrayOutput struct{ *pulumi.OutputState }

func (AutoinstallArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Autoinstall)(nil)).Elem()
}

func (o AutoinstallArrayOutput) ToAutoinstallArrayOutput() AutoinstallArrayOutput {
	return o
}

func (o AutoinstallArrayOutput) ToAutoinstallArrayOutputWithContext(ctx context.Context) AutoinstallArrayOutput {
	return o
}

func (o AutoinstallArrayOutput) Index(i pulumi.IntInput) AutoinstallOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Autoinstall {
		return vs[0].([]*Autoinstall)[vs[1].(int)]
	}).(AutoinstallOutput)
}

type AutoinstallMapOutput struct{ *pulumi.OutputState }

func (AutoinstallMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Autoinstall)(nil)).Elem()
}

func (o AutoinstallMapOutput) ToAutoinstallMapOutput() AutoinstallMapOutput {
	return o
}

func (o AutoinstallMapOutput) ToAutoinstallMapOutputWithContext(ctx context.Context) AutoinstallMapOutput {
	return o
}

func (o AutoinstallMapOutput) MapIndex(k pulumi.StringInput) AutoinstallOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Autoinstall {
		return vs[0].(map[string]*Autoinstall)[vs[1].(string)]
	}).(AutoinstallOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AutoinstallInput)(nil)).Elem(), &Autoinstall{})
	pulumi.RegisterInputType(reflect.TypeOf((*AutoinstallArrayInput)(nil)).Elem(), AutoinstallArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AutoinstallMapInput)(nil)).Elem(), AutoinstallMap{})
	pulumi.RegisterOutputType(AutoinstallOutput{})
	pulumi.RegisterOutputType(AutoinstallArrayOutput{})
	pulumi.RegisterOutputType(AutoinstallMapOutput{})
}