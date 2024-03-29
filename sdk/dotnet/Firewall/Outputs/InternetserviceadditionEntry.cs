// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Firewall.Outputs
{

    [OutputType]
    public sealed class InternetserviceadditionEntry
    {
        /// <summary>
        /// Address mode (IPv4 or IPv6) Valid values: `ipv4`, `ipv6`.
        /// </summary>
        public readonly string? AddrMode;
        /// <summary>
        /// Entry ID(1-255).
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// Port ranges in the custom entry. The structure of `port_range` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.InternetserviceadditionEntryPortRange> PortRanges;
        /// <summary>
        /// Integer value for the protocol type as defined by IANA (0 - 255).
        /// </summary>
        public readonly int? Protocol;

        [OutputConstructor]
        private InternetserviceadditionEntry(
            string? addrMode,

            int? id,

            ImmutableArray<Outputs.InternetserviceadditionEntryPortRange> portRanges,

            int? protocol)
        {
            AddrMode = addrMode;
            Id = id;
            PortRanges = portRanges;
            Protocol = protocol;
        }
    }
}
