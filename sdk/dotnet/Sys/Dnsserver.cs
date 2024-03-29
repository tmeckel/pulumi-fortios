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
    /// Configure DNS servers.
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
    ///     var trname = new Fortios.Sys.Dnsserver("trname", new()
    ///     {
    ///         DnsfilterProfile = "default",
    ///         Mode = "forward-only",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// System DnsServer can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:sys/dnsserver:Dnsserver labelname {{name}}
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:sys/dnsserver:Dnsserver labelname {{name}}
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:sys/dnsserver:Dnsserver")]
    public partial class Dnsserver : global::Pulumi.CustomResource
    {
        /// <summary>
        /// DNS filter profile.
        /// </summary>
        [Output("dnsfilterProfile")]
        public Output<string> DnsfilterProfile { get; private set; } = null!;

        /// <summary>
        /// DNS over HTTPS. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("doh")]
        public Output<string> Doh { get; private set; } = null!;

        /// <summary>
        /// DNS server mode. Valid values: `recursive`, `non-recursive`, `forward-only`.
        /// </summary>
        [Output("mode")]
        public Output<string> Mode { get; private set; } = null!;

        /// <summary>
        /// DNS server name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Dnsserver resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Dnsserver(string name, DnsserverArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:sys/dnsserver:Dnsserver", name, args ?? new DnsserverArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Dnsserver(string name, Input<string> id, DnsserverState? state = null, CustomResourceOptions? options = null)
            : base("fortios:sys/dnsserver:Dnsserver", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Dnsserver resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Dnsserver Get(string name, Input<string> id, DnsserverState? state = null, CustomResourceOptions? options = null)
        {
            return new Dnsserver(name, id, state, options);
        }
    }

    public sealed class DnsserverArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// DNS filter profile.
        /// </summary>
        [Input("dnsfilterProfile")]
        public Input<string>? DnsfilterProfile { get; set; }

        /// <summary>
        /// DNS over HTTPS. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("doh")]
        public Input<string>? Doh { get; set; }

        /// <summary>
        /// DNS server mode. Valid values: `recursive`, `non-recursive`, `forward-only`.
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        /// <summary>
        /// DNS server name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public DnsserverArgs()
        {
        }
        public static new DnsserverArgs Empty => new DnsserverArgs();
    }

    public sealed class DnsserverState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// DNS filter profile.
        /// </summary>
        [Input("dnsfilterProfile")]
        public Input<string>? DnsfilterProfile { get; set; }

        /// <summary>
        /// DNS over HTTPS. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("doh")]
        public Input<string>? Doh { get; set; }

        /// <summary>
        /// DNS server mode. Valid values: `recursive`, `non-recursive`, `forward-only`.
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        /// <summary>
        /// DNS server name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public DnsserverState()
        {
        }
        public static new DnsserverState Empty => new DnsserverState();
    }
}
