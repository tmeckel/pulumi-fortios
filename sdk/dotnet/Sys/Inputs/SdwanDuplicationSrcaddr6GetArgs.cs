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

    public sealed class SdwanDuplicationSrcaddr6GetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Physical interface name.
        /// 
        /// 
        /// 
        /// 
        /// 
        /// The `dst6` block supports:
        /// 
        /// 
        /// The `src6` block supports:
        /// 
        /// 
        /// 
        /// 
        /// 
        /// 
        /// 
        /// 
        /// 
        /// 
        /// 
        /// 
        /// 
        /// The `srcaddr6` block supports:
        /// 
        /// 
        /// The `dstaddr6` block supports:
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public SdwanDuplicationSrcaddr6GetArgs()
        {
        }
        public static new SdwanDuplicationSrcaddr6GetArgs Empty => new SdwanDuplicationSrcaddr6GetArgs();
    }
}
