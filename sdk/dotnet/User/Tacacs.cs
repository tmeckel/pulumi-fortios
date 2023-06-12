// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.User
{
    /// <summary>
    /// Configure TACACS+ server entries.
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
    ///     var trname = new Fortios.User.Tacacs("trname", new()
    ///     {
    ///         AuthenType = "auto",
    ///         Authorization = "disable",
    ///         Port = 2342,
    ///         Server = "1.1.1.1",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// User Tacacs can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:user/tacacs:Tacacs labelname {{name}}
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:user/tacacs:Tacacs labelname {{name}}
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:user/tacacs:Tacacs")]
    public partial class Tacacs : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Allowed authentication protocols/methods. Valid values: `mschap`, `chap`, `pap`, `ascii`, `auto`.
        /// </summary>
        [Output("authenType")]
        public Output<string> AuthenType { get; private set; } = null!;

        /// <summary>
        /// Enable/disable TACACS+ authorization. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("authorization")]
        public Output<string> Authorization { get; private set; } = null!;

        /// <summary>
        /// Specify outgoing interface to reach server.
        /// </summary>
        [Output("interface")]
        public Output<string> Interface { get; private set; } = null!;

        /// <summary>
        /// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
        /// </summary>
        [Output("interfaceSelectMethod")]
        public Output<string> InterfaceSelectMethod { get; private set; } = null!;

        /// <summary>
        /// Key to access the primary server.
        /// </summary>
        [Output("key")]
        public Output<string?> Key { get; private set; } = null!;

        /// <summary>
        /// TACACS+ server entry name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Port number of the TACACS+ server.
        /// </summary>
        [Output("port")]
        public Output<int> Port { get; private set; } = null!;

        /// <summary>
        /// Key to access the secondary server.
        /// </summary>
        [Output("secondaryKey")]
        public Output<string?> SecondaryKey { get; private set; } = null!;

        /// <summary>
        /// Secondary TACACS+ server CN domain name or IP address.
        /// </summary>
        [Output("secondaryServer")]
        public Output<string> SecondaryServer { get; private set; } = null!;

        /// <summary>
        /// Primary TACACS+ server CN domain name or IP address.
        /// </summary>
        [Output("server")]
        public Output<string> Server { get; private set; } = null!;

        /// <summary>
        /// source IP for communications to TACACS+ server.
        /// </summary>
        [Output("sourceIp")]
        public Output<string> SourceIp { get; private set; } = null!;

        /// <summary>
        /// Key to access the tertiary server.
        /// </summary>
        [Output("tertiaryKey")]
        public Output<string?> TertiaryKey { get; private set; } = null!;

        /// <summary>
        /// Tertiary TACACS+ server CN domain name or IP address.
        /// </summary>
        [Output("tertiaryServer")]
        public Output<string> TertiaryServer { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Tacacs resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Tacacs(string name, TacacsArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:user/tacacs:Tacacs", name, args ?? new TacacsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Tacacs(string name, Input<string> id, TacacsState? state = null, CustomResourceOptions? options = null)
            : base("fortios:user/tacacs:Tacacs", name, state, MakeResourceOptions(options, id))
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
                    "key",
                    "secondaryKey",
                    "tertiaryKey",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Tacacs resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Tacacs Get(string name, Input<string> id, TacacsState? state = null, CustomResourceOptions? options = null)
        {
            return new Tacacs(name, id, state, options);
        }
    }

    public sealed class TacacsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Allowed authentication protocols/methods. Valid values: `mschap`, `chap`, `pap`, `ascii`, `auto`.
        /// </summary>
        [Input("authenType")]
        public Input<string>? AuthenType { get; set; }

        /// <summary>
        /// Enable/disable TACACS+ authorization. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("authorization")]
        public Input<string>? Authorization { get; set; }

        /// <summary>
        /// Specify outgoing interface to reach server.
        /// </summary>
        [Input("interface")]
        public Input<string>? Interface { get; set; }

        /// <summary>
        /// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
        /// </summary>
        [Input("interfaceSelectMethod")]
        public Input<string>? InterfaceSelectMethod { get; set; }

        [Input("key")]
        private Input<string>? _key;

        /// <summary>
        /// Key to access the primary server.
        /// </summary>
        public Input<string>? Key
        {
            get => _key;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _key = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// TACACS+ server entry name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Port number of the TACACS+ server.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        [Input("secondaryKey")]
        private Input<string>? _secondaryKey;

        /// <summary>
        /// Key to access the secondary server.
        /// </summary>
        public Input<string>? SecondaryKey
        {
            get => _secondaryKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _secondaryKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Secondary TACACS+ server CN domain name or IP address.
        /// </summary>
        [Input("secondaryServer")]
        public Input<string>? SecondaryServer { get; set; }

        /// <summary>
        /// Primary TACACS+ server CN domain name or IP address.
        /// </summary>
        [Input("server")]
        public Input<string>? Server { get; set; }

        /// <summary>
        /// source IP for communications to TACACS+ server.
        /// </summary>
        [Input("sourceIp")]
        public Input<string>? SourceIp { get; set; }

        [Input("tertiaryKey")]
        private Input<string>? _tertiaryKey;

        /// <summary>
        /// Key to access the tertiary server.
        /// </summary>
        public Input<string>? TertiaryKey
        {
            get => _tertiaryKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _tertiaryKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Tertiary TACACS+ server CN domain name or IP address.
        /// </summary>
        [Input("tertiaryServer")]
        public Input<string>? TertiaryServer { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public TacacsArgs()
        {
        }
        public static new TacacsArgs Empty => new TacacsArgs();
    }

    public sealed class TacacsState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Allowed authentication protocols/methods. Valid values: `mschap`, `chap`, `pap`, `ascii`, `auto`.
        /// </summary>
        [Input("authenType")]
        public Input<string>? AuthenType { get; set; }

        /// <summary>
        /// Enable/disable TACACS+ authorization. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("authorization")]
        public Input<string>? Authorization { get; set; }

        /// <summary>
        /// Specify outgoing interface to reach server.
        /// </summary>
        [Input("interface")]
        public Input<string>? Interface { get; set; }

        /// <summary>
        /// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
        /// </summary>
        [Input("interfaceSelectMethod")]
        public Input<string>? InterfaceSelectMethod { get; set; }

        [Input("key")]
        private Input<string>? _key;

        /// <summary>
        /// Key to access the primary server.
        /// </summary>
        public Input<string>? Key
        {
            get => _key;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _key = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// TACACS+ server entry name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Port number of the TACACS+ server.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        [Input("secondaryKey")]
        private Input<string>? _secondaryKey;

        /// <summary>
        /// Key to access the secondary server.
        /// </summary>
        public Input<string>? SecondaryKey
        {
            get => _secondaryKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _secondaryKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Secondary TACACS+ server CN domain name or IP address.
        /// </summary>
        [Input("secondaryServer")]
        public Input<string>? SecondaryServer { get; set; }

        /// <summary>
        /// Primary TACACS+ server CN domain name or IP address.
        /// </summary>
        [Input("server")]
        public Input<string>? Server { get; set; }

        /// <summary>
        /// source IP for communications to TACACS+ server.
        /// </summary>
        [Input("sourceIp")]
        public Input<string>? SourceIp { get; set; }

        [Input("tertiaryKey")]
        private Input<string>? _tertiaryKey;

        /// <summary>
        /// Key to access the tertiary server.
        /// </summary>
        public Input<string>? TertiaryKey
        {
            get => _tertiaryKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _tertiaryKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Tertiary TACACS+ server CN domain name or IP address.
        /// </summary>
        [Input("tertiaryServer")]
        public Input<string>? TertiaryServer { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public TacacsState()
        {
        }
        public static new TacacsState Empty => new TacacsState();
    }
}