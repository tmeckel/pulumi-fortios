// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Router.Inputs
{

    public sealed class IsisSummaryAddress6GetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// isis-net ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// Level. Valid values: `level-1-2`, `level-1`, `level-2`.
        /// 
        /// The `summary_address6` block supports:
        /// </summary>
        [Input("level")]
        public Input<string>? Level { get; set; }

        /// <summary>
        /// IPv6 prefix.
        /// </summary>
        [Input("prefix6")]
        public Input<string>? Prefix6 { get; set; }

        public IsisSummaryAddress6GetArgs()
        {
        }
        public static new IsisSummaryAddress6GetArgs Empty => new IsisSummaryAddress6GetArgs();
    }
}