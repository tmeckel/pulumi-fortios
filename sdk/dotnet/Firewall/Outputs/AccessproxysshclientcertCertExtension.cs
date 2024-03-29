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
    public sealed class AccessproxysshclientcertCertExtension
    {
        /// <summary>
        /// Critical option. Valid values: `no`, `yes`.
        /// </summary>
        public readonly string? Critical;
        /// <summary>
        /// Data of certificate extension.
        /// </summary>
        public readonly string? Data;
        /// <summary>
        /// Name of certificate extension.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Type of certificate extension. Valid values: `fixed`, `user`.
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private AccessproxysshclientcertCertExtension(
            string? critical,

            string? data,

            string? name,

            string? type)
        {
            Critical = critical;
            Data = data;
            Name = name;
            Type = type;
        }
    }
}
