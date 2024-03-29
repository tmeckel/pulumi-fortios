// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Router.Inputs
{

    public sealed class Multicast6PimSmGlobalArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Limit of packets/sec per source registered through this RP (0 means unlimited).
        /// </summary>
        [Input("registerRateLimit")]
        public Input<int>? RegisterRateLimit { get; set; }

        [Input("rpAddresses")]
        private InputList<Inputs.Multicast6PimSmGlobalRpAddressArgs>? _rpAddresses;

        /// <summary>
        /// Statically configured RP addresses. The structure of `rp_address` block is documented below.
        /// </summary>
        public InputList<Inputs.Multicast6PimSmGlobalRpAddressArgs> RpAddresses
        {
            get => _rpAddresses ?? (_rpAddresses = new InputList<Inputs.Multicast6PimSmGlobalRpAddressArgs>());
            set => _rpAddresses = value;
        }

        public Multicast6PimSmGlobalArgs()
        {
        }
        public static new Multicast6PimSmGlobalArgs Empty => new Multicast6PimSmGlobalArgs();
    }
}
