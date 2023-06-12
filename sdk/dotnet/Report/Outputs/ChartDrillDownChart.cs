// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Report.Outputs
{

    [OutputType]
    public sealed class ChartDrillDownChart
    {
        /// <summary>
        /// Drill down chart name.
        /// </summary>
        public readonly string? ChartName;
        /// <summary>
        /// Drill down chart ID.
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// Enable/disable this drill down chart. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? Status;

        [OutputConstructor]
        private ChartDrillDownChart(
            string? chartName,

            int? id,

            string? status)
        {
            ChartName = chartName;
            Id = id;
            Status = status;
        }
    }
}
