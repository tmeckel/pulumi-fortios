// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firewall

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios firewall address6
func LookupAddress6(ctx *pulumi.Context, args *LookupAddress6Args, opts ...pulumi.InvokeOption) (*LookupAddress6Result, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupAddress6Result
	err := ctx.Invoke("fortios:firewall/getAddress6:getAddress6", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAddress6.
type LookupAddress6Args struct {
	// Specify the name of the desired firewall address6.
	Name string `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getAddress6.
type LookupAddress6Result struct {
	// Minimal TTL of individual IPv6 addresses in FQDN cache.
	CacheTtl int `pulumi:"cacheTtl"`
	// Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
	Color int `pulumi:"color"`
	// Comment.
	Comment string `pulumi:"comment"`
	// IPv6 addresses associated to a specific country.
	Country string `pulumi:"country"`
	// Final IP address (inclusive) in the range for the address (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).
	EndIp string `pulumi:"endIp"`
	// Last MAC address in the range.
	EndMac string `pulumi:"endMac"`
	// Endpoint group name.
	EpgName string `pulumi:"epgName"`
	// Security Fabric global object setting.
	FabricObject string `pulumi:"fabricObject"`
	// Fully qualified domain name.
	Fqdn string `pulumi:"fqdn"`
	// Host Address.
	Host string `pulumi:"host"`
	// Host type.
	HostType string `pulumi:"hostType"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// IPv6 address prefix (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx).
	Ip6 string `pulumi:"ip6"`
	// IP address list. The structure of `list` block is documented below.
	Lists []GetAddress6List `pulumi:"lists"`
	// MAC address ranges <start>[-<end>] separated by space.
	Macaddrs []GetAddress6Macaddr `pulumi:"macaddrs"`
	// Name.
	Name string `pulumi:"name"`
	// Object ID for NSX.
	ObjId string `pulumi:"objId"`
	// SDN.
	Sdn string `pulumi:"sdn"`
	// SDN Tag.
	SdnTag string `pulumi:"sdnTag"`
	// First IP address (inclusive) in the range for the address (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).
	StartIp string `pulumi:"startIp"`
	// First MAC address in the range.
	StartMac string `pulumi:"startMac"`
	// IPv6 subnet segments. The structure of `subnetSegment` block is documented below.
	SubnetSegments []GetAddress6SubnetSegment `pulumi:"subnetSegments"`
	// Config object tagging The structure of `tagging` block is documented below.
	Taggings []GetAddress6Tagging `pulumi:"taggings"`
	// IPv6 address template.
	Template string `pulumi:"template"`
	// Tenant.
	Tenant string `pulumi:"tenant"`
	// Subnet segment type.
	Type string `pulumi:"type"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid      string  `pulumi:"uuid"`
	Vdomparam *string `pulumi:"vdomparam"`
	// Enable/disable the visibility of the object in the GUI.
	Visibility string `pulumi:"visibility"`
}

func LookupAddress6Output(ctx *pulumi.Context, args LookupAddress6OutputArgs, opts ...pulumi.InvokeOption) LookupAddress6ResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAddress6Result, error) {
			args := v.(LookupAddress6Args)
			r, err := LookupAddress6(ctx, &args, opts...)
			var s LookupAddress6Result
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAddress6ResultOutput)
}

// A collection of arguments for invoking getAddress6.
type LookupAddress6OutputArgs struct {
	// Specify the name of the desired firewall address6.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupAddress6OutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAddress6Args)(nil)).Elem()
}

// A collection of values returned by getAddress6.
type LookupAddress6ResultOutput struct{ *pulumi.OutputState }

func (LookupAddress6ResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAddress6Result)(nil)).Elem()
}

func (o LookupAddress6ResultOutput) ToLookupAddress6ResultOutput() LookupAddress6ResultOutput {
	return o
}

func (o LookupAddress6ResultOutput) ToLookupAddress6ResultOutputWithContext(ctx context.Context) LookupAddress6ResultOutput {
	return o
}

// Minimal TTL of individual IPv6 addresses in FQDN cache.
func (o LookupAddress6ResultOutput) CacheTtl() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAddress6Result) int { return v.CacheTtl }).(pulumi.IntOutput)
}

// Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
func (o LookupAddress6ResultOutput) Color() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAddress6Result) int { return v.Color }).(pulumi.IntOutput)
}

// Comment.
func (o LookupAddress6ResultOutput) Comment() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAddress6Result) string { return v.Comment }).(pulumi.StringOutput)
}

// IPv6 addresses associated to a specific country.
func (o LookupAddress6ResultOutput) Country() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAddress6Result) string { return v.Country }).(pulumi.StringOutput)
}

// Final IP address (inclusive) in the range for the address (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).
func (o LookupAddress6ResultOutput) EndIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAddress6Result) string { return v.EndIp }).(pulumi.StringOutput)
}

// Last MAC address in the range.
func (o LookupAddress6ResultOutput) EndMac() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAddress6Result) string { return v.EndMac }).(pulumi.StringOutput)
}

// Endpoint group name.
func (o LookupAddress6ResultOutput) EpgName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAddress6Result) string { return v.EpgName }).(pulumi.StringOutput)
}

// Security Fabric global object setting.
func (o LookupAddress6ResultOutput) FabricObject() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAddress6Result) string { return v.FabricObject }).(pulumi.StringOutput)
}

// Fully qualified domain name.
func (o LookupAddress6ResultOutput) Fqdn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAddress6Result) string { return v.Fqdn }).(pulumi.StringOutput)
}

// Host Address.
func (o LookupAddress6ResultOutput) Host() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAddress6Result) string { return v.Host }).(pulumi.StringOutput)
}

// Host type.
func (o LookupAddress6ResultOutput) HostType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAddress6Result) string { return v.HostType }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupAddress6ResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAddress6Result) string { return v.Id }).(pulumi.StringOutput)
}

// IPv6 address prefix (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx).
func (o LookupAddress6ResultOutput) Ip6() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAddress6Result) string { return v.Ip6 }).(pulumi.StringOutput)
}

// IP address list. The structure of `list` block is documented below.
func (o LookupAddress6ResultOutput) Lists() GetAddress6ListArrayOutput {
	return o.ApplyT(func(v LookupAddress6Result) []GetAddress6List { return v.Lists }).(GetAddress6ListArrayOutput)
}

// MAC address ranges <start>[-<end>] separated by space.
func (o LookupAddress6ResultOutput) Macaddrs() GetAddress6MacaddrArrayOutput {
	return o.ApplyT(func(v LookupAddress6Result) []GetAddress6Macaddr { return v.Macaddrs }).(GetAddress6MacaddrArrayOutput)
}

// Name.
func (o LookupAddress6ResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAddress6Result) string { return v.Name }).(pulumi.StringOutput)
}

// Object ID for NSX.
func (o LookupAddress6ResultOutput) ObjId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAddress6Result) string { return v.ObjId }).(pulumi.StringOutput)
}

// SDN.
func (o LookupAddress6ResultOutput) Sdn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAddress6Result) string { return v.Sdn }).(pulumi.StringOutput)
}

// SDN Tag.
func (o LookupAddress6ResultOutput) SdnTag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAddress6Result) string { return v.SdnTag }).(pulumi.StringOutput)
}

// First IP address (inclusive) in the range for the address (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).
func (o LookupAddress6ResultOutput) StartIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAddress6Result) string { return v.StartIp }).(pulumi.StringOutput)
}

// First MAC address in the range.
func (o LookupAddress6ResultOutput) StartMac() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAddress6Result) string { return v.StartMac }).(pulumi.StringOutput)
}

// IPv6 subnet segments. The structure of `subnetSegment` block is documented below.
func (o LookupAddress6ResultOutput) SubnetSegments() GetAddress6SubnetSegmentArrayOutput {
	return o.ApplyT(func(v LookupAddress6Result) []GetAddress6SubnetSegment { return v.SubnetSegments }).(GetAddress6SubnetSegmentArrayOutput)
}

// Config object tagging The structure of `tagging` block is documented below.
func (o LookupAddress6ResultOutput) Taggings() GetAddress6TaggingArrayOutput {
	return o.ApplyT(func(v LookupAddress6Result) []GetAddress6Tagging { return v.Taggings }).(GetAddress6TaggingArrayOutput)
}

// IPv6 address template.
func (o LookupAddress6ResultOutput) Template() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAddress6Result) string { return v.Template }).(pulumi.StringOutput)
}

// Tenant.
func (o LookupAddress6ResultOutput) Tenant() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAddress6Result) string { return v.Tenant }).(pulumi.StringOutput)
}

// Subnet segment type.
func (o LookupAddress6ResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAddress6Result) string { return v.Type }).(pulumi.StringOutput)
}

// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
func (o LookupAddress6ResultOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAddress6Result) string { return v.Uuid }).(pulumi.StringOutput)
}

func (o LookupAddress6ResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAddress6Result) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// Enable/disable the visibility of the object in the GUI.
func (o LookupAddress6ResultOutput) Visibility() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAddress6Result) string { return v.Visibility }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAddress6ResultOutput{})
}
