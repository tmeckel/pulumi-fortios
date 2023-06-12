// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sys

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure Fortinet Single Sign On (FSSO) server.
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
//			_, err := sys.NewFssopolling(ctx, "trname", &sys.FssopollingArgs{
//				Authentication: pulumi.String("disable"),
//				ListeningPort:  pulumi.Int(8000),
//				Status:         pulumi.String("enable"),
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
// # System FssoPolling can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:sys/fssopolling:Fssopolling labelname SystemFssoPolling
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:sys/fssopolling:Fssopolling labelname SystemFssoPolling
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type Fssopolling struct {
	pulumi.CustomResourceState

	// Password to connect to FSSO Agent.
	AuthPassword pulumi.StringPtrOutput `pulumi:"authPassword"`
	// Enable/disable FSSO Agent Authentication. Valid values: `enable`, `disable`.
	Authentication pulumi.StringOutput `pulumi:"authentication"`
	// Listening port to accept clients (1 - 65535).
	ListeningPort pulumi.IntOutput `pulumi:"listeningPort"`
	// Enable/disable FSSO Polling Mode. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewFssopolling registers a new resource with the given unique name, arguments, and options.
func NewFssopolling(ctx *pulumi.Context,
	name string, args *FssopollingArgs, opts ...pulumi.ResourceOption) (*Fssopolling, error) {
	if args == nil {
		args = &FssopollingArgs{}
	}

	if args.AuthPassword != nil {
		args.AuthPassword = pulumi.ToSecret(args.AuthPassword).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"authPassword",
	})
	opts = append(opts, secrets)
	opts = pkgResourceDefaultOpts(opts)
	var resource Fssopolling
	err := ctx.RegisterResource("fortios:sys/fssopolling:Fssopolling", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFssopolling gets an existing Fssopolling resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFssopolling(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FssopollingState, opts ...pulumi.ResourceOption) (*Fssopolling, error) {
	var resource Fssopolling
	err := ctx.ReadResource("fortios:sys/fssopolling:Fssopolling", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Fssopolling resources.
type fssopollingState struct {
	// Password to connect to FSSO Agent.
	AuthPassword *string `pulumi:"authPassword"`
	// Enable/disable FSSO Agent Authentication. Valid values: `enable`, `disable`.
	Authentication *string `pulumi:"authentication"`
	// Listening port to accept clients (1 - 65535).
	ListeningPort *int `pulumi:"listeningPort"`
	// Enable/disable FSSO Polling Mode. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type FssopollingState struct {
	// Password to connect to FSSO Agent.
	AuthPassword pulumi.StringPtrInput
	// Enable/disable FSSO Agent Authentication. Valid values: `enable`, `disable`.
	Authentication pulumi.StringPtrInput
	// Listening port to accept clients (1 - 65535).
	ListeningPort pulumi.IntPtrInput
	// Enable/disable FSSO Polling Mode. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FssopollingState) ElementType() reflect.Type {
	return reflect.TypeOf((*fssopollingState)(nil)).Elem()
}

type fssopollingArgs struct {
	// Password to connect to FSSO Agent.
	AuthPassword *string `pulumi:"authPassword"`
	// Enable/disable FSSO Agent Authentication. Valid values: `enable`, `disable`.
	Authentication *string `pulumi:"authentication"`
	// Listening port to accept clients (1 - 65535).
	ListeningPort *int `pulumi:"listeningPort"`
	// Enable/disable FSSO Polling Mode. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Fssopolling resource.
type FssopollingArgs struct {
	// Password to connect to FSSO Agent.
	AuthPassword pulumi.StringPtrInput
	// Enable/disable FSSO Agent Authentication. Valid values: `enable`, `disable`.
	Authentication pulumi.StringPtrInput
	// Listening port to accept clients (1 - 65535).
	ListeningPort pulumi.IntPtrInput
	// Enable/disable FSSO Polling Mode. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FssopollingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fssopollingArgs)(nil)).Elem()
}

type FssopollingInput interface {
	pulumi.Input

	ToFssopollingOutput() FssopollingOutput
	ToFssopollingOutputWithContext(ctx context.Context) FssopollingOutput
}

func (*Fssopolling) ElementType() reflect.Type {
	return reflect.TypeOf((**Fssopolling)(nil)).Elem()
}

func (i *Fssopolling) ToFssopollingOutput() FssopollingOutput {
	return i.ToFssopollingOutputWithContext(context.Background())
}

func (i *Fssopolling) ToFssopollingOutputWithContext(ctx context.Context) FssopollingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FssopollingOutput)
}

// FssopollingArrayInput is an input type that accepts FssopollingArray and FssopollingArrayOutput values.
// You can construct a concrete instance of `FssopollingArrayInput` via:
//
//	FssopollingArray{ FssopollingArgs{...} }
type FssopollingArrayInput interface {
	pulumi.Input

	ToFssopollingArrayOutput() FssopollingArrayOutput
	ToFssopollingArrayOutputWithContext(context.Context) FssopollingArrayOutput
}

type FssopollingArray []FssopollingInput

func (FssopollingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Fssopolling)(nil)).Elem()
}

func (i FssopollingArray) ToFssopollingArrayOutput() FssopollingArrayOutput {
	return i.ToFssopollingArrayOutputWithContext(context.Background())
}

func (i FssopollingArray) ToFssopollingArrayOutputWithContext(ctx context.Context) FssopollingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FssopollingArrayOutput)
}

// FssopollingMapInput is an input type that accepts FssopollingMap and FssopollingMapOutput values.
// You can construct a concrete instance of `FssopollingMapInput` via:
//
//	FssopollingMap{ "key": FssopollingArgs{...} }
type FssopollingMapInput interface {
	pulumi.Input

	ToFssopollingMapOutput() FssopollingMapOutput
	ToFssopollingMapOutputWithContext(context.Context) FssopollingMapOutput
}

type FssopollingMap map[string]FssopollingInput

func (FssopollingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Fssopolling)(nil)).Elem()
}

func (i FssopollingMap) ToFssopollingMapOutput() FssopollingMapOutput {
	return i.ToFssopollingMapOutputWithContext(context.Background())
}

func (i FssopollingMap) ToFssopollingMapOutputWithContext(ctx context.Context) FssopollingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FssopollingMapOutput)
}

type FssopollingOutput struct{ *pulumi.OutputState }

func (FssopollingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Fssopolling)(nil)).Elem()
}

func (o FssopollingOutput) ToFssopollingOutput() FssopollingOutput {
	return o
}

func (o FssopollingOutput) ToFssopollingOutputWithContext(ctx context.Context) FssopollingOutput {
	return o
}

// Password to connect to FSSO Agent.
func (o FssopollingOutput) AuthPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Fssopolling) pulumi.StringPtrOutput { return v.AuthPassword }).(pulumi.StringPtrOutput)
}

// Enable/disable FSSO Agent Authentication. Valid values: `enable`, `disable`.
func (o FssopollingOutput) Authentication() pulumi.StringOutput {
	return o.ApplyT(func(v *Fssopolling) pulumi.StringOutput { return v.Authentication }).(pulumi.StringOutput)
}

// Listening port to accept clients (1 - 65535).
func (o FssopollingOutput) ListeningPort() pulumi.IntOutput {
	return o.ApplyT(func(v *Fssopolling) pulumi.IntOutput { return v.ListeningPort }).(pulumi.IntOutput)
}

// Enable/disable FSSO Polling Mode. Valid values: `enable`, `disable`.
func (o FssopollingOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Fssopolling) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o FssopollingOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Fssopolling) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type FssopollingArrayOutput struct{ *pulumi.OutputState }

func (FssopollingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Fssopolling)(nil)).Elem()
}

func (o FssopollingArrayOutput) ToFssopollingArrayOutput() FssopollingArrayOutput {
	return o
}

func (o FssopollingArrayOutput) ToFssopollingArrayOutputWithContext(ctx context.Context) FssopollingArrayOutput {
	return o
}

func (o FssopollingArrayOutput) Index(i pulumi.IntInput) FssopollingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Fssopolling {
		return vs[0].([]*Fssopolling)[vs[1].(int)]
	}).(FssopollingOutput)
}

type FssopollingMapOutput struct{ *pulumi.OutputState }

func (FssopollingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Fssopolling)(nil)).Elem()
}

func (o FssopollingMapOutput) ToFssopollingMapOutput() FssopollingMapOutput {
	return o
}

func (o FssopollingMapOutput) ToFssopollingMapOutputWithContext(ctx context.Context) FssopollingMapOutput {
	return o
}

func (o FssopollingMapOutput) MapIndex(k pulumi.StringInput) FssopollingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Fssopolling {
		return vs[0].(map[string]*Fssopolling)[vs[1].(string)]
	}).(FssopollingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FssopollingInput)(nil)).Elem(), &Fssopolling{})
	pulumi.RegisterInputType(reflect.TypeOf((*FssopollingArrayInput)(nil)).Elem(), FssopollingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FssopollingMapInput)(nil)).Elem(), FssopollingMap{})
	pulumi.RegisterOutputType(FssopollingOutput{})
	pulumi.RegisterOutputType(FssopollingArrayOutput{})
	pulumi.RegisterOutputType(FssopollingMapOutput{})
}
