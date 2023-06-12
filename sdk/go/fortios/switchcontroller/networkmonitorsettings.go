// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package switchcontroller

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure network monitor settings.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/switchcontroller"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := switchcontroller.NewNetworkmonitorsettings(ctx, "trname", &switchcontroller.NetworkmonitorsettingsArgs{
//				NetworkMonitoring: pulumi.String("disable"),
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
// # SwitchController NetworkMonitorSettings can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:switchcontroller/networkmonitorsettings:Networkmonitorsettings labelname SwitchControllerNetworkMonitorSettings
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:switchcontroller/networkmonitorsettings:Networkmonitorsettings labelname SwitchControllerNetworkMonitorSettings
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type Networkmonitorsettings struct {
	pulumi.CustomResourceState

	// Enable/disable passive gathering of information by FortiSwitch units concerning other network devices. Valid values: `enable`, `disable`.
	NetworkMonitoring pulumi.StringOutput `pulumi:"networkMonitoring"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewNetworkmonitorsettings registers a new resource with the given unique name, arguments, and options.
func NewNetworkmonitorsettings(ctx *pulumi.Context,
	name string, args *NetworkmonitorsettingsArgs, opts ...pulumi.ResourceOption) (*Networkmonitorsettings, error) {
	if args == nil {
		args = &NetworkmonitorsettingsArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource Networkmonitorsettings
	err := ctx.RegisterResource("fortios:switchcontroller/networkmonitorsettings:Networkmonitorsettings", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkmonitorsettings gets an existing Networkmonitorsettings resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkmonitorsettings(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkmonitorsettingsState, opts ...pulumi.ResourceOption) (*Networkmonitorsettings, error) {
	var resource Networkmonitorsettings
	err := ctx.ReadResource("fortios:switchcontroller/networkmonitorsettings:Networkmonitorsettings", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Networkmonitorsettings resources.
type networkmonitorsettingsState struct {
	// Enable/disable passive gathering of information by FortiSwitch units concerning other network devices. Valid values: `enable`, `disable`.
	NetworkMonitoring *string `pulumi:"networkMonitoring"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type NetworkmonitorsettingsState struct {
	// Enable/disable passive gathering of information by FortiSwitch units concerning other network devices. Valid values: `enable`, `disable`.
	NetworkMonitoring pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (NetworkmonitorsettingsState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkmonitorsettingsState)(nil)).Elem()
}

type networkmonitorsettingsArgs struct {
	// Enable/disable passive gathering of information by FortiSwitch units concerning other network devices. Valid values: `enable`, `disable`.
	NetworkMonitoring *string `pulumi:"networkMonitoring"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Networkmonitorsettings resource.
type NetworkmonitorsettingsArgs struct {
	// Enable/disable passive gathering of information by FortiSwitch units concerning other network devices. Valid values: `enable`, `disable`.
	NetworkMonitoring pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (NetworkmonitorsettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkmonitorsettingsArgs)(nil)).Elem()
}

type NetworkmonitorsettingsInput interface {
	pulumi.Input

	ToNetworkmonitorsettingsOutput() NetworkmonitorsettingsOutput
	ToNetworkmonitorsettingsOutputWithContext(ctx context.Context) NetworkmonitorsettingsOutput
}

func (*Networkmonitorsettings) ElementType() reflect.Type {
	return reflect.TypeOf((**Networkmonitorsettings)(nil)).Elem()
}

func (i *Networkmonitorsettings) ToNetworkmonitorsettingsOutput() NetworkmonitorsettingsOutput {
	return i.ToNetworkmonitorsettingsOutputWithContext(context.Background())
}

func (i *Networkmonitorsettings) ToNetworkmonitorsettingsOutputWithContext(ctx context.Context) NetworkmonitorsettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkmonitorsettingsOutput)
}

// NetworkmonitorsettingsArrayInput is an input type that accepts NetworkmonitorsettingsArray and NetworkmonitorsettingsArrayOutput values.
// You can construct a concrete instance of `NetworkmonitorsettingsArrayInput` via:
//
//	NetworkmonitorsettingsArray{ NetworkmonitorsettingsArgs{...} }
type NetworkmonitorsettingsArrayInput interface {
	pulumi.Input

	ToNetworkmonitorsettingsArrayOutput() NetworkmonitorsettingsArrayOutput
	ToNetworkmonitorsettingsArrayOutputWithContext(context.Context) NetworkmonitorsettingsArrayOutput
}

type NetworkmonitorsettingsArray []NetworkmonitorsettingsInput

func (NetworkmonitorsettingsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Networkmonitorsettings)(nil)).Elem()
}

func (i NetworkmonitorsettingsArray) ToNetworkmonitorsettingsArrayOutput() NetworkmonitorsettingsArrayOutput {
	return i.ToNetworkmonitorsettingsArrayOutputWithContext(context.Background())
}

func (i NetworkmonitorsettingsArray) ToNetworkmonitorsettingsArrayOutputWithContext(ctx context.Context) NetworkmonitorsettingsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkmonitorsettingsArrayOutput)
}

// NetworkmonitorsettingsMapInput is an input type that accepts NetworkmonitorsettingsMap and NetworkmonitorsettingsMapOutput values.
// You can construct a concrete instance of `NetworkmonitorsettingsMapInput` via:
//
//	NetworkmonitorsettingsMap{ "key": NetworkmonitorsettingsArgs{...} }
type NetworkmonitorsettingsMapInput interface {
	pulumi.Input

	ToNetworkmonitorsettingsMapOutput() NetworkmonitorsettingsMapOutput
	ToNetworkmonitorsettingsMapOutputWithContext(context.Context) NetworkmonitorsettingsMapOutput
}

type NetworkmonitorsettingsMap map[string]NetworkmonitorsettingsInput

func (NetworkmonitorsettingsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Networkmonitorsettings)(nil)).Elem()
}

func (i NetworkmonitorsettingsMap) ToNetworkmonitorsettingsMapOutput() NetworkmonitorsettingsMapOutput {
	return i.ToNetworkmonitorsettingsMapOutputWithContext(context.Background())
}

func (i NetworkmonitorsettingsMap) ToNetworkmonitorsettingsMapOutputWithContext(ctx context.Context) NetworkmonitorsettingsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkmonitorsettingsMapOutput)
}

type NetworkmonitorsettingsOutput struct{ *pulumi.OutputState }

func (NetworkmonitorsettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Networkmonitorsettings)(nil)).Elem()
}

func (o NetworkmonitorsettingsOutput) ToNetworkmonitorsettingsOutput() NetworkmonitorsettingsOutput {
	return o
}

func (o NetworkmonitorsettingsOutput) ToNetworkmonitorsettingsOutputWithContext(ctx context.Context) NetworkmonitorsettingsOutput {
	return o
}

// Enable/disable passive gathering of information by FortiSwitch units concerning other network devices. Valid values: `enable`, `disable`.
func (o NetworkmonitorsettingsOutput) NetworkMonitoring() pulumi.StringOutput {
	return o.ApplyT(func(v *Networkmonitorsettings) pulumi.StringOutput { return v.NetworkMonitoring }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o NetworkmonitorsettingsOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Networkmonitorsettings) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type NetworkmonitorsettingsArrayOutput struct{ *pulumi.OutputState }

func (NetworkmonitorsettingsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Networkmonitorsettings)(nil)).Elem()
}

func (o NetworkmonitorsettingsArrayOutput) ToNetworkmonitorsettingsArrayOutput() NetworkmonitorsettingsArrayOutput {
	return o
}

func (o NetworkmonitorsettingsArrayOutput) ToNetworkmonitorsettingsArrayOutputWithContext(ctx context.Context) NetworkmonitorsettingsArrayOutput {
	return o
}

func (o NetworkmonitorsettingsArrayOutput) Index(i pulumi.IntInput) NetworkmonitorsettingsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Networkmonitorsettings {
		return vs[0].([]*Networkmonitorsettings)[vs[1].(int)]
	}).(NetworkmonitorsettingsOutput)
}

type NetworkmonitorsettingsMapOutput struct{ *pulumi.OutputState }

func (NetworkmonitorsettingsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Networkmonitorsettings)(nil)).Elem()
}

func (o NetworkmonitorsettingsMapOutput) ToNetworkmonitorsettingsMapOutput() NetworkmonitorsettingsMapOutput {
	return o
}

func (o NetworkmonitorsettingsMapOutput) ToNetworkmonitorsettingsMapOutputWithContext(ctx context.Context) NetworkmonitorsettingsMapOutput {
	return o
}

func (o NetworkmonitorsettingsMapOutput) MapIndex(k pulumi.StringInput) NetworkmonitorsettingsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Networkmonitorsettings {
		return vs[0].(map[string]*Networkmonitorsettings)[vs[1].(string)]
	}).(NetworkmonitorsettingsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkmonitorsettingsInput)(nil)).Elem(), &Networkmonitorsettings{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkmonitorsettingsArrayInput)(nil)).Elem(), NetworkmonitorsettingsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkmonitorsettingsMapInput)(nil)).Elem(), NetworkmonitorsettingsMap{})
	pulumi.RegisterOutputType(NetworkmonitorsettingsOutput{})
	pulumi.RegisterOutputType(NetworkmonitorsettingsArrayOutput{})
	pulumi.RegisterOutputType(NetworkmonitorsettingsMapOutput{})
}
