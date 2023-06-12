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
    /// Configure IPv6 IP pools.
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
    ///     var trname = new Fortios.Firewall.Ippool6("trname", new()
    ///     {
    ///         Endip = "2001:3ca1:10f:1a:121b::19",
    ///         Startip = "2001:3ca1:10f:1a:121b::10",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Firewall Ippool6 can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:firewall/ippool6:Ippool6 labelname {{name}}
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:firewall/ippool6:Ippool6 labelname {{name}}
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:firewall/ippool6:Ippool6")]
    public partial class Ippool6 : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Enable/disable adding NAT46 route. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("addNat46Route")]
        public Output<string> AddNat46Route { get; private set; } = null!;

        /// <summary>
        /// Comment.
        /// </summary>
        [Output("comments")]
        public Output<string?> Comments { get; private set; } = null!;

        /// <summary>
        /// Final IPv6 address (inclusive) in the range for the address pool (format xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx, Default: ::).
        /// </summary>
        [Output("endip")]
        public Output<string> Endip { get; private set; } = null!;

        /// <summary>
        /// IPv6 IP pool name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Enable/disable NAT46. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("nat46")]
        public Output<string> Nat46 { get; private set; } = null!;

        /// <summary>
        /// First IPv6 address (inclusive) in the range for the address pool (format xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx, Default: ::).
        /// </summary>
        [Output("startip")]
        public Output<string> Startip { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Ippool6 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Ippool6(string name, Ippool6Args args, CustomResourceOptions? options = null)
            : base("fortios:firewall/ippool6:Ippool6", name, args ?? new Ippool6Args(), MakeResourceOptions(options, ""))
        {
        }

        private Ippool6(string name, Input<string> id, Ippool6State? state = null, CustomResourceOptions? options = null)
            : base("fortios:firewall/ippool6:Ippool6", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Ippool6 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Ippool6 Get(string name, Input<string> id, Ippool6State? state = null, CustomResourceOptions? options = null)
        {
            return new Ippool6(name, id, state, options);
        }
    }

    public sealed class Ippool6Args : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable adding NAT46 route. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("addNat46Route")]
        public Input<string>? AddNat46Route { get; set; }

        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comments")]
        public Input<string>? Comments { get; set; }

        /// <summary>
        /// Final IPv6 address (inclusive) in the range for the address pool (format xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx, Default: ::).
        /// </summary>
        [Input("endip", required: true)]
        public Input<string> Endip { get; set; } = null!;

        /// <summary>
        /// IPv6 IP pool name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Enable/disable NAT46. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("nat46")]
        public Input<string>? Nat46 { get; set; }

        /// <summary>
        /// First IPv6 address (inclusive) in the range for the address pool (format xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx, Default: ::).
        /// </summary>
        [Input("startip", required: true)]
        public Input<string> Startip { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public Ippool6Args()
        {
        }
        public static new Ippool6Args Empty => new Ippool6Args();
    }

    public sealed class Ippool6State : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable adding NAT46 route. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("addNat46Route")]
        public Input<string>? AddNat46Route { get; set; }

        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comments")]
        public Input<string>? Comments { get; set; }

        /// <summary>
        /// Final IPv6 address (inclusive) in the range for the address pool (format xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx, Default: ::).
        /// </summary>
        [Input("endip")]
        public Input<string>? Endip { get; set; }

        /// <summary>
        /// IPv6 IP pool name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Enable/disable NAT46. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("nat46")]
        public Input<string>? Nat46 { get; set; }

        /// <summary>
        /// First IPv6 address (inclusive) in the range for the address pool (format xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx, Default: ::).
        /// </summary>
        [Input("startip")]
        public Input<string>? Startip { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public Ippool6State()
        {
        }
        public static new Ippool6State Empty => new Ippool6State();
    }
}
