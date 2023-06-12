// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Router.Outputs
{

    [OutputType]
    public sealed class RipngOffsetList
    {
        /// <summary>
        /// IPv6 access list name.
        /// </summary>
        public readonly string? AccessList6;
        /// <summary>
        /// Offset list direction. Valid values: `in`, `out`.
        /// </summary>
        public readonly string? Direction;
        /// <summary>
        /// Offset-list ID.
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// Interface name.
        /// </summary>
        public readonly string? Interface;
        /// <summary>
        /// offset
        /// </summary>
        public readonly int? Offset;
        /// <summary>
        /// status Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? Status;

        [OutputConstructor]
        private RipngOffsetList(
            string? accessList6,

            string? direction,

            int? id,

            string? @interface,

            int? offset,

            string? status)
        {
            AccessList6 = accessList6;
            Direction = direction;
            Id = id;
            Interface = @interface;
            Offset = offset;
            Status = status;
        }
    }
}
