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
    /// Configure FIPS-CC mode.
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
    ///     var trname = new Fortios.Sys.Fipscc("trname", new()
    ///     {
    ///         EntropyToken = "enable",
    ///         KeyGenerationSelfTest = "disable",
    ///         SelfTestPeriod = 1440,
    ///         Status = "disable",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// System FipsCc can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:sys/fipscc:Fipscc labelname SystemFipsCc
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:sys/fipscc:Fipscc labelname SystemFipsCc
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:sys/fipscc:Fipscc")]
    public partial class Fipscc : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Enable/disable/dynamic entropy token. Valid values: `enable`, `disable`, `dynamic`.
        /// </summary>
        [Output("entropyToken")]
        public Output<string> EntropyToken { get; private set; } = null!;

        /// <summary>
        /// Enable/disable self tests after key generation. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("keyGenerationSelfTest")]
        public Output<string> KeyGenerationSelfTest { get; private set; } = null!;

        /// <summary>
        /// Self test period.
        /// </summary>
        [Output("selfTestPeriod")]
        public Output<int> SelfTestPeriod { get; private set; } = null!;

        /// <summary>
        /// Enable/disable FIPS-CC mode. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Fipscc resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Fipscc(string name, FipsccArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:sys/fipscc:Fipscc", name, args ?? new FipsccArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Fipscc(string name, Input<string> id, FipsccState? state = null, CustomResourceOptions? options = null)
            : base("fortios:sys/fipscc:Fipscc", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Fipscc resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Fipscc Get(string name, Input<string> id, FipsccState? state = null, CustomResourceOptions? options = null)
        {
            return new Fipscc(name, id, state, options);
        }
    }

    public sealed class FipsccArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable/dynamic entropy token. Valid values: `enable`, `disable`, `dynamic`.
        /// </summary>
        [Input("entropyToken")]
        public Input<string>? EntropyToken { get; set; }

        /// <summary>
        /// Enable/disable self tests after key generation. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("keyGenerationSelfTest")]
        public Input<string>? KeyGenerationSelfTest { get; set; }

        /// <summary>
        /// Self test period.
        /// </summary>
        [Input("selfTestPeriod")]
        public Input<int>? SelfTestPeriod { get; set; }

        /// <summary>
        /// Enable/disable FIPS-CC mode. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public FipsccArgs()
        {
        }
        public static new FipsccArgs Empty => new FipsccArgs();
    }

    public sealed class FipsccState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable/dynamic entropy token. Valid values: `enable`, `disable`, `dynamic`.
        /// </summary>
        [Input("entropyToken")]
        public Input<string>? EntropyToken { get; set; }

        /// <summary>
        /// Enable/disable self tests after key generation. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("keyGenerationSelfTest")]
        public Input<string>? KeyGenerationSelfTest { get; set; }

        /// <summary>
        /// Self test period.
        /// </summary>
        [Input("selfTestPeriod")]
        public Input<int>? SelfTestPeriod { get; set; }

        /// <summary>
        /// Enable/disable FIPS-CC mode. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public FipsccState()
        {
        }
        public static new FipsccState Empty => new FipsccState();
    }
}