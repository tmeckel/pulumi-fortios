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

    public sealed class SdwanHealthCheckSlaArgs : global::Pulumi.ResourceArgs
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
        /// Criteria on which to base link selection.
        /// </summary>
        [Input("linkCostFactor")]
        public Input<string>? LinkCostFactor { get; set; }

        /// <summary>
        /// Minimum Mean Opinion Score for SLA to be marked as pass. (1.0 - 5.0, default = 3.6).
        /// </summary>
        [Input("mosThreshold")]
        public Input<string>? MosThreshold { get; set; }

        /// <summary>
        /// Packet loss for SLA to make decision in percentage. (0 - 100, default = 0).
        /// </summary>
        [Input("packetlossThreshold")]
        public Input<int>? PacketlossThreshold { get; set; }

        /// <summary>
        /// Value to be distributed into routing table when in-sla (0 - 65535, default = 0).
        /// </summary>
        [Input("priorityInSla")]
        public Input<int>? PriorityInSla { get; set; }

        /// <summary>
        /// Value to be distributed into routing table when out-sla (0 - 65535, default = 0).
        /// </summary>
        [Input("priorityOutSla")]
        public Input<int>? PriorityOutSla { get; set; }

        public SdwanHealthCheckSlaArgs()
        {
        }
        public static new SdwanHealthCheckSlaArgs Empty => new SdwanHealthCheckSlaArgs();
    }
}
