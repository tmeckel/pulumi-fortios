// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Firewall.Outputs
{

    [OutputType]
    public sealed class GetProfileprotocoloptionsCifResult
    {
        /// <summary>
        /// Domain for which to decrypt CIFS traffic.
        /// </summary>
        public readonly string DomainController;
        /// <summary>
        /// One or more options that can be applied to the session.
        /// </summary>
        public readonly string Options;
        /// <summary>
        /// Maximum in-memory file size that can be scanned (1 - 383 MB, default = 10).
        /// </summary>
        public readonly int OversizeLimit;
        /// <summary>
        /// Ports to scan for content (1 - 65535, default = 445).
        /// </summary>
        public readonly int Ports;
        /// <summary>
        /// Enable/disable scanning of BZip2 compressed files.
        /// </summary>
        public readonly string ScanBzip2;
        /// <summary>
        /// CIFS server credential type.
        /// </summary>
        public readonly string ServerCredentialType;
        /// <summary>
        /// Server keytab. The structure of `server_keytab` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetProfileprotocoloptionsCifServerKeytabResult> ServerKeytabs;
        /// <summary>
        /// Enable/disable adding an email signature to SMTP email messages as they pass through the FortiGate.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// Maximum dynamic TCP window size (default = 8MB).
        /// </summary>
        public readonly int TcpWindowMaximum;
        /// <summary>
        /// Minimum dynamic TCP window size (default = 128KB).
        /// </summary>
        public readonly int TcpWindowMinimum;
        /// <summary>
        /// Set TCP static window size (default = 256KB).
        /// </summary>
        public readonly int TcpWindowSize;
        /// <summary>
        /// Specify type of TCP window to use for this protocol.
        /// </summary>
        public readonly string TcpWindowType;
        /// <summary>
        /// Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).
        /// </summary>
        public readonly int UncompressedNestLimit;
        /// <summary>
        /// Maximum in-memory uncompressed file size that can be scanned (0 - 383 MB, 0 = unlimited, default = 10).
        /// </summary>
        public readonly int UncompressedOversizeLimit;

        [OutputConstructor]
        private GetProfileprotocoloptionsCifResult(
            string domainController,

            string options,

            int oversizeLimit,

            int ports,

            string scanBzip2,

            string serverCredentialType,

            ImmutableArray<Outputs.GetProfileprotocoloptionsCifServerKeytabResult> serverKeytabs,

            string status,

            int tcpWindowMaximum,

            int tcpWindowMinimum,

            int tcpWindowSize,

            string tcpWindowType,

            int uncompressedNestLimit,

            int uncompressedOversizeLimit)
        {
            DomainController = domainController;
            Options = options;
            OversizeLimit = oversizeLimit;
            Ports = ports;
            ScanBzip2 = scanBzip2;
            ServerCredentialType = serverCredentialType;
            ServerKeytabs = serverKeytabs;
            Status = status;
            TcpWindowMaximum = tcpWindowMaximum;
            TcpWindowMinimum = tcpWindowMinimum;
            TcpWindowSize = tcpWindowSize;
            TcpWindowType = tcpWindowType;
            UncompressedNestLimit = uncompressedNestLimit;
            UncompressedOversizeLimit = uncompressedOversizeLimit;
        }
    }
}
