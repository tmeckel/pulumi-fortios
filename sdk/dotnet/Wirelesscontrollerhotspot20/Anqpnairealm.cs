// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Wirelesscontrollerhotspot20
{
    /// <summary>
    /// Configure network access identifier (NAI) realm.
    /// 
    /// ## Import
    /// 
    /// WirelessControllerHotspot20 AnqpNaiRealm can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:wirelesscontrollerhotspot20/anqpnairealm:Anqpnairealm labelname {{name}}
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:wirelesscontrollerhotspot20/anqpnairealm:Anqpnairealm labelname {{name}}
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:wirelesscontrollerhotspot20/anqpnairealm:Anqpnairealm")]
    public partial class Anqpnairealm : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// NAI list. The structure of `nai_list` block is documented below.
        /// </summary>
        [Output("naiLists")]
        public Output<ImmutableArray<Outputs.AnqpnairealmNaiList>> NaiLists { get; private set; } = null!;

        /// <summary>
        /// NAI realm list name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Anqpnairealm resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Anqpnairealm(string name, AnqpnairealmArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:wirelesscontrollerhotspot20/anqpnairealm:Anqpnairealm", name, args ?? new AnqpnairealmArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Anqpnairealm(string name, Input<string> id, AnqpnairealmState? state = null, CustomResourceOptions? options = null)
            : base("fortios:wirelesscontrollerhotspot20/anqpnairealm:Anqpnairealm", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Anqpnairealm resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Anqpnairealm Get(string name, Input<string> id, AnqpnairealmState? state = null, CustomResourceOptions? options = null)
        {
            return new Anqpnairealm(name, id, state, options);
        }
    }

    public sealed class AnqpnairealmArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        [Input("naiLists")]
        private InputList<Inputs.AnqpnairealmNaiListArgs>? _naiLists;

        /// <summary>
        /// NAI list. The structure of `nai_list` block is documented below.
        /// </summary>
        public InputList<Inputs.AnqpnairealmNaiListArgs> NaiLists
        {
            get => _naiLists ?? (_naiLists = new InputList<Inputs.AnqpnairealmNaiListArgs>());
            set => _naiLists = value;
        }

        /// <summary>
        /// NAI realm list name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public AnqpnairealmArgs()
        {
        }
        public static new AnqpnairealmArgs Empty => new AnqpnairealmArgs();
    }

    public sealed class AnqpnairealmState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        [Input("naiLists")]
        private InputList<Inputs.AnqpnairealmNaiListGetArgs>? _naiLists;

        /// <summary>
        /// NAI list. The structure of `nai_list` block is documented below.
        /// </summary>
        public InputList<Inputs.AnqpnairealmNaiListGetArgs> NaiLists
        {
            get => _naiLists ?? (_naiLists = new InputList<Inputs.AnqpnairealmNaiListGetArgs>());
            set => _naiLists = value;
        }

        /// <summary>
        /// NAI realm list name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public AnqpnairealmState()
        {
        }
        public static new AnqpnairealmState Empty => new AnqpnairealmState();
    }
}
