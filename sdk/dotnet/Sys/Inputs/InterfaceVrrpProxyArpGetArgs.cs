// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Sys.Inputs
{

    public sealed class InterfaceVrrpProxyArpGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// Set IP addresses of proxy ARP.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        public InterfaceVrrpProxyArpGetArgs()
        {
        }
        public static new InterfaceVrrpProxyArpGetArgs Empty => new InterfaceVrrpProxyArpGetArgs();
    }
}
