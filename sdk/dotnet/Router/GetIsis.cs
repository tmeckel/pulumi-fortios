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
    public static class GetIsis
    {
        /// <summary>
        /// Use this data source to get information on fortios router isis
        /// </summary>
        public static Task<GetIsisResult> InvokeAsync(GetIsisArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIsisResult>("fortios:router/getIsis:getIsis", args ?? new GetIsisArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on fortios router isis
        /// </summary>
        public static Output<GetIsisResult> Invoke(GetIsisInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetIsisResult>("fortios:router/getIsis:getIsis", args ?? new GetIsisInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIsisArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetIsisArgs()
        {
        }
        public static new GetIsisArgs Empty => new GetIsisArgs();
    }

    public sealed class GetIsisInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetIsisInvokeArgs()
        {
        }
        public static new GetIsisInvokeArgs Empty => new GetIsisInvokeArgs();
    }


    [OutputType]
    public sealed class GetIsisResult
    {
        /// <summary>
        /// Enable/disable adjacency check.
        /// </summary>
        public readonly string AdjacencyCheck;
        /// <summary>
        /// Enable/disable IPv6 adjacency check.
        /// </summary>
        public readonly string AdjacencyCheck6;
        /// <summary>
        /// Enable/disable IS-IS advertisement of passive interfaces only.
        /// </summary>
        public readonly string AdvPassiveOnly;
        /// <summary>
        /// Enable/disable IPv6 IS-IS advertisement of passive interfaces only.
        /// </summary>
        public readonly string AdvPassiveOnly6;
        /// <summary>
        /// Authentication key-chain for level 1 PDUs.
        /// </summary>
        public readonly string AuthKeychainL1;
        /// <summary>
        /// Authentication key-chain for level 2 PDUs.
        /// </summary>
        public readonly string AuthKeychainL2;
        /// <summary>
        /// Level 1 authentication mode.
        /// </summary>
        public readonly string AuthModeL1;
        /// <summary>
        /// Level 2 authentication mode.
        /// </summary>
        public readonly string AuthModeL2;
        /// <summary>
        /// Authentication password for level 1 PDUs.
        /// </summary>
        public readonly string AuthPasswordL1;
        /// <summary>
        /// Authentication password for level 2 PDUs.
        /// </summary>
        public readonly string AuthPasswordL2;
        /// <summary>
        /// Enable/disable level 1 authentication send-only.
        /// </summary>
        public readonly string AuthSendonlyL1;
        /// <summary>
        /// Enable/disable level 2 authentication send-only.
        /// </summary>
        public readonly string AuthSendonlyL2;
        /// <summary>
        /// Enable/disable distribution of default route information.
        /// </summary>
        public readonly string DefaultOriginate;
        /// <summary>
        /// Enable/disable distribution of default IPv6 route information.
        /// </summary>
        public readonly string DefaultOriginate6;
        /// <summary>
        /// Enable/disable dynamic hostname.
        /// </summary>
        public readonly string DynamicHostname;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Enable/disable ignoring of LSP errors with bad checksums.
        /// </summary>
        public readonly string IgnoreLspErrors;
        /// <summary>
        /// IS type.
        /// </summary>
        public readonly string IsType;
        /// <summary>
        /// IS-IS interface configuration. The structure of `isis_interface` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetIsisIsisInterfaceResult> IsisInterfaces;
        /// <summary>
        /// IS-IS net configuration. The structure of `isis_net` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetIsisIsisNetResult> IsisNets;
        /// <summary>
        /// Minimum interval for level 1 LSP regenerating.
        /// </summary>
        public readonly int LspGenIntervalL1;
        /// <summary>
        /// Minimum interval for level 2 LSP regenerating.
        /// </summary>
        public readonly int LspGenIntervalL2;
        /// <summary>
        /// LSP refresh time in seconds.
        /// </summary>
        public readonly int LspRefreshInterval;
        /// <summary>
        /// Maximum LSP lifetime in seconds.
        /// </summary>
        public readonly int MaxLspLifetime;
        /// <summary>
        /// Use old-style (ISO 10589) or new-style packet formats
        /// </summary>
        public readonly string MetricStyle;
        /// <summary>
        /// Enable/disable signal other routers not to use us in SPF.
        /// </summary>
        public readonly string OverloadBit;
        /// <summary>
        /// Overload-bit only temporarily after reboot.
        /// </summary>
        public readonly int OverloadBitOnStartup;
        /// <summary>
        /// Suppress overload-bit for the specific prefixes.
        /// </summary>
        public readonly string OverloadBitSuppress;
        /// <summary>
        /// Enable/disable redistribution of level 1 IPv6 routes into level 2.
        /// </summary>
        public readonly string Redistribute6L1;
        /// <summary>
        /// Access-list for IPv6 route redistribution from l1 to l2.
        /// </summary>
        public readonly string Redistribute6L1List;
        /// <summary>
        /// Enable/disable redistribution of level 2 IPv6 routes into level 1.
        /// </summary>
        public readonly string Redistribute6L2;
        /// <summary>
        /// Access-list for IPv6 route redistribution from l2 to l1.
        /// </summary>
        public readonly string Redistribute6L2List;
        /// <summary>
        /// IS-IS IPv6 redistribution for routing protocols. The structure of `redistribute6` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetIsisRedistribute6Result> Redistribute6s;
        /// <summary>
        /// Enable/disable redistribution of level 1 routes into level 2.
        /// </summary>
        public readonly string RedistributeL1;
        /// <summary>
        /// Access-list for route redistribution from l1 to l2.
        /// </summary>
        public readonly string RedistributeL1List;
        /// <summary>
        /// Enable/disable redistribution of level 2 routes into level 1.
        /// </summary>
        public readonly string RedistributeL2;
        /// <summary>
        /// Access-list for route redistribution from l2 to l1.
        /// </summary>
        public readonly string RedistributeL2List;
        /// <summary>
        /// IS-IS redistribute protocols. The structure of `redistribute` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetIsisRedistributeResult> Redistributes;
        /// <summary>
        /// Level 1 SPF calculation delay.
        /// </summary>
        public readonly string SpfIntervalExpL1;
        /// <summary>
        /// Level 2 SPF calculation delay.
        /// </summary>
        public readonly string SpfIntervalExpL2;
        /// <summary>
        /// IS-IS IPv6 summary address. The structure of `summary_address6` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetIsisSummaryAddress6Result> SummaryAddress6s;
        /// <summary>
        /// IS-IS summary addresses. The structure of `summary_address` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetIsisSummaryAddressResult> SummaryAddresses;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetIsisResult(
            string adjacencyCheck,

            string adjacencyCheck6,

            string advPassiveOnly,

            string advPassiveOnly6,

            string authKeychainL1,

            string authKeychainL2,

            string authModeL1,

            string authModeL2,

            string authPasswordL1,

            string authPasswordL2,

            string authSendonlyL1,

            string authSendonlyL2,

            string defaultOriginate,

            string defaultOriginate6,

            string dynamicHostname,

            string id,

            string ignoreLspErrors,

            string isType,

            ImmutableArray<Outputs.GetIsisIsisInterfaceResult> isisInterfaces,

            ImmutableArray<Outputs.GetIsisIsisNetResult> isisNets,

            int lspGenIntervalL1,

            int lspGenIntervalL2,

            int lspRefreshInterval,

            int maxLspLifetime,

            string metricStyle,

            string overloadBit,

            int overloadBitOnStartup,

            string overloadBitSuppress,

            string redistribute6L1,

            string redistribute6L1List,

            string redistribute6L2,

            string redistribute6L2List,

            ImmutableArray<Outputs.GetIsisRedistribute6Result> redistribute6s,

            string redistributeL1,

            string redistributeL1List,

            string redistributeL2,

            string redistributeL2List,

            ImmutableArray<Outputs.GetIsisRedistributeResult> redistributes,

            string spfIntervalExpL1,

            string spfIntervalExpL2,

            ImmutableArray<Outputs.GetIsisSummaryAddress6Result> summaryAddress6s,

            ImmutableArray<Outputs.GetIsisSummaryAddressResult> summaryAddresses,

            string? vdomparam)
        {
            AdjacencyCheck = adjacencyCheck;
            AdjacencyCheck6 = adjacencyCheck6;
            AdvPassiveOnly = advPassiveOnly;
            AdvPassiveOnly6 = advPassiveOnly6;
            AuthKeychainL1 = authKeychainL1;
            AuthKeychainL2 = authKeychainL2;
            AuthModeL1 = authModeL1;
            AuthModeL2 = authModeL2;
            AuthPasswordL1 = authPasswordL1;
            AuthPasswordL2 = authPasswordL2;
            AuthSendonlyL1 = authSendonlyL1;
            AuthSendonlyL2 = authSendonlyL2;
            DefaultOriginate = defaultOriginate;
            DefaultOriginate6 = defaultOriginate6;
            DynamicHostname = dynamicHostname;
            Id = id;
            IgnoreLspErrors = ignoreLspErrors;
            IsType = isType;
            IsisInterfaces = isisInterfaces;
            IsisNets = isisNets;
            LspGenIntervalL1 = lspGenIntervalL1;
            LspGenIntervalL2 = lspGenIntervalL2;
            LspRefreshInterval = lspRefreshInterval;
            MaxLspLifetime = maxLspLifetime;
            MetricStyle = metricStyle;
            OverloadBit = overloadBit;
            OverloadBitOnStartup = overloadBitOnStartup;
            OverloadBitSuppress = overloadBitSuppress;
            Redistribute6L1 = redistribute6L1;
            Redistribute6L1List = redistribute6L1List;
            Redistribute6L2 = redistribute6L2;
            Redistribute6L2List = redistribute6L2List;
            Redistribute6s = redistribute6s;
            RedistributeL1 = redistributeL1;
            RedistributeL1List = redistributeL1List;
            RedistributeL2 = redistributeL2;
            RedistributeL2List = redistributeL2List;
            Redistributes = redistributes;
            SpfIntervalExpL1 = spfIntervalExpL1;
            SpfIntervalExpL2 = spfIntervalExpL2;
            SummaryAddress6s = summaryAddress6s;
            SummaryAddresses = summaryAddresses;
            Vdomparam = vdomparam;
        }
    }
}
