// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Log
{
    /// <summary>
    /// Configure custom log fields.
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
    ///     var trname = new Fortios.Log.Customfield("trname", new()
    ///     {
    ///         Fosid = "1",
    ///         Value = "logteststr",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Log CustomField can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:log/customfield:Customfield labelname {{fosid}}
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:log/customfield:Customfield labelname {{fosid}}
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:log/customfield:Customfield")]
    public partial class Customfield : global::Pulumi.CustomResource
    {
        /// <summary>
        /// field ID &lt;string&gt;.
        /// </summary>
        [Output("fosid")]
        public Output<string> Fosid { get; private set; } = null!;

        /// <summary>
        /// Field name (max: 15 characters).
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Field value (max: 15 characters).
        /// </summary>
        [Output("value")]
        public Output<string> Value { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Customfield resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Customfield(string name, CustomfieldArgs args, CustomResourceOptions? options = null)
            : base("fortios:log/customfield:Customfield", name, args ?? new CustomfieldArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Customfield(string name, Input<string> id, CustomfieldState? state = null, CustomResourceOptions? options = null)
            : base("fortios:log/customfield:Customfield", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Customfield resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Customfield Get(string name, Input<string> id, CustomfieldState? state = null, CustomResourceOptions? options = null)
        {
            return new Customfield(name, id, state, options);
        }
    }

    public sealed class CustomfieldArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// field ID &lt;string&gt;.
        /// </summary>
        [Input("fosid")]
        public Input<string>? Fosid { get; set; }

        /// <summary>
        /// Field name (max: 15 characters).
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Field value (max: 15 characters).
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public CustomfieldArgs()
        {
        }
        public static new CustomfieldArgs Empty => new CustomfieldArgs();
    }

    public sealed class CustomfieldState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// field ID &lt;string&gt;.
        /// </summary>
        [Input("fosid")]
        public Input<string>? Fosid { get; set; }

        /// <summary>
        /// Field name (max: 15 characters).
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Field value (max: 15 characters).
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public CustomfieldState()
        {
        }
        public static new CustomfieldState Empty => new CustomfieldState();
    }
}
