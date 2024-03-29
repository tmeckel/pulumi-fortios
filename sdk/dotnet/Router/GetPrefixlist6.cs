// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Router
{
    public static class GetPrefixlist6
    {
        /// <summary>
        /// Use this data source to get information on an fortios router prefixlist6
        /// </summary>
        public static Task<GetPrefixlist6Result> InvokeAsync(GetPrefixlist6Args args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPrefixlist6Result>("fortios:router/getPrefixlist6:getPrefixlist6", args ?? new GetPrefixlist6Args(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on an fortios router prefixlist6
        /// </summary>
        public static Output<GetPrefixlist6Result> Invoke(GetPrefixlist6InvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPrefixlist6Result>("fortios:router/getPrefixlist6:getPrefixlist6", args ?? new GetPrefixlist6InvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPrefixlist6Args : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired router prefixlist6.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetPrefixlist6Args()
        {
        }
        public static new GetPrefixlist6Args Empty => new GetPrefixlist6Args();
    }

    public sealed class GetPrefixlist6InvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired router prefixlist6.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetPrefixlist6InvokeArgs()
        {
        }
        public static new GetPrefixlist6InvokeArgs Empty => new GetPrefixlist6InvokeArgs();
    }


    [OutputType]
    public sealed class GetPrefixlist6Result
    {
        /// <summary>
        /// Comment.
        /// </summary>
        public readonly string Comments;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// IPv6 prefix list rule. The structure of `rule` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetPrefixlist6RuleResult> Rules;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetPrefixlist6Result(
            string comments,

            string id,

            string name,

            ImmutableArray<Outputs.GetPrefixlist6RuleResult> rules,

            string? vdomparam)
        {
            Comments = comments;
            Id = id;
            Name = name;
            Rules = rules;
            Vdomparam = vdomparam;
        }
    }
}
