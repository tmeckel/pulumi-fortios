// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Webfilter
{
    /// <summary>
    /// Configure local FortiGuard Web Filter local ratings.
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
    ///     var trname = new Fortios.Webfilter.Ftgdlocalrating("trname", new()
    ///     {
    ///         Rating = "1",
    ///         Status = "enable",
    ///         Url = "sgala.com",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Webfilter FtgdLocalRating can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:webfilter/ftgdlocalrating:Ftgdlocalrating labelname {{url}}
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:webfilter/ftgdlocalrating:Ftgdlocalrating labelname {{url}}
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:webfilter/ftgdlocalrating:Ftgdlocalrating")]
    public partial class Ftgdlocalrating : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Comment.
        /// </summary>
        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        /// <summary>
        /// Local rating.
        /// </summary>
        [Output("rating")]
        public Output<string> Rating { get; private set; } = null!;

        /// <summary>
        /// Enable/disable local rating. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// URL to rate locally.
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Ftgdlocalrating resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Ftgdlocalrating(string name, FtgdlocalratingArgs args, CustomResourceOptions? options = null)
            : base("fortios:webfilter/ftgdlocalrating:Ftgdlocalrating", name, args ?? new FtgdlocalratingArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Ftgdlocalrating(string name, Input<string> id, FtgdlocalratingState? state = null, CustomResourceOptions? options = null)
            : base("fortios:webfilter/ftgdlocalrating:Ftgdlocalrating", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Ftgdlocalrating resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Ftgdlocalrating Get(string name, Input<string> id, FtgdlocalratingState? state = null, CustomResourceOptions? options = null)
        {
            return new Ftgdlocalrating(name, id, state, options);
        }
    }

    public sealed class FtgdlocalratingArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Local rating.
        /// </summary>
        [Input("rating", required: true)]
        public Input<string> Rating { get; set; } = null!;

        /// <summary>
        /// Enable/disable local rating. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// URL to rate locally.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public FtgdlocalratingArgs()
        {
        }
        public static new FtgdlocalratingArgs Empty => new FtgdlocalratingArgs();
    }

    public sealed class FtgdlocalratingState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Local rating.
        /// </summary>
        [Input("rating")]
        public Input<string>? Rating { get; set; }

        /// <summary>
        /// Enable/disable local rating. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// URL to rate locally.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public FtgdlocalratingState()
        {
        }
        public static new FtgdlocalratingState Empty => new FtgdlocalratingState();
    }
}