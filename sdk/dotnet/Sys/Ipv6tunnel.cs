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
    /// Configure IPv6/IPv4 in IPv6 tunnel.
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
    ///     var trname = new Fortios.Sys.Ipv6tunnel("trname", new()
    ///     {
    ///         Destination = "2001:db8:85a3::8a2e:370:7324",
    ///         Interface = "port3",
    ///         Source = "2001:db8:85a3::8a2e:370:7334",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// System Ipv6Tunnel can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:sys/ipv6tunnel:Ipv6tunnel labelname {{name}}
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:sys/ipv6tunnel:Ipv6tunnel labelname {{name}}
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:sys/ipv6tunnel:Ipv6tunnel")]
    public partial class Ipv6tunnel : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Enable/disable tunnel ASIC offloading. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("autoAsicOffload")]
        public Output<string> AutoAsicOffload { get; private set; } = null!;

        /// <summary>
        /// Remote IPv6 address of the tunnel.
        /// </summary>
        [Output("destination")]
        public Output<string> Destination { get; private set; } = null!;

        /// <summary>
        /// Interface name.
        /// </summary>
        [Output("interface")]
        public Output<string> Interface { get; private set; } = null!;

        /// <summary>
        /// IPv6 tunnel name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Local IPv6 address of the tunnel.
        /// </summary>
        [Output("source")]
        public Output<string> Source { get; private set; } = null!;

        /// <summary>
        /// Enable/disable use of SD-WAN to reach remote gateway. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("useSdwan")]
        public Output<string> UseSdwan { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Ipv6tunnel resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Ipv6tunnel(string name, Ipv6tunnelArgs args, CustomResourceOptions? options = null)
            : base("fortios:sys/ipv6tunnel:Ipv6tunnel", name, args ?? new Ipv6tunnelArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Ipv6tunnel(string name, Input<string> id, Ipv6tunnelState? state = null, CustomResourceOptions? options = null)
            : base("fortios:sys/ipv6tunnel:Ipv6tunnel", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Ipv6tunnel resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Ipv6tunnel Get(string name, Input<string> id, Ipv6tunnelState? state = null, CustomResourceOptions? options = null)
        {
            return new Ipv6tunnel(name, id, state, options);
        }
    }

    public sealed class Ipv6tunnelArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable tunnel ASIC offloading. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("autoAsicOffload")]
        public Input<string>? AutoAsicOffload { get; set; }

        /// <summary>
        /// Remote IPv6 address of the tunnel.
        /// </summary>
        [Input("destination", required: true)]
        public Input<string> Destination { get; set; } = null!;

        /// <summary>
        /// Interface name.
        /// </summary>
        [Input("interface")]
        public Input<string>? Interface { get; set; }

        /// <summary>
        /// IPv6 tunnel name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Local IPv6 address of the tunnel.
        /// </summary>
        [Input("source")]
        public Input<string>? Source { get; set; }

        /// <summary>
        /// Enable/disable use of SD-WAN to reach remote gateway. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("useSdwan")]
        public Input<string>? UseSdwan { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public Ipv6tunnelArgs()
        {
        }
        public static new Ipv6tunnelArgs Empty => new Ipv6tunnelArgs();
    }

    public sealed class Ipv6tunnelState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable tunnel ASIC offloading. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("autoAsicOffload")]
        public Input<string>? AutoAsicOffload { get; set; }

        /// <summary>
        /// Remote IPv6 address of the tunnel.
        /// </summary>
        [Input("destination")]
        public Input<string>? Destination { get; set; }

        /// <summary>
        /// Interface name.
        /// </summary>
        [Input("interface")]
        public Input<string>? Interface { get; set; }

        /// <summary>
        /// IPv6 tunnel name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Local IPv6 address of the tunnel.
        /// </summary>
        [Input("source")]
        public Input<string>? Source { get; set; }

        /// <summary>
        /// Enable/disable use of SD-WAN to reach remote gateway. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("useSdwan")]
        public Input<string>? UseSdwan { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public Ipv6tunnelState()
        {
        }
        public static new Ipv6tunnelState Empty => new Ipv6tunnelState();
    }
}
