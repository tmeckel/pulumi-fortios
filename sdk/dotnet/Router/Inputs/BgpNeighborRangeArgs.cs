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

    public sealed class BgpNeighborRangeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// IPv6 neighbor range ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// Maximum number of neighbors.
        /// </summary>
        [Input("maxNeighborNum")]
        public Input<int>? MaxNeighborNum { get; set; }

        /// <summary>
        /// Neighbor group name.
        /// </summary>
        [Input("neighborGroup")]
        public Input<string>? NeighborGroup { get; set; }

        /// <summary>
        /// Neighbor range prefix.
        /// </summary>
        [Input("prefix")]
        public Input<string>? Prefix { get; set; }

        public BgpNeighborRangeArgs()
        {
        }
        public static new BgpNeighborRangeArgs Empty => new BgpNeighborRangeArgs();
    }
}
