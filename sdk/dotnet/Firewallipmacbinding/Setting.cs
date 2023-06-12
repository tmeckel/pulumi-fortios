// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Firewallipmacbinding
{
    /// <summary>
    /// Configure IP to MAC binding settings.
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
    ///     var trname = new Fortios.Firewallipmacbinding.Setting("trname", new()
    ///     {
    ///         Bindthroughfw = "disable",
    ///         Bindtofw = "disable",
    ///         Undefinedhost = "block",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// FirewallIpmacbinding Setting can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:firewallipmacbinding/setting:Setting labelname FirewallIpmacbindingSetting
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:firewallipmacbinding/setting:Setting labelname FirewallIpmacbindingSetting
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:firewallipmacbinding/setting:Setting")]
    public partial class Setting : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Enable/disable use of IP/MAC binding to filter packets that would normally go through the firewall. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("bindthroughfw")]
        public Output<string> Bindthroughfw { get; private set; } = null!;

        /// <summary>
        /// Enable/disable use of IP/MAC binding to filter packets that would normally go to the firewall. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("bindtofw")]
        public Output<string> Bindtofw { get; private set; } = null!;

        /// <summary>
        /// Select action to take on packets with IP/MAC addresses not in the binding list (default = block). Valid values: `allow`, `block`.
        /// </summary>
        [Output("undefinedhost")]
        public Output<string> Undefinedhost { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Setting resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Setting(string name, SettingArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:firewallipmacbinding/setting:Setting", name, args ?? new SettingArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Setting(string name, Input<string> id, SettingState? state = null, CustomResourceOptions? options = null)
            : base("fortios:firewallipmacbinding/setting:Setting", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Setting resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Setting Get(string name, Input<string> id, SettingState? state = null, CustomResourceOptions? options = null)
        {
            return new Setting(name, id, state, options);
        }
    }

    public sealed class SettingArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable use of IP/MAC binding to filter packets that would normally go through the firewall. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("bindthroughfw")]
        public Input<string>? Bindthroughfw { get; set; }

        /// <summary>
        /// Enable/disable use of IP/MAC binding to filter packets that would normally go to the firewall. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("bindtofw")]
        public Input<string>? Bindtofw { get; set; }

        /// <summary>
        /// Select action to take on packets with IP/MAC addresses not in the binding list (default = block). Valid values: `allow`, `block`.
        /// </summary>
        [Input("undefinedhost")]
        public Input<string>? Undefinedhost { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SettingArgs()
        {
        }
        public static new SettingArgs Empty => new SettingArgs();
    }

    public sealed class SettingState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable use of IP/MAC binding to filter packets that would normally go through the firewall. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("bindthroughfw")]
        public Input<string>? Bindthroughfw { get; set; }

        /// <summary>
        /// Enable/disable use of IP/MAC binding to filter packets that would normally go to the firewall. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("bindtofw")]
        public Input<string>? Bindtofw { get; set; }

        /// <summary>
        /// Select action to take on packets with IP/MAC addresses not in the binding list (default = block). Valid values: `allow`, `block`.
        /// </summary>
        [Input("undefinedhost")]
        public Input<string>? Undefinedhost { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SettingState()
        {
        }
        public static new SettingState Empty => new SettingState();
    }
}
