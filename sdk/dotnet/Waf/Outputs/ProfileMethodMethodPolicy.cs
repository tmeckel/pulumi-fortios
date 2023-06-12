// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Waf.Outputs
{

    [OutputType]
    public sealed class ProfileMethodMethodPolicy
    {
        /// <summary>
        /// Host address.
        /// </summary>
        public readonly string? Address;
        /// <summary>
        /// Allowed Methods. Valid values: `get`, `post`, `put`, `head`, `connect`, `trace`, `options`, `delete`, `others`.
        /// </summary>
        public readonly string? AllowedMethods;
        /// <summary>
        /// HTTP method policy ID.
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// URL pattern.
        /// </summary>
        public readonly string? Pattern;
        /// <summary>
        /// Enable/disable regular expression based pattern match. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? Regex;

        [OutputConstructor]
        private ProfileMethodMethodPolicy(
            string? address,

            string? allowedMethods,

            int? id,

            string? pattern,

            string? regex)
        {
            Address = address;
            AllowedMethods = allowedMethods;
            Id = id;
            Pattern = pattern;
            Regex = regex;
        }
    }
}