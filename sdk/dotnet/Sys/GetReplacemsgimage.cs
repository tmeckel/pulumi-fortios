// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Sys
{
    public static class GetReplacemsgimage
    {
        /// <summary>
        /// Use this data source to get information on an fortios system replacemsgimage
        /// </summary>
        public static Task<GetReplacemsgimageResult> InvokeAsync(GetReplacemsgimageArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetReplacemsgimageResult>("fortios:sys/getReplacemsgimage:getReplacemsgimage", args ?? new GetReplacemsgimageArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on an fortios system replacemsgimage
        /// </summary>
        public static Output<GetReplacemsgimageResult> Invoke(GetReplacemsgimageInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetReplacemsgimageResult>("fortios:sys/getReplacemsgimage:getReplacemsgimage", args ?? new GetReplacemsgimageInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetReplacemsgimageArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired system replacemsgimage.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetReplacemsgimageArgs()
        {
        }
        public static new GetReplacemsgimageArgs Empty => new GetReplacemsgimageArgs();
    }

    public sealed class GetReplacemsgimageInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired system replacemsgimage.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetReplacemsgimageInvokeArgs()
        {
        }
        public static new GetReplacemsgimageInvokeArgs Empty => new GetReplacemsgimageInvokeArgs();
    }


    [OutputType]
    public sealed class GetReplacemsgimageResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Image data.
        /// </summary>
        public readonly string ImageBase64;
        /// <summary>
        /// Image type.
        /// </summary>
        public readonly string ImageType;
        /// <summary>
        /// Image name.
        /// </summary>
        public readonly string Name;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetReplacemsgimageResult(
            string id,

            string imageBase64,

            string imageType,

            string name,

            string? vdomparam)
        {
            Id = id;
            ImageBase64 = imageBase64;
            ImageType = imageType;
            Name = name;
            Vdomparam = vdomparam;
        }
    }
}
