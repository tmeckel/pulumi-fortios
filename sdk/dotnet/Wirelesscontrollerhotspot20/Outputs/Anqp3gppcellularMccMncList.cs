// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Wirelesscontrollerhotspot20.Outputs
{

    [OutputType]
    public sealed class Anqp3gppcellularMccMncList
    {
        /// <summary>
        /// ID.
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// Mobile country code.
        /// </summary>
        public readonly string? Mcc;
        /// <summary>
        /// Mobile network code.
        /// </summary>
        public readonly string? Mnc;

        [OutputConstructor]
        private Anqp3gppcellularMccMncList(
            int? id,

            string? mcc,

            string? mnc)
        {
            Id = id;
            Mcc = mcc;
            Mnc = mnc;
        }
    }
}
