// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.User
{
    /// <summary>
    /// Configure peer groups.
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
    ///     var trname2 = new Fortios.User.Peer("trname2", new()
    ///     {
    ///         Ca = "EC-ACC",
    ///         CnType = "string",
    ///         LdapMode = "password",
    ///         MandatoryCaVerify = "enable",
    ///         TwoFactor = "disable",
    ///     });
    /// 
    ///     var trname = new Fortios.User.Peergrp("trname", new()
    ///     {
    ///         Members = new[]
    ///         {
    ///             new Fortios.User.Inputs.PeergrpMemberArgs
    ///             {
    ///                 Name = trname2.Name,
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// User Peergrp can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:user/peergrp:Peergrp labelname {{name}}
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:user/peergrp:Peergrp labelname {{name}}
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:user/peergrp:Peergrp")]
    public partial class Peergrp : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// Peer group members. The structure of `member` block is documented below.
        /// </summary>
        [Output("members")]
        public Output<ImmutableArray<Outputs.PeergrpMember>> Members { get; private set; } = null!;

        /// <summary>
        /// Peer group name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Peergrp resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Peergrp(string name, PeergrpArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:user/peergrp:Peergrp", name, args ?? new PeergrpArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Peergrp(string name, Input<string> id, PeergrpState? state = null, CustomResourceOptions? options = null)
            : base("fortios:user/peergrp:Peergrp", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Peergrp resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Peergrp Get(string name, Input<string> id, PeergrpState? state = null, CustomResourceOptions? options = null)
        {
            return new Peergrp(name, id, state, options);
        }
    }

    public sealed class PeergrpArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        [Input("members")]
        private InputList<Inputs.PeergrpMemberArgs>? _members;

        /// <summary>
        /// Peer group members. The structure of `member` block is documented below.
        /// </summary>
        public InputList<Inputs.PeergrpMemberArgs> Members
        {
            get => _members ?? (_members = new InputList<Inputs.PeergrpMemberArgs>());
            set => _members = value;
        }

        /// <summary>
        /// Peer group name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public PeergrpArgs()
        {
        }
        public static new PeergrpArgs Empty => new PeergrpArgs();
    }

    public sealed class PeergrpState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        [Input("members")]
        private InputList<Inputs.PeergrpMemberGetArgs>? _members;

        /// <summary>
        /// Peer group members. The structure of `member` block is documented below.
        /// </summary>
        public InputList<Inputs.PeergrpMemberGetArgs> Members
        {
            get => _members ?? (_members = new InputList<Inputs.PeergrpMemberGetArgs>());
            set => _members = value;
        }

        /// <summary>
        /// Peer group name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public PeergrpState()
        {
        }
        public static new PeergrpState Empty => new PeergrpState();
    }
}
