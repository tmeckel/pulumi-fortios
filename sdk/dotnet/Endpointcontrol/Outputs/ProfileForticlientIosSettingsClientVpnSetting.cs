// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Endpointcontrol.Outputs
{

    [OutputType]
    public sealed class ProfileForticlientIosSettingsClientVpnSetting
    {
        /// <summary>
        /// Authentication method. Valid values: `psk`, `certificate`.
        /// </summary>
        public readonly string? AuthMethod;
        /// <summary>
        /// VPN name.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Pre-shared secret for PSK authentication.
        /// </summary>
        public readonly string? PresharedKey;
        /// <summary>
        /// IP address or FQDN of the remote VPN gateway.
        /// </summary>
        public readonly string? RemoteGw;
        /// <summary>
        /// SSL VPN access port (1 - 65535).
        /// </summary>
        public readonly int? SslvpnAccessPort;
        /// <summary>
        /// Enable/disable requiring SSL VPN client certificate. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? SslvpnRequireCertificate;
        /// <summary>
        /// VPN type (IPsec or SSL VPN). Valid values: `ipsec`, `ssl`.
        /// </summary>
        public readonly string? Type;
        /// <summary>
        /// Content of VPN configuration.
        /// </summary>
        public readonly string? VpnConfigurationContent;
        /// <summary>
        /// Name of VPN configuration.
        /// </summary>
        public readonly string? VpnConfigurationName;

        [OutputConstructor]
        private ProfileForticlientIosSettingsClientVpnSetting(
            string? authMethod,

            string? name,

            string? presharedKey,

            string? remoteGw,

            int? sslvpnAccessPort,

            string? sslvpnRequireCertificate,

            string? type,

            string? vpnConfigurationContent,

            string? vpnConfigurationName)
        {
            AuthMethod = authMethod;
            Name = name;
            PresharedKey = presharedKey;
            RemoteGw = remoteGw;
            SslvpnAccessPort = sslvpnAccessPort;
            SslvpnRequireCertificate = sslvpnRequireCertificate;
            Type = type;
            VpnConfigurationContent = vpnConfigurationContent;
            VpnConfigurationName = vpnConfigurationName;
        }
    }
}
