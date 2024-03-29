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

    public sealed class OcvpnForticlientAccessAuthGroupGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Authentication user group for FortiClient access.
        /// </summary>
        [Input("authGroup")]
        public Input<string>? AuthGroup { get; set; }

        /// <summary>
        /// Group name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("overlays")]
        private InputList<Inputs.OcvpnForticlientAccessAuthGroupOverlayGetArgs>? _overlays;

        /// <summary>
        /// OCVPN overlays to allow access to. The structure of `overlays` block is documented below.
        /// </summary>
        public InputList<Inputs.OcvpnForticlientAccessAuthGroupOverlayGetArgs> Overlays
        {
            get => _overlays ?? (_overlays = new InputList<Inputs.OcvpnForticlientAccessAuthGroupOverlayGetArgs>());
            set => _overlays = value;
        }

        public OcvpnForticlientAccessAuthGroupGetArgs()
        {
        }
        public static new OcvpnForticlientAccessAuthGroupGetArgs Empty => new OcvpnForticlientAccessAuthGroupGetArgs();
    }
}
