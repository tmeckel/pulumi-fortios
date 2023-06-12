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
    public sealed class GetVirtualwanlinkServiceResult
    {
        /// <summary>
        /// Address mode (IPv4 or IPv6).
        /// </summary>
        public readonly string AddrMode;
        /// <summary>
        /// Coefficient of reciprocal of available bidirectional bandwidth in the formula of custom-profile-1.
        /// </summary>
        public readonly int BandwidthWeight;
        /// <summary>
        /// Enable/disable use of SD-WAN as default service.
        /// </summary>
        public readonly string Default;
        /// <summary>
        /// Enable/disable forward traffic DSCP tag.
        /// </summary>
        public readonly string DscpForward;
        /// <summary>
        /// Forward traffic DSCP tag.
        /// </summary>
        public readonly string DscpForwardTag;
        /// <summary>
        /// Enable/disable reverse traffic DSCP tag.
        /// </summary>
        public readonly string DscpReverse;
        /// <summary>
        /// Reverse traffic DSCP tag.
        /// </summary>
        public readonly string DscpReverseTag;
        /// <summary>
        /// Destination address6 name. The structure of `dst6` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetVirtualwanlinkServiceDst6Result> Dst6s;
        /// <summary>
        /// Enable/disable negation of destination address match.
        /// </summary>
        public readonly string DstNegate;
        /// <summary>
        /// Destination address name. The structure of `dst` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetVirtualwanlinkServiceDstResult> Dsts;
        /// <summary>
        /// End destination port number.
        /// </summary>
        public readonly int EndPort;
        /// <summary>
        /// Enable/disable SD-WAN service gateway.
        /// </summary>
        public readonly string Gateway;
        /// <summary>
        /// User groups. The structure of `groups` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetVirtualwanlinkServiceGroupResult> Groups;
        /// <summary>
        /// Virtual WAN Link health-check.
        /// </summary>
        public readonly string HealthCheck;
        /// <summary>
        /// Waiting period in seconds when switching from the back-up member to the primary member (0 - 10000000, default = 0).
        /// </summary>
        public readonly int HoldDownTime;
        /// <summary>
        /// SLA ID.
        /// </summary>
        public readonly int Id;
        /// <summary>
        /// Enable/disable negation of input device match.
        /// </summary>
        public readonly string InputDeviceNegate;
        /// <summary>
        /// Source interface name. The structure of `input_device` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetVirtualwanlinkServiceInputDeviceResult> InputDevices;
        /// <summary>
        /// Enable/disable use of Internet service for application-based load balancing.
        /// </summary>
        public readonly string InternetService;
        /// <summary>
        /// Application control based Internet Service group list. The structure of `internet_service_app_ctrl_group` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetVirtualwanlinkServiceInternetServiceAppCtrlGroupResult> InternetServiceAppCtrlGroups;
        /// <summary>
        /// Application control based Internet Service ID list. The structure of `internet_service_app_ctrl` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetVirtualwanlinkServiceInternetServiceAppCtrlResult> InternetServiceAppCtrls;
        /// <summary>
        /// Control-based Internet Service group list. The structure of `internet_service_ctrl_group` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetVirtualwanlinkServiceInternetServiceCtrlGroupResult> InternetServiceCtrlGroups;
        /// <summary>
        /// Control-based Internet Service ID list. The structure of `internet_service_ctrl` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetVirtualwanlinkServiceInternetServiceCtrlResult> InternetServiceCtrls;
        /// <summary>
        /// Custom Internet Service group list. The structure of `internet_service_custom_group` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetVirtualwanlinkServiceInternetServiceCustomGroupResult> InternetServiceCustomGroups;
        /// <summary>
        /// Custom Internet service name list. The structure of `internet_service_custom` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetVirtualwanlinkServiceInternetServiceCustomResult> InternetServiceCustoms;
        /// <summary>
        /// Internet Service group list. The structure of `internet_service_group` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetVirtualwanlinkServiceInternetServiceGroupResult> InternetServiceGroups;
        /// <summary>
        /// Internet service ID list. The structure of `internet_service_id` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetVirtualwanlinkServiceInternetServiceIdResult> InternetServiceIds;
        /// <summary>
        /// Internet service name list. The structure of `internet_service_name` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetVirtualwanlinkServiceInternetServiceNameResult> InternetServiceNames;
        /// <summary>
        /// Coefficient of jitter in the formula of custom-profile-1.
        /// </summary>
        public readonly int JitterWeight;
        /// <summary>
        /// Coefficient of latency in the formula of custom-profile-1.
        /// </summary>
        public readonly int LatencyWeight;
        /// <summary>
        /// Link cost factor.
        /// </summary>
        public readonly string LinkCostFactor;
        /// <summary>
        /// Percentage threshold change of link cost values that will result in policy route regeneration (0 - 10000000, default = 10).
        /// </summary>
        public readonly int LinkCostThreshold;
        /// <summary>
        /// Member sequence number.
        /// </summary>
        public readonly int Member;
        /// <summary>
        /// Control how the priority rule sets the priority of interfaces in the SD-WAN.
        /// </summary>
        public readonly string Mode;
        /// <summary>
        /// Control-based Internet Service group name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Coefficient of packet-loss in the formula of custom-profile-1.
        /// </summary>
        public readonly int PacketLossWeight;
        /// <summary>
        /// Member sequence number list. The structure of `priority_members` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetVirtualwanlinkServicePriorityMemberResult> PriorityMembers;
        /// <summary>
        /// Protocol number.
        /// </summary>
        public readonly int Protocol;
        /// <summary>
        /// Quality grade.
        /// </summary>
        public readonly int QualityLink;
        /// <summary>
        /// Service role to work with neighbor.
        /// </summary>
        public readonly string Role;
        /// <summary>
        /// IPv4 route map route-tag.
        /// </summary>
        public readonly int RouteTag;
        /// <summary>
        /// Method to compare SLA value for sla and load balance mode.
        /// </summary>
        public readonly string SlaCompareMethod;
        /// <summary>
        /// Service level agreement (SLA). The structure of `sla` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetVirtualwanlinkServiceSlaResult> Slas;
        /// <summary>
        /// Source address6 name. The structure of `src6` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetVirtualwanlinkServiceSrc6Result> Src6s;
        /// <summary>
        /// Enable/disable negation of source address match.
        /// </summary>
        public readonly string SrcNegate;
        /// <summary>
        /// Source address name. The structure of `src` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetVirtualwanlinkServiceSrcResult> Srcs;
        /// <summary>
        /// Enable/disable service when selected neighbor role is standalone while service role is not standalone.
        /// </summary>
        public readonly string StandaloneAction;
        /// <summary>
        /// Start destination port number.
        /// </summary>
        public readonly int StartPort;
        /// <summary>
        /// Enable/disable SD-WAN service.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// Type of service bit pattern.
        /// </summary>
        public readonly string Tos;
        /// <summary>
        /// Type of service evaluated bits.
        /// </summary>
        public readonly string TosMask;
        /// <summary>
        /// User name. The structure of `users` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetVirtualwanlinkServiceUserResult> Users;

        [OutputConstructor]
        private GetVirtualwanlinkServiceResult(
            string addrMode,

            int bandwidthWeight,

            string @default,

            string dscpForward,

            string dscpForwardTag,

            string dscpReverse,

            string dscpReverseTag,

            ImmutableArray<Outputs.GetVirtualwanlinkServiceDst6Result> dst6s,

            string dstNegate,

            ImmutableArray<Outputs.GetVirtualwanlinkServiceDstResult> dsts,

            int endPort,

            string gateway,

            ImmutableArray<Outputs.GetVirtualwanlinkServiceGroupResult> groups,

            string healthCheck,

            int holdDownTime,

            int id,

            string inputDeviceNegate,

            ImmutableArray<Outputs.GetVirtualwanlinkServiceInputDeviceResult> inputDevices,

            string internetService,

            ImmutableArray<Outputs.GetVirtualwanlinkServiceInternetServiceAppCtrlGroupResult> internetServiceAppCtrlGroups,

            ImmutableArray<Outputs.GetVirtualwanlinkServiceInternetServiceAppCtrlResult> internetServiceAppCtrls,

            ImmutableArray<Outputs.GetVirtualwanlinkServiceInternetServiceCtrlGroupResult> internetServiceCtrlGroups,

            ImmutableArray<Outputs.GetVirtualwanlinkServiceInternetServiceCtrlResult> internetServiceCtrls,

            ImmutableArray<Outputs.GetVirtualwanlinkServiceInternetServiceCustomGroupResult> internetServiceCustomGroups,

            ImmutableArray<Outputs.GetVirtualwanlinkServiceInternetServiceCustomResult> internetServiceCustoms,

            ImmutableArray<Outputs.GetVirtualwanlinkServiceInternetServiceGroupResult> internetServiceGroups,

            ImmutableArray<Outputs.GetVirtualwanlinkServiceInternetServiceIdResult> internetServiceIds,

            ImmutableArray<Outputs.GetVirtualwanlinkServiceInternetServiceNameResult> internetServiceNames,

            int jitterWeight,

            int latencyWeight,

            string linkCostFactor,

            int linkCostThreshold,

            int member,

            string mode,

            string name,

            int packetLossWeight,

            ImmutableArray<Outputs.GetVirtualwanlinkServicePriorityMemberResult> priorityMembers,

            int protocol,

            int qualityLink,

            string role,

            int routeTag,

            string slaCompareMethod,

            ImmutableArray<Outputs.GetVirtualwanlinkServiceSlaResult> slas,

            ImmutableArray<Outputs.GetVirtualwanlinkServiceSrc6Result> src6s,

            string srcNegate,

            ImmutableArray<Outputs.GetVirtualwanlinkServiceSrcResult> srcs,

            string standaloneAction,

            int startPort,

            string status,

            string tos,

            string tosMask,

            ImmutableArray<Outputs.GetVirtualwanlinkServiceUserResult> users)
        {
            AddrMode = addrMode;
            BandwidthWeight = bandwidthWeight;
            Default = @default;
            DscpForward = dscpForward;
            DscpForwardTag = dscpForwardTag;
            DscpReverse = dscpReverse;
            DscpReverseTag = dscpReverseTag;
            Dst6s = dst6s;
            DstNegate = dstNegate;
            Dsts = dsts;
            EndPort = endPort;
            Gateway = gateway;
            Groups = groups;
            HealthCheck = healthCheck;
            HoldDownTime = holdDownTime;
            Id = id;
            InputDeviceNegate = inputDeviceNegate;
            InputDevices = inputDevices;
            InternetService = internetService;
            InternetServiceAppCtrlGroups = internetServiceAppCtrlGroups;
            InternetServiceAppCtrls = internetServiceAppCtrls;
            InternetServiceCtrlGroups = internetServiceCtrlGroups;
            InternetServiceCtrls = internetServiceCtrls;
            InternetServiceCustomGroups = internetServiceCustomGroups;
            InternetServiceCustoms = internetServiceCustoms;
            InternetServiceGroups = internetServiceGroups;
            InternetServiceIds = internetServiceIds;
            InternetServiceNames = internetServiceNames;
            JitterWeight = jitterWeight;
            LatencyWeight = latencyWeight;
            LinkCostFactor = linkCostFactor;
            LinkCostThreshold = linkCostThreshold;
            Member = member;
            Mode = mode;
            Name = name;
            PacketLossWeight = packetLossWeight;
            PriorityMembers = priorityMembers;
            Protocol = protocol;
            QualityLink = qualityLink;
            Role = role;
            RouteTag = routeTag;
            SlaCompareMethod = slaCompareMethod;
            Slas = slas;
            Src6s = src6s;
            SrcNegate = srcNegate;
            Srcs = srcs;
            StandaloneAction = standaloneAction;
            StartPort = startPort;
            Status = status;
            Tos = tos;
            TosMask = tosMask;
            Users = users;
        }
    }
}