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
    public sealed class GetAddrgrp6TaggingResult
    {
        /// <summary>
        /// Tag category.
        /// </summary>
        public readonly string Category;
        /// <summary>
        /// Specify the name of the desired firewall addrgrp6.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Tags. The structure of `tags` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAddrgrp6TaggingTagResult> Tags;

        [OutputConstructor]
        private GetAddrgrp6TaggingResult(
            string category,

            string name,

            ImmutableArray<Outputs.GetAddrgrp6TaggingTagResult> tags)
        {
            Category = category;
            Name = name;
            Tags = tags;
        }
    }
}
