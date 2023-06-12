// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Sys.Outputs
{

    [OutputType]
    public sealed class ClustersyncSessionSyncFilterCustomService
    {
        /// <summary>
        /// Custom service destination port range.
        /// </summary>
        public readonly string? DstPortRange;
        /// <summary>
        /// Custom service ID.
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// Custom service source port range.
        /// </summary>
        public readonly string? SrcPortRange;

        [OutputConstructor]
        private ClustersyncSessionSyncFilterCustomService(
            string? dstPortRange,

            int? id,

            string? srcPortRange)
        {
            DstPortRange = dstPortRange;
            Id = id;
            SrcPortRange = srcPortRange;
        }
    }
}
