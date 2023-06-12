// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Sys.Inputs
{

    public sealed class DnsServerHostnameGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// DNS server host name list separated by space (maximum 4 domains).
        /// </summary>
        [Input("hostname")]
        public Input<string>? Hostname { get; set; }

        public DnsServerHostnameGetArgs()
        {
        }
        public static new DnsServerHostnameGetArgs Empty => new DnsServerHostnameGetArgs();
    }
}