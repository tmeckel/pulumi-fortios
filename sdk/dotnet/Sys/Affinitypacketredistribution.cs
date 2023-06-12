// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Sys
{
    /// <summary>
    /// Configure packet redistribution.
    /// 
    /// ## Import
    /// 
    /// System AffinityPacketRedistribution can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:sys/affinitypacketredistribution:Affinitypacketredistribution labelname {{fosid}}
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:sys/affinitypacketredistribution:Affinitypacketredistribution labelname {{fosid}}
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:sys/affinitypacketredistribution:Affinitypacketredistribution")]
    public partial class Affinitypacketredistribution : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Affinity setting for VM throughput (64-bit hexadecimal value in the format of 0xxxxxxxxxxxxxxxxx).
        /// </summary>
        [Output("affinityCpumask")]
        public Output<string> AffinityCpumask { get; private set; } = null!;

        /// <summary>
        /// ID of the packet redistribution setting.
        /// </summary>
        [Output("fosid")]
        public Output<int> Fosid { get; private set; } = null!;

        /// <summary>
        /// Physical interface name on which to perform packet redistribution.
        /// </summary>
        [Output("interface")]
        public Output<string> Interface { get; private set; } = null!;

        /// <summary>
        /// ID of the receive queue (when the interface has multiple queues) on which to perform packet redistribution.
        /// </summary>
        [Output("rxqid")]
        public Output<int> Rxqid { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Affinitypacketredistribution resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Affinitypacketredistribution(string name, AffinitypacketredistributionArgs args, CustomResourceOptions? options = null)
            : base("fortios:sys/affinitypacketredistribution:Affinitypacketredistribution", name, args ?? new AffinitypacketredistributionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Affinitypacketredistribution(string name, Input<string> id, AffinitypacketredistributionState? state = null, CustomResourceOptions? options = null)
            : base("fortios:sys/affinitypacketredistribution:Affinitypacketredistribution", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Affinitypacketredistribution resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Affinitypacketredistribution Get(string name, Input<string> id, AffinitypacketredistributionState? state = null, CustomResourceOptions? options = null)
        {
            return new Affinitypacketredistribution(name, id, state, options);
        }
    }

    public sealed class AffinitypacketredistributionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Affinity setting for VM throughput (64-bit hexadecimal value in the format of 0xxxxxxxxxxxxxxxxx).
        /// </summary>
        [Input("affinityCpumask", required: true)]
        public Input<string> AffinityCpumask { get; set; } = null!;

        /// <summary>
        /// ID of the packet redistribution setting.
        /// </summary>
        [Input("fosid", required: true)]
        public Input<int> Fosid { get; set; } = null!;

        /// <summary>
        /// Physical interface name on which to perform packet redistribution.
        /// </summary>
        [Input("interface", required: true)]
        public Input<string> Interface { get; set; } = null!;

        /// <summary>
        /// ID of the receive queue (when the interface has multiple queues) on which to perform packet redistribution.
        /// </summary>
        [Input("rxqid", required: true)]
        public Input<int> Rxqid { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public AffinitypacketredistributionArgs()
        {
        }
        public static new AffinitypacketredistributionArgs Empty => new AffinitypacketredistributionArgs();
    }

    public sealed class AffinitypacketredistributionState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Affinity setting for VM throughput (64-bit hexadecimal value in the format of 0xxxxxxxxxxxxxxxxx).
        /// </summary>
        [Input("affinityCpumask")]
        public Input<string>? AffinityCpumask { get; set; }

        /// <summary>
        /// ID of the packet redistribution setting.
        /// </summary>
        [Input("fosid")]
        public Input<int>? Fosid { get; set; }

        /// <summary>
        /// Physical interface name on which to perform packet redistribution.
        /// </summary>
        [Input("interface")]
        public Input<string>? Interface { get; set; }

        /// <summary>
        /// ID of the receive queue (when the interface has multiple queues) on which to perform packet redistribution.
        /// </summary>
        [Input("rxqid")]
        public Input<int>? Rxqid { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public AffinitypacketredistributionState()
        {
        }
        public static new AffinitypacketredistributionState Empty => new AffinitypacketredistributionState();
    }
}
