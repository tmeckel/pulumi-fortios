// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Firewall.Outputs
{

    [OutputType]
    public sealed class GetProxypolicyDstintfResult
    {
        /// <summary>
        /// Group name.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetProxypolicyDstintfResult(string name)
        {
            Name = name;
        }
    }
}
