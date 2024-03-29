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
    public sealed class GetRipNetworkResult
    {
        /// <summary>
        /// Offset-list ID.
        /// </summary>
        public readonly int Id;
        /// <summary>
        /// Network prefix.
        /// </summary>
        public readonly string Prefix;

        [OutputConstructor]
        private GetRipNetworkResult(
            int id,

            string prefix)
        {
            Id = id;
            Prefix = prefix;
        }
    }
}
