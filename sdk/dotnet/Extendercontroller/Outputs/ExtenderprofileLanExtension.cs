// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Extendercontroller.Outputs
{

    [OutputType]
    public sealed class ExtenderprofileLanExtension
    {
        /// <summary>
        /// IPsec phase1 interface.
        /// </summary>
        public readonly string? BackhaulInterface;
        /// <summary>
        /// IPsec phase1 IPv4/FQDN. Used to specify the external IP/FQDN when the FortiGate unit is behind a NAT device.
        /// </summary>
        public readonly string? BackhaulIp;
        /// <summary>
        /// LAN extension backhaul tunnel configuration. The structure of `backhaul` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.ExtenderprofileLanExtensionBackhaul> Backhauls;
        /// <summary>
        /// IPsec tunnel name.
        /// </summary>
        public readonly string? IpsecTunnel;
        /// <summary>
        /// LAN extension link load balance strategy. Valid values: `activebackup`, `loadbalance`.
        /// </summary>
        public readonly string? LinkLoadbalance;

        [OutputConstructor]
        private ExtenderprofileLanExtension(
            string? backhaulInterface,

            string? backhaulIp,

            ImmutableArray<Outputs.ExtenderprofileLanExtensionBackhaul> backhauls,

            string? ipsecTunnel,

            string? linkLoadbalance)
        {
            BackhaulInterface = backhaulInterface;
            BackhaulIp = backhaulIp;
            Backhauls = backhauls;
            IpsecTunnel = ipsecTunnel;
            LinkLoadbalance = linkLoadbalance;
        }
    }
}
