// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Icap.Inputs
{

    public sealed class ProfileIcapHeaderGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable use of base64 encoding of HTTP content. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("base64Encoding")]
        public Input<string>? Base64Encoding { get; set; }

        /// <summary>
        /// HTTP header content.
        /// </summary>
        [Input("content")]
        public Input<string>? Content { get; set; }

        /// <summary>
        /// HTTP forwarded header ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// HTTP forwarded header name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public ProfileIcapHeaderGetArgs()
        {
        }
        public static new ProfileIcapHeaderGetArgs Empty => new ProfileIcapHeaderGetArgs();
    }
}
