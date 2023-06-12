// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package videofilter

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure YouTube API keys. Applies to FortiOS Version `>= 7.0.1`.
//
// ## Import
//
// # Videofilter YoutubeKey can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:videofilter/youtubekey:Youtubekey labelname {{fosid}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:videofilter/youtubekey:Youtubekey labelname {{fosid}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type Youtubekey struct {
	pulumi.CustomResourceState

	// ID.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// Key.
	Key pulumi.StringOutput `pulumi:"key"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewYoutubekey registers a new resource with the given unique name, arguments, and options.
func NewYoutubekey(ctx *pulumi.Context,
	name string, args *YoutubekeyArgs, opts ...pulumi.ResourceOption) (*Youtubekey, error) {
	if args == nil {
		args = &YoutubekeyArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource Youtubekey
	err := ctx.RegisterResource("fortios:videofilter/youtubekey:Youtubekey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetYoutubekey gets an existing Youtubekey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetYoutubekey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *YoutubekeyState, opts ...pulumi.ResourceOption) (*Youtubekey, error) {
	var resource Youtubekey
	err := ctx.ReadResource("fortios:videofilter/youtubekey:Youtubekey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Youtubekey resources.
type youtubekeyState struct {
	// ID.
	Fosid *int `pulumi:"fosid"`
	// Key.
	Key *string `pulumi:"key"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type YoutubekeyState struct {
	// ID.
	Fosid pulumi.IntPtrInput
	// Key.
	Key pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (YoutubekeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*youtubekeyState)(nil)).Elem()
}

type youtubekeyArgs struct {
	// ID.
	Fosid *int `pulumi:"fosid"`
	// Key.
	Key *string `pulumi:"key"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Youtubekey resource.
type YoutubekeyArgs struct {
	// ID.
	Fosid pulumi.IntPtrInput
	// Key.
	Key pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (YoutubekeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*youtubekeyArgs)(nil)).Elem()
}

type YoutubekeyInput interface {
	pulumi.Input

	ToYoutubekeyOutput() YoutubekeyOutput
	ToYoutubekeyOutputWithContext(ctx context.Context) YoutubekeyOutput
}

func (*Youtubekey) ElementType() reflect.Type {
	return reflect.TypeOf((**Youtubekey)(nil)).Elem()
}

func (i *Youtubekey) ToYoutubekeyOutput() YoutubekeyOutput {
	return i.ToYoutubekeyOutputWithContext(context.Background())
}

func (i *Youtubekey) ToYoutubekeyOutputWithContext(ctx context.Context) YoutubekeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(YoutubekeyOutput)
}

// YoutubekeyArrayInput is an input type that accepts YoutubekeyArray and YoutubekeyArrayOutput values.
// You can construct a concrete instance of `YoutubekeyArrayInput` via:
//
//	YoutubekeyArray{ YoutubekeyArgs{...} }
type YoutubekeyArrayInput interface {
	pulumi.Input

	ToYoutubekeyArrayOutput() YoutubekeyArrayOutput
	ToYoutubekeyArrayOutputWithContext(context.Context) YoutubekeyArrayOutput
}

type YoutubekeyArray []YoutubekeyInput

func (YoutubekeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Youtubekey)(nil)).Elem()
}

func (i YoutubekeyArray) ToYoutubekeyArrayOutput() YoutubekeyArrayOutput {
	return i.ToYoutubekeyArrayOutputWithContext(context.Background())
}

func (i YoutubekeyArray) ToYoutubekeyArrayOutputWithContext(ctx context.Context) YoutubekeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(YoutubekeyArrayOutput)
}

// YoutubekeyMapInput is an input type that accepts YoutubekeyMap and YoutubekeyMapOutput values.
// You can construct a concrete instance of `YoutubekeyMapInput` via:
//
//	YoutubekeyMap{ "key": YoutubekeyArgs{...} }
type YoutubekeyMapInput interface {
	pulumi.Input

	ToYoutubekeyMapOutput() YoutubekeyMapOutput
	ToYoutubekeyMapOutputWithContext(context.Context) YoutubekeyMapOutput
}

type YoutubekeyMap map[string]YoutubekeyInput

func (YoutubekeyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Youtubekey)(nil)).Elem()
}

func (i YoutubekeyMap) ToYoutubekeyMapOutput() YoutubekeyMapOutput {
	return i.ToYoutubekeyMapOutputWithContext(context.Background())
}

func (i YoutubekeyMap) ToYoutubekeyMapOutputWithContext(ctx context.Context) YoutubekeyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(YoutubekeyMapOutput)
}

type YoutubekeyOutput struct{ *pulumi.OutputState }

func (YoutubekeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Youtubekey)(nil)).Elem()
}

func (o YoutubekeyOutput) ToYoutubekeyOutput() YoutubekeyOutput {
	return o
}

func (o YoutubekeyOutput) ToYoutubekeyOutputWithContext(ctx context.Context) YoutubekeyOutput {
	return o
}

// ID.
func (o YoutubekeyOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *Youtubekey) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

// Key.
func (o YoutubekeyOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *Youtubekey) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o YoutubekeyOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Youtubekey) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type YoutubekeyArrayOutput struct{ *pulumi.OutputState }

func (YoutubekeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Youtubekey)(nil)).Elem()
}

func (o YoutubekeyArrayOutput) ToYoutubekeyArrayOutput() YoutubekeyArrayOutput {
	return o
}

func (o YoutubekeyArrayOutput) ToYoutubekeyArrayOutputWithContext(ctx context.Context) YoutubekeyArrayOutput {
	return o
}

func (o YoutubekeyArrayOutput) Index(i pulumi.IntInput) YoutubekeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Youtubekey {
		return vs[0].([]*Youtubekey)[vs[1].(int)]
	}).(YoutubekeyOutput)
}

type YoutubekeyMapOutput struct{ *pulumi.OutputState }

func (YoutubekeyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Youtubekey)(nil)).Elem()
}

func (o YoutubekeyMapOutput) ToYoutubekeyMapOutput() YoutubekeyMapOutput {
	return o
}

func (o YoutubekeyMapOutput) ToYoutubekeyMapOutputWithContext(ctx context.Context) YoutubekeyMapOutput {
	return o
}

func (o YoutubekeyMapOutput) MapIndex(k pulumi.StringInput) YoutubekeyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Youtubekey {
		return vs[0].(map[string]*Youtubekey)[vs[1].(string)]
	}).(YoutubekeyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*YoutubekeyInput)(nil)).Elem(), &Youtubekey{})
	pulumi.RegisterInputType(reflect.TypeOf((*YoutubekeyArrayInput)(nil)).Elem(), YoutubekeyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*YoutubekeyMapInput)(nil)).Elem(), YoutubekeyMap{})
	pulumi.RegisterOutputType(YoutubekeyOutput{})
	pulumi.RegisterOutputType(YoutubekeyArrayOutput{})
	pulumi.RegisterOutputType(YoutubekeyMapOutput{})
}
