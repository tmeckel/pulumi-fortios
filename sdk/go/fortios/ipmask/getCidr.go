// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ipmask

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Convert IP/Mask to CIDR
//
// ## Example Usage
// ### Example1
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/ipmask"
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/sys"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			trnameInterface, err := sys.LookupInterface(ctx, &sys.LookupInterfaceArgs{
//				Name: "port3",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			trnameCidr, err := ipmask.GetCidr(ctx, &ipmask.GetCidrArgs{
//				Ipmask: pulumi.StringRef(trnameInterface.Ip),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("output1", trnameCidr.Cidr)
//			return nil
//		})
//	}
//
// ```
// ### Example2
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/ipmask"
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/sys"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			trnameInterface, err := sys.LookupInterface(ctx, &sys.LookupInterfaceArgs{
//				Name: "port3",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			trnameCidr, err := ipmask.GetCidr(ctx, &ipmask.GetCidrArgs{
//				Ipmask: pulumi.StringRef(trnameInterface.Ip),
//				Ipmasklists: []string{
//					"21.1.1.1 255.255.255.0",
//					"22.1.1.1 255.255.255.240",
//					"23.1.1.1 255.255.255.224",
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("outputConv1", trnameCidr.Cidr)
//			ctx.Export("outputConv2", trnameCidr.Cidrlists)
//			ctx.Export("outputOrignal", trnameInterface.Ip)
//			return nil
//		})
//	}
//
// ```
func GetCidr(ctx *pulumi.Context, args *GetCidrArgs, opts ...pulumi.InvokeOption) (*GetCidrResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetCidrResult
	err := ctx.Invoke("fortios:ipmask/getCidr:getCidr", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCidr.
type GetCidrArgs struct {
	// Specify IP/MASK.
	Ipmask *string `pulumi:"ipmask"`
	// Specify IP/MASK list.
	Ipmasklists []string `pulumi:"ipmasklists"`
}

// A collection of values returned by getCidr.
type GetCidrResult struct {
	// Classless Inter-Domain Routing of the IP/MASK.
	Cidr string `pulumi:"cidr"`
	// Classless Inter-Domain Routing list converted from the IP/MASK list.
	Cidrlists []string `pulumi:"cidrlists"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// IP/MASK.
	Ipmask *string `pulumi:"ipmask"`
	// IP/MASK list.
	Ipmasklists []string `pulumi:"ipmasklists"`
}

func GetCidrOutput(ctx *pulumi.Context, args GetCidrOutputArgs, opts ...pulumi.InvokeOption) GetCidrResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetCidrResult, error) {
			args := v.(GetCidrArgs)
			r, err := GetCidr(ctx, &args, opts...)
			var s GetCidrResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetCidrResultOutput)
}

// A collection of arguments for invoking getCidr.
type GetCidrOutputArgs struct {
	// Specify IP/MASK.
	Ipmask pulumi.StringPtrInput `pulumi:"ipmask"`
	// Specify IP/MASK list.
	Ipmasklists pulumi.StringArrayInput `pulumi:"ipmasklists"`
}

func (GetCidrOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCidrArgs)(nil)).Elem()
}

// A collection of values returned by getCidr.
type GetCidrResultOutput struct{ *pulumi.OutputState }

func (GetCidrResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCidrResult)(nil)).Elem()
}

func (o GetCidrResultOutput) ToGetCidrResultOutput() GetCidrResultOutput {
	return o
}

func (o GetCidrResultOutput) ToGetCidrResultOutputWithContext(ctx context.Context) GetCidrResultOutput {
	return o
}

// Classless Inter-Domain Routing of the IP/MASK.
func (o GetCidrResultOutput) Cidr() pulumi.StringOutput {
	return o.ApplyT(func(v GetCidrResult) string { return v.Cidr }).(pulumi.StringOutput)
}

// Classless Inter-Domain Routing list converted from the IP/MASK list.
func (o GetCidrResultOutput) Cidrlists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetCidrResult) []string { return v.Cidrlists }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetCidrResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetCidrResult) string { return v.Id }).(pulumi.StringOutput)
}

// IP/MASK.
func (o GetCidrResultOutput) Ipmask() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCidrResult) *string { return v.Ipmask }).(pulumi.StringPtrOutput)
}

// IP/MASK list.
func (o GetCidrResultOutput) Ipmasklists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetCidrResult) []string { return v.Ipmasklists }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCidrResultOutput{})
}
