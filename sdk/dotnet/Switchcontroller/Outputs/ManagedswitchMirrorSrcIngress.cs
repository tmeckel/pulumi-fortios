// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Switchcontroller.Outputs
{

    [OutputType]
    public sealed class ManagedswitchMirrorSrcIngress
    {
        /// <summary>
        /// Interface name.
        /// </summary>
        public readonly string? Name;

        [OutputConstructor]
        private ManagedswitchMirrorSrcIngress(string? name)
        {
            Name = name;
        }
    }
}
