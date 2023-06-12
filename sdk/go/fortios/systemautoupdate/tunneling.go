// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package systemautoupdate

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure web proxy tunnelling for the FDN.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/systemautoupdate"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := systemautoupdate.NewTunneling(ctx, "trname", &systemautoupdate.TunnelingArgs{
//				Port:   pulumi.Int(0),
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
// # SystemAutoupdate Tunneling can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:systemautoupdate/tunneling:Tunneling labelname SystemAutoupdateTunneling
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:systemautoupdate/tunneling:Tunneling labelname SystemAutoupdateTunneling
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type Tunneling struct {
	pulumi.CustomResourceState

	// Web proxy IP address or FQDN.
	Address pulumi.StringOutput `pulumi:"address"`
	// Web proxy password.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// Web proxy port.
	Port pulumi.IntOutput `pulumi:"port"`
	// Enable/disable web proxy tunnelling. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Web proxy username.
	Username pulumi.StringOutput `pulumi:"username"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewTunneling registers a new resource with the given unique name, arguments, and options.
func NewTunneling(ctx *pulumi.Context,
	name string, args *TunnelingArgs, opts ...pulumi.ResourceOption) (*Tunneling, error) {
	if args == nil {
		args = &TunnelingArgs{}
	}

	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
	})
	opts = append(opts, secrets)
	opts = pkgResourceDefaultOpts(opts)
	var resource Tunneling
	err := ctx.RegisterResource("fortios:systemautoupdate/tunneling:Tunneling", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTunneling gets an existing Tunneling resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTunneling(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TunnelingState, opts ...pulumi.ResourceOption) (*Tunneling, error) {
	var resource Tunneling
	err := ctx.ReadResource("fortios:systemautoupdate/tunneling:Tunneling", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Tunneling resources.
type tunnelingState struct {
	// Web proxy IP address or FQDN.
	Address *string `pulumi:"address"`
	// Web proxy password.
	Password *string `pulumi:"password"`
	// Web proxy port.
	Port *int `pulumi:"port"`
	// Enable/disable web proxy tunnelling. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Web proxy username.
	Username *string `pulumi:"username"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type TunnelingState struct {
	// Web proxy IP address or FQDN.
	Address pulumi.StringPtrInput
	// Web proxy password.
	Password pulumi.StringPtrInput
	// Web proxy port.
	Port pulumi.IntPtrInput
	// Enable/disable web proxy tunnelling. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Web proxy username.
	Username pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (TunnelingState) ElementType() reflect.Type {
	return reflect.TypeOf((*tunnelingState)(nil)).Elem()
}

type tunnelingArgs struct {
	// Web proxy IP address or FQDN.
	Address *string `pulumi:"address"`
	// Web proxy password.
	Password *string `pulumi:"password"`
	// Web proxy port.
	Port *int `pulumi:"port"`
	// Enable/disable web proxy tunnelling. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Web proxy username.
	Username *string `pulumi:"username"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Tunneling resource.
type TunnelingArgs struct {
	// Web proxy IP address or FQDN.
	Address pulumi.StringPtrInput
	// Web proxy password.
	Password pulumi.StringPtrInput
	// Web proxy port.
	Port pulumi.IntPtrInput
	// Enable/disable web proxy tunnelling. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Web proxy username.
	Username pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (TunnelingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tunnelingArgs)(nil)).Elem()
}

type TunnelingInput interface {
	pulumi.Input

	ToTunnelingOutput() TunnelingOutput
	ToTunnelingOutputWithContext(ctx context.Context) TunnelingOutput
}

func (*Tunneling) ElementType() reflect.Type {
	return reflect.TypeOf((**Tunneling)(nil)).Elem()
}

func (i *Tunneling) ToTunnelingOutput() TunnelingOutput {
	return i.ToTunnelingOutputWithContext(context.Background())
}

func (i *Tunneling) ToTunnelingOutputWithContext(ctx context.Context) TunnelingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TunnelingOutput)
}

// TunnelingArrayInput is an input type that accepts TunnelingArray and TunnelingArrayOutput values.
// You can construct a concrete instance of `TunnelingArrayInput` via:
//
//	TunnelingArray{ TunnelingArgs{...} }
type TunnelingArrayInput interface {
	pulumi.Input

	ToTunnelingArrayOutput() TunnelingArrayOutput
	ToTunnelingArrayOutputWithContext(context.Context) TunnelingArrayOutput
}

type TunnelingArray []TunnelingInput

func (TunnelingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Tunneling)(nil)).Elem()
}

func (i TunnelingArray) ToTunnelingArrayOutput() TunnelingArrayOutput {
	return i.ToTunnelingArrayOutputWithContext(context.Background())
}

func (i TunnelingArray) ToTunnelingArrayOutputWithContext(ctx context.Context) TunnelingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TunnelingArrayOutput)
}

// TunnelingMapInput is an input type that accepts TunnelingMap and TunnelingMapOutput values.
// You can construct a concrete instance of `TunnelingMapInput` via:
//
//	TunnelingMap{ "key": TunnelingArgs{...} }
type TunnelingMapInput interface {
	pulumi.Input

	ToTunnelingMapOutput() TunnelingMapOutput
	ToTunnelingMapOutputWithContext(context.Context) TunnelingMapOutput
}

type TunnelingMap map[string]TunnelingInput

func (TunnelingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Tunneling)(nil)).Elem()
}

func (i TunnelingMap) ToTunnelingMapOutput() TunnelingMapOutput {
	return i.ToTunnelingMapOutputWithContext(context.Background())
}

func (i TunnelingMap) ToTunnelingMapOutputWithContext(ctx context.Context) TunnelingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TunnelingMapOutput)
}

type TunnelingOutput struct{ *pulumi.OutputState }

func (TunnelingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Tunneling)(nil)).Elem()
}

func (o TunnelingOutput) ToTunnelingOutput() TunnelingOutput {
	return o
}

func (o TunnelingOutput) ToTunnelingOutputWithContext(ctx context.Context) TunnelingOutput {
	return o
}

// Web proxy IP address or FQDN.
func (o TunnelingOutput) Address() pulumi.StringOutput {
	return o.ApplyT(func(v *Tunneling) pulumi.StringOutput { return v.Address }).(pulumi.StringOutput)
}

// Web proxy password.
func (o TunnelingOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Tunneling) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

// Web proxy port.
func (o TunnelingOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v *Tunneling) pulumi.IntOutput { return v.Port }).(pulumi.IntOutput)
}

// Enable/disable web proxy tunnelling. Valid values: `enable`, `disable`.
func (o TunnelingOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Tunneling) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Web proxy username.
func (o TunnelingOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v *Tunneling) pulumi.StringOutput { return v.Username }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o TunnelingOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Tunneling) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type TunnelingArrayOutput struct{ *pulumi.OutputState }

func (TunnelingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Tunneling)(nil)).Elem()
}

func (o TunnelingArrayOutput) ToTunnelingArrayOutput() TunnelingArrayOutput {
	return o
}

func (o TunnelingArrayOutput) ToTunnelingArrayOutputWithContext(ctx context.Context) TunnelingArrayOutput {
	return o
}

func (o TunnelingArrayOutput) Index(i pulumi.IntInput) TunnelingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Tunneling {
		return vs[0].([]*Tunneling)[vs[1].(int)]
	}).(TunnelingOutput)
}

type TunnelingMapOutput struct{ *pulumi.OutputState }

func (TunnelingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Tunneling)(nil)).Elem()
}

func (o TunnelingMapOutput) ToTunnelingMapOutput() TunnelingMapOutput {
	return o
}

func (o TunnelingMapOutput) ToTunnelingMapOutputWithContext(ctx context.Context) TunnelingMapOutput {
	return o
}

func (o TunnelingMapOutput) MapIndex(k pulumi.StringInput) TunnelingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Tunneling {
		return vs[0].(map[string]*Tunneling)[vs[1].(string)]
	}).(TunnelingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TunnelingInput)(nil)).Elem(), &Tunneling{})
	pulumi.RegisterInputType(reflect.TypeOf((*TunnelingArrayInput)(nil)).Elem(), TunnelingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TunnelingMapInput)(nil)).Elem(), TunnelingMap{})
	pulumi.RegisterOutputType(TunnelingOutput{})
	pulumi.RegisterOutputType(TunnelingArrayOutput{})
	pulumi.RegisterOutputType(TunnelingMapOutput{})
}
