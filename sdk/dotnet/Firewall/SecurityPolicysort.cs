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
    /// Resource to sort firewall security policies by policyid or policy name, in ascending or descending order.
    /// 
    /// ## Example Usage
    /// ### Example1
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Fortios = Pulumiverse.Fortios;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var test = new Fortios.Firewall.SecurityPolicysort("test", new()
    ///     {
    ///         Sortby = "policyid",
    ///         Sortdirection = "descending",
    ///     });
    /// 
    ///     return new Dictionary&lt;string, object?&gt;
    ///     {
    ///         ["policylistAfterApply"] = test.StatePolicyLists,
    ///     };
    /// });
    /// ```
    /// </summary>
    [FortiosResourceType("fortios:firewall/securityPolicysort:SecurityPolicysort")]
    public partial class SecurityPolicysort : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Comment.
        /// </summary>
        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        /// <summary>
        /// The argument is optional, if it is set, when the value changes, the resource will be re-created. It is usually used when new policies are added, or old policies are deleted, or when the policy name is changed when `sortby` is set to `name`, see Example2.
        /// </summary>
        [Output("forceRecreate")]
        public Output<string?> ForceRecreate { get; private set; } = null!;

        /// <summary>
        /// Sort security policies by the value, it currently supports "policyid" and "name".
        /// </summary>
        [Output("sortby")]
        public Output<string> Sortby { get; private set; } = null!;

        /// <summary>
        /// Sort dirction, supports "ascending" and "descending".
        /// </summary>
        [Output("sortdirection")]
        public Output<string> Sortdirection { get; private set; } = null!;

        [Output("statePolicyLists")]
        public Output<ImmutableArray<Outputs.SecurityPolicysortStatePolicyList>> StatePolicyLists { get; private set; } = null!;

        [Output("status")]
        public Output<string?> Status { get; private set; } = null!;

        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a SecurityPolicysort resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecurityPolicysort(string name, SecurityPolicysortArgs args, CustomResourceOptions? options = null)
            : base("fortios:firewall/securityPolicysort:SecurityPolicysort", name, args ?? new SecurityPolicysortArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecurityPolicysort(string name, Input<string> id, SecurityPolicysortState? state = null, CustomResourceOptions? options = null)
            : base("fortios:firewall/securityPolicysort:SecurityPolicysort", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SecurityPolicysort resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecurityPolicysort Get(string name, Input<string> id, SecurityPolicysortState? state = null, CustomResourceOptions? options = null)
        {
            return new SecurityPolicysort(name, id, state, options);
        }
    }

    public sealed class SecurityPolicysortArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// The argument is optional, if it is set, when the value changes, the resource will be re-created. It is usually used when new policies are added, or old policies are deleted, or when the policy name is changed when `sortby` is set to `name`, see Example2.
        /// </summary>
        [Input("forceRecreate")]
        public Input<string>? ForceRecreate { get; set; }

        /// <summary>
        /// Sort security policies by the value, it currently supports "policyid" and "name".
        /// </summary>
        [Input("sortby", required: true)]
        public Input<string> Sortby { get; set; } = null!;

        /// <summary>
        /// Sort dirction, supports "ascending" and "descending".
        /// </summary>
        [Input("sortdirection", required: true)]
        public Input<string> Sortdirection { get; set; } = null!;

        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SecurityPolicysortArgs()
        {
        }
        public static new SecurityPolicysortArgs Empty => new SecurityPolicysortArgs();
    }

    public sealed class SecurityPolicysortState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// The argument is optional, if it is set, when the value changes, the resource will be re-created. It is usually used when new policies are added, or old policies are deleted, or when the policy name is changed when `sortby` is set to `name`, see Example2.
        /// </summary>
        [Input("forceRecreate")]
        public Input<string>? ForceRecreate { get; set; }

        /// <summary>
        /// Sort security policies by the value, it currently supports "policyid" and "name".
        /// </summary>
        [Input("sortby")]
        public Input<string>? Sortby { get; set; }

        /// <summary>
        /// Sort dirction, supports "ascending" and "descending".
        /// </summary>
        [Input("sortdirection")]
        public Input<string>? Sortdirection { get; set; }

        [Input("statePolicyLists")]
        private InputList<Inputs.SecurityPolicysortStatePolicyListGetArgs>? _statePolicyLists;
        public InputList<Inputs.SecurityPolicysortStatePolicyListGetArgs> StatePolicyLists
        {
            get => _statePolicyLists ?? (_statePolicyLists = new InputList<Inputs.SecurityPolicysortStatePolicyListGetArgs>());
            set => _statePolicyLists = value;
        }

        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SecurityPolicysortState()
        {
        }
        public static new SecurityPolicysortState Empty => new SecurityPolicysortState();
    }
}
