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
    /// Configure IPS URL filter cache settings.
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
    ///     var trname = new Fortios.Webfilter.Ipsurlfiltercachesetting("trname", new()
    ///     {
    ///         DnsRetryInterval = 0,
    ///         ExtendedTtl = 0,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Webfilter IpsUrlfilterCacheSetting can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:webfilter/ipsurlfiltercachesetting:Ipsurlfiltercachesetting labelname WebfilterIpsUrlfilterCacheSetting
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:webfilter/ipsurlfiltercachesetting:Ipsurlfiltercachesetting labelname WebfilterIpsUrlfilterCacheSetting
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:webfilter/ipsurlfiltercachesetting:Ipsurlfiltercachesetting")]
    public partial class Ipsurlfiltercachesetting : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Retry interval. Refresh DNS faster than TTL to capture multiple IPs for hosts. 0 means use DNS server's TTL only.
        /// </summary>
        [Output("dnsRetryInterval")]
        public Output<int> DnsRetryInterval { get; private set; } = null!;

        /// <summary>
        /// Extend time to live beyond reported by DNS. 0 means use DNS server's TTL
        /// </summary>
        [Output("extendedTtl")]
        public Output<int> ExtendedTtl { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Ipsurlfiltercachesetting resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Ipsurlfiltercachesetting(string name, IpsurlfiltercachesettingArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:webfilter/ipsurlfiltercachesetting:Ipsurlfiltercachesetting", name, args ?? new IpsurlfiltercachesettingArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Ipsurlfiltercachesetting(string name, Input<string> id, IpsurlfiltercachesettingState? state = null, CustomResourceOptions? options = null)
            : base("fortios:webfilter/ipsurlfiltercachesetting:Ipsurlfiltercachesetting", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Ipsurlfiltercachesetting resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Ipsurlfiltercachesetting Get(string name, Input<string> id, IpsurlfiltercachesettingState? state = null, CustomResourceOptions? options = null)
        {
            return new Ipsurlfiltercachesetting(name, id, state, options);
        }
    }

    public sealed class IpsurlfiltercachesettingArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Retry interval. Refresh DNS faster than TTL to capture multiple IPs for hosts. 0 means use DNS server's TTL only.
        /// </summary>
        [Input("dnsRetryInterval")]
        public Input<int>? DnsRetryInterval { get; set; }

        /// <summary>
        /// Extend time to live beyond reported by DNS. 0 means use DNS server's TTL
        /// </summary>
        [Input("extendedTtl")]
        public Input<int>? ExtendedTtl { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public IpsurlfiltercachesettingArgs()
        {
        }
        public static new IpsurlfiltercachesettingArgs Empty => new IpsurlfiltercachesettingArgs();
    }

    public sealed class IpsurlfiltercachesettingState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Retry interval. Refresh DNS faster than TTL to capture multiple IPs for hosts. 0 means use DNS server's TTL only.
        /// </summary>
        [Input("dnsRetryInterval")]
        public Input<int>? DnsRetryInterval { get; set; }

        /// <summary>
        /// Extend time to live beyond reported by DNS. 0 means use DNS server's TTL
        /// </summary>
        [Input("extendedTtl")]
        public Input<int>? ExtendedTtl { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public IpsurlfiltercachesettingState()
        {
        }
        public static new IpsurlfiltercachesettingState Empty => new IpsurlfiltercachesettingState();
    }
}