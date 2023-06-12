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
    /// <summary>
    /// Configure OSPF.
    /// 
    /// &gt; The provider supports the definition of Ospf-Interface in Router Ospf `fortios.router.Ospf`, and also allows the definition of separate Ospf-Interface resources `fortios.routerospf.Ospfinterface`, but do not use a `fortios.router.Ospf` with in-line Ospf-Interface in conjunction with any `fortios.routerospf.Ospfinterface` resources, otherwise conflicts and overwrite will occur.
    /// 
    /// &gt; The provider supports the definition of Network in Router Ospf `fortios.router.Ospf`, and also allows the definition of separate Network resources `fortios.routerospf.Network`, but do not use a `fortios.router.Ospf` with in-line Network in conjunction with any `fortios.routerospf.Network` resources, otherwise conflicts and overwrite will occur.
    /// 
    /// &gt; The provider supports the definition of Neighbor in Router Ospf `fortios.router.Ospf`, and also allows the definition of separate Neighbor resources `fortios.routerospf.Neighbor`, but do not use a `fortios.router.Ospf` with in-line Neighbor in conjunction with any `fortios.routerospf.Neighbor` resources, otherwise conflicts and overwrite will occur.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Fortios = Pulumiverse.Fortios;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var trname = new Fortios.Router.Ospf("trname", new()
    ///     {
    ///         AbrType = "standard",
    ///         AutoCostRefBandwidth = 1000,
    ///         Bfd = "disable",
    ///         DatabaseOverflow = "disable",
    ///         DatabaseOverflowMaxLsas = 10000,
    ///         DatabaseOverflowTimeToRecover = 300,
    ///         DefaultInformationMetric = 10,
    ///         DefaultInformationMetricType = "2",
    ///         DefaultInformationOriginate = "disable",
    ///         DefaultMetric = 10,
    ///         Distance = 110,
    ///         DistanceExternal = 110,
    ///         DistanceInterArea = 110,
    ///         DistanceIntraArea = 110,
    ///         LogNeighbourChanges = "enable",
    ///         Redistributes = new[]
    ///         {
    ///             new Fortios.Router.Inputs.OspfRedistributeArgs
    ///             {
    ///                 Metric = 0,
    ///                 MetricType = "2",
    ///                 Name = "connected",
    ///                 Status = "disable",
    ///                 Tag = 0,
    ///             },
    ///             new Fortios.Router.Inputs.OspfRedistributeArgs
    ///             {
    ///                 Metric = 0,
    ///                 MetricType = "2",
    ///                 Name = "static",
    ///                 Status = "disable",
    ///                 Tag = 0,
    ///             },
    ///             new Fortios.Router.Inputs.OspfRedistributeArgs
    ///             {
    ///                 Metric = 0,
    ///                 MetricType = "2",
    ///                 Name = "rip",
    ///                 Status = "disable",
    ///                 Tag = 0,
    ///             },
    ///             new Fortios.Router.Inputs.OspfRedistributeArgs
    ///             {
    ///                 Metric = 0,
    ///                 MetricType = "2",
    ///                 Name = "bgp",
    ///                 Status = "disable",
    ///                 Tag = 0,
    ///             },
    ///             new Fortios.Router.Inputs.OspfRedistributeArgs
    ///             {
    ///                 Metric = 0,
    ///                 MetricType = "2",
    ///                 Name = "isis",
    ///                 Status = "disable",
    ///                 Tag = 0,
    ///             },
    ///         },
    ///         RestartMode = "none",
    ///         RestartPeriod = 120,
    ///         Rfc1583Compatible = "disable",
    ///         RouterId = "0.0.0.0",
    ///         SpfTimers = "5 10",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Router Ospf can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:router/ospf:Ospf labelname RouterOspf
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:router/ospf:Ospf labelname RouterOspf
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:router/ospf:Ospf")]
    public partial class Ospf : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Area border router type. Valid values: `cisco`, `ibm`, `shortcut`, `standard`.
        /// </summary>
        [Output("abrType")]
        public Output<string> AbrType { get; private set; } = null!;

        /// <summary>
        /// OSPF area configuration. The structure of `area` block is documented below.
        /// </summary>
        [Output("areas")]
        public Output<ImmutableArray<Outputs.OspfArea>> Areas { get; private set; } = null!;

        /// <summary>
        /// Reference bandwidth in terms of megabits per second.
        /// </summary>
        [Output("autoCostRefBandwidth")]
        public Output<int> AutoCostRefBandwidth { get; private set; } = null!;

        /// <summary>
        /// Bidirectional Forwarding Detection (BFD). Valid values: `enable`, `disable`.
        /// </summary>
        [Output("bfd")]
        public Output<string> Bfd { get; private set; } = null!;

        /// <summary>
        /// Enable/disable database overflow. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("databaseOverflow")]
        public Output<string> DatabaseOverflow { get; private set; } = null!;

        /// <summary>
        /// Database overflow maximum LSAs.
        /// </summary>
        [Output("databaseOverflowMaxLsas")]
        public Output<int> DatabaseOverflowMaxLsas { get; private set; } = null!;

        /// <summary>
        /// Database overflow time to recover (sec).
        /// </summary>
        [Output("databaseOverflowTimeToRecover")]
        public Output<int> DatabaseOverflowTimeToRecover { get; private set; } = null!;

        /// <summary>
        /// Default information metric.
        /// </summary>
        [Output("defaultInformationMetric")]
        public Output<int> DefaultInformationMetric { get; private set; } = null!;

        /// <summary>
        /// Default information metric type. Valid values: `1`, `2`.
        /// </summary>
        [Output("defaultInformationMetricType")]
        public Output<string> DefaultInformationMetricType { get; private set; } = null!;

        /// <summary>
        /// Enable/disable generation of default route. Valid values: `enable`, `always`, `disable`.
        /// </summary>
        [Output("defaultInformationOriginate")]
        public Output<string> DefaultInformationOriginate { get; private set; } = null!;

        /// <summary>
        /// Default information route map.
        /// </summary>
        [Output("defaultInformationRouteMap")]
        public Output<string> DefaultInformationRouteMap { get; private set; } = null!;

        /// <summary>
        /// Default metric of redistribute routes.
        /// </summary>
        [Output("defaultMetric")]
        public Output<int> DefaultMetric { get; private set; } = null!;

        /// <summary>
        /// Distance of the route.
        /// </summary>
        [Output("distance")]
        public Output<int> Distance { get; private set; } = null!;

        /// <summary>
        /// Administrative external distance.
        /// </summary>
        [Output("distanceExternal")]
        public Output<int> DistanceExternal { get; private set; } = null!;

        /// <summary>
        /// Administrative inter-area distance.
        /// </summary>
        [Output("distanceInterArea")]
        public Output<int> DistanceInterArea { get; private set; } = null!;

        /// <summary>
        /// Administrative intra-area distance.
        /// </summary>
        [Output("distanceIntraArea")]
        public Output<int> DistanceIntraArea { get; private set; } = null!;

        /// <summary>
        /// Filter incoming routes.
        /// </summary>
        [Output("distributeListIn")]
        public Output<string> DistributeListIn { get; private set; } = null!;

        /// <summary>
        /// Distribute list configuration. The structure of `distribute_list` block is documented below.
        /// </summary>
        [Output("distributeLists")]
        public Output<ImmutableArray<Outputs.OspfDistributeList>> DistributeLists { get; private set; } = null!;

        /// <summary>
        /// Filter incoming external routes by route-map.
        /// </summary>
        [Output("distributeRouteMapIn")]
        public Output<string> DistributeRouteMapIn { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// Enable logging of OSPF neighbour's changes Valid values: `enable`, `disable`.
        /// </summary>
        [Output("logNeighbourChanges")]
        public Output<string> LogNeighbourChanges { get; private set; } = null!;

        /// <summary>
        /// OSPF neighbor configuration are used when OSPF runs on non-broadcast media The structure of `neighbor` block is documented below.
        /// </summary>
        [Output("neighbors")]
        public Output<ImmutableArray<Outputs.OspfNeighbor>> Neighbors { get; private set; } = null!;

        /// <summary>
        /// OSPF network configuration. The structure of `network` block is documented below.
        /// </summary>
        [Output("networks")]
        public Output<ImmutableArray<Outputs.OspfNetwork>> Networks { get; private set; } = null!;

        /// <summary>
        /// OSPF interface configuration. The structure of `ospf_interface` block is documented below.
        /// </summary>
        [Output("ospfInterfaces")]
        public Output<ImmutableArray<Outputs.OspfOspfInterface>> OspfInterfaces { get; private set; } = null!;

        /// <summary>
        /// Passive interface configuration. The structure of `passive_interface` block is documented below.
        /// </summary>
        [Output("passiveInterfaces")]
        public Output<ImmutableArray<Outputs.OspfPassiveInterface>> PassiveInterfaces { get; private set; } = null!;

        /// <summary>
        /// Redistribute configuration. The structure of `redistribute` block is documented below.
        /// </summary>
        [Output("redistributes")]
        public Output<ImmutableArray<Outputs.OspfRedistribute>> Redistributes { get; private set; } = null!;

        /// <summary>
        /// OSPF restart mode (graceful or LLS). Valid values: `none`, `lls`, `graceful-restart`.
        /// </summary>
        [Output("restartMode")]
        public Output<string> RestartMode { get; private set; } = null!;

        /// <summary>
        /// Enable/disable continuing graceful restart upon topology change. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("restartOnTopologyChange")]
        public Output<string> RestartOnTopologyChange { get; private set; } = null!;

        /// <summary>
        /// Graceful restart period.
        /// </summary>
        [Output("restartPeriod")]
        public Output<int> RestartPeriod { get; private set; } = null!;

        /// <summary>
        /// Enable/disable RFC1583 compatibility. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("rfc1583Compatible")]
        public Output<string> Rfc1583Compatible { get; private set; } = null!;

        /// <summary>
        /// Router ID.
        /// </summary>
        [Output("routerId")]
        public Output<string> RouterId { get; private set; } = null!;

        /// <summary>
        /// SPF calculation frequency.
        /// </summary>
        [Output("spfTimers")]
        public Output<string> SpfTimers { get; private set; } = null!;

        /// <summary>
        /// IP address summary configuration. The structure of `summary_address` block is documented below.
        /// </summary>
        [Output("summaryAddresses")]
        public Output<ImmutableArray<Outputs.OspfSummaryAddress>> SummaryAddresses { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Ospf resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Ospf(string name, OspfArgs args, CustomResourceOptions? options = null)
            : base("fortios:router/ospf:Ospf", name, args ?? new OspfArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Ospf(string name, Input<string> id, OspfState? state = null, CustomResourceOptions? options = null)
            : base("fortios:router/ospf:Ospf", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse/pulumi-fortios",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Ospf resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Ospf Get(string name, Input<string> id, OspfState? state = null, CustomResourceOptions? options = null)
        {
            return new Ospf(name, id, state, options);
        }
    }

    public sealed class OspfArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Area border router type. Valid values: `cisco`, `ibm`, `shortcut`, `standard`.
        /// </summary>
        [Input("abrType")]
        public Input<string>? AbrType { get; set; }

        [Input("areas")]
        private InputList<Inputs.OspfAreaArgs>? _areas;

        /// <summary>
        /// OSPF area configuration. The structure of `area` block is documented below.
        /// </summary>
        public InputList<Inputs.OspfAreaArgs> Areas
        {
            get => _areas ?? (_areas = new InputList<Inputs.OspfAreaArgs>());
            set => _areas = value;
        }

        /// <summary>
        /// Reference bandwidth in terms of megabits per second.
        /// </summary>
        [Input("autoCostRefBandwidth")]
        public Input<int>? AutoCostRefBandwidth { get; set; }

        /// <summary>
        /// Bidirectional Forwarding Detection (BFD). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("bfd")]
        public Input<string>? Bfd { get; set; }

        /// <summary>
        /// Enable/disable database overflow. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("databaseOverflow")]
        public Input<string>? DatabaseOverflow { get; set; }

        /// <summary>
        /// Database overflow maximum LSAs.
        /// </summary>
        [Input("databaseOverflowMaxLsas")]
        public Input<int>? DatabaseOverflowMaxLsas { get; set; }

        /// <summary>
        /// Database overflow time to recover (sec).
        /// </summary>
        [Input("databaseOverflowTimeToRecover")]
        public Input<int>? DatabaseOverflowTimeToRecover { get; set; }

        /// <summary>
        /// Default information metric.
        /// </summary>
        [Input("defaultInformationMetric")]
        public Input<int>? DefaultInformationMetric { get; set; }

        /// <summary>
        /// Default information metric type. Valid values: `1`, `2`.
        /// </summary>
        [Input("defaultInformationMetricType")]
        public Input<string>? DefaultInformationMetricType { get; set; }

        /// <summary>
        /// Enable/disable generation of default route. Valid values: `enable`, `always`, `disable`.
        /// </summary>
        [Input("defaultInformationOriginate")]
        public Input<string>? DefaultInformationOriginate { get; set; }

        /// <summary>
        /// Default information route map.
        /// </summary>
        [Input("defaultInformationRouteMap")]
        public Input<string>? DefaultInformationRouteMap { get; set; }

        /// <summary>
        /// Default metric of redistribute routes.
        /// </summary>
        [Input("defaultMetric")]
        public Input<int>? DefaultMetric { get; set; }

        /// <summary>
        /// Distance of the route.
        /// </summary>
        [Input("distance")]
        public Input<int>? Distance { get; set; }

        /// <summary>
        /// Administrative external distance.
        /// </summary>
        [Input("distanceExternal")]
        public Input<int>? DistanceExternal { get; set; }

        /// <summary>
        /// Administrative inter-area distance.
        /// </summary>
        [Input("distanceInterArea")]
        public Input<int>? DistanceInterArea { get; set; }

        /// <summary>
        /// Administrative intra-area distance.
        /// </summary>
        [Input("distanceIntraArea")]
        public Input<int>? DistanceIntraArea { get; set; }

        /// <summary>
        /// Filter incoming routes.
        /// </summary>
        [Input("distributeListIn")]
        public Input<string>? DistributeListIn { get; set; }

        [Input("distributeLists")]
        private InputList<Inputs.OspfDistributeListArgs>? _distributeLists;

        /// <summary>
        /// Distribute list configuration. The structure of `distribute_list` block is documented below.
        /// </summary>
        public InputList<Inputs.OspfDistributeListArgs> DistributeLists
        {
            get => _distributeLists ?? (_distributeLists = new InputList<Inputs.OspfDistributeListArgs>());
            set => _distributeLists = value;
        }

        /// <summary>
        /// Filter incoming external routes by route-map.
        /// </summary>
        [Input("distributeRouteMapIn")]
        public Input<string>? DistributeRouteMapIn { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Enable logging of OSPF neighbour's changes Valid values: `enable`, `disable`.
        /// </summary>
        [Input("logNeighbourChanges")]
        public Input<string>? LogNeighbourChanges { get; set; }

        [Input("neighbors")]
        private InputList<Inputs.OspfNeighborArgs>? _neighbors;

        /// <summary>
        /// OSPF neighbor configuration are used when OSPF runs on non-broadcast media The structure of `neighbor` block is documented below.
        /// </summary>
        public InputList<Inputs.OspfNeighborArgs> Neighbors
        {
            get => _neighbors ?? (_neighbors = new InputList<Inputs.OspfNeighborArgs>());
            set => _neighbors = value;
        }

        [Input("networks")]
        private InputList<Inputs.OspfNetworkArgs>? _networks;

        /// <summary>
        /// OSPF network configuration. The structure of `network` block is documented below.
        /// </summary>
        public InputList<Inputs.OspfNetworkArgs> Networks
        {
            get => _networks ?? (_networks = new InputList<Inputs.OspfNetworkArgs>());
            set => _networks = value;
        }

        [Input("ospfInterfaces")]
        private InputList<Inputs.OspfOspfInterfaceArgs>? _ospfInterfaces;

        /// <summary>
        /// OSPF interface configuration. The structure of `ospf_interface` block is documented below.
        /// </summary>
        public InputList<Inputs.OspfOspfInterfaceArgs> OspfInterfaces
        {
            get => _ospfInterfaces ?? (_ospfInterfaces = new InputList<Inputs.OspfOspfInterfaceArgs>());
            set => _ospfInterfaces = value;
        }

        [Input("passiveInterfaces")]
        private InputList<Inputs.OspfPassiveInterfaceArgs>? _passiveInterfaces;

        /// <summary>
        /// Passive interface configuration. The structure of `passive_interface` block is documented below.
        /// </summary>
        public InputList<Inputs.OspfPassiveInterfaceArgs> PassiveInterfaces
        {
            get => _passiveInterfaces ?? (_passiveInterfaces = new InputList<Inputs.OspfPassiveInterfaceArgs>());
            set => _passiveInterfaces = value;
        }

        [Input("redistributes")]
        private InputList<Inputs.OspfRedistributeArgs>? _redistributes;

        /// <summary>
        /// Redistribute configuration. The structure of `redistribute` block is documented below.
        /// </summary>
        public InputList<Inputs.OspfRedistributeArgs> Redistributes
        {
            get => _redistributes ?? (_redistributes = new InputList<Inputs.OspfRedistributeArgs>());
            set => _redistributes = value;
        }

        /// <summary>
        /// OSPF restart mode (graceful or LLS). Valid values: `none`, `lls`, `graceful-restart`.
        /// </summary>
        [Input("restartMode")]
        public Input<string>? RestartMode { get; set; }

        /// <summary>
        /// Enable/disable continuing graceful restart upon topology change. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("restartOnTopologyChange")]
        public Input<string>? RestartOnTopologyChange { get; set; }

        /// <summary>
        /// Graceful restart period.
        /// </summary>
        [Input("restartPeriod")]
        public Input<int>? RestartPeriod { get; set; }

        /// <summary>
        /// Enable/disable RFC1583 compatibility. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("rfc1583Compatible")]
        public Input<string>? Rfc1583Compatible { get; set; }

        /// <summary>
        /// Router ID.
        /// </summary>
        [Input("routerId", required: true)]
        public Input<string> RouterId { get; set; } = null!;

        /// <summary>
        /// SPF calculation frequency.
        /// </summary>
        [Input("spfTimers")]
        public Input<string>? SpfTimers { get; set; }

        [Input("summaryAddresses")]
        private InputList<Inputs.OspfSummaryAddressArgs>? _summaryAddresses;

        /// <summary>
        /// IP address summary configuration. The structure of `summary_address` block is documented below.
        /// </summary>
        public InputList<Inputs.OspfSummaryAddressArgs> SummaryAddresses
        {
            get => _summaryAddresses ?? (_summaryAddresses = new InputList<Inputs.OspfSummaryAddressArgs>());
            set => _summaryAddresses = value;
        }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public OspfArgs()
        {
        }
        public static new OspfArgs Empty => new OspfArgs();
    }

    public sealed class OspfState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Area border router type. Valid values: `cisco`, `ibm`, `shortcut`, `standard`.
        /// </summary>
        [Input("abrType")]
        public Input<string>? AbrType { get; set; }

        [Input("areas")]
        private InputList<Inputs.OspfAreaGetArgs>? _areas;

        /// <summary>
        /// OSPF area configuration. The structure of `area` block is documented below.
        /// </summary>
        public InputList<Inputs.OspfAreaGetArgs> Areas
        {
            get => _areas ?? (_areas = new InputList<Inputs.OspfAreaGetArgs>());
            set => _areas = value;
        }

        /// <summary>
        /// Reference bandwidth in terms of megabits per second.
        /// </summary>
        [Input("autoCostRefBandwidth")]
        public Input<int>? AutoCostRefBandwidth { get; set; }

        /// <summary>
        /// Bidirectional Forwarding Detection (BFD). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("bfd")]
        public Input<string>? Bfd { get; set; }

        /// <summary>
        /// Enable/disable database overflow. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("databaseOverflow")]
        public Input<string>? DatabaseOverflow { get; set; }

        /// <summary>
        /// Database overflow maximum LSAs.
        /// </summary>
        [Input("databaseOverflowMaxLsas")]
        public Input<int>? DatabaseOverflowMaxLsas { get; set; }

        /// <summary>
        /// Database overflow time to recover (sec).
        /// </summary>
        [Input("databaseOverflowTimeToRecover")]
        public Input<int>? DatabaseOverflowTimeToRecover { get; set; }

        /// <summary>
        /// Default information metric.
        /// </summary>
        [Input("defaultInformationMetric")]
        public Input<int>? DefaultInformationMetric { get; set; }

        /// <summary>
        /// Default information metric type. Valid values: `1`, `2`.
        /// </summary>
        [Input("defaultInformationMetricType")]
        public Input<string>? DefaultInformationMetricType { get; set; }

        /// <summary>
        /// Enable/disable generation of default route. Valid values: `enable`, `always`, `disable`.
        /// </summary>
        [Input("defaultInformationOriginate")]
        public Input<string>? DefaultInformationOriginate { get; set; }

        /// <summary>
        /// Default information route map.
        /// </summary>
        [Input("defaultInformationRouteMap")]
        public Input<string>? DefaultInformationRouteMap { get; set; }

        /// <summary>
        /// Default metric of redistribute routes.
        /// </summary>
        [Input("defaultMetric")]
        public Input<int>? DefaultMetric { get; set; }

        /// <summary>
        /// Distance of the route.
        /// </summary>
        [Input("distance")]
        public Input<int>? Distance { get; set; }

        /// <summary>
        /// Administrative external distance.
        /// </summary>
        [Input("distanceExternal")]
        public Input<int>? DistanceExternal { get; set; }

        /// <summary>
        /// Administrative inter-area distance.
        /// </summary>
        [Input("distanceInterArea")]
        public Input<int>? DistanceInterArea { get; set; }

        /// <summary>
        /// Administrative intra-area distance.
        /// </summary>
        [Input("distanceIntraArea")]
        public Input<int>? DistanceIntraArea { get; set; }

        /// <summary>
        /// Filter incoming routes.
        /// </summary>
        [Input("distributeListIn")]
        public Input<string>? DistributeListIn { get; set; }

        [Input("distributeLists")]
        private InputList<Inputs.OspfDistributeListGetArgs>? _distributeLists;

        /// <summary>
        /// Distribute list configuration. The structure of `distribute_list` block is documented below.
        /// </summary>
        public InputList<Inputs.OspfDistributeListGetArgs> DistributeLists
        {
            get => _distributeLists ?? (_distributeLists = new InputList<Inputs.OspfDistributeListGetArgs>());
            set => _distributeLists = value;
        }

        /// <summary>
        /// Filter incoming external routes by route-map.
        /// </summary>
        [Input("distributeRouteMapIn")]
        public Input<string>? DistributeRouteMapIn { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Enable logging of OSPF neighbour's changes Valid values: `enable`, `disable`.
        /// </summary>
        [Input("logNeighbourChanges")]
        public Input<string>? LogNeighbourChanges { get; set; }

        [Input("neighbors")]
        private InputList<Inputs.OspfNeighborGetArgs>? _neighbors;

        /// <summary>
        /// OSPF neighbor configuration are used when OSPF runs on non-broadcast media The structure of `neighbor` block is documented below.
        /// </summary>
        public InputList<Inputs.OspfNeighborGetArgs> Neighbors
        {
            get => _neighbors ?? (_neighbors = new InputList<Inputs.OspfNeighborGetArgs>());
            set => _neighbors = value;
        }

        [Input("networks")]
        private InputList<Inputs.OspfNetworkGetArgs>? _networks;

        /// <summary>
        /// OSPF network configuration. The structure of `network` block is documented below.
        /// </summary>
        public InputList<Inputs.OspfNetworkGetArgs> Networks
        {
            get => _networks ?? (_networks = new InputList<Inputs.OspfNetworkGetArgs>());
            set => _networks = value;
        }

        [Input("ospfInterfaces")]
        private InputList<Inputs.OspfOspfInterfaceGetArgs>? _ospfInterfaces;

        /// <summary>
        /// OSPF interface configuration. The structure of `ospf_interface` block is documented below.
        /// </summary>
        public InputList<Inputs.OspfOspfInterfaceGetArgs> OspfInterfaces
        {
            get => _ospfInterfaces ?? (_ospfInterfaces = new InputList<Inputs.OspfOspfInterfaceGetArgs>());
            set => _ospfInterfaces = value;
        }

        [Input("passiveInterfaces")]
        private InputList<Inputs.OspfPassiveInterfaceGetArgs>? _passiveInterfaces;

        /// <summary>
        /// Passive interface configuration. The structure of `passive_interface` block is documented below.
        /// </summary>
        public InputList<Inputs.OspfPassiveInterfaceGetArgs> PassiveInterfaces
        {
            get => _passiveInterfaces ?? (_passiveInterfaces = new InputList<Inputs.OspfPassiveInterfaceGetArgs>());
            set => _passiveInterfaces = value;
        }

        [Input("redistributes")]
        private InputList<Inputs.OspfRedistributeGetArgs>? _redistributes;

        /// <summary>
        /// Redistribute configuration. The structure of `redistribute` block is documented below.
        /// </summary>
        public InputList<Inputs.OspfRedistributeGetArgs> Redistributes
        {
            get => _redistributes ?? (_redistributes = new InputList<Inputs.OspfRedistributeGetArgs>());
            set => _redistributes = value;
        }

        /// <summary>
        /// OSPF restart mode (graceful or LLS). Valid values: `none`, `lls`, `graceful-restart`.
        /// </summary>
        [Input("restartMode")]
        public Input<string>? RestartMode { get; set; }

        /// <summary>
        /// Enable/disable continuing graceful restart upon topology change. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("restartOnTopologyChange")]
        public Input<string>? RestartOnTopologyChange { get; set; }

        /// <summary>
        /// Graceful restart period.
        /// </summary>
        [Input("restartPeriod")]
        public Input<int>? RestartPeriod { get; set; }

        /// <summary>
        /// Enable/disable RFC1583 compatibility. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("rfc1583Compatible")]
        public Input<string>? Rfc1583Compatible { get; set; }

        /// <summary>
        /// Router ID.
        /// </summary>
        [Input("routerId")]
        public Input<string>? RouterId { get; set; }

        /// <summary>
        /// SPF calculation frequency.
        /// </summary>
        [Input("spfTimers")]
        public Input<string>? SpfTimers { get; set; }

        [Input("summaryAddresses")]
        private InputList<Inputs.OspfSummaryAddressGetArgs>? _summaryAddresses;

        /// <summary>
        /// IP address summary configuration. The structure of `summary_address` block is documented below.
        /// </summary>
        public InputList<Inputs.OspfSummaryAddressGetArgs> SummaryAddresses
        {
            get => _summaryAddresses ?? (_summaryAddresses = new InputList<Inputs.OspfSummaryAddressGetArgs>());
            set => _summaryAddresses = value;
        }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public OspfState()
        {
        }
        public static new OspfState Empty => new OspfState();
    }
}
