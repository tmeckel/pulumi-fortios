// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package systemreplacemsg

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Replacement messages.
//
// ## Import
//
// # SystemReplacemsg NacQuar can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:systemreplacemsg/nacquar:Nacquar labelname {{msg_type}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:systemreplacemsg/nacquar:Nacquar labelname {{msg_type}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type Nacquar struct {
	pulumi.CustomResourceState

	// Message string.
	Buffer pulumi.StringPtrOutput `pulumi:"buffer"`
	// Format flag.
	Format pulumi.StringOutput `pulumi:"format"`
	// Header flag. Valid values: `none`, `http`, `8bit`.
	Header pulumi.StringOutput `pulumi:"header"`
	// Message type.
	MsgType pulumi.StringOutput `pulumi:"msgType"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewNacquar registers a new resource with the given unique name, arguments, and options.
func NewNacquar(ctx *pulumi.Context,
	name string, args *NacquarArgs, opts ...pulumi.ResourceOption) (*Nacquar, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MsgType == nil {
		return nil, errors.New("invalid value for required argument 'MsgType'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Nacquar
	err := ctx.RegisterResource("fortios:systemreplacemsg/nacquar:Nacquar", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNacquar gets an existing Nacquar resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNacquar(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NacquarState, opts ...pulumi.ResourceOption) (*Nacquar, error) {
	var resource Nacquar
	err := ctx.ReadResource("fortios:systemreplacemsg/nacquar:Nacquar", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Nacquar resources.
type nacquarState struct {
	// Message string.
	Buffer *string `pulumi:"buffer"`
	// Format flag.
	Format *string `pulumi:"format"`
	// Header flag. Valid values: `none`, `http`, `8bit`.
	Header *string `pulumi:"header"`
	// Message type.
	MsgType *string `pulumi:"msgType"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type NacquarState struct {
	// Message string.
	Buffer pulumi.StringPtrInput
	// Format flag.
	Format pulumi.StringPtrInput
	// Header flag. Valid values: `none`, `http`, `8bit`.
	Header pulumi.StringPtrInput
	// Message type.
	MsgType pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (NacquarState) ElementType() reflect.Type {
	return reflect.TypeOf((*nacquarState)(nil)).Elem()
}

type nacquarArgs struct {
	// Message string.
	Buffer *string `pulumi:"buffer"`
	// Format flag.
	Format *string `pulumi:"format"`
	// Header flag. Valid values: `none`, `http`, `8bit`.
	Header *string `pulumi:"header"`
	// Message type.
	MsgType string `pulumi:"msgType"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Nacquar resource.
type NacquarArgs struct {
	// Message string.
	Buffer pulumi.StringPtrInput
	// Format flag.
	Format pulumi.StringPtrInput
	// Header flag. Valid values: `none`, `http`, `8bit`.
	Header pulumi.StringPtrInput
	// Message type.
	MsgType pulumi.StringInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (NacquarArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nacquarArgs)(nil)).Elem()
}

type NacquarInput interface {
	pulumi.Input

	ToNacquarOutput() NacquarOutput
	ToNacquarOutputWithContext(ctx context.Context) NacquarOutput
}

func (*Nacquar) ElementType() reflect.Type {
	return reflect.TypeOf((**Nacquar)(nil)).Elem()
}

func (i *Nacquar) ToNacquarOutput() NacquarOutput {
	return i.ToNacquarOutputWithContext(context.Background())
}

func (i *Nacquar) ToNacquarOutputWithContext(ctx context.Context) NacquarOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NacquarOutput)
}

// NacquarArrayInput is an input type that accepts NacquarArray and NacquarArrayOutput values.
// You can construct a concrete instance of `NacquarArrayInput` via:
//
//	NacquarArray{ NacquarArgs{...} }
type NacquarArrayInput interface {
	pulumi.Input

	ToNacquarArrayOutput() NacquarArrayOutput
	ToNacquarArrayOutputWithContext(context.Context) NacquarArrayOutput
}

type NacquarArray []NacquarInput

func (NacquarArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Nacquar)(nil)).Elem()
}

func (i NacquarArray) ToNacquarArrayOutput() NacquarArrayOutput {
	return i.ToNacquarArrayOutputWithContext(context.Background())
}

func (i NacquarArray) ToNacquarArrayOutputWithContext(ctx context.Context) NacquarArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NacquarArrayOutput)
}

// NacquarMapInput is an input type that accepts NacquarMap and NacquarMapOutput values.
// You can construct a concrete instance of `NacquarMapInput` via:
//
//	NacquarMap{ "key": NacquarArgs{...} }
type NacquarMapInput interface {
	pulumi.Input

	ToNacquarMapOutput() NacquarMapOutput
	ToNacquarMapOutputWithContext(context.Context) NacquarMapOutput
}

type NacquarMap map[string]NacquarInput

func (NacquarMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Nacquar)(nil)).Elem()
}

func (i NacquarMap) ToNacquarMapOutput() NacquarMapOutput {
	return i.ToNacquarMapOutputWithContext(context.Background())
}

func (i NacquarMap) ToNacquarMapOutputWithContext(ctx context.Context) NacquarMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NacquarMapOutput)
}

type NacquarOutput struct{ *pulumi.OutputState }

func (NacquarOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Nacquar)(nil)).Elem()
}

func (o NacquarOutput) ToNacquarOutput() NacquarOutput {
	return o
}

func (o NacquarOutput) ToNacquarOutputWithContext(ctx context.Context) NacquarOutput {
	return o
}

// Message string.
func (o NacquarOutput) Buffer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Nacquar) pulumi.StringPtrOutput { return v.Buffer }).(pulumi.StringPtrOutput)
}

// Format flag.
func (o NacquarOutput) Format() pulumi.StringOutput {
	return o.ApplyT(func(v *Nacquar) pulumi.StringOutput { return v.Format }).(pulumi.StringOutput)
}

// Header flag. Valid values: `none`, `http`, `8bit`.
func (o NacquarOutput) Header() pulumi.StringOutput {
	return o.ApplyT(func(v *Nacquar) pulumi.StringOutput { return v.Header }).(pulumi.StringOutput)
}

// Message type.
func (o NacquarOutput) MsgType() pulumi.StringOutput {
	return o.ApplyT(func(v *Nacquar) pulumi.StringOutput { return v.MsgType }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o NacquarOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Nacquar) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type NacquarArrayOutput struct{ *pulumi.OutputState }

func (NacquarArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Nacquar)(nil)).Elem()
}

func (o NacquarArrayOutput) ToNacquarArrayOutput() NacquarArrayOutput {
	return o
}

func (o NacquarArrayOutput) ToNacquarArrayOutputWithContext(ctx context.Context) NacquarArrayOutput {
	return o
}

func (o NacquarArrayOutput) Index(i pulumi.IntInput) NacquarOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Nacquar {
		return vs[0].([]*Nacquar)[vs[1].(int)]
	}).(NacquarOutput)
}

type NacquarMapOutput struct{ *pulumi.OutputState }

func (NacquarMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Nacquar)(nil)).Elem()
}

func (o NacquarMapOutput) ToNacquarMapOutput() NacquarMapOutput {
	return o
}

func (o NacquarMapOutput) ToNacquarMapOutputWithContext(ctx context.Context) NacquarMapOutput {
	return o
}

func (o NacquarMapOutput) MapIndex(k pulumi.StringInput) NacquarOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Nacquar {
		return vs[0].(map[string]*Nacquar)[vs[1].(string)]
	}).(NacquarOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NacquarInput)(nil)).Elem(), &Nacquar{})
	pulumi.RegisterInputType(reflect.TypeOf((*NacquarArrayInput)(nil)).Elem(), NacquarArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NacquarMapInput)(nil)).Elem(), NacquarMap{})
	pulumi.RegisterOutputType(NacquarOutput{})
	pulumi.RegisterOutputType(NacquarArrayOutput{})
	pulumi.RegisterOutputType(NacquarMapOutput{})
}