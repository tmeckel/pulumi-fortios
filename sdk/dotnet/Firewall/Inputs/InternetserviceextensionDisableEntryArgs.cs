// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Firewall.Inputs
{

    public sealed class InternetserviceextensionDisableEntryArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Address mode (IPv4 or IPv6) Valid values: `ipv4`, `ipv6`.
        /// </summary>
        [Input("addrMode")]
        public Input<string>? AddrMode { get; set; }

        /// <summary>
        /// Disable entry ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        [Input("ip6Ranges")]
        private InputList<Inputs.InternetserviceextensionDisableEntryIp6RangeArgs>? _ip6Ranges;

        /// <summary>
        /// IPv6 ranges in the disable entry. The structure of `ip6_range` block is documented below.
        /// </summary>
        public InputList<Inputs.InternetserviceextensionDisableEntryIp6RangeArgs> Ip6Ranges
        {
            get => _ip6Ranges ?? (_ip6Ranges = new InputList<Inputs.InternetserviceextensionDisableEntryIp6RangeArgs>());
            set => _ip6Ranges = value;
        }

        [Input("ipRanges")]
        private InputList<Inputs.InternetserviceextensionDisableEntryIpRangeArgs>? _ipRanges;

        /// <summary>
        /// IP ranges in the disable entry. The structure of `ip_range` block is documented below.
        /// </summary>
        public InputList<Inputs.InternetserviceextensionDisableEntryIpRangeArgs> IpRanges
        {
            get => _ipRanges ?? (_ipRanges = new InputList<Inputs.InternetserviceextensionDisableEntryIpRangeArgs>());
            set => _ipRanges = value;
        }

        /// <summary>
        /// Integer value for the TCP/IP port (0 - 65535).
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        [Input("portRanges")]
        private InputList<Inputs.InternetserviceextensionDisableEntryPortRangeArgs>? _portRanges;

        /// <summary>
        /// Port ranges in the disable entry. The structure of `port_range` block is documented below.
        /// </summary>
        public InputList<Inputs.InternetserviceextensionDisableEntryPortRangeArgs> PortRanges
        {
            get => _portRanges ?? (_portRanges = new InputList<Inputs.InternetserviceextensionDisableEntryPortRangeArgs>());
            set => _portRanges = value;
        }

        /// <summary>
        /// Integer value for the protocol type as defined by IANA (0 - 255).
        /// </summary>
        [Input("protocol")]
        public Input<int>? Protocol { get; set; }

        public InternetserviceextensionDisableEntryArgs()
        {
        }
        public static new InternetserviceextensionDisableEntryArgs Empty => new InternetserviceextensionDisableEntryArgs();
    }
}
