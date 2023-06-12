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

    public sealed class InternetserviceadditionEntryGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Address mode (IPv4 or IPv6) Valid values: `ipv4`, `ipv6`.
        /// </summary>
        [Input("addrMode")]
        public Input<string>? AddrMode { get; set; }

        /// <summary>
        /// Entry ID(1-255).
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        [Input("portRanges")]
        private InputList<Inputs.InternetserviceadditionEntryPortRangeGetArgs>? _portRanges;

        /// <summary>
        /// Port ranges in the custom entry. The structure of `port_range` block is documented below.
        /// </summary>
        public InputList<Inputs.InternetserviceadditionEntryPortRangeGetArgs> PortRanges
        {
            get => _portRanges ?? (_portRanges = new InputList<Inputs.InternetserviceadditionEntryPortRangeGetArgs>());
            set => _portRanges = value;
        }

        /// <summary>
        /// Integer value for the protocol type as defined by IANA (0 - 255).
        /// </summary>
        [Input("protocol")]
        public Input<int>? Protocol { get; set; }

        public InternetserviceadditionEntryGetArgs()
        {
        }
        public static new InternetserviceadditionEntryGetArgs Empty => new InternetserviceadditionEntryGetArgs();
    }
}
