// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package videofilter

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure VideoFilter profile. Applies to FortiOS Version `>= 7.0.1`.
//
// ## Import
//
// # Videofilter Profile can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:videofilter/profile:Profile labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:videofilter/profile:Profile labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type Profile struct {
	pulumi.CustomResourceState

	// Comment.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Enable/disable Dailymotion video source. Valid values: `enable`, `disable`.
	Dailymotion pulumi.StringOutput `pulumi:"dailymotion"`
	// Configure FortiGuard categories. The structure of `fortiguardCategory` block is documented below.
	FortiguardCategory ProfileFortiguardCategoryOutput `pulumi:"fortiguardCategory"`
	// Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Replacement message group.
	ReplacemsgGroup pulumi.StringOutput `pulumi:"replacemsgGroup"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// Enable/disable Vimeo video source. Valid values: `enable`, `disable`.
	Vimeo pulumi.StringOutput `pulumi:"vimeo"`
	// Enable/disable YouTube video source. Valid values: `enable`, `disable`.
	Youtube pulumi.StringOutput `pulumi:"youtube"`
	// Set YouTube channel filter.
	YoutubeChannelFilter pulumi.IntOutput `pulumi:"youtubeChannelFilter"`
}

// NewProfile registers a new resource with the given unique name, arguments, and options.
func NewProfile(ctx *pulumi.Context,
	name string, args *ProfileArgs, opts ...pulumi.ResourceOption) (*Profile, error) {
	if args == nil {
		args = &ProfileArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource Profile
	err := ctx.RegisterResource("fortios:videofilter/profile:Profile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProfile gets an existing Profile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProfileState, opts ...pulumi.ResourceOption) (*Profile, error) {
	var resource Profile
	err := ctx.ReadResource("fortios:videofilter/profile:Profile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Profile resources.
type profileState struct {
	// Comment.
	Comment *string `pulumi:"comment"`
	// Enable/disable Dailymotion video source. Valid values: `enable`, `disable`.
	Dailymotion *string `pulumi:"dailymotion"`
	// Configure FortiGuard categories. The structure of `fortiguardCategory` block is documented below.
	FortiguardCategory *ProfileFortiguardCategory `pulumi:"fortiguardCategory"`
	// Name.
	Name *string `pulumi:"name"`
	// Replacement message group.
	ReplacemsgGroup *string `pulumi:"replacemsgGroup"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Enable/disable Vimeo video source. Valid values: `enable`, `disable`.
	Vimeo *string `pulumi:"vimeo"`
	// Enable/disable YouTube video source. Valid values: `enable`, `disable`.
	Youtube *string `pulumi:"youtube"`
	// Set YouTube channel filter.
	YoutubeChannelFilter *int `pulumi:"youtubeChannelFilter"`
}

type ProfileState struct {
	// Comment.
	Comment pulumi.StringPtrInput
	// Enable/disable Dailymotion video source. Valid values: `enable`, `disable`.
	Dailymotion pulumi.StringPtrInput
	// Configure FortiGuard categories. The structure of `fortiguardCategory` block is documented below.
	FortiguardCategory ProfileFortiguardCategoryPtrInput
	// Name.
	Name pulumi.StringPtrInput
	// Replacement message group.
	ReplacemsgGroup pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Enable/disable Vimeo video source. Valid values: `enable`, `disable`.
	Vimeo pulumi.StringPtrInput
	// Enable/disable YouTube video source. Valid values: `enable`, `disable`.
	Youtube pulumi.StringPtrInput
	// Set YouTube channel filter.
	YoutubeChannelFilter pulumi.IntPtrInput
}

func (ProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*profileState)(nil)).Elem()
}

type profileArgs struct {
	// Comment.
	Comment *string `pulumi:"comment"`
	// Enable/disable Dailymotion video source. Valid values: `enable`, `disable`.
	Dailymotion *string `pulumi:"dailymotion"`
	// Configure FortiGuard categories. The structure of `fortiguardCategory` block is documented below.
	FortiguardCategory *ProfileFortiguardCategory `pulumi:"fortiguardCategory"`
	// Name.
	Name *string `pulumi:"name"`
	// Replacement message group.
	ReplacemsgGroup *string `pulumi:"replacemsgGroup"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Enable/disable Vimeo video source. Valid values: `enable`, `disable`.
	Vimeo *string `pulumi:"vimeo"`
	// Enable/disable YouTube video source. Valid values: `enable`, `disable`.
	Youtube *string `pulumi:"youtube"`
	// Set YouTube channel filter.
	YoutubeChannelFilter *int `pulumi:"youtubeChannelFilter"`
}

// The set of arguments for constructing a Profile resource.
type ProfileArgs struct {
	// Comment.
	Comment pulumi.StringPtrInput
	// Enable/disable Dailymotion video source. Valid values: `enable`, `disable`.
	Dailymotion pulumi.StringPtrInput
	// Configure FortiGuard categories. The structure of `fortiguardCategory` block is documented below.
	FortiguardCategory ProfileFortiguardCategoryPtrInput
	// Name.
	Name pulumi.StringPtrInput
	// Replacement message group.
	ReplacemsgGroup pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Enable/disable Vimeo video source. Valid values: `enable`, `disable`.
	Vimeo pulumi.StringPtrInput
	// Enable/disable YouTube video source. Valid values: `enable`, `disable`.
	Youtube pulumi.StringPtrInput
	// Set YouTube channel filter.
	YoutubeChannelFilter pulumi.IntPtrInput
}

func (ProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*profileArgs)(nil)).Elem()
}

type ProfileInput interface {
	pulumi.Input

	ToProfileOutput() ProfileOutput
	ToProfileOutputWithContext(ctx context.Context) ProfileOutput
}

func (*Profile) ElementType() reflect.Type {
	return reflect.TypeOf((**Profile)(nil)).Elem()
}

func (i *Profile) ToProfileOutput() ProfileOutput {
	return i.ToProfileOutputWithContext(context.Background())
}

func (i *Profile) ToProfileOutputWithContext(ctx context.Context) ProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileOutput)
}

// ProfileArrayInput is an input type that accepts ProfileArray and ProfileArrayOutput values.
// You can construct a concrete instance of `ProfileArrayInput` via:
//
//	ProfileArray{ ProfileArgs{...} }
type ProfileArrayInput interface {
	pulumi.Input

	ToProfileArrayOutput() ProfileArrayOutput
	ToProfileArrayOutputWithContext(context.Context) ProfileArrayOutput
}

type ProfileArray []ProfileInput

func (ProfileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Profile)(nil)).Elem()
}

func (i ProfileArray) ToProfileArrayOutput() ProfileArrayOutput {
	return i.ToProfileArrayOutputWithContext(context.Background())
}

func (i ProfileArray) ToProfileArrayOutputWithContext(ctx context.Context) ProfileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileArrayOutput)
}

// ProfileMapInput is an input type that accepts ProfileMap and ProfileMapOutput values.
// You can construct a concrete instance of `ProfileMapInput` via:
//
//	ProfileMap{ "key": ProfileArgs{...} }
type ProfileMapInput interface {
	pulumi.Input

	ToProfileMapOutput() ProfileMapOutput
	ToProfileMapOutputWithContext(context.Context) ProfileMapOutput
}

type ProfileMap map[string]ProfileInput

func (ProfileMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Profile)(nil)).Elem()
}

func (i ProfileMap) ToProfileMapOutput() ProfileMapOutput {
	return i.ToProfileMapOutputWithContext(context.Background())
}

func (i ProfileMap) ToProfileMapOutputWithContext(ctx context.Context) ProfileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileMapOutput)
}

type ProfileOutput struct{ *pulumi.OutputState }

func (ProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Profile)(nil)).Elem()
}

func (o ProfileOutput) ToProfileOutput() ProfileOutput {
	return o
}

func (o ProfileOutput) ToProfileOutputWithContext(ctx context.Context) ProfileOutput {
	return o
}

// Comment.
func (o ProfileOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// Enable/disable Dailymotion video source. Valid values: `enable`, `disable`.
func (o ProfileOutput) Dailymotion() pulumi.StringOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringOutput { return v.Dailymotion }).(pulumi.StringOutput)
}

// Configure FortiGuard categories. The structure of `fortiguardCategory` block is documented below.
func (o ProfileOutput) FortiguardCategory() ProfileFortiguardCategoryOutput {
	return o.ApplyT(func(v *Profile) ProfileFortiguardCategoryOutput { return v.FortiguardCategory }).(ProfileFortiguardCategoryOutput)
}

// Name.
func (o ProfileOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Replacement message group.
func (o ProfileOutput) ReplacemsgGroup() pulumi.StringOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringOutput { return v.ReplacemsgGroup }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o ProfileOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// Enable/disable Vimeo video source. Valid values: `enable`, `disable`.
func (o ProfileOutput) Vimeo() pulumi.StringOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringOutput { return v.Vimeo }).(pulumi.StringOutput)
}

// Enable/disable YouTube video source. Valid values: `enable`, `disable`.
func (o ProfileOutput) Youtube() pulumi.StringOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringOutput { return v.Youtube }).(pulumi.StringOutput)
}

// Set YouTube channel filter.
func (o ProfileOutput) YoutubeChannelFilter() pulumi.IntOutput {
	return o.ApplyT(func(v *Profile) pulumi.IntOutput { return v.YoutubeChannelFilter }).(pulumi.IntOutput)
}

type ProfileArrayOutput struct{ *pulumi.OutputState }

func (ProfileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Profile)(nil)).Elem()
}

func (o ProfileArrayOutput) ToProfileArrayOutput() ProfileArrayOutput {
	return o
}

func (o ProfileArrayOutput) ToProfileArrayOutputWithContext(ctx context.Context) ProfileArrayOutput {
	return o
}

func (o ProfileArrayOutput) Index(i pulumi.IntInput) ProfileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Profile {
		return vs[0].([]*Profile)[vs[1].(int)]
	}).(ProfileOutput)
}

type ProfileMapOutput struct{ *pulumi.OutputState }

func (ProfileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Profile)(nil)).Elem()
}

func (o ProfileMapOutput) ToProfileMapOutput() ProfileMapOutput {
	return o
}

func (o ProfileMapOutput) ToProfileMapOutputWithContext(ctx context.Context) ProfileMapOutput {
	return o
}

func (o ProfileMapOutput) MapIndex(k pulumi.StringInput) ProfileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Profile {
		return vs[0].(map[string]*Profile)[vs[1].(string)]
	}).(ProfileOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProfileInput)(nil)).Elem(), &Profile{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProfileArrayInput)(nil)).Elem(), ProfileArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProfileMapInput)(nil)).Elem(), ProfileMap{})
	pulumi.RegisterOutputType(ProfileOutput{})
	pulumi.RegisterOutputType(ProfileArrayOutput{})
	pulumi.RegisterOutputType(ProfileMapOutput{})
}
