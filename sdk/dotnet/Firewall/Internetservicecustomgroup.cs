// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Firewall
{
    /// <summary>
    /// Configure custom Internet Service group.
    /// 
    /// ## Import
    /// 
    /// Firewall InternetServiceCustomGroup can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:firewall/internetservicecustomgroup:Internetservicecustomgroup labelname {{name}}
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:firewall/internetservicecustomgroup:Internetservicecustomgroup labelname {{name}}
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:firewall/internetservicecustomgroup:Internetservicecustomgroup")]
    public partial class Internetservicecustomgroup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Comment.
        /// </summary>
        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// Custom Internet Service group members. The structure of `member` block is documented below.
        /// </summary>
        [Output("members")]
        public Output<ImmutableArray<Outputs.InternetservicecustomgroupMember>> Members { get; private set; } = null!;

        /// <summary>
        /// Custom Internet Service group name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Internetservicecustomgroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Internetservicecustomgroup(string name, InternetservicecustomgroupArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:firewall/internetservicecustomgroup:Internetservicecustomgroup", name, args ?? new InternetservicecustomgroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Internetservicecustomgroup(string name, Input<string> id, InternetservicecustomgroupState? state = null, CustomResourceOptions? options = null)
            : base("fortios:firewall/internetservicecustomgroup:Internetservicecustomgroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Internetservicecustomgroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Internetservicecustomgroup Get(string name, Input<string> id, InternetservicecustomgroupState? state = null, CustomResourceOptions? options = null)
        {
            return new Internetservicecustomgroup(name, id, state, options);
        }
    }

    public sealed class InternetservicecustomgroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        [Input("members")]
        private InputList<Inputs.InternetservicecustomgroupMemberArgs>? _members;

        /// <summary>
        /// Custom Internet Service group members. The structure of `member` block is documented below.
        /// </summary>
        public InputList<Inputs.InternetservicecustomgroupMemberArgs> Members
        {
            get => _members ?? (_members = new InputList<Inputs.InternetservicecustomgroupMemberArgs>());
            set => _members = value;
        }

        /// <summary>
        /// Custom Internet Service group name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public InternetservicecustomgroupArgs()
        {
        }
        public static new InternetservicecustomgroupArgs Empty => new InternetservicecustomgroupArgs();
    }

    public sealed class InternetservicecustomgroupState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        [Input("members")]
        private InputList<Inputs.InternetservicecustomgroupMemberGetArgs>? _members;

        /// <summary>
        /// Custom Internet Service group members. The structure of `member` block is documented below.
        /// </summary>
        public InputList<Inputs.InternetservicecustomgroupMemberGetArgs> Members
        {
            get => _members ?? (_members = new InputList<Inputs.InternetservicecustomgroupMemberGetArgs>());
            set => _members = value;
        }

        /// <summary>
        /// Custom Internet Service group name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public InternetservicecustomgroupState()
        {
        }
        public static new InternetservicecustomgroupState Empty => new InternetservicecustomgroupState();
    }
}