// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Switchcontroller.Outputs
{

    [OutputType]
    public sealed class FlowtrackingCollector
    {
        /// <summary>
        /// Collector IP address.
        /// </summary>
        public readonly string? Ip;
        /// <summary>
        /// Collector name.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Collector port number(0-65535, default:0, netflow:2055, ipfix:4739).
        /// </summary>
        public readonly int? Port;
        /// <summary>
        /// Collector L4 transport protocol for exporting packets. Valid values: `udp`, `tcp`, `sctp`.
        /// </summary>
        public readonly string? Transport;

        [OutputConstructor]
        private FlowtrackingCollector(
            string? ip,

            string? name,

            int? port,

            string? transport)
        {
            Ip = ip;
            Name = name;
            Port = port;
            Transport = transport;
        }
    }
}
