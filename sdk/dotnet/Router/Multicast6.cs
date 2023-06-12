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
    /// Configure IPv6 multicast.
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
    ///     var trname = new Fortios.Router.Multicast6("trname", new()
    ///     {
    ///         MulticastPmtu = "disable",
    ///         MulticastRouting = "disable",
    ///         PimSmGlobal = new Fortios.Router.Inputs.Multicast6PimSmGlobalArgs
    ///         {
    ///             RegisterRateLimit = 0,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Router Multicast6 can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:router/multicast6:Multicast6 labelname RouterMulticast6
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:router/multicast6:Multicast6 labelname RouterMulticast6
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:router/multicast6:Multicast6")]
    public partial class Multicast6 : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// Protocol Independent Multicast (PIM) interfaces. The structure of `interface` block is documented below.
        /// </summary>
        [Output("interfaces")]
        public Output<ImmutableArray<Outputs.Multicast6Interface>> Interfaces { get; private set; } = null!;

        /// <summary>
        /// Enable/disable PMTU for IPv6 multicast. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("multicastPmtu")]
        public Output<string> MulticastPmtu { get; private set; } = null!;

        /// <summary>
        /// Enable/disable IPv6 multicast routing. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("multicastRouting")]
        public Output<string> MulticastRouting { get; private set; } = null!;

        /// <summary>
        /// PIM sparse-mode global settings. The structure of `pim_sm_global` block is documented below.
        /// </summary>
        [Output("pimSmGlobal")]
        public Output<Outputs.Multicast6PimSmGlobal> PimSmGlobal { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Multicast6 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Multicast6(string name, Multicast6Args? args = null, CustomResourceOptions? options = null)
            : base("fortios:router/multicast6:Multicast6", name, args ?? new Multicast6Args(), MakeResourceOptions(options, ""))
        {
        }

        private Multicast6(string name, Input<string> id, Multicast6State? state = null, CustomResourceOptions? options = null)
            : base("fortios:router/multicast6:Multicast6", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Multicast6 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Multicast6 Get(string name, Input<string> id, Multicast6State? state = null, CustomResourceOptions? options = null)
        {
            return new Multicast6(name, id, state, options);
        }
    }

    public sealed class Multicast6Args : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        [Input("interfaces")]
        private InputList<Inputs.Multicast6InterfaceArgs>? _interfaces;

        /// <summary>
        /// Protocol Independent Multicast (PIM) interfaces. The structure of `interface` block is documented below.
        /// </summary>
        public InputList<Inputs.Multicast6InterfaceArgs> Interfaces
        {
            get => _interfaces ?? (_interfaces = new InputList<Inputs.Multicast6InterfaceArgs>());
            set => _interfaces = value;
        }

        /// <summary>
        /// Enable/disable PMTU for IPv6 multicast. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("multicastPmtu")]
        public Input<string>? MulticastPmtu { get; set; }

        /// <summary>
        /// Enable/disable IPv6 multicast routing. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("multicastRouting")]
        public Input<string>? MulticastRouting { get; set; }

        /// <summary>
        /// PIM sparse-mode global settings. The structure of `pim_sm_global` block is documented below.
        /// </summary>
        [Input("pimSmGlobal")]
        public Input<Inputs.Multicast6PimSmGlobalArgs>? PimSmGlobal { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public Multicast6Args()
        {
        }
        public static new Multicast6Args Empty => new Multicast6Args();
    }

    public sealed class Multicast6State : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        [Input("interfaces")]
        private InputList<Inputs.Multicast6InterfaceGetArgs>? _interfaces;

        /// <summary>
        /// Protocol Independent Multicast (PIM) interfaces. The structure of `interface` block is documented below.
        /// </summary>
        public InputList<Inputs.Multicast6InterfaceGetArgs> Interfaces
        {
            get => _interfaces ?? (_interfaces = new InputList<Inputs.Multicast6InterfaceGetArgs>());
            set => _interfaces = value;
        }

        /// <summary>
        /// Enable/disable PMTU for IPv6 multicast. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("multicastPmtu")]
        public Input<string>? MulticastPmtu { get; set; }

        /// <summary>
        /// Enable/disable IPv6 multicast routing. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("multicastRouting")]
        public Input<string>? MulticastRouting { get; set; }

        /// <summary>
        /// PIM sparse-mode global settings. The structure of `pim_sm_global` block is documented below.
        /// </summary>
        [Input("pimSmGlobal")]
        public Input<Inputs.Multicast6PimSmGlobalGetArgs>? PimSmGlobal { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public Multicast6State()
        {
        }
        public static new Multicast6State Empty => new Multicast6State();
    }
}
