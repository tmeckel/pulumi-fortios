// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fmg

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource supports handling JSON RPC request for FortiManager.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/fmg"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := fmg.NewJsonrpcRequest(ctx, "test1", &fmg.JsonrpcRequestArgs{
//				JsonContent: pulumi.String("{\n  \"method\": \"add\",\n  \"params\": [\n    {\n      \"data\": [\n        {\n          \"action\": \"accept\",\n          \"dstaddr\": [\"all\"],\n          \"dstintf\": \"any\",\n          \"name\": \"policytest\",\n          \"schedule\": \"none\",\n          \"service\": \"ALL\",\n          \"srcaddr\": \"all\",\n          \"srcintf\": \"any\",\n          \"internet-service\": \"enable\",\n          \"internet-service-id\": \"Alibaba-Web\",\n          \"internet-service-src\": \"enable\",\n          \"internet-service-src-id\": \"Alibaba-Web\",\n          \"users\": \"guest\",\n          \"groups\": \"Guest-group\"\n        }\n      ],\n      \"url\": \"/pm/config/adom/root/pkg/default/firewall/policy\"\n    }\n  ]\n}\n\n"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = fmg.NewJsonrpcRequest(ctx, "test2", &fmg.JsonrpcRequestArgs{
//				JsonContent: pulumi.String("{\n  \"method\": \"add\",\n  \"params\": [\n    {\n      \"data\": [\n        {\n          \"ip\": \"192.168.1.2\",\n          \"name\": \"logserver4\",\n          \"port\": \"514\"\n        }\n      ],\n      \"url\": \"/cli/global/system/syslog\"\n    }\n  ]\n}\n\n"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = fmg.NewJsonrpcRequest(ctx, "test3", &fmg.JsonrpcRequestArgs{
//				JsonContent: pulumi.String("{\n  \"method\": \"get\",\n  \"params\": [\n    {\n      \"url\": \"/cli/global/system/admin/user/APIUser\"\n    }\n  ]\n}\n\n"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type JsonrpcRequest struct {
	pulumi.CustomResourceState

	// Comment.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// JSON RPC request, which should contain 'method' and 'params' parameters.
	JsonContent pulumi.StringOutput `pulumi:"jsonContent"`
	// JSON RPC request response data.
	Response pulumi.StringOutput `pulumi:"response"`
}

// NewJsonrpcRequest registers a new resource with the given unique name, arguments, and options.
func NewJsonrpcRequest(ctx *pulumi.Context,
	name string, args *JsonrpcRequestArgs, opts ...pulumi.ResourceOption) (*JsonrpcRequest, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.JsonContent == nil {
		return nil, errors.New("invalid value for required argument 'JsonContent'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource JsonrpcRequest
	err := ctx.RegisterResource("fortios:fmg/jsonrpcRequest:JsonrpcRequest", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetJsonrpcRequest gets an existing JsonrpcRequest resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetJsonrpcRequest(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JsonrpcRequestState, opts ...pulumi.ResourceOption) (*JsonrpcRequest, error) {
	var resource JsonrpcRequest
	err := ctx.ReadResource("fortios:fmg/jsonrpcRequest:JsonrpcRequest", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering JsonrpcRequest resources.
type jsonrpcRequestState struct {
	// Comment.
	Comment *string `pulumi:"comment"`
	// JSON RPC request, which should contain 'method' and 'params' parameters.
	JsonContent *string `pulumi:"jsonContent"`
	// JSON RPC request response data.
	Response *string `pulumi:"response"`
}

type JsonrpcRequestState struct {
	// Comment.
	Comment pulumi.StringPtrInput
	// JSON RPC request, which should contain 'method' and 'params' parameters.
	JsonContent pulumi.StringPtrInput
	// JSON RPC request response data.
	Response pulumi.StringPtrInput
}

func (JsonrpcRequestState) ElementType() reflect.Type {
	return reflect.TypeOf((*jsonrpcRequestState)(nil)).Elem()
}

type jsonrpcRequestArgs struct {
	// Comment.
	Comment *string `pulumi:"comment"`
	// JSON RPC request, which should contain 'method' and 'params' parameters.
	JsonContent string `pulumi:"jsonContent"`
}

// The set of arguments for constructing a JsonrpcRequest resource.
type JsonrpcRequestArgs struct {
	// Comment.
	Comment pulumi.StringPtrInput
	// JSON RPC request, which should contain 'method' and 'params' parameters.
	JsonContent pulumi.StringInput
}

func (JsonrpcRequestArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*jsonrpcRequestArgs)(nil)).Elem()
}

type JsonrpcRequestInput interface {
	pulumi.Input

	ToJsonrpcRequestOutput() JsonrpcRequestOutput
	ToJsonrpcRequestOutputWithContext(ctx context.Context) JsonrpcRequestOutput
}

func (*JsonrpcRequest) ElementType() reflect.Type {
	return reflect.TypeOf((**JsonrpcRequest)(nil)).Elem()
}

func (i *JsonrpcRequest) ToJsonrpcRequestOutput() JsonrpcRequestOutput {
	return i.ToJsonrpcRequestOutputWithContext(context.Background())
}

func (i *JsonrpcRequest) ToJsonrpcRequestOutputWithContext(ctx context.Context) JsonrpcRequestOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JsonrpcRequestOutput)
}

// JsonrpcRequestArrayInput is an input type that accepts JsonrpcRequestArray and JsonrpcRequestArrayOutput values.
// You can construct a concrete instance of `JsonrpcRequestArrayInput` via:
//
//	JsonrpcRequestArray{ JsonrpcRequestArgs{...} }
type JsonrpcRequestArrayInput interface {
	pulumi.Input

	ToJsonrpcRequestArrayOutput() JsonrpcRequestArrayOutput
	ToJsonrpcRequestArrayOutputWithContext(context.Context) JsonrpcRequestArrayOutput
}

type JsonrpcRequestArray []JsonrpcRequestInput

func (JsonrpcRequestArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*JsonrpcRequest)(nil)).Elem()
}

func (i JsonrpcRequestArray) ToJsonrpcRequestArrayOutput() JsonrpcRequestArrayOutput {
	return i.ToJsonrpcRequestArrayOutputWithContext(context.Background())
}

func (i JsonrpcRequestArray) ToJsonrpcRequestArrayOutputWithContext(ctx context.Context) JsonrpcRequestArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JsonrpcRequestArrayOutput)
}

// JsonrpcRequestMapInput is an input type that accepts JsonrpcRequestMap and JsonrpcRequestMapOutput values.
// You can construct a concrete instance of `JsonrpcRequestMapInput` via:
//
//	JsonrpcRequestMap{ "key": JsonrpcRequestArgs{...} }
type JsonrpcRequestMapInput interface {
	pulumi.Input

	ToJsonrpcRequestMapOutput() JsonrpcRequestMapOutput
	ToJsonrpcRequestMapOutputWithContext(context.Context) JsonrpcRequestMapOutput
}

type JsonrpcRequestMap map[string]JsonrpcRequestInput

func (JsonrpcRequestMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*JsonrpcRequest)(nil)).Elem()
}

func (i JsonrpcRequestMap) ToJsonrpcRequestMapOutput() JsonrpcRequestMapOutput {
	return i.ToJsonrpcRequestMapOutputWithContext(context.Background())
}

func (i JsonrpcRequestMap) ToJsonrpcRequestMapOutputWithContext(ctx context.Context) JsonrpcRequestMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JsonrpcRequestMapOutput)
}

type JsonrpcRequestOutput struct{ *pulumi.OutputState }

func (JsonrpcRequestOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JsonrpcRequest)(nil)).Elem()
}

func (o JsonrpcRequestOutput) ToJsonrpcRequestOutput() JsonrpcRequestOutput {
	return o
}

func (o JsonrpcRequestOutput) ToJsonrpcRequestOutputWithContext(ctx context.Context) JsonrpcRequestOutput {
	return o
}

// Comment.
func (o JsonrpcRequestOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JsonrpcRequest) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// JSON RPC request, which should contain 'method' and 'params' parameters.
func (o JsonrpcRequestOutput) JsonContent() pulumi.StringOutput {
	return o.ApplyT(func(v *JsonrpcRequest) pulumi.StringOutput { return v.JsonContent }).(pulumi.StringOutput)
}

// JSON RPC request response data.
func (o JsonrpcRequestOutput) Response() pulumi.StringOutput {
	return o.ApplyT(func(v *JsonrpcRequest) pulumi.StringOutput { return v.Response }).(pulumi.StringOutput)
}

type JsonrpcRequestArrayOutput struct{ *pulumi.OutputState }

func (JsonrpcRequestArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*JsonrpcRequest)(nil)).Elem()
}

func (o JsonrpcRequestArrayOutput) ToJsonrpcRequestArrayOutput() JsonrpcRequestArrayOutput {
	return o
}

func (o JsonrpcRequestArrayOutput) ToJsonrpcRequestArrayOutputWithContext(ctx context.Context) JsonrpcRequestArrayOutput {
	return o
}

func (o JsonrpcRequestArrayOutput) Index(i pulumi.IntInput) JsonrpcRequestOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *JsonrpcRequest {
		return vs[0].([]*JsonrpcRequest)[vs[1].(int)]
	}).(JsonrpcRequestOutput)
}

type JsonrpcRequestMapOutput struct{ *pulumi.OutputState }

func (JsonrpcRequestMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*JsonrpcRequest)(nil)).Elem()
}

func (o JsonrpcRequestMapOutput) ToJsonrpcRequestMapOutput() JsonrpcRequestMapOutput {
	return o
}

func (o JsonrpcRequestMapOutput) ToJsonrpcRequestMapOutputWithContext(ctx context.Context) JsonrpcRequestMapOutput {
	return o
}

func (o JsonrpcRequestMapOutput) MapIndex(k pulumi.StringInput) JsonrpcRequestOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *JsonrpcRequest {
		return vs[0].(map[string]*JsonrpcRequest)[vs[1].(string)]
	}).(JsonrpcRequestOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*JsonrpcRequestInput)(nil)).Elem(), &JsonrpcRequest{})
	pulumi.RegisterInputType(reflect.TypeOf((*JsonrpcRequestArrayInput)(nil)).Elem(), JsonrpcRequestArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*JsonrpcRequestMapInput)(nil)).Elem(), JsonrpcRequestMap{})
	pulumi.RegisterOutputType(JsonrpcRequestOutput{})
	pulumi.RegisterOutputType(JsonrpcRequestArrayOutput{})
	pulumi.RegisterOutputType(JsonrpcRequestMapOutput{})
}