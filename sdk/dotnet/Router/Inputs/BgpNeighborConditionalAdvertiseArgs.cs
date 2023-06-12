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

    public sealed class BgpNeighborConditionalAdvertiseArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of advertising route map.
        /// </summary>
        [Input("advertiseRoutemap")]
        public Input<string>? AdvertiseRoutemap { get; set; }

        /// <summary>
        /// Name of condition route map.
        /// </summary>
        [Input("conditionRoutemap")]
        public Input<string>? ConditionRoutemap { get; set; }

        /// <summary>
        /// Type of condition. Valid values: `exist`, `non-exist`.
        /// </summary>
        [Input("conditionType")]
        public Input<string>? ConditionType { get; set; }

        public BgpNeighborConditionalAdvertiseArgs()
        {
        }
        public static new BgpNeighborConditionalAdvertiseArgs Empty => new BgpNeighborConditionalAdvertiseArgs();
    }
}
