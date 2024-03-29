// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Application.Inputs
{

    public sealed class ListEntryArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Pass or block traffic, or reset connection for traffic from this application. Valid values: `pass`, `block`, `reset`.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        [Input("applications")]
        private InputList<Inputs.ListEntryApplicationArgs>? _applications;

        /// <summary>
        /// ID of allowed applications. The structure of `application` block is documented below.
        /// </summary>
        public InputList<Inputs.ListEntryApplicationArgs> Applications
        {
            get => _applications ?? (_applications = new InputList<Inputs.ListEntryApplicationArgs>());
            set => _applications = value;
        }

        /// <summary>
        /// Application behavior filter.
        /// </summary>
        [Input("behavior")]
        public Input<string>? Behavior { get; set; }

        [Input("categories")]
        private InputList<Inputs.ListEntryCategoryArgs>? _categories;

        /// <summary>
        /// Category ID list. The structure of `category` block is documented below.
        /// </summary>
        public InputList<Inputs.ListEntryCategoryArgs> Categories
        {
            get => _categories ?? (_categories = new InputList<Inputs.ListEntryCategoryArgs>());
            set => _categories = value;
        }

        [Input("exclusions")]
        private InputList<Inputs.ListEntryExclusionArgs>? _exclusions;

        /// <summary>
        /// ID of excluded applications. The structure of `exclusion` block is documented below.
        /// </summary>
        public InputList<Inputs.ListEntryExclusionArgs> Exclusions
        {
            get => _exclusions ?? (_exclusions = new InputList<Inputs.ListEntryExclusionArgs>());
            set => _exclusions = value;
        }

        /// <summary>
        /// Entry ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// Enable/disable logging for this application list. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("log")]
        public Input<string>? Log { get; set; }

        /// <summary>
        /// Enable/disable packet logging. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("logPacket")]
        public Input<string>? LogPacket { get; set; }

        [Input("parameters")]
        private InputList<Inputs.ListEntryParameterArgs>? _parameters;

        /// <summary>
        /// Application parameters. The structure of `parameters` block is documented below.
        /// </summary>
        public InputList<Inputs.ListEntryParameterArgs> Parameters
        {
            get => _parameters ?? (_parameters = new InputList<Inputs.ListEntryParameterArgs>());
            set => _parameters = value;
        }

        /// <summary>
        /// Per-IP traffic shaper.
        /// </summary>
        [Input("perIpShaper")]
        public Input<string>? PerIpShaper { get; set; }

        /// <summary>
        /// Application popularity filter (1 - 5, from least to most popular). Valid values: `1`, `2`, `3`, `4`, `5`.
        /// </summary>
        [Input("popularity")]
        public Input<string>? Popularity { get; set; }

        /// <summary>
        /// Application protocol filter.
        /// </summary>
        [Input("protocols")]
        public Input<string>? Protocols { get; set; }

        /// <summary>
        /// Quarantine method. Valid values: `none`, `attacker`.
        /// </summary>
        [Input("quarantine")]
        public Input<string>? Quarantine { get; set; }

        /// <summary>
        /// Duration of quarantine. (Format ###d##h##m, minimum 1m, maximum 364d23h59m, default = 5m). Requires quarantine set to attacker.
        /// </summary>
        [Input("quarantineExpiry")]
        public Input<string>? QuarantineExpiry { get; set; }

        /// <summary>
        /// Enable/disable quarantine logging. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("quarantineLog")]
        public Input<string>? QuarantineLog { get; set; }

        /// <summary>
        /// Count of the rate.
        /// </summary>
        [Input("rateCount")]
        public Input<int>? RateCount { get; set; }

        /// <summary>
        /// Duration (sec) of the rate.
        /// </summary>
        [Input("rateDuration")]
        public Input<int>? RateDuration { get; set; }

        /// <summary>
        /// Rate limit mode. Valid values: `periodical`, `continuous`.
        /// </summary>
        [Input("rateMode")]
        public Input<string>? RateMode { get; set; }

        /// <summary>
        /// Track the packet protocol field. Valid values: `none`, `src-ip`, `dest-ip`, `dhcp-client-mac`, `dns-domain`.
        /// </summary>
        [Input("rateTrack")]
        public Input<string>? RateTrack { get; set; }

        [Input("risks")]
        private InputList<Inputs.ListEntryRiskArgs>? _risks;

        /// <summary>
        /// Risk, or impact, of allowing traffic from this application to occur (1 - 5; Low, Elevated, Medium, High, and Critical). The structure of `risk` block is documented below.
        /// </summary>
        public InputList<Inputs.ListEntryRiskArgs> Risks
        {
            get => _risks ?? (_risks = new InputList<Inputs.ListEntryRiskArgs>());
            set => _risks = value;
        }

        /// <summary>
        /// Session TTL (0 = default).
        /// </summary>
        [Input("sessionTtl")]
        public Input<int>? SessionTtl { get; set; }

        /// <summary>
        /// Traffic shaper.
        /// </summary>
        [Input("shaper")]
        public Input<string>? Shaper { get; set; }

        /// <summary>
        /// Reverse traffic shaper.
        /// </summary>
        [Input("shaperReverse")]
        public Input<string>? ShaperReverse { get; set; }

        [Input("subCategories")]
        private InputList<Inputs.ListEntrySubCategoryArgs>? _subCategories;

        /// <summary>
        /// Application Sub-category ID list. The structure of `sub_category` block is documented below.
        /// </summary>
        public InputList<Inputs.ListEntrySubCategoryArgs> SubCategories
        {
            get => _subCategories ?? (_subCategories = new InputList<Inputs.ListEntrySubCategoryArgs>());
            set => _subCategories = value;
        }

        /// <summary>
        /// Application technology filter.
        /// </summary>
        [Input("technology")]
        public Input<string>? Technology { get; set; }

        /// <summary>
        /// Application vendor filter.
        /// </summary>
        [Input("vendor")]
        public Input<string>? Vendor { get; set; }

        public ListEntryArgs()
        {
        }
        public static new ListEntryArgs Empty => new ListEntryArgs();
    }
}
