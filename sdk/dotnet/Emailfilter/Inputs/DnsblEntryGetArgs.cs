// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Emailfilter.Inputs
{

    public sealed class DnsblEntryGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Reject connection or mark as spam email. Valid values: `reject`, `spam`.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// DNSBL/ORBL entry ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// DNSBL or ORBL server name.
        /// </summary>
        [Input("server")]
        public Input<string>? Server { get; set; }

        /// <summary>
        /// Enable/disable status. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public DnsblEntryGetArgs()
        {
        }
        public static new DnsblEntryGetArgs Empty => new DnsblEntryGetArgs();
    }
}
