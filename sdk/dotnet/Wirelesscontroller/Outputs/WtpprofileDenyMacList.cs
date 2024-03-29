// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Wirelesscontroller.Outputs
{

    [OutputType]
    public sealed class WtpprofileDenyMacList
    {
        /// <summary>
        /// ID.
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// A WiFi device with this MAC address is denied access to this WTP, FortiAP or AP.
        /// </summary>
        public readonly string? Mac;

        [OutputConstructor]
        private WtpprofileDenyMacList(
            int? id,

            string? mac)
        {
            Id = id;
            Mac = mac;
        }
    }
}
