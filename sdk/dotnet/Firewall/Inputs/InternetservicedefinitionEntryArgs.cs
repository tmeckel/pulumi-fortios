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

    public sealed class InternetservicedefinitionEntryArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Internet Service category ID.
        /// </summary>
        [Input("categoryId")]
        public Input<int>? CategoryId { get; set; }

        /// <summary>
        /// Internet Service name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Integer value for ending TCP/UDP/SCTP destination port in range (0 to 65535). 0 means undefined.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        [Input("portRanges")]
        private InputList<Inputs.InternetservicedefinitionEntryPortRangeArgs>? _portRanges;

        /// <summary>
        /// Port ranges in the definition entry. The structure of `port_range` block is documented below.
        /// </summary>
        public InputList<Inputs.InternetservicedefinitionEntryPortRangeArgs> PortRanges
        {
            get => _portRanges ?? (_portRanges = new InputList<Inputs.InternetservicedefinitionEntryPortRangeArgs>());
            set => _portRanges = value;
        }

        /// <summary>
        /// Integer value for the protocol type as defined by IANA (0 - 255).
        /// </summary>
        [Input("protocol")]
        public Input<int>? Protocol { get; set; }

        /// <summary>
        /// Entry sequence number.
        /// </summary>
        [Input("seqNum")]
        public Input<int>? SeqNum { get; set; }

        public InternetservicedefinitionEntryArgs()
        {
        }
        public static new InternetservicedefinitionEntryArgs Empty => new InternetservicedefinitionEntryArgs();
    }
}
