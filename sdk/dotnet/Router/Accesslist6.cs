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
    /// Configure IPv6 access lists.
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
    ///     var trname = new Fortios.Router.Accesslist6("trname", new()
    ///     {
    ///         Comments = "access-list6 test",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Router AccessList6 can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:router/accesslist6:Accesslist6 labelname {{name}}
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:router/accesslist6:Accesslist6 labelname {{name}}
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:router/accesslist6:Accesslist6")]
    public partial class Accesslist6 : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Comment.
        /// </summary>
        [Output("comments")]
        public Output<string> Comments { get; private set; } = null!;

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
        /// Rule. The structure of `rule` block is documented below.
        /// </summary>
        [Output("rules")]
        public Output<ImmutableArray<Outputs.Accesslist6Rule>> Rules { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Accesslist6 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Accesslist6(string name, Accesslist6Args? args = null, CustomResourceOptions? options = null)
            : base("fortios:router/accesslist6:Accesslist6", name, args ?? new Accesslist6Args(), MakeResourceOptions(options, ""))
        {
        }

        private Accesslist6(string name, Input<string> id, Accesslist6State? state = null, CustomResourceOptions? options = null)
            : base("fortios:router/accesslist6:Accesslist6", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Accesslist6 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Accesslist6 Get(string name, Input<string> id, Accesslist6State? state = null, CustomResourceOptions? options = null)
        {
            return new Accesslist6(name, id, state, options);
        }
    }

    public sealed class Accesslist6Args : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comments")]
        public Input<string>? Comments { get; set; }

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

        [Input("rules")]
        private InputList<Inputs.Accesslist6RuleArgs>? _rules;

        /// <summary>
        /// Rule. The structure of `rule` block is documented below.
        /// </summary>
        public InputList<Inputs.Accesslist6RuleArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.Accesslist6RuleArgs>());
            set => _rules = value;
        }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public Accesslist6Args()
        {
        }
        public static new Accesslist6Args Empty => new Accesslist6Args();
    }

    public sealed class Accesslist6State : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comments")]
        public Input<string>? Comments { get; set; }

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

        [Input("rules")]
        private InputList<Inputs.Accesslist6RuleGetArgs>? _rules;

        /// <summary>
        /// Rule. The structure of `rule` block is documented below.
        /// </summary>
        public InputList<Inputs.Accesslist6RuleGetArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.Accesslist6RuleGetArgs>());
            set => _rules = value;
        }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public Accesslist6State()
        {
        }
        public static new Accesslist6State Empty => new Accesslist6State();
    }
}
