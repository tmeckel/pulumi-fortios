// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Vpnsslweb.Inputs
{

    public sealed class PortalSplitDnGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// DNS server 1.
        /// </summary>
        [Input("dnsServer1")]
        public Input<string>? DnsServer1 { get; set; }

        /// <summary>
        /// DNS server 2.
        /// </summary>
        [Input("dnsServer2")]
        public Input<string>? DnsServer2 { get; set; }

        /// <summary>
        /// Split DNS domains used for SSL-VPN clients separated by comma(,).
        /// </summary>
        [Input("domains")]
        public Input<string>? Domains { get; set; }

        /// <summary>
        /// ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// IPv6 DNS server 1.
        /// </summary>
        [Input("ipv6DnsServer1")]
        public Input<string>? Ipv6DnsServer1 { get; set; }

        /// <summary>
        /// IPv6 DNS server 2.
        /// </summary>
        [Input("ipv6DnsServer2")]
        public Input<string>? Ipv6DnsServer2 { get; set; }

        public PortalSplitDnGetArgs()
        {
        }
        public static new PortalSplitDnGetArgs Empty => new PortalSplitDnGetArgs();
    }
}
