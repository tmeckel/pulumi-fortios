// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Spamfilter.Inputs
{

    public sealed class ProfileMapiArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Action for spam email. Valid values: `pass`, `discard`.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// Enable/disable logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("log")]
        public Input<string>? Log { get; set; }

        public ProfileMapiArgs()
        {
        }
        public static new ProfileMapiArgs Empty => new ProfileMapiArgs();
    }
}
