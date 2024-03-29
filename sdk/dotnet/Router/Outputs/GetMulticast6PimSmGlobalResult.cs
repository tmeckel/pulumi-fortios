// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Router.Outputs
{

    [OutputType]
    public sealed class GetMulticast6PimSmGlobalResult
    {
        /// <summary>
        /// Limit of packets/sec per source registered through this RP (0 means unlimited).
        /// </summary>
        public readonly int RegisterRateLimit;
        /// <summary>
        /// Statically configured RP addresses. The structure of `rp_address` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetMulticast6PimSmGlobalRpAddressResult> RpAddresses;

        [OutputConstructor]
        private GetMulticast6PimSmGlobalResult(
            int registerRateLimit,

            ImmutableArray<Outputs.GetMulticast6PimSmGlobalRpAddressResult> rpAddresses)
        {
            RegisterRateLimit = registerRateLimit;
            RpAddresses = rpAddresses;
        }
    }
}
