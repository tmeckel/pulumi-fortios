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
    public sealed class GetBgpAdminDistanceResult
    {
        /// <summary>
        /// Administrative distance to apply (1 - 255).
        /// </summary>
        public readonly int Distance;
        /// <summary>
        /// ID.
        /// </summary>
        public readonly int Id;
        /// <summary>
        /// Neighbor address prefix.
        /// </summary>
        public readonly string NeighbourPrefix;
        /// <summary>
        /// Access list of routes to apply new distance to.
        /// </summary>
        public readonly string RouteList;

        [OutputConstructor]
        private GetBgpAdminDistanceResult(
            int distance,

            int id,

            string neighbourPrefix,

            string routeList)
        {
            Distance = distance;
            Id = id;
            NeighbourPrefix = neighbourPrefix;
            RouteList = routeList;
        }
    }
}
