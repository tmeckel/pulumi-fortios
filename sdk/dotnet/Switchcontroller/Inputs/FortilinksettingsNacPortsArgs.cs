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

    public sealed class FortilinksettingsNacPortsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable bouncing (administratively bring the link down, up) of a switch port when NAC mode is configured on the port. Helps to re-initiate the DHCP process for a device. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("bounceNacPort")]
        public Input<string>? BounceNacPort { get; set; }

        /// <summary>
        /// Enable/disable LAN segment feature on the FortiLink interface. Valid values: `enabled`, `disabled`.
        /// </summary>
        [Input("lanSegment")]
        public Input<string>? LanSegment { get; set; }

        /// <summary>
        /// Member change flag.
        /// </summary>
        [Input("memberChange")]
        public Input<int>? MemberChange { get; set; }

        /// <summary>
        /// Configure NAC LAN interface.
        /// </summary>
        [Input("nacLanInterface")]
        public Input<string>? NacLanInterface { get; set; }

        [Input("nacSegmentVlans")]
        private InputList<Inputs.FortilinksettingsNacPortsNacSegmentVlanArgs>? _nacSegmentVlans;

        /// <summary>
        /// Configure NAC segment VLANs. The structure of `nac_segment_vlans` block is documented below.
        /// </summary>
        public InputList<Inputs.FortilinksettingsNacPortsNacSegmentVlanArgs> NacSegmentVlans
        {
            get => _nacSegmentVlans ?? (_nacSegmentVlans = new InputList<Inputs.FortilinksettingsNacPortsNacSegmentVlanArgs>());
            set => _nacSegmentVlans = value;
        }

        /// <summary>
        /// Default NAC Onboarding VLAN when NAC devices are discovered.
        /// </summary>
        [Input("onboardingVlan")]
        public Input<string>? OnboardingVlan { get; set; }

        /// <summary>
        /// Parent key name.
        /// </summary>
        [Input("parentKey")]
        public Input<string>? ParentKey { get; set; }

        public FortilinksettingsNacPortsArgs()
        {
        }
        public static new FortilinksettingsNacPortsArgs Empty => new FortilinksettingsNacPortsArgs();
    }
}
