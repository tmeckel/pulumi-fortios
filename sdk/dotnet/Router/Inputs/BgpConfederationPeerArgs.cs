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

    public sealed class BgpConfederationPeerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Peer ID.
        /// </summary>
        [Input("peer")]
        public Input<string>? Peer { get; set; }

        public BgpConfederationPeerArgs()
        {
        }
        public static new BgpConfederationPeerArgs Empty => new BgpConfederationPeerArgs();
    }
}
