// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpncertificate

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// OCSP server configuration.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/vpncertificate"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vpncertificate.NewOcspserver(ctx, "trname", &vpncertificate.OcspserverArgs{
//				Cert:          pulumi.String("ACCVRAIZ1"),
//				SourceIp:      pulumi.String("0.0.0.0"),
//				UnavailAction: pulumi.String("revoke"),
//				Url:           pulumi.String("www.tetserv.com"),
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
// # VpnCertificate OcspServer can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:vpncertificate/ocspserver:Ocspserver labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:vpncertificate/ocspserver:Ocspserver labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type Ocspserver struct {
	pulumi.CustomResourceState

	// OCSP server certificate.
	Cert pulumi.StringOutput `pulumi:"cert"`
	// OCSP server entry name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Secondary OCSP server certificate.
	SecondaryCert pulumi.StringOutput `pulumi:"secondaryCert"`
	// Secondary OCSP server URL.
	SecondaryUrl pulumi.StringOutput `pulumi:"secondaryUrl"`
	// Source IP address for communications to the OCSP server.
	SourceIp pulumi.StringOutput `pulumi:"sourceIp"`
	// Action when server is unavailable (revoke the certificate or ignore the result of the check). Valid values: `revoke`, `ignore`.
	UnavailAction pulumi.StringOutput `pulumi:"unavailAction"`
	// OCSP server URL.
	Url pulumi.StringOutput `pulumi:"url"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewOcspserver registers a new resource with the given unique name, arguments, and options.
func NewOcspserver(ctx *pulumi.Context,
	name string, args *OcspserverArgs, opts ...pulumi.ResourceOption) (*Ocspserver, error) {
	if args == nil {
		args = &OcspserverArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource Ocspserver
	err := ctx.RegisterResource("fortios:vpncertificate/ocspserver:Ocspserver", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOcspserver gets an existing Ocspserver resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOcspserver(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OcspserverState, opts ...pulumi.ResourceOption) (*Ocspserver, error) {
	var resource Ocspserver
	err := ctx.ReadResource("fortios:vpncertificate/ocspserver:Ocspserver", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Ocspserver resources.
type ocspserverState struct {
	// OCSP server certificate.
	Cert *string `pulumi:"cert"`
	// OCSP server entry name.
	Name *string `pulumi:"name"`
	// Secondary OCSP server certificate.
	SecondaryCert *string `pulumi:"secondaryCert"`
	// Secondary OCSP server URL.
	SecondaryUrl *string `pulumi:"secondaryUrl"`
	// Source IP address for communications to the OCSP server.
	SourceIp *string `pulumi:"sourceIp"`
	// Action when server is unavailable (revoke the certificate or ignore the result of the check). Valid values: `revoke`, `ignore`.
	UnavailAction *string `pulumi:"unavailAction"`
	// OCSP server URL.
	Url *string `pulumi:"url"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type OcspserverState struct {
	// OCSP server certificate.
	Cert pulumi.StringPtrInput
	// OCSP server entry name.
	Name pulumi.StringPtrInput
	// Secondary OCSP server certificate.
	SecondaryCert pulumi.StringPtrInput
	// Secondary OCSP server URL.
	SecondaryUrl pulumi.StringPtrInput
	// Source IP address for communications to the OCSP server.
	SourceIp pulumi.StringPtrInput
	// Action when server is unavailable (revoke the certificate or ignore the result of the check). Valid values: `revoke`, `ignore`.
	UnavailAction pulumi.StringPtrInput
	// OCSP server URL.
	Url pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (OcspserverState) ElementType() reflect.Type {
	return reflect.TypeOf((*ocspserverState)(nil)).Elem()
}

type ocspserverArgs struct {
	// OCSP server certificate.
	Cert *string `pulumi:"cert"`
	// OCSP server entry name.
	Name *string `pulumi:"name"`
	// Secondary OCSP server certificate.
	SecondaryCert *string `pulumi:"secondaryCert"`
	// Secondary OCSP server URL.
	SecondaryUrl *string `pulumi:"secondaryUrl"`
	// Source IP address for communications to the OCSP server.
	SourceIp *string `pulumi:"sourceIp"`
	// Action when server is unavailable (revoke the certificate or ignore the result of the check). Valid values: `revoke`, `ignore`.
	UnavailAction *string `pulumi:"unavailAction"`
	// OCSP server URL.
	Url *string `pulumi:"url"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Ocspserver resource.
type OcspserverArgs struct {
	// OCSP server certificate.
	Cert pulumi.StringPtrInput
	// OCSP server entry name.
	Name pulumi.StringPtrInput
	// Secondary OCSP server certificate.
	SecondaryCert pulumi.StringPtrInput
	// Secondary OCSP server URL.
	SecondaryUrl pulumi.StringPtrInput
	// Source IP address for communications to the OCSP server.
	SourceIp pulumi.StringPtrInput
	// Action when server is unavailable (revoke the certificate or ignore the result of the check). Valid values: `revoke`, `ignore`.
	UnavailAction pulumi.StringPtrInput
	// OCSP server URL.
	Url pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (OcspserverArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ocspserverArgs)(nil)).Elem()
}

type OcspserverInput interface {
	pulumi.Input

	ToOcspserverOutput() OcspserverOutput
	ToOcspserverOutputWithContext(ctx context.Context) OcspserverOutput
}

func (*Ocspserver) ElementType() reflect.Type {
	return reflect.TypeOf((**Ocspserver)(nil)).Elem()
}

func (i *Ocspserver) ToOcspserverOutput() OcspserverOutput {
	return i.ToOcspserverOutputWithContext(context.Background())
}

func (i *Ocspserver) ToOcspserverOutputWithContext(ctx context.Context) OcspserverOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OcspserverOutput)
}

// OcspserverArrayInput is an input type that accepts OcspserverArray and OcspserverArrayOutput values.
// You can construct a concrete instance of `OcspserverArrayInput` via:
//
//	OcspserverArray{ OcspserverArgs{...} }
type OcspserverArrayInput interface {
	pulumi.Input

	ToOcspserverArrayOutput() OcspserverArrayOutput
	ToOcspserverArrayOutputWithContext(context.Context) OcspserverArrayOutput
}

type OcspserverArray []OcspserverInput

func (OcspserverArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ocspserver)(nil)).Elem()
}

func (i OcspserverArray) ToOcspserverArrayOutput() OcspserverArrayOutput {
	return i.ToOcspserverArrayOutputWithContext(context.Background())
}

func (i OcspserverArray) ToOcspserverArrayOutputWithContext(ctx context.Context) OcspserverArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OcspserverArrayOutput)
}

// OcspserverMapInput is an input type that accepts OcspserverMap and OcspserverMapOutput values.
// You can construct a concrete instance of `OcspserverMapInput` via:
//
//	OcspserverMap{ "key": OcspserverArgs{...} }
type OcspserverMapInput interface {
	pulumi.Input

	ToOcspserverMapOutput() OcspserverMapOutput
	ToOcspserverMapOutputWithContext(context.Context) OcspserverMapOutput
}

type OcspserverMap map[string]OcspserverInput

func (OcspserverMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ocspserver)(nil)).Elem()
}

func (i OcspserverMap) ToOcspserverMapOutput() OcspserverMapOutput {
	return i.ToOcspserverMapOutputWithContext(context.Background())
}

func (i OcspserverMap) ToOcspserverMapOutputWithContext(ctx context.Context) OcspserverMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OcspserverMapOutput)
}

type OcspserverOutput struct{ *pulumi.OutputState }

func (OcspserverOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Ocspserver)(nil)).Elem()
}

func (o OcspserverOutput) ToOcspserverOutput() OcspserverOutput {
	return o
}

func (o OcspserverOutput) ToOcspserverOutputWithContext(ctx context.Context) OcspserverOutput {
	return o
}

// OCSP server certificate.
func (o OcspserverOutput) Cert() pulumi.StringOutput {
	return o.ApplyT(func(v *Ocspserver) pulumi.StringOutput { return v.Cert }).(pulumi.StringOutput)
}

// OCSP server entry name.
func (o OcspserverOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Ocspserver) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Secondary OCSP server certificate.
func (o OcspserverOutput) SecondaryCert() pulumi.StringOutput {
	return o.ApplyT(func(v *Ocspserver) pulumi.StringOutput { return v.SecondaryCert }).(pulumi.StringOutput)
}

// Secondary OCSP server URL.
func (o OcspserverOutput) SecondaryUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *Ocspserver) pulumi.StringOutput { return v.SecondaryUrl }).(pulumi.StringOutput)
}

// Source IP address for communications to the OCSP server.
func (o OcspserverOutput) SourceIp() pulumi.StringOutput {
	return o.ApplyT(func(v *Ocspserver) pulumi.StringOutput { return v.SourceIp }).(pulumi.StringOutput)
}

// Action when server is unavailable (revoke the certificate or ignore the result of the check). Valid values: `revoke`, `ignore`.
func (o OcspserverOutput) UnavailAction() pulumi.StringOutput {
	return o.ApplyT(func(v *Ocspserver) pulumi.StringOutput { return v.UnavailAction }).(pulumi.StringOutput)
}

// OCSP server URL.
func (o OcspserverOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *Ocspserver) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o OcspserverOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ocspserver) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type OcspserverArrayOutput struct{ *pulumi.OutputState }

func (OcspserverArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ocspserver)(nil)).Elem()
}

func (o OcspserverArrayOutput) ToOcspserverArrayOutput() OcspserverArrayOutput {
	return o
}

func (o OcspserverArrayOutput) ToOcspserverArrayOutputWithContext(ctx context.Context) OcspserverArrayOutput {
	return o
}

func (o OcspserverArrayOutput) Index(i pulumi.IntInput) OcspserverOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Ocspserver {
		return vs[0].([]*Ocspserver)[vs[1].(int)]
	}).(OcspserverOutput)
}

type OcspserverMapOutput struct{ *pulumi.OutputState }

func (OcspserverMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ocspserver)(nil)).Elem()
}

func (o OcspserverMapOutput) ToOcspserverMapOutput() OcspserverMapOutput {
	return o
}

func (o OcspserverMapOutput) ToOcspserverMapOutputWithContext(ctx context.Context) OcspserverMapOutput {
	return o
}

func (o OcspserverMapOutput) MapIndex(k pulumi.StringInput) OcspserverOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Ocspserver {
		return vs[0].(map[string]*Ocspserver)[vs[1].(string)]
	}).(OcspserverOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OcspserverInput)(nil)).Elem(), &Ocspserver{})
	pulumi.RegisterInputType(reflect.TypeOf((*OcspserverArrayInput)(nil)).Elem(), OcspserverArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OcspserverMapInput)(nil)).Elem(), OcspserverMap{})
	pulumi.RegisterOutputType(OcspserverOutput{})
	pulumi.RegisterOutputType(OcspserverArrayOutput{})
	pulumi.RegisterOutputType(OcspserverMapOutput{})
}