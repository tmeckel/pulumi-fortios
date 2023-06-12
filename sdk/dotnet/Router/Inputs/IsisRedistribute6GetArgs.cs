// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Router.Inputs
{

    public sealed class IsisRedistribute6GetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Level. Valid values: `level-1-2`, `level-1`, `level-2`.
        /// 
        /// The `summary_address6` block supports:
        /// </summary>
        [Input("level")]
        public Input<string>? Level { get; set; }

        /// <summary>
        /// Metric.
        /// </summary>
        [Input("metric")]
        public Input<int>? Metric { get; set; }

        /// <summary>
        /// Metric type. Valid values: `external`, `internal`.
        /// </summary>
        [Input("metricType")]
        public Input<string>? MetricType { get; set; }

        /// <summary>
        /// Protocol name.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// Route map name.
        /// 
        /// The `redistribute6` block supports:
        /// </summary>
        [Input("routemap")]
        public Input<string>? Routemap { get; set; }

        /// <summary>
        /// Enable/disable interface for IS-IS. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public IsisRedistribute6GetArgs()
        {
        }
        public static new IsisRedistribute6GetArgs Empty => new IsisRedistribute6GetArgs();
    }
}
