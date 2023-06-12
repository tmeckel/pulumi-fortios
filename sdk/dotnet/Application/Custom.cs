// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Application
{
    /// <summary>
    /// Configure custom application signatures.
    /// 
    /// ## Import
    /// 
    /// Application Custom can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:application/custom:Custom labelname {{tag}}
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:application/custom:Custom labelname {{tag}}
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:application/custom:Custom")]
    public partial class Custom : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Custom application signature behavior.
        /// </summary>
        [Output("behavior")]
        public Output<string> Behavior { get; private set; } = null!;

        /// <summary>
        /// Custom application category ID (use ? to view available options).
        /// </summary>
        [Output("category")]
        public Output<int> Category { get; private set; } = null!;

        /// <summary>
        /// Comment.
        /// </summary>
        [Output("comment")]
        public Output<string> Comment { get; private set; } = null!;

        /// <summary>
        /// Custom application category ID (use ? to view available options).
        /// </summary>
        [Output("fosid")]
        public Output<int> Fosid { get; private set; } = null!;

        /// <summary>
        /// Name of this custom application signature.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Custom application signature protocol.
        /// </summary>
        [Output("protocol")]
        public Output<string> Protocol { get; private set; } = null!;

        /// <summary>
        /// The text that makes up the actual custom application signature.
        /// </summary>
        [Output("signature")]
        public Output<string> Signature { get; private set; } = null!;

        /// <summary>
        /// Signature tag.
        /// </summary>
        [Output("tag")]
        public Output<string> Tag { get; private set; } = null!;

        /// <summary>
        /// Custom application signature technology.
        /// </summary>
        [Output("technology")]
        public Output<string> Technology { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;

        /// <summary>
        /// Custom application signature vendor.
        /// </summary>
        [Output("vendor")]
        public Output<string> Vendor { get; private set; } = null!;


        /// <summary>
        /// Create a Custom resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Custom(string name, CustomArgs args, CustomResourceOptions? options = null)
            : base("fortios:application/custom:Custom", name, args ?? new CustomArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Custom(string name, Input<string> id, CustomState? state = null, CustomResourceOptions? options = null)
            : base("fortios:application/custom:Custom", name, state, MakeResourceOptions(options, id))
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
        /// Custom application signature behavior.
        /// </summary>
        [Input("behavior")]
        public Input<string>? Behavior { get; set; }

        /// <summary>
        /// Custom application category ID (use ? to view available options).
        /// </summary>
        [Input("category", required: true)]
        public Input<int> Category { get; set; } = null!;

        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Custom application category ID (use ? to view available options).
        /// </summary>
        [Input("fosid")]
        public Input<int>? Fosid { get; set; }

        /// <summary>
        /// Name of this custom application signature.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Custom application signature protocol.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// The text that makes up the actual custom application signature.
        /// </summary>
        [Input("signature")]
        public Input<string>? Signature { get; set; }

        /// <summary>
        /// Signature tag.
        /// </summary>
        [Input("tag")]
        public Input<string>? Tag { get; set; }

        /// <summary>
        /// Custom application signature technology.
        /// </summary>
        [Input("technology")]
        public Input<string>? Technology { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        /// <summary>
        /// Custom application signature vendor.
        /// </summary>
        [Input("vendor")]
        public Input<string>? Vendor { get; set; }

        public CustomArgs()
        {
        }
        public static new CustomArgs Empty => new CustomArgs();
    }

    public sealed class CustomState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Custom application signature behavior.
        /// </summary>
        [Input("behavior")]
        public Input<string>? Behavior { get; set; }

        /// <summary>
        /// Custom application category ID (use ? to view available options).
        /// </summary>
        [Input("category")]
        public Input<int>? Category { get; set; }

        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Custom application category ID (use ? to view available options).
        /// </summary>
        [Input("fosid")]
        public Input<int>? Fosid { get; set; }

        /// <summary>
        /// Name of this custom application signature.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Custom application signature protocol.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// The text that makes up the actual custom application signature.
        /// </summary>
        [Input("signature")]
        public Input<string>? Signature { get; set; }

        /// <summary>
        /// Signature tag.
        /// </summary>
        [Input("tag")]
        public Input<string>? Tag { get; set; }

        /// <summary>
        /// Custom application signature technology.
        /// </summary>
        [Input("technology")]
        public Input<string>? Technology { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        /// <summary>
        /// Custom application signature vendor.
        /// </summary>
        [Input("vendor")]
        public Input<string>? Vendor { get; set; }

        public CustomState()
        {
        }
        public static new CustomState Empty => new CustomState();
    }
}