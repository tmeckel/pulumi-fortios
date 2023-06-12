// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Firewallssh
{
    /// <summary>
    /// SSH proxy local CA.
    /// 
    /// ## Import
    /// 
    /// FirewallSsh LocalCa can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:firewallssh/localca:Localca labelname {{name}}
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:firewallssh/localca:Localca labelname {{name}}
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:firewallssh/localca:Localca")]
    public partial class Localca : global::Pulumi.CustomResource
    {
        /// <summary>
        /// SSH proxy local CA name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Password for SSH private key.
        /// </summary>
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        /// <summary>
        /// SSH proxy private key, encrypted with a password.
        /// </summary>
        [Output("privateKey")]
        public Output<string> PrivateKey { get; private set; } = null!;

        /// <summary>
        /// SSH proxy public key.
        /// </summary>
        [Output("publicKey")]
        public Output<string> PublicKey { get; private set; } = null!;

        /// <summary>
        /// SSH proxy local CA source type. Valid values: `built-in`, `user`.
        /// </summary>
        [Output("source")]
        public Output<string> Source { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Localca resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Localca(string name, LocalcaArgs args, CustomResourceOptions? options = null)
            : base("fortios:firewallssh/localca:Localca", name, args ?? new LocalcaArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Localca(string name, Input<string> id, LocalcaState? state = null, CustomResourceOptions? options = null)
            : base("fortios:firewallssh/localca:Localca", name, state, MakeResourceOptions(options, id))
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
                    "password",
                    "privateKey",
                    "publicKey",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Localca resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Localca Get(string name, Input<string> id, LocalcaState? state = null, CustomResourceOptions? options = null)
        {
            return new Localca(name, id, state, options);
        }
    }

    public sealed class LocalcaArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// SSH proxy local CA name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// Password for SSH private key.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("privateKey", required: true)]
        private Input<string>? _privateKey;

        /// <summary>
        /// SSH proxy private key, encrypted with a password.
        /// </summary>
        public Input<string>? PrivateKey
        {
            get => _privateKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _privateKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("publicKey", required: true)]
        private Input<string>? _publicKey;

        /// <summary>
        /// SSH proxy public key.
        /// </summary>
        public Input<string>? PublicKey
        {
            get => _publicKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _publicKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// SSH proxy local CA source type. Valid values: `built-in`, `user`.
        /// </summary>
        [Input("source")]
        public Input<string>? Source { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public LocalcaArgs()
        {
        }
        public static new LocalcaArgs Empty => new LocalcaArgs();
    }

    public sealed class LocalcaState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// SSH proxy local CA name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// Password for SSH private key.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("privateKey")]
        private Input<string>? _privateKey;

        /// <summary>
        /// SSH proxy private key, encrypted with a password.
        /// </summary>
        public Input<string>? PrivateKey
        {
            get => _privateKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _privateKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("publicKey")]
        private Input<string>? _publicKey;

        /// <summary>
        /// SSH proxy public key.
        /// </summary>
        public Input<string>? PublicKey
        {
            get => _publicKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _publicKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// SSH proxy local CA source type. Valid values: `built-in`, `user`.
        /// </summary>
        [Input("source")]
        public Input<string>? Source { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public LocalcaState()
        {
        }
        public static new LocalcaState Empty => new LocalcaState();
    }
}
