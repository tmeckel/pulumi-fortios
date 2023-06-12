// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Ips.Outputs
{

    [OutputType]
    public sealed class SensorEntryCfe
    {
        /// <summary>
        /// CVE IDs or CVE wildcards.
        /// </summary>
        public readonly string? CveEntry;

        [OutputConstructor]
        private SensorEntryCfe(string? cveEntry)
        {
            CveEntry = cveEntry;
        }
    }
}
