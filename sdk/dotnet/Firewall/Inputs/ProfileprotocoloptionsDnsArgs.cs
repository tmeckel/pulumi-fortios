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

    public sealed class ProfileprotocoloptionsDnsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Ports to scan for content (1 - 65535, default = 53).
        /// </summary>
        [Input("ports")]
        public Input<int>? Ports { get; set; }

        /// <summary>
        /// Enable/disable the active status of scanning for this protocol. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public ProfileprotocoloptionsDnsArgs()
        {
        }
        public static new ProfileprotocoloptionsDnsArgs Empty => new ProfileprotocoloptionsDnsArgs();
    }
}
