// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Router
{
    public static class GetSetting
    {
        /// <summary>
        /// Use this data source to get information on fortios router setting
        /// </summary>
        public static Task<GetSettingResult> InvokeAsync(GetSettingArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSettingResult>("fortios:router/getSetting:getSetting", args ?? new GetSettingArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on fortios router setting
        /// </summary>
        public static Output<GetSettingResult> Invoke(GetSettingInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSettingResult>("fortios:router/getSetting:getSetting", args ?? new GetSettingInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSettingArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetSettingArgs()
        {
        }
        public static new GetSettingArgs Empty => new GetSettingArgs();
    }

    public sealed class GetSettingInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetSettingInvokeArgs()
        {
        }
        public static new GetSettingInvokeArgs Empty => new GetSettingInvokeArgs();
    }


    [OutputType]
    public sealed class GetSettingResult
    {
        /// <summary>
        /// bgp_debug_flags
        /// </summary>
        public readonly string BgpDebugFlags;
        /// <summary>
        /// Hostname for this virtual domain router.
        /// </summary>
        public readonly string Hostname;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// igmp_debug_flags
        /// </summary>
        public readonly string IgmpDebugFlags;
        /// <summary>
        /// imi_debug_flags
        /// </summary>
        public readonly string ImiDebugFlags;
        /// <summary>
        /// isis_debug_flags
        /// </summary>
        public readonly string IsisDebugFlags;
        /// <summary>
        /// ospf6_debug_events_flags
        /// </summary>
        public readonly string Ospf6DebugEventsFlags;
        /// <summary>
        /// ospf6_debug_ifsm_flags
        /// </summary>
        public readonly string Ospf6DebugIfsmFlags;
        /// <summary>
        /// ospf6_debug_lsa_flags
        /// </summary>
        public readonly string Ospf6DebugLsaFlags;
        /// <summary>
        /// ospf6_debug_nfsm_flags
        /// </summary>
        public readonly string Ospf6DebugNfsmFlags;
        /// <summary>
        /// ospf6_debug_nsm_flags
        /// </summary>
        public readonly string Ospf6DebugNsmFlags;
        /// <summary>
        /// ospf6_debug_packet_flags
        /// </summary>
        public readonly string Ospf6DebugPacketFlags;
        /// <summary>
        /// ospf6_debug_route_flags
        /// </summary>
        public readonly string Ospf6DebugRouteFlags;
        /// <summary>
        /// ospf_debug_events_flags
        /// </summary>
        public readonly string OspfDebugEventsFlags;
        /// <summary>
        /// ospf_debug_ifsm_flags
        /// </summary>
        public readonly string OspfDebugIfsmFlags;
        /// <summary>
        /// ospf_debug_lsa_flags
        /// </summary>
        public readonly string OspfDebugLsaFlags;
        /// <summary>
        /// ospf_debug_nfsm_flags
        /// </summary>
        public readonly string OspfDebugNfsmFlags;
        /// <summary>
        /// ospf_debug_nsm_flags
        /// </summary>
        public readonly string OspfDebugNsmFlags;
        /// <summary>
        /// ospf_debug_packet_flags
        /// </summary>
        public readonly string OspfDebugPacketFlags;
        /// <summary>
        /// ospf_debug_route_flags
        /// </summary>
        public readonly string OspfDebugRouteFlags;
        /// <summary>
        /// pimdm_debug_flags
        /// </summary>
        public readonly string PimdmDebugFlags;
        /// <summary>
        /// pimsm_debug_joinprune_flags
        /// </summary>
        public readonly string PimsmDebugJoinpruneFlags;
        /// <summary>
        /// pimsm_debug_simple_flags
        /// </summary>
        public readonly string PimsmDebugSimpleFlags;
        /// <summary>
        /// pimsm_debug_timer_flags
        /// </summary>
        public readonly string PimsmDebugTimerFlags;
        /// <summary>
        /// rip_debug_flags
        /// </summary>
        public readonly string RipDebugFlags;
        /// <summary>
        /// ripng_debug_flags
        /// </summary>
        public readonly string RipngDebugFlags;
        /// <summary>
        /// Prefix-list as filter for showing routes.
        /// </summary>
        public readonly string ShowFilter;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetSettingResult(
            string bgpDebugFlags,

            string hostname,

            string id,

            string igmpDebugFlags,

            string imiDebugFlags,

            string isisDebugFlags,

            string ospf6DebugEventsFlags,

            string ospf6DebugIfsmFlags,

            string ospf6DebugLsaFlags,

            string ospf6DebugNfsmFlags,

            string ospf6DebugNsmFlags,

            string ospf6DebugPacketFlags,

            string ospf6DebugRouteFlags,

            string ospfDebugEventsFlags,

            string ospfDebugIfsmFlags,

            string ospfDebugLsaFlags,

            string ospfDebugNfsmFlags,

            string ospfDebugNsmFlags,

            string ospfDebugPacketFlags,

            string ospfDebugRouteFlags,

            string pimdmDebugFlags,

            string pimsmDebugJoinpruneFlags,

            string pimsmDebugSimpleFlags,

            string pimsmDebugTimerFlags,

            string ripDebugFlags,

            string ripngDebugFlags,

            string showFilter,

            string? vdomparam)
        {
            BgpDebugFlags = bgpDebugFlags;
            Hostname = hostname;
            Id = id;
            IgmpDebugFlags = igmpDebugFlags;
            ImiDebugFlags = imiDebugFlags;
            IsisDebugFlags = isisDebugFlags;
            Ospf6DebugEventsFlags = ospf6DebugEventsFlags;
            Ospf6DebugIfsmFlags = ospf6DebugIfsmFlags;
            Ospf6DebugLsaFlags = ospf6DebugLsaFlags;
            Ospf6DebugNfsmFlags = ospf6DebugNfsmFlags;
            Ospf6DebugNsmFlags = ospf6DebugNsmFlags;
            Ospf6DebugPacketFlags = ospf6DebugPacketFlags;
            Ospf6DebugRouteFlags = ospf6DebugRouteFlags;
            OspfDebugEventsFlags = ospfDebugEventsFlags;
            OspfDebugIfsmFlags = ospfDebugIfsmFlags;
            OspfDebugLsaFlags = ospfDebugLsaFlags;
            OspfDebugNfsmFlags = ospfDebugNfsmFlags;
            OspfDebugNsmFlags = ospfDebugNsmFlags;
            OspfDebugPacketFlags = ospfDebugPacketFlags;
            OspfDebugRouteFlags = ospfDebugRouteFlags;
            PimdmDebugFlags = pimdmDebugFlags;
            PimsmDebugJoinpruneFlags = pimsmDebugJoinpruneFlags;
            PimsmDebugSimpleFlags = pimsmDebugSimpleFlags;
            PimsmDebugTimerFlags = pimsmDebugTimerFlags;
            RipDebugFlags = ripDebugFlags;
            RipngDebugFlags = ripngDebugFlags;
            ShowFilter = showFilter;
            Vdomparam = vdomparam;
        }
    }
}
