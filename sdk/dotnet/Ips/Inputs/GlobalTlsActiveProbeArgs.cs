// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Ips.Inputs
{

    public sealed class GlobalTlsActiveProbeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specify outgoing interface to reach server.
        /// </summary>
        [Input("interface")]
        public Input<string>? Interface { get; set; }

        /// <summary>
        /// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
        /// </summary>
        [Input("interfaceSelectMethod")]
        public Input<string>? InterfaceSelectMethod { get; set; }

        /// <summary>
        /// Source IP address used for TLS active probe.
        /// </summary>
        [Input("sourceIp")]
        public Input<string>? SourceIp { get; set; }

        /// <summary>
        /// Source IPv6 address used for TLS active probe.
        /// </summary>
        [Input("sourceIp6")]
        public Input<string>? SourceIp6 { get; set; }

        /// <summary>
        /// Virtual domain name for TLS active probe.
        /// </summary>
        [Input("vdom")]
        public Input<string>? Vdom { get; set; }

        public GlobalTlsActiveProbeArgs()
        {
        }
        public static new GlobalTlsActiveProbeArgs Empty => new GlobalTlsActiveProbeArgs();
    }
}
