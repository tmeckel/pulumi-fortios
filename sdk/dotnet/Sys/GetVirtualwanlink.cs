// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Sys
{
    public static class GetVirtualwanlink
    {
        /// <summary>
        /// Use this data source to get information on fortios system virtualwanlink
        /// </summary>
        public static Task<GetVirtualwanlinkResult> InvokeAsync(GetVirtualwanlinkArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVirtualwanlinkResult>("fortios:sys/getVirtualwanlink:getVirtualwanlink", args ?? new GetVirtualwanlinkArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on fortios system virtualwanlink
        /// </summary>
        public static Output<GetVirtualwanlinkResult> Invoke(GetVirtualwanlinkInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVirtualwanlinkResult>("fortios:sys/getVirtualwanlink:getVirtualwanlink", args ?? new GetVirtualwanlinkInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVirtualwanlinkArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetVirtualwanlinkArgs()
        {
        }
        public static new GetVirtualwanlinkArgs Empty => new GetVirtualwanlinkArgs();
    }

    public sealed class GetVirtualwanlinkInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetVirtualwanlinkInvokeArgs()
        {
        }
        public static new GetVirtualwanlinkInvokeArgs Empty => new GetVirtualwanlinkInvokeArgs();
    }


    [OutputType]
    public sealed class GetVirtualwanlinkResult
    {
        /// <summary>
        /// Physical interfaces that will be alerted. The structure of `fail_alert_interfaces` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetVirtualwanlinkFailAlertInterfaceResult> FailAlertInterfaces;
        /// <summary>
        /// Enable/disable SD-WAN Internet connection status checking (failure detection).
        /// </summary>
        public readonly string FailDetect;
        /// <summary>
        /// Virtual WAN Link health-check.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetVirtualwanlinkHealthCheckResult> HealthChecks;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Algorithm or mode to use for load balancing Internet traffic to SD-WAN members.
        /// </summary>
        public readonly string LoadBalanceMode;
        /// <summary>
        /// Member sequence number list. The structure of `members` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetVirtualwanlinkMemberResult> Members;
        /// <summary>
        /// Waiting period in seconds when switching from the primary neighbor to the secondary neighbor from the neighbor start. (0 - 10000000, default = 0).
        /// </summary>
        public readonly int NeighborHoldBootTime;
        /// <summary>
        /// Enable/disable hold switching from the secondary neighbor to the primary neighbor.
        /// </summary>
        public readonly string NeighborHoldDown;
        /// <summary>
        /// Waiting period in seconds when switching from the secondary neighbor to the primary neighbor when hold-down is disabled. (0 - 10000000, default = 0).
        /// </summary>
        public readonly int NeighborHoldDownTime;
        /// <summary>
        /// Create SD-WAN neighbor from BGP neighbor table to control route advertisements according to SLA status. The structure of `neighbor` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetVirtualwanlinkNeighborResult> Neighbors;
        /// <summary>
        /// Create SD-WAN rules or priority rules (also called services) to control how sessions are distributed to physical interfaces in the SD-WAN. The structure of `service` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetVirtualwanlinkServiceResult> Services;
        /// <summary>
        /// Enable/disable SD-WAN service.
        /// </summary>
        public readonly string Status;
        public readonly string? Vdomparam;
        /// <summary>
        /// Configure SD-WAN zones. The structure of `zone` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetVirtualwanlinkZoneResult> Zones;

        [OutputConstructor]
        private GetVirtualwanlinkResult(
            ImmutableArray<Outputs.GetVirtualwanlinkFailAlertInterfaceResult> failAlertInterfaces,

            string failDetect,

            ImmutableArray<Outputs.GetVirtualwanlinkHealthCheckResult> healthChecks,

            string id,

            string loadBalanceMode,

            ImmutableArray<Outputs.GetVirtualwanlinkMemberResult> members,

            int neighborHoldBootTime,

            string neighborHoldDown,

            int neighborHoldDownTime,

            ImmutableArray<Outputs.GetVirtualwanlinkNeighborResult> neighbors,

            ImmutableArray<Outputs.GetVirtualwanlinkServiceResult> services,

            string status,

            string? vdomparam,

            ImmutableArray<Outputs.GetVirtualwanlinkZoneResult> zones)
        {
            FailAlertInterfaces = failAlertInterfaces;
            FailDetect = failDetect;
            HealthChecks = healthChecks;
            Id = id;
            LoadBalanceMode = loadBalanceMode;
            Members = members;
            NeighborHoldBootTime = neighborHoldBootTime;
            NeighborHoldDown = neighborHoldDown;
            NeighborHoldDownTime = neighborHoldDownTime;
            Neighbors = neighbors;
            Services = services;
            Status = status;
            Vdomparam = vdomparam;
            Zones = zones;
        }
    }
}