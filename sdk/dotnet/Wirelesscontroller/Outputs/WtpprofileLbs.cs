// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Wirelesscontroller.Outputs
{

    [OutputType]
    public sealed class WtpprofileLbs
    {
        /// <summary>
        /// Enable/disable AeroScout Real Time Location Service (RTLS) support. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? Aeroscout;
        /// <summary>
        /// Use BSSID or board MAC address as AP MAC address in the Aeroscout AP message. Valid values: `bssid`, `board-mac`.
        /// </summary>
        public readonly string? AeroscoutApMac;
        /// <summary>
        /// Enable/disable MU compounded report. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? AeroscoutMmuReport;
        /// <summary>
        /// Enable/disable AeroScout support. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? AeroscoutMu;
        /// <summary>
        /// AeroScout Mobile Unit (MU) mode dilution factor (default = 20).
        /// </summary>
        public readonly int? AeroscoutMuFactor;
        /// <summary>
        /// AeroScout MU mode timeout (0 - 65535 sec, default = 5).
        /// </summary>
        public readonly int? AeroscoutMuTimeout;
        /// <summary>
        /// IP address of AeroScout server.
        /// </summary>
        public readonly string? AeroscoutServerIp;
        /// <summary>
        /// AeroScout server UDP listening port.
        /// </summary>
        public readonly int? AeroscoutServerPort;
        /// <summary>
        /// Enable/disable Ekahua blink mode (also called AiRISTA Flow Blink Mode) to find the location of devices connected to a wireless LAN (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? EkahauBlinkMode;
        /// <summary>
        /// WiFi frame MAC address or WiFi Tag.
        /// </summary>
        public readonly string? EkahauTag;
        /// <summary>
        /// IP address of Ekahua RTLS Controller (ERC).
        /// </summary>
        public readonly string? ErcServerIp;
        /// <summary>
        /// Ekahua RTLS Controller (ERC) UDP listening port.
        /// </summary>
        public readonly int? ErcServerPort;
        /// <summary>
        /// Enable/disable FortiPresence to monitor the location and activity of WiFi clients even if they don't connect to this WiFi network (default = disable). Valid values: `foreign`, `both`, `disable`.
        /// </summary>
        public readonly string? Fortipresence;
        /// <summary>
        /// Enable/disable FortiPresence finding and reporting BLE devices. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? FortipresenceBle;
        /// <summary>
        /// FortiPresence report transmit frequency (5 - 65535 sec, default = 30).
        /// </summary>
        public readonly int? FortipresenceFrequency;
        /// <summary>
        /// FortiPresence server UDP listening port (default = 3000).
        /// </summary>
        public readonly int? FortipresencePort;
        /// <summary>
        /// FortiPresence project name (max. 16 characters, default = fortipresence).
        /// </summary>
        public readonly string? FortipresenceProject;
        /// <summary>
        /// Enable/disable FortiPresence finding and reporting rogue APs. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? FortipresenceRogue;
        /// <summary>
        /// FortiPresence secret password (max. 16 characters).
        /// </summary>
        public readonly string? FortipresenceSecret;
        /// <summary>
        /// FortiPresence server IP address.
        /// </summary>
        public readonly string? FortipresenceServer;
        /// <summary>
        /// FortiPresence server address type (default = ipv4). Valid values: `ipv4`, `fqdn`.
        /// </summary>
        public readonly string? FortipresenceServerAddrType;
        /// <summary>
        /// FQDN of FortiPresence server.
        /// </summary>
        public readonly string? FortipresenceServerFqdn;
        /// <summary>
        /// Enable/disable FortiPresence finding and reporting unassociated stations. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? FortipresenceUnassoc;
        /// <summary>
        /// Enable/disable client station locating services for all clients, whether associated or not (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? StationLocate;

        [OutputConstructor]
        private WtpprofileLbs(
            string? aeroscout,

            string? aeroscoutApMac,

            string? aeroscoutMmuReport,

            string? aeroscoutMu,

            int? aeroscoutMuFactor,

            int? aeroscoutMuTimeout,

            string? aeroscoutServerIp,

            int? aeroscoutServerPort,

            string? ekahauBlinkMode,

            string? ekahauTag,

            string? ercServerIp,

            int? ercServerPort,

            string? fortipresence,

            string? fortipresenceBle,

            int? fortipresenceFrequency,

            int? fortipresencePort,

            string? fortipresenceProject,

            string? fortipresenceRogue,

            string? fortipresenceSecret,

            string? fortipresenceServer,

            string? fortipresenceServerAddrType,

            string? fortipresenceServerFqdn,

            string? fortipresenceUnassoc,

            string? stationLocate)
        {
            Aeroscout = aeroscout;
            AeroscoutApMac = aeroscoutApMac;
            AeroscoutMmuReport = aeroscoutMmuReport;
            AeroscoutMu = aeroscoutMu;
            AeroscoutMuFactor = aeroscoutMuFactor;
            AeroscoutMuTimeout = aeroscoutMuTimeout;
            AeroscoutServerIp = aeroscoutServerIp;
            AeroscoutServerPort = aeroscoutServerPort;
            EkahauBlinkMode = ekahauBlinkMode;
            EkahauTag = ekahauTag;
            ErcServerIp = ercServerIp;
            ErcServerPort = ercServerPort;
            Fortipresence = fortipresence;
            FortipresenceBle = fortipresenceBle;
            FortipresenceFrequency = fortipresenceFrequency;
            FortipresencePort = fortipresencePort;
            FortipresenceProject = fortipresenceProject;
            FortipresenceRogue = fortipresenceRogue;
            FortipresenceSecret = fortipresenceSecret;
            FortipresenceServer = fortipresenceServer;
            FortipresenceServerAddrType = fortipresenceServerAddrType;
            FortipresenceServerFqdn = fortipresenceServerFqdn;
            FortipresenceUnassoc = fortipresenceUnassoc;
            StationLocate = stationLocate;
        }
    }
}
