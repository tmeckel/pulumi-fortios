// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sys

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// # System ApiUser can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:sys/apiuser:Apiuser labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:sys/apiuser:Apiuser labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type Apiuser struct {
	pulumi.CustomResourceState

	// Admin user access profile.
	Accprofile pulumi.StringOutput `pulumi:"accprofile"`
	// Admin user password.
	ApiKey pulumi.StringPtrOutput `pulumi:"apiKey"`
	// Comment.
	Comments pulumi.StringPtrOutput `pulumi:"comments"`
	// Value for Access-Control-Allow-Origin on API responses. Avoid using '*' if possible.
	CorsAllowOrigin pulumi.StringOutput `pulumi:"corsAllowOrigin"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// User name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Enable/disable peer authentication. Valid values: `enable`, `disable`.
	PeerAuth pulumi.StringOutput `pulumi:"peerAuth"`
	// Peer group name.
	PeerGroup pulumi.StringOutput `pulumi:"peerGroup"`
	// Schedule name.
	Schedule pulumi.StringOutput `pulumi:"schedule"`
	// Trusthost. The structure of `trusthost` block is documented below.
	Trusthosts ApiuserTrusthostArrayOutput `pulumi:"trusthosts"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// Virtual domains. The structure of `vdom` block is documented below.
	Vdoms ApiuserVdomArrayOutput `pulumi:"vdoms"`
}

// NewApiuser registers a new resource with the given unique name, arguments, and options.
func NewApiuser(ctx *pulumi.Context,
	name string, args *ApiuserArgs, opts ...pulumi.ResourceOption) (*Apiuser, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Accprofile == nil {
		return nil, errors.New("invalid value for required argument 'Accprofile'")
	}
	if args.ApiKey != nil {
		args.ApiKey = pulumi.ToSecret(args.ApiKey).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"apiKey",
	})
	opts = append(opts, secrets)
	opts = pkgResourceDefaultOpts(opts)
	var resource Apiuser
	err := ctx.RegisterResource("fortios:sys/apiuser:Apiuser", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApiuser gets an existing Apiuser resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApiuser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiuserState, opts ...pulumi.ResourceOption) (*Apiuser, error) {
	var resource Apiuser
	err := ctx.ReadResource("fortios:sys/apiuser:Apiuser", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Apiuser resources.
type apiuserState struct {
	// Admin user access profile.
	Accprofile *string `pulumi:"accprofile"`
	// Admin user password.
	ApiKey *string `pulumi:"apiKey"`
	// Comment.
	Comments *string `pulumi:"comments"`
	// Value for Access-Control-Allow-Origin on API responses. Avoid using '*' if possible.
	CorsAllowOrigin *string `pulumi:"corsAllowOrigin"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// User name.
	Name *string `pulumi:"name"`
	// Enable/disable peer authentication. Valid values: `enable`, `disable`.
	PeerAuth *string `pulumi:"peerAuth"`
	// Peer group name.
	PeerGroup *string `pulumi:"peerGroup"`
	// Schedule name.
	Schedule *string `pulumi:"schedule"`
	// Trusthost. The structure of `trusthost` block is documented below.
	Trusthosts []ApiuserTrusthost `pulumi:"trusthosts"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Virtual domains. The structure of `vdom` block is documented below.
	Vdoms []ApiuserVdom `pulumi:"vdoms"`
}

type ApiuserState struct {
	// Admin user access profile.
	Accprofile pulumi.StringPtrInput
	// Admin user password.
	ApiKey pulumi.StringPtrInput
	// Comment.
	Comments pulumi.StringPtrInput
	// Value for Access-Control-Allow-Origin on API responses. Avoid using '*' if possible.
	CorsAllowOrigin pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// User name.
	Name pulumi.StringPtrInput
	// Enable/disable peer authentication. Valid values: `enable`, `disable`.
	PeerAuth pulumi.StringPtrInput
	// Peer group name.
	PeerGroup pulumi.StringPtrInput
	// Schedule name.
	Schedule pulumi.StringPtrInput
	// Trusthost. The structure of `trusthost` block is documented below.
	Trusthosts ApiuserTrusthostArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Virtual domains. The structure of `vdom` block is documented below.
	Vdoms ApiuserVdomArrayInput
}

func (ApiuserState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiuserState)(nil)).Elem()
}

type apiuserArgs struct {
	// Admin user access profile.
	Accprofile string `pulumi:"accprofile"`
	// Admin user password.
	ApiKey *string `pulumi:"apiKey"`
	// Comment.
	Comments *string `pulumi:"comments"`
	// Value for Access-Control-Allow-Origin on API responses. Avoid using '*' if possible.
	CorsAllowOrigin *string `pulumi:"corsAllowOrigin"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// User name.
	Name *string `pulumi:"name"`
	// Enable/disable peer authentication. Valid values: `enable`, `disable`.
	PeerAuth *string `pulumi:"peerAuth"`
	// Peer group name.
	PeerGroup *string `pulumi:"peerGroup"`
	// Schedule name.
	Schedule *string `pulumi:"schedule"`
	// Trusthost. The structure of `trusthost` block is documented below.
	Trusthosts []ApiuserTrusthost `pulumi:"trusthosts"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Virtual domains. The structure of `vdom` block is documented below.
	Vdoms []ApiuserVdom `pulumi:"vdoms"`
}

// The set of arguments for constructing a Apiuser resource.
type ApiuserArgs struct {
	// Admin user access profile.
	Accprofile pulumi.StringInput
	// Admin user password.
	ApiKey pulumi.StringPtrInput
	// Comment.
	Comments pulumi.StringPtrInput
	// Value for Access-Control-Allow-Origin on API responses. Avoid using '*' if possible.
	CorsAllowOrigin pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// User name.
	Name pulumi.StringPtrInput
	// Enable/disable peer authentication. Valid values: `enable`, `disable`.
	PeerAuth pulumi.StringPtrInput
	// Peer group name.
	PeerGroup pulumi.StringPtrInput
	// Schedule name.
	Schedule pulumi.StringPtrInput
	// Trusthost. The structure of `trusthost` block is documented below.
	Trusthosts ApiuserTrusthostArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Virtual domains. The structure of `vdom` block is documented below.
	Vdoms ApiuserVdomArrayInput
}

func (ApiuserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiuserArgs)(nil)).Elem()
}

type ApiuserInput interface {
	pulumi.Input

	ToApiuserOutput() ApiuserOutput
	ToApiuserOutputWithContext(ctx context.Context) ApiuserOutput
}

func (*Apiuser) ElementType() reflect.Type {
	return reflect.TypeOf((**Apiuser)(nil)).Elem()
}

func (i *Apiuser) ToApiuserOutput() ApiuserOutput {
	return i.ToApiuserOutputWithContext(context.Background())
}

func (i *Apiuser) ToApiuserOutputWithContext(ctx context.Context) ApiuserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiuserOutput)
}

// ApiuserArrayInput is an input type that accepts ApiuserArray and ApiuserArrayOutput values.
// You can construct a concrete instance of `ApiuserArrayInput` via:
//
//	ApiuserArray{ ApiuserArgs{...} }
type ApiuserArrayInput interface {
	pulumi.Input

	ToApiuserArrayOutput() ApiuserArrayOutput
	ToApiuserArrayOutputWithContext(context.Context) ApiuserArrayOutput
}

type ApiuserArray []ApiuserInput

func (ApiuserArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Apiuser)(nil)).Elem()
}

func (i ApiuserArray) ToApiuserArrayOutput() ApiuserArrayOutput {
	return i.ToApiuserArrayOutputWithContext(context.Background())
}

func (i ApiuserArray) ToApiuserArrayOutputWithContext(ctx context.Context) ApiuserArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiuserArrayOutput)
}

// ApiuserMapInput is an input type that accepts ApiuserMap and ApiuserMapOutput values.
// You can construct a concrete instance of `ApiuserMapInput` via:
//
//	ApiuserMap{ "key": ApiuserArgs{...} }
type ApiuserMapInput interface {
	pulumi.Input

	ToApiuserMapOutput() ApiuserMapOutput
	ToApiuserMapOutputWithContext(context.Context) ApiuserMapOutput
}

type ApiuserMap map[string]ApiuserInput

func (ApiuserMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Apiuser)(nil)).Elem()
}

func (i ApiuserMap) ToApiuserMapOutput() ApiuserMapOutput {
	return i.ToApiuserMapOutputWithContext(context.Background())
}

func (i ApiuserMap) ToApiuserMapOutputWithContext(ctx context.Context) ApiuserMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiuserMapOutput)
}

type ApiuserOutput struct{ *pulumi.OutputState }

func (ApiuserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Apiuser)(nil)).Elem()
}

func (o ApiuserOutput) ToApiuserOutput() ApiuserOutput {
	return o
}

func (o ApiuserOutput) ToApiuserOutputWithContext(ctx context.Context) ApiuserOutput {
	return o
}

// Admin user access profile.
func (o ApiuserOutput) Accprofile() pulumi.StringOutput {
	return o.ApplyT(func(v *Apiuser) pulumi.StringOutput { return v.Accprofile }).(pulumi.StringOutput)
}

// Admin user password.
func (o ApiuserOutput) ApiKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Apiuser) pulumi.StringPtrOutput { return v.ApiKey }).(pulumi.StringPtrOutput)
}

// Comment.
func (o ApiuserOutput) Comments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Apiuser) pulumi.StringPtrOutput { return v.Comments }).(pulumi.StringPtrOutput)
}

// Value for Access-Control-Allow-Origin on API responses. Avoid using '*' if possible.
func (o ApiuserOutput) CorsAllowOrigin() pulumi.StringOutput {
	return o.ApplyT(func(v *Apiuser) pulumi.StringOutput { return v.CorsAllowOrigin }).(pulumi.StringOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o ApiuserOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Apiuser) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// User name.
func (o ApiuserOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Apiuser) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Enable/disable peer authentication. Valid values: `enable`, `disable`.
func (o ApiuserOutput) PeerAuth() pulumi.StringOutput {
	return o.ApplyT(func(v *Apiuser) pulumi.StringOutput { return v.PeerAuth }).(pulumi.StringOutput)
}

// Peer group name.
func (o ApiuserOutput) PeerGroup() pulumi.StringOutput {
	return o.ApplyT(func(v *Apiuser) pulumi.StringOutput { return v.PeerGroup }).(pulumi.StringOutput)
}

// Schedule name.
func (o ApiuserOutput) Schedule() pulumi.StringOutput {
	return o.ApplyT(func(v *Apiuser) pulumi.StringOutput { return v.Schedule }).(pulumi.StringOutput)
}

// Trusthost. The structure of `trusthost` block is documented below.
func (o ApiuserOutput) Trusthosts() ApiuserTrusthostArrayOutput {
	return o.ApplyT(func(v *Apiuser) ApiuserTrusthostArrayOutput { return v.Trusthosts }).(ApiuserTrusthostArrayOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o ApiuserOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Apiuser) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// Virtual domains. The structure of `vdom` block is documented below.
func (o ApiuserOutput) Vdoms() ApiuserVdomArrayOutput {
	return o.ApplyT(func(v *Apiuser) ApiuserVdomArrayOutput { return v.Vdoms }).(ApiuserVdomArrayOutput)
}

type ApiuserArrayOutput struct{ *pulumi.OutputState }

func (ApiuserArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Apiuser)(nil)).Elem()
}

func (o ApiuserArrayOutput) ToApiuserArrayOutput() ApiuserArrayOutput {
	return o
}

func (o ApiuserArrayOutput) ToApiuserArrayOutputWithContext(ctx context.Context) ApiuserArrayOutput {
	return o
}

func (o ApiuserArrayOutput) Index(i pulumi.IntInput) ApiuserOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Apiuser {
		return vs[0].([]*Apiuser)[vs[1].(int)]
	}).(ApiuserOutput)
}

type ApiuserMapOutput struct{ *pulumi.OutputState }

func (ApiuserMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Apiuser)(nil)).Elem()
}

func (o ApiuserMapOutput) ToApiuserMapOutput() ApiuserMapOutput {
	return o
}

func (o ApiuserMapOutput) ToApiuserMapOutputWithContext(ctx context.Context) ApiuserMapOutput {
	return o
}

func (o ApiuserMapOutput) MapIndex(k pulumi.StringInput) ApiuserOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Apiuser {
		return vs[0].(map[string]*Apiuser)[vs[1].(string)]
	}).(ApiuserOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApiuserInput)(nil)).Elem(), &Apiuser{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApiuserArrayInput)(nil)).Elem(), ApiuserArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApiuserMapInput)(nil)).Elem(), ApiuserMap{})
	pulumi.RegisterOutputType(ApiuserOutput{})
	pulumi.RegisterOutputType(ApiuserArrayOutput{})
	pulumi.RegisterOutputType(ApiuserMapOutput{})
}