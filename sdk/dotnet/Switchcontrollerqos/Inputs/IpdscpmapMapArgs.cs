// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Switchcontrollerqos.Inputs
{

    public sealed class IpdscpmapMapArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// COS queue number.
        /// </summary>
        [Input("cosQueue")]
        public Input<int>? CosQueue { get; set; }

        /// <summary>
        /// Differentiated service. Valid values: `CS0`, `CS1`, `AF11`, `AF12`, `AF13`, `CS2`, `AF21`, `AF22`, `AF23`, `CS3`, `AF31`, `AF32`, `AF33`, `CS4`, `AF41`, `AF42`, `AF43`, `CS5`, `EF`, `CS6`, `CS7`.
        /// </summary>
        [Input("diffserv")]
        public Input<string>? Diffserv { get; set; }

        /// <summary>
        /// IP Precedence. Valid values: `network-control`, `internetwork-control`, `critic-ecp`, `flashoverride`, `flash`, `immediate`, `priority`, `routine`.
        /// </summary>
        [Input("ipPrecedence")]
        public Input<string>? IpPrecedence { get; set; }

        /// <summary>
        /// Dscp mapping entry name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Raw values of DSCP (0 - 63).
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public IpdscpmapMapArgs()
        {
        }
        public static new IpdscpmapMapArgs Empty => new IpdscpmapMapArgs();
    }
}
