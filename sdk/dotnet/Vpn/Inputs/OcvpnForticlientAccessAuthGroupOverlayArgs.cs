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

    public sealed class OcvpnForticlientAccessAuthGroupOverlayArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Overlay name.
        /// </summary>
        [Input("overlayName")]
        public Input<string>? OverlayName { get; set; }

        public OcvpnForticlientAccessAuthGroupOverlayArgs()
        {
        }
        public static new OcvpnForticlientAccessAuthGroupOverlayArgs Empty => new OcvpnForticlientAccessAuthGroupOverlayArgs();
    }
}
