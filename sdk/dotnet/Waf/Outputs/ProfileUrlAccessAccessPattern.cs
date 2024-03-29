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
    public sealed class ProfileUrlAccessAccessPattern
    {
        /// <summary>
        /// URL access pattern ID.
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// Enable/disable match negation. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? Negate;
        /// <summary>
        /// URL pattern.
        /// </summary>
        public readonly string? Pattern;
        /// <summary>
        /// Enable/disable regular expression based pattern match. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? Regex;
        /// <summary>
        /// Source address.
        /// </summary>
        public readonly string? Srcaddr;

        [OutputConstructor]
        private ProfileUrlAccessAccessPattern(
            int? id,

            string? negate,

            string? pattern,

            string? regex,

            string? srcaddr)
        {
            Id = id;
            Negate = negate;
            Pattern = pattern;
            Regex = regex;
            Srcaddr = srcaddr;
        }
    }
}
