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

    public sealed class FortilinksettingsNacPortsNacSegmentVlanGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// VLAN interface name.
        /// </summary>
        [Input("vlanName")]
        public Input<string>? VlanName { get; set; }

        public FortilinksettingsNacPortsNacSegmentVlanGetArgs()
        {
        }
        public static new FortilinksettingsNacPortsNacSegmentVlanGetArgs Empty => new FortilinksettingsNacPortsNacSegmentVlanGetArgs();
    }
}
