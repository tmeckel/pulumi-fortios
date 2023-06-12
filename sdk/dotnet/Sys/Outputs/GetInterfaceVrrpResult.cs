// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Sys.Outputs
{

    [OutputType]
    public sealed class GetInterfaceVrrpResult
    {
        /// <summary>
        /// Enable/disable accept mode.
        /// </summary>
        public readonly string AcceptMode;
        /// <summary>
        /// Advertisement interval (1 - 255 seconds).
        /// </summary>
        public readonly int AdvInterval;
        /// <summary>
        /// Enable/disable ignoring of default route when checking destination.
        /// </summary>
        public readonly string IgnoreDefaultRoute;
        /// <summary>
        /// Enable/disable preempt mode.
        /// </summary>
        public readonly string Preempt;
        /// <summary>
        /// Priority of the virtual router (1 - 255).
        /// </summary>
        public readonly int Priority;
        /// <summary>
        /// VRRP Proxy ARP configuration. The structure of `proxy_arp` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInterfaceVrrpProxyArpResult> ProxyArps;
        /// <summary>
        /// Startup time (1 - 255 seconds).
        /// </summary>
        public readonly int StartTime;
        /// <summary>
        /// Enable/disable VRRP.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// VRRP version.
        /// </summary>
        public readonly string Version;
        /// <summary>
        /// Monitor the route to this destination.
        /// </summary>
        public readonly string Vrdst;
        /// <summary>
        /// Priority of the virtual router when the virtual router destination becomes unreachable (0 - 254).
        /// </summary>
        public readonly int VrdstPriority;
        /// <summary>
        /// VRRP group ID (1 - 65535).
        /// </summary>
        public readonly int Vrgrp;
        /// <summary>
        /// Virtual router identifier (1 - 255).
        /// </summary>
        public readonly int Vrid;
        /// <summary>
        /// IP address of the virtual router.
        /// </summary>
        public readonly string Vrip;

        [OutputConstructor]
        private GetInterfaceVrrpResult(
            string acceptMode,

            int advInterval,

            string ignoreDefaultRoute,

            string preempt,

            int priority,

            ImmutableArray<Outputs.GetInterfaceVrrpProxyArpResult> proxyArps,

            int startTime,

            string status,

            string version,

            string vrdst,

            int vrdstPriority,

            int vrgrp,

            int vrid,

            string vrip)
        {
            AcceptMode = acceptMode;
            AdvInterval = advInterval;
            IgnoreDefaultRoute = ignoreDefaultRoute;
            Preempt = preempt;
            Priority = priority;
            ProxyArps = proxyArps;
            StartTime = startTime;
            Status = status;
            Version = version;
            Vrdst = vrdst;
            VrdstPriority = vrdstPriority;
            Vrgrp = vrgrp;
            Vrid = vrid;
            Vrip = vrip;
        }
    }
}
