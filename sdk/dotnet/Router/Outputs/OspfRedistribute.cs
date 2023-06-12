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
    public sealed class OspfRedistribute
    {
        /// <summary>
        /// Redistribute metric setting.
        /// </summary>
        public readonly int? Metric;
        /// <summary>
        /// Metric type. Valid values: `1`, `2`.
        /// </summary>
        public readonly string? MetricType;
        /// <summary>
        /// Redistribute name.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Route map name.
        /// </summary>
        public readonly string? Routemap;
        /// <summary>
        /// status Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? Status;
        /// <summary>
        /// Tag value.
        /// </summary>
        public readonly int? Tag;

        [OutputConstructor]
        private OspfRedistribute(
            int? metric,

            string? metricType,

            string? name,

            string? routemap,

            string? status,

            int? tag)
        {
            Metric = metric;
            MetricType = metricType;
            Name = name;
            Routemap = routemap;
            Status = status;
            Tag = tag;
        }
    }
}
