// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firewall

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure shaping profiles.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/firewall"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := firewall.NewShapingprofile(ctx, "trname", &firewall.ShapingprofileArgs{
//				DefaultClassId: pulumi.Int(2),
//				ProfileName:    pulumi.String("shapingprofiles1"),
//				ShapingEntries: firewall.ShapingprofileShapingEntryArray{
//					&firewall.ShapingprofileShapingEntryArgs{
//						ClassId:                       pulumi.Int(2),
//						GuaranteedBandwidthPercentage: pulumi.Int(33),
//						Id:                            pulumi.Int(1),
//						MaximumBandwidthPercentage:    pulumi.Int(88),
//						Priority:                      pulumi.String("high"),
//					},
//				},
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
// # Firewall ShapingProfile can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:firewall/shapingprofile:Shapingprofile labelname {{profile_name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:firewall/shapingprofile:Shapingprofile labelname {{profile_name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type Shapingprofile struct {
	pulumi.CustomResourceState

	// Comment.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Default class ID to handle unclassified packets (including all local traffic).
	DefaultClassId pulumi.IntOutput `pulumi:"defaultClassId"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Shaping profile name.
	ProfileName pulumi.StringOutput `pulumi:"profileName"`
	// Define shaping entries of this shaping profile. The structure of `shapingEntries` block is documented below.
	ShapingEntries ShapingprofileShapingEntryArrayOutput `pulumi:"shapingEntries"`
	// Select shaping profile type: policing / queuing. Valid values: `policing`, `queuing`.
	Type pulumi.StringOutput `pulumi:"type"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewShapingprofile registers a new resource with the given unique name, arguments, and options.
func NewShapingprofile(ctx *pulumi.Context,
	name string, args *ShapingprofileArgs, opts ...pulumi.ResourceOption) (*Shapingprofile, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DefaultClassId == nil {
		return nil, errors.New("invalid value for required argument 'DefaultClassId'")
	}
	if args.ProfileName == nil {
		return nil, errors.New("invalid value for required argument 'ProfileName'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Shapingprofile
	err := ctx.RegisterResource("fortios:firewall/shapingprofile:Shapingprofile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetShapingprofile gets an existing Shapingprofile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetShapingprofile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ShapingprofileState, opts ...pulumi.ResourceOption) (*Shapingprofile, error) {
	var resource Shapingprofile
	err := ctx.ReadResource("fortios:firewall/shapingprofile:Shapingprofile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Shapingprofile resources.
type shapingprofileState struct {
	// Comment.
	Comment *string `pulumi:"comment"`
	// Default class ID to handle unclassified packets (including all local traffic).
	DefaultClassId *int `pulumi:"defaultClassId"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Shaping profile name.
	ProfileName *string `pulumi:"profileName"`
	// Define shaping entries of this shaping profile. The structure of `shapingEntries` block is documented below.
	ShapingEntries []ShapingprofileShapingEntry `pulumi:"shapingEntries"`
	// Select shaping profile type: policing / queuing. Valid values: `policing`, `queuing`.
	Type *string `pulumi:"type"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type ShapingprofileState struct {
	// Comment.
	Comment pulumi.StringPtrInput
	// Default class ID to handle unclassified packets (including all local traffic).
	DefaultClassId pulumi.IntPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Shaping profile name.
	ProfileName pulumi.StringPtrInput
	// Define shaping entries of this shaping profile. The structure of `shapingEntries` block is documented below.
	ShapingEntries ShapingprofileShapingEntryArrayInput
	// Select shaping profile type: policing / queuing. Valid values: `policing`, `queuing`.
	Type pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (ShapingprofileState) ElementType() reflect.Type {
	return reflect.TypeOf((*shapingprofileState)(nil)).Elem()
}

type shapingprofileArgs struct {
	// Comment.
	Comment *string `pulumi:"comment"`
	// Default class ID to handle unclassified packets (including all local traffic).
	DefaultClassId int `pulumi:"defaultClassId"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Shaping profile name.
	ProfileName string `pulumi:"profileName"`
	// Define shaping entries of this shaping profile. The structure of `shapingEntries` block is documented below.
	ShapingEntries []ShapingprofileShapingEntry `pulumi:"shapingEntries"`
	// Select shaping profile type: policing / queuing. Valid values: `policing`, `queuing`.
	Type *string `pulumi:"type"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Shapingprofile resource.
type ShapingprofileArgs struct {
	// Comment.
	Comment pulumi.StringPtrInput
	// Default class ID to handle unclassified packets (including all local traffic).
	DefaultClassId pulumi.IntInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Shaping profile name.
	ProfileName pulumi.StringInput
	// Define shaping entries of this shaping profile. The structure of `shapingEntries` block is documented below.
	ShapingEntries ShapingprofileShapingEntryArrayInput
	// Select shaping profile type: policing / queuing. Valid values: `policing`, `queuing`.
	Type pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (ShapingprofileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*shapingprofileArgs)(nil)).Elem()
}

type ShapingprofileInput interface {
	pulumi.Input

	ToShapingprofileOutput() ShapingprofileOutput
	ToShapingprofileOutputWithContext(ctx context.Context) ShapingprofileOutput
}

func (*Shapingprofile) ElementType() reflect.Type {
	return reflect.TypeOf((**Shapingprofile)(nil)).Elem()
}

func (i *Shapingprofile) ToShapingprofileOutput() ShapingprofileOutput {
	return i.ToShapingprofileOutputWithContext(context.Background())
}

func (i *Shapingprofile) ToShapingprofileOutputWithContext(ctx context.Context) ShapingprofileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShapingprofileOutput)
}

// ShapingprofileArrayInput is an input type that accepts ShapingprofileArray and ShapingprofileArrayOutput values.
// You can construct a concrete instance of `ShapingprofileArrayInput` via:
//
//	ShapingprofileArray{ ShapingprofileArgs{...} }
type ShapingprofileArrayInput interface {
	pulumi.Input

	ToShapingprofileArrayOutput() ShapingprofileArrayOutput
	ToShapingprofileArrayOutputWithContext(context.Context) ShapingprofileArrayOutput
}

type ShapingprofileArray []ShapingprofileInput

func (ShapingprofileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Shapingprofile)(nil)).Elem()
}

func (i ShapingprofileArray) ToShapingprofileArrayOutput() ShapingprofileArrayOutput {
	return i.ToShapingprofileArrayOutputWithContext(context.Background())
}

func (i ShapingprofileArray) ToShapingprofileArrayOutputWithContext(ctx context.Context) ShapingprofileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShapingprofileArrayOutput)
}

// ShapingprofileMapInput is an input type that accepts ShapingprofileMap and ShapingprofileMapOutput values.
// You can construct a concrete instance of `ShapingprofileMapInput` via:
//
//	ShapingprofileMap{ "key": ShapingprofileArgs{...} }
type ShapingprofileMapInput interface {
	pulumi.Input

	ToShapingprofileMapOutput() ShapingprofileMapOutput
	ToShapingprofileMapOutputWithContext(context.Context) ShapingprofileMapOutput
}

type ShapingprofileMap map[string]ShapingprofileInput

func (ShapingprofileMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Shapingprofile)(nil)).Elem()
}

func (i ShapingprofileMap) ToShapingprofileMapOutput() ShapingprofileMapOutput {
	return i.ToShapingprofileMapOutputWithContext(context.Background())
}

func (i ShapingprofileMap) ToShapingprofileMapOutputWithContext(ctx context.Context) ShapingprofileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShapingprofileMapOutput)
}

type ShapingprofileOutput struct{ *pulumi.OutputState }

func (ShapingprofileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Shapingprofile)(nil)).Elem()
}

func (o ShapingprofileOutput) ToShapingprofileOutput() ShapingprofileOutput {
	return o
}

func (o ShapingprofileOutput) ToShapingprofileOutputWithContext(ctx context.Context) ShapingprofileOutput {
	return o
}

// Comment.
func (o ShapingprofileOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Shapingprofile) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// Default class ID to handle unclassified packets (including all local traffic).
func (o ShapingprofileOutput) DefaultClassId() pulumi.IntOutput {
	return o.ApplyT(func(v *Shapingprofile) pulumi.IntOutput { return v.DefaultClassId }).(pulumi.IntOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o ShapingprofileOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Shapingprofile) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Shaping profile name.
func (o ShapingprofileOutput) ProfileName() pulumi.StringOutput {
	return o.ApplyT(func(v *Shapingprofile) pulumi.StringOutput { return v.ProfileName }).(pulumi.StringOutput)
}

// Define shaping entries of this shaping profile. The structure of `shapingEntries` block is documented below.
func (o ShapingprofileOutput) ShapingEntries() ShapingprofileShapingEntryArrayOutput {
	return o.ApplyT(func(v *Shapingprofile) ShapingprofileShapingEntryArrayOutput { return v.ShapingEntries }).(ShapingprofileShapingEntryArrayOutput)
}

// Select shaping profile type: policing / queuing. Valid values: `policing`, `queuing`.
func (o ShapingprofileOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Shapingprofile) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o ShapingprofileOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Shapingprofile) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type ShapingprofileArrayOutput struct{ *pulumi.OutputState }

func (ShapingprofileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Shapingprofile)(nil)).Elem()
}

func (o ShapingprofileArrayOutput) ToShapingprofileArrayOutput() ShapingprofileArrayOutput {
	return o
}

func (o ShapingprofileArrayOutput) ToShapingprofileArrayOutputWithContext(ctx context.Context) ShapingprofileArrayOutput {
	return o
}

func (o ShapingprofileArrayOutput) Index(i pulumi.IntInput) ShapingprofileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Shapingprofile {
		return vs[0].([]*Shapingprofile)[vs[1].(int)]
	}).(ShapingprofileOutput)
}

type ShapingprofileMapOutput struct{ *pulumi.OutputState }

func (ShapingprofileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Shapingprofile)(nil)).Elem()
}

func (o ShapingprofileMapOutput) ToShapingprofileMapOutput() ShapingprofileMapOutput {
	return o
}

func (o ShapingprofileMapOutput) ToShapingprofileMapOutputWithContext(ctx context.Context) ShapingprofileMapOutput {
	return o
}

func (o ShapingprofileMapOutput) MapIndex(k pulumi.StringInput) ShapingprofileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Shapingprofile {
		return vs[0].(map[string]*Shapingprofile)[vs[1].(string)]
	}).(ShapingprofileOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ShapingprofileInput)(nil)).Elem(), &Shapingprofile{})
	pulumi.RegisterInputType(reflect.TypeOf((*ShapingprofileArrayInput)(nil)).Elem(), ShapingprofileArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ShapingprofileMapInput)(nil)).Elem(), ShapingprofileMap{})
	pulumi.RegisterOutputType(ShapingprofileOutput{})
	pulumi.RegisterOutputType(ShapingprofileArrayOutput{})
	pulumi.RegisterOutputType(ShapingprofileMapOutput{})
}
