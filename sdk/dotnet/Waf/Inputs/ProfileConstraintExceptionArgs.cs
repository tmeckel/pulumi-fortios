// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Waf.Inputs
{

    public sealed class ProfileConstraintExceptionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Host address.
        /// </summary>
        [Input("address")]
        public Input<string>? Address { get; set; }

        /// <summary>
        /// HTTP content length in request. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("contentLength")]
        public Input<string>? ContentLength { get; set; }

        /// <summary>
        /// HTTP header length in request. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("headerLength")]
        public Input<string>? HeaderLength { get; set; }

        /// <summary>
        /// Enable/disable hostname check. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("hostname")]
        public Input<string>? Hostname { get; set; }

        /// <summary>
        /// Exception ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// HTTP line length in request. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("lineLength")]
        public Input<string>? LineLength { get; set; }

        /// <summary>
        /// Enable/disable malformed HTTP request check. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("malformed")]
        public Input<string>? Malformed { get; set; }

        /// <summary>
        /// Maximum number of cookies in HTTP request. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("maxCookie")]
        public Input<string>? MaxCookie { get; set; }

        /// <summary>
        /// Maximum number of HTTP header line. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("maxHeaderLine")]
        public Input<string>? MaxHeaderLine { get; set; }

        /// <summary>
        /// Maximum number of range segments in HTTP range line. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("maxRangeSegment")]
        public Input<string>? MaxRangeSegment { get; set; }

        /// <summary>
        /// Maximum number of parameters in URL. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("maxUrlParam")]
        public Input<string>? MaxUrlParam { get; set; }

        /// <summary>
        /// Enable/disable HTTP method check. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("method")]
        public Input<string>? Method { get; set; }

        /// <summary>
        /// Maximum length of parameter in URL, HTTP POST request or HTTP body. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("paramLength")]
        public Input<string>? ParamLength { get; set; }

        /// <summary>
        /// URL pattern.
        /// </summary>
        [Input("pattern")]
        public Input<string>? Pattern { get; set; }

        /// <summary>
        /// Enable/disable regular expression based pattern match. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("regex")]
        public Input<string>? Regex { get; set; }

        /// <summary>
        /// Maximum length of parameter in URL. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("urlParamLength")]
        public Input<string>? UrlParamLength { get; set; }

        /// <summary>
        /// Enable/disable HTTP version check. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public ProfileConstraintExceptionArgs()
        {
        }
        public static new ProfileConstraintExceptionArgs Empty => new ProfileConstraintExceptionArgs();
    }
}
