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

    public sealed class ContentdeliverynetworkruleRuleContentIdArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Search direction from end-str match. Valid values: `forward`, `backward`.
        /// </summary>
        [Input("endDirection")]
        public Input<string>? EndDirection { get; set; }

        /// <summary>
        /// Number of characters in URL to skip after end-str has been matched.
        /// </summary>
        [Input("endSkip")]
        public Input<int>? EndSkip { get; set; }

        /// <summary>
        /// String from which to end search.
        /// </summary>
        [Input("endStr")]
        public Input<string>? EndStr { get; set; }

        /// <summary>
        /// Name of content ID within the start string and end string.
        /// </summary>
        [Input("rangeStr")]
        public Input<string>? RangeStr { get; set; }

        /// <summary>
        /// Search direction from start-str match. Valid values: `forward`, `backward`.
        /// </summary>
        [Input("startDirection")]
        public Input<string>? StartDirection { get; set; }

        /// <summary>
        /// Number of characters in URL to skip after start-str has been matched.
        /// </summary>
        [Input("startSkip")]
        public Input<int>? StartSkip { get; set; }

        /// <summary>
        /// String from which to start search.
        /// </summary>
        [Input("startStr")]
        public Input<string>? StartStr { get; set; }

        /// <summary>
        /// Option in HTTP header or URL parameter to match. Valid values: `path`, `parameter`, `referrer`, `youtube-map`, `youtube-id`, `youku-id`, `hls-manifest`, `dash-manifest`, `hls-fragment`, `dash-fragment`.
        /// </summary>
        [Input("target")]
        public Input<string>? Target { get; set; }

        public ContentdeliverynetworkruleRuleContentIdArgs()
        {
        }
        public static new ContentdeliverynetworkruleRuleContentIdArgs Empty => new ContentdeliverynetworkruleRuleContentIdArgs();
    }
}
