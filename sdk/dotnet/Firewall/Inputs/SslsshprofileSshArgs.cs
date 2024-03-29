// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Firewall.Inputs
{

    public sealed class SslsshprofileSshArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Level of SSL inspection. Valid values: `disable`, `deep-inspection`.
        /// </summary>
        [Input("inspectAll")]
        public Input<string>? InspectAll { get; set; }

        /// <summary>
        /// Ports to use for scanning (1 - 65535, default = 443).
        /// </summary>
        [Input("ports")]
        public Input<string>? Ports { get; set; }

        /// <summary>
        /// Proxy traffic after the TCP 3-way handshake has been established (not before). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("proxyAfterTcpHandshake")]
        public Input<string>? ProxyAfterTcpHandshake { get; set; }

        /// <summary>
        /// Relative strength of encryption algorithms accepted during negotiation. Valid values: `compatible`, `high-encryption`.
        /// </summary>
        [Input("sshAlgorithm")]
        public Input<string>? SshAlgorithm { get; set; }

        /// <summary>
        /// Enable/disable SSH policy check. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("sshPolicyCheck")]
        public Input<string>? SshPolicyCheck { get; set; }

        /// <summary>
        /// Enable/disable SSH tunnel policy check. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("sshTunPolicyCheck")]
        public Input<string>? SshTunPolicyCheck { get; set; }

        /// <summary>
        /// Configure protocol inspection status. Valid values: `disable`, `deep-inspection`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Action based on SSH version being unsupported. Valid values: `bypass`, `block`.
        /// </summary>
        [Input("unsupportedVersion")]
        public Input<string>? UnsupportedVersion { get; set; }

        public SslsshprofileSshArgs()
        {
        }
        public static new SslsshprofileSshArgs Empty => new SslsshprofileSshArgs();
    }
}
