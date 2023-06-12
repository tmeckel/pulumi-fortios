// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Router.Inputs
{

    public sealed class BgpAggregateAddress6Args : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable generate AS set path information. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("asSet")]
        public Input<string>? AsSet { get; set; }

        /// <summary>
        /// ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// Aggregate IPv6 prefix.
        /// </summary>
        [Input("prefix6")]
        public Input<string>? Prefix6 { get; set; }

        /// <summary>
        /// Enable/disable filter more specific routes from updates. Valid values: `enable`, `disable`.
        /// 
        /// The `aggregate_address6` block supports:
        /// </summary>
        [Input("summaryOnly")]
        public Input<string>? SummaryOnly { get; set; }

        public BgpAggregateAddress6Args()
        {
        }
        public static new BgpAggregateAddress6Args Empty => new BgpAggregateAddress6Args();
    }
}
