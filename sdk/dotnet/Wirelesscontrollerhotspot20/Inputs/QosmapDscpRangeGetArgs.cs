// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Wirelesscontrollerhotspot20.Inputs
{

    public sealed class QosmapDscpRangeGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// DSCP high value.
        /// </summary>
        [Input("high")]
        public Input<int>? High { get; set; }

        /// <summary>
        /// DSCP range index.
        /// </summary>
        [Input("index")]
        public Input<int>? Index { get; set; }

        /// <summary>
        /// DSCP low value.
        /// </summary>
        [Input("low")]
        public Input<int>? Low { get; set; }

        /// <summary>
        /// User priority.
        /// </summary>
        [Input("up")]
        public Input<int>? Up { get; set; }

        public QosmapDscpRangeGetArgs()
        {
        }
        public static new QosmapDscpRangeGetArgs Empty => new QosmapDscpRangeGetArgs();
    }
}