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
    /// Configure IPv6 to IPv4 virtual IP groups. Applies to FortiOS Version `&lt;= 7.0.0`.
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
    ///     var trname1 = new Fortios.Firewall.Vip64("trname1", new()
    ///     {
    ///         ArpReply = "enable",
    ///         Color = 0,
    ///         Extip = "2001:db8:99:503::22",
    ///         Extport = "0-65535",
    ///         Fosid = 0,
    ///         LdbMethod = "static",
    ///         Mappedip = "1.1.3.1",
    ///         Mappedport = "0-65535",
    ///         Portforward = "disable",
    ///         Protocol = "tcp",
    ///         Type = "static-nat",
    ///     });
    /// 
    ///     var trname = new Fortios.Firewall.Vipgrp64("trname", new()
    ///     {
    ///         Color = 0,
    ///         Members = new[]
    ///         {
    ///             new Fortios.Firewall.Inputs.Vipgrp64MemberArgs
    ///             {
    ///                 Name = trname1.Name,
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Firewall Vipgrp64 can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:firewall/vipgrp64:Vipgrp64 labelname {{name}}
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:firewall/vipgrp64:Vipgrp64 labelname {{name}}
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:firewall/vipgrp64:Vipgrp64")]
    public partial class Vipgrp64 : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
        /// </summary>
        [Output("color")]
        public Output<int> Color { get; private set; } = null!;

        /// <summary>
        /// Comment.
        /// </summary>
        [Output("comments")]
        public Output<string?> Comments { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// Member VIP objects of the group (Separate multiple objects with a space). The structure of `member` block is documented below.
        /// </summary>
        [Output("members")]
        public Output<ImmutableArray<Outputs.Vipgrp64Member>> Members { get; private set; } = null!;

        /// <summary>
        /// VIP64 group name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        /// </summary>
        [Output("uuid")]
        public Output<string> Uuid { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Vipgrp64 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Vipgrp64(string name, Vipgrp64Args args, CustomResourceOptions? options = null)
            : base("fortios:firewall/vipgrp64:Vipgrp64", name, args ?? new Vipgrp64Args(), MakeResourceOptions(options, ""))
        {
        }

        private Vipgrp64(string name, Input<string> id, Vipgrp64State? state = null, CustomResourceOptions? options = null)
            : base("fortios:firewall/vipgrp64:Vipgrp64", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Vipgrp64 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Vipgrp64 Get(string name, Input<string> id, Vipgrp64State? state = null, CustomResourceOptions? options = null)
        {
            return new Vipgrp64(name, id, state, options);
        }
    }

    public sealed class Vipgrp64Args : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
        /// </summary>
        [Input("color")]
        public Input<int>? Color { get; set; }

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

        [Input("members", required: true)]
        private InputList<Inputs.Vipgrp64MemberArgs>? _members;

        /// <summary>
        /// Member VIP objects of the group (Separate multiple objects with a space). The structure of `member` block is documented below.
        /// </summary>
        public InputList<Inputs.Vipgrp64MemberArgs> Members
        {
            get => _members ?? (_members = new InputList<Inputs.Vipgrp64MemberArgs>());
            set => _members = value;
        }

        /// <summary>
        /// VIP64 group name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        /// </summary>
        [Input("uuid")]
        public Input<string>? Uuid { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public Vipgrp64Args()
        {
        }
        public static new Vipgrp64Args Empty => new Vipgrp64Args();
    }

    public sealed class Vipgrp64State : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
        /// </summary>
        [Input("color")]
        public Input<int>? Color { get; set; }

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

        [Input("members")]
        private InputList<Inputs.Vipgrp64MemberGetArgs>? _members;

        /// <summary>
        /// Member VIP objects of the group (Separate multiple objects with a space). The structure of `member` block is documented below.
        /// </summary>
        public InputList<Inputs.Vipgrp64MemberGetArgs> Members
        {
            get => _members ?? (_members = new InputList<Inputs.Vipgrp64MemberGetArgs>());
            set => _members = value;
        }

        /// <summary>
        /// VIP64 group name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        /// </summary>
        [Input("uuid")]
        public Input<string>? Uuid { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public Vipgrp64State()
        {
        }
        public static new Vipgrp64State Empty => new Vipgrp64State();
    }
}
