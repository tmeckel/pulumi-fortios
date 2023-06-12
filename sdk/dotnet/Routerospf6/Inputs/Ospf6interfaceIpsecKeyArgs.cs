// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Routerospf6.Inputs
{

    public sealed class Ospf6interfaceIpsecKeyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Authentication key.
        /// </summary>
        [Input("authKey")]
        public Input<string>? AuthKey { get; set; }

        /// <summary>
        /// Encryption key.
        /// </summary>
        [Input("encKey")]
        public Input<string>? EncKey { get; set; }

        /// <summary>
        /// Security Parameters Index.
        /// </summary>
        [Input("spi")]
        public Input<int>? Spi { get; set; }

        public Ospf6interfaceIpsecKeyArgs()
        {
        }
        public static new Ospf6interfaceIpsecKeyArgs Empty => new Ospf6interfaceIpsecKeyArgs();
    }
}