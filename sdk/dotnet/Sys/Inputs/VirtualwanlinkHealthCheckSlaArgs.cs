// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Sys.Inputs
{

    public sealed class VirtualwanlinkHealthCheckSlaArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// SLA ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// Jitter for SLA to make decision in milliseconds. (0 - 10000000, default = 5).
        /// </summary>
        [Input("jitterThreshold")]
        public Input<int>? JitterThreshold { get; set; }

        /// <summary>
        /// Latency for SLA to make decision in milliseconds. (0 - 10000000, default = 5).
        /// </summary>
        [Input("latencyThreshold")]
        public Input<int>? LatencyThreshold { get; set; }

        /// <summary>
        /// Criteria on which to base link selection. Valid values: `latency`, `jitter`, `packet-loss`.
        /// </summary>
        [Input("linkCostFactor")]
        public Input<string>? LinkCostFactor { get; set; }

        /// <summary>
        /// Packet loss for SLA to make decision in percentage. (0 - 100, default = 0).
        /// </summary>
        [Input("packetlossThreshold")]
        public Input<int>? PacketlossThreshold { get; set; }

        public VirtualwanlinkHealthCheckSlaArgs()
        {
        }
        public static new VirtualwanlinkHealthCheckSlaArgs Empty => new VirtualwanlinkHealthCheckSlaArgs();
    }
}
