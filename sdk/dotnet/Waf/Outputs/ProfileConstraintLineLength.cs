// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Waf.Outputs
{

    [OutputType]
    public sealed class ProfileConstraintLineLength
    {
        /// <summary>
        /// Action. Valid values: `allow`, `block`.
        /// </summary>
        public readonly string? Action;
        /// <summary>
        /// Length of HTTP line in bytes (0 to 2147483647).
        /// </summary>
        public readonly int? Length;
        /// <summary>
        /// Enable/disable logging. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? Log;
        /// <summary>
        /// Severity. Valid values: `high`, `medium`, `low`.
        /// </summary>
        public readonly string? Severity;
        /// <summary>
        /// Enable/disable the constraint. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? Status;

        [OutputConstructor]
        private ProfileConstraintLineLength(
            string? action,

            int? length,

            string? log,

            string? severity,

            string? status)
        {
            Action = action;
            Length = length;
            Log = log;
            Severity = severity;
            Status = status;
        }
    }
}
