// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Sys.Outputs
{

    [OutputType]
    public sealed class VirtualwanlinkMember
    {
        /// <summary>
        /// Comments.
        /// </summary>
        public readonly string? Comment;
        /// <summary>
        /// Cost of this interface for services in SLA mode (0 - 4294967295, default = 0).
        /// </summary>
        public readonly int? Cost;
        /// <summary>
        /// The default gateway for this interface. Usually the default gateway of the Internet service provider that this interface is connected to.
        /// </summary>
        public readonly string? Gateway;
        /// <summary>
        /// IPv6 gateway.
        /// </summary>
        public readonly string? Gateway6;
        /// <summary>
        /// Ingress spillover threshold for this interface (0 - 16776000 kbit/s). When this traffic volume threshold is reached, new sessions spill over to other interfaces in the SD-WAN.
        /// </summary>
        public readonly int? IngressSpilloverThreshold;
        /// <summary>
        /// Interface name.
        /// </summary>
        public readonly string? Interface;
        /// <summary>
        /// Priority of the interface (0 - 4294967295). Used for SD-WAN rules or priority rules.
        /// </summary>
        public readonly int? Priority;
        /// <summary>
        /// Member sequence number.
        /// </summary>
        public readonly int? SeqNum;
        /// <summary>
        /// Source IP address used in the health-check packet to the server.
        /// </summary>
        public readonly string? Source;
        /// <summary>
        /// Source IPv6 address used in the health-check packet to the server.
        /// </summary>
        public readonly string? Source6;
        /// <summary>
        /// Egress spillover threshold for this interface (0 - 16776000 kbit/s). When this traffic volume threshold is reached, new sessions spill over to other interfaces in the SD-WAN.
        /// </summary>
        public readonly int? SpilloverThreshold;
        /// <summary>
        /// Enable/disable this interface in the SD-WAN. Valid values: `disable`, `enable`.
        /// </summary>
        public readonly string? Status;
        /// <summary>
        /// Measured volume ratio (this value / sum of all values = percentage of link volume, 0 - 255).
        /// </summary>
        public readonly int? VolumeRatio;
        /// <summary>
        /// Weight of this interface for weighted load balancing. (0 - 255) More traffic is directed to interfaces with higher weights.
        /// </summary>
        public readonly int? Weight;

        [OutputConstructor]
        private VirtualwanlinkMember(
            string? comment,

            int? cost,

            string? gateway,

            string? gateway6,

            int? ingressSpilloverThreshold,

            string? @interface,

            int? priority,

            int? seqNum,

            string? source,

            string? source6,

            int? spilloverThreshold,

            string? status,

            int? volumeRatio,

            int? weight)
        {
            Comment = comment;
            Cost = cost;
            Gateway = gateway;
            Gateway6 = gateway6;
            IngressSpilloverThreshold = ingressSpilloverThreshold;
            Interface = @interface;
            Priority = priority;
            SeqNum = seqNum;
            Source = source;
            Source6 = source6;
            SpilloverThreshold = spilloverThreshold;
            Status = status;
            VolumeRatio = volumeRatio;
            Weight = weight;
        }
    }
}
