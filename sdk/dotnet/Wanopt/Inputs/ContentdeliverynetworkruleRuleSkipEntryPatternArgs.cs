// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Wanopt.Inputs
{

    public sealed class ContentdeliverynetworkruleRuleSkipEntryPatternArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Pattern strings.
        /// </summary>
        [Input("string")]
        public Input<string>? String { get; set; }

        public ContentdeliverynetworkruleRuleSkipEntryPatternArgs()
        {
        }
        public static new ContentdeliverynetworkruleRuleSkipEntryPatternArgs Empty => new ContentdeliverynetworkruleRuleSkipEntryPatternArgs();
    }
}
