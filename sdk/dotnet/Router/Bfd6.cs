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
    /// Configure IPv6 BFD.
    /// 
    /// ## Import
    /// 
    /// Router Bfd6 can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:router/bfd6:Bfd6 labelname RouterBfd6
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:router/bfd6:Bfd6 labelname RouterBfd6
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:router/bfd6:Bfd6")]
    public partial class Bfd6 : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// BFD IPv6 multi-hop template table. The structure of `multihop_template` block is documented below.
        /// </summary>
        [Output("multihopTemplates")]
        public Output<ImmutableArray<Outputs.Bfd6MultihopTemplate>> MultihopTemplates { get; private set; } = null!;

        /// <summary>
        /// Configure neighbor of IPv6 BFD. The structure of `neighbor` block is documented below.
        /// </summary>
        [Output("neighbors")]
        public Output<ImmutableArray<Outputs.Bfd6Neighbor>> Neighbors { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Bfd6 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Bfd6(string name, Bfd6Args? args = null, CustomResourceOptions? options = null)
            : base("fortios:router/bfd6:Bfd6", name, args ?? new Bfd6Args(), MakeResourceOptions(options, ""))
        {
        }

        private Bfd6(string name, Input<string> id, Bfd6State? state = null, CustomResourceOptions? options = null)
            : base("fortios:router/bfd6:Bfd6", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Bfd6 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Bfd6 Get(string name, Input<string> id, Bfd6State? state = null, CustomResourceOptions? options = null)
        {
            return new Bfd6(name, id, state, options);
        }
    }

    public sealed class Bfd6Args : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        [Input("multihopTemplates")]
        private InputList<Inputs.Bfd6MultihopTemplateArgs>? _multihopTemplates;

        /// <summary>
        /// BFD IPv6 multi-hop template table. The structure of `multihop_template` block is documented below.
        /// </summary>
        public InputList<Inputs.Bfd6MultihopTemplateArgs> MultihopTemplates
        {
            get => _multihopTemplates ?? (_multihopTemplates = new InputList<Inputs.Bfd6MultihopTemplateArgs>());
            set => _multihopTemplates = value;
        }

        [Input("neighbors")]
        private InputList<Inputs.Bfd6NeighborArgs>? _neighbors;

        /// <summary>
        /// Configure neighbor of IPv6 BFD. The structure of `neighbor` block is documented below.
        /// </summary>
        public InputList<Inputs.Bfd6NeighborArgs> Neighbors
        {
            get => _neighbors ?? (_neighbors = new InputList<Inputs.Bfd6NeighborArgs>());
            set => _neighbors = value;
        }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public Bfd6Args()
        {
        }
        public static new Bfd6Args Empty => new Bfd6Args();
    }

    public sealed class Bfd6State : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        [Input("multihopTemplates")]
        private InputList<Inputs.Bfd6MultihopTemplateGetArgs>? _multihopTemplates;

        /// <summary>
        /// BFD IPv6 multi-hop template table. The structure of `multihop_template` block is documented below.
        /// </summary>
        public InputList<Inputs.Bfd6MultihopTemplateGetArgs> MultihopTemplates
        {
            get => _multihopTemplates ?? (_multihopTemplates = new InputList<Inputs.Bfd6MultihopTemplateGetArgs>());
            set => _multihopTemplates = value;
        }

        [Input("neighbors")]
        private InputList<Inputs.Bfd6NeighborGetArgs>? _neighbors;

        /// <summary>
        /// Configure neighbor of IPv6 BFD. The structure of `neighbor` block is documented below.
        /// </summary>
        public InputList<Inputs.Bfd6NeighborGetArgs> Neighbors
        {
            get => _neighbors ?? (_neighbors = new InputList<Inputs.Bfd6NeighborGetArgs>());
            set => _neighbors = value;
        }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public Bfd6State()
        {
        }
        public static new Bfd6State Empty => new Bfd6State();
    }
}
