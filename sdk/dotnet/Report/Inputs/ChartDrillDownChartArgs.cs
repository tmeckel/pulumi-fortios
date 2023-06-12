// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Report.Inputs
{

    public sealed class ChartDrillDownChartArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Drill down chart name.
        /// </summary>
        [Input("chartName")]
        public Input<string>? ChartName { get; set; }

        /// <summary>
        /// Drill down chart ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// Enable/disable this drill down chart. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public ChartDrillDownChartArgs()
        {
        }
        public static new ChartDrillDownChartArgs Empty => new ChartDrillDownChartArgs();
    }
}
