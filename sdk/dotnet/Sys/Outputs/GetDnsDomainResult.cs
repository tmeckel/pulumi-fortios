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
    public sealed class GetDnsDomainResult
    {
        /// <summary>
        /// DNS search domain list separated by space (maximum 8 domains)
        /// </summary>
        public readonly string Domain;

        [OutputConstructor]
        private GetDnsDomainResult(string domain)
        {
            Domain = domain;
        }
    }
}
