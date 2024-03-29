// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Systemsnmp
{
    /// <summary>
    /// SNMP user configuration.
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
    ///     var trname = new Fortios.Systemsnmp.User("trname", new()
    ///     {
    ///         AuthProto = "sha",
    ///         Events = "cpu-high mem-low log-full intf-ip vpn-tun-up vpn-tun-down ha-switch ha-hb-failure ips-signature ips-anomaly av-virus av-oversize av-pattern av-fragmented fm-if-change bgp-established bgp-backward-transition ha-member-up ha-member-down ent-conf-change av-conserve av-bypass av-oversize-passed av-oversize-blocked ips-pkg-update ips-fail-open faz-disconnect wc-ap-up wc-ap-down fswctl-session-up fswctl-session-down load-balance-real-server-down per-cpu-high",
    ///         HaDirect = "disable",
    ///         PrivProto = "aes",
    ///         Queries = "disable",
    ///         QueryPort = 161,
    ///         SecurityLevel = "no-auth-no-priv",
    ///         SourceIp = "0.0.0.0",
    ///         SourceIpv6 = "::",
    ///         Status = "disable",
    ///         TrapLport = 162,
    ///         TrapRport = 162,
    ///         TrapStatus = "enable",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// SystemSnmp User can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:systemsnmp/user:User labelname {{name}}
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:systemsnmp/user:User labelname {{name}}
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:systemsnmp/user:User")]
    public partial class User : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Authentication protocol.
        /// </summary>
        [Output("authProto")]
        public Output<string> AuthProto { get; private set; } = null!;

        /// <summary>
        /// Password for authentication protocol.
        /// </summary>
        [Output("authPwd")]
        public Output<string?> AuthPwd { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// SNMP notifications (traps) to send.
        /// </summary>
        [Output("events")]
        public Output<string> Events { get; private set; } = null!;

        /// <summary>
        /// Enable/disable direct management of HA cluster members. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("haDirect")]
        public Output<string> HaDirect { get; private set; } = null!;

        /// <summary>
        /// SNMP access control MIB view.
        /// </summary>
        [Output("mibView")]
        public Output<string> MibView { get; private set; } = null!;

        /// <summary>
        /// SNMP user name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// SNMP managers to send notifications (traps) to.
        /// </summary>
        [Output("notifyHosts")]
        public Output<string> NotifyHosts { get; private set; } = null!;

        /// <summary>
        /// IPv6 SNMP managers to send notifications (traps) to.
        /// </summary>
        [Output("notifyHosts6")]
        public Output<string> NotifyHosts6 { get; private set; } = null!;

        /// <summary>
        /// Privacy (encryption) protocol. Valid values: `aes`, `des`, `aes256`, `aes256cisco`.
        /// </summary>
        [Output("privProto")]
        public Output<string> PrivProto { get; private set; } = null!;

        /// <summary>
        /// Password for privacy (encryption) protocol.
        /// </summary>
        [Output("privPwd")]
        public Output<string?> PrivPwd { get; private set; } = null!;

        /// <summary>
        /// Enable/disable SNMP queries for this user. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("queries")]
        public Output<string> Queries { get; private set; } = null!;

        /// <summary>
        /// SNMPv3 query port (default = 161).
        /// </summary>
        [Output("queryPort")]
        public Output<int> QueryPort { get; private set; } = null!;

        /// <summary>
        /// Security level for message authentication and encryption. Valid values: `no-auth-no-priv`, `auth-no-priv`, `auth-priv`.
        /// </summary>
        [Output("securityLevel")]
        public Output<string> SecurityLevel { get; private set; } = null!;

        /// <summary>
        /// Source IP for SNMP trap.
        /// </summary>
        [Output("sourceIp")]
        public Output<string> SourceIp { get; private set; } = null!;

        /// <summary>
        /// Source IPv6 for SNMP trap.
        /// </summary>
        [Output("sourceIpv6")]
        public Output<string> SourceIpv6 { get; private set; } = null!;

        /// <summary>
        /// Enable/disable this SNMP user. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// SNMPv3 local trap port (default = 162).
        /// </summary>
        [Output("trapLport")]
        public Output<int> TrapLport { get; private set; } = null!;

        /// <summary>
        /// SNMPv3 trap remote port (default = 162).
        /// </summary>
        [Output("trapRport")]
        public Output<int> TrapRport { get; private set; } = null!;

        /// <summary>
        /// Enable/disable traps for this SNMP user. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("trapStatus")]
        public Output<string> TrapStatus { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;

        /// <summary>
        /// SNMP access control VDOMs. The structure of `vdoms` block is documented below.
        /// </summary>
        [Output("vdoms")]
        public Output<ImmutableArray<Outputs.UserVdom>> Vdoms { get; private set; } = null!;


        /// <summary>
        /// Create a User resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public User(string name, UserArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:systemsnmp/user:User", name, args ?? new UserArgs(), MakeResourceOptions(options, ""))
        {
        }

        private User(string name, Input<string> id, UserState? state = null, CustomResourceOptions? options = null)
            : base("fortios:systemsnmp/user:User", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse/pulumi-fortios",
                AdditionalSecretOutputs =
                {
                    "authPwd",
                    "privPwd",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing User resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static User Get(string name, Input<string> id, UserState? state = null, CustomResourceOptions? options = null)
        {
            return new User(name, id, state, options);
        }
    }

    public sealed class UserArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Authentication protocol.
        /// </summary>
        [Input("authProto")]
        public Input<string>? AuthProto { get; set; }

        [Input("authPwd")]
        private Input<string>? _authPwd;

        /// <summary>
        /// Password for authentication protocol.
        /// </summary>
        public Input<string>? AuthPwd
        {
            get => _authPwd;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _authPwd = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// SNMP notifications (traps) to send.
        /// </summary>
        [Input("events")]
        public Input<string>? Events { get; set; }

        /// <summary>
        /// Enable/disable direct management of HA cluster members. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("haDirect")]
        public Input<string>? HaDirect { get; set; }

        /// <summary>
        /// SNMP access control MIB view.
        /// </summary>
        [Input("mibView")]
        public Input<string>? MibView { get; set; }

        /// <summary>
        /// SNMP user name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// SNMP managers to send notifications (traps) to.
        /// </summary>
        [Input("notifyHosts")]
        public Input<string>? NotifyHosts { get; set; }

        /// <summary>
        /// IPv6 SNMP managers to send notifications (traps) to.
        /// </summary>
        [Input("notifyHosts6")]
        public Input<string>? NotifyHosts6 { get; set; }

        /// <summary>
        /// Privacy (encryption) protocol. Valid values: `aes`, `des`, `aes256`, `aes256cisco`.
        /// </summary>
        [Input("privProto")]
        public Input<string>? PrivProto { get; set; }

        [Input("privPwd")]
        private Input<string>? _privPwd;

        /// <summary>
        /// Password for privacy (encryption) protocol.
        /// </summary>
        public Input<string>? PrivPwd
        {
            get => _privPwd;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _privPwd = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Enable/disable SNMP queries for this user. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("queries")]
        public Input<string>? Queries { get; set; }

        /// <summary>
        /// SNMPv3 query port (default = 161).
        /// </summary>
        [Input("queryPort")]
        public Input<int>? QueryPort { get; set; }

        /// <summary>
        /// Security level for message authentication and encryption. Valid values: `no-auth-no-priv`, `auth-no-priv`, `auth-priv`.
        /// </summary>
        [Input("securityLevel")]
        public Input<string>? SecurityLevel { get; set; }

        /// <summary>
        /// Source IP for SNMP trap.
        /// </summary>
        [Input("sourceIp")]
        public Input<string>? SourceIp { get; set; }

        /// <summary>
        /// Source IPv6 for SNMP trap.
        /// </summary>
        [Input("sourceIpv6")]
        public Input<string>? SourceIpv6 { get; set; }

        /// <summary>
        /// Enable/disable this SNMP user. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// SNMPv3 local trap port (default = 162).
        /// </summary>
        [Input("trapLport")]
        public Input<int>? TrapLport { get; set; }

        /// <summary>
        /// SNMPv3 trap remote port (default = 162).
        /// </summary>
        [Input("trapRport")]
        public Input<int>? TrapRport { get; set; }

        /// <summary>
        /// Enable/disable traps for this SNMP user. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("trapStatus")]
        public Input<string>? TrapStatus { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        [Input("vdoms")]
        private InputList<Inputs.UserVdomArgs>? _vdoms;

        /// <summary>
        /// SNMP access control VDOMs. The structure of `vdoms` block is documented below.
        /// </summary>
        public InputList<Inputs.UserVdomArgs> Vdoms
        {
            get => _vdoms ?? (_vdoms = new InputList<Inputs.UserVdomArgs>());
            set => _vdoms = value;
        }

        public UserArgs()
        {
        }
        public static new UserArgs Empty => new UserArgs();
    }

    public sealed class UserState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Authentication protocol.
        /// </summary>
        [Input("authProto")]
        public Input<string>? AuthProto { get; set; }

        [Input("authPwd")]
        private Input<string>? _authPwd;

        /// <summary>
        /// Password for authentication protocol.
        /// </summary>
        public Input<string>? AuthPwd
        {
            get => _authPwd;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _authPwd = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// SNMP notifications (traps) to send.
        /// </summary>
        [Input("events")]
        public Input<string>? Events { get; set; }

        /// <summary>
        /// Enable/disable direct management of HA cluster members. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("haDirect")]
        public Input<string>? HaDirect { get; set; }

        /// <summary>
        /// SNMP access control MIB view.
        /// </summary>
        [Input("mibView")]
        public Input<string>? MibView { get; set; }

        /// <summary>
        /// SNMP user name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// SNMP managers to send notifications (traps) to.
        /// </summary>
        [Input("notifyHosts")]
        public Input<string>? NotifyHosts { get; set; }

        /// <summary>
        /// IPv6 SNMP managers to send notifications (traps) to.
        /// </summary>
        [Input("notifyHosts6")]
        public Input<string>? NotifyHosts6 { get; set; }

        /// <summary>
        /// Privacy (encryption) protocol. Valid values: `aes`, `des`, `aes256`, `aes256cisco`.
        /// </summary>
        [Input("privProto")]
        public Input<string>? PrivProto { get; set; }

        [Input("privPwd")]
        private Input<string>? _privPwd;

        /// <summary>
        /// Password for privacy (encryption) protocol.
        /// </summary>
        public Input<string>? PrivPwd
        {
            get => _privPwd;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _privPwd = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Enable/disable SNMP queries for this user. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("queries")]
        public Input<string>? Queries { get; set; }

        /// <summary>
        /// SNMPv3 query port (default = 161).
        /// </summary>
        [Input("queryPort")]
        public Input<int>? QueryPort { get; set; }

        /// <summary>
        /// Security level for message authentication and encryption. Valid values: `no-auth-no-priv`, `auth-no-priv`, `auth-priv`.
        /// </summary>
        [Input("securityLevel")]
        public Input<string>? SecurityLevel { get; set; }

        /// <summary>
        /// Source IP for SNMP trap.
        /// </summary>
        [Input("sourceIp")]
        public Input<string>? SourceIp { get; set; }

        /// <summary>
        /// Source IPv6 for SNMP trap.
        /// </summary>
        [Input("sourceIpv6")]
        public Input<string>? SourceIpv6 { get; set; }

        /// <summary>
        /// Enable/disable this SNMP user. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// SNMPv3 local trap port (default = 162).
        /// </summary>
        [Input("trapLport")]
        public Input<int>? TrapLport { get; set; }

        /// <summary>
        /// SNMPv3 trap remote port (default = 162).
        /// </summary>
        [Input("trapRport")]
        public Input<int>? TrapRport { get; set; }

        /// <summary>
        /// Enable/disable traps for this SNMP user. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("trapStatus")]
        public Input<string>? TrapStatus { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        [Input("vdoms")]
        private InputList<Inputs.UserVdomGetArgs>? _vdoms;

        /// <summary>
        /// SNMP access control VDOMs. The structure of `vdoms` block is documented below.
        /// </summary>
        public InputList<Inputs.UserVdomGetArgs> Vdoms
        {
            get => _vdoms ?? (_vdoms = new InputList<Inputs.UserVdomGetArgs>());
            set => _vdoms = value;
        }

        public UserState()
        {
        }
        public static new UserState Empty => new UserState();
    }
}
