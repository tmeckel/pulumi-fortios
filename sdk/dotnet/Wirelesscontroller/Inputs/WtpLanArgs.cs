// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Wirelesscontroller.Inputs
{

    public sealed class WtpLanArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// LAN port 1 mode. Valid values: `offline`, `nat-to-wan`, `bridge-to-wan`, `bridge-to-ssid`.
        /// </summary>
        [Input("port1Mode")]
        public Input<string>? Port1Mode { get; set; }

        /// <summary>
        /// Bridge LAN port 1 to SSID.
        /// </summary>
        [Input("port1Ssid")]
        public Input<string>? Port1Ssid { get; set; }

        /// <summary>
        /// LAN port 2 mode. Valid values: `offline`, `nat-to-wan`, `bridge-to-wan`, `bridge-to-ssid`.
        /// </summary>
        [Input("port2Mode")]
        public Input<string>? Port2Mode { get; set; }

        /// <summary>
        /// Bridge LAN port 2 to SSID.
        /// </summary>
        [Input("port2Ssid")]
        public Input<string>? Port2Ssid { get; set; }

        /// <summary>
        /// LAN port 3 mode. Valid values: `offline`, `nat-to-wan`, `bridge-to-wan`, `bridge-to-ssid`.
        /// </summary>
        [Input("port3Mode")]
        public Input<string>? Port3Mode { get; set; }

        /// <summary>
        /// Bridge LAN port 3 to SSID.
        /// </summary>
        [Input("port3Ssid")]
        public Input<string>? Port3Ssid { get; set; }

        /// <summary>
        /// LAN port 4 mode. Valid values: `offline`, `nat-to-wan`, `bridge-to-wan`, `bridge-to-ssid`.
        /// </summary>
        [Input("port4Mode")]
        public Input<string>? Port4Mode { get; set; }

        /// <summary>
        /// Bridge LAN port 4 to SSID.
        /// </summary>
        [Input("port4Ssid")]
        public Input<string>? Port4Ssid { get; set; }

        /// <summary>
        /// LAN port 5 mode. Valid values: `offline`, `nat-to-wan`, `bridge-to-wan`, `bridge-to-ssid`.
        /// </summary>
        [Input("port5Mode")]
        public Input<string>? Port5Mode { get; set; }

        /// <summary>
        /// Bridge LAN port 5 to SSID.
        /// </summary>
        [Input("port5Ssid")]
        public Input<string>? Port5Ssid { get; set; }

        /// <summary>
        /// LAN port 6 mode. Valid values: `offline`, `nat-to-wan`, `bridge-to-wan`, `bridge-to-ssid`.
        /// </summary>
        [Input("port6Mode")]
        public Input<string>? Port6Mode { get; set; }

        /// <summary>
        /// Bridge LAN port 6 to SSID.
        /// </summary>
        [Input("port6Ssid")]
        public Input<string>? Port6Ssid { get; set; }

        /// <summary>
        /// LAN port 7 mode. Valid values: `offline`, `nat-to-wan`, `bridge-to-wan`, `bridge-to-ssid`.
        /// </summary>
        [Input("port7Mode")]
        public Input<string>? Port7Mode { get; set; }

        /// <summary>
        /// Bridge LAN port 7 to SSID.
        /// </summary>
        [Input("port7Ssid")]
        public Input<string>? Port7Ssid { get; set; }

        /// <summary>
        /// LAN port 8 mode. Valid values: `offline`, `nat-to-wan`, `bridge-to-wan`, `bridge-to-ssid`.
        /// </summary>
        [Input("port8Mode")]
        public Input<string>? Port8Mode { get; set; }

        /// <summary>
        /// Bridge LAN port 8 to SSID.
        /// </summary>
        [Input("port8Ssid")]
        public Input<string>? Port8Ssid { get; set; }

        /// <summary>
        /// ESL port mode. Valid values: `offline`, `nat-to-wan`, `bridge-to-wan`, `bridge-to-ssid`.
        /// </summary>
        [Input("portEslMode")]
        public Input<string>? PortEslMode { get; set; }

        /// <summary>
        /// Bridge ESL port to SSID.
        /// 
        /// The `radio_1` block supports:
        /// </summary>
        [Input("portEslSsid")]
        public Input<string>? PortEslSsid { get; set; }

        /// <summary>
        /// LAN port mode. Valid values: `offline`, `nat-to-wan`, `bridge-to-wan`, `bridge-to-ssid`.
        /// </summary>
        [Input("portMode")]
        public Input<string>? PortMode { get; set; }

        /// <summary>
        /// Bridge LAN port to SSID.
        /// </summary>
        [Input("portSsid")]
        public Input<string>? PortSsid { get; set; }

        public WtpLanArgs()
        {
        }
        public static new WtpLanArgs Empty => new WtpLanArgs();
    }
}
