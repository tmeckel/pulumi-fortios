// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Webfilter
{
    /// <summary>
    /// Configure Web filter profiles.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Fortios = Pulumiverse.Fortios;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var trname = new Fortios.Webfilter.Profile("trname", new()
    ///     {
    ///         ExtendedLog = "disable",
    ///         FtgdWf = new Fortios.Webfilter.Inputs.ProfileFtgdWfArgs
    ///         {
    ///             ExemptQuota = "17",
    ///             Filters = new[]
    ///             {
    ///                 new Fortios.Webfilter.Inputs.ProfileFtgdWfFilterArgs
    ///                 {
    ///                     Action = "warning",
    ///                     Category = 2,
    ///                     Id = 1,
    ///                     Log = "enable",
    ///                     WarnDuration = "5m",
    ///                     WarningDurationType = "timeout",
    ///                     WarningPrompt = "per-category",
    ///                 },
    ///                 new Fortios.Webfilter.Inputs.ProfileFtgdWfFilterArgs
    ///                 {
    ///                     Action = "warning",
    ///                     Category = 7,
    ///                     Id = 2,
    ///                     Log = "enable",
    ///                     WarnDuration = "5m",
    ///                     WarningDurationType = "timeout",
    ///                     WarningPrompt = "per-category",
    ///                 },
    ///             },
    ///             MaxQuotaTimeout = 300,
    ///             RateCrlUrls = "enable",
    ///             RateCssUrls = "enable",
    ///             RateImageUrls = "enable",
    ///             RateJavascriptUrls = "enable",
    ///         },
    ///         HttpsReplacemsg = "enable",
    ///         InspectionMode = "flow-based",
    ///         LogAllUrl = "disable",
    ///         Override = new Fortios.Webfilter.Inputs.ProfileOverrideArgs
    ///         {
    ///             OvrdCookie = "deny",
    ///             OvrdDur = "15m",
    ///             OvrdDurMode = "constant",
    ///             OvrdScope = "user",
    ///             ProfileAttribute = "Login-LAT-Service",
    ///             ProfileType = "list",
    ///         },
    ///         PostAction = "normal",
    ///         Web = new Fortios.Webfilter.Inputs.ProfileWebArgs
    ///         {
    ///             Blacklist = "disable",
    ///             BwordTable = 0,
    ///             BwordThreshold = 10,
    ///             ContentHeaderList = 0,
    ///             LogSearch = "disable",
    ///             UrlfilterTable = 0,
    ///             YoutubeRestrict = "none",
    ///         },
    ///         WebContentLog = "enable",
    ///         WebExtendedAllActionLog = "disable",
    ///         WebFilterActivexLog = "enable",
    ///         WebFilterAppletLog = "enable",
    ///         WebFilterCommandBlockLog = "enable",
    ///         WebFilterCookieLog = "enable",
    ///         WebFilterCookieRemovalLog = "enable",
    ///         WebFilterJsLog = "enable",
    ///         WebFilterJscriptLog = "enable",
    ///         WebFilterRefererLog = "enable",
    ///         WebFilterUnknownLog = "enable",
    ///         WebFilterVbsLog = "enable",
    ///         WebFtgdErrLog = "enable",
    ///         WebFtgdQuotaUsage = "enable",
    ///         WebInvalidDomainLog = "enable",
    ///         WebUrlLog = "enable",
    ///         Wisp = "disable",
    ///         WispAlgorithm = "auto-learning",
    ///         YoutubeChannelStatus = "disable",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Webfilter Profile can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:webfilter/profile:Profile labelname {{name}}
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:webfilter/profile:Profile labelname {{name}}
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:webfilter/profile:Profile")]
    public partial class Profile : global::Pulumi.CustomResource
    {
        /// <summary>
        /// AntiPhishing profile. The structure of `antiphish` block is documented below.
        /// </summary>
        [Output("antiphish")]
        public Output<Outputs.ProfileAntiphish> Antiphish { get; private set; } = null!;

        /// <summary>
        /// Optional comments.
        /// </summary>
        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// Enable/disable extended logging for web filtering. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("extendedLog")]
        public Output<string> ExtendedLog { get; private set; } = null!;

        /// <summary>
        /// Flow/proxy feature set. Valid values: `flow`, `proxy`.
        /// </summary>
        [Output("featureSet")]
        public Output<string> FeatureSet { get; private set; } = null!;

        /// <summary>
        /// File filter. The structure of `file_filter` block is documented below.
        /// </summary>
        [Output("fileFilter")]
        public Output<Outputs.ProfileFileFilter> FileFilter { get; private set; } = null!;

        /// <summary>
        /// FortiGuard Web Filter settings. The structure of `ftgd_wf` block is documented below.
        /// </summary>
        [Output("ftgdWf")]
        public Output<Outputs.ProfileFtgdWf> FtgdWf { get; private set; } = null!;

        /// <summary>
        /// Enable replacement messages for HTTPS. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("httpsReplacemsg")]
        public Output<string> HttpsReplacemsg { get; private set; } = null!;

        /// <summary>
        /// Web filtering inspection mode. Valid values: `proxy`, `flow-based`.
        /// </summary>
        [Output("inspectionMode")]
        public Output<string> InspectionMode { get; private set; } = null!;

        /// <summary>
        /// Enable/disable logging all URLs visited. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("logAllUrl")]
        public Output<string> LogAllUrl { get; private set; } = null!;

        /// <summary>
        /// Profile name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Options.
        /// </summary>
        [Output("options")]
        public Output<string> Options { get; private set; } = null!;

        /// <summary>
        /// Web Filter override settings. The structure of `override` block is documented below.
        /// </summary>
        [Output("override")]
        public Output<Outputs.ProfileOverride> Override { get; private set; } = null!;

        /// <summary>
        /// Permitted override types. Valid values: `bannedword-override`, `urlfilter-override`, `fortiguard-wf-override`, `contenttype-check-override`.
        /// </summary>
        [Output("ovrdPerm")]
        public Output<string> OvrdPerm { get; private set; } = null!;

        /// <summary>
        /// Action taken for HTTP POST traffic. Valid values: `normal`, `block`.
        /// </summary>
        [Output("postAction")]
        public Output<string> PostAction { get; private set; } = null!;

        /// <summary>
        /// Replacement message group.
        /// </summary>
        [Output("replacemsgGroup")]
        public Output<string> ReplacemsgGroup { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;

        /// <summary>
        /// Web content filtering settings. The structure of `web` block is documented below.
        /// </summary>
        [Output("web")]
        public Output<Outputs.ProfileWeb> Web { get; private set; } = null!;

        /// <summary>
        /// Enable/disable logging of AntiPhishing checks. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("webAntiphishingLog")]
        public Output<string> WebAntiphishingLog { get; private set; } = null!;

        /// <summary>
        /// Enable/disable logging logging blocked web content. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("webContentLog")]
        public Output<string> WebContentLog { get; private set; } = null!;

        /// <summary>
        /// Enable/disable extended any filter action logging for web filtering. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("webExtendedAllActionLog")]
        public Output<string> WebExtendedAllActionLog { get; private set; } = null!;

        /// <summary>
        /// Enable/disable logging ActiveX. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("webFilterActivexLog")]
        public Output<string> WebFilterActivexLog { get; private set; } = null!;

        /// <summary>
        /// Enable/disable logging Java applets. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("webFilterAppletLog")]
        public Output<string> WebFilterAppletLog { get; private set; } = null!;

        /// <summary>
        /// Enable/disable logging blocked commands. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("webFilterCommandBlockLog")]
        public Output<string> WebFilterCommandBlockLog { get; private set; } = null!;

        /// <summary>
        /// Enable/disable logging cookie filtering. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("webFilterCookieLog")]
        public Output<string> WebFilterCookieLog { get; private set; } = null!;

        /// <summary>
        /// Enable/disable logging blocked cookies. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("webFilterCookieRemovalLog")]
        public Output<string> WebFilterCookieRemovalLog { get; private set; } = null!;

        /// <summary>
        /// Enable/disable logging Java scripts. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("webFilterJsLog")]
        public Output<string> WebFilterJsLog { get; private set; } = null!;

        /// <summary>
        /// Enable/disable logging JScripts. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("webFilterJscriptLog")]
        public Output<string> WebFilterJscriptLog { get; private set; } = null!;

        /// <summary>
        /// Enable/disable logging referrers. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("webFilterRefererLog")]
        public Output<string> WebFilterRefererLog { get; private set; } = null!;

        /// <summary>
        /// Enable/disable logging unknown scripts. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("webFilterUnknownLog")]
        public Output<string> WebFilterUnknownLog { get; private set; } = null!;

        /// <summary>
        /// Enable/disable logging VBS scripts. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("webFilterVbsLog")]
        public Output<string> WebFilterVbsLog { get; private set; } = null!;

        /// <summary>
        /// Enable/disable logging rating errors. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("webFtgdErrLog")]
        public Output<string> WebFtgdErrLog { get; private set; } = null!;

        /// <summary>
        /// Enable/disable logging daily quota usage. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("webFtgdQuotaUsage")]
        public Output<string> WebFtgdQuotaUsage { get; private set; } = null!;

        /// <summary>
        /// Enable/disable logging invalid domain names. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("webInvalidDomainLog")]
        public Output<string> WebInvalidDomainLog { get; private set; } = null!;

        /// <summary>
        /// Enable/disable logging URL filtering. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("webUrlLog")]
        public Output<string> WebUrlLog { get; private set; } = null!;

        /// <summary>
        /// Enable/disable web proxy WISP. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("wisp")]
        public Output<string> Wisp { get; private set; } = null!;

        /// <summary>
        /// WISP server selection algorithm. Valid values: `primary-secondary`, `round-robin`, `auto-learning`.
        /// </summary>
        [Output("wispAlgorithm")]
        public Output<string> WispAlgorithm { get; private set; } = null!;

        /// <summary>
        /// WISP servers. The structure of `wisp_servers` block is documented below.
        /// </summary>
        [Output("wispServers")]
        public Output<ImmutableArray<Outputs.ProfileWispServer>> WispServers { get; private set; } = null!;

        /// <summary>
        /// YouTube channel filter. The structure of `youtube_channel_filter` block is documented below.
        /// </summary>
        [Output("youtubeChannelFilters")]
        public Output<ImmutableArray<Outputs.ProfileYoutubeChannelFilter>> YoutubeChannelFilters { get; private set; } = null!;

        /// <summary>
        /// YouTube channel filter status.
        /// </summary>
        [Output("youtubeChannelStatus")]
        public Output<string> YoutubeChannelStatus { get; private set; } = null!;


        /// <summary>
        /// Create a Profile resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Profile(string name, ProfileArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:webfilter/profile:Profile", name, args ?? new ProfileArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Profile(string name, Input<string> id, ProfileState? state = null, CustomResourceOptions? options = null)
            : base("fortios:webfilter/profile:Profile", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse/pulumi-fortios",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Profile resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Profile Get(string name, Input<string> id, ProfileState? state = null, CustomResourceOptions? options = null)
        {
            return new Profile(name, id, state, options);
        }
    }

    public sealed class ProfileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// AntiPhishing profile. The structure of `antiphish` block is documented below.
        /// </summary>
        [Input("antiphish")]
        public Input<Inputs.ProfileAntiphishArgs>? Antiphish { get; set; }

        /// <summary>
        /// Optional comments.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Enable/disable extended logging for web filtering. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("extendedLog")]
        public Input<string>? ExtendedLog { get; set; }

        /// <summary>
        /// Flow/proxy feature set. Valid values: `flow`, `proxy`.
        /// </summary>
        [Input("featureSet")]
        public Input<string>? FeatureSet { get; set; }

        /// <summary>
        /// File filter. The structure of `file_filter` block is documented below.
        /// </summary>
        [Input("fileFilter")]
        public Input<Inputs.ProfileFileFilterArgs>? FileFilter { get; set; }

        /// <summary>
        /// FortiGuard Web Filter settings. The structure of `ftgd_wf` block is documented below.
        /// </summary>
        [Input("ftgdWf")]
        public Input<Inputs.ProfileFtgdWfArgs>? FtgdWf { get; set; }

        /// <summary>
        /// Enable replacement messages for HTTPS. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("httpsReplacemsg")]
        public Input<string>? HttpsReplacemsg { get; set; }

        /// <summary>
        /// Web filtering inspection mode. Valid values: `proxy`, `flow-based`.
        /// </summary>
        [Input("inspectionMode")]
        public Input<string>? InspectionMode { get; set; }

        /// <summary>
        /// Enable/disable logging all URLs visited. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("logAllUrl")]
        public Input<string>? LogAllUrl { get; set; }

        /// <summary>
        /// Profile name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Options.
        /// </summary>
        [Input("options")]
        public Input<string>? Options { get; set; }

        /// <summary>
        /// Web Filter override settings. The structure of `override` block is documented below.
        /// </summary>
        [Input("override")]
        public Input<Inputs.ProfileOverrideArgs>? Override { get; set; }

        /// <summary>
        /// Permitted override types. Valid values: `bannedword-override`, `urlfilter-override`, `fortiguard-wf-override`, `contenttype-check-override`.
        /// </summary>
        [Input("ovrdPerm")]
        public Input<string>? OvrdPerm { get; set; }

        /// <summary>
        /// Action taken for HTTP POST traffic. Valid values: `normal`, `block`.
        /// </summary>
        [Input("postAction")]
        public Input<string>? PostAction { get; set; }

        /// <summary>
        /// Replacement message group.
        /// </summary>
        [Input("replacemsgGroup")]
        public Input<string>? ReplacemsgGroup { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        /// <summary>
        /// Web content filtering settings. The structure of `web` block is documented below.
        /// </summary>
        [Input("web")]
        public Input<Inputs.ProfileWebArgs>? Web { get; set; }

        /// <summary>
        /// Enable/disable logging of AntiPhishing checks. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("webAntiphishingLog")]
        public Input<string>? WebAntiphishingLog { get; set; }

        /// <summary>
        /// Enable/disable logging logging blocked web content. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("webContentLog")]
        public Input<string>? WebContentLog { get; set; }

        /// <summary>
        /// Enable/disable extended any filter action logging for web filtering. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("webExtendedAllActionLog")]
        public Input<string>? WebExtendedAllActionLog { get; set; }

        /// <summary>
        /// Enable/disable logging ActiveX. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("webFilterActivexLog")]
        public Input<string>? WebFilterActivexLog { get; set; }

        /// <summary>
        /// Enable/disable logging Java applets. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("webFilterAppletLog")]
        public Input<string>? WebFilterAppletLog { get; set; }

        /// <summary>
        /// Enable/disable logging blocked commands. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("webFilterCommandBlockLog")]
        public Input<string>? WebFilterCommandBlockLog { get; set; }

        /// <summary>
        /// Enable/disable logging cookie filtering. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("webFilterCookieLog")]
        public Input<string>? WebFilterCookieLog { get; set; }

        /// <summary>
        /// Enable/disable logging blocked cookies. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("webFilterCookieRemovalLog")]
        public Input<string>? WebFilterCookieRemovalLog { get; set; }

        /// <summary>
        /// Enable/disable logging Java scripts. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("webFilterJsLog")]
        public Input<string>? WebFilterJsLog { get; set; }

        /// <summary>
        /// Enable/disable logging JScripts. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("webFilterJscriptLog")]
        public Input<string>? WebFilterJscriptLog { get; set; }

        /// <summary>
        /// Enable/disable logging referrers. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("webFilterRefererLog")]
        public Input<string>? WebFilterRefererLog { get; set; }

        /// <summary>
        /// Enable/disable logging unknown scripts. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("webFilterUnknownLog")]
        public Input<string>? WebFilterUnknownLog { get; set; }

        /// <summary>
        /// Enable/disable logging VBS scripts. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("webFilterVbsLog")]
        public Input<string>? WebFilterVbsLog { get; set; }

        /// <summary>
        /// Enable/disable logging rating errors. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("webFtgdErrLog")]
        public Input<string>? WebFtgdErrLog { get; set; }

        /// <summary>
        /// Enable/disable logging daily quota usage. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("webFtgdQuotaUsage")]
        public Input<string>? WebFtgdQuotaUsage { get; set; }

        /// <summary>
        /// Enable/disable logging invalid domain names. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("webInvalidDomainLog")]
        public Input<string>? WebInvalidDomainLog { get; set; }

        /// <summary>
        /// Enable/disable logging URL filtering. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("webUrlLog")]
        public Input<string>? WebUrlLog { get; set; }

        /// <summary>
        /// Enable/disable web proxy WISP. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("wisp")]
        public Input<string>? Wisp { get; set; }

        /// <summary>
        /// WISP server selection algorithm. Valid values: `primary-secondary`, `round-robin`, `auto-learning`.
        /// </summary>
        [Input("wispAlgorithm")]
        public Input<string>? WispAlgorithm { get; set; }

        [Input("wispServers")]
        private InputList<Inputs.ProfileWispServerArgs>? _wispServers;

        /// <summary>
        /// WISP servers. The structure of `wisp_servers` block is documented below.
        /// </summary>
        public InputList<Inputs.ProfileWispServerArgs> WispServers
        {
            get => _wispServers ?? (_wispServers = new InputList<Inputs.ProfileWispServerArgs>());
            set => _wispServers = value;
        }

        [Input("youtubeChannelFilters")]
        private InputList<Inputs.ProfileYoutubeChannelFilterArgs>? _youtubeChannelFilters;

        /// <summary>
        /// YouTube channel filter. The structure of `youtube_channel_filter` block is documented below.
        /// </summary>
        public InputList<Inputs.ProfileYoutubeChannelFilterArgs> YoutubeChannelFilters
        {
            get => _youtubeChannelFilters ?? (_youtubeChannelFilters = new InputList<Inputs.ProfileYoutubeChannelFilterArgs>());
            set => _youtubeChannelFilters = value;
        }

        /// <summary>
        /// YouTube channel filter status.
        /// </summary>
        [Input("youtubeChannelStatus")]
        public Input<string>? YoutubeChannelStatus { get; set; }

        public ProfileArgs()
        {
        }
        public static new ProfileArgs Empty => new ProfileArgs();
    }

    public sealed class ProfileState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// AntiPhishing profile. The structure of `antiphish` block is documented below.
        /// </summary>
        [Input("antiphish")]
        public Input<Inputs.ProfileAntiphishGetArgs>? Antiphish { get; set; }

        /// <summary>
        /// Optional comments.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Enable/disable extended logging for web filtering. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("extendedLog")]
        public Input<string>? ExtendedLog { get; set; }

        /// <summary>
        /// Flow/proxy feature set. Valid values: `flow`, `proxy`.
        /// </summary>
        [Input("featureSet")]
        public Input<string>? FeatureSet { get; set; }

        /// <summary>
        /// File filter. The structure of `file_filter` block is documented below.
        /// </summary>
        [Input("fileFilter")]
        public Input<Inputs.ProfileFileFilterGetArgs>? FileFilter { get; set; }

        /// <summary>
        /// FortiGuard Web Filter settings. The structure of `ftgd_wf` block is documented below.
        /// </summary>
        [Input("ftgdWf")]
        public Input<Inputs.ProfileFtgdWfGetArgs>? FtgdWf { get; set; }

        /// <summary>
        /// Enable replacement messages for HTTPS. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("httpsReplacemsg")]
        public Input<string>? HttpsReplacemsg { get; set; }

        /// <summary>
        /// Web filtering inspection mode. Valid values: `proxy`, `flow-based`.
        /// </summary>
        [Input("inspectionMode")]
        public Input<string>? InspectionMode { get; set; }

        /// <summary>
        /// Enable/disable logging all URLs visited. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("logAllUrl")]
        public Input<string>? LogAllUrl { get; set; }

        /// <summary>
        /// Profile name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Options.
        /// </summary>
        [Input("options")]
        public Input<string>? Options { get; set; }

        /// <summary>
        /// Web Filter override settings. The structure of `override` block is documented below.
        /// </summary>
        [Input("override")]
        public Input<Inputs.ProfileOverrideGetArgs>? Override { get; set; }

        /// <summary>
        /// Permitted override types. Valid values: `bannedword-override`, `urlfilter-override`, `fortiguard-wf-override`, `contenttype-check-override`.
        /// </summary>
        [Input("ovrdPerm")]
        public Input<string>? OvrdPerm { get; set; }

        /// <summary>
        /// Action taken for HTTP POST traffic. Valid values: `normal`, `block`.
        /// </summary>
        [Input("postAction")]
        public Input<string>? PostAction { get; set; }

        /// <summary>
        /// Replacement message group.
        /// </summary>
        [Input("replacemsgGroup")]
        public Input<string>? ReplacemsgGroup { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        /// <summary>
        /// Web content filtering settings. The structure of `web` block is documented below.
        /// </summary>
        [Input("web")]
        public Input<Inputs.ProfileWebGetArgs>? Web { get; set; }

        /// <summary>
        /// Enable/disable logging of AntiPhishing checks. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("webAntiphishingLog")]
        public Input<string>? WebAntiphishingLog { get; set; }

        /// <summary>
        /// Enable/disable logging logging blocked web content. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("webContentLog")]
        public Input<string>? WebContentLog { get; set; }

        /// <summary>
        /// Enable/disable extended any filter action logging for web filtering. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("webExtendedAllActionLog")]
        public Input<string>? WebExtendedAllActionLog { get; set; }

        /// <summary>
        /// Enable/disable logging ActiveX. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("webFilterActivexLog")]
        public Input<string>? WebFilterActivexLog { get; set; }

        /// <summary>
        /// Enable/disable logging Java applets. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("webFilterAppletLog")]
        public Input<string>? WebFilterAppletLog { get; set; }

        /// <summary>
        /// Enable/disable logging blocked commands. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("webFilterCommandBlockLog")]
        public Input<string>? WebFilterCommandBlockLog { get; set; }

        /// <summary>
        /// Enable/disable logging cookie filtering. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("webFilterCookieLog")]
        public Input<string>? WebFilterCookieLog { get; set; }

        /// <summary>
        /// Enable/disable logging blocked cookies. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("webFilterCookieRemovalLog")]
        public Input<string>? WebFilterCookieRemovalLog { get; set; }

        /// <summary>
        /// Enable/disable logging Java scripts. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("webFilterJsLog")]
        public Input<string>? WebFilterJsLog { get; set; }

        /// <summary>
        /// Enable/disable logging JScripts. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("webFilterJscriptLog")]
        public Input<string>? WebFilterJscriptLog { get; set; }

        /// <summary>
        /// Enable/disable logging referrers. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("webFilterRefererLog")]
        public Input<string>? WebFilterRefererLog { get; set; }

        /// <summary>
        /// Enable/disable logging unknown scripts. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("webFilterUnknownLog")]
        public Input<string>? WebFilterUnknownLog { get; set; }

        /// <summary>
        /// Enable/disable logging VBS scripts. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("webFilterVbsLog")]
        public Input<string>? WebFilterVbsLog { get; set; }

        /// <summary>
        /// Enable/disable logging rating errors. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("webFtgdErrLog")]
        public Input<string>? WebFtgdErrLog { get; set; }

        /// <summary>
        /// Enable/disable logging daily quota usage. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("webFtgdQuotaUsage")]
        public Input<string>? WebFtgdQuotaUsage { get; set; }

        /// <summary>
        /// Enable/disable logging invalid domain names. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("webInvalidDomainLog")]
        public Input<string>? WebInvalidDomainLog { get; set; }

        /// <summary>
        /// Enable/disable logging URL filtering. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("webUrlLog")]
        public Input<string>? WebUrlLog { get; set; }

        /// <summary>
        /// Enable/disable web proxy WISP. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("wisp")]
        public Input<string>? Wisp { get; set; }

        /// <summary>
        /// WISP server selection algorithm. Valid values: `primary-secondary`, `round-robin`, `auto-learning`.
        /// </summary>
        [Input("wispAlgorithm")]
        public Input<string>? WispAlgorithm { get; set; }

        [Input("wispServers")]
        private InputList<Inputs.ProfileWispServerGetArgs>? _wispServers;

        /// <summary>
        /// WISP servers. The structure of `wisp_servers` block is documented below.
        /// </summary>
        public InputList<Inputs.ProfileWispServerGetArgs> WispServers
        {
            get => _wispServers ?? (_wispServers = new InputList<Inputs.ProfileWispServerGetArgs>());
            set => _wispServers = value;
        }

        [Input("youtubeChannelFilters")]
        private InputList<Inputs.ProfileYoutubeChannelFilterGetArgs>? _youtubeChannelFilters;

        /// <summary>
        /// YouTube channel filter. The structure of `youtube_channel_filter` block is documented below.
        /// </summary>
        public InputList<Inputs.ProfileYoutubeChannelFilterGetArgs> YoutubeChannelFilters
        {
            get => _youtubeChannelFilters ?? (_youtubeChannelFilters = new InputList<Inputs.ProfileYoutubeChannelFilterGetArgs>());
            set => _youtubeChannelFilters = value;
        }

        /// <summary>
        /// YouTube channel filter status.
        /// </summary>
        [Input("youtubeChannelStatus")]
        public Input<string>? YoutubeChannelStatus { get; set; }

        public ProfileState()
        {
        }
        public static new ProfileState Empty => new ProfileState();
    }
}
