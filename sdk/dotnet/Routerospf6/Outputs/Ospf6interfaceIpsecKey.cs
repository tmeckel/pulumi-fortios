// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Routerospf6.Outputs
{

    [OutputType]
    public sealed class Ospf6interfaceIpsecKey
    {
        /// <summary>
        /// Authentication key.
        /// </summary>
        public readonly string? AuthKey;
        /// <summary>
        /// Encryption key.
        /// </summary>
        public readonly string? EncKey;
        /// <summary>
        /// Security Parameters Index.
        /// </summary>
        public readonly int? Spi;

        [OutputConstructor]
        private Ospf6interfaceIpsecKey(
            string? authKey,

            string? encKey,

            int? spi)
        {
            AuthKey = authKey;
            EncKey = encKey;
            Spi = spi;
        }
    }
}