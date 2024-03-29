// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Router.Outputs
{

    [OutputType]
    public sealed class BgpNeighborRange
    {
        /// <summary>
        /// IPv6 neighbor range ID.
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// Maximum number of neighbors.
        /// </summary>
        public readonly int? MaxNeighborNum;
        /// <summary>
        /// Neighbor group name.
        /// </summary>
        public readonly string? NeighborGroup;
        /// <summary>
        /// Neighbor range prefix.
        /// </summary>
        public readonly string? Prefix;

        [OutputConstructor]
        private BgpNeighborRange(
            int? id,

            int? maxNeighborNum,

            string? neighborGroup,

            string? prefix)
        {
            Id = id;
            MaxNeighborNum = maxNeighborNum;
            NeighborGroup = neighborGroup;
            Prefix = prefix;
        }
    }
}
