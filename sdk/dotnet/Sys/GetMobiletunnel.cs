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
    public static class GetMobiletunnel
    {
        /// <summary>
        /// Use this data source to get information on an fortios system mobiletunnel
        /// </summary>
        public static Task<GetMobiletunnelResult> InvokeAsync(GetMobiletunnelArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetMobiletunnelResult>("fortios:sys/getMobiletunnel:getMobiletunnel", args ?? new GetMobiletunnelArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on an fortios system mobiletunnel
        /// </summary>
        public static Output<GetMobiletunnelResult> Invoke(GetMobiletunnelInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetMobiletunnelResult>("fortios:sys/getMobiletunnel:getMobiletunnel", args ?? new GetMobiletunnelInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetMobiletunnelArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired system mobiletunnel.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetMobiletunnelArgs()
        {
        }
        public static new GetMobiletunnelArgs Empty => new GetMobiletunnelArgs();
    }

    public sealed class GetMobiletunnelInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired system mobiletunnel.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetMobiletunnelInvokeArgs()
        {
        }
        public static new GetMobiletunnelInvokeArgs Empty => new GetMobiletunnelInvokeArgs();
    }


    [OutputType]
    public sealed class GetMobiletunnelResult
    {
        /// <summary>
        /// Hash Algorithm (Keyed MD5).
        /// </summary>
        public readonly string HashAlgorithm;
        /// <summary>
        /// Home IP address (Format: xxx.xxx.xxx.xxx).
        /// </summary>
        public readonly string HomeAddress;
        /// <summary>
        /// IPv4 address of the NEMO HA (Format: xxx.xxx.xxx.xxx).
        /// </summary>
        public readonly string HomeAgent;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// NMMO HA registration request lifetime (180 - 65535 sec, default = 65535).
        /// </summary>
        public readonly int Lifetime;
        /// <summary>
        /// NEMO authentication key.
        /// </summary>
        public readonly string NMhaeKey;
        /// <summary>
        /// NEMO authentication key type (ascii or base64).
        /// </summary>
        public readonly string NMhaeKeyType;
        /// <summary>
        /// NEMO authentication SPI (default: 256).
        /// </summary>
        public readonly int NMhaeSpi;
        /// <summary>
        /// Tunnel name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// NEMO network configuration. The structure of `network` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetMobiletunnelNetworkResult> Networks;
        /// <summary>
        /// NMMO HA registration interval (5 - 300, default = 5).
        /// </summary>
        public readonly int RegInterval;
        /// <summary>
        /// Maximum number of NMMO HA registration retries (1 to 30, default = 3).
        /// </summary>
        public readonly int RegRetry;
        /// <summary>
        /// Time before lifetime expiraton to send NMMO HA re-registration (5 - 60, default = 60).
        /// </summary>
        public readonly int RenewInterval;
        /// <summary>
        /// Select the associated interface name from available options.
        /// </summary>
        public readonly string RoamingInterface;
        /// <summary>
        /// Enable/disable this mobile tunnel.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// NEMO tunnnel mode (GRE tunnel).
        /// </summary>
        public readonly string TunnelMode;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetMobiletunnelResult(
            string hashAlgorithm,

            string homeAddress,

            string homeAgent,

            string id,

            int lifetime,

            string nMhaeKey,

            string nMhaeKeyType,

            int nMhaeSpi,

            string name,

            ImmutableArray<Outputs.GetMobiletunnelNetworkResult> networks,

            int regInterval,

            int regRetry,

            int renewInterval,

            string roamingInterface,

            string status,

            string tunnelMode,

            string? vdomparam)
        {
            HashAlgorithm = hashAlgorithm;
            HomeAddress = homeAddress;
            HomeAgent = homeAgent;
            Id = id;
            Lifetime = lifetime;
            NMhaeKey = nMhaeKey;
            NMhaeKeyType = nMhaeKeyType;
            NMhaeSpi = nMhaeSpi;
            Name = name;
            Networks = networks;
            RegInterval = regInterval;
            RegRetry = regRetry;
            RenewInterval = renewInterval;
            RoamingInterface = roamingInterface;
            Status = status;
            TunnelMode = tunnelMode;
            Vdomparam = vdomparam;
        }
    }
}