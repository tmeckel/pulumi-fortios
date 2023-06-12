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
    /// Configure FortiToken Mobile push services.
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
    ///     var trname = new Fortios.Sys.Ftmpush("trname", new()
    ///     {
    ///         ServerIp = "0.0.0.0",
    ///         ServerPort = 4433,
    ///         Status = "disable",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// System FtmPush can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:sys/ftmpush:Ftmpush labelname SystemFtmPush
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:sys/ftmpush:Ftmpush labelname SystemFtmPush
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:sys/ftmpush:Ftmpush")]
    public partial class Ftmpush : global::Pulumi.CustomResource
    {
        /// <summary>
        /// IPv4 address or domain name of FortiToken Mobile push services server.
        /// </summary>
        [Output("server")]
        public Output<string> Server { get; private set; } = null!;

        /// <summary>
        /// Name of the server certificate to be used for SSL (default = Fortinet_Factory).
        /// </summary>
        [Output("serverCert")]
        public Output<string> ServerCert { get; private set; } = null!;

        /// <summary>
        /// IPv4 address of FortiToken Mobile push services server (format: xxx.xxx.xxx.xxx).
        /// </summary>
        [Output("serverIp")]
        public Output<string> ServerIp { get; private set; } = null!;

        /// <summary>
        /// Port to communicate with FortiToken Mobile push services server (1 - 65535, default = 4433).
        /// </summary>
        [Output("serverPort")]
        public Output<int> ServerPort { get; private set; } = null!;

        /// <summary>
        /// Enable/disable the use of FortiToken Mobile push services. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Ftmpush resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Ftmpush(string name, FtmpushArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:sys/ftmpush:Ftmpush", name, args ?? new FtmpushArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Ftmpush(string name, Input<string> id, FtmpushState? state = null, CustomResourceOptions? options = null)
            : base("fortios:sys/ftmpush:Ftmpush", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Ftmpush resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Ftmpush Get(string name, Input<string> id, FtmpushState? state = null, CustomResourceOptions? options = null)
        {
            return new Ftmpush(name, id, state, options);
        }
    }

    public sealed class FtmpushArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// IPv4 address or domain name of FortiToken Mobile push services server.
        /// </summary>
        [Input("server")]
        public Input<string>? Server { get; set; }

        /// <summary>
        /// Name of the server certificate to be used for SSL (default = Fortinet_Factory).
        /// </summary>
        [Input("serverCert")]
        public Input<string>? ServerCert { get; set; }

        /// <summary>
        /// IPv4 address of FortiToken Mobile push services server (format: xxx.xxx.xxx.xxx).
        /// </summary>
        [Input("serverIp")]
        public Input<string>? ServerIp { get; set; }

        /// <summary>
        /// Port to communicate with FortiToken Mobile push services server (1 - 65535, default = 4433).
        /// </summary>
        [Input("serverPort")]
        public Input<int>? ServerPort { get; set; }

        /// <summary>
        /// Enable/disable the use of FortiToken Mobile push services. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public FtmpushArgs()
        {
        }
        public static new FtmpushArgs Empty => new FtmpushArgs();
    }

    public sealed class FtmpushState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// IPv4 address or domain name of FortiToken Mobile push services server.
        /// </summary>
        [Input("server")]
        public Input<string>? Server { get; set; }

        /// <summary>
        /// Name of the server certificate to be used for SSL (default = Fortinet_Factory).
        /// </summary>
        [Input("serverCert")]
        public Input<string>? ServerCert { get; set; }

        /// <summary>
        /// IPv4 address of FortiToken Mobile push services server (format: xxx.xxx.xxx.xxx).
        /// </summary>
        [Input("serverIp")]
        public Input<string>? ServerIp { get; set; }

        /// <summary>
        /// Port to communicate with FortiToken Mobile push services server (1 - 65535, default = 4433).
        /// </summary>
        [Input("serverPort")]
        public Input<int>? ServerPort { get; set; }

        /// <summary>
        /// Enable/disable the use of FortiToken Mobile push services. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public FtmpushState()
        {
        }
        public static new FtmpushState Empty => new FtmpushState();
    }
}