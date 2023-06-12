// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Dnsfilter.Inputs
{

    public sealed class DomainfilterEntryGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Action to take for domain filter matches. Valid values: `block`, `allow`, `monitor`.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// Domain entries to be filtered.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// Id.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// Enable/disable this domain filter. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// DNS domain filter type. Valid values: `simple`, `regex`, `wildcard`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public DomainfilterEntryGetArgs()
        {
        }
        public static new DomainfilterEntryGetArgs Empty => new DomainfilterEntryGetArgs();
    }
}
