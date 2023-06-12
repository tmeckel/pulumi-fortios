// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package wanopt

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure WAN optimization profiles.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/wanopt"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := wanopt.NewProfile(ctx, "trname", &wanopt.ProfileArgs{
//				Cifs: &wanopt.ProfileCifsArgs{
//					ByteCaching:    pulumi.String("enable"),
//					LogTraffic:     pulumi.String("enable"),
//					Port:           pulumi.Int(445),
//					PreferChunking: pulumi.String("fix"),
//					SecureTunnel:   pulumi.String("disable"),
//					Status:         pulumi.String("disable"),
//					TunnelSharing:  pulumi.String("private"),
//				},
//				Comments: pulumi.String("test"),
//				Ftp: &wanopt.ProfileFtpArgs{
//					ByteCaching:    pulumi.String("enable"),
//					LogTraffic:     pulumi.String("enable"),
//					Port:           pulumi.Int(21),
//					PreferChunking: pulumi.String("fix"),
//					SecureTunnel:   pulumi.String("disable"),
//					Status:         pulumi.String("disable"),
//					TunnelSharing:  pulumi.String("private"),
//				},
//				Http: &wanopt.ProfileHttpArgs{
//					ByteCaching:        pulumi.String("enable"),
//					LogTraffic:         pulumi.String("enable"),
//					Port:               pulumi.Int(80),
//					PreferChunking:     pulumi.String("fix"),
//					SecureTunnel:       pulumi.String("disable"),
//					Ssl:                pulumi.String("disable"),
//					SslPort:            pulumi.Int(443),
//					Status:             pulumi.String("disable"),
//					TunnelNonHttp:      pulumi.String("disable"),
//					TunnelSharing:      pulumi.String("private"),
//					UnknownHttpVersion: pulumi.String("tunnel"),
//				},
//				Mapi: &wanopt.ProfileMapiArgs{
//					ByteCaching:   pulumi.String("enable"),
//					LogTraffic:    pulumi.String("enable"),
//					Port:          pulumi.Int(135),
//					SecureTunnel:  pulumi.String("disable"),
//					Status:        pulumi.String("disable"),
//					TunnelSharing: pulumi.String("private"),
//				},
//				Tcp: &wanopt.ProfileTcpArgs{
//					ByteCaching:    pulumi.String("disable"),
//					ByteCachingOpt: pulumi.String("mem-only"),
//					LogTraffic:     pulumi.String("enable"),
//					Port:           pulumi.String("1-65535"),
//					SecureTunnel:   pulumi.String("disable"),
//					Ssl:            pulumi.String("disable"),
//					SslPort:        pulumi.Int(443),
//					Status:         pulumi.String("disable"),
//					TunnelSharing:  pulumi.String("private"),
//				},
//				Transparent: pulumi.String("enable"),
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
// # Wanopt Profile can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:wanopt/profile:Profile labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:wanopt/profile:Profile labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type Profile struct {
	pulumi.CustomResourceState

	// Optionally add an authentication group to restrict access to the WAN Optimization tunnel to peers in the authentication group.
	AuthGroup pulumi.StringOutput `pulumi:"authGroup"`
	// Enable/disable CIFS (Windows sharing) WAN Optimization and configure CIFS WAN Optimization features. The structure of `cifs` block is documented below.
	Cifs ProfileCifsOutput `pulumi:"cifs"`
	// Comment.
	Comments pulumi.StringPtrOutput `pulumi:"comments"`
	// Enable/disable FTP WAN Optimization and configure FTP WAN Optimization features. The structure of `ftp` block is documented below.
	Ftp ProfileFtpOutput `pulumi:"ftp"`
	// Enable/disable HTTP WAN Optimization and configure HTTP WAN Optimization features. The structure of `http` block is documented below.
	Http ProfileHttpOutput `pulumi:"http"`
	// Enable/disable MAPI email WAN Optimization and configure MAPI WAN Optimization features. The structure of `mapi` block is documented below.
	Mapi ProfileMapiOutput `pulumi:"mapi"`
	// Profile name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Enable/disable TCP WAN Optimization and configure TCP WAN Optimization features. The structure of `tcp` block is documented below.
	Tcp ProfileTcpOutput `pulumi:"tcp"`
	// Enable/disable transparent mode. Valid values: `enable`, `disable`.
	Transparent pulumi.StringOutput `pulumi:"transparent"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewProfile registers a new resource with the given unique name, arguments, and options.
func NewProfile(ctx *pulumi.Context,
	name string, args *ProfileArgs, opts ...pulumi.ResourceOption) (*Profile, error) {
	if args == nil {
		args = &ProfileArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource Profile
	err := ctx.RegisterResource("fortios:wanopt/profile:Profile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProfile gets an existing Profile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProfileState, opts ...pulumi.ResourceOption) (*Profile, error) {
	var resource Profile
	err := ctx.ReadResource("fortios:wanopt/profile:Profile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Profile resources.
type profileState struct {
	// Optionally add an authentication group to restrict access to the WAN Optimization tunnel to peers in the authentication group.
	AuthGroup *string `pulumi:"authGroup"`
	// Enable/disable CIFS (Windows sharing) WAN Optimization and configure CIFS WAN Optimization features. The structure of `cifs` block is documented below.
	Cifs *ProfileCifs `pulumi:"cifs"`
	// Comment.
	Comments *string `pulumi:"comments"`
	// Enable/disable FTP WAN Optimization and configure FTP WAN Optimization features. The structure of `ftp` block is documented below.
	Ftp *ProfileFtp `pulumi:"ftp"`
	// Enable/disable HTTP WAN Optimization and configure HTTP WAN Optimization features. The structure of `http` block is documented below.
	Http *ProfileHttp `pulumi:"http"`
	// Enable/disable MAPI email WAN Optimization and configure MAPI WAN Optimization features. The structure of `mapi` block is documented below.
	Mapi *ProfileMapi `pulumi:"mapi"`
	// Profile name.
	Name *string `pulumi:"name"`
	// Enable/disable TCP WAN Optimization and configure TCP WAN Optimization features. The structure of `tcp` block is documented below.
	Tcp *ProfileTcp `pulumi:"tcp"`
	// Enable/disable transparent mode. Valid values: `enable`, `disable`.
	Transparent *string `pulumi:"transparent"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type ProfileState struct {
	// Optionally add an authentication group to restrict access to the WAN Optimization tunnel to peers in the authentication group.
	AuthGroup pulumi.StringPtrInput
	// Enable/disable CIFS (Windows sharing) WAN Optimization and configure CIFS WAN Optimization features. The structure of `cifs` block is documented below.
	Cifs ProfileCifsPtrInput
	// Comment.
	Comments pulumi.StringPtrInput
	// Enable/disable FTP WAN Optimization and configure FTP WAN Optimization features. The structure of `ftp` block is documented below.
	Ftp ProfileFtpPtrInput
	// Enable/disable HTTP WAN Optimization and configure HTTP WAN Optimization features. The structure of `http` block is documented below.
	Http ProfileHttpPtrInput
	// Enable/disable MAPI email WAN Optimization and configure MAPI WAN Optimization features. The structure of `mapi` block is documented below.
	Mapi ProfileMapiPtrInput
	// Profile name.
	Name pulumi.StringPtrInput
	// Enable/disable TCP WAN Optimization and configure TCP WAN Optimization features. The structure of `tcp` block is documented below.
	Tcp ProfileTcpPtrInput
	// Enable/disable transparent mode. Valid values: `enable`, `disable`.
	Transparent pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (ProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*profileState)(nil)).Elem()
}

type profileArgs struct {
	// Optionally add an authentication group to restrict access to the WAN Optimization tunnel to peers in the authentication group.
	AuthGroup *string `pulumi:"authGroup"`
	// Enable/disable CIFS (Windows sharing) WAN Optimization and configure CIFS WAN Optimization features. The structure of `cifs` block is documented below.
	Cifs *ProfileCifs `pulumi:"cifs"`
	// Comment.
	Comments *string `pulumi:"comments"`
	// Enable/disable FTP WAN Optimization and configure FTP WAN Optimization features. The structure of `ftp` block is documented below.
	Ftp *ProfileFtp `pulumi:"ftp"`
	// Enable/disable HTTP WAN Optimization and configure HTTP WAN Optimization features. The structure of `http` block is documented below.
	Http *ProfileHttp `pulumi:"http"`
	// Enable/disable MAPI email WAN Optimization and configure MAPI WAN Optimization features. The structure of `mapi` block is documented below.
	Mapi *ProfileMapi `pulumi:"mapi"`
	// Profile name.
	Name *string `pulumi:"name"`
	// Enable/disable TCP WAN Optimization and configure TCP WAN Optimization features. The structure of `tcp` block is documented below.
	Tcp *ProfileTcp `pulumi:"tcp"`
	// Enable/disable transparent mode. Valid values: `enable`, `disable`.
	Transparent *string `pulumi:"transparent"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Profile resource.
type ProfileArgs struct {
	// Optionally add an authentication group to restrict access to the WAN Optimization tunnel to peers in the authentication group.
	AuthGroup pulumi.StringPtrInput
	// Enable/disable CIFS (Windows sharing) WAN Optimization and configure CIFS WAN Optimization features. The structure of `cifs` block is documented below.
	Cifs ProfileCifsPtrInput
	// Comment.
	Comments pulumi.StringPtrInput
	// Enable/disable FTP WAN Optimization and configure FTP WAN Optimization features. The structure of `ftp` block is documented below.
	Ftp ProfileFtpPtrInput
	// Enable/disable HTTP WAN Optimization and configure HTTP WAN Optimization features. The structure of `http` block is documented below.
	Http ProfileHttpPtrInput
	// Enable/disable MAPI email WAN Optimization and configure MAPI WAN Optimization features. The structure of `mapi` block is documented below.
	Mapi ProfileMapiPtrInput
	// Profile name.
	Name pulumi.StringPtrInput
	// Enable/disable TCP WAN Optimization and configure TCP WAN Optimization features. The structure of `tcp` block is documented below.
	Tcp ProfileTcpPtrInput
	// Enable/disable transparent mode. Valid values: `enable`, `disable`.
	Transparent pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (ProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*profileArgs)(nil)).Elem()
}

type ProfileInput interface {
	pulumi.Input

	ToProfileOutput() ProfileOutput
	ToProfileOutputWithContext(ctx context.Context) ProfileOutput
}

func (*Profile) ElementType() reflect.Type {
	return reflect.TypeOf((**Profile)(nil)).Elem()
}

func (i *Profile) ToProfileOutput() ProfileOutput {
	return i.ToProfileOutputWithContext(context.Background())
}

func (i *Profile) ToProfileOutputWithContext(ctx context.Context) ProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileOutput)
}

// ProfileArrayInput is an input type that accepts ProfileArray and ProfileArrayOutput values.
// You can construct a concrete instance of `ProfileArrayInput` via:
//
//	ProfileArray{ ProfileArgs{...} }
type ProfileArrayInput interface {
	pulumi.Input

	ToProfileArrayOutput() ProfileArrayOutput
	ToProfileArrayOutputWithContext(context.Context) ProfileArrayOutput
}

type ProfileArray []ProfileInput

func (ProfileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Profile)(nil)).Elem()
}

func (i ProfileArray) ToProfileArrayOutput() ProfileArrayOutput {
	return i.ToProfileArrayOutputWithContext(context.Background())
}

func (i ProfileArray) ToProfileArrayOutputWithContext(ctx context.Context) ProfileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileArrayOutput)
}

// ProfileMapInput is an input type that accepts ProfileMap and ProfileMapOutput values.
// You can construct a concrete instance of `ProfileMapInput` via:
//
//	ProfileMap{ "key": ProfileArgs{...} }
type ProfileMapInput interface {
	pulumi.Input

	ToProfileMapOutput() ProfileMapOutput
	ToProfileMapOutputWithContext(context.Context) ProfileMapOutput
}

type ProfileMap map[string]ProfileInput

func (ProfileMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Profile)(nil)).Elem()
}

func (i ProfileMap) ToProfileMapOutput() ProfileMapOutput {
	return i.ToProfileMapOutputWithContext(context.Background())
}

func (i ProfileMap) ToProfileMapOutputWithContext(ctx context.Context) ProfileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileMapOutput)
}

type ProfileOutput struct{ *pulumi.OutputState }

func (ProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Profile)(nil)).Elem()
}

func (o ProfileOutput) ToProfileOutput() ProfileOutput {
	return o
}

func (o ProfileOutput) ToProfileOutputWithContext(ctx context.Context) ProfileOutput {
	return o
}

// Optionally add an authentication group to restrict access to the WAN Optimization tunnel to peers in the authentication group.
func (o ProfileOutput) AuthGroup() pulumi.StringOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringOutput { return v.AuthGroup }).(pulumi.StringOutput)
}

// Enable/disable CIFS (Windows sharing) WAN Optimization and configure CIFS WAN Optimization features. The structure of `cifs` block is documented below.
func (o ProfileOutput) Cifs() ProfileCifsOutput {
	return o.ApplyT(func(v *Profile) ProfileCifsOutput { return v.Cifs }).(ProfileCifsOutput)
}

// Comment.
func (o ProfileOutput) Comments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringPtrOutput { return v.Comments }).(pulumi.StringPtrOutput)
}

// Enable/disable FTP WAN Optimization and configure FTP WAN Optimization features. The structure of `ftp` block is documented below.
func (o ProfileOutput) Ftp() ProfileFtpOutput {
	return o.ApplyT(func(v *Profile) ProfileFtpOutput { return v.Ftp }).(ProfileFtpOutput)
}

// Enable/disable HTTP WAN Optimization and configure HTTP WAN Optimization features. The structure of `http` block is documented below.
func (o ProfileOutput) Http() ProfileHttpOutput {
	return o.ApplyT(func(v *Profile) ProfileHttpOutput { return v.Http }).(ProfileHttpOutput)
}

// Enable/disable MAPI email WAN Optimization and configure MAPI WAN Optimization features. The structure of `mapi` block is documented below.
func (o ProfileOutput) Mapi() ProfileMapiOutput {
	return o.ApplyT(func(v *Profile) ProfileMapiOutput { return v.Mapi }).(ProfileMapiOutput)
}

// Profile name.
func (o ProfileOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Enable/disable TCP WAN Optimization and configure TCP WAN Optimization features. The structure of `tcp` block is documented below.
func (o ProfileOutput) Tcp() ProfileTcpOutput {
	return o.ApplyT(func(v *Profile) ProfileTcpOutput { return v.Tcp }).(ProfileTcpOutput)
}

// Enable/disable transparent mode. Valid values: `enable`, `disable`.
func (o ProfileOutput) Transparent() pulumi.StringOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringOutput { return v.Transparent }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o ProfileOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type ProfileArrayOutput struct{ *pulumi.OutputState }

func (ProfileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Profile)(nil)).Elem()
}

func (o ProfileArrayOutput) ToProfileArrayOutput() ProfileArrayOutput {
	return o
}

func (o ProfileArrayOutput) ToProfileArrayOutputWithContext(ctx context.Context) ProfileArrayOutput {
	return o
}

func (o ProfileArrayOutput) Index(i pulumi.IntInput) ProfileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Profile {
		return vs[0].([]*Profile)[vs[1].(int)]
	}).(ProfileOutput)
}

type ProfileMapOutput struct{ *pulumi.OutputState }

func (ProfileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Profile)(nil)).Elem()
}

func (o ProfileMapOutput) ToProfileMapOutput() ProfileMapOutput {
	return o
}

func (o ProfileMapOutput) ToProfileMapOutputWithContext(ctx context.Context) ProfileMapOutput {
	return o
}

func (o ProfileMapOutput) MapIndex(k pulumi.StringInput) ProfileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Profile {
		return vs[0].(map[string]*Profile)[vs[1].(string)]
	}).(ProfileOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProfileInput)(nil)).Elem(), &Profile{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProfileArrayInput)(nil)).Elem(), ProfileArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProfileMapInput)(nil)).Elem(), ProfileMap{})
	pulumi.RegisterOutputType(ProfileOutput{})
	pulumi.RegisterOutputType(ProfileArrayOutput{})
	pulumi.RegisterOutputType(ProfileMapOutput{})
}