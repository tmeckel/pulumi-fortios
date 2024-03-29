// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Vpnipsec.Inputs
{

    public sealed class Phase1Ipv4ExcludeRangeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// End of IPv4 exclusive range.
        /// 
        /// The `ipv6_exclude_range` block supports:
        /// </summary>
        [Input("endIp")]
        public Input<string>? EndIp { get; set; }

        /// <summary>
        /// ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// Start of IPv4 exclusive range.
        /// </summary>
        [Input("startIp")]
        public Input<string>? StartIp { get; set; }

        public Phase1Ipv4ExcludeRangeArgs()
        {
        }
        public static new Phase1Ipv4ExcludeRangeArgs Empty => new Phase1Ipv4ExcludeRangeArgs();
    }
}
