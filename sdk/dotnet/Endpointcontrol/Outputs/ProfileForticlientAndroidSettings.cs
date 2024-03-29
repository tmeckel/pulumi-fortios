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
    public sealed class ProfileForticlientAndroidSettings
    {
        /// <summary>
        /// Enable/disable FortiClient web category filtering when protected by FortiGate. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? DisableWfWhenProtected;
        /// <summary>
        /// Enable/disable advanced FortiClient VPN configuration. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? ForticlientAdvancedVpn;
        /// <summary>
        /// Advanced FortiClient VPN configuration.
        /// </summary>
        public readonly string? ForticlientAdvancedVpnBuffer;
        /// <summary>
        /// Enable/disable FortiClient VPN provisioning. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? ForticlientVpnProvisioning;
        /// <summary>
        /// FortiClient VPN settings. The structure of `forticlient_vpn_settings` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.ProfileForticlientAndroidSettingsForticlientVpnSetting> ForticlientVpnSettings;
        /// <summary>
        /// Enable/disable FortiClient web filtering. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? ForticlientWf;
        /// <summary>
        /// The FortiClient web filter profile to apply.
        /// </summary>
        public readonly string? ForticlientWfProfile;

        [OutputConstructor]
        private ProfileForticlientAndroidSettings(
            string? disableWfWhenProtected,

            string? forticlientAdvancedVpn,

            string? forticlientAdvancedVpnBuffer,

            string? forticlientVpnProvisioning,

            ImmutableArray<Outputs.ProfileForticlientAndroidSettingsForticlientVpnSetting> forticlientVpnSettings,

            string? forticlientWf,

            string? forticlientWfProfile)
        {
            DisableWfWhenProtected = disableWfWhenProtected;
            ForticlientAdvancedVpn = forticlientAdvancedVpn;
            ForticlientAdvancedVpnBuffer = forticlientAdvancedVpnBuffer;
            ForticlientVpnProvisioning = forticlientVpnProvisioning;
            ForticlientVpnSettings = forticlientVpnSettings;
            ForticlientWf = forticlientWf;
            ForticlientWfProfile = forticlientWfProfile;
        }
    }
}
