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

    public sealed class OspfNeighborArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cost of the interface, value range from 0 to 65535, 0 means auto-cost.
        /// </summary>
        [Input("cost")]
        public Input<int>? Cost { get; set; }

        /// <summary>
        /// Neighbor entry ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// Interface IP address of the neighbor.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// Poll interval time in seconds.
        /// </summary>
        [Input("pollInterval")]
        public Input<int>? PollInterval { get; set; }

        /// <summary>
        /// Priority.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        public OspfNeighborArgs()
        {
        }
        public static new OspfNeighborArgs Empty => new OspfNeighborArgs();
    }
}
