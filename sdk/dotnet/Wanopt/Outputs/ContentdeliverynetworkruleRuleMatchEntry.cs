// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Wanopt.Outputs
{

    [OutputType]
    public sealed class ContentdeliverynetworkruleRuleMatchEntry
    {
        /// <summary>
        /// Rule ID.
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// Pattern string for matching target (Referrer or URL pattern, eg. "a", "a*c", "*a*", "a*c*e", and "*"). The structure of `pattern` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.ContentdeliverynetworkruleRuleMatchEntryPattern> Patterns;
        /// <summary>
        /// Option in HTTP header or URL parameter to match. Valid values: `path`, `parameter`, `referrer`, `youtube-map`, `youtube-id`, `youku-id`.
        /// </summary>
        public readonly string? Target;

        [OutputConstructor]
        private ContentdeliverynetworkruleRuleMatchEntry(
            int? id,

            ImmutableArray<Outputs.ContentdeliverynetworkruleRuleMatchEntryPattern> patterns,

            string? target)
        {
            Id = id;
            Patterns = patterns;
            Target = target;
        }
    }
}
