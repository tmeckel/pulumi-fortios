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

    public sealed class VipRealserverGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Dynamic address of the real server.
        /// </summary>
        [Input("address")]
        public Input<string>? Address { get; set; }

        /// <summary>
        /// Only clients in this IP range can connect to this real server.
        /// </summary>
        [Input("clientIp")]
        public Input<string>? ClientIp { get; set; }

        /// <summary>
        /// Enable to check the responsiveness of the real server before forwarding traffic. Valid values: `disable`, `enable`, `vip`.
        /// </summary>
        [Input("healthcheck")]
        public Input<string>? Healthcheck { get; set; }

        /// <summary>
        /// Time in seconds that the health check monitor continues to monitor and unresponsive server that should be active.
        /// </summary>
        [Input("holddownInterval")]
        public Input<int>? HolddownInterval { get; set; }

        /// <summary>
        /// HTTP server domain name in HTTP header.
        /// </summary>
        [Input("httpHost")]
        public Input<string>? HttpHost { get; set; }

        /// <summary>
        /// Real server ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// IP address of the real server.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// Max number of active connections that can be directed to the real server. When reached, sessions are sent to other real servers.
        /// </summary>
        [Input("maxConnections")]
        public Input<int>? MaxConnections { get; set; }

        /// <summary>
        /// Name of the health check monitor to use when polling to determine a virtual server's connectivity status.
        /// </summary>
        [Input("monitor")]
        public Input<string>? Monitor { get; set; }

        /// <summary>
        /// Port for communicating with the real server. Required if port forwarding is enabled.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// Set the status of the real server to active so that it can accept traffic, or on standby or disabled so no traffic is sent. Valid values: `active`, `standby`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Type of address. Valid values: `ip`, `address`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Weight of the real server. If weighted load balancing is enabled, the server with the highest weight gets more connections.
        /// </summary>
        [Input("weight")]
        public Input<int>? Weight { get; set; }

        public VipRealserverGetArgs()
        {
        }
        public static new VipRealserverGetArgs Empty => new VipRealserverGetArgs();
    }
}
