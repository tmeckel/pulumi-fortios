// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Sys.Outputs
{

    [OutputType]
    public sealed class HaUnicastPeer
    {
        public readonly int? Id;
        public readonly string? PeerIp;

        [OutputConstructor]
        private HaUnicastPeer(
            int? id,

            string? peerIp)
        {
            Id = id;
            PeerIp = peerIp;
        }
    }
}
