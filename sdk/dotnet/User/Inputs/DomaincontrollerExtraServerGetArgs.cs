// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.User.Inputs
{

    public sealed class DomaincontrollerExtraServerGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Server ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// Domain controller IP address.
        /// </summary>
        [Input("ipAddress")]
        public Input<string>? IpAddress { get; set; }

        /// <summary>
        /// Port to be used for communication with the domain controller (default = 445).
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// FortiGate IPv4 address to be used for communication with the domain controller.
        /// </summary>
        [Input("sourceIpAddress")]
        public Input<string>? SourceIpAddress { get; set; }

        /// <summary>
        /// Source port to be used for communication with the domain controller.
        /// </summary>
        [Input("sourcePort")]
        public Input<int>? SourcePort { get; set; }

        public DomaincontrollerExtraServerGetArgs()
        {
        }
        public static new DomaincontrollerExtraServerGetArgs Empty => new DomaincontrollerExtraServerGetArgs();
    }
}
