// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fmg

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource supports modifying system ntp setting for FortiManager.
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
//			_, err := fmg.NewSystemNtp(ctx, "test1", &fmg.SystemNtpArgs{
//				Server:       pulumi.String("ntp1.fortinet.com"),
//				Status:       pulumi.String("enable"),
//				SyncInterval: pulumi.Int(30),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type SystemNtp struct {
	pulumi.CustomResourceState

	// IP address/hostname of NTP Server.
	Server pulumi.StringOutput `pulumi:"server"`
	// Enable/disable NTP.
	Status pulumi.StringPtrOutput `pulumi:"status"`
	// NTP sync interval (minute).
	SyncInterval pulumi.IntPtrOutput `pulumi:"syncInterval"`
}

// NewSystemNtp registers a new resource with the given unique name, arguments, and options.
func NewSystemNtp(ctx *pulumi.Context,
	name string, args *SystemNtpArgs, opts ...pulumi.ResourceOption) (*SystemNtp, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Server == nil {
		return nil, errors.New("invalid value for required argument 'Server'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource SystemNtp
	err := ctx.RegisterResource("fortios:fmg/systemNtp:SystemNtp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemNtp gets an existing SystemNtp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemNtp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemNtpState, opts ...pulumi.ResourceOption) (*SystemNtp, error) {
	var resource SystemNtp
	err := ctx.ReadResource("fortios:fmg/systemNtp:SystemNtp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemNtp resources.
type systemNtpState struct {
	// IP address/hostname of NTP Server.
	Server *string `pulumi:"server"`
	// Enable/disable NTP.
	Status *string `pulumi:"status"`
	// NTP sync interval (minute).
	SyncInterval *int `pulumi:"syncInterval"`
}

type SystemNtpState struct {
	// IP address/hostname of NTP Server.
	Server pulumi.StringPtrInput
	// Enable/disable NTP.
	Status pulumi.StringPtrInput
	// NTP sync interval (minute).
	SyncInterval pulumi.IntPtrInput
}

func (SystemNtpState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemNtpState)(nil)).Elem()
}

type systemNtpArgs struct {
	// IP address/hostname of NTP Server.
	Server string `pulumi:"server"`
	// Enable/disable NTP.
	Status *string `pulumi:"status"`
	// NTP sync interval (minute).
	SyncInterval *int `pulumi:"syncInterval"`
}

// The set of arguments for constructing a SystemNtp resource.
type SystemNtpArgs struct {
	// IP address/hostname of NTP Server.
	Server pulumi.StringInput
	// Enable/disable NTP.
	Status pulumi.StringPtrInput
	// NTP sync interval (minute).
	SyncInterval pulumi.IntPtrInput
}

func (SystemNtpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemNtpArgs)(nil)).Elem()
}

type SystemNtpInput interface {
	pulumi.Input

	ToSystemNtpOutput() SystemNtpOutput
	ToSystemNtpOutputWithContext(ctx context.Context) SystemNtpOutput
}

func (*SystemNtp) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemNtp)(nil)).Elem()
}

func (i *SystemNtp) ToSystemNtpOutput() SystemNtpOutput {
	return i.ToSystemNtpOutputWithContext(context.Background())
}

func (i *SystemNtp) ToSystemNtpOutputWithContext(ctx context.Context) SystemNtpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemNtpOutput)
}

// SystemNtpArrayInput is an input type that accepts SystemNtpArray and SystemNtpArrayOutput values.
// You can construct a concrete instance of `SystemNtpArrayInput` via:
//
//	SystemNtpArray{ SystemNtpArgs{...} }
type SystemNtpArrayInput interface {
	pulumi.Input

	ToSystemNtpArrayOutput() SystemNtpArrayOutput
	ToSystemNtpArrayOutputWithContext(context.Context) SystemNtpArrayOutput
}

type SystemNtpArray []SystemNtpInput

func (SystemNtpArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemNtp)(nil)).Elem()
}

func (i SystemNtpArray) ToSystemNtpArrayOutput() SystemNtpArrayOutput {
	return i.ToSystemNtpArrayOutputWithContext(context.Background())
}

func (i SystemNtpArray) ToSystemNtpArrayOutputWithContext(ctx context.Context) SystemNtpArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemNtpArrayOutput)
}

// SystemNtpMapInput is an input type that accepts SystemNtpMap and SystemNtpMapOutput values.
// You can construct a concrete instance of `SystemNtpMapInput` via:
//
//	SystemNtpMap{ "key": SystemNtpArgs{...} }
type SystemNtpMapInput interface {
	pulumi.Input

	ToSystemNtpMapOutput() SystemNtpMapOutput
	ToSystemNtpMapOutputWithContext(context.Context) SystemNtpMapOutput
}

type SystemNtpMap map[string]SystemNtpInput

func (SystemNtpMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemNtp)(nil)).Elem()
}

func (i SystemNtpMap) ToSystemNtpMapOutput() SystemNtpMapOutput {
	return i.ToSystemNtpMapOutputWithContext(context.Background())
}

func (i SystemNtpMap) ToSystemNtpMapOutputWithContext(ctx context.Context) SystemNtpMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemNtpMapOutput)
}

type SystemNtpOutput struct{ *pulumi.OutputState }

func (SystemNtpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemNtp)(nil)).Elem()
}

func (o SystemNtpOutput) ToSystemNtpOutput() SystemNtpOutput {
	return o
}

func (o SystemNtpOutput) ToSystemNtpOutputWithContext(ctx context.Context) SystemNtpOutput {
	return o
}

// IP address/hostname of NTP Server.
func (o SystemNtpOutput) Server() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemNtp) pulumi.StringOutput { return v.Server }).(pulumi.StringOutput)
}

// Enable/disable NTP.
func (o SystemNtpOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemNtp) pulumi.StringPtrOutput { return v.Status }).(pulumi.StringPtrOutput)
}

// NTP sync interval (minute).
func (o SystemNtpOutput) SyncInterval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SystemNtp) pulumi.IntPtrOutput { return v.SyncInterval }).(pulumi.IntPtrOutput)
}

type SystemNtpArrayOutput struct{ *pulumi.OutputState }

func (SystemNtpArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemNtp)(nil)).Elem()
}

func (o SystemNtpArrayOutput) ToSystemNtpArrayOutput() SystemNtpArrayOutput {
	return o
}

func (o SystemNtpArrayOutput) ToSystemNtpArrayOutputWithContext(ctx context.Context) SystemNtpArrayOutput {
	return o
}

func (o SystemNtpArrayOutput) Index(i pulumi.IntInput) SystemNtpOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemNtp {
		return vs[0].([]*SystemNtp)[vs[1].(int)]
	}).(SystemNtpOutput)
}

type SystemNtpMapOutput struct{ *pulumi.OutputState }

func (SystemNtpMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemNtp)(nil)).Elem()
}

func (o SystemNtpMapOutput) ToSystemNtpMapOutput() SystemNtpMapOutput {
	return o
}

func (o SystemNtpMapOutput) ToSystemNtpMapOutputWithContext(ctx context.Context) SystemNtpMapOutput {
	return o
}

func (o SystemNtpMapOutput) MapIndex(k pulumi.StringInput) SystemNtpOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemNtp {
		return vs[0].(map[string]*SystemNtp)[vs[1].(string)]
	}).(SystemNtpOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemNtpInput)(nil)).Elem(), &SystemNtp{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemNtpArrayInput)(nil)).Elem(), SystemNtpArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemNtpMapInput)(nil)).Elem(), SystemNtpMap{})
	pulumi.RegisterOutputType(SystemNtpOutput{})
	pulumi.RegisterOutputType(SystemNtpArrayOutput{})
	pulumi.RegisterOutputType(SystemNtpMapOutput{})
}
