// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Switchcontrollersecuritypolicy
{
    /// <summary>
    /// Configure 802.1x MAC Authentication Bypass (MAB) policies.
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
    ///     var trname = new Fortios.Switchcontrollersecuritypolicy.Policy8021X("trname", new()
    ///     {
    ///         AuthFailVlan = "disable",
    ///         AuthFailVlanid = 0,
    ///         EapPassthru = "disable",
    ///         FramevidApply = "enable",
    ///         GuestAuthDelay = 30,
    ///         GuestVlan = "disable",
    ///         GuestVlanid = 100,
    ///         MacAuthBypass = "disable",
    ///         OpenAuth = "disable",
    ///         PolicyType = "802.1X",
    ///         RadiusTimeoutOverwrite = "disable",
    ///         SecurityMode = "802.1X",
    ///         UserGroups = new[]
    ///         {
    ///             new Fortios.Switchcontrollersecuritypolicy.Inputs.Policy8021XUserGroupArgs
    ///             {
    ///                 Name = "Guest-group",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// SwitchControllerSecurityPolicy 8021X can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:switchcontrollersecuritypolicy/policy8021X:Policy8021X labelname {{name}}
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:switchcontrollersecuritypolicy/policy8021X:Policy8021X labelname {{name}}
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:switchcontrollersecuritypolicy/policy8021X:Policy8021X")]
    public partial class Policy8021X : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Enable to allow limited access to clients that cannot authenticate. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("authFailVlan")]
        public Output<string> AuthFailVlan { get; private set; } = null!;

        /// <summary>
        /// VLAN ID on which authentication failed.
        /// </summary>
        [Output("authFailVlanId")]
        public Output<string> AuthFailVlanId { get; private set; } = null!;

        /// <summary>
        /// VLAN ID on which authentication failed.
        /// </summary>
        [Output("authFailVlanid")]
        public Output<int> AuthFailVlanid { get; private set; } = null!;

        /// <summary>
        /// Authentication server timeout period (3 - 15 sec, default = 3).
        /// </summary>
        [Output("authserverTimeoutPeriod")]
        public Output<int> AuthserverTimeoutPeriod { get; private set; } = null!;

        /// <summary>
        /// Enable/disable the authentication server timeout VLAN to allow limited access when RADIUS is unavailable.  Valid values: `disable`, `enable`.
        /// </summary>
        [Output("authserverTimeoutVlan")]
        public Output<string> AuthserverTimeoutVlan { get; private set; } = null!;

        /// <summary>
        /// Authentication server timeout VLAN name.
        /// </summary>
        [Output("authserverTimeoutVlanid")]
        public Output<string> AuthserverTimeoutVlanid { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// Enable/disable automatic inclusion of untagged VLANs. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("eapAutoUntaggedVlans")]
        public Output<string> EapAutoUntaggedVlans { get; private set; } = null!;

        /// <summary>
        /// Enable/disable EAP pass-through mode, allowing protocols (such as LLDP) to pass through ports for more flexible authentication. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("eapPassthru")]
        public Output<string> EapPassthru { get; private set; } = null!;

        /// <summary>
        /// Enable/disable the capability to apply the EAP/MAB frame VLAN to the port native VLAN. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("framevidApply")]
        public Output<string> FramevidApply { get; private set; } = null!;

        /// <summary>
        /// Guest authentication delay (1 - 900  sec, default = 30).
        /// </summary>
        [Output("guestAuthDelay")]
        public Output<int> GuestAuthDelay { get; private set; } = null!;

        /// <summary>
        /// Enable the guest VLAN feature to allow limited access to non-802.1X-compliant clients. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("guestVlan")]
        public Output<string> GuestVlan { get; private set; } = null!;

        /// <summary>
        /// Guest VLAN name.
        /// </summary>
        [Output("guestVlanId")]
        public Output<string> GuestVlanId { get; private set; } = null!;

        /// <summary>
        /// Guest VLAN ID.
        /// </summary>
        [Output("guestVlanid")]
        public Output<int> GuestVlanid { get; private set; } = null!;

        /// <summary>
        /// Enable/disable MAB for this policy. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("macAuthBypass")]
        public Output<string> MacAuthBypass { get; private set; } = null!;

        /// <summary>
        /// Policy name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Enable/disable open authentication for this policy. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("openAuth")]
        public Output<string> OpenAuth { get; private set; } = null!;

        /// <summary>
        /// Policy type. Valid values: `802.1X`.
        /// </summary>
        [Output("policyType")]
        public Output<string> PolicyType { get; private set; } = null!;

        /// <summary>
        /// Enable to override the global RADIUS session timeout. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("radiusTimeoutOverwrite")]
        public Output<string> RadiusTimeoutOverwrite { get; private set; } = null!;

        /// <summary>
        /// Port or MAC based 802.1X security mode. Valid values: `802.1X`, `802.1X-mac-based`.
        /// </summary>
        [Output("securityMode")]
        public Output<string> SecurityMode { get; private set; } = null!;

        /// <summary>
        /// Name of user-group to assign to this MAC Authentication Bypass (MAB) policy. The structure of `user_group` block is documented below.
        /// </summary>
        [Output("userGroups")]
        public Output<ImmutableArray<Outputs.Policy8021XUserGroup>> UserGroups { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Policy8021X resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Policy8021X(string name, Policy8021XArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:switchcontrollersecuritypolicy/policy8021X:Policy8021X", name, args ?? new Policy8021XArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Policy8021X(string name, Input<string> id, Policy8021XState? state = null, CustomResourceOptions? options = null)
            : base("fortios:switchcontrollersecuritypolicy/policy8021X:Policy8021X", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Policy8021X resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Policy8021X Get(string name, Input<string> id, Policy8021XState? state = null, CustomResourceOptions? options = null)
        {
            return new Policy8021X(name, id, state, options);
        }
    }

    public sealed class Policy8021XArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable to allow limited access to clients that cannot authenticate. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("authFailVlan")]
        public Input<string>? AuthFailVlan { get; set; }

        /// <summary>
        /// VLAN ID on which authentication failed.
        /// </summary>
        [Input("authFailVlanId")]
        public Input<string>? AuthFailVlanId { get; set; }

        /// <summary>
        /// VLAN ID on which authentication failed.
        /// </summary>
        [Input("authFailVlanid")]
        public Input<int>? AuthFailVlanid { get; set; }

        /// <summary>
        /// Authentication server timeout period (3 - 15 sec, default = 3).
        /// </summary>
        [Input("authserverTimeoutPeriod")]
        public Input<int>? AuthserverTimeoutPeriod { get; set; }

        /// <summary>
        /// Enable/disable the authentication server timeout VLAN to allow limited access when RADIUS is unavailable.  Valid values: `disable`, `enable`.
        /// </summary>
        [Input("authserverTimeoutVlan")]
        public Input<string>? AuthserverTimeoutVlan { get; set; }

        /// <summary>
        /// Authentication server timeout VLAN name.
        /// </summary>
        [Input("authserverTimeoutVlanid")]
        public Input<string>? AuthserverTimeoutVlanid { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Enable/disable automatic inclusion of untagged VLANs. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("eapAutoUntaggedVlans")]
        public Input<string>? EapAutoUntaggedVlans { get; set; }

        /// <summary>
        /// Enable/disable EAP pass-through mode, allowing protocols (such as LLDP) to pass through ports for more flexible authentication. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("eapPassthru")]
        public Input<string>? EapPassthru { get; set; }

        /// <summary>
        /// Enable/disable the capability to apply the EAP/MAB frame VLAN to the port native VLAN. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("framevidApply")]
        public Input<string>? FramevidApply { get; set; }

        /// <summary>
        /// Guest authentication delay (1 - 900  sec, default = 30).
        /// </summary>
        [Input("guestAuthDelay")]
        public Input<int>? GuestAuthDelay { get; set; }

        /// <summary>
        /// Enable the guest VLAN feature to allow limited access to non-802.1X-compliant clients. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("guestVlan")]
        public Input<string>? GuestVlan { get; set; }

        /// <summary>
        /// Guest VLAN name.
        /// </summary>
        [Input("guestVlanId")]
        public Input<string>? GuestVlanId { get; set; }

        /// <summary>
        /// Guest VLAN ID.
        /// </summary>
        [Input("guestVlanid")]
        public Input<int>? GuestVlanid { get; set; }

        /// <summary>
        /// Enable/disable MAB for this policy. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("macAuthBypass")]
        public Input<string>? MacAuthBypass { get; set; }

        /// <summary>
        /// Policy name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Enable/disable open authentication for this policy. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("openAuth")]
        public Input<string>? OpenAuth { get; set; }

        /// <summary>
        /// Policy type. Valid values: `802.1X`.
        /// </summary>
        [Input("policyType")]
        public Input<string>? PolicyType { get; set; }

        /// <summary>
        /// Enable to override the global RADIUS session timeout. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("radiusTimeoutOverwrite")]
        public Input<string>? RadiusTimeoutOverwrite { get; set; }

        /// <summary>
        /// Port or MAC based 802.1X security mode. Valid values: `802.1X`, `802.1X-mac-based`.
        /// </summary>
        [Input("securityMode")]
        public Input<string>? SecurityMode { get; set; }

        [Input("userGroups")]
        private InputList<Inputs.Policy8021XUserGroupArgs>? _userGroups;

        /// <summary>
        /// Name of user-group to assign to this MAC Authentication Bypass (MAB) policy. The structure of `user_group` block is documented below.
        /// </summary>
        public InputList<Inputs.Policy8021XUserGroupArgs> UserGroups
        {
            get => _userGroups ?? (_userGroups = new InputList<Inputs.Policy8021XUserGroupArgs>());
            set => _userGroups = value;
        }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public Policy8021XArgs()
        {
        }
        public static new Policy8021XArgs Empty => new Policy8021XArgs();
    }

    public sealed class Policy8021XState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable to allow limited access to clients that cannot authenticate. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("authFailVlan")]
        public Input<string>? AuthFailVlan { get; set; }

        /// <summary>
        /// VLAN ID on which authentication failed.
        /// </summary>
        [Input("authFailVlanId")]
        public Input<string>? AuthFailVlanId { get; set; }

        /// <summary>
        /// VLAN ID on which authentication failed.
        /// </summary>
        [Input("authFailVlanid")]
        public Input<int>? AuthFailVlanid { get; set; }

        /// <summary>
        /// Authentication server timeout period (3 - 15 sec, default = 3).
        /// </summary>
        [Input("authserverTimeoutPeriod")]
        public Input<int>? AuthserverTimeoutPeriod { get; set; }

        /// <summary>
        /// Enable/disable the authentication server timeout VLAN to allow limited access when RADIUS is unavailable.  Valid values: `disable`, `enable`.
        /// </summary>
        [Input("authserverTimeoutVlan")]
        public Input<string>? AuthserverTimeoutVlan { get; set; }

        /// <summary>
        /// Authentication server timeout VLAN name.
        /// </summary>
        [Input("authserverTimeoutVlanid")]
        public Input<string>? AuthserverTimeoutVlanid { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Enable/disable automatic inclusion of untagged VLANs. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("eapAutoUntaggedVlans")]
        public Input<string>? EapAutoUntaggedVlans { get; set; }

        /// <summary>
        /// Enable/disable EAP pass-through mode, allowing protocols (such as LLDP) to pass through ports for more flexible authentication. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("eapPassthru")]
        public Input<string>? EapPassthru { get; set; }

        /// <summary>
        /// Enable/disable the capability to apply the EAP/MAB frame VLAN to the port native VLAN. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("framevidApply")]
        public Input<string>? FramevidApply { get; set; }

        /// <summary>
        /// Guest authentication delay (1 - 900  sec, default = 30).
        /// </summary>
        [Input("guestAuthDelay")]
        public Input<int>? GuestAuthDelay { get; set; }

        /// <summary>
        /// Enable the guest VLAN feature to allow limited access to non-802.1X-compliant clients. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("guestVlan")]
        public Input<string>? GuestVlan { get; set; }

        /// <summary>
        /// Guest VLAN name.
        /// </summary>
        [Input("guestVlanId")]
        public Input<string>? GuestVlanId { get; set; }

        /// <summary>
        /// Guest VLAN ID.
        /// </summary>
        [Input("guestVlanid")]
        public Input<int>? GuestVlanid { get; set; }

        /// <summary>
        /// Enable/disable MAB for this policy. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("macAuthBypass")]
        public Input<string>? MacAuthBypass { get; set; }

        /// <summary>
        /// Policy name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Enable/disable open authentication for this policy. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("openAuth")]
        public Input<string>? OpenAuth { get; set; }

        /// <summary>
        /// Policy type. Valid values: `802.1X`.
        /// </summary>
        [Input("policyType")]
        public Input<string>? PolicyType { get; set; }

        /// <summary>
        /// Enable to override the global RADIUS session timeout. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("radiusTimeoutOverwrite")]
        public Input<string>? RadiusTimeoutOverwrite { get; set; }

        /// <summary>
        /// Port or MAC based 802.1X security mode. Valid values: `802.1X`, `802.1X-mac-based`.
        /// </summary>
        [Input("securityMode")]
        public Input<string>? SecurityMode { get; set; }

        [Input("userGroups")]
        private InputList<Inputs.Policy8021XUserGroupGetArgs>? _userGroups;

        /// <summary>
        /// Name of user-group to assign to this MAC Authentication Bypass (MAB) policy. The structure of `user_group` block is documented below.
        /// </summary>
        public InputList<Inputs.Policy8021XUserGroupGetArgs> UserGroups
        {
            get => _userGroups ?? (_userGroups = new InputList<Inputs.Policy8021XUserGroupGetArgs>());
            set => _userGroups = value;
        }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public Policy8021XState()
        {
        }
        public static new Policy8021XState Empty => new Policy8021XState();
    }
}
