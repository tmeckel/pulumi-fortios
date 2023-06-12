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

    public sealed class SdwanHealthCheckArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Address mode (IPv4 or IPv6). Valid values: `ipv4`, `ipv6`.
        /// </summary>
        [Input("addrMode")]
        public Input<string>? AddrMode { get; set; }

        /// <summary>
        /// The mode determining how to detect the server.
        /// </summary>
        [Input("detectMode")]
        public Input<string>? DetectMode { get; set; }

        /// <summary>
        /// Differentiated services code point (DSCP) in the IP header of the probe packet.
        /// </summary>
        [Input("diffservcode")]
        public Input<string>? Diffservcode { get; set; }

        /// <summary>
        /// Response IP expected from DNS server if the protocol is DNS.
        /// </summary>
        [Input("dnsMatchIp")]
        public Input<string>? DnsMatchIp { get; set; }

        /// <summary>
        /// Fully qualified domain name to resolve for the DNS probe.
        /// </summary>
        [Input("dnsRequestDomain")]
        public Input<string>? DnsRequestDomain { get; set; }

        /// <summary>
        /// Enable/disable embedding measured health information. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("embedMeasuredHealth")]
        public Input<string>? EmbedMeasuredHealth { get; set; }

        /// <summary>
        /// Number of failures before server is considered lost (1 - 3600, default = 5).
        /// </summary>
        [Input("failtime")]
        public Input<int>? Failtime { get; set; }

        /// <summary>
        /// Full path and file name on the FTP server to download for FTP health-check to probe.
        /// </summary>
        [Input("ftpFile")]
        public Input<string>? FtpFile { get; set; }

        /// <summary>
        /// FTP mode. Valid values: `passive`, `port`.
        /// </summary>
        [Input("ftpMode")]
        public Input<string>? FtpMode { get; set; }

        /// <summary>
        /// HA election priority (1 - 50).
        /// </summary>
        [Input("haPriority")]
        public Input<int>? HaPriority { get; set; }

        /// <summary>
        /// String in the http-agent field in the HTTP header.
        /// </summary>
        [Input("httpAgent")]
        public Input<string>? HttpAgent { get; set; }

        /// <summary>
        /// URL used to communicate with the server if the protocol if the protocol is HTTP.
        /// </summary>
        [Input("httpGet")]
        public Input<string>? HttpGet { get; set; }

        /// <summary>
        /// Response string expected from the server if the protocol is HTTP.
        /// </summary>
        [Input("httpMatch")]
        public Input<string>? HttpMatch { get; set; }

        /// <summary>
        /// Status check interval in milliseconds, or the time between attempting to connect to the server (500 - 3600*1000 msec, default = 500).
        /// </summary>
        [Input("interval")]
        public Input<int>? Interval { get; set; }

        [Input("members")]
        private InputList<Inputs.SdwanHealthCheckMemberArgs>? _members;

        /// <summary>
        /// Member sequence number list. The structure of `members` block is documented below.
        /// </summary>
        public InputList<Inputs.SdwanHealthCheckMemberArgs> Members
        {
            get => _members ?? (_members = new InputList<Inputs.SdwanHealthCheckMemberArgs>());
            set => _members = value;
        }

        /// <summary>
        /// Codec to use for MOS calculation (default = g711). Valid values: `g711`, `g722`, `g729`.
        /// </summary>
        [Input("mosCodec")]
        public Input<string>? MosCodec { get; set; }

        /// <summary>
        /// Health check name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Packet size of a twamp test session,
        /// </summary>
        [Input("packetSize")]
        public Input<int>? PacketSize { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// Twamp controller password in authentication mode
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Port number used to communicate with the server over the selected protocol (0-65535, default = 0, auto select. http, twamp: 80, udp-echo, tcp-echo: 7, dns: 53, ftp: 21).
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// Number of most recent probes that should be used to calculate latency and jitter (5 - 30, default = 30).
        /// </summary>
        [Input("probeCount")]
        public Input<int>? ProbeCount { get; set; }

        /// <summary>
        /// Enable/disable transmission of probe packets. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("probePackets")]
        public Input<string>? ProbePackets { get; set; }

        /// <summary>
        /// Time to wait before a probe packet is considered lost (500 - 3600*1000 msec, default = 500).
        /// </summary>
        [Input("probeTimeout")]
        public Input<int>? ProbeTimeout { get; set; }

        /// <summary>
        /// Protocol used to determine if the FortiGate can communicate with the server.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// Method to measure the quality of tcp-connect. Valid values: `half-open`, `half-close`.
        /// </summary>
        [Input("qualityMeasuredMethod")]
        public Input<string>? QualityMeasuredMethod { get; set; }

        /// <summary>
        /// Number of successful responses received before server is considered recovered (1 - 3600, default = 5).
        /// </summary>
        [Input("recoverytime")]
        public Input<int>? Recoverytime { get; set; }

        /// <summary>
        /// Twamp controller security mode. Valid values: `none`, `authentication`.
        /// </summary>
        [Input("securityMode")]
        public Input<string>? SecurityMode { get; set; }

        /// <summary>
        /// IP address or FQDN name of the server.
        /// </summary>
        [Input("server")]
        public Input<string>? Server { get; set; }

        /// <summary>
        /// Time interval in seconds that SLA fail log messages will be generated (0 - 3600, default = 0).
        /// </summary>
        [Input("slaFailLogPeriod")]
        public Input<int>? SlaFailLogPeriod { get; set; }

        /// <summary>
        /// Select the ID from the SLA sub-table. The selected SLA's priority value will be distributed into the routing table (0 - 32, default = 0).
        /// </summary>
        [Input("slaIdRedistribute")]
        public Input<int>? SlaIdRedistribute { get; set; }

        /// <summary>
        /// Time interval in seconds that SLA pass log messages will be generated (0 - 3600, default = 0).
        /// </summary>
        [Input("slaPassLogPeriod")]
        public Input<int>? SlaPassLogPeriod { get; set; }

        [Input("slas")]
        private InputList<Inputs.SdwanHealthCheckSlaArgs>? _slas;

        /// <summary>
        /// Service level agreement (SLA). The structure of `sla` block is documented below.
        /// </summary>
        public InputList<Inputs.SdwanHealthCheckSlaArgs> Slas
        {
            get => _slas ?? (_slas = new InputList<Inputs.SdwanHealthCheckSlaArgs>());
            set => _slas = value;
        }

        /// <summary>
        /// Source IP address used in the health-check packet to the server.
        /// </summary>
        [Input("source")]
        public Input<string>? Source { get; set; }

        /// <summary>
        /// Enable/disable system DNS as the probe server. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("systemDns")]
        public Input<string>? SystemDns { get; set; }

        /// <summary>
        /// Alert threshold for jitter (ms, default = 0).
        /// </summary>
        [Input("thresholdAlertJitter")]
        public Input<int>? ThresholdAlertJitter { get; set; }

        /// <summary>
        /// Alert threshold for latency (ms, default = 0).
        /// </summary>
        [Input("thresholdAlertLatency")]
        public Input<int>? ThresholdAlertLatency { get; set; }

        /// <summary>
        /// Alert threshold for packet loss (percentage, default = 0).
        /// </summary>
        [Input("thresholdAlertPacketloss")]
        public Input<int>? ThresholdAlertPacketloss { get; set; }

        /// <summary>
        /// Warning threshold for jitter (ms, default = 0).
        /// </summary>
        [Input("thresholdWarningJitter")]
        public Input<int>? ThresholdWarningJitter { get; set; }

        /// <summary>
        /// Warning threshold for latency (ms, default = 0).
        /// </summary>
        [Input("thresholdWarningLatency")]
        public Input<int>? ThresholdWarningLatency { get; set; }

        /// <summary>
        /// Warning threshold for packet loss (percentage, default = 0).
        /// </summary>
        [Input("thresholdWarningPacketloss")]
        public Input<int>? ThresholdWarningPacketloss { get; set; }

        /// <summary>
        /// Enable/disable update cascade interface. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("updateCascadeInterface")]
        public Input<string>? UpdateCascadeInterface { get; set; }

        /// <summary>
        /// Enable/disable updating the static route. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("updateStaticRoute")]
        public Input<string>? UpdateStaticRoute { get; set; }

        /// <summary>
        /// The user name to access probe server.
        /// </summary>
        [Input("user")]
        public Input<string>? User { get; set; }

        /// <summary>
        /// Virtual Routing Forwarding ID.
        /// </summary>
        [Input("vrf")]
        public Input<int>? Vrf { get; set; }

        public SdwanHealthCheckArgs()
        {
        }
        public static new SdwanHealthCheckArgs Empty => new SdwanHealthCheckArgs();
    }
}
