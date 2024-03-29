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
    public sealed class RipngDistance
    {
        /// <summary>
        /// Access list for route destination.
        /// </summary>
        public readonly string? AccessList6;
        /// <summary>
        /// Distance (1 - 255).
        /// </summary>
        public readonly int? Distance;
        /// <summary>
        /// Distance ID.
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// Distance prefix6.
        /// </summary>
        public readonly string? Prefix6;

        [OutputConstructor]
        private RipngDistance(
            string? accessList6,

            int? distance,

            int? id,

            string? prefix6)
        {
            AccessList6 = accessList6;
            Distance = distance;
            Id = id;
            Prefix6 = prefix6;
        }
    }
}
