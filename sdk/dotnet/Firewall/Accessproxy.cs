// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Firewall
{
    /// <summary>
    /// Configure IPv4 access proxy. Applies to FortiOS Version `&gt;= 7.0.1`.
    /// 
    /// ## Import
    /// 
    /// Firewall AccessProxy can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:firewall/accessproxy:Accessproxy labelname {{name}}
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:firewall/accessproxy:Accessproxy labelname {{name}}
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:firewall/accessproxy:Accessproxy")]
    public partial class Accessproxy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Enable/disable adding vhost/domain to dnsdb for ztna dox tunnel. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("addVhostDomainToDnsdb")]
        public Output<string> AddVhostDomainToDnsdb { get; private set; } = null!;

        /// <summary>
        /// Set IPv6 API Gateway. The structure of `api_gateway6` block is documented below.
        /// </summary>
        [Output("apiGateway6s")]
        public Output<ImmutableArray<Outputs.AccessproxyApiGateway6>> ApiGateway6s { get; private set; } = null!;

        /// <summary>
        /// Set IPv4 API Gateway. The structure of `api_gateway` block is documented below.
        /// </summary>
        [Output("apiGateways")]
        public Output<ImmutableArray<Outputs.AccessproxyApiGateway>> ApiGateways { get; private set; } = null!;

        /// <summary>
        /// Enable/disable authentication portal. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("authPortal")]
        public Output<string> AuthPortal { get; private set; } = null!;

        /// <summary>
        /// Virtual host for authentication portal.
        /// </summary>
        [Output("authVirtualHost")]
        public Output<string> AuthVirtualHost { get; private set; } = null!;

        /// <summary>
        /// Enable/disable to request client certificate. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("clientCert")]
        public Output<string> ClientCert { get; private set; } = null!;

        /// <summary>
        /// Decrypted traffic mirror.
        /// </summary>
        [Output("decryptedTrafficMirror")]
        public Output<string> DecryptedTrafficMirror { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// Action of an empty client certificate.
        /// </summary>
        [Output("emptyCertAction")]
        public Output<string> EmptyCertAction { get; private set; } = null!;

        /// <summary>
        /// Enable/disable logging of blocked traffic. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("logBlockedTraffic")]
        public Output<string> LogBlockedTraffic { get; private set; } = null!;

        /// <summary>
        /// Access Proxy name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Enable/disable to detect device type by HTTP user-agent if no client certificate provided. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("userAgentDetect")]
        public Output<string> UserAgentDetect { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;

        /// <summary>
        /// Virtual IP name.
        /// </summary>
        [Output("vip")]
        public Output<string> Vip { get; private set; } = null!;


        /// <summary>
        /// Create a Accessproxy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Accessproxy(string name, AccessproxyArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:firewall/accessproxy:Accessproxy", name, args ?? new AccessproxyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Accessproxy(string name, Input<string> id, AccessproxyState? state = null, CustomResourceOptions? options = null)
            : base("fortios:firewall/accessproxy:Accessproxy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Accessproxy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Accessproxy Get(string name, Input<string> id, AccessproxyState? state = null, CustomResourceOptions? options = null)
        {
            return new Accessproxy(name, id, state, options);
        }
    }

    public sealed class AccessproxyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable adding vhost/domain to dnsdb for ztna dox tunnel. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("addVhostDomainToDnsdb")]
        public Input<string>? AddVhostDomainToDnsdb { get; set; }

        [Input("apiGateway6s")]
        private InputList<Inputs.AccessproxyApiGateway6Args>? _apiGateway6s;

        /// <summary>
        /// Set IPv6 API Gateway. The structure of `api_gateway6` block is documented below.
        /// </summary>
        public InputList<Inputs.AccessproxyApiGateway6Args> ApiGateway6s
        {
            get => _apiGateway6s ?? (_apiGateway6s = new InputList<Inputs.AccessproxyApiGateway6Args>());
            set => _apiGateway6s = value;
        }

        [Input("apiGateways")]
        private InputList<Inputs.AccessproxyApiGatewayArgs>? _apiGateways;

        /// <summary>
        /// Set IPv4 API Gateway. The structure of `api_gateway` block is documented below.
        /// </summary>
        public InputList<Inputs.AccessproxyApiGatewayArgs> ApiGateways
        {
            get => _apiGateways ?? (_apiGateways = new InputList<Inputs.AccessproxyApiGatewayArgs>());
            set => _apiGateways = value;
        }

        /// <summary>
        /// Enable/disable authentication portal. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("authPortal")]
        public Input<string>? AuthPortal { get; set; }

        /// <summary>
        /// Virtual host for authentication portal.
        /// </summary>
        [Input("authVirtualHost")]
        public Input<string>? AuthVirtualHost { get; set; }

        /// <summary>
        /// Enable/disable to request client certificate. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("clientCert")]
        public Input<string>? ClientCert { get; set; }

        /// <summary>
        /// Decrypted traffic mirror.
        /// </summary>
        [Input("decryptedTrafficMirror")]
        public Input<string>? DecryptedTrafficMirror { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Action of an empty client certificate.
        /// </summary>
        [Input("emptyCertAction")]
        public Input<string>? EmptyCertAction { get; set; }

        /// <summary>
        /// Enable/disable logging of blocked traffic. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("logBlockedTraffic")]
        public Input<string>? LogBlockedTraffic { get; set; }

        /// <summary>
        /// Access Proxy name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Enable/disable to detect device type by HTTP user-agent if no client certificate provided. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("userAgentDetect")]
        public Input<string>? UserAgentDetect { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        /// <summary>
        /// Virtual IP name.
        /// </summary>
        [Input("vip")]
        public Input<string>? Vip { get; set; }

        public AccessproxyArgs()
        {
        }
        public static new AccessproxyArgs Empty => new AccessproxyArgs();
    }

    public sealed class AccessproxyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable adding vhost/domain to dnsdb for ztna dox tunnel. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("addVhostDomainToDnsdb")]
        public Input<string>? AddVhostDomainToDnsdb { get; set; }

        [Input("apiGateway6s")]
        private InputList<Inputs.AccessproxyApiGateway6GetArgs>? _apiGateway6s;

        /// <summary>
        /// Set IPv6 API Gateway. The structure of `api_gateway6` block is documented below.
        /// </summary>
        public InputList<Inputs.AccessproxyApiGateway6GetArgs> ApiGateway6s
        {
            get => _apiGateway6s ?? (_apiGateway6s = new InputList<Inputs.AccessproxyApiGateway6GetArgs>());
            set => _apiGateway6s = value;
        }

        [Input("apiGateways")]
        private InputList<Inputs.AccessproxyApiGatewayGetArgs>? _apiGateways;

        /// <summary>
        /// Set IPv4 API Gateway. The structure of `api_gateway` block is documented below.
        /// </summary>
        public InputList<Inputs.AccessproxyApiGatewayGetArgs> ApiGateways
        {
            get => _apiGateways ?? (_apiGateways = new InputList<Inputs.AccessproxyApiGatewayGetArgs>());
            set => _apiGateways = value;
        }

        /// <summary>
        /// Enable/disable authentication portal. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("authPortal")]
        public Input<string>? AuthPortal { get; set; }

        /// <summary>
        /// Virtual host for authentication portal.
        /// </summary>
        [Input("authVirtualHost")]
        public Input<string>? AuthVirtualHost { get; set; }

        /// <summary>
        /// Enable/disable to request client certificate. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("clientCert")]
        public Input<string>? ClientCert { get; set; }

        /// <summary>
        /// Decrypted traffic mirror.
        /// </summary>
        [Input("decryptedTrafficMirror")]
        public Input<string>? DecryptedTrafficMirror { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Action of an empty client certificate.
        /// </summary>
        [Input("emptyCertAction")]
        public Input<string>? EmptyCertAction { get; set; }

        /// <summary>
        /// Enable/disable logging of blocked traffic. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("logBlockedTraffic")]
        public Input<string>? LogBlockedTraffic { get; set; }

        /// <summary>
        /// Access Proxy name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Enable/disable to detect device type by HTTP user-agent if no client certificate provided. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("userAgentDetect")]
        public Input<string>? UserAgentDetect { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        /// <summary>
        /// Virtual IP name.
        /// </summary>
        [Input("vip")]
        public Input<string>? Vip { get; set; }

        public AccessproxyState()
        {
        }
        public static new AccessproxyState Empty => new AccessproxyState();
    }
}
