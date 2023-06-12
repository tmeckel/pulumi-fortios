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
    public sealed class BgpAggregateAddress6
    {
        /// <summary>
        /// Enable/disable generate AS set path information. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? AsSet;
        /// <summary>
        /// ID.
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// Aggregate IPv6 prefix.
        /// </summary>
        public readonly string? Prefix6;
        /// <summary>
        /// Enable/disable filter more specific routes from updates. Valid values: `enable`, `disable`.
        /// 
        /// The `aggregate_address6` block supports:
        /// </summary>
        public readonly string? SummaryOnly;

        [OutputConstructor]
        private BgpAggregateAddress6(
            string? asSet,

            int? id,

            string? prefix6,

            string? summaryOnly)
        {
            AsSet = asSet;
            Id = id;
            Prefix6 = prefix6;
            SummaryOnly = summaryOnly;
        }
    }
}
