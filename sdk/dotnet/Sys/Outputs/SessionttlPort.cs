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
    public sealed class SessionttlPort
    {
        /// <summary>
        /// End port number.
        /// </summary>
        public readonly int? EndPort;
        /// <summary>
        /// Table entry ID.
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// Protocol (0 - 255).
        /// </summary>
        public readonly int? Protocol;
        /// <summary>
        /// Start port number.
        /// </summary>
        public readonly int? StartPort;
        /// <summary>
        /// Session timeout (TTL).
        /// </summary>
        public readonly string? Timeout;

        [OutputConstructor]
        private SessionttlPort(
            int? endPort,

            int? id,

            int? protocol,

            int? startPort,

            string? timeout)
        {
            EndPort = endPort;
            Id = id;
            Protocol = protocol;
            StartPort = startPort;
            Timeout = timeout;
        }
    }
}
