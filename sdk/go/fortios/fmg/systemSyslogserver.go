// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fmg

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource supports Create/Delete system syslog server for FortiManager.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/fmg"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := fmg.NewSystemSyslogserver(ctx, "test1", &fmg.SystemSyslogserverArgs{
//				Ip:   pulumi.String("1.1.1.1"),
//				Port: pulumi.Int(99),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type SystemSyslogserver struct {
	pulumi.CustomResourceState

	// Ipaddress of the syslog server.
	Ip pulumi.StringPtrOutput `pulumi:"ip"`
	// Syslog server name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Port of the syslog server.
	Port pulumi.IntPtrOutput `pulumi:"port"`
}

// NewSystemSyslogserver registers a new resource with the given unique name, arguments, and options.
func NewSystemSyslogserver(ctx *pulumi.Context,
	name string, args *SystemSyslogserverArgs, opts ...pulumi.ResourceOption) (*SystemSyslogserver, error) {
	if args == nil {
		args = &SystemSyslogserverArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SystemSyslogserver
	err := ctx.RegisterResource("fortios:fmg/systemSyslogserver:SystemSyslogserver", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemSyslogserver gets an existing SystemSyslogserver resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemSyslogserver(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemSyslogserverState, opts ...pulumi.ResourceOption) (*SystemSyslogserver, error) {
	var resource SystemSyslogserver
	err := ctx.ReadResource("fortios:fmg/systemSyslogserver:SystemSyslogserver", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemSyslogserver resources.
type systemSyslogserverState struct {
	// Ipaddress of the syslog server.
	Ip *string `pulumi:"ip"`
	// Syslog server name.
	Name *string `pulumi:"name"`
	// Port of the syslog server.
	Port *int `pulumi:"port"`
}

type SystemSyslogserverState struct {
	// Ipaddress of the syslog server.
	Ip pulumi.StringPtrInput
	// Syslog server name.
	Name pulumi.StringPtrInput
	// Port of the syslog server.
	Port pulumi.IntPtrInput
}

func (SystemSyslogserverState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemSyslogserverState)(nil)).Elem()
}

type systemSyslogserverArgs struct {
	// Ipaddress of the syslog server.
	Ip *string `pulumi:"ip"`
	// Syslog server name.
	Name *string `pulumi:"name"`
	// Port of the syslog server.
	Port *int `pulumi:"port"`
}

// The set of arguments for constructing a SystemSyslogserver resource.
type SystemSyslogserverArgs struct {
	// Ipaddress of the syslog server.
	Ip pulumi.StringPtrInput
	// Syslog server name.
	Name pulumi.StringPtrInput
	// Port of the syslog server.
	Port pulumi.IntPtrInput
}

func (SystemSyslogserverArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemSyslogserverArgs)(nil)).Elem()
}

type SystemSyslogserverInput interface {
	pulumi.Input

	ToSystemSyslogserverOutput() SystemSyslogserverOutput
	ToSystemSyslogserverOutputWithContext(ctx context.Context) SystemSyslogserverOutput
}

func (*SystemSyslogserver) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemSyslogserver)(nil)).Elem()
}

func (i *SystemSyslogserver) ToSystemSyslogserverOutput() SystemSyslogserverOutput {
	return i.ToSystemSyslogserverOutputWithContext(context.Background())
}

func (i *SystemSyslogserver) ToSystemSyslogserverOutputWithContext(ctx context.Context) SystemSyslogserverOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemSyslogserverOutput)
}

// SystemSyslogserverArrayInput is an input type that accepts SystemSyslogserverArray and SystemSyslogserverArrayOutput values.
// You can construct a concrete instance of `SystemSyslogserverArrayInput` via:
//
//	SystemSyslogserverArray{ SystemSyslogserverArgs{...} }
type SystemSyslogserverArrayInput interface {
	pulumi.Input

	ToSystemSyslogserverArrayOutput() SystemSyslogserverArrayOutput
	ToSystemSyslogserverArrayOutputWithContext(context.Context) SystemSyslogserverArrayOutput
}

type SystemSyslogserverArray []SystemSyslogserverInput

func (SystemSyslogserverArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemSyslogserver)(nil)).Elem()
}

func (i SystemSyslogserverArray) ToSystemSyslogserverArrayOutput() SystemSyslogserverArrayOutput {
	return i.ToSystemSyslogserverArrayOutputWithContext(context.Background())
}

func (i SystemSyslogserverArray) ToSystemSyslogserverArrayOutputWithContext(ctx context.Context) SystemSyslogserverArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemSyslogserverArrayOutput)
}

// SystemSyslogserverMapInput is an input type that accepts SystemSyslogserverMap and SystemSyslogserverMapOutput values.
// You can construct a concrete instance of `SystemSyslogserverMapInput` via:
//
//	SystemSyslogserverMap{ "key": SystemSyslogserverArgs{...} }
type SystemSyslogserverMapInput interface {
	pulumi.Input

	ToSystemSyslogserverMapOutput() SystemSyslogserverMapOutput
	ToSystemSyslogserverMapOutputWithContext(context.Context) SystemSyslogserverMapOutput
}

type SystemSyslogserverMap map[string]SystemSyslogserverInput

func (SystemSyslogserverMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemSyslogserver)(nil)).Elem()
}

func (i SystemSyslogserverMap) ToSystemSyslogserverMapOutput() SystemSyslogserverMapOutput {
	return i.ToSystemSyslogserverMapOutputWithContext(context.Background())
}

func (i SystemSyslogserverMap) ToSystemSyslogserverMapOutputWithContext(ctx context.Context) SystemSyslogserverMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemSyslogserverMapOutput)
}

type SystemSyslogserverOutput struct{ *pulumi.OutputState }

func (SystemSyslogserverOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemSyslogserver)(nil)).Elem()
}

func (o SystemSyslogserverOutput) ToSystemSyslogserverOutput() SystemSyslogserverOutput {
	return o
}

func (o SystemSyslogserverOutput) ToSystemSyslogserverOutputWithContext(ctx context.Context) SystemSyslogserverOutput {
	return o
}

// Ipaddress of the syslog server.
func (o SystemSyslogserverOutput) Ip() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemSyslogserver) pulumi.StringPtrOutput { return v.Ip }).(pulumi.StringPtrOutput)
}

// Syslog server name.
func (o SystemSyslogserverOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemSyslogserver) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Port of the syslog server.
func (o SystemSyslogserverOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SystemSyslogserver) pulumi.IntPtrOutput { return v.Port }).(pulumi.IntPtrOutput)
}

type SystemSyslogserverArrayOutput struct{ *pulumi.OutputState }

func (SystemSyslogserverArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemSyslogserver)(nil)).Elem()
}

func (o SystemSyslogserverArrayOutput) ToSystemSyslogserverArrayOutput() SystemSyslogserverArrayOutput {
	return o
}

func (o SystemSyslogserverArrayOutput) ToSystemSyslogserverArrayOutputWithContext(ctx context.Context) SystemSyslogserverArrayOutput {
	return o
}

func (o SystemSyslogserverArrayOutput) Index(i pulumi.IntInput) SystemSyslogserverOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemSyslogserver {
		return vs[0].([]*SystemSyslogserver)[vs[1].(int)]
	}).(SystemSyslogserverOutput)
}

type SystemSyslogserverMapOutput struct{ *pulumi.OutputState }

func (SystemSyslogserverMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemSyslogserver)(nil)).Elem()
}

func (o SystemSyslogserverMapOutput) ToSystemSyslogserverMapOutput() SystemSyslogserverMapOutput {
	return o
}

func (o SystemSyslogserverMapOutput) ToSystemSyslogserverMapOutputWithContext(ctx context.Context) SystemSyslogserverMapOutput {
	return o
}

func (o SystemSyslogserverMapOutput) MapIndex(k pulumi.StringInput) SystemSyslogserverOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemSyslogserver {
		return vs[0].(map[string]*SystemSyslogserver)[vs[1].(string)]
	}).(SystemSyslogserverOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemSyslogserverInput)(nil)).Elem(), &SystemSyslogserver{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemSyslogserverArrayInput)(nil)).Elem(), SystemSyslogserverArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemSyslogserverMapInput)(nil)).Elem(), SystemSyslogserverMap{})
	pulumi.RegisterOutputType(SystemSyslogserverOutput{})
	pulumi.RegisterOutputType(SystemSyslogserverArrayOutput{})
	pulumi.RegisterOutputType(SystemSyslogserverMapOutput{})
}