// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Firewall.Inputs
{

    public sealed class ShapingprofileShapingEntryGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Number of bytes that can be burst at maximum-bandwidth speed. Formula: burst = maximum-bandwidth*burst-in-msec.
        /// </summary>
        [Input("burstInMsec")]
        public Input<int>? BurstInMsec { get; set; }

        /// <summary>
        /// Number of bytes that can be burst as fast as the interface can transmit. Formula: cburst = maximum-bandwidth*cburst-in-msec.
        /// </summary>
        [Input("cburstInMsec")]
        public Input<int>? CburstInMsec { get; set; }

        /// <summary>
        /// Class ID.
        /// </summary>
        [Input("classId")]
        public Input<int>? ClassId { get; set; }

        /// <summary>
        /// Guaranteed bandwith in percentage.
        /// </summary>
        [Input("guaranteedBandwidthPercentage")]
        public Input<int>? GuaranteedBandwidthPercentage { get; set; }

        /// <summary>
        /// ID number.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// Hard limit on the real queue size in packets.
        /// </summary>
        [Input("limit")]
        public Input<int>? Limit { get; set; }

        /// <summary>
        /// Average queue size in packets at which RED drop probability is maximal.
        /// </summary>
        [Input("max")]
        public Input<int>? Max { get; set; }

        /// <summary>
        /// Maximum bandwith in percentage.
        /// </summary>
        [Input("maximumBandwidthPercentage")]
        public Input<int>? MaximumBandwidthPercentage { get; set; }

        /// <summary>
        /// Average queue size in packets at which RED drop becomes a possibility.
        /// </summary>
        [Input("min")]
        public Input<int>? Min { get; set; }

        /// <summary>
        /// Priority.
        /// </summary>
        [Input("priority")]
        public Input<string>? Priority { get; set; }

        /// <summary>
        /// Maximum probability (in percentage) for RED marking.
        /// </summary>
        [Input("redProbability")]
        public Input<int>? RedProbability { get; set; }

        public ShapingprofileShapingEntryGetArgs()
        {
        }
        public static new ShapingprofileShapingEntryGetArgs Empty => new ShapingprofileShapingEntryGetArgs();
    }
}