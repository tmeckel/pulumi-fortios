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

    public sealed class SdwanDuplicationSrcintfGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Interface, zone or SDWAN zone name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public SdwanDuplicationSrcintfGetArgs()
        {
        }
        public static new SdwanDuplicationSrcintfGetArgs Empty => new SdwanDuplicationSrcintfGetArgs();
    }
}
