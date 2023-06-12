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

    public sealed class SdnconnectorNicGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("ips")]
        private InputList<Inputs.SdnconnectorNicIpGetArgs>? _ips;

        /// <summary>
        /// Configure IP configuration. The structure of `ip` block is documented below.
        /// </summary>
        public InputList<Inputs.SdnconnectorNicIpGetArgs> Ips
        {
            get => _ips ?? (_ips = new InputList<Inputs.SdnconnectorNicIpGetArgs>());
            set => _ips = value;
        }

        /// <summary>
        /// Network interface name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public SdnconnectorNicGetArgs()
        {
        }
        public static new SdnconnectorNicGetArgs Empty => new SdnconnectorNicGetArgs();
    }
}
