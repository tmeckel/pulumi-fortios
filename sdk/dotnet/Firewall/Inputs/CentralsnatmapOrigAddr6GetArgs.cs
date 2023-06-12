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

    public sealed class CentralsnatmapOrigAddr6GetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Address name.
        /// 
        /// The `orig_addr6` block supports:
        /// 
        /// 
        /// 
        /// 
        /// The `dst_addr6` block supports:
        /// 
        /// 
        /// 
        /// 
        /// The `nat_ippool6` block supports:
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public CentralsnatmapOrigAddr6GetArgs()
        {
        }
        public static new CentralsnatmapOrigAddr6GetArgs Empty => new CentralsnatmapOrigAddr6GetArgs();
    }
}
