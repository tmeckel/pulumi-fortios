// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Switchcontroller.Inputs
{

    public sealed class FlowtrackingCollectorArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Collector IP address.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// Collector name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Collector port number(0-65535, default:0, netflow:2055, ipfix:4739).
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// Collector L4 transport protocol for exporting packets. Valid values: `udp`, `tcp`, `sctp`.
        /// </summary>
        [Input("transport")]
        public Input<string>? Transport { get; set; }

        public FlowtrackingCollectorArgs()
        {
        }
        public static new FlowtrackingCollectorArgs Empty => new FlowtrackingCollectorArgs();
    }
}
