// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.User.Inputs
{

    public sealed class PeergrpMemberGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Peer group member name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public PeergrpMemberGetArgs()
        {
        }
        public static new PeergrpMemberGetArgs Empty => new PeergrpMemberGetArgs();
    }
}
