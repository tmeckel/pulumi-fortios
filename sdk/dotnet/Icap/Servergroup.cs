// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Icap
{
    /// <summary>
    /// Configure an ICAP server group consisting of multiple forward servers. Supports failover and load balancing. Applies to FortiOS Version `&gt;= 7.2.0`.
    /// 
    /// ## Import
    /// 
    /// Icap ServerGroup can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:icap/servergroup:Servergroup labelname {{name}}
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:icap/servergroup:Servergroup labelname {{name}}
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:icap/servergroup:Servergroup")]
    public partial class Servergroup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// Load balance method. Valid values: `weighted`, `least-session`, `active-passive`.
        /// </summary>
        [Output("ldbMethod")]
        public Output<string> LdbMethod { get; private set; } = null!;

        /// <summary>
        /// Configure an ICAP server group consisting one or multiple forward servers. Supports failover and load balancing.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Add ICAP servers to a list to form a server group. Optionally assign weights to each server. The structure of `server_list` block is documented below.
        /// </summary>
        [Output("serverLists")]
        public Output<ImmutableArray<Outputs.ServergroupServerList>> ServerLists { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Servergroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Servergroup(string name, ServergroupArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:icap/servergroup:Servergroup", name, args ?? new ServergroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Servergroup(string name, Input<string> id, ServergroupState? state = null, CustomResourceOptions? options = null)
            : base("fortios:icap/servergroup:Servergroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Servergroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Servergroup Get(string name, Input<string> id, ServergroupState? state = null, CustomResourceOptions? options = null)
        {
            return new Servergroup(name, id, state, options);
        }
    }

    public sealed class ServergroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Load balance method. Valid values: `weighted`, `least-session`, `active-passive`.
        /// </summary>
        [Input("ldbMethod")]
        public Input<string>? LdbMethod { get; set; }

        /// <summary>
        /// Configure an ICAP server group consisting one or multiple forward servers. Supports failover and load balancing.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("serverLists")]
        private InputList<Inputs.ServergroupServerListArgs>? _serverLists;

        /// <summary>
        /// Add ICAP servers to a list to form a server group. Optionally assign weights to each server. The structure of `server_list` block is documented below.
        /// </summary>
        public InputList<Inputs.ServergroupServerListArgs> ServerLists
        {
            get => _serverLists ?? (_serverLists = new InputList<Inputs.ServergroupServerListArgs>());
            set => _serverLists = value;
        }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public ServergroupArgs()
        {
        }
        public static new ServergroupArgs Empty => new ServergroupArgs();
    }

    public sealed class ServergroupState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Load balance method. Valid values: `weighted`, `least-session`, `active-passive`.
        /// </summary>
        [Input("ldbMethod")]
        public Input<string>? LdbMethod { get; set; }

        /// <summary>
        /// Configure an ICAP server group consisting one or multiple forward servers. Supports failover and load balancing.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("serverLists")]
        private InputList<Inputs.ServergroupServerListGetArgs>? _serverLists;

        /// <summary>
        /// Add ICAP servers to a list to form a server group. Optionally assign weights to each server. The structure of `server_list` block is documented below.
        /// </summary>
        public InputList<Inputs.ServergroupServerListGetArgs> ServerLists
        {
            get => _serverLists ?? (_serverLists = new InputList<Inputs.ServergroupServerListGetArgs>());
            set => _serverLists = value;
        }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public ServergroupState()
        {
        }
        public static new ServergroupState Empty => new ServergroupState();
    }
}
