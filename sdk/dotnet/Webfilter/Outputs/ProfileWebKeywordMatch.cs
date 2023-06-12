// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Webfilter.Outputs
{

    [OutputType]
    public sealed class ProfileWebKeywordMatch
    {
        /// <summary>
        /// Pattern/keyword to search for.
        /// </summary>
        public readonly string? Pattern;

        [OutputConstructor]
        private ProfileWebKeywordMatch(string? pattern)
        {
            Pattern = pattern;
        }
    }
}
