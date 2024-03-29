// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Vpnipsec.Outputs
{

    [OutputType]
    public sealed class Phase1Ipv4ExcludeRange
    {
        /// <summary>
        /// End of IPv4 exclusive range.
        /// 
        /// The `ipv6_exclude_range` block supports:
        /// </summary>
        public readonly string? EndIp;
        /// <summary>
        /// ID.
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// Start of IPv4 exclusive range.
        /// </summary>
        public readonly string? StartIp;

        [OutputConstructor]
        private Phase1Ipv4ExcludeRange(
            string? endIp,

            int? id,

            string? startIp)
        {
            EndIp = endIp;
            Id = id;
            StartIp = startIp;
        }
    }
}
