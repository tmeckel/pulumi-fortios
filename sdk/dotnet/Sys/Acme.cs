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
    /// Configure ACME client. Applies to FortiOS Version `&gt;= 7.0.1`.
    /// 
    /// ## Import
    /// 
    /// System Acme can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:sys/acme:Acme labelname SystemAcme
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:sys/acme:Acme labelname SystemAcme
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:sys/acme:Acme")]
    public partial class Acme : global::Pulumi.CustomResource
    {
        /// <summary>
        /// ACME accounts list. The structure of `accounts` block is documented below.
        /// </summary>
        [Output("accounts")]
        public Output<ImmutableArray<Outputs.AcmeAccount>> Accounts { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// Interface(s) on which the ACME client will listen for challenges. The structure of `interface` block is documented below.
        /// </summary>
        [Output("interfaces")]
        public Output<ImmutableArray<Outputs.AcmeInterface>> Interfaces { get; private set; } = null!;

        /// <summary>
        /// Source IPv4 address used to connect to the ACME server.
        /// </summary>
        [Output("sourceIp")]
        public Output<string> SourceIp { get; private set; } = null!;

        /// <summary>
        /// Source IPv6 address used to connect to the ACME server.
        /// </summary>
        [Output("sourceIp6")]
        public Output<string> SourceIp6 { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Acme resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Acme(string name, AcmeArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:sys/acme:Acme", name, args ?? new AcmeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Acme(string name, Input<string> id, AcmeState? state = null, CustomResourceOptions? options = null)
            : base("fortios:sys/acme:Acme", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Acme resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Acme Get(string name, Input<string> id, AcmeState? state = null, CustomResourceOptions? options = null)
        {
            return new Acme(name, id, state, options);
        }
    }

    public sealed class AcmeArgs : global::Pulumi.ResourceArgs
    {
        [Input("accounts")]
        private InputList<Inputs.AcmeAccountArgs>? _accounts;

        /// <summary>
        /// ACME accounts list. The structure of `accounts` block is documented below.
        /// </summary>
        public InputList<Inputs.AcmeAccountArgs> Accounts
        {
            get => _accounts ?? (_accounts = new InputList<Inputs.AcmeAccountArgs>());
            set => _accounts = value;
        }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        [Input("interfaces")]
        private InputList<Inputs.AcmeInterfaceArgs>? _interfaces;

        /// <summary>
        /// Interface(s) on which the ACME client will listen for challenges. The structure of `interface` block is documented below.
        /// </summary>
        public InputList<Inputs.AcmeInterfaceArgs> Interfaces
        {
            get => _interfaces ?? (_interfaces = new InputList<Inputs.AcmeInterfaceArgs>());
            set => _interfaces = value;
        }

        /// <summary>
        /// Source IPv4 address used to connect to the ACME server.
        /// </summary>
        [Input("sourceIp")]
        public Input<string>? SourceIp { get; set; }

        /// <summary>
        /// Source IPv6 address used to connect to the ACME server.
        /// </summary>
        [Input("sourceIp6")]
        public Input<string>? SourceIp6 { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public AcmeArgs()
        {
        }
        public static new AcmeArgs Empty => new AcmeArgs();
    }

    public sealed class AcmeState : global::Pulumi.ResourceArgs
    {
        [Input("accounts")]
        private InputList<Inputs.AcmeAccountGetArgs>? _accounts;

        /// <summary>
        /// ACME accounts list. The structure of `accounts` block is documented below.
        /// </summary>
        public InputList<Inputs.AcmeAccountGetArgs> Accounts
        {
            get => _accounts ?? (_accounts = new InputList<Inputs.AcmeAccountGetArgs>());
            set => _accounts = value;
        }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        [Input("interfaces")]
        private InputList<Inputs.AcmeInterfaceGetArgs>? _interfaces;

        /// <summary>
        /// Interface(s) on which the ACME client will listen for challenges. The structure of `interface` block is documented below.
        /// </summary>
        public InputList<Inputs.AcmeInterfaceGetArgs> Interfaces
        {
            get => _interfaces ?? (_interfaces = new InputList<Inputs.AcmeInterfaceGetArgs>());
            set => _interfaces = value;
        }

        /// <summary>
        /// Source IPv4 address used to connect to the ACME server.
        /// </summary>
        [Input("sourceIp")]
        public Input<string>? SourceIp { get; set; }

        /// <summary>
        /// Source IPv6 address used to connect to the ACME server.
        /// </summary>
        [Input("sourceIp6")]
        public Input<string>? SourceIp6 { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public AcmeState()
        {
        }
        public static new AcmeState Empty => new AcmeState();
    }
}
