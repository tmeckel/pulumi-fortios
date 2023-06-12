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

    public sealed class Address6templateSubnetSegmentGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Number of bits.
        /// </summary>
        [Input("bits")]
        public Input<int>? Bits { get; set; }

        /// <summary>
        /// Enable/disable exclusive value. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("exclusive")]
        public Input<string>? Exclusive { get; set; }

        /// <summary>
        /// Subnet segment ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// Subnet segment name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("values")]
        private InputList<Inputs.Address6templateSubnetSegmentValueGetArgs>? _values;

        /// <summary>
        /// Subnet segment values. The structure of `values` block is documented below.
        /// </summary>
        public InputList<Inputs.Address6templateSubnetSegmentValueGetArgs> Values
        {
            get => _values ?? (_values = new InputList<Inputs.Address6templateSubnetSegmentValueGetArgs>());
            set => _values = value;
        }

        public Address6templateSubnetSegmentGetArgs()
        {
        }
        public static new Address6templateSubnetSegmentGetArgs Empty => new Address6templateSubnetSegmentGetArgs();
    }
}
