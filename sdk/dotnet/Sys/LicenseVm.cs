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
    /// Provides a resource to update VM license using uploaded file for FortiOS. Reboots immediately if successful.
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
    ///     var test2 = new Fortios.Sys.LicenseVm("test2", new()
    ///     {
    ///         FileContent = "LS0tLS1CRUdJTiBGR1QgVk0gTElDRU5TRS0tLS0tDQpRQUFBQUxXaTdCVnVkV2x3QXJZcC92S2J2Yk5zME5YNWluUW9sVldmcFoxWldJQi9pL2g4c01oR0psWWc5Vkl1DQorSlBJRis1aFphMWwyNm9yNHdiEQE3RnJDeVZnQUFBQWhxWjliWHFLK1hGN2o3dnB3WTB6QXRTaTdOMVM1ZWNxDQpWYmRRREZyYklUdnRvUWNyRU1jV0ltQzFqWWs5dmVoeGlYTG1OV0MwN25BSitYTTJFNmh2b29DMjE1YUwxK2wrDQovUHl5M0VLVnNTNjJDT2hMZHc3UndXajB3V3RqMmZiWg0KLS0tLS1FTkQgRkdUIFZNIExJQ0VOU0UtLS0tLQ0K",
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [FortiosResourceType("fortios:sys/licenseVm:LicenseVm")]
    public partial class LicenseVm : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The license file, it needs to be base64 encoded, must not contain whitespace or other invalid base64 characters, and must be included in HTTP body.
        /// </summary>
        [Output("fileContent")]
        public Output<string> FileContent { get; private set; } = null!;


        /// <summary>
        /// Create a LicenseVm resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LicenseVm(string name, LicenseVmArgs args, CustomResourceOptions? options = null)
            : base("fortios:sys/licenseVm:LicenseVm", name, args ?? new LicenseVmArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LicenseVm(string name, Input<string> id, LicenseVmState? state = null, CustomResourceOptions? options = null)
            : base("fortios:sys/licenseVm:LicenseVm", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing LicenseVm resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LicenseVm Get(string name, Input<string> id, LicenseVmState? state = null, CustomResourceOptions? options = null)
        {
            return new LicenseVm(name, id, state, options);
        }
    }

    public sealed class LicenseVmArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The license file, it needs to be base64 encoded, must not contain whitespace or other invalid base64 characters, and must be included in HTTP body.
        /// </summary>
        [Input("fileContent", required: true)]
        public Input<string> FileContent { get; set; } = null!;

        public LicenseVmArgs()
        {
        }
        public static new LicenseVmArgs Empty => new LicenseVmArgs();
    }

    public sealed class LicenseVmState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The license file, it needs to be base64 encoded, must not contain whitespace or other invalid base64 characters, and must be included in HTTP body.
        /// </summary>
        [Input("fileContent")]
        public Input<string>? FileContent { get; set; }

        public LicenseVmState()
        {
        }
        public static new LicenseVmState Empty => new LicenseVmState();
    }
}
