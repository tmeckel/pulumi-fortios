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
    public sealed class NtpNtpserver
    {
        /// <summary>
        /// Enable/disable MD5/SHA1 authentication. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? Authentication;
        /// <summary>
        /// NTP server ID.
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// Specify outgoing interface to reach server.
        /// </summary>
        public readonly string? Interface;
        /// <summary>
        /// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
        /// </summary>
        public readonly string? InterfaceSelectMethod;
        /// <summary>
        /// Key for MD5/SHA1 authentication.
        /// </summary>
        public readonly string? Key;
        /// <summary>
        /// Key ID for authentication.
        /// </summary>
        public readonly int? KeyId;
        /// <summary>
        /// Enable to use NTPv3 instead of NTPv4. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? Ntpv3;
        /// <summary>
        /// IP address or hostname of the NTP Server.
        /// </summary>
        public readonly string? Server;

        [OutputConstructor]
        private NtpNtpserver(
            string? authentication,

            int? id,

            string? @interface,

            string? interfaceSelectMethod,

            string? key,

            int? keyId,

            string? ntpv3,

            string? server)
        {
            Authentication = authentication;
            Id = id;
            Interface = @interface;
            InterfaceSelectMethod = interfaceSelectMethod;
            Key = key;
            KeyId = keyId;
            Ntpv3 = ntpv3;
            Server = server;
        }
    }
}
