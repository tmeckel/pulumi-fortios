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
    public sealed class GetZoneTaggingTagResult
    {
        /// <summary>
        /// Specify the name of the desired system zone.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetZoneTaggingTagResult(string name)
        {
            Name = name;
        }
    }
}