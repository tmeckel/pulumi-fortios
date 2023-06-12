// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Vpn.Inputs
{

    public sealed class OcvpnOverlayGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable client address assignment. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("assignIp")]
        public Input<string>? AssignIp { get; set; }

        /// <summary>
        /// ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// Allow or deny traffic from other overlays. Valid values: `allow`, `deny`.
        /// </summary>
        [Input("interOverlay")]
        public Input<string>? InterOverlay { get; set; }

        /// <summary>
        /// End of client IPv4 range.
        /// </summary>
        [Input("ipv4EndIp")]
        public Input<string>? Ipv4EndIp { get; set; }

        /// <summary>
        /// Start of client IPv4 range.
        /// </summary>
        [Input("ipv4StartIp")]
        public Input<string>? Ipv4StartIp { get; set; }

        /// <summary>
        /// Overlay name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Overlay name.
        /// </summary>
        [Input("overlayName")]
        public Input<string>? OverlayName { get; set; }

        [Input("subnets")]
        private InputList<Inputs.OcvpnOverlaySubnetGetArgs>? _subnets;

        /// <summary>
        /// Internal subnets to register with OCVPN service. The structure of `subnets` block is documented below.
        /// </summary>
        public InputList<Inputs.OcvpnOverlaySubnetGetArgs> Subnets
        {
            get => _subnets ?? (_subnets = new InputList<Inputs.OcvpnOverlaySubnetGetArgs>());
            set => _subnets = value;
        }

        public OcvpnOverlayGetArgs()
        {
        }
        public static new OcvpnOverlayGetArgs Empty => new OcvpnOverlayGetArgs();
    }
}
