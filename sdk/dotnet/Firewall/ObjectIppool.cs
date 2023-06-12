// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Firewall
{
    /// <summary>
    /// Provides a resource to configure IPv4 IP address pools of FortiOS.
    /// 
    /// !&gt; **Warning:** The resource will be deprecated and replaced by new resource `fortios.firewall.Ippool`, we recommend that you use the new resource.
    /// 
    /// ## Example Usage
    /// ### Overload Ippool
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Fortios = Pulumiverse.Fortios;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var s1 = new Fortios.Firewall.ObjectIppool("s1", new()
    ///     {
    ///         ArpReply = "enable",
    ///         Comments = "fdsaf",
    ///         Endip = "22.0.0.0",
    ///         Startip = "11.0.0.0",
    ///         Type = "overload",
    ///     });
    /// 
    /// });
    /// ```
    /// ### One-To-One Ippool
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Fortios = Pulumiverse.Fortios;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var s2 = new Fortios.Firewall.ObjectIppool("s2", new()
    ///     {
    ///         ArpReply = "enable",
    ///         Comments = "fdsaf",
    ///         Endip = "222.0.0.0",
    ///         Startip = "121.0.0.0",
    ///         Type = "one-to-one",
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [FortiosResourceType("fortios:firewall/objectIppool:ObjectIppool")]
    public partial class ObjectIppool : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Enable/disable replying to ARP requests when an IP Pool is added to a policy.
        /// </summary>
        [Output("arpReply")]
        public Output<string> ArpReply { get; private set; } = null!;

        /// <summary>
        /// Comment.
        /// </summary>
        [Output("comments")]
        public Output<string?> Comments { get; private set; } = null!;

        /// <summary>
        /// Final IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx).
        /// </summary>
        [Output("endip")]
        public Output<string> Endip { get; private set; } = null!;

        /// <summary>
        /// IP pool name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// First IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx).
        /// </summary>
        [Output("startip")]
        public Output<string> Startip { get; private set; } = null!;

        /// <summary>
        /// IP pool type(Support overload and one-to-one).
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a ObjectIppool resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ObjectIppool(string name, ObjectIppoolArgs args, CustomResourceOptions? options = null)
            : base("fortios:firewall/objectIppool:ObjectIppool", name, args ?? new ObjectIppoolArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ObjectIppool(string name, Input<string> id, ObjectIppoolState? state = null, CustomResourceOptions? options = null)
            : base("fortios:firewall/objectIppool:ObjectIppool", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ObjectIppool resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ObjectIppool Get(string name, Input<string> id, ObjectIppoolState? state = null, CustomResourceOptions? options = null)
        {
            return new ObjectIppool(name, id, state, options);
        }
    }

    public sealed class ObjectIppoolArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable replying to ARP requests when an IP Pool is added to a policy.
        /// </summary>
        [Input("arpReply")]
        public Input<string>? ArpReply { get; set; }

        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comments")]
        public Input<string>? Comments { get; set; }

        /// <summary>
        /// Final IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx).
        /// </summary>
        [Input("endip", required: true)]
        public Input<string> Endip { get; set; } = null!;

        /// <summary>
        /// IP pool name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// First IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx).
        /// </summary>
        [Input("startip", required: true)]
        public Input<string> Startip { get; set; } = null!;

        /// <summary>
        /// IP pool type(Support overload and one-to-one).
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public ObjectIppoolArgs()
        {
        }
        public static new ObjectIppoolArgs Empty => new ObjectIppoolArgs();
    }

    public sealed class ObjectIppoolState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable replying to ARP requests when an IP Pool is added to a policy.
        /// </summary>
        [Input("arpReply")]
        public Input<string>? ArpReply { get; set; }

        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comments")]
        public Input<string>? Comments { get; set; }

        /// <summary>
        /// Final IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx).
        /// </summary>
        [Input("endip")]
        public Input<string>? Endip { get; set; }

        /// <summary>
        /// IP pool name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// First IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx).
        /// </summary>
        [Input("startip")]
        public Input<string>? Startip { get; set; }

        /// <summary>
        /// IP pool type(Support overload and one-to-one).
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public ObjectIppoolState()
        {
        }
        public static new ObjectIppoolState Empty => new ObjectIppoolState();
    }
}
