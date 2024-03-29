// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Router.Outputs
{

    [OutputType]
    public sealed class RipngAggregateAddress
    {
        /// <summary>
        /// Aggregate address entry ID.
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// Aggregate address prefix.
        /// </summary>
        public readonly string? Prefix6;

        [OutputConstructor]
        private RipngAggregateAddress(
            int? id,

            string? prefix6)
        {
            Id = id;
            Prefix6 = prefix6;
        }
    }
}
