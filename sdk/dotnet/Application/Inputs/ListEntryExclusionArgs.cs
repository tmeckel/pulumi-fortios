// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Application.Inputs
{

    public sealed class ListEntryExclusionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Excluded application IDs.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        public ListEntryExclusionArgs()
        {
        }
        public static new ListEntryExclusionArgs Empty => new ListEntryExclusionArgs();
    }
}
