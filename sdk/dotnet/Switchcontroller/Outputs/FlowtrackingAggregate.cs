// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Switchcontroller.Outputs
{

    [OutputType]
    public sealed class FlowtrackingAggregate
    {
        /// <summary>
        /// Aggregate id.
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// IP address to group all matching traffic sessions to a flow.
        /// </summary>
        public readonly string? Ip;

        [OutputConstructor]
        private FlowtrackingAggregate(
            int? id,

            string? ip)
        {
            Id = id;
            Ip = ip;
        }
    }
}
