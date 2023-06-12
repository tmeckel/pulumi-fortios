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
    public static class GetPppoeinterface
    {
        /// <summary>
        /// Use this data source to get information on an fortios system pppoeinterface
        /// </summary>
        public static Task<GetPppoeinterfaceResult> InvokeAsync(GetPppoeinterfaceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPppoeinterfaceResult>("fortios:sys/getPppoeinterface:getPppoeinterface", args ?? new GetPppoeinterfaceArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on an fortios system pppoeinterface
        /// </summary>
        public static Output<GetPppoeinterfaceResult> Invoke(GetPppoeinterfaceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPppoeinterfaceResult>("fortios:sys/getPppoeinterface:getPppoeinterface", args ?? new GetPppoeinterfaceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPppoeinterfaceArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired system pppoeinterface.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetPppoeinterfaceArgs()
        {
        }
        public static new GetPppoeinterfaceArgs Empty => new GetPppoeinterfaceArgs();
    }

    public sealed class GetPppoeinterfaceInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired system pppoeinterface.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetPppoeinterfaceInvokeArgs()
        {
        }
        public static new GetPppoeinterfaceInvokeArgs Empty => new GetPppoeinterfaceInvokeArgs();
    }


    [OutputType]
    public sealed class GetPppoeinterfaceResult
    {
        /// <summary>
        /// PPPoE AC name.
        /// </summary>
        public readonly string AcName;
        /// <summary>
        /// PPP authentication type to use.
        /// </summary>
        public readonly string AuthType;
        /// <summary>
        /// Name for the physical interface.
        /// </summary>
        public readonly string Device;
        /// <summary>
        /// Enable/disable dial on demand to dial the PPPoE interface when packets are routed to the PPPoE interface.
        /// </summary>
        public readonly string DialOnDemand;
        /// <summary>
        /// PPPoE discovery init timeout value in (0-4294967295 sec).
        /// </summary>
        public readonly int DiscRetryTimeout;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// PPPoE auto disconnect after idle timeout (0-4294967295 sec).
        /// </summary>
        public readonly int IdleTimeout;
        /// <summary>
        /// PPPoE unnumbered IP.
        /// </summary>
        public readonly string Ipunnumbered;
        /// <summary>
        /// Enable/disable IPv6 Control Protocol (IPv6CP).
        /// </summary>
        public readonly string Ipv6;
        /// <summary>
        /// PPPoE LCP echo interval in (0-4294967295 sec, default = 5).
        /// </summary>
        public readonly int LcpEchoInterval;
        /// <summary>
        /// Maximum missed LCP echo messages before disconnect (0-4294967295, default = 3).
        /// </summary>
        public readonly int LcpMaxEchoFails;
        /// <summary>
        /// Name of the PPPoE interface.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// PPPoE terminate timeout value in (0-4294967295 sec).
        /// </summary>
        public readonly int PadtRetryTimeout;
        /// <summary>
        /// Enter the password.
        /// </summary>
        public readonly string Password;
        /// <summary>
        /// Enable/disable PPPoE unnumbered negotiation.
        /// </summary>
        public readonly string PppoeUnnumberedNegotiate;
        /// <summary>
        /// PPPoE service name.
        /// </summary>
        public readonly string ServiceName;
        /// <summary>
        /// User name.
        /// </summary>
        public readonly string Username;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetPppoeinterfaceResult(
            string acName,

            string authType,

            string device,

            string dialOnDemand,

            int discRetryTimeout,

            string id,

            int idleTimeout,

            string ipunnumbered,

            string ipv6,

            int lcpEchoInterval,

            int lcpMaxEchoFails,

            string name,

            int padtRetryTimeout,

            string password,

            string pppoeUnnumberedNegotiate,

            string serviceName,

            string username,

            string? vdomparam)
        {
            AcName = acName;
            AuthType = authType;
            Device = device;
            DialOnDemand = dialOnDemand;
            DiscRetryTimeout = discRetryTimeout;
            Id = id;
            IdleTimeout = idleTimeout;
            Ipunnumbered = ipunnumbered;
            Ipv6 = ipv6;
            LcpEchoInterval = lcpEchoInterval;
            LcpMaxEchoFails = lcpMaxEchoFails;
            Name = name;
            PadtRetryTimeout = padtRetryTimeout;
            Password = password;
            PppoeUnnumberedNegotiate = pppoeUnnumberedNegotiate;
            ServiceName = serviceName;
            Username = username;
            Vdomparam = vdomparam;
        }
    }
}
