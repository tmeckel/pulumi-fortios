// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Switchcontroller.Inputs
{

    public sealed class LldpprofileMedNetworkPolicyGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable VLAN assignment when this profile is applied on managed FortiSwitch port. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("assignVlan")]
        public Input<string>? AssignVlan { get; set; }

        /// <summary>
        /// Advertised Differentiated Services Code Point (DSCP) value, a packet header value indicating the level of service requested for traffic, such as high priority or best effort delivery.
        /// </summary>
        [Input("dscp")]
        public Input<int>? Dscp { get; set; }

        /// <summary>
        /// Policy type name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Advertised Layer 2 priority (0 - 7; from lowest to highest priority).
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// Enable or disable this TLV. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// ID of VLAN to advertise, if configured on port (0 - 4094, 0 = priority tag).
        /// </summary>
        [Input("vlan")]
        public Input<int>? Vlan { get; set; }

        /// <summary>
        /// VLAN interface to advertise; if configured on port.
        /// </summary>
        [Input("vlanIntf")]
        public Input<string>? VlanIntf { get; set; }

        public LldpprofileMedNetworkPolicyGetArgs()
        {
        }
        public static new LldpprofileMedNetworkPolicyGetArgs Empty => new LldpprofileMedNetworkPolicyGetArgs();
    }
}
