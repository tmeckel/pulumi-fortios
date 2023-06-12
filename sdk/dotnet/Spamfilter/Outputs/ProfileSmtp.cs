// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Spamfilter.Outputs
{

    [OutputType]
    public sealed class ProfileSmtp
    {
        /// <summary>
        /// Action for spam email. Valid values: `pass`, `tag`, `discard`.
        /// </summary>
        public readonly string? Action;
        /// <summary>
        /// Enable/disable SMTP email header IP checks for spamfsip, spamrbl and spambwl filters. Valid values: `disable`, `enable`.
        /// </summary>
        public readonly string? Hdrip;
        /// <summary>
        /// Enable/disable local filter to override SMTP remote check result. Valid values: `disable`, `enable`.
        /// </summary>
        public readonly string? LocalOverride;
        /// <summary>
        /// Enable/disable logging. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? Log;
        /// <summary>
        /// Subject text or header added to spam email.
        /// </summary>
        public readonly string? TagMsg;
        /// <summary>
        /// Tag subject or header for spam email. Valid values: `subject`, `header`, `spaminfo`.
        /// </summary>
        public readonly string? TagType;

        [OutputConstructor]
        private ProfileSmtp(
            string? action,

            string? hdrip,

            string? localOverride,

            string? log,

            string? tagMsg,

            string? tagType)
        {
            Action = action;
            Hdrip = hdrip;
            LocalOverride = localOverride;
            Log = log;
            TagMsg = tagMsg;
            TagType = tagType;
        }
    }
}
