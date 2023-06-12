// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Dlp
{
    /// <summary>
    /// Configure DLP sensors.
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
    ///     var trname = new Fortios.Dlp.Sensor("trname", new()
    ///     {
    ///         DlpLog = "enable",
    ///         ExtendedLog = "disable",
    ///         FlowBased = "enable",
    ///         NacQuarLog = "disable",
    ///         SummaryProto = "smtp pop3",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Dlp Sensor can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:dlp/sensor:Sensor labelname {{name}}
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:dlp/sensor:Sensor labelname {{name}}
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:dlp/sensor:Sensor")]
    public partial class Sensor : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Comment.
        /// </summary>
        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        /// <summary>
        /// Enable/disable DLP logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("dlpLog")]
        public Output<string> DlpLog { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// DLP sensor entries. The structure of `entries` block is documented below.
        /// </summary>
        [Output("entries")]
        public Output<ImmutableArray<Outputs.SensorEntry>> Entries { get; private set; } = null!;

        /// <summary>
        /// Expression to evaluate.
        /// </summary>
        [Output("eval")]
        public Output<string> Eval { get; private set; } = null!;

        /// <summary>
        /// Enable/disable extended logging for data leak prevention. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("extendedLog")]
        public Output<string> ExtendedLog { get; private set; } = null!;

        /// <summary>
        /// Flow/proxy feature set. Valid values: `flow`, `proxy`.
        /// </summary>
        [Output("featureSet")]
        public Output<string> FeatureSet { get; private set; } = null!;

        /// <summary>
        /// Set up DLP filters for this sensor. The structure of `filter` block is documented below.
        /// </summary>
        [Output("filters")]
        public Output<ImmutableArray<Outputs.SensorFilter>> Filters { get; private set; } = null!;

        /// <summary>
        /// Enable/disable flow-based DLP. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("flowBased")]
        public Output<string> FlowBased { get; private set; } = null!;

        /// <summary>
        /// Protocols to always content archive.
        /// </summary>
        [Output("fullArchiveProto")]
        public Output<string> FullArchiveProto { get; private set; } = null!;

        /// <summary>
        /// Logical relation between entries (default = match-any). Valid values: `match-all`, `match-any`, `match-eval`.
        /// </summary>
        [Output("matchType")]
        public Output<string> MatchType { get; private set; } = null!;

        /// <summary>
        /// Enable/disable NAC quarantine logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("nacQuarLog")]
        public Output<string> NacQuarLog { get; private set; } = null!;

        /// <summary>
        /// Name of the DLP sensor.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Configure DLP options.
        /// </summary>
        [Output("options")]
        public Output<string> Options { get; private set; } = null!;

        /// <summary>
        /// Replacement message group used by this DLP sensor.
        /// </summary>
        [Output("replacemsgGroup")]
        public Output<string> ReplacemsgGroup { get; private set; } = null!;

        /// <summary>
        /// Protocols to always log summary.
        /// </summary>
        [Output("summaryProto")]
        public Output<string> SummaryProto { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Sensor resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Sensor(string name, SensorArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:dlp/sensor:Sensor", name, args ?? new SensorArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Sensor(string name, Input<string> id, SensorState? state = null, CustomResourceOptions? options = null)
            : base("fortios:dlp/sensor:Sensor", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Sensor resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Sensor Get(string name, Input<string> id, SensorState? state = null, CustomResourceOptions? options = null)
        {
            return new Sensor(name, id, state, options);
        }
    }

    public sealed class SensorArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Enable/disable DLP logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("dlpLog")]
        public Input<string>? DlpLog { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        [Input("entries")]
        private InputList<Inputs.SensorEntryArgs>? _entries;

        /// <summary>
        /// DLP sensor entries. The structure of `entries` block is documented below.
        /// </summary>
        public InputList<Inputs.SensorEntryArgs> Entries
        {
            get => _entries ?? (_entries = new InputList<Inputs.SensorEntryArgs>());
            set => _entries = value;
        }

        /// <summary>
        /// Expression to evaluate.
        /// </summary>
        [Input("eval")]
        public Input<string>? Eval { get; set; }

        /// <summary>
        /// Enable/disable extended logging for data leak prevention. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("extendedLog")]
        public Input<string>? ExtendedLog { get; set; }

        /// <summary>
        /// Flow/proxy feature set. Valid values: `flow`, `proxy`.
        /// </summary>
        [Input("featureSet")]
        public Input<string>? FeatureSet { get; set; }

        [Input("filters")]
        private InputList<Inputs.SensorFilterArgs>? _filters;

        /// <summary>
        /// Set up DLP filters for this sensor. The structure of `filter` block is documented below.
        /// </summary>
        public InputList<Inputs.SensorFilterArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.SensorFilterArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// Enable/disable flow-based DLP. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("flowBased")]
        public Input<string>? FlowBased { get; set; }

        /// <summary>
        /// Protocols to always content archive.
        /// </summary>
        [Input("fullArchiveProto")]
        public Input<string>? FullArchiveProto { get; set; }

        /// <summary>
        /// Logical relation between entries (default = match-any). Valid values: `match-all`, `match-any`, `match-eval`.
        /// </summary>
        [Input("matchType")]
        public Input<string>? MatchType { get; set; }

        /// <summary>
        /// Enable/disable NAC quarantine logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("nacQuarLog")]
        public Input<string>? NacQuarLog { get; set; }

        /// <summary>
        /// Name of the DLP sensor.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Configure DLP options.
        /// </summary>
        [Input("options")]
        public Input<string>? Options { get; set; }

        /// <summary>
        /// Replacement message group used by this DLP sensor.
        /// </summary>
        [Input("replacemsgGroup")]
        public Input<string>? ReplacemsgGroup { get; set; }

        /// <summary>
        /// Protocols to always log summary.
        /// </summary>
        [Input("summaryProto")]
        public Input<string>? SummaryProto { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SensorArgs()
        {
        }
        public static new SensorArgs Empty => new SensorArgs();
    }

    public sealed class SensorState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Enable/disable DLP logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("dlpLog")]
        public Input<string>? DlpLog { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        [Input("entries")]
        private InputList<Inputs.SensorEntryGetArgs>? _entries;

        /// <summary>
        /// DLP sensor entries. The structure of `entries` block is documented below.
        /// </summary>
        public InputList<Inputs.SensorEntryGetArgs> Entries
        {
            get => _entries ?? (_entries = new InputList<Inputs.SensorEntryGetArgs>());
            set => _entries = value;
        }

        /// <summary>
        /// Expression to evaluate.
        /// </summary>
        [Input("eval")]
        public Input<string>? Eval { get; set; }

        /// <summary>
        /// Enable/disable extended logging for data leak prevention. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("extendedLog")]
        public Input<string>? ExtendedLog { get; set; }

        /// <summary>
        /// Flow/proxy feature set. Valid values: `flow`, `proxy`.
        /// </summary>
        [Input("featureSet")]
        public Input<string>? FeatureSet { get; set; }

        [Input("filters")]
        private InputList<Inputs.SensorFilterGetArgs>? _filters;

        /// <summary>
        /// Set up DLP filters for this sensor. The structure of `filter` block is documented below.
        /// </summary>
        public InputList<Inputs.SensorFilterGetArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.SensorFilterGetArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// Enable/disable flow-based DLP. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("flowBased")]
        public Input<string>? FlowBased { get; set; }

        /// <summary>
        /// Protocols to always content archive.
        /// </summary>
        [Input("fullArchiveProto")]
        public Input<string>? FullArchiveProto { get; set; }

        /// <summary>
        /// Logical relation between entries (default = match-any). Valid values: `match-all`, `match-any`, `match-eval`.
        /// </summary>
        [Input("matchType")]
        public Input<string>? MatchType { get; set; }

        /// <summary>
        /// Enable/disable NAC quarantine logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("nacQuarLog")]
        public Input<string>? NacQuarLog { get; set; }

        /// <summary>
        /// Name of the DLP sensor.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Configure DLP options.
        /// </summary>
        [Input("options")]
        public Input<string>? Options { get; set; }

        /// <summary>
        /// Replacement message group used by this DLP sensor.
        /// </summary>
        [Input("replacemsgGroup")]
        public Input<string>? ReplacemsgGroup { get; set; }

        /// <summary>
        /// Protocols to always log summary.
        /// </summary>
        [Input("summaryProto")]
        public Input<string>? SummaryProto { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SensorState()
        {
        }
        public static new SensorState Empty => new SensorState();
    }
}
