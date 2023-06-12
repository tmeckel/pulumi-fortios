// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Router.Outputs
{

    [OutputType]
    public sealed class GetRipInterfaceResult
    {
        /// <summary>
        /// Authentication key-chain name.
        /// </summary>
        public readonly string AuthKeychain;
        /// <summary>
        /// Authentication mode.
        /// </summary>
        public readonly string AuthMode;
        /// <summary>
        /// Authentication string/password.
        /// </summary>
        public readonly string AuthString;
        /// <summary>
        /// flags
        /// </summary>
        public readonly int Flags;
        /// <summary>
        /// Interface name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Receive version.
        /// </summary>
        public readonly string ReceiveVersion;
        /// <summary>
        /// Send version.
        /// </summary>
        public readonly string SendVersion;
        /// <summary>
        /// Enable/disable broadcast version 1 compatible packets.
        /// </summary>
        public readonly string SendVersion2Broadcast;
        /// <summary>
        /// Enable/disable split horizon.
        /// </summary>
        public readonly string SplitHorizon;
        /// <summary>
        /// Enable/disable split horizon.
        /// </summary>
        public readonly string SplitHorizonStatus;

        [OutputConstructor]
        private GetRipInterfaceResult(
            string authKeychain,

            string authMode,

            string authString,

            int flags,

            string name,

            string receiveVersion,

            string sendVersion,

            string sendVersion2Broadcast,

            string splitHorizon,

            string splitHorizonStatus)
        {
            AuthKeychain = authKeychain;
            AuthMode = authMode;
            AuthString = authString;
            Flags = flags;
            Name = name;
            ReceiveVersion = receiveVersion;
            SendVersion = sendVersion;
            SendVersion2Broadcast = sendVersion2Broadcast;
            SplitHorizon = splitHorizon;
            SplitHorizonStatus = splitHorizonStatus;
        }
    }
}