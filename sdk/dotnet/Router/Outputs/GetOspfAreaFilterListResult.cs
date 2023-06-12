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
    public sealed class GetOspfAreaFilterListResult
    {
        /// <summary>
        /// Direction.
        /// </summary>
        public readonly string Direction;
        /// <summary>
        /// Distribute list entry ID.
        /// </summary>
        public readonly int Id;
        /// <summary>
        /// Access-list or prefix-list name.
        /// </summary>
        public readonly string List;

        [OutputConstructor]
        private GetOspfAreaFilterListResult(
            string direction,

            int id,

            string list)
        {
            Direction = direction;
            Id = id;
            List = list;
        }
    }
}
