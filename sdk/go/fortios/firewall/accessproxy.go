// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firewall

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure IPv4 access proxy. Applies to FortiOS Version `>= 7.0.1`.
//
// ## Import
//
// # Firewall AccessProxy can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:firewall/accessproxy:Accessproxy labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:firewall/accessproxy:Accessproxy labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type Accessproxy struct {
	pulumi.CustomResourceState

	// Enable/disable adding vhost/domain to dnsdb for ztna dox tunnel. Valid values: `enable`, `disable`.
	AddVhostDomainToDnsdb pulumi.StringOutput `pulumi:"addVhostDomainToDnsdb"`
	// Set IPv6 API Gateway. The structure of `apiGateway6` block is documented below.
	ApiGateway6s AccessproxyApiGateway6ArrayOutput `pulumi:"apiGateway6s"`
	// Set IPv4 API Gateway. The structure of `apiGateway` block is documented below.
	ApiGateways AccessproxyApiGatewayArrayOutput `pulumi:"apiGateways"`
	// Enable/disable authentication portal. Valid values: `disable`, `enable`.
	AuthPortal pulumi.StringOutput `pulumi:"authPortal"`
	// Virtual host for authentication portal.
	AuthVirtualHost pulumi.StringOutput `pulumi:"authVirtualHost"`
	// Enable/disable to request client certificate. Valid values: `disable`, `enable`.
	ClientCert pulumi.StringOutput `pulumi:"clientCert"`
	// Decrypted traffic mirror.
	DecryptedTrafficMirror pulumi.StringOutput `pulumi:"decryptedTrafficMirror"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Action of an empty client certificate.
	EmptyCertAction pulumi.StringOutput `pulumi:"emptyCertAction"`
	// Enable/disable logging of blocked traffic. Valid values: `enable`, `disable`.
	LogBlockedTraffic pulumi.StringOutput `pulumi:"logBlockedTraffic"`
	// Access Proxy name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Enable/disable to detect device type by HTTP user-agent if no client certificate provided. Valid values: `disable`, `enable`.
	UserAgentDetect pulumi.StringOutput `pulumi:"userAgentDetect"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// Virtual IP name.
	Vip pulumi.StringOutput `pulumi:"vip"`
}

// NewAccessproxy registers a new resource with the given unique name, arguments, and options.
func NewAccessproxy(ctx *pulumi.Context,
	name string, args *AccessproxyArgs, opts ...pulumi.ResourceOption) (*Accessproxy, error) {
	if args == nil {
		args = &AccessproxyArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource Accessproxy
	err := ctx.RegisterResource("fortios:firewall/accessproxy:Accessproxy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccessproxy gets an existing Accessproxy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccessproxy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccessproxyState, opts ...pulumi.ResourceOption) (*Accessproxy, error) {
	var resource Accessproxy
	err := ctx.ReadResource("fortios:firewall/accessproxy:Accessproxy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Accessproxy resources.
type accessproxyState struct {
	// Enable/disable adding vhost/domain to dnsdb for ztna dox tunnel. Valid values: `enable`, `disable`.
	AddVhostDomainToDnsdb *string `pulumi:"addVhostDomainToDnsdb"`
	// Set IPv6 API Gateway. The structure of `apiGateway6` block is documented below.
	ApiGateway6s []AccessproxyApiGateway6 `pulumi:"apiGateway6s"`
	// Set IPv4 API Gateway. The structure of `apiGateway` block is documented below.
	ApiGateways []AccessproxyApiGateway `pulumi:"apiGateways"`
	// Enable/disable authentication portal. Valid values: `disable`, `enable`.
	AuthPortal *string `pulumi:"authPortal"`
	// Virtual host for authentication portal.
	AuthVirtualHost *string `pulumi:"authVirtualHost"`
	// Enable/disable to request client certificate. Valid values: `disable`, `enable`.
	ClientCert *string `pulumi:"clientCert"`
	// Decrypted traffic mirror.
	DecryptedTrafficMirror *string `pulumi:"decryptedTrafficMirror"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Action of an empty client certificate.
	EmptyCertAction *string `pulumi:"emptyCertAction"`
	// Enable/disable logging of blocked traffic. Valid values: `enable`, `disable`.
	LogBlockedTraffic *string `pulumi:"logBlockedTraffic"`
	// Access Proxy name.
	Name *string `pulumi:"name"`
	// Enable/disable to detect device type by HTTP user-agent if no client certificate provided. Valid values: `disable`, `enable`.
	UserAgentDetect *string `pulumi:"userAgentDetect"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Virtual IP name.
	Vip *string `pulumi:"vip"`
}

type AccessproxyState struct {
	// Enable/disable adding vhost/domain to dnsdb for ztna dox tunnel. Valid values: `enable`, `disable`.
	AddVhostDomainToDnsdb pulumi.StringPtrInput
	// Set IPv6 API Gateway. The structure of `apiGateway6` block is documented below.
	ApiGateway6s AccessproxyApiGateway6ArrayInput
	// Set IPv4 API Gateway. The structure of `apiGateway` block is documented below.
	ApiGateways AccessproxyApiGatewayArrayInput
	// Enable/disable authentication portal. Valid values: `disable`, `enable`.
	AuthPortal pulumi.StringPtrInput
	// Virtual host for authentication portal.
	AuthVirtualHost pulumi.StringPtrInput
	// Enable/disable to request client certificate. Valid values: `disable`, `enable`.
	ClientCert pulumi.StringPtrInput
	// Decrypted traffic mirror.
	DecryptedTrafficMirror pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Action of an empty client certificate.
	EmptyCertAction pulumi.StringPtrInput
	// Enable/disable logging of blocked traffic. Valid values: `enable`, `disable`.
	LogBlockedTraffic pulumi.StringPtrInput
	// Access Proxy name.
	Name pulumi.StringPtrInput
	// Enable/disable to detect device type by HTTP user-agent if no client certificate provided. Valid values: `disable`, `enable`.
	UserAgentDetect pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Virtual IP name.
	Vip pulumi.StringPtrInput
}

func (AccessproxyState) ElementType() reflect.Type {
	return reflect.TypeOf((*accessproxyState)(nil)).Elem()
}

type accessproxyArgs struct {
	// Enable/disable adding vhost/domain to dnsdb for ztna dox tunnel. Valid values: `enable`, `disable`.
	AddVhostDomainToDnsdb *string `pulumi:"addVhostDomainToDnsdb"`
	// Set IPv6 API Gateway. The structure of `apiGateway6` block is documented below.
	ApiGateway6s []AccessproxyApiGateway6 `pulumi:"apiGateway6s"`
	// Set IPv4 API Gateway. The structure of `apiGateway` block is documented below.
	ApiGateways []AccessproxyApiGateway `pulumi:"apiGateways"`
	// Enable/disable authentication portal. Valid values: `disable`, `enable`.
	AuthPortal *string `pulumi:"authPortal"`
	// Virtual host for authentication portal.
	AuthVirtualHost *string `pulumi:"authVirtualHost"`
	// Enable/disable to request client certificate. Valid values: `disable`, `enable`.
	ClientCert *string `pulumi:"clientCert"`
	// Decrypted traffic mirror.
	DecryptedTrafficMirror *string `pulumi:"decryptedTrafficMirror"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Action of an empty client certificate.
	EmptyCertAction *string `pulumi:"emptyCertAction"`
	// Enable/disable logging of blocked traffic. Valid values: `enable`, `disable`.
	LogBlockedTraffic *string `pulumi:"logBlockedTraffic"`
	// Access Proxy name.
	Name *string `pulumi:"name"`
	// Enable/disable to detect device type by HTTP user-agent if no client certificate provided. Valid values: `disable`, `enable`.
	UserAgentDetect *string `pulumi:"userAgentDetect"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Virtual IP name.
	Vip *string `pulumi:"vip"`
}

// The set of arguments for constructing a Accessproxy resource.
type AccessproxyArgs struct {
	// Enable/disable adding vhost/domain to dnsdb for ztna dox tunnel. Valid values: `enable`, `disable`.
	AddVhostDomainToDnsdb pulumi.StringPtrInput
	// Set IPv6 API Gateway. The structure of `apiGateway6` block is documented below.
	ApiGateway6s AccessproxyApiGateway6ArrayInput
	// Set IPv4 API Gateway. The structure of `apiGateway` block is documented below.
	ApiGateways AccessproxyApiGatewayArrayInput
	// Enable/disable authentication portal. Valid values: `disable`, `enable`.
	AuthPortal pulumi.StringPtrInput
	// Virtual host for authentication portal.
	AuthVirtualHost pulumi.StringPtrInput
	// Enable/disable to request client certificate. Valid values: `disable`, `enable`.
	ClientCert pulumi.StringPtrInput
	// Decrypted traffic mirror.
	DecryptedTrafficMirror pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Action of an empty client certificate.
	EmptyCertAction pulumi.StringPtrInput
	// Enable/disable logging of blocked traffic. Valid values: `enable`, `disable`.
	LogBlockedTraffic pulumi.StringPtrInput
	// Access Proxy name.
	Name pulumi.StringPtrInput
	// Enable/disable to detect device type by HTTP user-agent if no client certificate provided. Valid values: `disable`, `enable`.
	UserAgentDetect pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Virtual IP name.
	Vip pulumi.StringPtrInput
}

func (AccessproxyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accessproxyArgs)(nil)).Elem()
}

type AccessproxyInput interface {
	pulumi.Input

	ToAccessproxyOutput() AccessproxyOutput
	ToAccessproxyOutputWithContext(ctx context.Context) AccessproxyOutput
}

func (*Accessproxy) ElementType() reflect.Type {
	return reflect.TypeOf((**Accessproxy)(nil)).Elem()
}

func (i *Accessproxy) ToAccessproxyOutput() AccessproxyOutput {
	return i.ToAccessproxyOutputWithContext(context.Background())
}

func (i *Accessproxy) ToAccessproxyOutputWithContext(ctx context.Context) AccessproxyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessproxyOutput)
}

// AccessproxyArrayInput is an input type that accepts AccessproxyArray and AccessproxyArrayOutput values.
// You can construct a concrete instance of `AccessproxyArrayInput` via:
//
//	AccessproxyArray{ AccessproxyArgs{...} }
type AccessproxyArrayInput interface {
	pulumi.Input

	ToAccessproxyArrayOutput() AccessproxyArrayOutput
	ToAccessproxyArrayOutputWithContext(context.Context) AccessproxyArrayOutput
}

type AccessproxyArray []AccessproxyInput

func (AccessproxyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Accessproxy)(nil)).Elem()
}

func (i AccessproxyArray) ToAccessproxyArrayOutput() AccessproxyArrayOutput {
	return i.ToAccessproxyArrayOutputWithContext(context.Background())
}

func (i AccessproxyArray) ToAccessproxyArrayOutputWithContext(ctx context.Context) AccessproxyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessproxyArrayOutput)
}

// AccessproxyMapInput is an input type that accepts AccessproxyMap and AccessproxyMapOutput values.
// You can construct a concrete instance of `AccessproxyMapInput` via:
//
//	AccessproxyMap{ "key": AccessproxyArgs{...} }
type AccessproxyMapInput interface {
	pulumi.Input

	ToAccessproxyMapOutput() AccessproxyMapOutput
	ToAccessproxyMapOutputWithContext(context.Context) AccessproxyMapOutput
}

type AccessproxyMap map[string]AccessproxyInput

func (AccessproxyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Accessproxy)(nil)).Elem()
}

func (i AccessproxyMap) ToAccessproxyMapOutput() AccessproxyMapOutput {
	return i.ToAccessproxyMapOutputWithContext(context.Background())
}

func (i AccessproxyMap) ToAccessproxyMapOutputWithContext(ctx context.Context) AccessproxyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessproxyMapOutput)
}

type AccessproxyOutput struct{ *pulumi.OutputState }

func (AccessproxyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Accessproxy)(nil)).Elem()
}

func (o AccessproxyOutput) ToAccessproxyOutput() AccessproxyOutput {
	return o
}

func (o AccessproxyOutput) ToAccessproxyOutputWithContext(ctx context.Context) AccessproxyOutput {
	return o
}

// Enable/disable adding vhost/domain to dnsdb for ztna dox tunnel. Valid values: `enable`, `disable`.
func (o AccessproxyOutput) AddVhostDomainToDnsdb() pulumi.StringOutput {
	return o.ApplyT(func(v *Accessproxy) pulumi.StringOutput { return v.AddVhostDomainToDnsdb }).(pulumi.StringOutput)
}

// Set IPv6 API Gateway. The structure of `apiGateway6` block is documented below.
func (o AccessproxyOutput) ApiGateway6s() AccessproxyApiGateway6ArrayOutput {
	return o.ApplyT(func(v *Accessproxy) AccessproxyApiGateway6ArrayOutput { return v.ApiGateway6s }).(AccessproxyApiGateway6ArrayOutput)
}

// Set IPv4 API Gateway. The structure of `apiGateway` block is documented below.
func (o AccessproxyOutput) ApiGateways() AccessproxyApiGatewayArrayOutput {
	return o.ApplyT(func(v *Accessproxy) AccessproxyApiGatewayArrayOutput { return v.ApiGateways }).(AccessproxyApiGatewayArrayOutput)
}

// Enable/disable authentication portal. Valid values: `disable`, `enable`.
func (o AccessproxyOutput) AuthPortal() pulumi.StringOutput {
	return o.ApplyT(func(v *Accessproxy) pulumi.StringOutput { return v.AuthPortal }).(pulumi.StringOutput)
}

// Virtual host for authentication portal.
func (o AccessproxyOutput) AuthVirtualHost() pulumi.StringOutput {
	return o.ApplyT(func(v *Accessproxy) pulumi.StringOutput { return v.AuthVirtualHost }).(pulumi.StringOutput)
}

// Enable/disable to request client certificate. Valid values: `disable`, `enable`.
func (o AccessproxyOutput) ClientCert() pulumi.StringOutput {
	return o.ApplyT(func(v *Accessproxy) pulumi.StringOutput { return v.ClientCert }).(pulumi.StringOutput)
}

// Decrypted traffic mirror.
func (o AccessproxyOutput) DecryptedTrafficMirror() pulumi.StringOutput {
	return o.ApplyT(func(v *Accessproxy) pulumi.StringOutput { return v.DecryptedTrafficMirror }).(pulumi.StringOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o AccessproxyOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Accessproxy) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Action of an empty client certificate.
func (o AccessproxyOutput) EmptyCertAction() pulumi.StringOutput {
	return o.ApplyT(func(v *Accessproxy) pulumi.StringOutput { return v.EmptyCertAction }).(pulumi.StringOutput)
}

// Enable/disable logging of blocked traffic. Valid values: `enable`, `disable`.
func (o AccessproxyOutput) LogBlockedTraffic() pulumi.StringOutput {
	return o.ApplyT(func(v *Accessproxy) pulumi.StringOutput { return v.LogBlockedTraffic }).(pulumi.StringOutput)
}

// Access Proxy name.
func (o AccessproxyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Accessproxy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Enable/disable to detect device type by HTTP user-agent if no client certificate provided. Valid values: `disable`, `enable`.
func (o AccessproxyOutput) UserAgentDetect() pulumi.StringOutput {
	return o.ApplyT(func(v *Accessproxy) pulumi.StringOutput { return v.UserAgentDetect }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o AccessproxyOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Accessproxy) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// Virtual IP name.
func (o AccessproxyOutput) Vip() pulumi.StringOutput {
	return o.ApplyT(func(v *Accessproxy) pulumi.StringOutput { return v.Vip }).(pulumi.StringOutput)
}

type AccessproxyArrayOutput struct{ *pulumi.OutputState }

func (AccessproxyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Accessproxy)(nil)).Elem()
}

func (o AccessproxyArrayOutput) ToAccessproxyArrayOutput() AccessproxyArrayOutput {
	return o
}

func (o AccessproxyArrayOutput) ToAccessproxyArrayOutputWithContext(ctx context.Context) AccessproxyArrayOutput {
	return o
}

func (o AccessproxyArrayOutput) Index(i pulumi.IntInput) AccessproxyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Accessproxy {
		return vs[0].([]*Accessproxy)[vs[1].(int)]
	}).(AccessproxyOutput)
}

type AccessproxyMapOutput struct{ *pulumi.OutputState }

func (AccessproxyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Accessproxy)(nil)).Elem()
}

func (o AccessproxyMapOutput) ToAccessproxyMapOutput() AccessproxyMapOutput {
	return o
}

func (o AccessproxyMapOutput) ToAccessproxyMapOutputWithContext(ctx context.Context) AccessproxyMapOutput {
	return o
}

func (o AccessproxyMapOutput) MapIndex(k pulumi.StringInput) AccessproxyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Accessproxy {
		return vs[0].(map[string]*Accessproxy)[vs[1].(string)]
	}).(AccessproxyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AccessproxyInput)(nil)).Elem(), &Accessproxy{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccessproxyArrayInput)(nil)).Elem(), AccessproxyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccessproxyMapInput)(nil)).Elem(), AccessproxyMap{})
	pulumi.RegisterOutputType(AccessproxyOutput{})
	pulumi.RegisterOutputType(AccessproxyArrayOutput{})
	pulumi.RegisterOutputType(AccessproxyMapOutput{})
}