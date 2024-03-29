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
    public sealed class AnqpvenueurlValueList
    {
        /// <summary>
        /// URL index.
        /// </summary>
        public readonly int? Index;
        /// <summary>
        /// Venue number.
        /// </summary>
        public readonly int? Number;
        /// <summary>
        /// Venue URL value.
        /// </summary>
        public readonly string? Value;

        [OutputConstructor]
        private AnqpvenueurlValueList(
            int? index,

            int? number,

            string? value)
        {
            Index = index;
            Number = number;
            Value = value;
        }
    }
}
