// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Firewall.Inputs
{

    public sealed class VipSrcFilterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Source-filter range.
        /// </summary>
        [Input("range")]
        public Input<string>? Range { get; set; }

        public VipSrcFilterArgs()
        {
        }
        public static new VipSrcFilterArgs Empty => new VipSrcFilterArgs();
    }
}
