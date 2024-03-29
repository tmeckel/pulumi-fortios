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
    public static class GetPrefixlist
    {
        /// <summary>
        /// Use this data source to get information on an fortios router prefixlist
        /// </summary>
        public static Task<GetPrefixlistResult> InvokeAsync(GetPrefixlistArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPrefixlistResult>("fortios:router/getPrefixlist:getPrefixlist", args ?? new GetPrefixlistArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on an fortios router prefixlist
        /// </summary>
        public static Output<GetPrefixlistResult> Invoke(GetPrefixlistInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPrefixlistResult>("fortios:router/getPrefixlist:getPrefixlist", args ?? new GetPrefixlistInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPrefixlistArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired router prefixlist.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetPrefixlistArgs()
        {
        }
        public static new GetPrefixlistArgs Empty => new GetPrefixlistArgs();
    }

    public sealed class GetPrefixlistInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired router prefixlist.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetPrefixlistInvokeArgs()
        {
        }
        public static new GetPrefixlistInvokeArgs Empty => new GetPrefixlistInvokeArgs();
    }


    [OutputType]
    public sealed class GetPrefixlistResult
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
        /// IPv4 prefix list rule. The structure of `rule` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetPrefixlistRuleResult> Rules;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetPrefixlistResult(
            string comments,

            string id,

            string name,

            ImmutableArray<Outputs.GetPrefixlistRuleResult> rules,

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
