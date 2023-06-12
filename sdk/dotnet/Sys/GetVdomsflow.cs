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
    public static class GetVdomsflow
    {
        /// <summary>
        /// Use this data source to get information on fortios system vdomsflow
        /// </summary>
        public static Task<GetVdomsflowResult> InvokeAsync(GetVdomsflowArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVdomsflowResult>("fortios:sys/getVdomsflow:getVdomsflow", args ?? new GetVdomsflowArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on fortios system vdomsflow
        /// </summary>
        public static Output<GetVdomsflowResult> Invoke(GetVdomsflowInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVdomsflowResult>("fortios:sys/getVdomsflow:getVdomsflow", args ?? new GetVdomsflowInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVdomsflowArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetVdomsflowArgs()
        {
        }
        public static new GetVdomsflowArgs Empty => new GetVdomsflowArgs();
    }

    public sealed class GetVdomsflowInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetVdomsflowInvokeArgs()
        {
        }
        public static new GetVdomsflowInvokeArgs Empty => new GetVdomsflowInvokeArgs();
    }


    [OutputType]
    public sealed class GetVdomsflowResult
    {
        /// <summary>
        /// IP address of the sFlow collector that sFlow agents added to interfaces in this VDOM send sFlow datagrams to (default = 0.0.0.0).
        /// </summary>
        public readonly string CollectorIp;
        /// <summary>
        /// UDP port number used for sending sFlow datagrams (configure only if required by your sFlow collector or your network configuration) (0 - 65535, default = 6343).
        /// </summary>
        public readonly int CollectorPort;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Specify outgoing interface to reach server.
        /// </summary>
        public readonly string Interface;
        /// <summary>
        /// Specify how to select outgoing interface to reach server.
        /// </summary>
        public readonly string InterfaceSelectMethod;
        /// <summary>
        /// Source IP address for sFlow agent.
        /// </summary>
        public readonly string SourceIp;
        /// <summary>
        /// Enable/disable the sFlow configuration for the current VDOM.
        /// </summary>
        public readonly string VdomSflow;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetVdomsflowResult(
            string collectorIp,

            int collectorPort,

            string id,

            string @interface,

            string interfaceSelectMethod,

            string sourceIp,

            string vdomSflow,

            string? vdomparam)
        {
            CollectorIp = collectorIp;
            CollectorPort = collectorPort;
            Id = id;
            Interface = @interface;
            InterfaceSelectMethod = interfaceSelectMethod;
            SourceIp = sourceIp;
            VdomSflow = vdomSflow;
            Vdomparam = vdomparam;
        }
    }
}