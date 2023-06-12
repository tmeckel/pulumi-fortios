// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.System3gmodem
{
    /// <summary>
    /// 3G MODEM custom. Applies to FortiOS Version `7.0.4`.
    /// 
    /// ## Import
    /// 
    /// System3GModem Custom can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:system3gmodem/custom:Custom labelname {{fosid}}
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:system3gmodem/custom:Custom labelname {{fosid}}
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:system3gmodem/custom:Custom")]
    public partial class Custom : global::Pulumi.CustomResource
    {
        /// <summary>
        /// USB interface class in hexadecimal format (00-ff).
        /// </summary>
        [Output("classId")]
        public Output<string> ClassId { get; private set; } = null!;

        /// <summary>
        /// ID.
        /// </summary>
        [Output("fosid")]
        public Output<int> Fosid { get; private set; } = null!;

        /// <summary>
        /// Init string in hexadecimal format (even length).
        /// </summary>
        [Output("initString")]
        public Output<string> InitString { get; private set; } = null!;

        /// <summary>
        /// MODEM model name.
        /// </summary>
        [Output("model")]
        public Output<string> Model { get; private set; } = null!;

        /// <summary>
        /// USB modeswitch arguments. e.g: '-v 1410 -p 9030 -V 1410 -P 9032 -u 3'
        /// </summary>
        [Output("modeswitchString")]
        public Output<string> ModeswitchString { get; private set; } = null!;

        /// <summary>
        /// USB product ID in hexadecimal format (0000-ffff).
        /// </summary>
        [Output("productId")]
        public Output<string> ProductId { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;

        /// <summary>
        /// MODEM vendor name.
        /// </summary>
        [Output("vendor")]
        public Output<string> Vendor { get; private set; } = null!;

        /// <summary>
        /// USB vendor ID in hexadecimal format (0000-ffff).
        /// </summary>
        [Output("vendorId")]
        public Output<string> VendorId { get; private set; } = null!;


        /// <summary>
        /// Create a Custom resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Custom(string name, CustomArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:system3gmodem/custom:Custom", name, args ?? new CustomArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Custom(string name, Input<string> id, CustomState? state = null, CustomResourceOptions? options = null)
            : base("fortios:system3gmodem/custom:Custom", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Custom resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Custom Get(string name, Input<string> id, CustomState? state = null, CustomResourceOptions? options = null)
        {
            return new Custom(name, id, state, options);
        }
    }

    public sealed class CustomArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// USB interface class in hexadecimal format (00-ff).
        /// </summary>
        [Input("classId")]
        public Input<string>? ClassId { get; set; }

        /// <summary>
        /// ID.
        /// </summary>
        [Input("fosid")]
        public Input<int>? Fosid { get; set; }

        /// <summary>
        /// Init string in hexadecimal format (even length).
        /// </summary>
        [Input("initString")]
        public Input<string>? InitString { get; set; }

        /// <summary>
        /// MODEM model name.
        /// </summary>
        [Input("model")]
        public Input<string>? Model { get; set; }

        /// <summary>
        /// USB modeswitch arguments. e.g: '-v 1410 -p 9030 -V 1410 -P 9032 -u 3'
        /// </summary>
        [Input("modeswitchString")]
        public Input<string>? ModeswitchString { get; set; }

        /// <summary>
        /// USB product ID in hexadecimal format (0000-ffff).
        /// </summary>
        [Input("productId")]
        public Input<string>? ProductId { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        /// <summary>
        /// MODEM vendor name.
        /// </summary>
        [Input("vendor")]
        public Input<string>? Vendor { get; set; }

        /// <summary>
        /// USB vendor ID in hexadecimal format (0000-ffff).
        /// </summary>
        [Input("vendorId")]
        public Input<string>? VendorId { get; set; }

        public CustomArgs()
        {
        }
        public static new CustomArgs Empty => new CustomArgs();
    }

    public sealed class CustomState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// USB interface class in hexadecimal format (00-ff).
        /// </summary>
        [Input("classId")]
        public Input<string>? ClassId { get; set; }

        /// <summary>
        /// ID.
        /// </summary>
        [Input("fosid")]
        public Input<int>? Fosid { get; set; }

        /// <summary>
        /// Init string in hexadecimal format (even length).
        /// </summary>
        [Input("initString")]
        public Input<string>? InitString { get; set; }

        /// <summary>
        /// MODEM model name.
        /// </summary>
        [Input("model")]
        public Input<string>? Model { get; set; }

        /// <summary>
        /// USB modeswitch arguments. e.g: '-v 1410 -p 9030 -V 1410 -P 9032 -u 3'
        /// </summary>
        [Input("modeswitchString")]
        public Input<string>? ModeswitchString { get; set; }

        /// <summary>
        /// USB product ID in hexadecimal format (0000-ffff).
        /// </summary>
        [Input("productId")]
        public Input<string>? ProductId { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        /// <summary>
        /// MODEM vendor name.
        /// </summary>
        [Input("vendor")]
        public Input<string>? Vendor { get; set; }

        /// <summary>
        /// USB vendor ID in hexadecimal format (0000-ffff).
        /// </summary>
        [Input("vendorId")]
        public Input<string>? VendorId { get; set; }

        public CustomState()
        {
        }
        public static new CustomState Empty => new CustomState();
    }
}
