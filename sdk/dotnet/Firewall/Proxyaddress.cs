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
    /// Web proxy address configuration.
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
    ///     var trname = new Fortios.Firewall.Proxyaddress("trname", new()
    ///     {
    ///         CaseSensitivity = "disable",
    ///         Color = 2,
    ///         Referrer = "enable",
    ///         Type = "host-regex",
    ///         Visibility = "enable",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Firewall ProxyAddress can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:firewall/proxyaddress:Proxyaddress labelname {{name}}
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:firewall/proxyaddress:Proxyaddress labelname {{name}}
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:firewall/proxyaddress:Proxyaddress")]
    public partial class Proxyaddress : global::Pulumi.CustomResource
    {
        /// <summary>
        /// SaaS application. The structure of `application` block is documented below.
        /// </summary>
        [Output("applications")]
        public Output<ImmutableArray<Outputs.ProxyaddressApplication>> Applications { get; private set; } = null!;

        /// <summary>
        /// Enable to make the pattern case sensitive. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("caseSensitivity")]
        public Output<string> CaseSensitivity { get; private set; } = null!;

        /// <summary>
        /// FortiGuard category ID. The structure of `category` block is documented below.
        /// </summary>
        [Output("categories")]
        public Output<ImmutableArray<Outputs.ProxyaddressCategory>> Categories { get; private set; } = null!;

        /// <summary>
        /// Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets value to 1).
        /// </summary>
        [Output("color")]
        public Output<int> Color { get; private set; } = null!;

        /// <summary>
        /// Optional comments.
        /// </summary>
        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// HTTP header name as a regular expression.
        /// </summary>
        [Output("header")]
        public Output<string> Header { get; private set; } = null!;

        /// <summary>
        /// HTTP header group. The structure of `header_group` block is documented below.
        /// </summary>
        [Output("headerGroups")]
        public Output<ImmutableArray<Outputs.ProxyaddressHeaderGroup>> HeaderGroups { get; private set; } = null!;

        /// <summary>
        /// Name of HTTP header.
        /// </summary>
        [Output("headerName")]
        public Output<string> HeaderName { get; private set; } = null!;

        /// <summary>
        /// Address object for the host.
        /// </summary>
        [Output("host")]
        public Output<string> Host { get; private set; } = null!;

        /// <summary>
        /// Host name as a regular expression.
        /// </summary>
        [Output("hostRegex")]
        public Output<string> HostRegex { get; private set; } = null!;

        /// <summary>
        /// HTTP request methods to be used. Valid values: `get`, `post`, `put`, `head`, `connect`, `trace`, `options`, `delete`.
        /// </summary>
        [Output("method")]
        public Output<string> Method { get; private set; } = null!;

        /// <summary>
        /// Address name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// URL path as a regular expression.
        /// </summary>
        [Output("path")]
        public Output<string> Path { get; private set; } = null!;

        /// <summary>
        /// Match the query part of the URL as a regular expression.
        /// </summary>
        [Output("query")]
        public Output<string> Query { get; private set; } = null!;

        /// <summary>
        /// Enable/disable use of referrer field in the HTTP header to match the address. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("referrer")]
        public Output<string> Referrer { get; private set; } = null!;

        /// <summary>
        /// Config object tagging. The structure of `tagging` block is documented below.
        /// </summary>
        [Output("taggings")]
        public Output<ImmutableArray<Outputs.ProxyaddressTagging>> Taggings { get; private set; } = null!;

        /// <summary>
        /// Proxy address type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Names of browsers to be used as user agent. Valid values: `chrome`, `ms`, `firefox`, `safari`, `other`.
        /// </summary>
        [Output("ua")]
        public Output<string> Ua { get; private set; } = null!;

        /// <summary>
        /// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        /// </summary>
        [Output("uuid")]
        public Output<string> Uuid { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;

        /// <summary>
        /// Enable/disable visibility of the object in the GUI. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("visibility")]
        public Output<string> Visibility { get; private set; } = null!;


        /// <summary>
        /// Create a Proxyaddress resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Proxyaddress(string name, ProxyaddressArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:firewall/proxyaddress:Proxyaddress", name, args ?? new ProxyaddressArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Proxyaddress(string name, Input<string> id, ProxyaddressState? state = null, CustomResourceOptions? options = null)
            : base("fortios:firewall/proxyaddress:Proxyaddress", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Proxyaddress resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Proxyaddress Get(string name, Input<string> id, ProxyaddressState? state = null, CustomResourceOptions? options = null)
        {
            return new Proxyaddress(name, id, state, options);
        }
    }

    public sealed class ProxyaddressArgs : global::Pulumi.ResourceArgs
    {
        [Input("applications")]
        private InputList<Inputs.ProxyaddressApplicationArgs>? _applications;

        /// <summary>
        /// SaaS application. The structure of `application` block is documented below.
        /// </summary>
        public InputList<Inputs.ProxyaddressApplicationArgs> Applications
        {
            get => _applications ?? (_applications = new InputList<Inputs.ProxyaddressApplicationArgs>());
            set => _applications = value;
        }

        /// <summary>
        /// Enable to make the pattern case sensitive. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("caseSensitivity")]
        public Input<string>? CaseSensitivity { get; set; }

        [Input("categories")]
        private InputList<Inputs.ProxyaddressCategoryArgs>? _categories;

        /// <summary>
        /// FortiGuard category ID. The structure of `category` block is documented below.
        /// </summary>
        public InputList<Inputs.ProxyaddressCategoryArgs> Categories
        {
            get => _categories ?? (_categories = new InputList<Inputs.ProxyaddressCategoryArgs>());
            set => _categories = value;
        }

        /// <summary>
        /// Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets value to 1).
        /// </summary>
        [Input("color")]
        public Input<int>? Color { get; set; }

        /// <summary>
        /// Optional comments.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// HTTP header name as a regular expression.
        /// </summary>
        [Input("header")]
        public Input<string>? Header { get; set; }

        [Input("headerGroups")]
        private InputList<Inputs.ProxyaddressHeaderGroupArgs>? _headerGroups;

        /// <summary>
        /// HTTP header group. The structure of `header_group` block is documented below.
        /// </summary>
        public InputList<Inputs.ProxyaddressHeaderGroupArgs> HeaderGroups
        {
            get => _headerGroups ?? (_headerGroups = new InputList<Inputs.ProxyaddressHeaderGroupArgs>());
            set => _headerGroups = value;
        }

        /// <summary>
        /// Name of HTTP header.
        /// </summary>
        [Input("headerName")]
        public Input<string>? HeaderName { get; set; }

        /// <summary>
        /// Address object for the host.
        /// </summary>
        [Input("host")]
        public Input<string>? Host { get; set; }

        /// <summary>
        /// Host name as a regular expression.
        /// </summary>
        [Input("hostRegex")]
        public Input<string>? HostRegex { get; set; }

        /// <summary>
        /// HTTP request methods to be used. Valid values: `get`, `post`, `put`, `head`, `connect`, `trace`, `options`, `delete`.
        /// </summary>
        [Input("method")]
        public Input<string>? Method { get; set; }

        /// <summary>
        /// Address name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// URL path as a regular expression.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// Match the query part of the URL as a regular expression.
        /// </summary>
        [Input("query")]
        public Input<string>? Query { get; set; }

        /// <summary>
        /// Enable/disable use of referrer field in the HTTP header to match the address. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("referrer")]
        public Input<string>? Referrer { get; set; }

        [Input("taggings")]
        private InputList<Inputs.ProxyaddressTaggingArgs>? _taggings;

        /// <summary>
        /// Config object tagging. The structure of `tagging` block is documented below.
        /// </summary>
        public InputList<Inputs.ProxyaddressTaggingArgs> Taggings
        {
            get => _taggings ?? (_taggings = new InputList<Inputs.ProxyaddressTaggingArgs>());
            set => _taggings = value;
        }

        /// <summary>
        /// Proxy address type.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Names of browsers to be used as user agent. Valid values: `chrome`, `ms`, `firefox`, `safari`, `other`.
        /// </summary>
        [Input("ua")]
        public Input<string>? Ua { get; set; }

        /// <summary>
        /// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        /// </summary>
        [Input("uuid")]
        public Input<string>? Uuid { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        /// <summary>
        /// Enable/disable visibility of the object in the GUI. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("visibility")]
        public Input<string>? Visibility { get; set; }

        public ProxyaddressArgs()
        {
        }
        public static new ProxyaddressArgs Empty => new ProxyaddressArgs();
    }

    public sealed class ProxyaddressState : global::Pulumi.ResourceArgs
    {
        [Input("applications")]
        private InputList<Inputs.ProxyaddressApplicationGetArgs>? _applications;

        /// <summary>
        /// SaaS application. The structure of `application` block is documented below.
        /// </summary>
        public InputList<Inputs.ProxyaddressApplicationGetArgs> Applications
        {
            get => _applications ?? (_applications = new InputList<Inputs.ProxyaddressApplicationGetArgs>());
            set => _applications = value;
        }

        /// <summary>
        /// Enable to make the pattern case sensitive. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("caseSensitivity")]
        public Input<string>? CaseSensitivity { get; set; }

        [Input("categories")]
        private InputList<Inputs.ProxyaddressCategoryGetArgs>? _categories;

        /// <summary>
        /// FortiGuard category ID. The structure of `category` block is documented below.
        /// </summary>
        public InputList<Inputs.ProxyaddressCategoryGetArgs> Categories
        {
            get => _categories ?? (_categories = new InputList<Inputs.ProxyaddressCategoryGetArgs>());
            set => _categories = value;
        }

        /// <summary>
        /// Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets value to 1).
        /// </summary>
        [Input("color")]
        public Input<int>? Color { get; set; }

        /// <summary>
        /// Optional comments.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// HTTP header name as a regular expression.
        /// </summary>
        [Input("header")]
        public Input<string>? Header { get; set; }

        [Input("headerGroups")]
        private InputList<Inputs.ProxyaddressHeaderGroupGetArgs>? _headerGroups;

        /// <summary>
        /// HTTP header group. The structure of `header_group` block is documented below.
        /// </summary>
        public InputList<Inputs.ProxyaddressHeaderGroupGetArgs> HeaderGroups
        {
            get => _headerGroups ?? (_headerGroups = new InputList<Inputs.ProxyaddressHeaderGroupGetArgs>());
            set => _headerGroups = value;
        }

        /// <summary>
        /// Name of HTTP header.
        /// </summary>
        [Input("headerName")]
        public Input<string>? HeaderName { get; set; }

        /// <summary>
        /// Address object for the host.
        /// </summary>
        [Input("host")]
        public Input<string>? Host { get; set; }

        /// <summary>
        /// Host name as a regular expression.
        /// </summary>
        [Input("hostRegex")]
        public Input<string>? HostRegex { get; set; }

        /// <summary>
        /// HTTP request methods to be used. Valid values: `get`, `post`, `put`, `head`, `connect`, `trace`, `options`, `delete`.
        /// </summary>
        [Input("method")]
        public Input<string>? Method { get; set; }

        /// <summary>
        /// Address name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// URL path as a regular expression.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// Match the query part of the URL as a regular expression.
        /// </summary>
        [Input("query")]
        public Input<string>? Query { get; set; }

        /// <summary>
        /// Enable/disable use of referrer field in the HTTP header to match the address. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("referrer")]
        public Input<string>? Referrer { get; set; }

        [Input("taggings")]
        private InputList<Inputs.ProxyaddressTaggingGetArgs>? _taggings;

        /// <summary>
        /// Config object tagging. The structure of `tagging` block is documented below.
        /// </summary>
        public InputList<Inputs.ProxyaddressTaggingGetArgs> Taggings
        {
            get => _taggings ?? (_taggings = new InputList<Inputs.ProxyaddressTaggingGetArgs>());
            set => _taggings = value;
        }

        /// <summary>
        /// Proxy address type.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Names of browsers to be used as user agent. Valid values: `chrome`, `ms`, `firefox`, `safari`, `other`.
        /// </summary>
        [Input("ua")]
        public Input<string>? Ua { get; set; }

        /// <summary>
        /// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        /// </summary>
        [Input("uuid")]
        public Input<string>? Uuid { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        /// <summary>
        /// Enable/disable visibility of the object in the GUI. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("visibility")]
        public Input<string>? Visibility { get; set; }

        public ProxyaddressState()
        {
        }
        public static new ProxyaddressState Empty => new ProxyaddressState();
    }
}