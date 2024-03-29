// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package user

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure Kerberos keytab entries.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/user"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			trname2, err := user.NewLdap(ctx, "trname2", &user.LdapArgs{
//				AccountKeyFilter:      pulumi.String("(&(userPrincipalName=%s)(!(UserAccountControl:1.2.840.113556.1.4.803:=2)))"),
//				AccountKeyProcessing:  pulumi.String("same"),
//				Cnid:                  pulumi.String("cn"),
//				Dn:                    pulumi.String("EIWNCIEW"),
//				GroupMemberCheck:      pulumi.String("user-attr"),
//				GroupObjectFilter:     pulumi.String("(&(objectcategory=group)(member=*))"),
//				MemberAttr:            pulumi.String("memberOf"),
//				PasswordExpiryWarning: pulumi.String("disable"),
//				PasswordRenewal:       pulumi.String("disable"),
//				Port:                  pulumi.Int(389),
//				Secure:                pulumi.String("disable"),
//				Server:                pulumi.String("1.1.1.1"),
//				ServerIdentityCheck:   pulumi.String("disable"),
//				SourceIp:              pulumi.String("0.0.0.0"),
//				SslMinProtoVersion:    pulumi.String("default"),
//				Type:                  pulumi.String("simple"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = user.NewKrbkeytab(ctx, "trname", &user.KrbkeytabArgs{
//				Keytab:     pulumi.String("ZXdlY2VxcmVxd3Jld3E="),
//				LdapServer: trname2.Name,
//				Principal:  pulumi.String("testprin"),
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
// # User KrbKeytab can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:user/krbkeytab:Krbkeytab labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:user/krbkeytab:Krbkeytab labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type Krbkeytab struct {
	pulumi.CustomResourceState

	// base64 coded keytab file containing a pre-shared key.
	Keytab pulumi.StringOutput `pulumi:"keytab"`
	// LDAP server name.
	LdapServer pulumi.StringOutput `pulumi:"ldapServer"`
	// Kerberos keytab entry name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Enable/disable parsing PAC data in the ticket. Valid values: `enable`, `disable`.
	PacData pulumi.StringOutput `pulumi:"pacData"`
	// Kerberos service principal, e.g. HTTP/fgt.example.com@EXAMPLE.COM.
	Principal pulumi.StringOutput `pulumi:"principal"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewKrbkeytab registers a new resource with the given unique name, arguments, and options.
func NewKrbkeytab(ctx *pulumi.Context,
	name string, args *KrbkeytabArgs, opts ...pulumi.ResourceOption) (*Krbkeytab, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Keytab == nil {
		return nil, errors.New("invalid value for required argument 'Keytab'")
	}
	if args.LdapServer == nil {
		return nil, errors.New("invalid value for required argument 'LdapServer'")
	}
	if args.Principal == nil {
		return nil, errors.New("invalid value for required argument 'Principal'")
	}
	if args.Keytab != nil {
		args.Keytab = pulumi.ToSecret(args.Keytab).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"keytab",
	})
	opts = append(opts, secrets)
	opts = pkgResourceDefaultOpts(opts)
	var resource Krbkeytab
	err := ctx.RegisterResource("fortios:user/krbkeytab:Krbkeytab", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKrbkeytab gets an existing Krbkeytab resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKrbkeytab(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KrbkeytabState, opts ...pulumi.ResourceOption) (*Krbkeytab, error) {
	var resource Krbkeytab
	err := ctx.ReadResource("fortios:user/krbkeytab:Krbkeytab", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Krbkeytab resources.
type krbkeytabState struct {
	// base64 coded keytab file containing a pre-shared key.
	Keytab *string `pulumi:"keytab"`
	// LDAP server name.
	LdapServer *string `pulumi:"ldapServer"`
	// Kerberos keytab entry name.
	Name *string `pulumi:"name"`
	// Enable/disable parsing PAC data in the ticket. Valid values: `enable`, `disable`.
	PacData *string `pulumi:"pacData"`
	// Kerberos service principal, e.g. HTTP/fgt.example.com@EXAMPLE.COM.
	Principal *string `pulumi:"principal"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type KrbkeytabState struct {
	// base64 coded keytab file containing a pre-shared key.
	Keytab pulumi.StringPtrInput
	// LDAP server name.
	LdapServer pulumi.StringPtrInput
	// Kerberos keytab entry name.
	Name pulumi.StringPtrInput
	// Enable/disable parsing PAC data in the ticket. Valid values: `enable`, `disable`.
	PacData pulumi.StringPtrInput
	// Kerberos service principal, e.g. HTTP/fgt.example.com@EXAMPLE.COM.
	Principal pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (KrbkeytabState) ElementType() reflect.Type {
	return reflect.TypeOf((*krbkeytabState)(nil)).Elem()
}

type krbkeytabArgs struct {
	// base64 coded keytab file containing a pre-shared key.
	Keytab string `pulumi:"keytab"`
	// LDAP server name.
	LdapServer string `pulumi:"ldapServer"`
	// Kerberos keytab entry name.
	Name *string `pulumi:"name"`
	// Enable/disable parsing PAC data in the ticket. Valid values: `enable`, `disable`.
	PacData *string `pulumi:"pacData"`
	// Kerberos service principal, e.g. HTTP/fgt.example.com@EXAMPLE.COM.
	Principal string `pulumi:"principal"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Krbkeytab resource.
type KrbkeytabArgs struct {
	// base64 coded keytab file containing a pre-shared key.
	Keytab pulumi.StringInput
	// LDAP server name.
	LdapServer pulumi.StringInput
	// Kerberos keytab entry name.
	Name pulumi.StringPtrInput
	// Enable/disable parsing PAC data in the ticket. Valid values: `enable`, `disable`.
	PacData pulumi.StringPtrInput
	// Kerberos service principal, e.g. HTTP/fgt.example.com@EXAMPLE.COM.
	Principal pulumi.StringInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (KrbkeytabArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*krbkeytabArgs)(nil)).Elem()
}

type KrbkeytabInput interface {
	pulumi.Input

	ToKrbkeytabOutput() KrbkeytabOutput
	ToKrbkeytabOutputWithContext(ctx context.Context) KrbkeytabOutput
}

func (*Krbkeytab) ElementType() reflect.Type {
	return reflect.TypeOf((**Krbkeytab)(nil)).Elem()
}

func (i *Krbkeytab) ToKrbkeytabOutput() KrbkeytabOutput {
	return i.ToKrbkeytabOutputWithContext(context.Background())
}

func (i *Krbkeytab) ToKrbkeytabOutputWithContext(ctx context.Context) KrbkeytabOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KrbkeytabOutput)
}

// KrbkeytabArrayInput is an input type that accepts KrbkeytabArray and KrbkeytabArrayOutput values.
// You can construct a concrete instance of `KrbkeytabArrayInput` via:
//
//	KrbkeytabArray{ KrbkeytabArgs{...} }
type KrbkeytabArrayInput interface {
	pulumi.Input

	ToKrbkeytabArrayOutput() KrbkeytabArrayOutput
	ToKrbkeytabArrayOutputWithContext(context.Context) KrbkeytabArrayOutput
}

type KrbkeytabArray []KrbkeytabInput

func (KrbkeytabArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Krbkeytab)(nil)).Elem()
}

func (i KrbkeytabArray) ToKrbkeytabArrayOutput() KrbkeytabArrayOutput {
	return i.ToKrbkeytabArrayOutputWithContext(context.Background())
}

func (i KrbkeytabArray) ToKrbkeytabArrayOutputWithContext(ctx context.Context) KrbkeytabArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KrbkeytabArrayOutput)
}

// KrbkeytabMapInput is an input type that accepts KrbkeytabMap and KrbkeytabMapOutput values.
// You can construct a concrete instance of `KrbkeytabMapInput` via:
//
//	KrbkeytabMap{ "key": KrbkeytabArgs{...} }
type KrbkeytabMapInput interface {
	pulumi.Input

	ToKrbkeytabMapOutput() KrbkeytabMapOutput
	ToKrbkeytabMapOutputWithContext(context.Context) KrbkeytabMapOutput
}

type KrbkeytabMap map[string]KrbkeytabInput

func (KrbkeytabMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Krbkeytab)(nil)).Elem()
}

func (i KrbkeytabMap) ToKrbkeytabMapOutput() KrbkeytabMapOutput {
	return i.ToKrbkeytabMapOutputWithContext(context.Background())
}

func (i KrbkeytabMap) ToKrbkeytabMapOutputWithContext(ctx context.Context) KrbkeytabMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KrbkeytabMapOutput)
}

type KrbkeytabOutput struct{ *pulumi.OutputState }

func (KrbkeytabOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Krbkeytab)(nil)).Elem()
}

func (o KrbkeytabOutput) ToKrbkeytabOutput() KrbkeytabOutput {
	return o
}

func (o KrbkeytabOutput) ToKrbkeytabOutputWithContext(ctx context.Context) KrbkeytabOutput {
	return o
}

// base64 coded keytab file containing a pre-shared key.
func (o KrbkeytabOutput) Keytab() pulumi.StringOutput {
	return o.ApplyT(func(v *Krbkeytab) pulumi.StringOutput { return v.Keytab }).(pulumi.StringOutput)
}

// LDAP server name.
func (o KrbkeytabOutput) LdapServer() pulumi.StringOutput {
	return o.ApplyT(func(v *Krbkeytab) pulumi.StringOutput { return v.LdapServer }).(pulumi.StringOutput)
}

// Kerberos keytab entry name.
func (o KrbkeytabOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Krbkeytab) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Enable/disable parsing PAC data in the ticket. Valid values: `enable`, `disable`.
func (o KrbkeytabOutput) PacData() pulumi.StringOutput {
	return o.ApplyT(func(v *Krbkeytab) pulumi.StringOutput { return v.PacData }).(pulumi.StringOutput)
}

// Kerberos service principal, e.g. HTTP/fgt.example.com@EXAMPLE.COM.
func (o KrbkeytabOutput) Principal() pulumi.StringOutput {
	return o.ApplyT(func(v *Krbkeytab) pulumi.StringOutput { return v.Principal }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o KrbkeytabOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Krbkeytab) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type KrbkeytabArrayOutput struct{ *pulumi.OutputState }

func (KrbkeytabArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Krbkeytab)(nil)).Elem()
}

func (o KrbkeytabArrayOutput) ToKrbkeytabArrayOutput() KrbkeytabArrayOutput {
	return o
}

func (o KrbkeytabArrayOutput) ToKrbkeytabArrayOutputWithContext(ctx context.Context) KrbkeytabArrayOutput {
	return o
}

func (o KrbkeytabArrayOutput) Index(i pulumi.IntInput) KrbkeytabOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Krbkeytab {
		return vs[0].([]*Krbkeytab)[vs[1].(int)]
	}).(KrbkeytabOutput)
}

type KrbkeytabMapOutput struct{ *pulumi.OutputState }

func (KrbkeytabMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Krbkeytab)(nil)).Elem()
}

func (o KrbkeytabMapOutput) ToKrbkeytabMapOutput() KrbkeytabMapOutput {
	return o
}

func (o KrbkeytabMapOutput) ToKrbkeytabMapOutputWithContext(ctx context.Context) KrbkeytabMapOutput {
	return o
}

func (o KrbkeytabMapOutput) MapIndex(k pulumi.StringInput) KrbkeytabOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Krbkeytab {
		return vs[0].(map[string]*Krbkeytab)[vs[1].(string)]
	}).(KrbkeytabOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KrbkeytabInput)(nil)).Elem(), &Krbkeytab{})
	pulumi.RegisterInputType(reflect.TypeOf((*KrbkeytabArrayInput)(nil)).Elem(), KrbkeytabArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*KrbkeytabMapInput)(nil)).Elem(), KrbkeytabMap{})
	pulumi.RegisterOutputType(KrbkeytabOutput{})
	pulumi.RegisterOutputType(KrbkeytabArrayOutput{})
	pulumi.RegisterOutputType(KrbkeytabMapOutput{})
}
