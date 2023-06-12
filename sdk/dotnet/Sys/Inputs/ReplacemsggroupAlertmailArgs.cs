// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Sys.Inputs
{

    public sealed class ReplacemsggroupAlertmailArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Message string.
        /// </summary>
        [Input("buffer")]
        public Input<string>? Buffer { get; set; }

        /// <summary>
        /// Format flag.
        /// </summary>
        [Input("format")]
        public Input<string>? Format { get; set; }

        /// <summary>
        /// Header flag. Valid values: `none`, `http`, `8bit`.
        /// </summary>
        [Input("header")]
        public Input<string>? Header { get; set; }

        /// <summary>
        /// Message type.
        /// </summary>
        [Input("msgType")]
        public Input<string>? MsgType { get; set; }

        public ReplacemsggroupAlertmailArgs()
        {
        }
        public static new ReplacemsggroupAlertmailArgs Empty => new ReplacemsggroupAlertmailArgs();
    }
}
