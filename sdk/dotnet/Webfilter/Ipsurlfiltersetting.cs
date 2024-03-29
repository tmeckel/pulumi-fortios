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
    /// Configure IPS URL filter settings.
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
    ///     var trname = new Fortios.Webfilter.Ipsurlfiltersetting("trname", new()
    ///     {
    ///         Distance = 1,
    ///         Gateway = "0.0.0.0",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Webfilter IpsUrlfilterSetting can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:webfilter/ipsurlfiltersetting:Ipsurlfiltersetting labelname WebfilterIpsUrlfilterSetting
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:webfilter/ipsurlfiltersetting:Ipsurlfiltersetting labelname WebfilterIpsUrlfilterSetting
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:webfilter/ipsurlfiltersetting:Ipsurlfiltersetting")]
    public partial class Ipsurlfiltersetting : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Interface for this route.
        /// </summary>
        [Output("device")]
        public Output<string> Device { get; private set; } = null!;

        /// <summary>
        /// Administrative distance (1 - 255) for this route.
        /// </summary>
        [Output("distance")]
        public Output<int> Distance { get; private set; } = null!;

        /// <summary>
        /// Gateway IP address for this route.
        /// </summary>
        [Output("gateway")]
        public Output<string> Gateway { get; private set; } = null!;

        /// <summary>
        /// Filter based on geographical location. Route will NOT be installed if the resolved IP address belongs to the country in the filter.
        /// </summary>
        [Output("geoFilter")]
        public Output<string?> GeoFilter { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Ipsurlfiltersetting resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Ipsurlfiltersetting(string name, IpsurlfiltersettingArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:webfilter/ipsurlfiltersetting:Ipsurlfiltersetting", name, args ?? new IpsurlfiltersettingArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Ipsurlfiltersetting(string name, Input<string> id, IpsurlfiltersettingState? state = null, CustomResourceOptions? options = null)
            : base("fortios:webfilter/ipsurlfiltersetting:Ipsurlfiltersetting", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Ipsurlfiltersetting resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Ipsurlfiltersetting Get(string name, Input<string> id, IpsurlfiltersettingState? state = null, CustomResourceOptions? options = null)
        {
            return new Ipsurlfiltersetting(name, id, state, options);
        }
    }

    public sealed class IpsurlfiltersettingArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Interface for this route.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// Administrative distance (1 - 255) for this route.
        /// </summary>
        [Input("distance")]
        public Input<int>? Distance { get; set; }

        /// <summary>
        /// Gateway IP address for this route.
        /// </summary>
        [Input("gateway")]
        public Input<string>? Gateway { get; set; }

        /// <summary>
        /// Filter based on geographical location. Route will NOT be installed if the resolved IP address belongs to the country in the filter.
        /// </summary>
        [Input("geoFilter")]
        public Input<string>? GeoFilter { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public IpsurlfiltersettingArgs()
        {
        }
        public static new IpsurlfiltersettingArgs Empty => new IpsurlfiltersettingArgs();
    }

    public sealed class IpsurlfiltersettingState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Interface for this route.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// Administrative distance (1 - 255) for this route.
        /// </summary>
        [Input("distance")]
        public Input<int>? Distance { get; set; }

        /// <summary>
        /// Gateway IP address for this route.
        /// </summary>
        [Input("gateway")]
        public Input<string>? Gateway { get; set; }

        /// <summary>
        /// Filter based on geographical location. Route will NOT be installed if the resolved IP address belongs to the country in the filter.
        /// </summary>
        [Input("geoFilter")]
        public Input<string>? GeoFilter { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public IpsurlfiltersettingState()
        {
        }
        public static new IpsurlfiltersettingState Empty => new IpsurlfiltersettingState();
    }
}
