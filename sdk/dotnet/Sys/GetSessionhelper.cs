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
    public static class GetSessionhelper
    {
        /// <summary>
        /// Use this data source to get information on an fortios system sessionhelper
        /// </summary>
        public static Task<GetSessionhelperResult> InvokeAsync(GetSessionhelperArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSessionhelperResult>("fortios:sys/getSessionhelper:getSessionhelper", args ?? new GetSessionhelperArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on an fortios system sessionhelper
        /// </summary>
        public static Output<GetSessionhelperResult> Invoke(GetSessionhelperInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSessionhelperResult>("fortios:sys/getSessionhelper:getSessionhelper", args ?? new GetSessionhelperInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSessionhelperArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the fosid of the desired system sessionhelper.
        /// </summary>
        [Input("fosid", required: true)]
        public int Fosid { get; set; }

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetSessionhelperArgs()
        {
        }
        public static new GetSessionhelperArgs Empty => new GetSessionhelperArgs();
    }

    public sealed class GetSessionhelperInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the fosid of the desired system sessionhelper.
        /// </summary>
        [Input("fosid", required: true)]
        public Input<int> Fosid { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetSessionhelperInvokeArgs()
        {
        }
        public static new GetSessionhelperInvokeArgs Empty => new GetSessionhelperInvokeArgs();
    }


    [OutputType]
    public sealed class GetSessionhelperResult
    {
        /// <summary>
        /// Session helper ID.
        /// </summary>
        public readonly int Fosid;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Helper name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Protocol port.
        /// </summary>
        public readonly int Port;
        /// <summary>
        /// Protocol number.
        /// </summary>
        public readonly int Protocol;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetSessionhelperResult(
            int fosid,

            string id,

            string name,

            int port,

            int protocol,

            string? vdomparam)
        {
            Fosid = fosid;
            Id = id;
            Name = name;
            Port = port;
            Protocol = protocol;
            Vdomparam = vdomparam;
        }
    }
}
