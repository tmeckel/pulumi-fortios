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
    /// <summary>
    /// Configure alarm.
    /// 
    /// ## Import
    /// 
    /// System Alarm can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:sys/alarm:Alarm labelname SystemAlarm
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:sys/alarm:Alarm labelname SystemAlarm
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:sys/alarm:Alarm")]
    public partial class Alarm : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Enable/disable audible alarm. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("audible")]
        public Output<string> Audible { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// Alarm groups. The structure of `groups` block is documented below.
        /// </summary>
        [Output("groups")]
        public Output<ImmutableArray<Outputs.AlarmGroup>> Groups { get; private set; } = null!;

        /// <summary>
        /// Enable/disable alarm. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Alarm resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Alarm(string name, AlarmArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:sys/alarm:Alarm", name, args ?? new AlarmArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Alarm(string name, Input<string> id, AlarmState? state = null, CustomResourceOptions? options = null)
            : base("fortios:sys/alarm:Alarm", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Alarm resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Alarm Get(string name, Input<string> id, AlarmState? state = null, CustomResourceOptions? options = null)
        {
            return new Alarm(name, id, state, options);
        }
    }

    public sealed class AlarmArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable audible alarm. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("audible")]
        public Input<string>? Audible { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        [Input("groups")]
        private InputList<Inputs.AlarmGroupArgs>? _groups;

        /// <summary>
        /// Alarm groups. The structure of `groups` block is documented below.
        /// </summary>
        public InputList<Inputs.AlarmGroupArgs> Groups
        {
            get => _groups ?? (_groups = new InputList<Inputs.AlarmGroupArgs>());
            set => _groups = value;
        }

        /// <summary>
        /// Enable/disable alarm. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public AlarmArgs()
        {
        }
        public static new AlarmArgs Empty => new AlarmArgs();
    }

    public sealed class AlarmState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable audible alarm. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("audible")]
        public Input<string>? Audible { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        [Input("groups")]
        private InputList<Inputs.AlarmGroupGetArgs>? _groups;

        /// <summary>
        /// Alarm groups. The structure of `groups` block is documented below.
        /// </summary>
        public InputList<Inputs.AlarmGroupGetArgs> Groups
        {
            get => _groups ?? (_groups = new InputList<Inputs.AlarmGroupGetArgs>());
            set => _groups = value;
        }

        /// <summary>
        /// Enable/disable alarm. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public AlarmState()
        {
        }
        public static new AlarmState Empty => new AlarmState();
    }
}