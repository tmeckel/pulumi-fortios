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

    public sealed class Ospf6AreaVirtualLinkArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Authentication mode. Valid values: `none`, `ah`, `esp`, `area`.
        /// </summary>
        [Input("authentication")]
        public Input<string>? Authentication { get; set; }

        /// <summary>
        /// Dead interval.
        /// </summary>
        [Input("deadInterval")]
        public Input<int>? DeadInterval { get; set; }

        /// <summary>
        /// Hello interval.
        /// </summary>
        [Input("helloInterval")]
        public Input<int>? HelloInterval { get; set; }

        /// <summary>
        /// Authentication algorithm. Valid values: `md5`, `sha1`, `sha256`, `sha384`, `sha512`.
        /// </summary>
        [Input("ipsecAuthAlg")]
        public Input<string>? IpsecAuthAlg { get; set; }

        /// <summary>
        /// Encryption algorithm. Valid values: `null`, `des`, `3des`, `aes128`, `aes192`, `aes256`.
        /// </summary>
        [Input("ipsecEncAlg")]
        public Input<string>? IpsecEncAlg { get; set; }

        [Input("ipsecKeys")]
        private InputList<Inputs.Ospf6AreaVirtualLinkIpsecKeyArgs>? _ipsecKeys;

        /// <summary>
        /// IPsec authentication and encryption keys. The structure of `ipsec_keys` block is documented below.
        /// </summary>
        public InputList<Inputs.Ospf6AreaVirtualLinkIpsecKeyArgs> IpsecKeys
        {
            get => _ipsecKeys ?? (_ipsecKeys = new InputList<Inputs.Ospf6AreaVirtualLinkIpsecKeyArgs>());
            set => _ipsecKeys = value;
        }

        /// <summary>
        /// Key roll-over interval.
        /// </summary>
        [Input("keyRolloverInterval")]
        public Input<int>? KeyRolloverInterval { get; set; }

        /// <summary>
        /// Virtual link entry name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A.B.C.D, peer router ID.
        /// </summary>
        [Input("peer")]
        public Input<string>? Peer { get; set; }

        /// <summary>
        /// Retransmit interval.
        /// </summary>
        [Input("retransmitInterval")]
        public Input<int>? RetransmitInterval { get; set; }

        /// <summary>
        /// Transmit delay.
        /// </summary>
        [Input("transmitDelay")]
        public Input<int>? TransmitDelay { get; set; }

        public Ospf6AreaVirtualLinkArgs()
        {
        }
        public static new Ospf6AreaVirtualLinkArgs Empty => new Ospf6AreaVirtualLinkArgs();
    }
}
