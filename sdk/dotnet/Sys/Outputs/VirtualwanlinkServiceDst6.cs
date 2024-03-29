// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Sys.Outputs
{

    [OutputType]
    public sealed class VirtualwanlinkServiceDst6
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
        /// </summary>
        public readonly string? Name;

        [OutputConstructor]
        private VirtualwanlinkServiceDst6(string? name)
        {
            Name = name;
        }
    }
}
