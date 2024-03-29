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
    public static class GetProxyaddress
    {
        /// <summary>
        /// Use this data source to get information on an fortios firewall proxyaddress
        /// </summary>
        public static Task<GetProxyaddressResult> InvokeAsync(GetProxyaddressArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetProxyaddressResult>("fortios:firewall/getProxyaddress:getProxyaddress", args ?? new GetProxyaddressArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on an fortios firewall proxyaddress
        /// </summary>
        public static Output<GetProxyaddressResult> Invoke(GetProxyaddressInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetProxyaddressResult>("fortios:firewall/getProxyaddress:getProxyaddress", args ?? new GetProxyaddressInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetProxyaddressArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired firewall proxyaddress.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetProxyaddressArgs()
        {
        }
        public static new GetProxyaddressArgs Empty => new GetProxyaddressArgs();
    }

    public sealed class GetProxyaddressInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired firewall proxyaddress.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetProxyaddressInvokeArgs()
        {
        }
        public static new GetProxyaddressInvokeArgs Empty => new GetProxyaddressInvokeArgs();
    }


    [OutputType]
    public sealed class GetProxyaddressResult
    {
        /// <summary>
        /// SaaS application. The structure of `application` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetProxyaddressApplicationResult> Applications;
        /// <summary>
        /// Case sensitivity in pattern.
        /// </summary>
        public readonly string CaseSensitivity;
        /// <summary>
        /// Tag category.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetProxyaddressCategoryResult> Categories;
        /// <summary>
        /// Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets value to 1).
        /// </summary>
        public readonly int Color;
        /// <summary>
        /// Optional comments.
        /// </summary>
        public readonly string Comment;
        /// <summary>
        /// HTTP header regular expression.
        /// </summary>
        public readonly string Header;
        /// <summary>
        /// HTTP header group. The structure of `header_group` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetProxyaddressHeaderGroupResult> HeaderGroups;
        /// <summary>
        /// HTTP header.
        /// </summary>
        public readonly string HeaderName;
        /// <summary>
        /// Address object for the host.
        /// </summary>
        public readonly string Host;
        /// <summary>
        /// Host name as a regular expression.
        /// </summary>
        public readonly string HostRegex;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// HTTP request methods to be used.
        /// </summary>
        public readonly string Method;
        /// <summary>
        /// SaaS applicaton name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// URL path as a regular expression.
        /// </summary>
        public readonly string Path;
        /// <summary>
        /// Match the query part of the URL as a regular expression.
        /// </summary>
        public readonly string Query;
        /// <summary>
        /// Enable/disable use of referrer field in the HTTP header to match the address.
        /// </summary>
        public readonly string Referrer;
        /// <summary>
        /// Config object tagging. The structure of `tagging` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetProxyaddressTaggingResult> Taggings;
        /// <summary>
        /// Proxy address type.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Names of browsers to be used as user agent.
        /// </summary>
        public readonly string Ua;
        /// <summary>
        /// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        /// </summary>
        public readonly string Uuid;
        public readonly string? Vdomparam;
        /// <summary>
        /// Enable/disable visibility of the object in the GUI.
        /// </summary>
        public readonly string Visibility;

        [OutputConstructor]
        private GetProxyaddressResult(
            ImmutableArray<Outputs.GetProxyaddressApplicationResult> applications,

            string caseSensitivity,

            ImmutableArray<Outputs.GetProxyaddressCategoryResult> categories,

            int color,

            string comment,

            string header,

            ImmutableArray<Outputs.GetProxyaddressHeaderGroupResult> headerGroups,

            string headerName,

            string host,

            string hostRegex,

            string id,

            string method,

            string name,

            string path,

            string query,

            string referrer,

            ImmutableArray<Outputs.GetProxyaddressTaggingResult> taggings,

            string type,

            string ua,

            string uuid,

            string? vdomparam,

            string visibility)
        {
            Applications = applications;
            CaseSensitivity = caseSensitivity;
            Categories = categories;
            Color = color;
            Comment = comment;
            Header = header;
            HeaderGroups = headerGroups;
            HeaderName = headerName;
            Host = host;
            HostRegex = hostRegex;
            Id = id;
            Method = method;
            Name = name;
            Path = path;
            Query = query;
            Referrer = referrer;
            Taggings = taggings;
            Type = type;
            Ua = ua;
            Uuid = uuid;
            Vdomparam = vdomparam;
            Visibility = visibility;
        }
    }
}
