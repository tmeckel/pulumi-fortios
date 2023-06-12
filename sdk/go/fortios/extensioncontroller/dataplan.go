// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package extensioncontroller

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// FortiExtender dataplan configuration. Applies to FortiOS Version `>= 7.2.1`.
//
// ## Import
//
// # ExtensionController Dataplan can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:extensioncontroller/dataplan:Dataplan labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:extensioncontroller/dataplan:Dataplan labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type Dataplan struct {
	pulumi.CustomResourceState

	// APN configuration.
	Apn pulumi.StringOutput `pulumi:"apn"`
	// Authentication type. Valid values: `none`, `pap`, `chap`.
	AuthType pulumi.StringOutput `pulumi:"authType"`
	// Billing day of the month (1 - 31).
	BillingDate pulumi.IntOutput `pulumi:"billingDate"`
	// Capacity in MB (0 - 102400000).
	Capacity pulumi.IntOutput `pulumi:"capacity"`
	// Carrier configuration.
	Carrier pulumi.StringOutput `pulumi:"carrier"`
	// ICCID configuration.
	Iccid pulumi.StringOutput `pulumi:"iccid"`
	// Dataplan's modem specifics, if any. Valid values: `modem1`, `modem2`, `all`.
	ModemId pulumi.StringOutput `pulumi:"modemId"`
	// Monthly fee of dataplan (0 - 100000, in local currency).
	MonthlyFee pulumi.IntOutput `pulumi:"monthlyFee"`
	// FortiExtender data plan name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Enable/disable dataplan overage detection. Valid values: `disable`, `enable`.
	Overage pulumi.StringOutput `pulumi:"overage"`
	// Password.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// PDN type. Valid values: `ipv4-only`, `ipv6-only`, `ipv4-ipv6`.
	Pdn pulumi.StringOutput `pulumi:"pdn"`
	// Preferred subnet mask (0 - 32).
	PreferredSubnet pulumi.IntOutput `pulumi:"preferredSubnet"`
	// Enable/disable dataplan private network support. Valid values: `disable`, `enable`.
	PrivateNetwork pulumi.StringOutput `pulumi:"privateNetwork"`
	// Signal period (600 to 18000 seconds).
	SignalPeriod pulumi.IntOutput `pulumi:"signalPeriod"`
	// Signal threshold. Specify the range between 50 - 100, where 50/100 means -50/-100 dBm.
	SignalThreshold pulumi.IntOutput `pulumi:"signalThreshold"`
	// SIM slot configuration. Valid values: `sim1`, `sim2`.
	Slot pulumi.StringOutput `pulumi:"slot"`
	// Type preferences configuration. Valid values: `carrier`, `slot`, `iccid`, `generic`.
	Type pulumi.StringOutput `pulumi:"type"`
	// Username.
	Username pulumi.StringOutput `pulumi:"username"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewDataplan registers a new resource with the given unique name, arguments, and options.
func NewDataplan(ctx *pulumi.Context,
	name string, args *DataplanArgs, opts ...pulumi.ResourceOption) (*Dataplan, error) {
	if args == nil {
		args = &DataplanArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource Dataplan
	err := ctx.RegisterResource("fortios:extensioncontroller/dataplan:Dataplan", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataplan gets an existing Dataplan resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataplan(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataplanState, opts ...pulumi.ResourceOption) (*Dataplan, error) {
	var resource Dataplan
	err := ctx.ReadResource("fortios:extensioncontroller/dataplan:Dataplan", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Dataplan resources.
type dataplanState struct {
	// APN configuration.
	Apn *string `pulumi:"apn"`
	// Authentication type. Valid values: `none`, `pap`, `chap`.
	AuthType *string `pulumi:"authType"`
	// Billing day of the month (1 - 31).
	BillingDate *int `pulumi:"billingDate"`
	// Capacity in MB (0 - 102400000).
	Capacity *int `pulumi:"capacity"`
	// Carrier configuration.
	Carrier *string `pulumi:"carrier"`
	// ICCID configuration.
	Iccid *string `pulumi:"iccid"`
	// Dataplan's modem specifics, if any. Valid values: `modem1`, `modem2`, `all`.
	ModemId *string `pulumi:"modemId"`
	// Monthly fee of dataplan (0 - 100000, in local currency).
	MonthlyFee *int `pulumi:"monthlyFee"`
	// FortiExtender data plan name.
	Name *string `pulumi:"name"`
	// Enable/disable dataplan overage detection. Valid values: `disable`, `enable`.
	Overage *string `pulumi:"overage"`
	// Password.
	Password *string `pulumi:"password"`
	// PDN type. Valid values: `ipv4-only`, `ipv6-only`, `ipv4-ipv6`.
	Pdn *string `pulumi:"pdn"`
	// Preferred subnet mask (0 - 32).
	PreferredSubnet *int `pulumi:"preferredSubnet"`
	// Enable/disable dataplan private network support. Valid values: `disable`, `enable`.
	PrivateNetwork *string `pulumi:"privateNetwork"`
	// Signal period (600 to 18000 seconds).
	SignalPeriod *int `pulumi:"signalPeriod"`
	// Signal threshold. Specify the range between 50 - 100, where 50/100 means -50/-100 dBm.
	SignalThreshold *int `pulumi:"signalThreshold"`
	// SIM slot configuration. Valid values: `sim1`, `sim2`.
	Slot *string `pulumi:"slot"`
	// Type preferences configuration. Valid values: `carrier`, `slot`, `iccid`, `generic`.
	Type *string `pulumi:"type"`
	// Username.
	Username *string `pulumi:"username"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type DataplanState struct {
	// APN configuration.
	Apn pulumi.StringPtrInput
	// Authentication type. Valid values: `none`, `pap`, `chap`.
	AuthType pulumi.StringPtrInput
	// Billing day of the month (1 - 31).
	BillingDate pulumi.IntPtrInput
	// Capacity in MB (0 - 102400000).
	Capacity pulumi.IntPtrInput
	// Carrier configuration.
	Carrier pulumi.StringPtrInput
	// ICCID configuration.
	Iccid pulumi.StringPtrInput
	// Dataplan's modem specifics, if any. Valid values: `modem1`, `modem2`, `all`.
	ModemId pulumi.StringPtrInput
	// Monthly fee of dataplan (0 - 100000, in local currency).
	MonthlyFee pulumi.IntPtrInput
	// FortiExtender data plan name.
	Name pulumi.StringPtrInput
	// Enable/disable dataplan overage detection. Valid values: `disable`, `enable`.
	Overage pulumi.StringPtrInput
	// Password.
	Password pulumi.StringPtrInput
	// PDN type. Valid values: `ipv4-only`, `ipv6-only`, `ipv4-ipv6`.
	Pdn pulumi.StringPtrInput
	// Preferred subnet mask (0 - 32).
	PreferredSubnet pulumi.IntPtrInput
	// Enable/disable dataplan private network support. Valid values: `disable`, `enable`.
	PrivateNetwork pulumi.StringPtrInput
	// Signal period (600 to 18000 seconds).
	SignalPeriod pulumi.IntPtrInput
	// Signal threshold. Specify the range between 50 - 100, where 50/100 means -50/-100 dBm.
	SignalThreshold pulumi.IntPtrInput
	// SIM slot configuration. Valid values: `sim1`, `sim2`.
	Slot pulumi.StringPtrInput
	// Type preferences configuration. Valid values: `carrier`, `slot`, `iccid`, `generic`.
	Type pulumi.StringPtrInput
	// Username.
	Username pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (DataplanState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataplanState)(nil)).Elem()
}

type dataplanArgs struct {
	// APN configuration.
	Apn *string `pulumi:"apn"`
	// Authentication type. Valid values: `none`, `pap`, `chap`.
	AuthType *string `pulumi:"authType"`
	// Billing day of the month (1 - 31).
	BillingDate *int `pulumi:"billingDate"`
	// Capacity in MB (0 - 102400000).
	Capacity *int `pulumi:"capacity"`
	// Carrier configuration.
	Carrier *string `pulumi:"carrier"`
	// ICCID configuration.
	Iccid *string `pulumi:"iccid"`
	// Dataplan's modem specifics, if any. Valid values: `modem1`, `modem2`, `all`.
	ModemId *string `pulumi:"modemId"`
	// Monthly fee of dataplan (0 - 100000, in local currency).
	MonthlyFee *int `pulumi:"monthlyFee"`
	// FortiExtender data plan name.
	Name *string `pulumi:"name"`
	// Enable/disable dataplan overage detection. Valid values: `disable`, `enable`.
	Overage *string `pulumi:"overage"`
	// Password.
	Password *string `pulumi:"password"`
	// PDN type. Valid values: `ipv4-only`, `ipv6-only`, `ipv4-ipv6`.
	Pdn *string `pulumi:"pdn"`
	// Preferred subnet mask (0 - 32).
	PreferredSubnet *int `pulumi:"preferredSubnet"`
	// Enable/disable dataplan private network support. Valid values: `disable`, `enable`.
	PrivateNetwork *string `pulumi:"privateNetwork"`
	// Signal period (600 to 18000 seconds).
	SignalPeriod *int `pulumi:"signalPeriod"`
	// Signal threshold. Specify the range between 50 - 100, where 50/100 means -50/-100 dBm.
	SignalThreshold *int `pulumi:"signalThreshold"`
	// SIM slot configuration. Valid values: `sim1`, `sim2`.
	Slot *string `pulumi:"slot"`
	// Type preferences configuration. Valid values: `carrier`, `slot`, `iccid`, `generic`.
	Type *string `pulumi:"type"`
	// Username.
	Username *string `pulumi:"username"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Dataplan resource.
type DataplanArgs struct {
	// APN configuration.
	Apn pulumi.StringPtrInput
	// Authentication type. Valid values: `none`, `pap`, `chap`.
	AuthType pulumi.StringPtrInput
	// Billing day of the month (1 - 31).
	BillingDate pulumi.IntPtrInput
	// Capacity in MB (0 - 102400000).
	Capacity pulumi.IntPtrInput
	// Carrier configuration.
	Carrier pulumi.StringPtrInput
	// ICCID configuration.
	Iccid pulumi.StringPtrInput
	// Dataplan's modem specifics, if any. Valid values: `modem1`, `modem2`, `all`.
	ModemId pulumi.StringPtrInput
	// Monthly fee of dataplan (0 - 100000, in local currency).
	MonthlyFee pulumi.IntPtrInput
	// FortiExtender data plan name.
	Name pulumi.StringPtrInput
	// Enable/disable dataplan overage detection. Valid values: `disable`, `enable`.
	Overage pulumi.StringPtrInput
	// Password.
	Password pulumi.StringPtrInput
	// PDN type. Valid values: `ipv4-only`, `ipv6-only`, `ipv4-ipv6`.
	Pdn pulumi.StringPtrInput
	// Preferred subnet mask (0 - 32).
	PreferredSubnet pulumi.IntPtrInput
	// Enable/disable dataplan private network support. Valid values: `disable`, `enable`.
	PrivateNetwork pulumi.StringPtrInput
	// Signal period (600 to 18000 seconds).
	SignalPeriod pulumi.IntPtrInput
	// Signal threshold. Specify the range between 50 - 100, where 50/100 means -50/-100 dBm.
	SignalThreshold pulumi.IntPtrInput
	// SIM slot configuration. Valid values: `sim1`, `sim2`.
	Slot pulumi.StringPtrInput
	// Type preferences configuration. Valid values: `carrier`, `slot`, `iccid`, `generic`.
	Type pulumi.StringPtrInput
	// Username.
	Username pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (DataplanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataplanArgs)(nil)).Elem()
}

type DataplanInput interface {
	pulumi.Input

	ToDataplanOutput() DataplanOutput
	ToDataplanOutputWithContext(ctx context.Context) DataplanOutput
}

func (*Dataplan) ElementType() reflect.Type {
	return reflect.TypeOf((**Dataplan)(nil)).Elem()
}

func (i *Dataplan) ToDataplanOutput() DataplanOutput {
	return i.ToDataplanOutputWithContext(context.Background())
}

func (i *Dataplan) ToDataplanOutputWithContext(ctx context.Context) DataplanOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataplanOutput)
}

// DataplanArrayInput is an input type that accepts DataplanArray and DataplanArrayOutput values.
// You can construct a concrete instance of `DataplanArrayInput` via:
//
//	DataplanArray{ DataplanArgs{...} }
type DataplanArrayInput interface {
	pulumi.Input

	ToDataplanArrayOutput() DataplanArrayOutput
	ToDataplanArrayOutputWithContext(context.Context) DataplanArrayOutput
}

type DataplanArray []DataplanInput

func (DataplanArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Dataplan)(nil)).Elem()
}

func (i DataplanArray) ToDataplanArrayOutput() DataplanArrayOutput {
	return i.ToDataplanArrayOutputWithContext(context.Background())
}

func (i DataplanArray) ToDataplanArrayOutputWithContext(ctx context.Context) DataplanArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataplanArrayOutput)
}

// DataplanMapInput is an input type that accepts DataplanMap and DataplanMapOutput values.
// You can construct a concrete instance of `DataplanMapInput` via:
//
//	DataplanMap{ "key": DataplanArgs{...} }
type DataplanMapInput interface {
	pulumi.Input

	ToDataplanMapOutput() DataplanMapOutput
	ToDataplanMapOutputWithContext(context.Context) DataplanMapOutput
}

type DataplanMap map[string]DataplanInput

func (DataplanMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Dataplan)(nil)).Elem()
}

func (i DataplanMap) ToDataplanMapOutput() DataplanMapOutput {
	return i.ToDataplanMapOutputWithContext(context.Background())
}

func (i DataplanMap) ToDataplanMapOutputWithContext(ctx context.Context) DataplanMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataplanMapOutput)
}

type DataplanOutput struct{ *pulumi.OutputState }

func (DataplanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Dataplan)(nil)).Elem()
}

func (o DataplanOutput) ToDataplanOutput() DataplanOutput {
	return o
}

func (o DataplanOutput) ToDataplanOutputWithContext(ctx context.Context) DataplanOutput {
	return o
}

// APN configuration.
func (o DataplanOutput) Apn() pulumi.StringOutput {
	return o.ApplyT(func(v *Dataplan) pulumi.StringOutput { return v.Apn }).(pulumi.StringOutput)
}

// Authentication type. Valid values: `none`, `pap`, `chap`.
func (o DataplanOutput) AuthType() pulumi.StringOutput {
	return o.ApplyT(func(v *Dataplan) pulumi.StringOutput { return v.AuthType }).(pulumi.StringOutput)
}

// Billing day of the month (1 - 31).
func (o DataplanOutput) BillingDate() pulumi.IntOutput {
	return o.ApplyT(func(v *Dataplan) pulumi.IntOutput { return v.BillingDate }).(pulumi.IntOutput)
}

// Capacity in MB (0 - 102400000).
func (o DataplanOutput) Capacity() pulumi.IntOutput {
	return o.ApplyT(func(v *Dataplan) pulumi.IntOutput { return v.Capacity }).(pulumi.IntOutput)
}

// Carrier configuration.
func (o DataplanOutput) Carrier() pulumi.StringOutput {
	return o.ApplyT(func(v *Dataplan) pulumi.StringOutput { return v.Carrier }).(pulumi.StringOutput)
}

// ICCID configuration.
func (o DataplanOutput) Iccid() pulumi.StringOutput {
	return o.ApplyT(func(v *Dataplan) pulumi.StringOutput { return v.Iccid }).(pulumi.StringOutput)
}

// Dataplan's modem specifics, if any. Valid values: `modem1`, `modem2`, `all`.
func (o DataplanOutput) ModemId() pulumi.StringOutput {
	return o.ApplyT(func(v *Dataplan) pulumi.StringOutput { return v.ModemId }).(pulumi.StringOutput)
}

// Monthly fee of dataplan (0 - 100000, in local currency).
func (o DataplanOutput) MonthlyFee() pulumi.IntOutput {
	return o.ApplyT(func(v *Dataplan) pulumi.IntOutput { return v.MonthlyFee }).(pulumi.IntOutput)
}

// FortiExtender data plan name.
func (o DataplanOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Dataplan) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Enable/disable dataplan overage detection. Valid values: `disable`, `enable`.
func (o DataplanOutput) Overage() pulumi.StringOutput {
	return o.ApplyT(func(v *Dataplan) pulumi.StringOutput { return v.Overage }).(pulumi.StringOutput)
}

// Password.
func (o DataplanOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Dataplan) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

// PDN type. Valid values: `ipv4-only`, `ipv6-only`, `ipv4-ipv6`.
func (o DataplanOutput) Pdn() pulumi.StringOutput {
	return o.ApplyT(func(v *Dataplan) pulumi.StringOutput { return v.Pdn }).(pulumi.StringOutput)
}

// Preferred subnet mask (0 - 32).
func (o DataplanOutput) PreferredSubnet() pulumi.IntOutput {
	return o.ApplyT(func(v *Dataplan) pulumi.IntOutput { return v.PreferredSubnet }).(pulumi.IntOutput)
}

// Enable/disable dataplan private network support. Valid values: `disable`, `enable`.
func (o DataplanOutput) PrivateNetwork() pulumi.StringOutput {
	return o.ApplyT(func(v *Dataplan) pulumi.StringOutput { return v.PrivateNetwork }).(pulumi.StringOutput)
}

// Signal period (600 to 18000 seconds).
func (o DataplanOutput) SignalPeriod() pulumi.IntOutput {
	return o.ApplyT(func(v *Dataplan) pulumi.IntOutput { return v.SignalPeriod }).(pulumi.IntOutput)
}

// Signal threshold. Specify the range between 50 - 100, where 50/100 means -50/-100 dBm.
func (o DataplanOutput) SignalThreshold() pulumi.IntOutput {
	return o.ApplyT(func(v *Dataplan) pulumi.IntOutput { return v.SignalThreshold }).(pulumi.IntOutput)
}

// SIM slot configuration. Valid values: `sim1`, `sim2`.
func (o DataplanOutput) Slot() pulumi.StringOutput {
	return o.ApplyT(func(v *Dataplan) pulumi.StringOutput { return v.Slot }).(pulumi.StringOutput)
}

// Type preferences configuration. Valid values: `carrier`, `slot`, `iccid`, `generic`.
func (o DataplanOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Dataplan) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Username.
func (o DataplanOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v *Dataplan) pulumi.StringOutput { return v.Username }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o DataplanOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Dataplan) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type DataplanArrayOutput struct{ *pulumi.OutputState }

func (DataplanArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Dataplan)(nil)).Elem()
}

func (o DataplanArrayOutput) ToDataplanArrayOutput() DataplanArrayOutput {
	return o
}

func (o DataplanArrayOutput) ToDataplanArrayOutputWithContext(ctx context.Context) DataplanArrayOutput {
	return o
}

func (o DataplanArrayOutput) Index(i pulumi.IntInput) DataplanOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Dataplan {
		return vs[0].([]*Dataplan)[vs[1].(int)]
	}).(DataplanOutput)
}

type DataplanMapOutput struct{ *pulumi.OutputState }

func (DataplanMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Dataplan)(nil)).Elem()
}

func (o DataplanMapOutput) ToDataplanMapOutput() DataplanMapOutput {
	return o
}

func (o DataplanMapOutput) ToDataplanMapOutputWithContext(ctx context.Context) DataplanMapOutput {
	return o
}

func (o DataplanMapOutput) MapIndex(k pulumi.StringInput) DataplanOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Dataplan {
		return vs[0].(map[string]*Dataplan)[vs[1].(string)]
	}).(DataplanOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DataplanInput)(nil)).Elem(), &Dataplan{})
	pulumi.RegisterInputType(reflect.TypeOf((*DataplanArrayInput)(nil)).Elem(), DataplanArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DataplanMapInput)(nil)).Elem(), DataplanMap{})
	pulumi.RegisterOutputType(DataplanOutput{})
	pulumi.RegisterOutputType(DataplanArrayOutput{})
	pulumi.RegisterOutputType(DataplanMapOutput{})
}
