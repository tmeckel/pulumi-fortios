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
    public sealed class PtpServerInterface
    {
        /// <summary>
        /// End to end delay detection or peer to peer delay detection. Valid values: `E2E`, `P2P`.
        /// </summary>
        public readonly string? DelayMechanism;
        /// <summary>
        /// ID.
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// Interface name.
        /// </summary>
        public readonly string? ServerInterfaceName;

        [OutputConstructor]
        private PtpServerInterface(
            string? delayMechanism,

            int? id,

            string? serverInterfaceName)
        {
            DelayMechanism = delayMechanism;
            Id = id;
            ServerInterfaceName = serverInterfaceName;
        }
    }
}
