// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Spamfilter.Inputs
{

    public sealed class BwlEntryArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Reject, mark as spam or good email. Valid values: `reject`, `spam`, `clear`.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// IP address type. Valid values: `ipv4`, `ipv6`.
        /// </summary>
        [Input("addrType")]
        public Input<string>? AddrType { get; set; }

        /// <summary>
        /// Email address pattern.
        /// </summary>
        [Input("emailPattern")]
        public Input<string>? EmailPattern { get; set; }

        /// <summary>
        /// Entry ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// IPv4 network address/subnet mask bits.
        /// </summary>
        [Input("ip4Subnet")]
        public Input<string>? Ip4Subnet { get; set; }

        /// <summary>
        /// IPv6 network address/subnet mask bits.
        /// </summary>
        [Input("ip6Subnet")]
        public Input<string>? Ip6Subnet { get; set; }

        /// <summary>
        /// Wildcard pattern or regular expression. Valid values: `wildcard`, `regexp`.
        /// </summary>
        [Input("patternType")]
        public Input<string>? PatternType { get; set; }

        /// <summary>
        /// Enable/disable status. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Entry type. Valid values: `ip`, `email`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public BwlEntryArgs()
        {
        }
        public static new BwlEntryArgs Empty => new BwlEntryArgs();
    }
}