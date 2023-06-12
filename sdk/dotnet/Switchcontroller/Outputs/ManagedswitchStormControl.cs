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
    public sealed class ManagedswitchStormControl
    {
        /// <summary>
        /// Enable/disable storm control to drop broadcast traffic. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? Broadcast;
        /// <summary>
        /// Enable to override global FortiSwitch storm control settings for this FortiSwitch. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? LocalOverride;
        /// <summary>
        /// Rate in packets per second at which storm traffic is controlled (1 - 10000000, default = 500). Storm control drops excess traffic data rates beyond this threshold.
        /// </summary>
        public readonly int? Rate;
        /// <summary>
        /// Enable/disable storm control to drop unknown multicast traffic. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? UnknownMulticast;
        /// <summary>
        /// Enable/disable storm control to drop unknown unicast traffic. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? UnknownUnicast;

        [OutputConstructor]
        private ManagedswitchStormControl(
            string? broadcast,

            string? localOverride,

            int? rate,

            string? unknownMulticast,

            string? unknownUnicast)
        {
            Broadcast = broadcast;
            LocalOverride = localOverride;
            Rate = rate;
            UnknownMulticast = unknownMulticast;
            UnknownUnicast = unknownUnicast;
        }
    }
}
