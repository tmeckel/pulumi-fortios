// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Wirelesscontroller.Outputs
{

    [OutputType]
    public sealed class IntercontrollerInterControllerPeer
    {
        /// <summary>
        /// ID.
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// Peer wireless controller's IP address.
        /// </summary>
        public readonly string? PeerIp;
        /// <summary>
        /// Port used by the wireless controller's for inter-controller communications (1024 - 49150, default = 5246).
        /// </summary>
        public readonly int? PeerPort;
        /// <summary>
        /// Peer wireless controller's priority (primary or secondary, default = primary). Valid values: `primary`, `secondary`.
        /// </summary>
        public readonly string? PeerPriority;

        [OutputConstructor]
        private IntercontrollerInterControllerPeer(
            int? id,

            string? peerIp,

            int? peerPort,

            string? peerPriority)
        {
            Id = id;
            PeerIp = peerIp;
            PeerPort = peerPort;
            PeerPriority = peerPriority;
        }
    }
}
