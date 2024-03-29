// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package user

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure local users.
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
//			trname3, err := user.NewLdap(ctx, "trname3", &user.LdapArgs{
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
//			_, err = user.NewLocal(ctx, "trname", &user.LocalArgs{
//				AuthConcurrentOverride: pulumi.String("disable"),
//				AuthConcurrentValue:    pulumi.Int(0),
//				Authtimeout:            pulumi.Int(0),
//				LdapServer:             trname3.Name,
//				PasswdTime:             pulumi.String("0000-00-00 00:00:00"),
//				SmsServer:              pulumi.String("fortiguard"),
//				Status:                 pulumi.String("enable"),
//				TwoFactor:              pulumi.String("disable"),
//				Type:                   pulumi.String("ldap"),
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
// # User Local can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:user/local:Local labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:user/local:Local labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type Local struct {
	pulumi.CustomResourceState

	// Enable/disable overriding the policy-auth-concurrent under config system global. Valid values: `enable`, `disable`.
	AuthConcurrentOverride pulumi.StringOutput `pulumi:"authConcurrentOverride"`
	// Maximum number of concurrent logins permitted from the same user.
	AuthConcurrentValue pulumi.IntOutput `pulumi:"authConcurrentValue"`
	// Time in minutes before the authentication timeout for a user is reached.
	Authtimeout pulumi.IntOutput `pulumi:"authtimeout"`
	// Two-factor recipient's email address.
	EmailTo pulumi.StringOutput `pulumi:"emailTo"`
	// Two-factor recipient's FortiToken serial number.
	Fortitoken pulumi.StringOutput `pulumi:"fortitoken"`
	// User ID.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// Name of LDAP server with which the user must authenticate.
	LdapServer pulumi.StringOutput `pulumi:"ldapServer"`
	// User name.
	Name pulumi.StringOutput `pulumi:"name"`
	// User's password.
	Passwd pulumi.StringPtrOutput `pulumi:"passwd"`
	// Password policy to apply to this user, as defined in config user password-policy.
	PasswdPolicy pulumi.StringOutput `pulumi:"passwdPolicy"`
	// Time of the last password update.
	PasswdTime pulumi.StringOutput `pulumi:"passwdTime"`
	// IKEv2 Postquantum Preshared Key Identity.
	PpkIdentity pulumi.StringOutput `pulumi:"ppkIdentity"`
	// IKEv2 Postquantum Preshared Key (ASCII string or hexadecimal encoded with a leading 0x).
	PpkSecret pulumi.StringPtrOutput `pulumi:"ppkSecret"`
	// Name of RADIUS server with which the user must authenticate.
	RadiusServer pulumi.StringOutput `pulumi:"radiusServer"`
	// Two-factor recipient's SMS server.
	SmsCustomServer pulumi.StringOutput `pulumi:"smsCustomServer"`
	// Two-factor recipient's mobile phone number.
	SmsPhone pulumi.StringOutput `pulumi:"smsPhone"`
	// Send SMS through FortiGuard or other external server. Valid values: `fortiguard`, `custom`.
	SmsServer pulumi.StringOutput `pulumi:"smsServer"`
	// Enable/disable allowing the local user to authenticate with the FortiGate unit. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Name of TACACS+ server with which the user must authenticate.
	TacacsServer pulumi.StringOutput `pulumi:"tacacsServer"`
	// Enable/disable two-factor authentication.
	TwoFactor pulumi.StringOutput `pulumi:"twoFactor"`
	// Authentication method by FortiToken Cloud. Valid values: `fortitoken`, `email`, `sms`.
	TwoFactorAuthentication pulumi.StringOutput `pulumi:"twoFactorAuthentication"`
	// Notification method for user activation by FortiToken Cloud. Valid values: `email`, `sms`.
	TwoFactorNotification pulumi.StringOutput `pulumi:"twoFactorNotification"`
	// Authentication method. Valid values: `password`, `radius`, `tacacs+`, `ldap`.
	Type pulumi.StringOutput `pulumi:"type"`
	// Enable/disable case sensitivity when performing username matching (uppercase and lowercase letters are treated either as distinct or equivalent). Valid values: `enable`, `disable`.
	UsernameCaseInsensitivity pulumi.StringOutput `pulumi:"usernameCaseInsensitivity"`
	// Enable/disable case sensitivity when performing username matching (uppercase and lowercase letters are treated either as distinct or equivalent). Valid values: `disable`, `enable`.
	UsernameCaseSensitivity pulumi.StringOutput `pulumi:"usernameCaseSensitivity"`
	// Enable/disable case and accent sensitivity when performing username matching (accents are stripped and case is ignored when disabled). Valid values: `disable`, `enable`.
	UsernameSensitivity pulumi.StringOutput `pulumi:"usernameSensitivity"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// Name of the remote user workstation, if you want to limit the user to authenticate only from a particular workstation.
	Workstation pulumi.StringOutput `pulumi:"workstation"`
}

// NewLocal registers a new resource with the given unique name, arguments, and options.
func NewLocal(ctx *pulumi.Context,
	name string, args *LocalArgs, opts ...pulumi.ResourceOption) (*Local, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Status == nil {
		return nil, errors.New("invalid value for required argument 'Status'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	if args.Passwd != nil {
		args.Passwd = pulumi.ToSecret(args.Passwd).(pulumi.StringPtrInput)
	}
	if args.PpkSecret != nil {
		args.PpkSecret = pulumi.ToSecret(args.PpkSecret).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"passwd",
		"ppkSecret",
	})
	opts = append(opts, secrets)
	opts = pkgResourceDefaultOpts(opts)
	var resource Local
	err := ctx.RegisterResource("fortios:user/local:Local", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLocal gets an existing Local resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLocal(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LocalState, opts ...pulumi.ResourceOption) (*Local, error) {
	var resource Local
	err := ctx.ReadResource("fortios:user/local:Local", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Local resources.
type localState struct {
	// Enable/disable overriding the policy-auth-concurrent under config system global. Valid values: `enable`, `disable`.
	AuthConcurrentOverride *string `pulumi:"authConcurrentOverride"`
	// Maximum number of concurrent logins permitted from the same user.
	AuthConcurrentValue *int `pulumi:"authConcurrentValue"`
	// Time in minutes before the authentication timeout for a user is reached.
	Authtimeout *int `pulumi:"authtimeout"`
	// Two-factor recipient's email address.
	EmailTo *string `pulumi:"emailTo"`
	// Two-factor recipient's FortiToken serial number.
	Fortitoken *string `pulumi:"fortitoken"`
	// User ID.
	Fosid *int `pulumi:"fosid"`
	// Name of LDAP server with which the user must authenticate.
	LdapServer *string `pulumi:"ldapServer"`
	// User name.
	Name *string `pulumi:"name"`
	// User's password.
	Passwd *string `pulumi:"passwd"`
	// Password policy to apply to this user, as defined in config user password-policy.
	PasswdPolicy *string `pulumi:"passwdPolicy"`
	// Time of the last password update.
	PasswdTime *string `pulumi:"passwdTime"`
	// IKEv2 Postquantum Preshared Key Identity.
	PpkIdentity *string `pulumi:"ppkIdentity"`
	// IKEv2 Postquantum Preshared Key (ASCII string or hexadecimal encoded with a leading 0x).
	PpkSecret *string `pulumi:"ppkSecret"`
	// Name of RADIUS server with which the user must authenticate.
	RadiusServer *string `pulumi:"radiusServer"`
	// Two-factor recipient's SMS server.
	SmsCustomServer *string `pulumi:"smsCustomServer"`
	// Two-factor recipient's mobile phone number.
	SmsPhone *string `pulumi:"smsPhone"`
	// Send SMS through FortiGuard or other external server. Valid values: `fortiguard`, `custom`.
	SmsServer *string `pulumi:"smsServer"`
	// Enable/disable allowing the local user to authenticate with the FortiGate unit. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Name of TACACS+ server with which the user must authenticate.
	TacacsServer *string `pulumi:"tacacsServer"`
	// Enable/disable two-factor authentication.
	TwoFactor *string `pulumi:"twoFactor"`
	// Authentication method by FortiToken Cloud. Valid values: `fortitoken`, `email`, `sms`.
	TwoFactorAuthentication *string `pulumi:"twoFactorAuthentication"`
	// Notification method for user activation by FortiToken Cloud. Valid values: `email`, `sms`.
	TwoFactorNotification *string `pulumi:"twoFactorNotification"`
	// Authentication method. Valid values: `password`, `radius`, `tacacs+`, `ldap`.
	Type *string `pulumi:"type"`
	// Enable/disable case sensitivity when performing username matching (uppercase and lowercase letters are treated either as distinct or equivalent). Valid values: `enable`, `disable`.
	UsernameCaseInsensitivity *string `pulumi:"usernameCaseInsensitivity"`
	// Enable/disable case sensitivity when performing username matching (uppercase and lowercase letters are treated either as distinct or equivalent). Valid values: `disable`, `enable`.
	UsernameCaseSensitivity *string `pulumi:"usernameCaseSensitivity"`
	// Enable/disable case and accent sensitivity when performing username matching (accents are stripped and case is ignored when disabled). Valid values: `disable`, `enable`.
	UsernameSensitivity *string `pulumi:"usernameSensitivity"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Name of the remote user workstation, if you want to limit the user to authenticate only from a particular workstation.
	Workstation *string `pulumi:"workstation"`
}

type LocalState struct {
	// Enable/disable overriding the policy-auth-concurrent under config system global. Valid values: `enable`, `disable`.
	AuthConcurrentOverride pulumi.StringPtrInput
	// Maximum number of concurrent logins permitted from the same user.
	AuthConcurrentValue pulumi.IntPtrInput
	// Time in minutes before the authentication timeout for a user is reached.
	Authtimeout pulumi.IntPtrInput
	// Two-factor recipient's email address.
	EmailTo pulumi.StringPtrInput
	// Two-factor recipient's FortiToken serial number.
	Fortitoken pulumi.StringPtrInput
	// User ID.
	Fosid pulumi.IntPtrInput
	// Name of LDAP server with which the user must authenticate.
	LdapServer pulumi.StringPtrInput
	// User name.
	Name pulumi.StringPtrInput
	// User's password.
	Passwd pulumi.StringPtrInput
	// Password policy to apply to this user, as defined in config user password-policy.
	PasswdPolicy pulumi.StringPtrInput
	// Time of the last password update.
	PasswdTime pulumi.StringPtrInput
	// IKEv2 Postquantum Preshared Key Identity.
	PpkIdentity pulumi.StringPtrInput
	// IKEv2 Postquantum Preshared Key (ASCII string or hexadecimal encoded with a leading 0x).
	PpkSecret pulumi.StringPtrInput
	// Name of RADIUS server with which the user must authenticate.
	RadiusServer pulumi.StringPtrInput
	// Two-factor recipient's SMS server.
	SmsCustomServer pulumi.StringPtrInput
	// Two-factor recipient's mobile phone number.
	SmsPhone pulumi.StringPtrInput
	// Send SMS through FortiGuard or other external server. Valid values: `fortiguard`, `custom`.
	SmsServer pulumi.StringPtrInput
	// Enable/disable allowing the local user to authenticate with the FortiGate unit. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Name of TACACS+ server with which the user must authenticate.
	TacacsServer pulumi.StringPtrInput
	// Enable/disable two-factor authentication.
	TwoFactor pulumi.StringPtrInput
	// Authentication method by FortiToken Cloud. Valid values: `fortitoken`, `email`, `sms`.
	TwoFactorAuthentication pulumi.StringPtrInput
	// Notification method for user activation by FortiToken Cloud. Valid values: `email`, `sms`.
	TwoFactorNotification pulumi.StringPtrInput
	// Authentication method. Valid values: `password`, `radius`, `tacacs+`, `ldap`.
	Type pulumi.StringPtrInput
	// Enable/disable case sensitivity when performing username matching (uppercase and lowercase letters are treated either as distinct or equivalent). Valid values: `enable`, `disable`.
	UsernameCaseInsensitivity pulumi.StringPtrInput
	// Enable/disable case sensitivity when performing username matching (uppercase and lowercase letters are treated either as distinct or equivalent). Valid values: `disable`, `enable`.
	UsernameCaseSensitivity pulumi.StringPtrInput
	// Enable/disable case and accent sensitivity when performing username matching (accents are stripped and case is ignored when disabled). Valid values: `disable`, `enable`.
	UsernameSensitivity pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Name of the remote user workstation, if you want to limit the user to authenticate only from a particular workstation.
	Workstation pulumi.StringPtrInput
}

func (LocalState) ElementType() reflect.Type {
	return reflect.TypeOf((*localState)(nil)).Elem()
}

type localArgs struct {
	// Enable/disable overriding the policy-auth-concurrent under config system global. Valid values: `enable`, `disable`.
	AuthConcurrentOverride *string `pulumi:"authConcurrentOverride"`
	// Maximum number of concurrent logins permitted from the same user.
	AuthConcurrentValue *int `pulumi:"authConcurrentValue"`
	// Time in minutes before the authentication timeout for a user is reached.
	Authtimeout *int `pulumi:"authtimeout"`
	// Two-factor recipient's email address.
	EmailTo *string `pulumi:"emailTo"`
	// Two-factor recipient's FortiToken serial number.
	Fortitoken *string `pulumi:"fortitoken"`
	// User ID.
	Fosid *int `pulumi:"fosid"`
	// Name of LDAP server with which the user must authenticate.
	LdapServer *string `pulumi:"ldapServer"`
	// User name.
	Name *string `pulumi:"name"`
	// User's password.
	Passwd *string `pulumi:"passwd"`
	// Password policy to apply to this user, as defined in config user password-policy.
	PasswdPolicy *string `pulumi:"passwdPolicy"`
	// Time of the last password update.
	PasswdTime *string `pulumi:"passwdTime"`
	// IKEv2 Postquantum Preshared Key Identity.
	PpkIdentity *string `pulumi:"ppkIdentity"`
	// IKEv2 Postquantum Preshared Key (ASCII string or hexadecimal encoded with a leading 0x).
	PpkSecret *string `pulumi:"ppkSecret"`
	// Name of RADIUS server with which the user must authenticate.
	RadiusServer *string `pulumi:"radiusServer"`
	// Two-factor recipient's SMS server.
	SmsCustomServer *string `pulumi:"smsCustomServer"`
	// Two-factor recipient's mobile phone number.
	SmsPhone *string `pulumi:"smsPhone"`
	// Send SMS through FortiGuard or other external server. Valid values: `fortiguard`, `custom`.
	SmsServer *string `pulumi:"smsServer"`
	// Enable/disable allowing the local user to authenticate with the FortiGate unit. Valid values: `enable`, `disable`.
	Status string `pulumi:"status"`
	// Name of TACACS+ server with which the user must authenticate.
	TacacsServer *string `pulumi:"tacacsServer"`
	// Enable/disable two-factor authentication.
	TwoFactor *string `pulumi:"twoFactor"`
	// Authentication method by FortiToken Cloud. Valid values: `fortitoken`, `email`, `sms`.
	TwoFactorAuthentication *string `pulumi:"twoFactorAuthentication"`
	// Notification method for user activation by FortiToken Cloud. Valid values: `email`, `sms`.
	TwoFactorNotification *string `pulumi:"twoFactorNotification"`
	// Authentication method. Valid values: `password`, `radius`, `tacacs+`, `ldap`.
	Type string `pulumi:"type"`
	// Enable/disable case sensitivity when performing username matching (uppercase and lowercase letters are treated either as distinct or equivalent). Valid values: `enable`, `disable`.
	UsernameCaseInsensitivity *string `pulumi:"usernameCaseInsensitivity"`
	// Enable/disable case sensitivity when performing username matching (uppercase and lowercase letters are treated either as distinct or equivalent). Valid values: `disable`, `enable`.
	UsernameCaseSensitivity *string `pulumi:"usernameCaseSensitivity"`
	// Enable/disable case and accent sensitivity when performing username matching (accents are stripped and case is ignored when disabled). Valid values: `disable`, `enable`.
	UsernameSensitivity *string `pulumi:"usernameSensitivity"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Name of the remote user workstation, if you want to limit the user to authenticate only from a particular workstation.
	Workstation *string `pulumi:"workstation"`
}

// The set of arguments for constructing a Local resource.
type LocalArgs struct {
	// Enable/disable overriding the policy-auth-concurrent under config system global. Valid values: `enable`, `disable`.
	AuthConcurrentOverride pulumi.StringPtrInput
	// Maximum number of concurrent logins permitted from the same user.
	AuthConcurrentValue pulumi.IntPtrInput
	// Time in minutes before the authentication timeout for a user is reached.
	Authtimeout pulumi.IntPtrInput
	// Two-factor recipient's email address.
	EmailTo pulumi.StringPtrInput
	// Two-factor recipient's FortiToken serial number.
	Fortitoken pulumi.StringPtrInput
	// User ID.
	Fosid pulumi.IntPtrInput
	// Name of LDAP server with which the user must authenticate.
	LdapServer pulumi.StringPtrInput
	// User name.
	Name pulumi.StringPtrInput
	// User's password.
	Passwd pulumi.StringPtrInput
	// Password policy to apply to this user, as defined in config user password-policy.
	PasswdPolicy pulumi.StringPtrInput
	// Time of the last password update.
	PasswdTime pulumi.StringPtrInput
	// IKEv2 Postquantum Preshared Key Identity.
	PpkIdentity pulumi.StringPtrInput
	// IKEv2 Postquantum Preshared Key (ASCII string or hexadecimal encoded with a leading 0x).
	PpkSecret pulumi.StringPtrInput
	// Name of RADIUS server with which the user must authenticate.
	RadiusServer pulumi.StringPtrInput
	// Two-factor recipient's SMS server.
	SmsCustomServer pulumi.StringPtrInput
	// Two-factor recipient's mobile phone number.
	SmsPhone pulumi.StringPtrInput
	// Send SMS through FortiGuard or other external server. Valid values: `fortiguard`, `custom`.
	SmsServer pulumi.StringPtrInput
	// Enable/disable allowing the local user to authenticate with the FortiGate unit. Valid values: `enable`, `disable`.
	Status pulumi.StringInput
	// Name of TACACS+ server with which the user must authenticate.
	TacacsServer pulumi.StringPtrInput
	// Enable/disable two-factor authentication.
	TwoFactor pulumi.StringPtrInput
	// Authentication method by FortiToken Cloud. Valid values: `fortitoken`, `email`, `sms`.
	TwoFactorAuthentication pulumi.StringPtrInput
	// Notification method for user activation by FortiToken Cloud. Valid values: `email`, `sms`.
	TwoFactorNotification pulumi.StringPtrInput
	// Authentication method. Valid values: `password`, `radius`, `tacacs+`, `ldap`.
	Type pulumi.StringInput
	// Enable/disable case sensitivity when performing username matching (uppercase and lowercase letters are treated either as distinct or equivalent). Valid values: `enable`, `disable`.
	UsernameCaseInsensitivity pulumi.StringPtrInput
	// Enable/disable case sensitivity when performing username matching (uppercase and lowercase letters are treated either as distinct or equivalent). Valid values: `disable`, `enable`.
	UsernameCaseSensitivity pulumi.StringPtrInput
	// Enable/disable case and accent sensitivity when performing username matching (accents are stripped and case is ignored when disabled). Valid values: `disable`, `enable`.
	UsernameSensitivity pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Name of the remote user workstation, if you want to limit the user to authenticate only from a particular workstation.
	Workstation pulumi.StringPtrInput
}

func (LocalArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*localArgs)(nil)).Elem()
}

type LocalInput interface {
	pulumi.Input

	ToLocalOutput() LocalOutput
	ToLocalOutputWithContext(ctx context.Context) LocalOutput
}

func (*Local) ElementType() reflect.Type {
	return reflect.TypeOf((**Local)(nil)).Elem()
}

func (i *Local) ToLocalOutput() LocalOutput {
	return i.ToLocalOutputWithContext(context.Background())
}

func (i *Local) ToLocalOutputWithContext(ctx context.Context) LocalOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalOutput)
}

// LocalArrayInput is an input type that accepts LocalArray and LocalArrayOutput values.
// You can construct a concrete instance of `LocalArrayInput` via:
//
//	LocalArray{ LocalArgs{...} }
type LocalArrayInput interface {
	pulumi.Input

	ToLocalArrayOutput() LocalArrayOutput
	ToLocalArrayOutputWithContext(context.Context) LocalArrayOutput
}

type LocalArray []LocalInput

func (LocalArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Local)(nil)).Elem()
}

func (i LocalArray) ToLocalArrayOutput() LocalArrayOutput {
	return i.ToLocalArrayOutputWithContext(context.Background())
}

func (i LocalArray) ToLocalArrayOutputWithContext(ctx context.Context) LocalArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalArrayOutput)
}

// LocalMapInput is an input type that accepts LocalMap and LocalMapOutput values.
// You can construct a concrete instance of `LocalMapInput` via:
//
//	LocalMap{ "key": LocalArgs{...} }
type LocalMapInput interface {
	pulumi.Input

	ToLocalMapOutput() LocalMapOutput
	ToLocalMapOutputWithContext(context.Context) LocalMapOutput
}

type LocalMap map[string]LocalInput

func (LocalMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Local)(nil)).Elem()
}

func (i LocalMap) ToLocalMapOutput() LocalMapOutput {
	return i.ToLocalMapOutputWithContext(context.Background())
}

func (i LocalMap) ToLocalMapOutputWithContext(ctx context.Context) LocalMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalMapOutput)
}

type LocalOutput struct{ *pulumi.OutputState }

func (LocalOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Local)(nil)).Elem()
}

func (o LocalOutput) ToLocalOutput() LocalOutput {
	return o
}

func (o LocalOutput) ToLocalOutputWithContext(ctx context.Context) LocalOutput {
	return o
}

// Enable/disable overriding the policy-auth-concurrent under config system global. Valid values: `enable`, `disable`.
func (o LocalOutput) AuthConcurrentOverride() pulumi.StringOutput {
	return o.ApplyT(func(v *Local) pulumi.StringOutput { return v.AuthConcurrentOverride }).(pulumi.StringOutput)
}

// Maximum number of concurrent logins permitted from the same user.
func (o LocalOutput) AuthConcurrentValue() pulumi.IntOutput {
	return o.ApplyT(func(v *Local) pulumi.IntOutput { return v.AuthConcurrentValue }).(pulumi.IntOutput)
}

// Time in minutes before the authentication timeout for a user is reached.
func (o LocalOutput) Authtimeout() pulumi.IntOutput {
	return o.ApplyT(func(v *Local) pulumi.IntOutput { return v.Authtimeout }).(pulumi.IntOutput)
}

// Two-factor recipient's email address.
func (o LocalOutput) EmailTo() pulumi.StringOutput {
	return o.ApplyT(func(v *Local) pulumi.StringOutput { return v.EmailTo }).(pulumi.StringOutput)
}

// Two-factor recipient's FortiToken serial number.
func (o LocalOutput) Fortitoken() pulumi.StringOutput {
	return o.ApplyT(func(v *Local) pulumi.StringOutput { return v.Fortitoken }).(pulumi.StringOutput)
}

// User ID.
func (o LocalOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *Local) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

// Name of LDAP server with which the user must authenticate.
func (o LocalOutput) LdapServer() pulumi.StringOutput {
	return o.ApplyT(func(v *Local) pulumi.StringOutput { return v.LdapServer }).(pulumi.StringOutput)
}

// User name.
func (o LocalOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Local) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// User's password.
func (o LocalOutput) Passwd() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Local) pulumi.StringPtrOutput { return v.Passwd }).(pulumi.StringPtrOutput)
}

// Password policy to apply to this user, as defined in config user password-policy.
func (o LocalOutput) PasswdPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v *Local) pulumi.StringOutput { return v.PasswdPolicy }).(pulumi.StringOutput)
}

// Time of the last password update.
func (o LocalOutput) PasswdTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Local) pulumi.StringOutput { return v.PasswdTime }).(pulumi.StringOutput)
}

// IKEv2 Postquantum Preshared Key Identity.
func (o LocalOutput) PpkIdentity() pulumi.StringOutput {
	return o.ApplyT(func(v *Local) pulumi.StringOutput { return v.PpkIdentity }).(pulumi.StringOutput)
}

// IKEv2 Postquantum Preshared Key (ASCII string or hexadecimal encoded with a leading 0x).
func (o LocalOutput) PpkSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Local) pulumi.StringPtrOutput { return v.PpkSecret }).(pulumi.StringPtrOutput)
}

// Name of RADIUS server with which the user must authenticate.
func (o LocalOutput) RadiusServer() pulumi.StringOutput {
	return o.ApplyT(func(v *Local) pulumi.StringOutput { return v.RadiusServer }).(pulumi.StringOutput)
}

// Two-factor recipient's SMS server.
func (o LocalOutput) SmsCustomServer() pulumi.StringOutput {
	return o.ApplyT(func(v *Local) pulumi.StringOutput { return v.SmsCustomServer }).(pulumi.StringOutput)
}

// Two-factor recipient's mobile phone number.
func (o LocalOutput) SmsPhone() pulumi.StringOutput {
	return o.ApplyT(func(v *Local) pulumi.StringOutput { return v.SmsPhone }).(pulumi.StringOutput)
}

// Send SMS through FortiGuard or other external server. Valid values: `fortiguard`, `custom`.
func (o LocalOutput) SmsServer() pulumi.StringOutput {
	return o.ApplyT(func(v *Local) pulumi.StringOutput { return v.SmsServer }).(pulumi.StringOutput)
}

// Enable/disable allowing the local user to authenticate with the FortiGate unit. Valid values: `enable`, `disable`.
func (o LocalOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Local) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Name of TACACS+ server with which the user must authenticate.
func (o LocalOutput) TacacsServer() pulumi.StringOutput {
	return o.ApplyT(func(v *Local) pulumi.StringOutput { return v.TacacsServer }).(pulumi.StringOutput)
}

// Enable/disable two-factor authentication.
func (o LocalOutput) TwoFactor() pulumi.StringOutput {
	return o.ApplyT(func(v *Local) pulumi.StringOutput { return v.TwoFactor }).(pulumi.StringOutput)
}

// Authentication method by FortiToken Cloud. Valid values: `fortitoken`, `email`, `sms`.
func (o LocalOutput) TwoFactorAuthentication() pulumi.StringOutput {
	return o.ApplyT(func(v *Local) pulumi.StringOutput { return v.TwoFactorAuthentication }).(pulumi.StringOutput)
}

// Notification method for user activation by FortiToken Cloud. Valid values: `email`, `sms`.
func (o LocalOutput) TwoFactorNotification() pulumi.StringOutput {
	return o.ApplyT(func(v *Local) pulumi.StringOutput { return v.TwoFactorNotification }).(pulumi.StringOutput)
}

// Authentication method. Valid values: `password`, `radius`, `tacacs+`, `ldap`.
func (o LocalOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Local) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Enable/disable case sensitivity when performing username matching (uppercase and lowercase letters are treated either as distinct or equivalent). Valid values: `enable`, `disable`.
func (o LocalOutput) UsernameCaseInsensitivity() pulumi.StringOutput {
	return o.ApplyT(func(v *Local) pulumi.StringOutput { return v.UsernameCaseInsensitivity }).(pulumi.StringOutput)
}

// Enable/disable case sensitivity when performing username matching (uppercase and lowercase letters are treated either as distinct or equivalent). Valid values: `disable`, `enable`.
func (o LocalOutput) UsernameCaseSensitivity() pulumi.StringOutput {
	return o.ApplyT(func(v *Local) pulumi.StringOutput { return v.UsernameCaseSensitivity }).(pulumi.StringOutput)
}

// Enable/disable case and accent sensitivity when performing username matching (accents are stripped and case is ignored when disabled). Valid values: `disable`, `enable`.
func (o LocalOutput) UsernameSensitivity() pulumi.StringOutput {
	return o.ApplyT(func(v *Local) pulumi.StringOutput { return v.UsernameSensitivity }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o LocalOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Local) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// Name of the remote user workstation, if you want to limit the user to authenticate only from a particular workstation.
func (o LocalOutput) Workstation() pulumi.StringOutput {
	return o.ApplyT(func(v *Local) pulumi.StringOutput { return v.Workstation }).(pulumi.StringOutput)
}

type LocalArrayOutput struct{ *pulumi.OutputState }

func (LocalArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Local)(nil)).Elem()
}

func (o LocalArrayOutput) ToLocalArrayOutput() LocalArrayOutput {
	return o
}

func (o LocalArrayOutput) ToLocalArrayOutputWithContext(ctx context.Context) LocalArrayOutput {
	return o
}

func (o LocalArrayOutput) Index(i pulumi.IntInput) LocalOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Local {
		return vs[0].([]*Local)[vs[1].(int)]
	}).(LocalOutput)
}

type LocalMapOutput struct{ *pulumi.OutputState }

func (LocalMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Local)(nil)).Elem()
}

func (o LocalMapOutput) ToLocalMapOutput() LocalMapOutput {
	return o
}

func (o LocalMapOutput) ToLocalMapOutputWithContext(ctx context.Context) LocalMapOutput {
	return o
}

func (o LocalMapOutput) MapIndex(k pulumi.StringInput) LocalOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Local {
		return vs[0].(map[string]*Local)[vs[1].(string)]
	}).(LocalOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LocalInput)(nil)).Elem(), &Local{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalArrayInput)(nil)).Elem(), LocalArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalMapInput)(nil)).Elem(), LocalMap{})
	pulumi.RegisterOutputType(LocalOutput{})
	pulumi.RegisterOutputType(LocalArrayOutput{})
	pulumi.RegisterOutputType(LocalMapOutput{})
}
