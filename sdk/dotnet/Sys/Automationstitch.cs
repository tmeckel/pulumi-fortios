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
    /// Automation stitches.
    /// 
    /// ## Import
    /// 
    /// System AutomationStitch can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:sys/automationstitch:Automationstitch labelname {{name}}
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:sys/automationstitch:Automationstitch labelname {{name}}
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:sys/automationstitch:Automationstitch")]
    public partial class Automationstitch : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Action names. The structure of `action` block is documented below.
        /// </summary>
        [Output("action")]
        public Output<ImmutableArray<Outputs.AutomationstitchAction>> Action { get; private set; } = null!;

        /// <summary>
        /// Configure stitch actions. The structure of `actions` block is documented below.
        /// </summary>
        [Output("actions")]
        public Output<ImmutableArray<Outputs.AutomationstitchAction>> Actions { get; private set; } = null!;

        /// <summary>
        /// Description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Serial number/HA group-name of destination devices. The structure of `destination` block is documented below.
        /// </summary>
        [Output("destinations")]
        public Output<ImmutableArray<Outputs.AutomationstitchDestination>> Destinations { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// Name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Enable/disable this stitch. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Trigger name.
        /// </summary>
        [Output("trigger")]
        public Output<string> Trigger { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Automationstitch resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Automationstitch(string name, AutomationstitchArgs args, CustomResourceOptions? options = null)
            : base("fortios:sys/automationstitch:Automationstitch", name, args ?? new AutomationstitchArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Automationstitch(string name, Input<string> id, AutomationstitchState? state = null, CustomResourceOptions? options = null)
            : base("fortios:sys/automationstitch:Automationstitch", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Automationstitch resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Automationstitch Get(string name, Input<string> id, AutomationstitchState? state = null, CustomResourceOptions? options = null)
        {
            return new Automationstitch(name, id, state, options);
        }
    }

    public sealed class AutomationstitchArgs : global::Pulumi.ResourceArgs
    {
        [Input("action")]
        private InputList<Inputs.AutomationstitchActionArgs>? _action;

        /// <summary>
        /// Action names. The structure of `action` block is documented below.
        /// </summary>
        public InputList<Inputs.AutomationstitchActionArgs> Action
        {
            get => _action ?? (_action = new InputList<Inputs.AutomationstitchActionArgs>());
            set => _action = value;
        }

        [Input("actions")]
        private InputList<Inputs.AutomationstitchActionArgs>? _actions;

        /// <summary>
        /// Configure stitch actions. The structure of `actions` block is documented below.
        /// </summary>
        public InputList<Inputs.AutomationstitchActionArgs> Actions
        {
            get => _actions ?? (_actions = new InputList<Inputs.AutomationstitchActionArgs>());
            set => _actions = value;
        }

        /// <summary>
        /// Description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("destinations")]
        private InputList<Inputs.AutomationstitchDestinationArgs>? _destinations;

        /// <summary>
        /// Serial number/HA group-name of destination devices. The structure of `destination` block is documented below.
        /// </summary>
        public InputList<Inputs.AutomationstitchDestinationArgs> Destinations
        {
            get => _destinations ?? (_destinations = new InputList<Inputs.AutomationstitchDestinationArgs>());
            set => _destinations = value;
        }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Enable/disable this stitch. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status", required: true)]
        public Input<string> Status { get; set; } = null!;

        /// <summary>
        /// Trigger name.
        /// </summary>
        [Input("trigger", required: true)]
        public Input<string> Trigger { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public AutomationstitchArgs()
        {
        }
        public static new AutomationstitchArgs Empty => new AutomationstitchArgs();
    }

    public sealed class AutomationstitchState : global::Pulumi.ResourceArgs
    {
        [Input("action")]
        private InputList<Inputs.AutomationstitchActionGetArgs>? _action;

        /// <summary>
        /// Action names. The structure of `action` block is documented below.
        /// </summary>
        public InputList<Inputs.AutomationstitchActionGetArgs> Action
        {
            get => _action ?? (_action = new InputList<Inputs.AutomationstitchActionGetArgs>());
            set => _action = value;
        }

        [Input("actions")]
        private InputList<Inputs.AutomationstitchActionGetArgs>? _actions;

        /// <summary>
        /// Configure stitch actions. The structure of `actions` block is documented below.
        /// </summary>
        public InputList<Inputs.AutomationstitchActionGetArgs> Actions
        {
            get => _actions ?? (_actions = new InputList<Inputs.AutomationstitchActionGetArgs>());
            set => _actions = value;
        }

        /// <summary>
        /// Description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("destinations")]
        private InputList<Inputs.AutomationstitchDestinationGetArgs>? _destinations;

        /// <summary>
        /// Serial number/HA group-name of destination devices. The structure of `destination` block is documented below.
        /// </summary>
        public InputList<Inputs.AutomationstitchDestinationGetArgs> Destinations
        {
            get => _destinations ?? (_destinations = new InputList<Inputs.AutomationstitchDestinationGetArgs>());
            set => _destinations = value;
        }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Enable/disable this stitch. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Trigger name.
        /// </summary>
        [Input("trigger")]
        public Input<string>? Trigger { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public AutomationstitchState()
        {
        }
        public static new AutomationstitchState Empty => new AutomationstitchState();
    }
}
