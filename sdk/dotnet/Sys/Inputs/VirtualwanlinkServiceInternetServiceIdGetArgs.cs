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

    public sealed class VirtualwanlinkServiceInternetServiceIdGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Internet service ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        public VirtualwanlinkServiceInternetServiceIdGetArgs()
        {
        }
        public static new VirtualwanlinkServiceInternetServiceIdGetArgs Empty => new VirtualwanlinkServiceInternetServiceIdGetArgs();
    }
}
