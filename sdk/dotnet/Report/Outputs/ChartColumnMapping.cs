// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Report.Outputs
{

    [OutputType]
    public sealed class ChartColumnMapping
    {
        /// <summary>
        /// Display name.
        /// </summary>
        public readonly string? Displayname;
        /// <summary>
        /// id
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// Comparision operater. Valid values: `none`, `greater`, `greater-equal`, `less`, `less-equal`, `equal`, `between`.
        /// </summary>
        public readonly string? Op;
        /// <summary>
        /// Value 1.
        /// </summary>
        public readonly string? Value1;
        /// <summary>
        /// Value 2.
        /// </summary>
        public readonly string? Value2;
        /// <summary>
        /// Value type. Valid values: `integer`, `string`.
        /// </summary>
        public readonly string? ValueType;

        [OutputConstructor]
        private ChartColumnMapping(
            string? displayname,

            int? id,

            string? op,

            string? value1,

            string? value2,

            string? valueType)
        {
            Displayname = displayname;
            Id = id;
            Op = op;
            Value1 = value1;
            Value2 = value2;
            ValueType = valueType;
        }
    }
}
