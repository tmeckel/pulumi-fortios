// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Extensioncontroller
{
    /// <summary>
    /// FortiExtender dataplan configuration. Applies to FortiOS Version `&gt;= 7.2.1`.
    /// 
    /// ## Import
    /// 
    /// ExtensionController Dataplan can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:extensioncontroller/dataplan:Dataplan labelname {{name}}
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:extensioncontroller/dataplan:Dataplan labelname {{name}}
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:extensioncontroller/dataplan:Dataplan")]
    public partial class Dataplan : global::Pulumi.CustomResource
    {
        /// <summary>
        /// APN configuration.
        /// </summary>
        [Output("apn")]
        public Output<string> Apn { get; private set; } = null!;

        /// <summary>
        /// Authentication type. Valid values: `none`, `pap`, `chap`.
        /// </summary>
        [Output("authType")]
        public Output<string> AuthType { get; private set; } = null!;

        /// <summary>
        /// Billing day of the month (1 - 31).
        /// </summary>
        [Output("billingDate")]
        public Output<int> BillingDate { get; private set; } = null!;

        /// <summary>
        /// Capacity in MB (0 - 102400000).
        /// </summary>
        [Output("capacity")]
        public Output<int> Capacity { get; private set; } = null!;

        /// <summary>
        /// Carrier configuration.
        /// </summary>
        [Output("carrier")]
        public Output<string> Carrier { get; private set; } = null!;

        /// <summary>
        /// ICCID configuration.
        /// </summary>
        [Output("iccid")]
        public Output<string> Iccid { get; private set; } = null!;

        /// <summary>
        /// Dataplan's modem specifics, if any. Valid values: `modem1`, `modem2`, `all`.
        /// </summary>
        [Output("modemId")]
        public Output<string> ModemId { get; private set; } = null!;

        /// <summary>
        /// Monthly fee of dataplan (0 - 100000, in local currency).
        /// </summary>
        [Output("monthlyFee")]
        public Output<int> MonthlyFee { get; private set; } = null!;

        /// <summary>
        /// FortiExtender data plan name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Enable/disable dataplan overage detection. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("overage")]
        public Output<string> Overage { get; private set; } = null!;

        /// <summary>
        /// Password.
        /// </summary>
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        /// <summary>
        /// PDN type. Valid values: `ipv4-only`, `ipv6-only`, `ipv4-ipv6`.
        /// </summary>
        [Output("pdn")]
        public Output<string> Pdn { get; private set; } = null!;

        /// <summary>
        /// Preferred subnet mask (0 - 32).
        /// </summary>
        [Output("preferredSubnet")]
        public Output<int> PreferredSubnet { get; private set; } = null!;

        /// <summary>
        /// Enable/disable dataplan private network support. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("privateNetwork")]
        public Output<string> PrivateNetwork { get; private set; } = null!;

        /// <summary>
        /// Signal period (600 to 18000 seconds).
        /// </summary>
        [Output("signalPeriod")]
        public Output<int> SignalPeriod { get; private set; } = null!;

        /// <summary>
        /// Signal threshold. Specify the range between 50 - 100, where 50/100 means -50/-100 dBm.
        /// </summary>
        [Output("signalThreshold")]
        public Output<int> SignalThreshold { get; private set; } = null!;

        /// <summary>
        /// SIM slot configuration. Valid values: `sim1`, `sim2`.
        /// </summary>
        [Output("slot")]
        public Output<string> Slot { get; private set; } = null!;

        /// <summary>
        /// Type preferences configuration. Valid values: `carrier`, `slot`, `iccid`, `generic`.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Username.
        /// </summary>
        [Output("username")]
        public Output<string> Username { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Dataplan resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Dataplan(string name, DataplanArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:extensioncontroller/dataplan:Dataplan", name, args ?? new DataplanArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Dataplan(string name, Input<string> id, DataplanState? state = null, CustomResourceOptions? options = null)
            : base("fortios:extensioncontroller/dataplan:Dataplan", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Dataplan resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Dataplan Get(string name, Input<string> id, DataplanState? state = null, CustomResourceOptions? options = null)
        {
            return new Dataplan(name, id, state, options);
        }
    }

    public sealed class DataplanArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// APN configuration.
        /// </summary>
        [Input("apn")]
        public Input<string>? Apn { get; set; }

        /// <summary>
        /// Authentication type. Valid values: `none`, `pap`, `chap`.
        /// </summary>
        [Input("authType")]
        public Input<string>? AuthType { get; set; }

        /// <summary>
        /// Billing day of the month (1 - 31).
        /// </summary>
        [Input("billingDate")]
        public Input<int>? BillingDate { get; set; }

        /// <summary>
        /// Capacity in MB (0 - 102400000).
        /// </summary>
        [Input("capacity")]
        public Input<int>? Capacity { get; set; }

        /// <summary>
        /// Carrier configuration.
        /// </summary>
        [Input("carrier")]
        public Input<string>? Carrier { get; set; }

        /// <summary>
        /// ICCID configuration.
        /// </summary>
        [Input("iccid")]
        public Input<string>? Iccid { get; set; }

        /// <summary>
        /// Dataplan's modem specifics, if any. Valid values: `modem1`, `modem2`, `all`.
        /// </summary>
        [Input("modemId")]
        public Input<string>? ModemId { get; set; }

        /// <summary>
        /// Monthly fee of dataplan (0 - 100000, in local currency).
        /// </summary>
        [Input("monthlyFee")]
        public Input<int>? MonthlyFee { get; set; }

        /// <summary>
        /// FortiExtender data plan name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Enable/disable dataplan overage detection. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("overage")]
        public Input<string>? Overage { get; set; }

        /// <summary>
        /// Password.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// PDN type. Valid values: `ipv4-only`, `ipv6-only`, `ipv4-ipv6`.
        /// </summary>
        [Input("pdn")]
        public Input<string>? Pdn { get; set; }

        /// <summary>
        /// Preferred subnet mask (0 - 32).
        /// </summary>
        [Input("preferredSubnet")]
        public Input<int>? PreferredSubnet { get; set; }

        /// <summary>
        /// Enable/disable dataplan private network support. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("privateNetwork")]
        public Input<string>? PrivateNetwork { get; set; }

        /// <summary>
        /// Signal period (600 to 18000 seconds).
        /// </summary>
        [Input("signalPeriod")]
        public Input<int>? SignalPeriod { get; set; }

        /// <summary>
        /// Signal threshold. Specify the range between 50 - 100, where 50/100 means -50/-100 dBm.
        /// </summary>
        [Input("signalThreshold")]
        public Input<int>? SignalThreshold { get; set; }

        /// <summary>
        /// SIM slot configuration. Valid values: `sim1`, `sim2`.
        /// </summary>
        [Input("slot")]
        public Input<string>? Slot { get; set; }

        /// <summary>
        /// Type preferences configuration. Valid values: `carrier`, `slot`, `iccid`, `generic`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Username.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public DataplanArgs()
        {
        }
        public static new DataplanArgs Empty => new DataplanArgs();
    }

    public sealed class DataplanState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// APN configuration.
        /// </summary>
        [Input("apn")]
        public Input<string>? Apn { get; set; }

        /// <summary>
        /// Authentication type. Valid values: `none`, `pap`, `chap`.
        /// </summary>
        [Input("authType")]
        public Input<string>? AuthType { get; set; }

        /// <summary>
        /// Billing day of the month (1 - 31).
        /// </summary>
        [Input("billingDate")]
        public Input<int>? BillingDate { get; set; }

        /// <summary>
        /// Capacity in MB (0 - 102400000).
        /// </summary>
        [Input("capacity")]
        public Input<int>? Capacity { get; set; }

        /// <summary>
        /// Carrier configuration.
        /// </summary>
        [Input("carrier")]
        public Input<string>? Carrier { get; set; }

        /// <summary>
        /// ICCID configuration.
        /// </summary>
        [Input("iccid")]
        public Input<string>? Iccid { get; set; }

        /// <summary>
        /// Dataplan's modem specifics, if any. Valid values: `modem1`, `modem2`, `all`.
        /// </summary>
        [Input("modemId")]
        public Input<string>? ModemId { get; set; }

        /// <summary>
        /// Monthly fee of dataplan (0 - 100000, in local currency).
        /// </summary>
        [Input("monthlyFee")]
        public Input<int>? MonthlyFee { get; set; }

        /// <summary>
        /// FortiExtender data plan name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Enable/disable dataplan overage detection. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("overage")]
        public Input<string>? Overage { get; set; }

        /// <summary>
        /// Password.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// PDN type. Valid values: `ipv4-only`, `ipv6-only`, `ipv4-ipv6`.
        /// </summary>
        [Input("pdn")]
        public Input<string>? Pdn { get; set; }

        /// <summary>
        /// Preferred subnet mask (0 - 32).
        /// </summary>
        [Input("preferredSubnet")]
        public Input<int>? PreferredSubnet { get; set; }

        /// <summary>
        /// Enable/disable dataplan private network support. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("privateNetwork")]
        public Input<string>? PrivateNetwork { get; set; }

        /// <summary>
        /// Signal period (600 to 18000 seconds).
        /// </summary>
        [Input("signalPeriod")]
        public Input<int>? SignalPeriod { get; set; }

        /// <summary>
        /// Signal threshold. Specify the range between 50 - 100, where 50/100 means -50/-100 dBm.
        /// </summary>
        [Input("signalThreshold")]
        public Input<int>? SignalThreshold { get; set; }

        /// <summary>
        /// SIM slot configuration. Valid values: `sim1`, `sim2`.
        /// </summary>
        [Input("slot")]
        public Input<string>? Slot { get; set; }

        /// <summary>
        /// Type preferences configuration. Valid values: `carrier`, `slot`, `iccid`, `generic`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Username.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public DataplanState()
        {
        }
        public static new DataplanState Empty => new DataplanState();
    }
}
